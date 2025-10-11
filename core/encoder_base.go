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
	
	// High-frequency scratch buffers (40 bytes total)
	uintScratch   [9]byte  // Integer encoding: 1 byte header + 8 bytes max value
	floatBuf      [9]byte  // Float encoding: 1 byte header + 8 bytes IEEE 754
	intBuf        [10]byte // Int encoding: 1 byte header + 9 bytes varint
	varintScratch [5]byte  // Varint encoding
	stringLenBuf  [5]byte  // String length encoding
	single        [1]byte  // Single byte writes
	batchLen      int      // Current batch length (8 bytes on 64-bit)
	
	// = 16 (pointers) + 40 (buffers) + 8 (int) = 64 bytes (one cache line)
	
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
var encoderPool = sync.Pool{
	New: func() interface{} {
		return &Encoder{
			Buf: acquireBuffer(getOptimalBufferCapacity()),
		}
	},
}

// GetEncoderFromPool acquires an encoder from the pool.
func GetEncoderFromPool() *Encoder {
	return encoderPool.Get().(*Encoder)
}

// PutEncoderToPool returns an encoder to the pool for reuse.
//
// Encoders keep buffers up to maxBufferPoolCapacity (1MB) so large payloads
// benefit from pooling without unbounded memory growth.
func PutEncoderToPool(enc *Encoder) {
	// Reset state
	enc.w = nil
	enc.batchLen = 0

	// Only pool encoders with reasonable buffer sizes
	if enc.Buf != nil && cap(enc.Buf.data) <= maxBufferPoolCapacity {
		encoderPool.Put(enc)
	}
	// Large encoders are discarded and will be GC'd
}

// NewEncoder creates a new encoder writing to w.
//
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
func (e *Encoder) Encode(v reflect.Value) error {
	if !v.IsValid() {
		return e.EncodeNull()
	}

	// Get cached encoder function for this type
	// This avoids the type switch on every call - major performance win!
	fn := getEncoderFunc(v.Type())
	return fn(e, v)
}

// DetachBytes transfers ownership of the encoder's buffered bytes to the caller.
//
// The returned lease shares the encoder's underlying storage and must be
// released when no longer needed so the buffer can be recycled. The encoder is
// immediately equipped with a fresh buffer so it can be returned to the pool.
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
	e.Buf = acquireBuffer(reuseCap)
	e.batchLen = 0

	return lease
}

// Common types and functions are in common.go
