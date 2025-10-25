package translatornative

import (
	"encoding/binary"
	"fmt"
	"math"
)

// BEVEDecoder decodes BEVE binary format to generic values.
// Zero dependencies on beve-go package.
type BEVEDecoder struct {
	data []byte
	pos  int
}

// NewBEVEDecoder creates a new BEVE decoder.
func NewBEVEDecoder(data []byte) *BEVEDecoder {
	return &BEVEDecoder{
		data: data,
		pos:  0,
	}
}

// Decode decodes BEVE binary to a generic value.
func (d *BEVEDecoder) Decode() (interface{}, error) {
	if d.pos >= len(d.data) {
		return nil, fmt.Errorf("unexpected end of BEVE data")
	}
	return d.decodeValue()
}

// decodeValue decodes any BEVE value.
func (d *BEVEDecoder) decodeValue() (interface{}, error) {
	if d.pos >= len(d.data) {
		return nil, fmt.Errorf("unexpected end of data at position %d", d.pos)
	}

	header := d.data[d.pos]
	d.pos++

	// Extract type from first 3 bits
	typeTag := header & 0x07

	switch typeTag {
	case 0x00: // null or boolean
		return d.decodeNullOrBool(header)
	case 0x01: // number
		return d.decodeNumber(header)
	case 0x02: // string
		return d.decodeString()
	case 0x03: // object
		return d.decodeObject(header)
	case 0x04: // typed array
		return d.decodeTypedArray(header)
	case 0x05: // generic array
		return d.decodeGenericArray()
	case 0x06: // extension
		return nil, fmt.Errorf("extensions not supported in translator-native")
	default:
		return nil, fmt.Errorf("unknown type tag: 0x%02x", typeTag)
	}
}

// decodeNullOrBool decodes null or boolean.
// null:  0b00000'000 = 0x00
// false: 0b000'01'000 = 0x08
// true:  0b000'11'000 = 0x18
func (d *BEVEDecoder) decodeNullOrBool(header byte) (interface{}, error) {
	if header == 0x00 {
		return nil, nil // null
	}
	if header == 0x08 {
		return false, nil
	}
	if header == 0x18 {
		return true, nil
	}
	return nil, fmt.Errorf("invalid null/bool header: 0x%02x", header)
}

// decodeNumber decodes a number (float or int).
// Bits 3-4: subtype (00=float, 01=signed, 10=unsigned)
// Bits 5-7: byte count indicator
func (d *BEVEDecoder) decodeNumber(header byte) (interface{}, error) {
	subtype := (header >> 3) & 0x03      // bits 3-4
	byteCountIdx := (header >> 5) & 0x07 // bits 5-7

	// Calculate actual byte count from indicator
	byteCount := 1 << byteCountIdx // 1, 2, 4, 8, 16, etc.

	if d.pos+byteCount > len(d.data) {
		return nil, fmt.Errorf("not enough data for number (%d bytes needed)", byteCount)
	}

	switch subtype {
	case 0x00: // floating point
		return d.decodeFloat(byteCount)
	case 0x01: // signed integer
		return d.decodeSignedInt(byteCount)
	case 0x02: // unsigned integer
		return d.decodeUnsignedInt(byteCount)
	default:
		return nil, fmt.Errorf("invalid number subtype: 0x%02x", subtype)
	}
}

// decodeFloat decodes floating point number.
func (d *BEVEDecoder) decodeFloat(byteCount int) (float64, error) {
	switch byteCount {
	case 2: // bfloat16 (not fully supported, approximate as float32)
		// Read as uint16 then convert (simplified)
		if d.pos+2 > len(d.data) {
			return 0, fmt.Errorf("not enough data for bfloat16")
		}
		// For now, treat as regular float16 approximation
		bits := binary.LittleEndian.Uint16(d.data[d.pos:])
		d.pos += 2
		// Convert to float32 approximation (shift left by 16 bits)
		f32bits := uint32(bits) << 16
		return float64(math.Float32frombits(f32bits)), nil

	case 4: // float32
		bits := binary.LittleEndian.Uint32(d.data[d.pos:])
		d.pos += 4
		return float64(math.Float32frombits(bits)), nil

	case 8: // float64
		bits := binary.LittleEndian.Uint64(d.data[d.pos:])
		d.pos += 8
		return math.Float64frombits(bits), nil

	default:
		return 0, fmt.Errorf("unsupported float byte count: %d", byteCount)
	}
}

