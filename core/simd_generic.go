//go:build (!amd64 && !arm64) || purego
// +build !amd64,!arm64 purego

package core

import (
	"encoding/binary"
	"unsafe"
)

// encodeInt32ArraySIMD is a fallback for platforms without SIMD support.
// Delegates to scalar implementation.
func (e *Encoder) encodeInt32ArraySIMD(data []int32) error {
	return e.encodeInt32ArrayScalar(data)
}

// encodeInt64ArraySIMD is a fallback for platforms without SIMD support.
// Delegates to scalar implementation.
func (e *Encoder) encodeInt64ArraySIMD(data []int64) error {
	return e.encodeInt64ArrayScalar(data)
}

// encodeFloat32ArraySIMD is a fallback for platforms without SIMD support.
// Delegates to scalar implementation.
func (e *Encoder) encodeFloat32ArraySIMD(data []float32) error {
	return e.encodeFloat32ArrayScalar(data)
}

// encodeFloat64ArraySIMD is a fallback for platforms without SIMD support.
// Delegates to scalar implementation.
func (e *Encoder) encodeFloat64ArraySIMD(data []float64) error {
	return e.encodeFloat64ArrayScalar(data)
}

// Helper functions for scalar operations
func (e *Encoder) writeInt32LE(val int32) error {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(val))
	return e.WriteBytes(buf[:])
}

func (e *Encoder) writeInt64LE(val int64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(val))
	return e.WriteBytes(buf[:])
}

func (e *Encoder) writeFloat32LE(val float32) error {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], *(*uint32)(unsafe.Pointer(&val)))
	return e.WriteBytes(buf[:])
}

func (e *Encoder) writeFloat64LE(val float64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], *(*uint64)(unsafe.Pointer(&val)))
	return e.WriteBytes(buf[:])
}
