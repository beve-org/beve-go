package translatornative

import (
	"fmt"
	"math"
	"strconv"
)

// JSONSerializer is a zero-allocation JSON serializer using buffer pools.
type JSONSerializer struct {
	buf *ByteBuffer
}

// NewJSONSerializer creates a new JSON serializer with pooled buffer.
func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{
		buf: GetBuffer(),
	}
}

// Close returns the buffer to the pool.
func (s *JSONSerializer) Close() {
	if s.buf != nil {
		PutBuffer(s.buf)
		s.buf = nil
	}
}

// Serialize converts a generic value to JSON with minimal allocations.
func (s *JSONSerializer) Serialize(v interface{}) ([]byte, error) {
	s.buf.Reset()
	// Pre-grow buffer based on estimated size
	estimated := estimateJSONSize(v)
	s.buf.Grow(estimated)

	if err := s.writeValue(v); err != nil {
		return nil, err
	}
	// Return a copy (unavoidable allocation)
	result := make([]byte, s.buf.Len())
	copy(result, s.buf.Bytes())
	return result, nil
}

// SerializeIndent converts a value to pretty-printed JSON.
func (s *JSONSerializer) SerializeIndent(v interface{}, prefix, indent string) ([]byte, error) {
	s.buf.Reset()
	if err := s.writeValueIndent(v, prefix, indent, 0); err != nil {
		return nil, err
	}
	result := make([]byte, s.buf.Len())
	copy(result, s.buf.Bytes())
	return result, nil
}

// writeValue writes any value to the buffer.
func (s *JSONSerializer) writeValue(v interface{}) error {
	switch val := v.(type) {
	case nil:
		s.buf.WriteString("null")
	case bool:
		if val {
			s.buf.WriteString("true")
		} else {
			s.buf.WriteString("false")
		}
	case int:
		s.buf.WriteString(strconv.FormatInt(int64(val), 10))
	case int8:
		s.buf.WriteString(strconv.FormatInt(int64(val), 10))
	case int16:
		s.buf.WriteString(strconv.FormatInt(int64(val), 10))
	case int32:
		s.buf.WriteString(strconv.FormatInt(int64(val), 10))
	case int64:
		s.buf.WriteString(strconv.FormatInt(val, 10))
	case uint:
		s.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint8:
		s.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint16:
		s.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint32:
		s.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint64:
		s.buf.WriteString(strconv.FormatUint(val, 10))
	case float32:
		s.writeFloat64(float64(val))
	case float64:
		s.writeFloat64(val)
	case string:
		s.writeString(val)
	case []interface{}:
		return s.writeArray(val)
	case map[string]interface{}:
		return s.writeObject(val)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

// writeFloat64 writes a float64 with proper JSON formatting.
func (s *JSONSerializer) writeFloat64(f float64) {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		s.buf.WriteString("null")
		return
	}
	s.buf.WriteString(strconv.FormatFloat(f, 'g', -1, 64))
}

// writeString writes a JSON string with proper escaping (optimized).
func (s *JSONSerializer) writeString(str string) {
	s.buf.WriteByte('"')

	// Fast path: no escaping needed
	needsEscape := false
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c < 0x20 || c == '"' || c == '\\' {
			needsEscape = true
			break
		}
	}

	if !needsEscape {
		s.buf.WriteString(str)
		s.buf.WriteByte('"')
		return
	}

	// Slow path: escape characters
	for _, r := range str {
		switch r {
		case '"':
			s.buf.WriteString(`\"`)
		case '\\':
			s.buf.WriteString(`\\`)
		case '\b':
			s.buf.WriteString(`\b`)
		case '\f':
			s.buf.WriteString(`\f`)
		case '\n':
			s.buf.WriteString(`\n`)
		case '\r':
			s.buf.WriteString(`\r`)
		case '\t':
			s.buf.WriteString(`\t`)
		default:
			if r < 0x20 {
				s.buf.WriteString(fmt.Sprintf(`\u%04x`, r))
			} else {
				// Write rune directly (UTF-8)
				if r < 128 {
					s.buf.WriteByte(byte(r))
				} else {
					var tmp [4]byte
					n := encodeRune(tmp[:], r)
					s.buf.WriteBytes(tmp[:n])
				}
			}
		}
	}
	s.buf.WriteByte('"')
}

