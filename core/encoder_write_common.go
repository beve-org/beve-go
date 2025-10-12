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
