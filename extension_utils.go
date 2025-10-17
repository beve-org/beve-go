package beve

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// extractFieldSchema builds a FieldSchema slice from a struct value
func extractFieldSchema(v reflect.Value) ([]FieldSchema, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", v.Kind())
	}

	t := v.Type()
	var fields []FieldSchema

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get field name from beve tag or use field name
		name := field.Name
		if tag := field.Tag.Get("beve"); tag != "" {
			tagParts := strings.Split(tag, ",")
			if tagParts[0] != "" && tagParts[0] != "-" {
				name = tagParts[0]
			}
		} else if tag := field.Tag.Get("json"); tag != "" {
			tagParts := strings.Split(tag, ",")
			if tagParts[0] != "" && tagParts[0] != "-" {
				name = tagParts[0]
			}
		}

		// Determine type code
		typeCode := inferTypeCode(field.Type)

		fields = append(fields, FieldSchema{
			Name:     name,
			TypeCode: typeCode,
		})
	}

	return fields, nil
}

// extractFieldSchemaFromMap builds a FieldSchema slice from a map
func extractFieldSchemaFromMap(m map[string]interface{}) []FieldSchema {
	var fields []FieldSchema
	keys := make([]string, 0, len(m))

	// Sort keys for deterministic ordering
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := m[key]
		typeCode := inferTypeCodeFromValue(value)

		fields = append(fields, FieldSchema{
			Name:     key,
			TypeCode: typeCode,
		})
	}

	return fields
}

// inferTypeCode determines type code from reflect.Type
func inferTypeCode(t reflect.Type) byte {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return TypeInt
	case reflect.String:
		return TypeString
	case reflect.Float32, reflect.Float64:
		return TypeFloat
	case reflect.Bool:
		return TypeBool
	case reflect.Struct, reflect.Map:
		return TypeObject
	case reflect.Slice, reflect.Array:
		return TypeArray
	default:
		return TypeAny
	}
}

// inferTypeCodeFromValue determines type code from runtime value
func inferTypeCodeFromValue(v interface{}) byte {
	if v == nil {
		return TypeAny
	}

	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return TypeInt
	case string:
		return TypeString
	case float32, float64:
		return TypeFloat
	case bool:
		return TypeBool
	case map[string]interface{}:
		return TypeObject
	case []interface{}:
		return TypeArray
	default:
		return TypeAny
	}
}

// buildNestedSchema recursively builds schema nodes for nested structures
func buildNestedSchema(v reflect.Value, depth int) ([]SchemaNode, int, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	schemas := []SchemaNode{}
	currentSchemaID := 0
	maxDepth := depth

	// Build root schema
	rootFields, err := extractFieldSchemaWithNesting(v, &schemas, &currentSchemaID, depth)
	if err != nil {
		return nil, 0, err
	}

	rootSchema := SchemaNode{
		ID:     currentSchemaID,
		Fields: rootFields,
	}
	schemas = append([]SchemaNode{rootSchema}, schemas...)

	return schemas, maxDepth, nil
}

// extractFieldSchemaWithNesting builds field schema and recursively processes nested objects
func extractFieldSchemaWithNesting(v reflect.Value, schemas *[]SchemaNode, schemaID *int, currentDepth int) ([]FieldSchema, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", v.Kind())
	}

	t := v.Type()
	var fields []FieldSchema

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		name := field.Name
		if tag := field.Tag.Get("beve"); tag != "" {
			tagParts := strings.Split(tag, ",")
			if tagParts[0] != "" && tagParts[0] != "-" {
				name = tagParts[0]
			}
		} else if tag := field.Tag.Get("json"); tag != "" {
			tagParts := strings.Split(tag, ",")
			if tagParts[0] != "" && tagParts[0] != "-" {
				name = tagParts[0]
			}
		}

		typeCode := inferTypeCode(field.Type)
		fieldSchema := FieldSchema{
			Name:     name,
			TypeCode: typeCode,
		}

		// If field is a struct (nested object), create a nested schema
		if typeCode == TypeObject && field.Type.Kind() == reflect.Struct {
			*schemaID++
			nestedSchemaID := *schemaID

			// Get nested struct value (zero value if not available)
			nestedValue := reflect.New(field.Type).Elem()
			if v.IsValid() && i < v.NumField() {
				nestedValue = v.Field(i)
			}

			// Recursively build nested schema
			nestedFields, err := extractFieldSchemaWithNesting(nestedValue, schemas, schemaID, currentDepth+1)
			if err != nil {
				return nil, err
			}

			nestedSchema := SchemaNode{
				ID:     nestedSchemaID,
				Fields: nestedFields,
			}
			*schemas = append(*schemas, nestedSchema)

			fieldSchema.NestedSchemaID = nestedSchemaID
		}

		fields = append(fields, fieldSchema)
	}

	return fields, nil
}

// writeCompressedSize writes a compressed unsigned integer (BEVE SIZE format)
func writeCompressedSize(buf []byte, size int) int {
	if size < 64 {
		// 1 byte: 2-bit indicator (0) + 6-bit value
		buf[0] = byte(size << 2)
		return 1
	} else if size < 16384 {
		// 2 bytes: 2-bit indicator (1) + 14-bit value
		buf[0] = byte((size << 2) | 0x01)
		buf[1] = byte(size >> 6)
		return 2
	} else if size < 1073741824 {
		// 4 bytes: 2-bit indicator (2) + 30-bit value
		binary.LittleEndian.PutUint32(buf, uint32((size<<2)|0x02))
		return 4
	} else {
		// 8 bytes: 2-bit indicator (3) + 62-bit value
		binary.LittleEndian.PutUint64(buf, uint64((size<<2)|0x03))
		return 8
	}
}

// readCompressedSize reads a compressed unsigned integer
func readCompressedSize(data []byte, offset int) (int, int, error) {
	if offset >= len(data) {
		return 0, 0, fmt.Errorf("unexpected end of data")
	}

	indicator := data[offset] & 0x03
	switch indicator {
	case 0: // 1 byte
		size := int(data[offset] >> 2)
		return size, 1, nil
	case 1: // 2 bytes
		if offset+1 >= len(data) {
			return 0, 0, fmt.Errorf("unexpected end of data")
		}
		size := int(data[offset]>>2) | (int(data[offset+1]) << 6)
		return size, 2, nil
	case 2: // 4 bytes
		if offset+3 >= len(data) {
			return 0, 0, fmt.Errorf("unexpected end of data")
		}
		val := binary.LittleEndian.Uint32(data[offset:])
		size := int(val >> 2)
		return size, 4, nil
	case 3: // 8 bytes
		if offset+7 >= len(data) {
			return 0, 0, fmt.Errorf("unexpected end of data")
		}
		val := binary.LittleEndian.Uint64(data[offset:])
		size := int(val >> 2)
		return size, 8, nil
	}

	return 0, 0, fmt.Errorf("invalid size indicator")
}

// isArrayOfStructs checks if value is a slice/array of structs
func isArrayOfStructs(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return false
	}

	if rv.Len() == 0 {
		return false
	}

	elemType := rv.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	return elemType.Kind() == reflect.Struct
}

// isArrayOfMaps checks if value is a slice/array of maps
func isArrayOfMaps(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return false
	}

	if rv.Len() == 0 {
		return false
	}

	elemType := rv.Type().Elem()
	return elemType.Kind() == reflect.Map
}

// arraySize returns the size of an array/slice
func arraySize(v interface{}) int {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return 0
	}

	return rv.Len()
}
