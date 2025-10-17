package beve

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"sort"
)

// FieldIndex represents an indexed field in Extension 0
type FieldIndex struct {
	Name   string
	Offset uint32 // Offset from start of data section
	Size   uint16 // Size of the value in bytes
	Flags  byte   // Type flags for quick access
}

// EncodeIndexedObject encodes an object with field index (Extension 0)
// This allows O(1) field access by name
func EncodeIndexedObject(obj map[string]interface{}) ([]byte, error) {
	if len(obj) == 0 {
		return []byte{ExtFieldIndex, TypeObject, 0}, nil
	}

	// Sort keys for deterministic encoding
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Encode each value first to calculate offsets
	valueBuffers := make([][]byte, len(keys))
	totalValueSize := 0

	e := getEncoderFromPool()
	defer putEncoderToPool(e)

	for i, key := range keys {
		e.Buf.Reset()
		if err := e.Encode(reflect.ValueOf(obj[key])); err != nil {
			return nil, fmt.Errorf("failed to encode field %q: %w", key, err)
		}
		valueBuffers[i] = make([]byte, e.Buf.Len())
		copy(valueBuffers[i], e.Buf.Bytes())
		totalValueSize += len(valueBuffers[i])
	}

	// Build index table
	indices := make([]FieldIndex, len(keys))
	currentOffset := uint32(0)

	for i, key := range keys {
		valueType := inferTypeFromBytes(valueBuffers[i])

		indices[i] = FieldIndex{
			Name:   key,
			Offset: currentOffset,
			Size:   uint16(len(valueBuffers[i])),
			Flags:  valueType,
		}

		currentOffset += uint32(len(valueBuffers[i]))
	}

	// Calculate header size
	// Layout: header(1) + object_type(1) + field_count(varint) + index_table + values
	headerSize := 1 + 1 + sizeOfCompressedInt(len(keys))
	for _, idx := range indices {
		// Each index entry: name_size(varint) + name + offset(4) + size(2) + flags(1)
		headerSize += sizeOfCompressedInt(len(idx.Name)) + len(idx.Name) + 4 + 2 + 1
	}

	// Allocate buffer
	buf := make([]byte, headerSize+totalValueSize)
	offset := 0

	// Write header
	buf[offset] = ExtFieldIndex
	offset++

	// Write object type
	buf[offset] = TypeObject
	offset++

	// Write field count
	offset += writeCompressedSize(buf[offset:], len(keys))

	// Write index table
	for _, idx := range indices {
		// Write field name
		offset += writeCompressedSize(buf[offset:], len(idx.Name))
		copy(buf[offset:], idx.Name)
		offset += len(idx.Name)

		// Write offset
		binary.LittleEndian.PutUint32(buf[offset:], idx.Offset)
		offset += 4

		// Write size
		binary.LittleEndian.PutUint16(buf[offset:], idx.Size)
		offset += 2

		// Write flags
		buf[offset] = idx.Flags
		offset++
	}

	// Write values
	for _, valueBytes := range valueBuffers {
		copy(buf[offset:], valueBytes)
		offset += len(valueBytes)
	}

	return buf[:offset], nil
}

// DecodeIndexedObject decodes Extension 0 indexed object
func DecodeIndexedObject(data []byte) (map[string]interface{}, error) {
	if len(data) < 3 || data[0] != ExtFieldIndex {
		return nil, fmt.Errorf("invalid indexed object header")
	}

	offset := 1

	// Read object type
	objectType := data[offset]
	if objectType != TypeObject {
		return nil, fmt.Errorf("expected object type, got 0x%02X", objectType)
	}
	offset++

	// Read field count
	fieldCount, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to read field count: %w", err)
	}
	offset += bytesRead

	// Read index table
	indices := make([]FieldIndex, fieldCount)
	for i := 0; i < fieldCount; i++ {
		// Read field name
		nameLen, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, fmt.Errorf("failed to read field name length: %w", err)
		}
		offset += bytesRead

		name := string(data[offset : offset+nameLen])
		offset += nameLen

		// Read offset
		fieldOffset := binary.LittleEndian.Uint32(data[offset:])
		offset += 4

		// Read size
		size := binary.LittleEndian.Uint16(data[offset:])
		offset += 2

		// Read flags
		flags := data[offset]
		offset++

		indices[i] = FieldIndex{
			Name:   name,
			Offset: fieldOffset,
			Size:   size,
			Flags:  flags,
		}
	}

	// Data section starts here
	dataStartOffset := offset

	// Decode each field
	result := make(map[string]interface{}, fieldCount)
	for _, idx := range indices {
		valueOffset := dataStartOffset + int(idx.Offset)
		valueData := data[valueOffset : valueOffset+int(idx.Size)]

		value, _, err := decodeValueAt(valueData, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to decode field %q: %w", idx.Name, err)
		}

		result[idx.Name] = value
	}

	return result, nil
}

// ReadFieldByName reads a single field from an indexed object without decoding everything
func ReadFieldByName(data []byte, fieldName string) (interface{}, error) {
	if len(data) < 3 || data[0] != ExtFieldIndex {
		return nil, fmt.Errorf("invalid indexed object header")
	}

	offset := 2 // Skip header + object type

	// Read field count
	fieldCount, bytesRead, err := readCompressedSize(data, offset)
	if err != nil {
		return nil, err
	}
	offset += bytesRead

	// Search index table
	for i := 0; i < fieldCount; i++ {
		// Read field name
		nameLen, bytesRead, err := readCompressedSize(data, offset)
		if err != nil {
			return nil, err
		}
		offset += bytesRead

		name := string(data[offset : offset+nameLen])
		offset += nameLen

		// Read offset, size, flags
		fieldOffset := binary.LittleEndian.Uint32(data[offset:])
		offset += 4
		size := binary.LittleEndian.Uint16(data[offset:])
		offset += 2
		offset++ // Skip flags

		// If this is the field we're looking for
		if name == fieldName {
			// Calculate absolute offset
			// Need to skip remaining index entries to find data start
			remainingEntries := fieldCount - i - 1
			for j := 0; j < remainingEntries; j++ {
				// Skip name
				nameLen, bytesRead, err := readCompressedSize(data, offset)
				if err != nil {
					return nil, err
				}
				offset += bytesRead + nameLen
				// Skip offset(4) + size(2) + flags(1)
				offset += 7
			}

			// Data section starts here
			dataStartOffset := offset
			valueOffset := dataStartOffset + int(fieldOffset)
			valueData := data[valueOffset : valueOffset+int(size)]

			value, _, err := decodeValueAt(valueData, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to decode field %q: %w", fieldName, err)
			}

			return value, nil
		}
	}

	return nil, fmt.Errorf("field %q not found", fieldName)
}

// inferTypeFromBytes infers BEVE type from encoded bytes
func inferTypeFromBytes(data []byte) byte {
	if len(data) == 0 {
		return 0x00 // null type
	}

	header := data[0]
	return header & 0x07 // Extract type bits (0-2)
}

// sizeOfCompressedInt returns the size in bytes of a compressed integer
func sizeOfCompressedInt(n int) int {
	if n < 64 {
		return 1
	} else if n < 16384 {
		return 2
	} else if n < 1073741824 {
		return 4
	}
	return 8
}