// decodeSignedInt decodes signed integer.
func (d *BEVEDecoder) decodeSignedInt(byteCount int) (float64, error) {
	switch byteCount {
	case 1:
		val := int8(d.data[d.pos])
		d.pos++
		return float64(val), nil
	case 2:
		val := int16(binary.LittleEndian.Uint16(d.data[d.pos:]))
		d.pos += 2
		return float64(val), nil
	case 4:
		val := int32(binary.LittleEndian.Uint32(d.data[d.pos:]))
		d.pos += 4
		return float64(val), nil
	case 8:
		val := int64(binary.LittleEndian.Uint64(d.data[d.pos:]))
		d.pos += 8
		return float64(val), nil
	default:
		return 0, fmt.Errorf("unsupported int byte count: %d", byteCount)
	}
}

// decodeUnsignedInt decodes unsigned integer.
func (d *BEVEDecoder) decodeUnsignedInt(byteCount int) (float64, error) {
	switch byteCount {
	case 1:
		val := d.data[d.pos]
		d.pos++
		return float64(val), nil
	case 2:
		val := binary.LittleEndian.Uint16(d.data[d.pos:])
		d.pos += 2
		return float64(val), nil
	case 4:
		val := binary.LittleEndian.Uint32(d.data[d.pos:])
		d.pos += 4
		return float64(val), nil
	case 8:
		val := binary.LittleEndian.Uint64(d.data[d.pos:])
		d.pos += 8
		return float64(val), nil
	default:
		return 0, fmt.Errorf("unsupported uint byte count: %d", byteCount)
	}
}

// decodeString decodes UTF-8 string.
// Layout: (HEADER already consumed) | SIZE | DATA
func (d *BEVEDecoder) decodeString() (string, error) {
	size, err := d.decodeSize()
	if err != nil {
		return "", err
	}

	if d.pos+size > len(d.data) {
		return "", fmt.Errorf("not enough data for string (%d bytes needed)", size)
	}

	str := string(d.data[d.pos : d.pos+size])
	d.pos += size
	return str, nil
}

// decodeObject decodes object (map with string keys).
// Header bits 3-4: key type (00=string, 01=signed int, 10=unsigned int)
// Layout: (HEADER consumed) | SIZE | KEY[0] | VALUE[0] | ... KEY[N] | VALUE[N]
func (d *BEVEDecoder) decodeObject(header byte) (map[string]interface{}, error) {
	keyType := (header >> 3) & 0x03

	// Only support string keys for now (JSON compatible)
	if keyType != 0x00 {
		return nil, fmt.Errorf("only string keys supported, got key type: 0x%02x", keyType)
	}

	size, err := d.decodeSize()
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, size)

	for i := 0; i < size; i++ {
		// Decode key (string without header)
		keySize, err := d.decodeSize()
		if err != nil {
			return nil, err
		}

		if d.pos+keySize > len(d.data) {
			return nil, fmt.Errorf("not enough data for object key")
		}

		key := string(d.data[d.pos : d.pos+keySize])
		d.pos += keySize

		// Decode value (with header)
		value, err := d.decodeValue()
		if err != nil {
			return nil, err
		}

		result[key] = value
	}

	return result, nil
}

