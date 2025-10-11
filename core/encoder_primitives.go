package core

import (
	"encoding/binary"
	"math"
	"reflect"
)

// encodeNull encodes a null value (0x00).
//
// BEVE null encoding:
//
//	Single byte: 0x00
func (e *Encoder) EncodeNull() error {
	return e.WriteByte(0x00)
}

// encodeBool encodes a boolean value.
//
// BEVE boolean encoding:
//
//	false: 0x08
//	true:  0x18
//
// Performance: Inlined for zero-overhead boolean encoding.
//
//go:inline
func (e *Encoder) encodeBool(b bool) error {
	// Inline write: avoid function call overhead
	if e.Buf != nil {
		if b {
			return e.Buf.WriteByte(0x18) // true
		}
		return e.Buf.WriteByte(0x08) // false
	}

	// Fallback for io.Writer
	if b {
		return e.WriteByte(0x18)
	}
	return e.WriteByte(0x08)
}

// encodeInt encodes a signed integer with optimal byte count.
//
// BEVE int encoding uses variable-length encoding based on value range:
//
//	[-128, 127]:              2 bytes (header + 1 byte)
//	[-32768, 32767]:          3 bytes (header + 2 bytes)
//	[-2147483648, 2147483647]: 5 bytes (header + 4 bytes)
//	Otherwise:                9 bytes (header + 8 bytes)
//
// Header format: type=1 | signed=1 | byteCount (2 bits)
//
// Performance: Uses scratch buffer to batch header+value write.
//
//go:inline
func (e *Encoder) encodeInt(i int64) error {
	// Determine optimal byte count for value
	var byteCount int
	var byteCountBits byte

	if i >= -128 && i <= 127 {
		byteCount = 1
		byteCountBits = 0
	} else if i >= -32768 && i <= 32767 {
		byteCount = 2
		byteCountBits = 1
	} else if i >= -2147483648 && i <= 2147483647 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}

	// Construct header: type=1 (number) | mod=1 (signed) | byteCount
	header := byte(0x01) | (1 << 3) | (byteCountBits << 5)

	// Use scratch buffer to batch the write (avoids 2 write calls)
	e.uintScratch[0] = header
	for j := 0; j < byteCount; j++ {
		e.uintScratch[j+1] = byte(i >> (j * 8))
	}

	return e.WriteBytes(e.uintScratch[:byteCount+1])
}

// encodeUint encodes an unsigned integer with optimal byte count.
//
// BEVE uint encoding uses variable-length encoding based on value range:
//
//	[0, 255]:         2 bytes (header + 1 byte)
//	[0, 65535]:       3 bytes (header + 2 bytes)
//	[0, 4294967295]:  5 bytes (header + 4 bytes)
//	Otherwise:        9 bytes (header + 8 bytes)
//
// Header format: type=1 | unsigned=2 | byteCount (2 bits)
//
//go:inline
func (e *Encoder) encodeUint(u uint64) error {
	var byteCount int
	var byteCountBits byte

	if u <= 255 {
		byteCount = 1
		byteCountBits = 0
	} else if u <= 65535 {
		byteCount = 2
		byteCountBits = 1
	} else if u <= 4294967295 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}

	// Construct header: type=1 (number) | mod=2 (unsigned) | byteCount
	header := byte(0x01) | (2 << 3) | (byteCountBits << 5)

	// Batch write using scratch buffer
	e.uintScratch[0] = header
	for j := 0; j < byteCount; j++ {
		e.uintScratch[j+1] = byte(u >> (j * 8))
	}

	return e.WriteBytes(e.uintScratch[:byteCount+1])
}

// encodeFloat encodes a floating point number using IEEE 754 format.
//
// BEVE float encoding:
//
//	Float32: 5 bytes (1 header + 4 bytes IEEE 754)
//	Float64: 9 bytes (1 header + 8 bytes IEEE 754)
//
// Header format:
//
//	Float32: 0x01 | (0 << 3) | (2 << 5) = 0x41
//	Float64: 0x01 | (0 << 3) | (3 << 5) = 0x61
//
// Phase 1 optimization: Uses pre-allocated floatBuf to avoid allocation.
// This reduced float encoding allocations from 2.1M to zero!
//
//go:inline
func (e *Encoder) encodeFloat(f float64, kind reflect.Kind) error {
	if kind == reflect.Float32 {
		val := float32(f)
		uintVal := math.Float32bits(val)
		header := byte(0x01) | (0 << 3) | (2 << 5) // float, 4 bytes

		// Use scratch buffer for float32
		e.uintScratch[0] = header
		binary.LittleEndian.PutUint32(e.uintScratch[1:5], uintVal)
		return e.WriteBytes(e.uintScratch[:5])
	}

	// Float64 path
	uintVal := math.Float64bits(f)
	header := byte(0x01) | (0 << 3) | (3 << 5) // float, 8 bytes

	// Phase 1 optimization: Use pre-allocated floatBuf (NO allocation!)
	e.floatBuf[0] = header
	binary.LittleEndian.PutUint64(e.floatBuf[1:9], uintVal)
	return e.WriteBytes(e.floatBuf[:9])
}

// encodeString encodes a string.
//
// BEVE string encoding:
//
//	1 byte:  header (0x02 = string type)
//	1-5 bytes: compressed uint (string length)
//	N bytes: UTF-8 string data
//
// Total: 2 + len(s) bytes for typical strings
//
// Performance note: String data is written using writeStringBytes()
// which uses unsafe conversion to avoid allocation.
func (e *Encoder) EncodeString(s string) error {
	header := byte(0x02) // string type
	if err := e.WriteByte(header); err != nil {
		return err
	}

	// Write size as compressed unsigned integer
	size := uint64(len(s))
	if err := e.WriteCompressedUint(size); err != nil {
		return err
	}

	return e.WriteStringBytes(s)
}

// encodeRawMessage encodes pre-encoded BEVE data.
//
// RawMessage is used to embed already-encoded BEVE data without
// re-encoding it. This is useful for:
//   - Forwarding pre-encoded messages
//   - Caching encoded values
//   - Building BEVE manually
//
// The data must be valid BEVE format. No validation is performed.
func (e *Encoder) encodeRawMessage(data []byte) error {
	if len(data) == 0 {
		return &UnsupportedError{"RawMessage payload must contain a value"}
	}
	return e.WriteBytes(data)
}

// encodeBinaryMarshaler encodes a type that implements BinaryMarshaler.
//
// BinaryMarshaler allows custom types to provide their own BEVE encoding.
// The MarshalBEVE() method must return valid BEVE format data.
//
// This is checked during encode() dispatch via cached type information
// to avoid repeated interface type assertions.
func (e *Encoder) encodeBinaryMarshaler(m BinaryMarshaler) error {
	data, err := m.MarshalBEVE()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return &UnsupportedError{"BinaryMarshaler returned empty payload"}
	}
	return e.WriteBytes(data)
}
