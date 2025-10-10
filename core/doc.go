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
