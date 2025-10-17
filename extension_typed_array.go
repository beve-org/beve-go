package beve

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// EncodeTypedArray encodes an array of objects with shared schema (Extension 1)
// This eliminates field name repetition, saving ~48% space for large arrays
func EncodeTypedArray(v interface{}) ([]byte, error) {
	e := getEncoderFromPool()
	defer putEncoderToPool(e)

	if err := encodeTypedArrayToEncoder(e, v); err != nil {
		return nil, err
	}

	return e.Buf.Bytes(), nil
}

// encodeTypedArrayToEncoder is the internal implementation
func encodeTypedArrayToEncoder(e *encoder, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return fmt.Errorf("typed array requires slice or array, got %v", rv.Kind())
	}

	arrayLen := rv.Len()
	if arrayLen == 0 {
		// Empty array - use generic encoding
		buf := e.Buf
		buf.WriteByte(0x85) // Generic array header
		sizeBuf := make([]byte, 8)
		sizeBytes := writeCompressedSize(sizeBuf, 0)
		buf.Write(sizeBuf[:sizeBytes])
		return nil
	}

	// Extract schema from first element
	firstElem := rv.Index(0)
	var schema []FieldSchema
	var err error

	if firstElem.Kind() == reflect.Struct || (firstElem.Kind() == reflect.Ptr && firstElem.Elem().Kind() == reflect.Struct) {
		schema, err = extractFieldSchema(firstElem)
		if err != nil {
			return err
		}
	} else if firstElem.Kind() == reflect.Map && firstElem.Type().Key().Kind() == reflect.String {
		// Map with string keys
		m, ok := firstElem.Interface().(map[string]interface{})
		if !ok {
			return fmt.Errorf("unsupported map type for typed array")
		}
		schema = extractFieldSchemaFromMap(m)
	} else {
		return fmt.Errorf("typed array requires structs or string-keyed maps, got %v", firstElem.Kind())
	}

	// Write header
	e.Buf.WriteByte(ExtTypedArray)

	// Write field count
	sizeBuf := make([]byte, 8)
	sizeBytes := writeCompressedSize(sizeBuf, len(schema))
	e.Buf.Write(sizeBuf[:sizeBytes])

	// Write schema (field names once)
	for _, field := range schema {
		// Write field name as string (SIZE + UTF-8 bytes)
		nameBuf := []byte(field.Name)
		nameSize := len(nameBuf)

		nameSizeBuf := make([]byte, 8)
		nameSizeBytes := writeCompressedSize(nameSizeBuf, nameSize)
		e.Buf.Write(nameSizeBuf[:nameSizeBytes])
		e.Buf.Write(nameBuf)
	}

	// Write array size
	arraySizeBuf := make([]byte, 8)
	arraySizeBytes := writeCompressedSize(arraySizeBuf, arrayLen)
	e.Buf.Write(arraySizeBuf[:arraySizeBytes])

	// Write objects (values only, no keys!)
	for i := 0; i < arrayLen; i++ {
		elem := rv.Index(i)

		if elem.Kind() == reflect.Struct || (elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct) {
			// Encode struct fields in schema order
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			for _, field := range schema {
				fieldValue := getStructFieldByName(elem, field.Name)
				if err := e.Encode(fieldValue); err != nil {
					return fmt.Errorf("encode field %s: %w", field.Name, err)
				}
			}
		} else if elem.Kind() == reflect.Map {
			// Encode map values in schema order
			m := elem.Interface().(map[string]interface{})

			for _, field := range schema {
				value, ok := m[field.Name]
				if !ok {
					// Field not present - encode null
					e.Buf.WriteByte(0x00)
					continue
				}

				if err := e.Encode(reflect.ValueOf(value)); err != nil {
					return fmt.Errorf("encode field %s: %w", field.Name, err)
				}
			}
		}
	}

	return nil
}

// getStructFieldByName finds a struct field by BEVE/JSON tag name
func getStructFieldByName(v reflect.Value, name string) reflect.Value {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Check beve tag
		if tag := field.Tag.Get("beve"); tag != "" {
			if tag == name {
				return v.Field(i)
			}
		}

		// Check json tag
		if tag := field.Tag.Get("json"); tag != "" {
			if tag == name {
				return v.Field(i)
			}
		}

		// Check field name
		if field.Name == name {
			return v.Field(i)
		}
	}

	return reflect.Value{}
}

// DecodeTypedArray decodes Extension 1 typed object array
func DecodeTypedArray(data []byte) ([]map[string]interface{}, error) {
	if len(data) < 2 || data[0] != ExtTypedArray {
		return nil, fmt.Errorf("invalid typed array header")
	}

	offset := 1

	// Read field count
	fieldCount, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, fmt.Errorf("read field count: %w", err)
	}
	offset += bytesRead

	// Read schema (field names)
	schema := make([]string, fieldCount)
	for i := 0; i < fieldCount; i++ {
		// Read field name size
		nameSize, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, fmt.Errorf("read field name size: %w", err)
		}
		offset += bytesRead

		// Read field name
		if offset+nameSize > len(data) {
			return nil, fmt.Errorf("unexpected end of data reading field name")
		}
		schema[i] = string(data[offset : offset+nameSize])
		offset += nameSize
	}

	// Read array size
	arraySize, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, fmt.Errorf("read array size: %w", err)
	}
	offset += bytesRead

	// Decode objects (values only)
	objects := make([]map[string]interface{}, arraySize)

	for i := 0; i < arraySize; i++ {
		obj := make(map[string]interface{}, fieldCount)

		for _, fieldName := range schema {
			// Decode value
			value, bytesRead, err := decodeValueAt(data, offset)
			if err != nil {
				return nil, fmt.Errorf("decode field %s: %w", fieldName, err)
			}

			obj[fieldName] = value
			offset += bytesRead
		}

		objects[i] = obj
	}

	return objects, nil
}

