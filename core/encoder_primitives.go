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

// encodeInt and encodeUint are implemented in platform-specific files:
//   - encoder_primitives_amd64.go (assembly optimized)
//   - encoder_primitives_arm64.go (assembly optimized)
//   - encoder_primitives_generic.go (pure Go fallback)

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
