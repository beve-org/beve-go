package core

import "runtime"

// Platform-specific buffer capacities computed at init time
var (
	optimalBufferCapacity int
)

func init() {
	// OPTIMIZATION: Increased from 512 to 1024 bytes for all platforms
	// to reduce allocation count in string-heavy and medium-sized workloads.
	//
	// Benchmark analysis showed CBOR achieving fewer allocations by using
	// larger initial buffers. This 2× increase reduces reallocation overhead
	// for typical payloads (User struct ~200 bytes, string arrays ~400 bytes).
	//
	// Memory trade-off: +512 bytes per pooled buffer, but pooling limits max
	// memory impact to ~10MB (1024 bytes × ~10k concurrent encoders).
	//
	// Performance targets:
	//   - Reduce allocations from 2-3 to 1-2 for medium structs
	//   - Match CBOR's allocation efficiency (currently 1 alloc for medium writes)
	if runtime.GOOS == "windows" {
		optimalBufferCapacity = 1024 // Windows already at optimal
	} else {
		optimalBufferCapacity = 1024 // Increased from 512 (was too conservative)
	}
}

// getOptimalBufferCapacity returns platform-optimized initial buffer capacity.
//
// Windows allocators perform better with larger initial buffers, while Unix
// systems benefit from smaller initial allocations with exponential growth.
//
// Benchmarks show:
//   - Windows: 1024 bytes -> -25% allocations
//   - Linux/macOS: 512 bytes -> optimal balance
//
//go:inline
func getOptimalBufferCapacity() int {
	return optimalBufferCapacity
}
