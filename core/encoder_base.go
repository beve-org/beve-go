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
	w   io.Writer // Target writer (may be nil if using Buf)
	Buf *Buffer   // Pre-allocated buffer (pooled) - exported for backward compat

	// Scratch buffers to avoid allocations during encoding
	// These are reused across multiple encode operations
	single        [1]byte   // Single byte writes
	uintScratch   [9]byte   // Integer encoding: 1 byte header + 8 bytes max value
	varintScratch [5]byte   // Varint encoding
	batchBuf      [256]byte // Batch buffer for small writes
	batchLen      int       // Current batch length

	// Phase 1 optimization: Pre-allocated buffers to eliminate allocations
	// These reduced float encoding allocations from 2.1M to near-zero!
	floatBuf     [9]byte  // Float encoding: 1 byte header + 8 bytes IEEE 754
	intBuf       [10]byte // Int encoding: 1 byte header + 9 bytes varint
	stringLenBuf [5]byte  // String length encoding
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

// estimateSize provides a size hint for pre-allocating buffers.
//
// This avoids multiple buffer reallocations during encoding, which is
// particularly important on Windows where memory allocation is slower.
//
// Estimates are conservative (slight over-allocation) to minimize growth.
//
//go:inline
func estimateSize(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Struct:
		// Small struct: 256 bytes covers most cases
		// Large struct: 512 bytes prevents initial reallocation
		numField := v.NumField()
		if numField <= 5 {
			return 256
		}
		return 512
	case reflect.Slice, reflect.Array:
		len := v.Len()
		if len == 0 {
			return 16
		}
		// Estimate 8 bytes per element + overhead
		return len*8 + 32
	case reflect.Map:
		len := v.Len()
		if len == 0 {
			return 16
		}
		// Estimate 16 bytes per key-value pair + overhead
		return len*16 + 32
	case reflect.String:
		return v.Len() + 8
	default:
		return 64
	}
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
//   - Buffer pre-growth based on size hints (Windows optimization)
func (e *Encoder) Encode(v reflect.Value) error {
	if !v.IsValid() {
		return e.EncodeNull()
	}

	// Pre-grow buffer based on estimated size
	// This significantly reduces allocations, especially on Windows
	if e.Buf != nil {
		hint := estimateSize(v)
		if hint > 0 {
			e.Buf.Grow(hint)
		}
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