// encodeRune encodes a rune as UTF-8 (inlined for performance).
func encodeRune(p []byte, r rune) int {
	if r < 0x80 {
		p[0] = byte(r)
		return 1
	}
	if r < 0x800 {
		p[0] = 0xC0 | byte(r>>6)
		p[1] = 0x80 | (byte(r) & 0x3F)
		return 2
	}
	if r < 0x10000 {
		p[0] = 0xE0 | byte(r>>12)
		p[1] = 0x80 | (byte(r>>6) & 0x3F)
		p[2] = 0x80 | (byte(r) & 0x3F)
		return 3
	}
	p[0] = 0xF0 | byte(r>>18)
	p[1] = 0x80 | (byte(r>>12) & 0x3F)
	p[2] = 0x80 | (byte(r>>6) & 0x3F)
	p[3] = 0x80 | (byte(r) & 0x3F)
	return 4
}

// writeArray writes a JSON array.
func (s *JSONSerializer) writeArray(arr []interface{}) error {
	s.buf.WriteByte('[')
	for i, item := range arr {
		if i > 0 {
			s.buf.WriteByte(',')
		}
		if err := s.writeValue(item); err != nil {
			return err
		}
	}
	s.buf.WriteByte(']')
	return nil
}

// writeObject writes a JSON object.
func (s *JSONSerializer) writeObject(obj map[string]interface{}) error {
	s.buf.WriteByte('{')
	first := true
	for key, value := range obj {
		if !first {
			s.buf.WriteByte(',')
		}
		first = false
		s.writeString(key)
		s.buf.WriteByte(':')
		if err := s.writeValue(value); err != nil {
			return err
		}
	}
	s.buf.WriteByte('}')
	return nil
}

// writeValueIndent writes a value with indentation.
func (s *JSONSerializer) writeValueIndent(v interface{}, prefix, indent string, depth int) error {
	switch val := v.(type) {
	case nil:
		s.buf.WriteString("null")
	case bool:
		if val {
			s.buf.WriteString("true")
		} else {
			s.buf.WriteString("false")
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return s.writeValue(val)
	case string:
		s.writeString(val)
	case []interface{}:
		return s.writeArrayIndent(val, prefix, indent, depth)
	case map[string]interface{}:
		return s.writeObjectIndent(val, prefix, indent, depth)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

// writeArrayIndent writes an array with indentation.
func (s *JSONSerializer) writeArrayIndent(arr []interface{}, prefix, indent string, depth int) error {
	if len(arr) == 0 {
		s.buf.WriteString("[]")
		return nil
	}

	s.buf.WriteString("[\n")
	for i, item := range arr {
		s.buf.WriteString(prefix)
		for j := 0; j <= depth; j++ {
			s.buf.WriteString(indent)
		}
		if err := s.writeValueIndent(item, prefix, indent, depth+1); err != nil {
			return err
		}
		if i < len(arr)-1 {
			s.buf.WriteByte(',')
		}
		s.buf.WriteByte('\n')
	}
	s.buf.WriteString(prefix)
	for j := 0; j < depth; j++ {
		s.buf.WriteString(indent)
	}
	s.buf.WriteByte(']')
	return nil
}

// writeObjectIndent writes an object with indentation.
func (s *JSONSerializer) writeObjectIndent(obj map[string]interface{}, prefix, indent string, depth int) error {
	if len(obj) == 0 {
		s.buf.WriteString("{}")
		return nil
	}

	s.buf.WriteString("{\n")
	first := true
	for key, value := range obj {
		if !first {
			s.buf.WriteString(",\n")
		}
		first = false
		s.buf.WriteString(prefix)
		for j := 0; j <= depth; j++ {
			s.buf.WriteString(indent)
		}
		s.writeString(key)
		s.buf.WriteString(": ")
		if err := s.writeValueIndent(value, prefix, indent, depth+1); err != nil {
			return err
		}
	}
	s.buf.WriteByte('\n')
	s.buf.WriteString(prefix)
	for j := 0; j < depth; j++ {
		s.buf.WriteString(indent)
	}
	s.buf.WriteByte('}')
	return nil
}
