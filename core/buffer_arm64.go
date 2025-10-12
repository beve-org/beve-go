// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

//go:build arm64 && !purego
// +build arm64,!purego

package core

// writeByteAsm is the arm64 assembly implementation of Buffer.WriteByte.
//
//go:noescape
func writeByteAsm(b *Buffer, c byte) bool

// WriteByte appends a single byte to the buffer using Assembly fast path.
func (b *Buffer) WriteByte(c byte) error {
	// Try assembly fast path first
	if writeByteAsm(b, c) {
		return nil
	}

	// Slow path: buffer needs growth
	b.data = append(b.data, c)
	return nil
}
