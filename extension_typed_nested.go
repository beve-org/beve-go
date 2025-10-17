package beve

import (
	"fmt"
	"reflect"
)

// EncodeTypedNestedArray encodes nested struct arrays (Extension 2)
// This optimizes nested structures by removing redundant field names at all levels
func EncodeTypedNestedArray(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, fmt.Errorf("expected slice or array, got %s", rv.Kind())
	}

	if rv.Len() == 0 {
		return []byte{ExtTypedNestedArray, 0}, nil
	}

	// Build nested schema
	schemas, schemaCount, err := buildNestedSchema(rv.Index(0), MaxNestingDepth)
	if err != nil {
		return nil, fmt.Errorf("failed to build nested schema: %w", err)
	}

	e := getEncoderFromPool()
	defer putEncoderToPool(e)

	// Write header
	e.Buf.WriteByte(ExtTypedNestedArray)

	// Write schema count
	e.Buf.Write(encodeCompressedInt(schemaCount))

	// Write all schemas (depth-first)
	for _, schema := range schemas {
		// Schema ID
		e.Buf.Write(encodeCompressedInt(schema.ID))

		// Parent ID
		e.Buf.Write(encodeCompressedInt(schema.ParentID))

		// Field count
		e.Buf.Write(encodeCompressedInt(len(schema.Fields)))

		// Write fields
		for _, field := range schema.Fields {
			// Field name
			nameBytes := stringToBytes(field.Name)
			e.Buf.Write(encodeCompressedInt(len(nameBytes)))
			e.Buf.Write(nameBytes)

			// Type code
			e.Buf.WriteByte(field.TypeCode)

			// Nested schema ID (if applicable)
			if field.NestedSchemaID >= 0 {
				e.Buf.Write(encodeCompressedInt(field.NestedSchemaID))
			}
		}
	}

	// Write array size
	e.Buf.Write(encodeCompressedInt(rv.Len()))

	// Encode all objects
	for i := 0; i < rv.Len(); i++ {
		elem := rv.Index(i)
		if err := encodeTypedNestedValue(e, elem, schemas, 0); err != nil {
			return nil, fmt.Errorf("failed to encode element %d: %w", i, err)
		}
	}

	return e.Buf.Bytes(), nil
}

// encodeTypedNestedValue encodes a value using nested schema
func encodeTypedNestedValue(e *encoder, v reflect.Value, schemas []SchemaNode, schemaID int) error {
	// Find schema
	var schema *SchemaNode
	for i := range schemas {
		if schemas[i].ID == schemaID {
			schema = &schemas[i]
			break
		}
	}

	if schema == nil {
		return fmt.Errorf("schema %d not found", schemaID)
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			e.Buf.WriteByte(0x00) // null
			return nil
		}
		v = v.Elem()
	}

	// Encode each field according to schema
	for _, field := range schema.Fields {
		var fieldValue reflect.Value

		switch v.Kind() {
		case reflect.Struct:
			fieldValue = getStructFieldByName(v, field.Name)
			if !fieldValue.IsValid() {
				e.Buf.WriteByte(0x00) // null for missing field
				continue
			}

		case reflect.Map:
			key := reflect.ValueOf(field.Name)
			fieldValue = v.MapIndex(key)
			if !fieldValue.IsValid() {
				e.Buf.WriteByte(0x00) // null for missing field
				continue
			}

		default:
			return fmt.Errorf("unexpected type %s", v.Kind())
		}

		// If this field is a nested structure, use nested schema
		if field.NestedSchemaID >= 0 {
			// Check if it's an array of nested structures
			if fieldValue.Kind() == reflect.Slice || fieldValue.Kind() == reflect.Array {
				// Write array size
				e.Buf.Write(encodeCompressedInt(fieldValue.Len()))

				// Encode each element
				for i := 0; i < fieldValue.Len(); i++ {
					elem := fieldValue.Index(i)
					if err := encodeTypedNestedValue(e, elem, schemas, field.NestedSchemaID); err != nil {
						return err
					}
				}
			} else {
				// Single nested object
				if err := encodeTypedNestedValue(e, fieldValue, schemas, field.NestedSchemaID); err != nil {
					return err
				}
			}
		} else {
			// Primitive value - use core encoder
			if err := e.Encode(fieldValue); err != nil {
				return err
			}
		}
	}

	return nil
}

