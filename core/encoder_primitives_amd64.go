// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

//go:build amd64 && !purego
// +build amd64,!purego

package core

// encodeUintAsm is the amd64 assembly implementation of encodeUint.
// Returns the number of bytes written to scratch buffer.
//
//go:noescape
func encodeUintAsm(scratch *[9]byte, u uint64) int

// encodeUint encodes an unsigned integer using Assembly optimization.
//
//go:inline
func (e *Encoder) encodeUint(u uint64) error {
	// Use assembly for encoding
	n := encodeUintAsm(&e.uintScratch, u)

	// Write the encoded bytes
	return e.WriteBytes(e.uintScratch[:n])
}

// encodeInt encodes a signed integer (Pure Go for now, TODO: Assembly).
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