// decodeTypedArray decodes typed array.
// For simplicity, convert to generic array for JSON compatibility.
func (d *BEVEDecoder) decodeTypedArray(header byte) ([]interface{}, error) {
	dataType := (header >> 3) & 0x03     // bits 3-4
	byteCountIdx := (header >> 5) & 0x07 // bits 5-7

	size, err := d.decodeSize()
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0, size)

	switch dataType {
	case 0x00: // floating point array
		byteCount := 1 << byteCountIdx
		for i := 0; i < size; i++ {
			val, err := d.decodeFloat(byteCount)
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}

	case 0x01: // signed int array
		byteCount := 1 << byteCountIdx
		for i := 0; i < size; i++ {
			val, err := d.decodeSignedInt(byteCount)
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}

	case 0x02: // unsigned int array
		byteCount := 1 << byteCountIdx
		for i := 0; i < size; i++ {
			val, err := d.decodeUnsignedInt(byteCount)
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}

	case 0x03: // boolean or string array
		boolOrString := (header >> 6) & 0x01
		if boolOrString == 0 { // boolean array
			// Booleans packed as bits
			byteCount := (size + 7) / 8 // Round up to nearest byte
			if d.pos+byteCount > len(d.data) {
				return nil, fmt.Errorf("not enough data for boolean array")
			}

			for i := 0; i < size; i++ {
				byteIdx := i / 8
				bitIdx := i % 8
				bit := (d.data[d.pos+byteIdx] >> bitIdx) & 0x01
				result = append(result, bit == 1)
			}
			d.pos += byteCount
		} else { // string array
			for i := 0; i < size; i++ {
				str, err := d.decodeString()
				if err != nil {
					return nil, err
				}
				result = append(result, str)
			}
		}

	default:
		return nil, fmt.Errorf("invalid typed array data type: 0x%02x", dataType)
	}

	return result, nil
}

// decodeGenericArray decodes generic array.
// Layout: (HEADER consumed) | SIZE | VALUE[0] | ... VALUE[N]
func (d *BEVEDecoder) decodeGenericArray() ([]interface{}, error) {
	size, err := d.decodeSize()
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0, size)

	for i := 0; i < size; i++ {
		value, err := d.decodeValue()
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}

// decodeSize decodes a compressed unsigned integer.
// First 2 bits indicate byte count, remaining bits hold value.
func (d *BEVEDecoder) decodeSize() (int, error) {
	if d.pos >= len(d.data) {
		return 0, fmt.Errorf("unexpected end of data while reading size")
	}

	b0 := d.data[d.pos]
	indicator := b0 & 0x03 // Lower 2 bits (bit-shifted format)

	switch indicator {
	case 0x00: // 1 byte: bits 2-7 contain value
		d.pos++
		return int(b0 >> 2), nil

	case 0x01: // 2 bytes: bits 2-7 of b0 + full b1
		if d.pos+1 >= len(d.data) {
			return 0, fmt.Errorf("not enough data for 2-byte size")
		}
		high := int(b0>>2) << 8
		low := int(d.data[d.pos+1])
		d.pos += 2
		return high | low, nil

	case 0x02: // 4 bytes: bits 2-7 of b0 + b1-b3
		if d.pos+3 >= len(d.data) {
			return 0, fmt.Errorf("not enough data for 4-byte size")
		}
		val := int(b0>>2) << 24
		val |= int(d.data[d.pos+1]) << 16
		val |= int(d.data[d.pos+2]) << 8
		val |= int(d.data[d.pos+3])
		d.pos += 4
		return val, nil

	case 0x03: // 8 bytes: bits 2-7 of b0 + b1-b7
		if d.pos+7 >= len(d.data) {
			return 0, fmt.Errorf("not enough data for 8-byte size")
		}
		val := uint64(b0>>2) << 56
		for i := 1; i < 8; i++ {
			val |= uint64(d.data[d.pos+i]) << uint(56-i*8)
		}
		d.pos += 8
		return int(val), nil

	default:
		return 0, fmt.Errorf("invalid size indicator: 0x%02x", indicator)
	}
}