// decodeValueAt decodes a single BEVE value starting at offset
func decodeValueAt(data []byte, offset int) (interface{}, int, error) {
	if offset >= len(data) {
		return nil, 0, fmt.Errorf("unexpected end of data")
	}

	header := data[offset]
	typeCode := header & 0x07 // First 3 bits

	switch typeCode {
	case 0x00: // null or boolean
		if header == 0x00 {
			return nil, 1, nil
		}
		if header&0x08 != 0 {
			// Boolean
			return header&0x10 != 0, 1, nil
		}
		return nil, 1, nil

	case 0x01: // number
		return decodeNumberAt(data, offset)

	case 0x02: // string
		return decodeStringAt(data, offset)

	case 0x03: // object
		return decodeObjectAt(data, offset)

	case 0x05: // generic array
		return decodeArrayAt(data, offset)

	default:
		return nil, 0, fmt.Errorf("unsupported type code: 0x%02x", typeCode)
	}
}

// decodeNumberAt decodes a number starting at offset
func decodeNumberAt(data []byte, offset int) (interface{}, int, error) {
	if offset >= len(data) {
		return nil, 0, fmt.Errorf("unexpected end of data")
	}

	header := data[offset]
	numType := (header >> 3) & 0x03         // Bits 3-4
	byteCount := 1 << ((header >> 5) & 0x07) // Bits 5-7

	if offset+1+byteCount > len(data) {
		return nil, 0, fmt.Errorf("unexpected end of data reading number")
	}

	switch numType {
	case 0: // float
		switch byteCount {
		case 4:
			bits := binary.LittleEndian.Uint32(data[offset+1:])
			return float64(math.Float32frombits(bits)), 1 + byteCount, nil
		case 8:
			bits := binary.LittleEndian.Uint64(data[offset+1:])
			return math.Float64frombits(bits), 1 + byteCount, nil
		}

	case 1: // signed int
		switch byteCount {
		case 1:
			return int64(int8(data[offset+1])), 2, nil
		case 2:
			return int64(int16(binary.LittleEndian.Uint16(data[offset+1:]))), 3, nil
		case 4:
			return int64(int32(binary.LittleEndian.Uint32(data[offset+1:]))), 5, nil
		case 8:
			return int64(binary.LittleEndian.Uint64(data[offset+1:])), 9, nil
		}

	case 2: // unsigned int
		switch byteCount {
		case 1:
			return uint64(data[offset+1]), 2, nil
		case 2:
			return uint64(binary.LittleEndian.Uint16(data[offset+1:])), 3, nil
		case 4:
			return uint64(binary.LittleEndian.Uint32(data[offset+1:])), 5, nil
		case 8:
			return binary.LittleEndian.Uint64(data[offset+1:]), 9, nil
		}
	}

	return nil, 0, fmt.Errorf("unsupported number format")
}

// decodeStringAt decodes a string starting at offset
func decodeStringAt(data []byte, offset int) (string, int, error) {
	if offset >= len(data) {
		return "", 0, fmt.Errorf("unexpected end of data")
	}

	offset++ // Skip header

	// Read string size
	size, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return "", 0, err
	}
	offset += bytesRead

	if offset+size > len(data) {
		return "", 0, fmt.Errorf("unexpected end of data reading string")
	}

	str := string(data[offset : offset+size])
	return str, 1 + bytesRead + size, nil
}

// decodeObjectAt decodes an object starting at offset
func decodeObjectAt(data []byte, offset int) (map[string]interface{}, int, error) {
	startOffset := offset

	if offset >= len(data) {
		return nil, 0, fmt.Errorf("unexpected end of data")
	}

	offset++ // Skip header

	// Read object size
	size, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, 0, err
	}
	offset += bytesRead

	obj := make(map[string]interface{}, size)

	for i := 0; i < size; i++ {
		// Read key (string without header)
		keySize, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, 0, err
		}
		offset += bytesRead

		if offset+keySize > len(data) {
			return nil, 0, fmt.Errorf("unexpected end of data reading key")
		}

		key := string(data[offset : offset+keySize])
		offset += keySize

		// Read value
		value, bytesRead, err := decodeValueAt(data, offset)
		if err != nil {
			return nil, 0, err
		}

		obj[key] = value
		offset += bytesRead
	}

	return obj, offset - startOffset, nil
}

// decodeArrayAt decodes an array starting at offset
func decodeArrayAt(data []byte, offset int) ([]interface{}, int, error) {
	startOffset := offset

	if offset >= len(data) {
		return nil, 0, fmt.Errorf("unexpected end of data")
	}

	offset++ // Skip header

	// Read array size
	size, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, 0, err
	}
	offset += bytesRead

	arr := make([]interface{}, size)

	for i := 0; i < size; i++ {
		value, bytesRead, err := decodeValueAt(data, offset)
		if err != nil {
			return nil, 0, err
		}

		arr[i] = value
		offset += bytesRead
	}

	return arr, offset - startOffset, nil
}

// Helper functions for float conversion
func float32frombits(b uint32) float32 {
	return *(*float32)(unsafe.Pointer(&b))
}

func float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}
