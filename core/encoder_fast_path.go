package core

import (
	"reflect"
	"unsafe"
)

// Fast path optimizations for common patterns
// These functions bypass generic reflection where possible

// isWideStructSmallValues checks if a struct has many fields with small values
// Wide structs (50+ fields) with primitive types benefit from fast path encoding
func isWideStructSmallValues(info *encoderStructInfo) bool {
	if len(info.fields) < 20 {
		return false // Not wide enough
	}
	
	// Check if most fields are primitives (int, bool, float)
	primitiveCount := 0
	for i := range info.fields {
		field := &info.fields[i]
		switch field.kind {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			primitiveCount++
		}
	}
	
	// If 80%+ fields are primitives, use fast path
	return primitiveCount*100/len(info.fields) >= 80
}

// encodeWideStructFastPath is optimized for structs with many primitive fields
// It pre-allocates buffer space and uses inline encoding to minimize overhead
//
//go:inline
func (e *Encoder) encodeWideStructFastPath(info *encoderStructInfo, base unsafe.Pointer) error {
	if e.Buf == nil {
		// Fast path only works with buffered encoding
		return e.encodeStructPtr(info, base)
	}
	
	// Pre-allocate generous buffer for wide struct
	estimate := len(info.fields) * 8 // Assume ~8 bytes per field
	e.Buf.Grow(estimate)
	
	// Write struct header
	if err := e.WriteByte(0x03); err != nil {
		return err
	}
	
	// Count non-empty fields
	count := info.staticCount
	for _, idx := range info.omitEmpty {
		field := &info.fields[idx]
		fieldPtr := unsafe.Add(base, field.offset)
		if !isStructFieldEmpty(field, fieldPtr) {
			count++
		}
	}
	
	if err := e.WriteCompressedUint(uint64(count)); err != nil {
		return err
	}
	
	// Inline field encoding (critical hot path)
	buf := e.Buf.data
	for i := range info.fields {
		field := &info.fields[i]
		
		// Skip empty fields
		if field.omitEmpty {
			fieldPtr := unsafe.Add(base, field.offset)
			if isStructFieldEmpty(field, fieldPtr) {
				continue
			}
		}
		
		// Write field key (pre-computed)
		buf = append(buf, field.key...)
		
		// Write field value (inline for primitives)
		fieldPtr := unsafe.Add(base, field.offset)
		buf = appendFieldValueInline(buf, field, fieldPtr)
	}
	
	e.Buf.data = buf
	return nil
}

// appendFieldValueInline encodes field value directly into buffer
// This is the hottest path for wide struct encoding
//
//go:inline
func appendFieldValueInline(buf []byte, field *encoderStructField, ptr unsafe.Pointer) []byte {
	switch field.kind {
	case reflect.Bool:
		return appendEncodedBool(buf, *(*bool)(ptr))
	case reflect.Int:
		return appendEncodedInt(buf, int64(*(*int)(ptr)))
	case reflect.Int8:
		return appendEncodedInt(buf, int64(*(*int8)(ptr)))
	case reflect.Int16:
		return appendEncodedInt(buf, int64(*(*int16)(ptr)))
	case reflect.Int32:
		return appendEncodedInt(buf, int64(*(*int32)(ptr)))
	case reflect.Int64:
		return appendEncodedInt(buf, *(*int64)(ptr))
	case reflect.Uint:
		return appendEncodedUint(buf, uint64(*(*uint)(ptr)))
	case reflect.Uint8:
		return appendEncodedUint(buf, uint64(*(*uint8)(ptr)))
	case reflect.Uint16:
		return appendEncodedUint(buf, uint64(*(*uint16)(ptr)))
	case reflect.Uint32:
		return appendEncodedUint(buf, uint64(*(*uint32)(ptr)))
	case reflect.Uint64:
		return appendEncodedUint(buf, *(*uint64)(ptr))
	case reflect.Float32:
		return appendEncodedFloat32(buf, *(*float32)(ptr))
	case reflect.Float64:
		return appendEncodedFloat64(buf, *(*float64)(ptr))
	case reflect.String:
		return appendEncodedString(buf, *(*string)(ptr))
	default:
		// Fallback to slow path for complex types
		// This should rarely happen if isWideStructSmallValues is correct
		return buf
	}
}
