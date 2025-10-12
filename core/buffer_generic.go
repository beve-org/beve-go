// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Pure Go implementation of Buffer.WriteByte for non-assembly platforms.

//go:build (!amd64 && !arm64) || purego
// +build !amd64,!arm64 purego

package core

// WriteByte appends a single byte to the buffer (Pure Go version).
func (b *Buffer) WriteByte(c byte) error {
	b.data = append(b.data, c)
	return nil
}
