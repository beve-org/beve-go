package core

import (
	"io"
	"reflect"
	"sync"
)

// Encoder handles the encoding of values to BEVE format.
//
// The encoder is the central component of BEVE serialization.
// It maintains state during encoding and delegates to type-specific
// encoding methods based on reflection.
//
// Architecture:
//   - encode() dispatches to type-specific encoders
//   - Pre-allocated scratch buffers reduce allocations
//   - Object pooling amortizes encoder creation cost
//
// Thread Safety:
//   - Encoders are NOT thread-safe
//   - Each goroutine should acquire its own encoder from the pool
//   - The pool itself (encoderPool) IS thread-safe
type Encoder struct {
	// Phase 4: Optimized field layout for cache efficiency
	// Hot path fields (first 64 bytes - one cache line):

	// Most frequently accessed fields (pointers: 8 bytes each)
	Buf *Buffer   // Pre-allocated buffer (pooled) - exported for backward compat
	w   io.Writer // Target writer (may be nil if using Buf)

	// High-frequency scratch buffers (24 bytes total)
	uintScratch   [9]byte // Integer encoding: 1 byte header + 8 bytes max value
	floatBuf      [9]byte // Float encoding: 1 byte header + 8 bytes IEEE 754
	varintScratch [5]byte // Varint encoding
	single        [1]byte // Single byte writes
	batchLen      int     // Current batch length (8 bytes on 64-bit)

	// = 16 (pointers) + 24 (buffers) + 8 (batchLen) = 48 bytes (fits in 1 cache line)

	// Cold path (rarely accessed, second cache line):
	batchBuf [256]byte // Batch buffer for small writes (cold path)
}

// encoderPool is a global pool of encoders for reuse.
//
// Benefits:
//   - Reduces allocations (encoders are ~100 bytes each)
//   - Pre-warmed buffers (512 bytes initial capacity)
//   - Amortizes encoder creation across many Marshal() calls
//
// Hygiene:
//   - Encoders retain buffers up to maxBufferPoolCapacity (1MB)
//   - Encoders are reset before being returned to pool
//
//go:inline
var encoderPool = sync.Pool{
	New: func() interface{} {
		return &Encoder{
			Buf: AcquireBuffer(getOptimalBufferCapacity()),
		}
	},
}

// GetEncoderFromPool acquires an encoder from the pool.
//
//go:inline
func GetEncoderFromPool() *Encoder {
	return encoderPool.Get().(*Encoder)
}

// PutEncoderToPool returns an encoder to the pool for reuse.
//
// Encoders keep buffers up to maxBufferPoolCapacity (1MB) so large payloads
// benefit from pooling without unbounded memory growth.
//
//go:inline
func PutEncoderToPool(enc *Encoder) {
	if enc == nil || enc.Buf == nil {
		return
	}

	bufCap := cap(enc.Buf.data)
	if bufCap <= maxBufferPoolCapacity {
		enc.Buf.Reset()
		enc.batchLen = 0
		encoderPool.Put(enc)
	} else {
		// Buffer is too large to pool, release it
		ReleaseBuffer(enc.Buf)
	}
} // NewEncoder creates a new encoder writing to w.
// For most use cases, prefer GetEncoderFromPool() instead.
// This function is used when you need a specific io.Writer target.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// encode is the main entry point for encoding a reflect.Value.
//
// This method dispatches to type-specific encoders based on reflection.
// It handles:
//   - Null/invalid values
//   - BinaryMarshaler interface
//   - All Go primitive types
//   - Collections (slice, map, struct)
//   - Pointers and interfaces (dereferencing)
//
// Performance notes:
//   - Encoder functions are cached by type (30-40% speedup!)
//   - Type info is cached to avoid repeated interface checks
//   - Primitive types use fast paths with unsafe extraction
//   - Struct encoding uses cached field accessors
//
//go:inline
func (e *Encoder) Encode(v reflect.Value) error {
	if !v.IsValid() {
		return e.EncodeNull()
	}

	// Get cached encoder function for this type
	// This avoids the type switch on every call - major performance win!
	fn := getEncoderFunc(v.Type())
	return fn(e, v)
}

// EncodeAndDetach encodes the value and returns the bytes directly.
// This is optimized for Marshal() use case and leverages multiple optimization layers.
//
// Optimization layers (tried in order):
//
//	Phase 1.1: Stack encoding (primitives only) - 143ns
//	Phase 1.2: Cached encoding (all structs) - ~250ns target
//	Fallback: Standard reflection path - ~600ns
//
//go:inline
func (e *Encoder) EncodeAndDetach(v reflect.Value) ([]byte, error) {
	if v.Kind() == reflect.Struct && v.IsValid() {
		// Phase 1.1: Try stack encoding first (primitives only, fastest)
		if data, ok := e.tryStackEncode(v); ok {
			return data, nil // 143ns - single allocation
		}

		// Phase 1.2: Try cached encoding (all structs with ≤12 fields)
		cache := getOrBuildEncoderCache(v.Type())
		if cache.fieldCount > 0 && cache.fieldCount <= 12 {
			// Reset buffer for cached encoding
			if e.Buf != nil {
				e.Buf.Reset()
			}

			if e.tryEncodeCached(v, cache) {
				// Success! Copy to result
				encoded := e.Buf.Bytes()
				result := make([]byte, len(encoded))
				copy(result, encoded)
				return result, nil
			}

			// Cache encoding failed, reset buffer
			if e.Buf != nil {
				e.Buf.Reset()
			}
		}
	}

	// Standard encoding path - encode to buffer
	if err := e.Encode(v); err != nil {
		return nil, err
	}

	// Copy buffer to new slice
	encoded := e.Buf.Bytes()
	result := make([]byte, len(encoded))
	copy(result, encoded)
	return result, nil
}

// DetachBytes transfers ownership of the encoder's buffered bytes to the caller.
//
// The returned lease shares the encoder's underlying storage and must be
// released when no longer needed so the buffer can be recycled. The encoder is
// immediately equipped with a fresh buffer so it can be returned to the pool.
//
//go:inline
func (e *Encoder) DetachBytes() BufferLease {
	if e.Buf == nil {
		return BufferLease{}
	}

	buf := e.Buf
	lease := newBufferLease(buf)

	reuseCap := cap(buf.data)
	if reuseCap == 0 {
		reuseCap = defaultBufferCapacity
	} else if reuseCap > maxBufferPoolCapacity {
		reuseCap = maxBufferPoolCapacity
	}
	e.Buf = AcquireBuffer(reuseCap)
	e.batchLen = 0

	return lease
}

// Reset clears the encoder for reuse while keeping the underlying buffer.
// This is useful for benchmarks and batch processing scenarios.
//
//go:inline
func (e *Encoder) Reset() {
	if e.Buf != nil {
		e.Buf.Reset()
	}
	e.batchLen = 0
}

// Grow ensures the encoder's buffer has capacity for at least n more bytes.
// This can improve performance when the final size is known in advance.
//
//go:inline
func (e *Encoder) Grow(n int) {
	if e.Buf != nil {
		e.Buf.Grow(n)
	}
}

// Common types and functions are in common.go
