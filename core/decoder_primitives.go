package core

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

// decoder_primitives.go - Decoders for primitive types

// DecodeNumber decodes a numeric value based on the header.
//
// BEVE number header format:
//   - Bits 0-2: Type category (1 = number)
//   - Bits 3-4: Number type (0=float, 1=signed int, 2=unsigned int)
//   - Bits 5-7: Byte count bits
func (d *Decoder) DecodeNumber(v reflect.Value, header byte) error {
	typeBits := (header >> 3) & 0x03
	byteCountBits := (header >> 5) & 0x07

	byteCount := d.GetByteCount(byteCountBits)

	switch typeBits {
	case 0: // float
		return d.DecodeFloat(v, byteCount)
	case 1: // signed int
		return d.DecodeInt(v, byteCount)
	case 2: // unsigned int
		return d.DecodeUint(v, byteCount)
	}
	return &UnsupportedError{"invalid number type"}
}

// DecodeFloat decodes a floating-point number.
//
// Supported formats:
//   - 2 bytes: bfloat16 (approximated as float32)
//   - 4 bytes: float32 (IEEE 754)
//   - 8 bytes: float64 (IEEE 754)
func (d *Decoder) DecodeFloat(v reflect.Value, byteCount int) error {
	var value float64

	switch byteCount {
	case 1: // bfloat16 - 2 bytes, approximated
		data, err := d.ReadBytes(2)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint16(data)
		value = float64(math.Float32frombits(uint32(uintVal) << 16))
	case 4: // float32
		data, err := d.ReadBytes(4)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint32(data)
		value = float64(math.Float32frombits(uintVal))
	case 8: // float64
		data, err := d.ReadBytes(8)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint64(data)
		value = math.Float64frombits(uintVal)
	default:
		return &UnsupportedError{"unsupported float size"}
	}

	return setFloatValue(v, value)
}

// DecodeInt decodes a signed integer.
//
// Supports 1, 2, 4, or 8 byte integers with sign extension.
func (d *Decoder) DecodeInt(v reflect.Value, byteCount int) error {
	data, err := d.ReadBytes(byteCount)
	if err != nil {
		return err
	}

	var val int64
	for i, b := range data {
		val |= int64(b) << (i * 8)
	}

	// Sign extend if necessary
	if byteCount < 8 && (val&(1<<((byteCount*8)-1))) != 0 {
		val |= -1 << (byteCount * 8)
	}

	return setIntValue(v, val)
}

// DecodeUint decodes an unsigned integer.
//
// Supports 1, 2, 4, or 8 byte unsigned integers.
func (d *Decoder) DecodeUint(v reflect.Value, byteCount int) error {
	data, err := d.ReadBytes(byteCount)
	if err != nil {
		return err
	}

	var val uint64
	for i, b := range data {
		val |= uint64(b) << (i * 8)
	}

	return setUintValue(v, val)
}

// DecodeString decodes a string value.
//
// BEVE string encoding:
//   - Compressed uint: string length
//   - N bytes: UTF-8 string data
//
// Uses zero-copy conversion for performance.
//
//go:inline
func (d *Decoder) DecodeString(v reflect.Value) error {
	size, err := d.ReadCompressedUint()
	if err != nil {
		return err
	}

	data, err := d.ReadBytes(int(size))
	if err != nil {
		return err
	}

	// Zero-copy string conversion
	str := bytesToString(data)
	return setStringValue(v, str)
}

// bytesToString converts []byte to string without allocation.
//
// SAFETY: This is safe because:
//  1. The data slice comes from decoder's immutable input
//  2. The string lifetime doesn't outlive the decoder
//  3. Go strings are immutable, so no modification risk
//
// This optimization avoids allocating a copy of the string data.
//
//go:inline
func bytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}
