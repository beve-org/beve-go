//go:build (!amd64 && !arm64) || purego
// +build !amd64,!arm64 purego

package core

// WriteCompressedUint writes a variable-length encoded unsigned integer (Pure Go version).
//
// BEVE varint encoding:
//
//	[0, 63]:              1 byte  (value << 2)
//	[64, 16383]:          2 bytes (0x01 | (value << 2 for high bits))
//	[16384, 1073741823]:  3 bytes (0x02 | ...)
//	[1073741824, ...]:    4 bytes (0x03 | ...)
//
// The last 2 bits of the first byte indicate the number of additional bytes.
//
// Performance: Fast path for small numbers (<64) which covers most lengths.
// Uses scratch buffer to avoid allocation for multi-byte encoding.
//
//go:inline
func (e *Encoder) WriteCompressedUint(n uint64) error {
	// Fast path for small numbers (most common case)
	// Covers typical string lengths, array sizes, etc.
	if n < 64 {
		return e.WriteByte(byte(n << 2))
	}

	// 2-byte encoding (14-bit value)
	if n < 16384 {
		e.varintScratch[0] = byte(0x01 | ((n >> 8) << 2))
		e.varintScratch[1] = byte(n)
		return e.WriteBytes(e.varintScratch[:2])
	}

	// 3-byte encoding (30-bit value)
	if n < 1073741824 {
		e.varintScratch[0] = byte(0x02 | ((n >> 16) << 2))
		e.varintScratch[1] = byte(n >> 8)
		e.varintScratch[2] = byte(n)
		return e.WriteBytes(e.varintScratch[:3])
	}

	// 4-byte encoding (32-bit value)
	e.varintScratch[0] = byte(0x03 | ((n >> 24) << 2))
	e.varintScratch[1] = byte(n >> 16)
	e.varintScratch[2] = byte(n >> 8)
	e.varintScratch[3] = byte(n)
	return e.WriteBytes(e.varintScratch[:4])
}
