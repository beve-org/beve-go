//go:build (!amd64 && !arm64) || purego

package core

import (
	"io"
	"unsafe"
)

// Fallback write functions for non-AMD64/ARM64 architectures or purego builds.
//
// This file provides basic implementations when the optimized encoder_write_common.go
// cannot be used (e.g., 386, MIPS, RISC-V, or when -tags=purego is specified).

// WriteByte writes a single byte to the encoder's output.
//
//go:inline
func (e *Encoder) WriteByte(b byte) error {
	if e.Buf != nil {
		return e.Buf.WriteByte(b)
	}

	if bw, ok := e.w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}

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

	if e.Buf != nil {
		_, err := e.Buf.Write(data)
		return err
	}

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

	if e.Buf != nil {
		b := stringToBytes(s)
		_, err := e.Buf.Write(b)
		return err
	}

	if sw, ok := e.w.(io.StringWriter); ok {
		_, err := sw.WriteString(s)
		return err
	}

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

// WriteCompressedUint writes a variable-length encoded unsigned integer (Fallback version).
//
// BEVE varint encoding:
//
//	[0, 63]:              1 byte  (value << 2)
//	[64, 16383]:          2 bytes (0x01 | (value << 2 for high bits))
//	[16384, 1073741823]:  4 bytes (0x02 | ...)
//	[1073741824, ...]:    8 bytes (0x03 | ...)
//
// The last 2 bits of the first byte indicate the number of additional bytes.
//
// Performance: Fast path for small numbers (<64) which covers most lengths.
// Uses scratch buffer to avoid allocation for multi-byte encoding.
//
//go:inline
func (e *Encoder) WriteCompressedUint(n uint64) error {
	// Fast path for small numbers (most common case)
	// Covers typical string lengths, array sizes, etc.
	if n < 64 {
		return e.WriteByte(byte(n << 2))
	}

	// 2-byte encoding (14-bit value)
	if n < 16384 {
		e.varintScratch[0] = byte(0x01 | ((n >> 8) << 2))
		e.varintScratch[1] = byte(n)
		return e.WriteBytes(e.varintScratch[:2])
	}

	// 3-byte encoding (30-bit value)
	if n < 1073741824 {
		e.varintScratch[0] = byte(0x02 | ((n >> 16) << 2))
		e.varintScratch[1] = byte(n >> 8)
		e.varintScratch[2] = byte(n)
		return e.WriteBytes(e.varintScratch[:3])
	}

	// 4-byte encoding (32-bit value)
	e.varintScratch[0] = byte(0x03 | ((n >> 24) << 2))
	e.varintScratch[1] = byte(n >> 16)
	e.varintScratch[2] = byte(n >> 8)
	e.varintScratch[3] = byte(n)
	return e.WriteBytes(e.varintScratch[:4])
}
