package core

import (
	"io"
	"unsafe"
)

// writeByte writes a single byte to the encoder's output.
//
// Fast path: If using a pooled buffer, writes directly.
// Slow path: If using an io.Writer, checks for io.ByteWriter interface.
//
// This optimization avoids slice allocation for single-byte writes.
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

// writeBytes writes a byte slice to the encoder's output.
//
// Fast path: If using a pooled buffer, writes directly.
// Slow path: If using an io.Writer, writes through io.Writer interface.
//
// Performance: Inline hint helps compiler optimize for small slices.
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

// writeStringBytes writes a string as bytes to the encoder's output.
//
// This uses unsafe conversion (via stringToBytes) to avoid allocating
// a copy of the string data. This is safe because:
//  1. The string data is immediately written
//  2. The data is not retained after the write
//  3. The string cannot be modified during the write
//
// Fast path: Uses pooled buffer and zero-copy conversion.
// Slow path: Tries io.StringWriter interface, falls back to byte conversion.
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
	// Safe because data is immediately written and not retained
	_, err := e.w.Write(stringToBytes(s))
	return err
}

// writeCompressedUint writes a variable-length encoded unsigned integer.
//
// BEVE varint encoding:
//
//	[0, 63]:              1 byte  (value << 2)
//	[64, 16383]:          2 bytes (0x01 | (value << 2 for high bits))
//	[16384, 1073741823]:  3 bytes (0x02 | ...)
//	[1073741824, ...]:    4 bytes (0x03 | ...)
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

// stringToBytes converts string to []byte without allocation.
//
// SAFETY: The returned slice must not be modified and should not outlive the string.
// This is safe in our encoding context because:
//  1. The data is immediately written to the output
//  2. The slice is not retained after the write
//  3. Strings are immutable in Go
//
// This optimization avoids allocating a copy of string data,
// which can save significant memory for large strings.
//
//go:inline
func stringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
