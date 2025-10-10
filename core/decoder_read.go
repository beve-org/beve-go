package core

// decoder_read.go - Low-level read operations for decoder

// ReadByte reads a single byte from the input data.
//
// Returns error if we've reached the end of data.
//
//go:inline
func (d *Decoder) ReadByte() (byte, error) {
	if d.Pos >= len(d.Data) {
		return 0, &UnsupportedError{"unexpected end of data"}
	}
	b := d.Data[d.Pos]
	d.Pos++
	return b, nil
}

// ReadBytes reads n bytes from the input data.
//
// Returns error if there aren't enough bytes remaining.
//
//go:inline
func (d *Decoder) ReadBytes(n int) ([]byte, error) {
	if d.Pos+n > len(d.Data) {
		return nil, &UnsupportedError{"unexpected end of data"}
	}
	result := d.Data[d.Pos : d.Pos+n]
	d.Pos += n
	return result, nil
}

// ReadCompressedUint reads a variable-length encoded unsigned integer.
//
// BEVE varint encoding (matches encoder):
//   - Bits 0-1: Size indicator (0=1 byte, 1=2 bytes, 2=3 bytes, 3=8 bytes)
//   - Bits 2-7: Value (or high bits for multi-byte)
//
// Performance: Fast path for small values (most common).
func (d *Decoder) ReadCompressedUint() (uint64, error) {
	b, err := d.ReadByte()
	if err != nil {
		return 0, err
	}

	sizeIndicator := b & 0x03
	value := uint64(b >> 2)

	switch sizeIndicator {
	case 0:
		// 1 byte total - value already extracted
		return value, nil
	case 1:
		// 2 bytes total
		next, err := d.ReadByte()
		if err != nil {
			return 0, err
		}
		value = (value << 8) | uint64(next)
		return value, nil
	case 2:
		// 3 bytes total
		for i := 0; i < 2; i++ {
			next, err := d.ReadByte()
			if err != nil {
				return 0, err
			}
			value = (value << 8) | uint64(next)
		}
		return value, nil
	case 3:
		// 4 bytes total
		for i := 0; i < 3; i++ {
			next, err := d.ReadByte()
			if err != nil {
				return 0, err
			}
			value = (value << 8) | uint64(next)
		}
		return value, nil
	}

	return value, nil
}

// GetByteCount converts byte count bits to actual byte count.
//
// This is used for decoding number types where the header
// contains a byte count field.
//
//go:inline
func (d *Decoder) GetByteCount(bits byte) int {
	switch bits {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 8
	default:
		return 0
	}
}
