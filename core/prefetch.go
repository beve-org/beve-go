//go:build !purego
// +build !purego

package core

import (
	"unsafe"
)

// prefetch.go - Software prefetch hints for cache optimization
//
// Phase 2A: SIMD + Prefetching hybrid optimization
// Target: Reduce cache miss rate by 30-40% for Medium/Large payloads
//
// Strategy:
//   - Prefetch next data chunk while encoding current chunk
//   - Optimal prefetch distance: 64-128 bytes ahead (1-2 cache lines)
//   - M2 Max ARM64: L1D 64KB, L2 4MB, cache line 128 bytes
//
// Expected Performance Gains:
//   - Medium payload: 4.9μs → 3.0μs (1.6× faster)
//   - Large payload: 49.6μs → 30μs (1.65× faster)
//   - Cache miss rate: 30-40% reduction

// Prefetch distance constants optimized for M2 Max
const (
	// ARM64 M2 Max: 128-byte cache line
	// Prefetch 1-2 cache lines ahead for optimal L1D hit rate
	prefetchDistanceBytes = 128 // 128 bytes = 1 cache line

	// Element-based prefetch distances
	prefetchDistanceInt32   = 32 // 32 elements = 128 bytes
	prefetchDistanceInt64   = 16 // 16 elements = 128 bytes
	prefetchDistanceFloat32 = 32 // 32 elements = 128 bytes
	prefetchDistanceFloat64 = 16 // 16 elements = 128 bytes
)

// prefetchRead issues a software prefetch hint to load data into L1 cache.
//
// Platform-specific implementations:
//   - ARM64: PRFM PLDL1KEEP instruction
//   - AMD64: PREFETCHT0 instruction
//   - Generic: No-op (compiler may optimize)
//
// SAFETY: This is a hint, not a requirement. CPU may ignore if busy.
//
// Implementation Note: This function is implemented in assembly for ARM64 and AMD64.
// See prefetch_arm64.s and prefetch_amd64.s for platform-specific implementations.
func prefetchRead(addr unsafe.Pointer, len int)

// prefetchInt32Array prefetches next chunk of int32 array.
//
// Usage:
//
//	for i := 0; i < len(data); i += batchSize {
//	    prefetchInt32Array(data, i, batchSize)
//	    // Process data[i:i+batchSize]
//	}
func prefetchInt32Array(data []int32, offset int, batchSize int) {
	if offset+prefetchDistanceInt32 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceInt32])
		// Prefetch 128 bytes (32 elements × 4 bytes)
		prefetchRead(ptr, 128)
	}
}

// prefetchInt64Array prefetches next chunk of int64 array.
func prefetchInt64Array(data []int64, offset int, batchSize int) {
	if offset+prefetchDistanceInt64 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceInt64])
		// Prefetch 128 bytes (16 elements × 8 bytes)
		prefetchRead(ptr, 128)
	}
}

// prefetchFloat32Array prefetches next chunk of float32 array.
func prefetchFloat32Array(data []float32, offset int, batchSize int) {
	if offset+prefetchDistanceFloat32 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceFloat32])
		prefetchRead(ptr, 128)
	}
}

// prefetchFloat64Array prefetches next chunk of float64 array.
func prefetchFloat64Array(data []float64, offset int, batchSize int) {
	if offset+prefetchDistanceFloat64 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceFloat64])
		prefetchRead(ptr, 128)
	}
}

// prefetchUint32Array prefetches next chunk of uint32 array.
func prefetchUint32Array(data []uint32, offset int, batchSize int) {
	if offset+prefetchDistanceInt32 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceInt32])
		prefetchRead(ptr, 128)
	}
}

// prefetchUint64Array prefetches next chunk of uint64 array.
func prefetchUint64Array(data []uint64, offset int, batchSize int) {
	if offset+prefetchDistanceInt64 < len(data) {
		ptr := unsafe.Pointer(&data[offset+prefetchDistanceInt64])
		prefetchRead(ptr, 128)
	}
}
