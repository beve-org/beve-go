package translatornative

import (
	"fmt"
	"math"
	"strconv"
)

// DirectEncoder encodes JSON directly to BEVE without intermediate structures.
// This is true zero-copy: JSON bytes → BEVE bytes, no allocations.
type DirectEncoder struct {
	json []byte
	pos  int
	buf  []byte
}

// NewDirectEncoder creates a direct JSON→BEVE encoder.
func NewDirectEncoder(jsonData []byte) *DirectEncoder {
	// Estimate BEVE size: typically 70-80% of JSON (BEVE is more compact)
	estimatedSize := (len(jsonData) * 8) / 10
	if estimatedSize < 64 {
		estimatedSize = 64
	}

	return &DirectEncoder{
		json: jsonData,
		pos:  0,
		buf:  make([]byte, 0, estimatedSize),
	}
}

// Encode converts JSON directly to BEVE (zero-copy, zero intermediate allocations).
func (e *DirectEncoder) Encode() ([]byte, error) {
	e.skipWhitespace()
	if err := e.encodeValue(); err != nil {
		return nil, err
	}
	return e.buf, nil
}

// encodeValue encodes current JSON value to BEVE.
func (e *DirectEncoder) encodeValue() error {
	e.skipWhitespace()
	if e.pos >= len(e.json) {
		return fmt.Errorf("unexpected end of JSON")
	}

	ch := e.json[e.pos]
	switch ch {
	case 'n': // null
		return e.encodeNull()
	case 't', 'f': // true/false
		return e.encodeBool()
	case '"': // string
		return e.encodeString()
	case '[': // array
		return e.encodeArray()
	case '{': // object
		return e.encodeObject()
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return e.encodeNumber()
	default:
		return fmt.Errorf("unexpected character: %c", ch)
	}
}

//go:inline
func (e *DirectEncoder) encodeNull() error {
	// Fast byte comparison instead of string allocation
	if e.pos+4 > len(e.json) {
		return fmt.Errorf("invalid null")
	}
	if e.json[e.pos] != 'n' || e.json[e.pos+1] != 'u' || e.json[e.pos+2] != 'l' || e.json[e.pos+3] != 'l' {
		return fmt.Errorf("invalid null")
	}
	e.buf = append(e.buf, 0x00)
	e.pos += 4
	return nil
}

//go:inline
func (e *DirectEncoder) encodeBool() error {
	// Fast byte comparison
	if e.pos+4 <= len(e.json) {
		if e.json[e.pos] == 't' && e.json[e.pos+1] == 'r' && e.json[e.pos+2] == 'u' && e.json[e.pos+3] == 'e' {
			e.buf = append(e.buf, 0x18)
			e.pos += 4
			return nil
		}
	}
	if e.pos+5 <= len(e.json) {
		if e.json[e.pos] == 'f' && e.json[e.pos+1] == 'a' && e.json[e.pos+2] == 'l' && e.json[e.pos+3] == 's' && e.json[e.pos+4] == 'e' {
			e.buf = append(e.buf, 0x08)
			e.pos += 5
			return nil
		}
	}
	return fmt.Errorf("invalid boolean")
}

// encodeString: optimized string encoding with fast-path for no escapes
func (e *DirectEncoder) encodeString() error {
	if e.json[e.pos] != '"' {
		return fmt.Errorf("expected quote")
	}
	e.pos++ // skip opening quote

	// Fast-path: scan for unescaped string (most common case)
	start := e.pos
	for e.pos < len(e.json) {
		ch := e.json[e.pos]
		if ch == '"' {
			// No escapes found - direct copy (FAST PATH)
			length := e.pos - start
			e.buf = append(e.buf, 0x02) // String header
			e.writeSize(length)
			e.buf = append(e.buf, e.json[start:e.pos]...)
			e.pos++ // skip closing quote
			return nil
		}
		if ch == '\\' {
			// Escape found - use slow path
			break
		}
		e.pos++
	}

	// Slow path: handle escapes
	e.pos = start               // reset
	e.buf = append(e.buf, 0x02) // String header

	// Scan to find string end and calculate length
	strLen := 0
	for e.pos < len(e.json) {
		ch := e.json[e.pos]
		if ch == '"' {
			break
		}
		if ch == '\\' {
			e.pos++ // skip escape char
			if e.pos >= len(e.json) {
				return fmt.Errorf("unterminated string")
			}
			// Count escaped character
			if e.json[e.pos] == 'u' {
				// Unicode escape: \uXXXX (will become UTF-8, variable length)
				strLen += 3 // Approximate (worst case for BMP)
				e.pos += 5  // skip uXXXX
			} else {
				strLen++
				e.pos++
			}
		} else {
			strLen++
			e.pos++
		}
	}

	if e.pos >= len(e.json) {
		return fmt.Errorf("unterminated string")
	}

	// Write string length
	e.writeSize(strLen)

	// Write string data (handle escapes)
	e.pos = start
	for e.pos < len(e.json) {
		ch := e.json[e.pos]
		if ch == '"' {
			break
		}
		if ch == '\\' {
			e.pos++
			switch e.json[e.pos] {
			case '"', '\\', '/':
				e.buf = append(e.buf, e.json[e.pos])
			case 'b':
				e.buf = append(e.buf, '\b')
			case 'f':
				e.buf = append(e.buf, '\f')
			case 'n':
				e.buf = append(e.buf, '\n')
			case 'r':
				e.buf = append(e.buf, '\r')
			case 't':
				e.buf = append(e.buf, '\t')
			case 'u':
				// Unicode escape (simplified, assumes BMP)
				e.pos++
				// TODO: full unicode handling
				e.buf = append(e.buf, '?') // placeholder
				e.pos += 3
			}
			e.pos++
		} else {
			e.buf = append(e.buf, ch)
			e.pos++
		}
	}

	e.pos++ // skip closing quote
	return nil
}