// DecodeTypedNestedArray decodes Extension 2 nested typed arrays
func DecodeTypedNestedArray(data []byte) ([]map[string]interface{}, error) {
	if len(data) < 2 || data[0] != ExtTypedNestedArray {
		return nil, fmt.Errorf("invalid typed nested array header")
	}

	offset := 1

	// Read schema count
	schemaCount, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema count: %w", err)
	}
	offset += bytesRead

	// Read all schemas
	schemas := make([]SchemaNode, schemaCount)
	for i := 0; i < schemaCount; i++ {
		// Read schema ID
		schemaID, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, err
		}
		offset += bytesRead

		// Read parent ID
		parentID, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, err
		}
		offset += bytesRead

		// Read field count
		fieldCount, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, err
		}
		offset += bytesRead

		// Read fields
		fields := make([]FieldSchema, fieldCount)
		for j := 0; j < fieldCount; j++ {
			// Read field name
			nameLen, bytesRead, err := readCompressedSize(data, offset)
			if err != nil {
				return nil, err
			}
			offset += bytesRead

			name := bytesToString(data[offset : offset+nameLen])
			offset += nameLen

			// Read type code
			typeCode := data[offset]
			offset++

			// Read nested schema ID (if applicable)
			nestedSchemaID := -1
			if typeCode == TypeObject || typeCode == TypeArray {
				nestedSchemaID, bytesRead, err = readCompressedSize(data, offset)
				if err != nil {
					return nil, err
				}
				offset += bytesRead
			}

			fields[j] = FieldSchema{
				Name:           name,
				TypeCode:       typeCode,
				NestedSchemaID: nestedSchemaID,
			}
		}

		schemas[i] = SchemaNode{
			ID:       schemaID,
			ParentID: parentID,
			Fields:   fields,
		}
	}

	// Read array size
	arraySize, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to read array size: %w", err)
	}
	offset += bytesRead

	// Decode all objects
	result := make([]map[string]interface{}, arraySize)
	for i := 0; i < arraySize; i++ {
		obj, bytesRead, err := decodeTypedNestedObject(data, offset, schemas, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to decode object %d: %w", i, err)
		}
		result[i] = obj
		offset += bytesRead
	}

	return result, nil
}

// decodeTypedNestedObject decodes a single nested object
func decodeTypedNestedObject(data []byte, offset int, schemas []SchemaNode, schemaID int) (map[string]interface{}, int, error) {
	// Find schema
	var schema *SchemaNode
	for i := range schemas {
		if schemas[i].ID == schemaID {
			schema = &schemas[i]
			break
		}
	}

	if schema == nil {
		return nil, 0, fmt.Errorf("schema %d not found", schemaID)
	}

	result := make(map[string]interface{}, len(schema.Fields))
	startOffset := offset

	// Decode each field
	for _, field := range schema.Fields {
		// Check for null
		if data[offset] == 0x00 {
			result[field.Name] = nil
			offset++
			continue
		}

		// If nested structure
		if field.NestedSchemaID >= 0 {
			// Check if array
			if field.TypeCode == TypeArray {
				// Read array size
				arraySize, bytesRead, err := readCompressedSize(data, offset)
				if err != nil {
					return nil, 0, err
				}
				offset += bytesRead

				// Decode array elements
				array := make([]map[string]interface{}, arraySize)
				for i := 0; i < arraySize; i++ {
					obj, bytesRead, err := decodeTypedNestedObject(data, offset, schemas, field.NestedSchemaID)
					if err != nil {
						return nil, 0, err
					}
					array[i] = obj
					offset += bytesRead
				}

				result[field.Name] = array
			} else {
				// Single nested object
				obj, bytesRead, err := decodeTypedNestedObject(data, offset, schemas, field.NestedSchemaID)
				if err != nil {
					return nil, 0, err
				}
				result[field.Name] = obj
				offset += bytesRead
			}
		} else {
			// Primitive value
			value, bytesRead, err := decodeValueAt(data, offset)
			if err != nil {
				return nil, 0, err
			}
			result[field.Name] = value
			offset += bytesRead
		}
	}

	return result, offset - startOffset, nil
}

// encodeCompressedInt encodes an integer using BEVE compressed format
func encodeCompressedInt(n int) []byte {
	if n < 64 {
		return []byte{byte(n << 2)}
	} else if n < 16384 {
		buf := make([]byte, 2)
		buf[0] = byte((n<<2)|0x01) & 0xFF
		buf[1] = byte(n >> 6)
		return buf
	} else if n < 1073741824 {
		buf := make([]byte, 4)
		buf[0] = byte((n<<2)|0x02) & 0xFF
		buf[1] = byte(n >> 6)
		buf[2] = byte(n >> 14)
		buf[3] = byte(n >> 22)
		return buf
	}

	buf := make([]byte, 8)
	buf[0] = byte((n<<2)|0x03) & 0xFF
	buf[1] = byte(n >> 6)
	buf[2] = byte(n >> 14)
	buf[3] = byte(n >> 22)
	buf[4] = byte(n >> 30)
	buf[5] = byte(n >> 38)
	buf[6] = byte(n >> 46)
	buf[7] = byte(n >> 54)
	return buf
}
