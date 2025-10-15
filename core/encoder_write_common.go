// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Unified write functions for all platforms (Pure Go implementation).
//
// Phase 11 Migration: Replaced assembly with pure Go based on benchmark data
// showing 6.2% overall improvement and 35% improvement in string-heavy workloads.

package core

import (
	"io"
	"unsafe"
)

// Common write functions used by all platforms

// WriteByte writes a single byte to the encoder's output.
//
// Fast path: If using a pooled buffer, writes directly.
// Slow path: If using an io.Writer, checks for io.ByteWriter interface.
func (e *Encoder) WriteByte(b byte) error {
	// Fast path: write directly to pooled buffer
	if e.Buf != nil {
		return e.Buf.WriteByte(b)
	}

	// Slow path: write to io.Writer
	if bw, ok := e.w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}

	// Fallback: create single-byte slice using scratch buffer
	e.single[0] = b
	_, err := e.w.Write(e.single[:])
	return err
}

// WriteBytes writes a byte slice to the encoder's output.
//
//go:inline
func (e *Encoder) WriteBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	// Fast path: write directly to pooled buffer
	if e.Buf != nil {
		_, err := e.Buf.Write(data)
		return err
	}

	// Slow path: write to io.Writer
	_, err := e.w.Write(data)
	return err
}

// WriteStringBytes writes a string as bytes to the encoder's output.
//
//go:inline
func (e *Encoder) WriteStringBytes(s string) error {
	if len(s) == 0 {
		return nil
	}

	// Fast path: write directly to pooled buffer
	if e.Buf != nil {
		b := stringToBytes(s)
		_, err := e.Buf.Write(b)
		return err
	}

	// Slow path: write to io.Writer
	if sw, ok := e.w.(io.StringWriter); ok {
		_, err := sw.WriteString(s)
		return err
	}

	// Fallback: use zero-copy conversion
	_, err := e.w.Write(stringToBytes(s))
	return err
}

// stringToBytes converts string to []byte without allocation.
//
//go:inline
func stringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// writeCompressedUintPure is the pure Go implementation of compressed uint encoding.
// This is used by WriteCompressedUint for edge cases (n >= 64).
//
// PERFORMANCE: Pure Go version (migrated from assembly in Phase 11).
// End-to-end benchmarks showed pure Go is 6.2% faster overall.
//
//go:inline
func writeCompressedUintPure(scratch *[5]byte, n uint64) int {
	// Fast path: n < 64 (should be handled at caller level)
	// This is here for completeness
	if n < 64 {
		scratch[0] = byte(n << 2)
		return 1
	}

	// Two byte encoding: n < 16384
	if n < 16384 {
		scratch[0] = byte((n>>8)<<2) | 0x01
		scratch[1] = byte(n)
		return 2
	}

	// Three byte encoding: n < 1073741824 (2^30)
	if n < 1073741824 {
		scratch[0] = byte((n>>16)<<2) | 0x02
		scratch[1] = byte(n >> 8)
		scratch[2] = byte(n)
		return 3
	}

	// Four byte encoding: n >= 1073741824
	scratch[0] = byte((n>>24)<<2) | 0x03
	scratch[1] = byte(n >> 16)
	scratch[2] = byte(n >> 8)
	scratch[3] = byte(n)
	return 4
}

// WriteCompressedUint writes a variable-length encoded unsigned integer.
//
// PERFORMANCE CRITICAL: Second hotspot in Phase 11 profiling.
//
// Format:
//   - 1 byte:  0-63        → [vv vv vv vv vv vv 00]
//   - 2 bytes: 64-16383    → [vv vv vv vv vv vv 01] [vvvvvvvv]
//   - 3 bytes: 16384-1B    → [vv vv vv vv vv vv 10] [vvvvvvvv] [vvvvvvvv]
//   - 4 bytes: 1B+         → [vv vv vv vv vv vv 11] [vvvvvvvv] [vvvvvvvv] [vvvvvvvv]
//
// WriteCompressedUint is implemented in encoder_write.go
// This file only contains WriteByte and WriteBytes
