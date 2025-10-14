package core

import "reflect"

// EncodeIntFast encodes a signed integer without using reflection.
//
//go:inline
func EncodeIntFast(enc *Encoder, v int64) error {
	return enc.encodeInt(v)
}

// EncodeUintFast encodes an unsigned integer without using reflection.
//
//go:inline
func EncodeUintFast(enc *Encoder, v uint64) error {
	return enc.encodeUint(v)
}

// EncodeBoolFast encodes a boolean without using reflection.
//
//go:inline
func EncodeBoolFast(enc *Encoder, v bool) error {
	return enc.encodeBool(v)
}

// EncodeStringFast encodes a string without using reflection.
//
//go:inline
func EncodeStringFast(enc *Encoder, s string) error {
	return enc.EncodeString(s)
}

// EncodeFloat32Fast encodes a float32 without using reflection.
//
//go:inline
func EncodeFloat32Fast(enc *Encoder, v float32) error {
	return enc.encodeFloat(float64(v), reflect.Float32)
}

// EncodeFloat64Fast encodes a float64 without using reflection.
//
//go:inline
func EncodeFloat64Fast(enc *Encoder, v float64) error {
	return enc.encodeFloat(v, reflect.Float64)
}

// EncodeBytesFast encodes a raw byte slice using the typed array fast path.
//
//go:inline
func EncodeBytesFast(enc *Encoder, b []byte) error {
	return enc.encodeUint8SliceDirect(b)
}

// EncodeStringSliceFast encodes a []string using the typed array fast path.
//
//go:inline
func EncodeStringSliceFast(enc *Encoder, slice []string) error {
	return enc.encodeStringSliceDirect(slice)
}

// EncodeBoolSliceFast encodes a []bool using the typed array fast path.
//
//go:inline
func EncodeBoolSliceFast(enc *Encoder, slice []bool) error {
	return enc.encodeBoolSliceDirect(slice)
}

// EncodeInt8SliceFast encodes a []int8 using the typed array fast path.
//
//go:inline
func EncodeInt8SliceFast(enc *Encoder, slice []int8) error {
	return enc.encodeInt8SliceDirect(slice)
}

// EncodeInt16SliceFast encodes a []int16 using the typed array fast path.
//
//go:inline
func EncodeInt16SliceFast(enc *Encoder, slice []int16) error {
	return enc.encodeInt16SliceDirect(slice)
}

// EncodeInt32SliceFast encodes a []int32 using the typed array fast path.
//
//go:inline
func EncodeInt32SliceFast(enc *Encoder, slice []int32) error {
	return enc.encodeInt32SliceDirect(slice)
}

// EncodeInt64SliceFast encodes a []int64 using the typed array fast path.
//
//go:inline
func EncodeInt64SliceFast(enc *Encoder, slice []int64) error {
	return enc.encodeInt64SliceDirect(slice)
}

// EncodeUint8SliceFast encodes a []uint8 using the typed array fast path.
//
//go:inline
func EncodeUint8SliceFast(enc *Encoder, slice []uint8) error {
	return enc.encodeUint8SliceDirect(slice)
}

// EncodeUint16SliceFast encodes a []uint16 using the typed array fast path.
//
//go:inline
func EncodeUint16SliceFast(enc *Encoder, slice []uint16) error {
	return enc.encodeUint16SliceDirect(slice)
}

// EncodeUint32SliceFast encodes a []uint32 using the typed array fast path.
//
//go:inline
func EncodeUint32SliceFast(enc *Encoder, slice []uint32) error {
	return enc.encodeUint32SliceDirect(slice)
}

// EncodeUint64SliceFast encodes a []uint64 using the typed array fast path.
//
//go:inline
func EncodeUint64SliceFast(enc *Encoder, slice []uint64) error {
	return enc.encodeUint64SliceDirect(slice)
}

// EncodeFloat32SliceFast encodes a []float32 using the typed array fast path.
//
//go:inline
func EncodeFloat32SliceFast(enc *Encoder, slice []float32) error {
	return enc.encodeFloat32SliceDirect(slice)
}

// EncodeFloat64SliceFast encodes a []float64 using the typed array fast path.
//
//go:inline
func EncodeFloat64SliceFast(enc *Encoder, slice []float64) error {
	return enc.encodeFloat64SliceDirect(slice)
}
