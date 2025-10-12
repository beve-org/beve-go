// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Pure Go implementation of encodeInt/encodeUint for non-assembly platforms.

//go:build (!amd64 && !arm64) || purego
// +build !amd64,!arm64 purego

package core

// encodeInt encodes a signed integer with optimal byte count (Pure Go version).
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

	// Use scratch buffer to batch the write
	e.uintScratch[0] = header
	for j := 0; j < byteCount; j++ {
		e.uintScratch[j+1] = byte(i >> (j * 8))
	}

	return e.WriteBytes(e.uintScratch[:byteCount+1])
}

// encodeUint encodes an unsigned integer with optimal byte count (Pure Go version).
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