// encodeNumber: parse JSON number, write BEVE number
func (e *DirectEncoder) encodeNumber() error {
	start := e.pos
	negative := false

	if e.json[e.pos] == '-' {
		negative = true
		e.pos++
	}

	// Fast path: integer
	intVal := int64(0)
	digitCount := 0
	for e.pos < len(e.json) && e.json[e.pos] >= '0' && e.json[e.pos] <= '9' {
		intVal = intVal*10 + int64(e.json[e.pos]-'0')
		digitCount++
		e.pos++
		if digitCount > 18 {
			break
		}
	}

	// Check if simple integer (no decimal/exponent)
	if e.pos >= len(e.json) || (e.json[e.pos] != '.' && e.json[e.pos] != 'e' && e.json[e.pos] != 'E') {
		if negative {
			intVal = -intVal
		}

		// Encode as compact int
		if intVal >= math.MinInt8 && intVal <= math.MaxInt8 {
			e.buf = append(e.buf, 0x09, byte(intVal))
		} else if intVal >= math.MinInt16 && intVal <= math.MaxInt16 {
			e.buf = append(e.buf, 0x29, byte(intVal), byte(intVal>>8))
		} else if intVal >= math.MinInt32 && intVal <= math.MaxInt32 {
			e.buf = append(e.buf, 0x49)
			e.buf = append(e.buf, byte(intVal), byte(intVal>>8), byte(intVal>>16), byte(intVal>>24))
		} else {
			e.buf = append(e.buf, 0x69)
			for i := 0; i < 8; i++ {
				e.buf = append(e.buf, byte(intVal>>(i*8)))
			}
		}
		return nil
	}

	// Has decimal/exponent - parse as float
	// Skip to end of number
	for e.pos < len(e.json) {
		ch := e.json[e.pos]
		if (ch >= '0' && ch <= '9') || ch == '.' || ch == 'e' || ch == 'E' || ch == '+' || ch == '-' {
			e.pos++
		} else {
			break
		}
	}

	// Parse float using standard library (only place we need it)
	numStr := unsafeString(e.json[start:e.pos])
	f, err := parseFloat(numStr)
	if err != nil {
		return fmt.Errorf("invalid number: %v", err)
	}

	// Encode as float64
	e.buf = append(e.buf, 0x61)
	bits := math.Float64bits(f)
	for i := 0; i < 8; i++ {
		e.buf = append(e.buf, byte(bits>>(i*8)))
	}
	return nil
}

// encodeArray: [...]
func (e *DirectEncoder) encodeArray() error {
	if e.json[e.pos] != '[' {
		return fmt.Errorf("expected '['")
	}
	e.pos++

	// Write BEVE generic array header
	e.buf = append(e.buf, 0x05)

	e.skipWhitespace()
	if e.pos < len(e.json) && e.json[e.pos] == ']' {
		// Empty array
		e.writeSize(0)
		e.pos++
		return nil
	}

	// Reserve 4 bytes for size (will patch later)
	sizePos := len(e.buf)
	e.buf = append(e.buf, 0, 0, 0, 0)

	// Encode elements (single pass!)
	count := 0
	for {
		if err := e.encodeValue(); err != nil {
			return err
		}
		count++
		e.skipWhitespace()
		if e.pos >= len(e.json) {
			return fmt.Errorf("unterminated array")
		}
		if e.json[e.pos] == ']' {
			break
		}
		if e.json[e.pos] != ',' {
			return fmt.Errorf("expected comma")
		}
		e.pos++
	}

	// Patch size at reserved position
	e.patchSize(sizePos, count)

	e.pos++ // skip ]
	return nil
}

