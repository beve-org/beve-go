// Package core provides the core encoding functionality for BEVE (Binary Encoded Values).
//
// BEVE is a high-performance binary serialization format optimized for Go.
// This package contains the internal encoding implementation, organized into
// logical modules for maintainability and clarity.
//
// # Architecture
//
// The core package is organized as follows:
//
//   - buffer.go:             Buffer management, pooling, and growth strategies
//   - encoder_base.go:       Core encoder structure and type dispatch logic
//   - encoder_primitives.go: Encoders for primitive types (int, float, bool, string)
//   - encoder_collections.go: Encoders for collections (slice, map, struct)
//   - encoder_arrays.go:      Encoders for typed arrays (optimized bulk encoding)
//   - encoder_write.go:       Low-level write operations and helpers
//   - encoder_utils.go:       Utility functions (varint, type checks, etc.)
//
// # Performance Characteristics
//
// The encoder is optimized for:
//   - Memory efficiency: Pre-allocated buffers, object pooling
//   - Speed: Unsafe reflection, primitive fast paths, batch encoding
//   - Low allocations: 17-20 allocs for complex payloads
//
// # Thread Safety
//
// Encoders are NOT thread-safe. Each goroutine should get its own encoder
// from the pool. The pool itself is thread-safe.
//
// # Runtime Configuration
//
// BEVE supports runtime configuration via environment variables:
//
//   - BEVE_ENABLE_PREFETCH: Enable software prefetching for SIMD operations (default: false)
//     Software prefetching was tested but showed performance regression on M2 Max due to
//     strong hardware prefetcher. Set to "1", "true", or "yes" to enable for testing.
//
// Example:
//
//	export BEVE_ENABLE_PREFETCH=true
//	go run main.go
//
// # Example Usage
//
//	import "github.com/beve-org/beve-go/core"
//
//	// Internal use only - use beve.Marshal() instead
//	enc := NewEncoder(buf)
//	err := enc.Encode(reflect.ValueOf(data))
//
// For public API, see package beve.
package core

import (
	"os"
	"strconv"
	"strings"
)

// Global runtime configuration flags.
// These are initialized once at package load time from environment variables.
var (
	// EnablePrefetch controls software prefetching in SIMD operations.
	// Default: false (disabled)
	//
	// Phase 2A testing showed that software prefetching causes performance
	// regression on Apple M2 Max due to its strong hardware prefetcher.
	// The CPU already does aggressive prefetching, and software hints add overhead.
	//
	// Results with prefetch enabled:
	//   - Medium payload: 4.9μs → 5.2μs (6% slower)
	//   - Large payload:  49.6μs → 46.1μs (7% faster, but inconsistent)
	//
	// Enable this flag only for:
	//   1. Testing on different CPU architectures
	//   2. Workloads with non-sequential memory access
	//   3. Very large arrays (>1MB) where hardware prefetcher may struggle
	//
	// Set via environment variable: BEVE_ENABLE_PREFETCH=true
	EnablePrefetch = false

	// EnableSIMD is controlled by runtime CPU detection in simd.go
	// This is just documentation - actual flag is in simd.go
	// Set via environment variable: BEVE_DISABLE_SIMD=true to disable
)

func init() {
	// Initialize EnablePrefetch from environment variable
	if val := os.Getenv("BEVE_ENABLE_PREFETCH"); val != "" {
		EnablePrefetch = parseBool(val)
	}
}

// parseBool parses environment variable as boolean.
// Accepts: "1", "true", "yes", "on" (case-insensitive) as true.
func parseBool(val string) bool {
	val = strings.ToLower(strings.TrimSpace(val))
	if val == "1" || val == "true" || val == "yes" || val == "on" {
		return true
	}
	if b, err := strconv.ParseBool(val); err == nil {
		return b
	}
	return false
}
