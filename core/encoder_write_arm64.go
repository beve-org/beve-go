// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

//go:build arm64 && !purego
// +build arm64,!purego

package core

// writeCompressedUintAsm is the arm64 assembly implementation of WriteCompressedUint.
//
//go:noescape
func writeCompressedUintAsm(scratch *[5]byte, n uint64) int

// WriteCompressedUint writes a variable-length encoded unsigned integer using Assembly.
func (e *Encoder) WriteCompressedUint(n uint64) error {
	// Use assembly for the encoding
	length := writeCompressedUintAsm(&e.varintScratch, n)

	// Write the encoded bytes
	return e.WriteBytes(e.varintScratch[:length])
}