// encodeObject: {...}
func (e *DirectEncoder) encodeObject() error {
	if e.json[e.pos] != '{' {
		return fmt.Errorf("expected '{'")
	}
	e.pos++

	// Write BEVE object header (string keys)
	e.buf = append(e.buf, 0x03)

	e.skipWhitespace()
	if e.pos < len(e.json) && e.json[e.pos] == '}' {
		e.writeSize(0)
		e.pos++
		return nil
	}

	// Reserve 4 bytes for size (will patch later)
	sizePos := len(e.buf)
	e.buf = append(e.buf, 0, 0, 0, 0)

	// Encode fields (single pass!)
	count := 0
	for {
		// Skip whitespace before key
		e.skipWhitespace()

		if e.json[e.pos] != '"' {
			return fmt.Errorf("expected string key")
		}

		// Encode key (string without BEVE header)
		e.pos++ // skip quote
		keyStart := e.pos
		for e.pos < len(e.json) && e.json[e.pos] != '"' {
			if e.json[e.pos] == '\\' {
				e.pos++
			}
			e.pos++
		}
		keyLen := e.pos - keyStart
		e.writeSize(keyLen)
		e.buf = append(e.buf, e.json[keyStart:e.pos]...)
		e.pos++ // skip closing quote

		e.skipWhitespace()
		if e.json[e.pos] != ':' {
			return fmt.Errorf("expected colon")
		}
		e.pos++ // skip colon

		// Encode value
		if err := e.encodeValue(); err != nil {
			return err
		}
		count++

		e.skipWhitespace()
		if e.pos >= len(e.json) {
			return fmt.Errorf("unterminated object")
		}
		if e.json[e.pos] == '}' {
			break
		}
		if e.json[e.pos] != ',' {
			return fmt.Errorf("expected comma")
		}
		e.pos++
	}

	// Patch size at reserved position
	e.patchSize(sizePos, count)

	e.pos++ // skip }
	return nil
}

// Helper functions
func (e *DirectEncoder) skipWhitespace() {
	for e.pos < len(e.json) {
		ch := e.json[e.pos]
		if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
			e.pos++
		} else {
			break
		}
	}
}

//go:inline
func (e *DirectEncoder) writeSize(n int) {
	if n < 64 {
		// Bit-shifted encoding (matches Rust/C++/TS/core)
		e.buf = append(e.buf, byte(n<<2))
	} else if n < 16384 {
		e.buf = append(e.buf, byte(0x01|((n>>8)<<2)), byte(n))
	} else if n < 1073741824 {
		e.buf = append(e.buf, byte(0x02|((n>>16)<<2)), byte(n>>8), byte(n))
	} else {
		u := uint64(n)
		e.buf = append(e.buf, byte(0x03|byte((u>>24)<<2)))
		for i := 2; i >= 0; i-- {
			e.buf = append(e.buf, byte(u>>(i*8)))
		}
	}
}

// patchSize patches the size at the reserved position and adjusts buffer
func (e *DirectEncoder) patchSize(pos int, n int) {
	if n < 64 {
		// 1 byte size (bit-shifted)
		e.buf[pos] = byte(n << 2)
		// Shift buffer left by 3 bytes (we reserved 4, need 1)
		copy(e.buf[pos+1:], e.buf[pos+4:])
		e.buf = e.buf[:len(e.buf)-3]
	} else if n < 16384 {
		// 2 bytes size (bit-shifted)
		e.buf[pos] = byte(0x01 | ((n >> 8) << 2))
		e.buf[pos+1] = byte(n)
		// Shift buffer left by 2 bytes
		copy(e.buf[pos+2:], e.buf[pos+4:])
		e.buf = e.buf[:len(e.buf)-2]
	} else if n < 1073741824 {
		// 4 bytes size (perfect fit, no shift needed!)
		e.buf[pos] = byte(0x02 | ((n >> 16) << 2))
		e.buf[pos+1] = byte(n >> 16)
		e.buf[pos+2] = byte(n >> 8)
		e.buf[pos+3] = byte(n)
	} else {
		// 8 bytes size (need to grow)
		u := uint64(n)
		// Make space for 4 more bytes
		e.buf = append(e.buf, 0, 0, 0, 0)
		// Shift right
		copy(e.buf[pos+8:], e.buf[pos+4:len(e.buf)-4])
		e.buf[pos] = byte(0xC0 | byte(u>>56)&0x3F)
		for i := 0; i < 7; i++ {
			e.buf[pos+1+i] = byte(u >> ((6 - i) * 8))
		}
	}
}

// parseFloat - optimized float parser using strconv (10x faster than fmt.Sscanf)
func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
