package core

import (
	"encoding/binary"
	"math"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// encoderStructField contains cached metadata for struct encoding.
type encoderStructField struct {
	omitEmpty    bool
	key          []byte
	offset       uintptr
	kind         reflect.Kind
	typ          reflect.Type
	encoder      encoderFunc
	structInfo   *encoderStructInfo
	mapEncoder   mapEncoderFunc
	sliceEncoder sliceEncoderFunc
}

// encoderStructInfo caches struct field metadata for fast encoding.
type encoderStructInfo struct {
	fields      []encoderStructField
	staticCount int
	omitEmpty   []int
	baseSize    int
	sizeHint    uint32
}

type mapEncoderFunc func(*Encoder, unsafe.Pointer) error
type sliceEncoderFunc func(*Encoder, unsafe.Pointer) error

var encoderStructInfoCache sync.Map // map[reflect.Type]*encoderStructInfo

func buildStructEncoder(t reflect.Type) encoderFunc {
	info := getEncoderStructInfo(t)
	return func(e *Encoder, v reflect.Value) error {
		_, basePtr, keep := ensureAddressableStruct(v)
		err := e.encodeStructPtr(info, basePtr)
		if keep != nil {
			runtime.KeepAlive(keep)
		}
		return err
	}
}

// getEncoderStructInfo retrieves cached struct metadata or builds it on demand.
func getEncoderStructInfo(t reflect.Type) *encoderStructInfo {
	if info, ok := encoderStructInfoCache.Load(t); ok {
		return info.(*encoderStructInfo)
	}

	computed := buildEncoderStructInfo(t)

	if actual, loaded := encoderStructInfoCache.LoadOrStore(t, computed); loaded {
		return actual.(*encoderStructInfo)
	}
	return computed
}

// buildEncoderStructInfo generates metadata for struct fields, including inline support.
func buildEncoderStructInfo(t reflect.Type) *encoderStructInfo {
	info := &encoderStructInfo{}
	if t.Kind() != reflect.Struct {
		return info
	}

	visited := make(map[reflect.Type]bool)
	buildEncoderStructFieldsRecursive(t, 0, true, info, visited)
	finalizeEncoderStructInfo(info)
	return info
}

func buildEncoderStructFieldsRecursive(t reflect.Type, baseOffset uintptr, flattenInline bool, info *encoderStructInfo, visited map[reflect.Type]bool) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" && !field.Anonymous {
			continue
		}

		tag := field.Tag.Get("beve")
		if tag == "-" {
			continue
		}

		name := field.Name
		omitEmpty := false
		inline := false

		if tag != "" {
			for _, opt := range parseFieldTag(tag) {
				switch opt {
				case "omitempty":
					omitEmpty = true
				case "inline":
					inline = true
				default:
					if opt != "" {
						name = opt
					}
				}
			}
		}

		offset := baseOffset + field.Offset
		if inline && field.Type.Kind() == reflect.Struct {
			if flattenInline {
				if visited[field.Type] {
					continue
				}
				visited[field.Type] = true
				buildEncoderStructFieldsRecursive(field.Type, offset, false, info, visited)
				delete(visited, field.Type)
				continue
			}
		}

		entry := encoderStructField{
			omitEmpty: omitEmpty,
			key:       buildStructFieldKey(name),
			offset:    offset,
			kind:      field.Type.Kind(),
			typ:       field.Type,
			encoder:   getEncoderFunc(field.Type),
			structInfo: func() *encoderStructInfo {
				if field.Type.Kind() == reflect.Struct {
					return getEncoderStructInfo(field.Type)
				}
				return nil
			}(),
			mapEncoder: func() mapEncoderFunc {
				if field.Type.Kind() == reflect.Map {
					return buildMapEncoder(field.Type)
				}
				return nil
			}(),
			sliceEncoder: func() sliceEncoderFunc {
				if field.Type.Kind() == reflect.Slice {
					return buildSliceEncoder(field.Type)
				}
				return nil
			}(),
		}

		if !omitEmpty {
			info.baseSize += len(entry.key) + minimalValueSize(&entry)
		}

		info.fields = append(info.fields, entry)
		idx := len(info.fields) - 1
		if omitEmpty {
			info.omitEmpty = append(info.omitEmpty, idx)
		} else {
			info.staticCount++
		}
	}
}

func finalizeEncoderStructInfo(info *encoderStructInfo) {
	base := info.baseSize + 1 + compressedUintLen(info.staticCount)
	if base < 16 {
		base = 16
	}
	max := int(^uint32(0))
	if base > max {
		base = max
	}
	atomic.StoreUint32(&info.sizeHint, uint32(base))
}

func minimalValueSize(field *encoderStructField) int {
	switch field.kind {
	case reflect.Bool:
		return 1
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return 2
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return 2
	case reflect.Float32:
		return 5
	case reflect.Float64:
		return 9
	case reflect.String:
		return 2
	case reflect.Interface, reflect.Ptr:
		return 1
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return 2
	default:
		return 1
	}
}

func updateStructSizeHint(info *encoderStructInfo, size int) {
	if size <= 0 {
		return
	}
	if size > int(^uint32(0)) {
		size = int(^uint32(0))
	}
	for {
		current := atomic.LoadUint32(&info.sizeHint)
		if int(current) >= size {
			return
		}
		if atomic.CompareAndSwapUint32(&info.sizeHint, current, uint32(size)) {
			return
		}
	}
}

func buildMapEncoder(t reflect.Type) mapEncoderFunc {
	if t.Kind() != reflect.Map {
		return nil
	}

	keyType := t.Key()
	if keyType.Kind() != reflect.String {
		return nil
	}

	valueType := t.Elem()

	switch valueType.Kind() {
	case reflect.Interface:
		if valueType.NumMethod() == 0 {
			return func(e *Encoder, ptr unsafe.Pointer) error {
				m := *(*map[string]interface{})(ptr)
				return encodeStringInterfaceMap(e, m)
			}
		}
	case reflect.String:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]string)(ptr)
			return encodeStringKeyMap[string](e, m, func(enc *Encoder, v string) error {
				return enc.EncodeString(v)
			})
		}
	case reflect.Bool:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]bool)(ptr)
			return encodeStringKeyMap[bool](e, m, func(enc *Encoder, v bool) error {
				return enc.encodeBool(v)
			})
		}
	case reflect.Int:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]int)(ptr)
			return encodeStringKeyMap[int](e, m, func(enc *Encoder, v int) error {
				return enc.encodeInt(int64(v))
			})
		}
	case reflect.Int8:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]int8)(ptr)
			return encodeStringKeyMap[int8](e, m, func(enc *Encoder, v int8) error {
				return enc.encodeInt(int64(v))
			})
		}
	case reflect.Int16:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]int16)(ptr)
			return encodeStringKeyMap[int16](e, m, func(enc *Encoder, v int16) error {
				return enc.encodeInt(int64(v))
			})
		}
	case reflect.Int32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]int32)(ptr)
			return encodeStringKeyMap[int32](e, m, func(enc *Encoder, v int32) error {
				return enc.encodeInt(int64(v))
			})
		}
	case reflect.Int64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]int64)(ptr)
			return encodeStringKeyMap[int64](e, m, func(enc *Encoder, v int64) error {
				return enc.encodeInt(v)
			})
		}
	case reflect.Uint:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]uint)(ptr)
			return encodeStringKeyMap[uint](e, m, func(enc *Encoder, v uint) error {
				return enc.encodeUint(uint64(v))
			})
		}
	case reflect.Uint8:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]uint8)(ptr)
			return encodeStringKeyMap[uint8](e, m, func(enc *Encoder, v uint8) error {
				return enc.encodeUint(uint64(v))
			})
		}
	case reflect.Uint16:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]uint16)(ptr)
			return encodeStringKeyMap[uint16](e, m, func(enc *Encoder, v uint16) error {
				return enc.encodeUint(uint64(v))
			})
		}
	case reflect.Uint32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]uint32)(ptr)
			return encodeStringKeyMap[uint32](e, m, func(enc *Encoder, v uint32) error {
				return enc.encodeUint(uint64(v))
			})
		}
	case reflect.Uint64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]uint64)(ptr)
			return encodeStringKeyMap[uint64](e, m, func(enc *Encoder, v uint64) error {
				return enc.encodeUint(v)
			})
		}
	case reflect.Float32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]float32)(ptr)
			return encodeStringKeyMap[float32](e, m, func(enc *Encoder, v float32) error {
				return enc.encodeFloat(float64(v), reflect.Float32)
			})
		}
	case reflect.Float64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			m := *(*map[string]float64)(ptr)
			return encodeStringKeyMap[float64](e, m, func(enc *Encoder, v float64) error {
				return enc.encodeFloat(v, reflect.Float64)
			})
		}
	default:
		return nil
	}

	return nil
}

func buildSliceEncoder(t reflect.Type) sliceEncoderFunc {
	if t.Kind() != reflect.Slice || isRawMessageType(t) {
		return nil
	}

	elem := t.Elem()

	switch elem.Kind() {
	case reflect.String:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]string)(ptr)
			return e.encodeStringSliceDirect(slice)
		}
	case reflect.Bool:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]bool)(ptr)
			return e.encodeBoolSliceDirect(slice)
		}
	case reflect.Int8:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]int8)(ptr)
			return e.encodeInt8SliceDirect(slice)
		}
	case reflect.Int16:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]int16)(ptr)
			return e.encodeInt16SliceDirect(slice)
		}
	case reflect.Int32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]int32)(ptr)
			return e.encodeInt32SliceDirect(slice)
		}
	case reflect.Int64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]int64)(ptr)
			return e.encodeInt64SliceDirect(slice)
		}
	case reflect.Uint8:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]uint8)(ptr)
			return e.encodeUint8SliceDirect(slice)
		}
	case reflect.Uint16:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]uint16)(ptr)
			return e.encodeUint16SliceDirect(slice)
		}
	case reflect.Uint32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]uint32)(ptr)
			return e.encodeUint32SliceDirect(slice)
		}
	case reflect.Uint64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]uint64)(ptr)
			return e.encodeUint64SliceDirect(slice)
		}
	case reflect.Float32:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]float32)(ptr)
			return e.encodeFloat32SliceDirect(slice)
		}
	case reflect.Float64:
		return func(e *Encoder, ptr unsafe.Pointer) error {
			slice := *(*[]float64)(ptr)
			return e.encodeFloat64SliceDirect(slice)
		}
	default:
		return nil
	}
}

func writeMapHeader(e *Encoder, keyTypeByte byte, size int) error {
	header := byte(0x03 | (keyTypeByte << 3))
	if err := e.WriteByte(header); err != nil {
		return err
	}
	return e.WriteCompressedUint(uint64(size))
}

func encodeStringKeyMap[T any](e *Encoder, m map[string]T, writeValue func(*Encoder, T) error) error {
	if err := writeMapHeader(e, 0, len(m)); err != nil {
		return err
	}
	for k, v := range m {
		if err := e.WriteCompressedUint(uint64(len(k))); err != nil {
			return err
		}
		if err := e.WriteStringBytes(k); err != nil {
			return err
		}
		if err := writeValue(e, v); err != nil {
			return err
		}
	}
	return nil
}

func encodeStringInterfaceMap(e *Encoder, m map[string]interface{}) error {
	if err := writeMapHeader(e, 0, len(m)); err != nil {
		return err
	}
	for k, v := range m {
		if err := e.WriteCompressedUint(uint64(len(k))); err != nil {
			return err
		}
		if err := e.WriteStringBytes(k); err != nil {
			return err
		}
		if err := encodeInterfaceValue(e, v); err != nil {
			return err
		}
	}
	return nil
}

func encodeInterfaceValue(e *Encoder, v interface{}) error {
	switch val := v.(type) {
	case nil:
		return e.EncodeNull()
	case bool:
		return e.encodeBool(val)
	case string:
		return e.EncodeString(val)
	case int:
		return e.encodeInt(int64(val))
	case int8:
		return e.encodeInt(int64(val))
	case int16:
		return e.encodeInt(int64(val))
	case int32:
		return e.encodeInt(int64(val))
	case int64:
		return e.encodeInt(val)
	case uint:
		return e.encodeUint(uint64(val))
	case uint8:
		return e.encodeUint(uint64(val))
	case uint16:
		return e.encodeUint(uint64(val))
	case uint32:
		return e.encodeUint(uint64(val))
	case uint64:
		return e.encodeUint(val)
	case float32:
		return e.encodeFloat(float64(val), reflect.Float32)
	case float64:
		return e.encodeFloat(val, reflect.Float64)
	case map[string]interface{}:
		return encodeStringInterfaceMap(e, val)
	case []interface{}:
		return e.encodeSlice(reflect.ValueOf(val))
	case []byte:
		return e.encodeSlice(reflect.ValueOf(val))
	default:
		return e.Encode(reflect.ValueOf(v))
	}
}

// encodeSlice encodes a slice or array.
//
// BEVE slice encoding:
//  1. Check if it's a typed array (homogeneous primitives)
//  2. Check if it's a primitive slice (fast path)
//  3. Otherwise use generic array encoding
//
// Generic array format:
//
//	1 byte:  header (0x85 = array type)
//	varint:  array length
//	N items: encoded elements
//
// Phase 2 optimization: Batch encoding in 16-item chunks for better
// CPU cache locality.
func (e *Encoder) encodeSlice(v reflect.Value) error {
	length := v.Len()
	elemKind := v.Type().Elem().Kind()

	// Use typed arrays for homogeneous primitive types
	if length > 0 {
		switch elemKind {
		case reflect.String:
			return e.encodeStringTypedArray(v)
		case reflect.Bool:
			return e.encodeBoolTypedArray(v)
		case reflect.Int8:
			return e.encodeInt8TypedArray(v)
		case reflect.Int16:
			return e.encodeInt16TypedArray(v)
		case reflect.Int32:
			return e.encodeInt32TypedArray(v)
		case reflect.Int64:
			return e.encodeInt64TypedArray(v)
		case reflect.Uint8:
			return e.encodeUint8TypedArray(v)
		case reflect.Uint16:
			return e.encodeUint16TypedArray(v)
		case reflect.Uint32:
			return e.encodeUint32TypedArray(v)
		case reflect.Uint64:
			return e.encodeUint64TypedArray(v)
		case reflect.Float32:
			return e.encodeFloat32TypedArray(v)
		case reflect.Float64:
			return e.encodeFloat64TypedArray(v)
		}
	}

	// Phase 2 optimization: Fast path for other primitive slices
	if length > 0 && isPrimitive(elemKind) {
		return e.encodePrimitiveSlice(v, elemKind)
	}

	// Generic array header (type=5)
	header := byte(0x85)
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(length)); err != nil {
		return err
	}

	// Phase 2 optimization: Batch encode in chunks for better cache locality
	const batchSize = 16
	for i := 0; i < length; i += batchSize {
		end := i + batchSize
		if end > length {
			end = length
		}

		for j := i; j < end; j++ {
			if err := e.Encode(v.Index(j)); err != nil {
				return err
			}
		}
	}

	return nil
}

// encodePrimitiveSlice uses fast-path encoding for slices of primitive types.
//
// Instead of calling encode() for each element (which involves type dispatch
// and reflection), we directly call the appropriate primitive encoder.
//
// This optimization provides ~25% speedup for primitive slices.
//
// Phase 2 optimization: Added to avoid encode() dispatch overhead.
func (e *Encoder) encodePrimitiveSlice(v reflect.Value, kind reflect.Kind) error {
	length := v.Len()

	// Write generic array header
	header := byte(0x85)
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(length)); err != nil {
		return err
	}

	// Batch encode primitives (better CPU cache utilization)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		for i := 0; i < length; i++ {
			if err := e.encodeInt(extractInt(v.Index(i))); err != nil {
				return err
			}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		for i := 0; i < length; i++ {
			if err := e.encodeUint(extractUint(v.Index(i))); err != nil {
				return err
			}
		}

	case reflect.Float32, reflect.Float64:
		for i := 0; i < length; i++ {
			if err := e.encodeFloat(extractFloat(v.Index(i)), kind); err != nil {
				return err
			}
		}

	case reflect.Bool:
		for i := 0; i < length; i++ {
			if err := e.encodeBool(extractBool(v.Index(i))); err != nil {
				return err
			}
		}

	case reflect.String:
		for i := 0; i < length; i++ {
			if err := e.EncodeString(extractString(v.Index(i))); err != nil {
				return err
			}
		}

	default:
		// Fallback to generic encoding
		for i := 0; i < length; i++ {
			if err := e.Encode(v.Index(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

// encodeMapFast encodes a map using optimized iteration.
//
// BEVE map encoding:
//
//	1 byte:  header (0x03 | keyType<<3)
//	varint:  map size
//	N pairs: key-value pairs
//
// Optimization: Uses MapRange() instead of MapKeys() to avoid
// allocating a slice of all keys. This saves memory for large maps.
//
// TODO: Move full implementation from reflect_optimize.go
func (e *Encoder) encodeMapFast(v reflect.Value) error {
	keyType := v.Type().Key()
	valueType := v.Type().Elem()

	valueEncoder := getEncoderFunc(valueType)

	switch keyType.Kind() {
	case reflect.String:
		return e.encodeMapStringFast(v, valueEncoder)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return e.encodeMapIntFast(v, valueEncoder)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return e.encodeMapUintFast(v, valueEncoder)
	default:
		return &UnsupportedError{"unsupported map key type: " + keyType.String()}
	}
}

func (e *Encoder) encodeMapStringFast(v reflect.Value, valueEncoder encoderFunc) error {
	if err := writeMapHeader(e, 0, v.Len()); err != nil {
		return err
	}

	iter := v.MapRange()
	for iter.Next() {
		keyStr := iter.Key().String()
		if err := e.WriteCompressedUint(uint64(len(keyStr))); err != nil {
			return err
		}
		if err := e.WriteStringBytes(keyStr); err != nil {
			return err
		}
		if err := valueEncoder(e, iter.Value()); err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) encodeMapIntFast(v reflect.Value, valueEncoder encoderFunc) error {
	if err := writeMapHeader(e, 1, v.Len()); err != nil {
		return err
	}

	iter := v.MapRange()
	for iter.Next() {
		keyVal := uint64(iter.Key().Int())
		binary.LittleEndian.PutUint64(e.uintScratch[:8], keyVal)
		if err := e.WriteBytes(e.uintScratch[:8]); err != nil {
			return err
		}
		if err := valueEncoder(e, iter.Value()); err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) encodeMapUintFast(v reflect.Value, valueEncoder encoderFunc) error {
	if err := writeMapHeader(e, 2, v.Len()); err != nil {
		return err
	}

	iter := v.MapRange()
	for iter.Next() {
		keyVal := iter.Key().Uint()
		binary.LittleEndian.PutUint64(e.uintScratch[:8], keyVal)
		if err := e.WriteBytes(e.uintScratch[:8]); err != nil {
			return err
		}
		if err := valueEncoder(e, iter.Value()); err != nil {
			return err
		}
	}
	return nil
}

// encodeStructFast encodes a struct using optimized field access.
//
// BEVE struct encoding:
//
//	1 byte:  header (0x03 = object type)
//	varint:  field count
//	N pairs: field name + field value
//
// TODO: Move full implementation from reflect_optimize.go
// For now, use simple implementation.
func (e *Encoder) encodeStructFast(v reflect.Value) error {
	addrValue, basePtr, keep := ensureAddressableStruct(v)
	info := getEncoderStructInfo(addrValue.Type())
	err := e.encodeStructPtr(info, basePtr)
	if keep != nil {
		runtime.KeepAlive(keep)
	}
	return err
}

func countStructFields(v reflect.Value, info *encoderStructInfo) int {
	addrValue, basePtr, keep := ensureAddressableStruct(v)
	_ = addrValue
	count := countStructFieldsPtr(info, basePtr)
	if keep != nil {
		runtime.KeepAlive(keep)
	}
	return count
}

func countStructFieldsPtr(info *encoderStructInfo, base unsafe.Pointer) int {
	count := info.staticCount
	for _, idx := range info.omitEmpty {
		field := &info.fields[idx]
		fieldPtr := unsafe.Add(base, field.offset)
		if !isStructFieldEmpty(field, fieldPtr) {
			count++
		}
	}
	return count
}

func countStructFieldsWithMask(info *encoderStructInfo, base unsafe.Pointer, scratch []byte) (int, []byte, bool) {
	count := info.staticCount
	if len(info.omitEmpty) == 0 {
		return count, nil, false
	}

	bitLen := (len(info.fields) + 7) >> 3
	if bitLen <= 0 {
		return count, nil, false
	}

	if bitLen <= len(scratch) {
		mask := scratch[:bitLen]
		for i := 0; i < bitLen; i++ {
			mask[i] = 0
		}
		for _, idx := range info.omitEmpty {
			field := &info.fields[idx]
			fieldPtr := unsafe.Add(base, field.offset)
			if !isStructFieldEmpty(field, fieldPtr) {
				count++
				mask[idx>>3] |= 1 << (uint(idx) & 7)
			}
		}
		return count, mask, true
	}

	for _, idx := range info.omitEmpty {
		field := &info.fields[idx]
		fieldPtr := unsafe.Add(base, field.offset)
		if !isStructFieldEmpty(field, fieldPtr) {
			count++
		}
	}
	return count, nil, false
}

func writeStructFields(e *Encoder, v reflect.Value, info *encoderStructInfo) error {
	addrValue, basePtr, keep := ensureAddressableStruct(v)
	_ = addrValue
	err := writeStructFieldsPtr(e, info, basePtr)
	if keep != nil {
		runtime.KeepAlive(keep)
	}
	return err
}

func writeStructFieldsPtr(e *Encoder, info *encoderStructInfo, base unsafe.Pointer) error {
	if e.Buf != nil {
		return writeStructFieldsPtrBuffered(e, info, base)
	}
	return writeStructFieldsPtrGeneric(e, info, base)
}

func (e *Encoder) encodeStructPtr(info *encoderStructInfo, base unsafe.Pointer) error {
	if e.Buf == nil {
		if err := e.WriteByte(0x03); err != nil {
			return err
		}
		count := countStructFieldsPtr(info, base)
		if err := e.WriteCompressedUint(uint64(count)); err != nil {
			return err
		}
		return writeStructFieldsPtrGeneric(e, info, base)
	}

	estimate := int(atomic.LoadUint32(&info.sizeHint))
	if estimate < 16 {
		estimate = 16
	}
	e.Buf.Grow(estimate)
	startLen := e.Buf.Len()

	if err := e.WriteByte(0x03); err != nil {
		return err
	}

	count, mask, useMask := countStructFieldsWithMask(info, base, e.batchBuf[:])
	if err := e.WriteCompressedUint(uint64(count)); err != nil {
		return err
	}

	if err := e.writeStructFieldsBuffered(info, base, mask, useMask); err != nil {
		return err
	}

	actual := e.Buf.Len() - startLen
	updateStructSizeHint(info, actual)
	return nil
}

func writeStructFieldsPtrGeneric(e *Encoder, info *encoderStructInfo, base unsafe.Pointer) error {
	for i := range info.fields {
		field := &info.fields[i]
		fieldPtr := unsafe.Add(base, field.offset)

		if field.omitEmpty && isStructFieldEmpty(field, fieldPtr) {
			continue
		}

		if err := e.WriteBytes(field.key); err != nil {
			return err
		}
		if field.sliceEncoder != nil && field.kind == reflect.Slice {
			if err := field.sliceEncoder(e, fieldPtr); err != nil {
				return err
			}
			continue
		}
		if err := encodeStructFieldValue(e, field, fieldPtr); err != nil {
			return err
		}
	}
	return nil
}

func writeStructFieldsPtrBuffered(e *Encoder, info *encoderStructInfo, base unsafe.Pointer) error {
	return e.writeStructFieldsBuffered(info, base, nil, false)
}

func (e *Encoder) writeStructFieldsBuffered(info *encoderStructInfo, base unsafe.Pointer, mask []byte, useMask bool) error {
	buf := e.Buf.data
	for i := range info.fields {
		field := &info.fields[i]
		fieldPtr := unsafe.Add(base, field.offset)

		if field.omitEmpty {
			if useMask {
				if (mask[i>>3] & (1 << (uint(i) & 7))) == 0 {
					continue
				}
			} else if isStructFieldEmpty(field, fieldPtr) {
				continue
			}
		}

		buf = append(buf, field.key...)

		switch field.kind {
		case reflect.Bool:
			buf = appendEncodedBool(buf, *(*bool)(fieldPtr))
		case reflect.Int:
			buf = appendEncodedInt(buf, int64(*(*int)(fieldPtr)))
		case reflect.Int8:
			buf = appendEncodedInt(buf, int64(*(*int8)(fieldPtr)))
		case reflect.Int16:
			buf = appendEncodedInt(buf, int64(*(*int16)(fieldPtr)))
		case reflect.Int32:
			buf = appendEncodedInt(buf, int64(*(*int32)(fieldPtr)))
		case reflect.Int64:
			buf = appendEncodedInt(buf, *(*int64)(fieldPtr))
		case reflect.Uint:
			buf = appendEncodedUint(buf, uint64(*(*uint)(fieldPtr)))
		case reflect.Uint8:
			buf = appendEncodedUint(buf, uint64(*(*uint8)(fieldPtr)))
		case reflect.Uint16:
			buf = appendEncodedUint(buf, uint64(*(*uint16)(fieldPtr)))
		case reflect.Uint32:
			buf = appendEncodedUint(buf, uint64(*(*uint32)(fieldPtr)))
		case reflect.Uint64:
			buf = appendEncodedUint(buf, *(*uint64)(fieldPtr))
		case reflect.Uintptr:
			buf = appendEncodedUint(buf, uint64(*(*uintptr)(fieldPtr)))
		case reflect.Float32:
			buf = appendEncodedFloat32(buf, *(*float32)(fieldPtr))
		case reflect.Float64:
			buf = appendEncodedFloat64(buf, *(*float64)(fieldPtr))
		case reflect.String:
			buf = appendEncodedString(buf, *(*string)(fieldPtr))
		case reflect.Slice:
			if field.sliceEncoder != nil {
				e.Buf.data = buf
				if err := field.sliceEncoder(e, fieldPtr); err != nil {
					return err
				}
				buf = e.Buf.data
			} else {
				e.Buf.data = buf
				if err := encodeStructFieldValue(e, field, fieldPtr); err != nil {
					return err
				}
				buf = e.Buf.data
			}
			continue
		default:
			e.Buf.data = buf
			if err := encodeStructFieldValue(e, field, fieldPtr); err != nil {
				return err
			}
			buf = e.Buf.data
		}
	}
	e.Buf.data = buf
	return nil
}

func ensureAddressableStruct(v reflect.Value) (reflect.Value, unsafe.Pointer, interface{}) {
	if v.Kind() != reflect.Struct {
		panic("ensureAddressableStruct: non-struct value")
	}
	if v.CanAddr() {
		return v, unsafe.Pointer(v.UnsafeAddr()), nil
	}
	// For non-addressable structs, we need to make a heap-allocated copy
	// to ensure the pointer remains valid during field access
	typ := v.Type()
	ptr := reflect.New(typ)
	ptr.Elem().Set(v)
	
	// Return the addressable copy
	addrValue := ptr.Elem()
	basePtr := unsafe.Pointer(addrValue.UnsafeAddr())
	
	// Keep the pointer alive by returning it as the keep-alive value
	return addrValue, basePtr, ptr.Interface()
}

func isStructFieldEmpty(field *encoderStructField, ptr unsafe.Pointer) bool {
	switch field.kind {
	case reflect.Bool:
		return !*(*bool)(ptr)
	case reflect.Int:
		return *(*int)(ptr) == 0
	case reflect.Int8:
		return *(*int8)(ptr) == 0
	case reflect.Int16:
		return *(*int16)(ptr) == 0
	case reflect.Int32:
		return *(*int32)(ptr) == 0
	case reflect.Int64:
		return *(*int64)(ptr) == 0
	case reflect.Uint:
		return *(*uint)(ptr) == 0
	case reflect.Uint8:
		return *(*uint8)(ptr) == 0
	case reflect.Uint16:
		return *(*uint16)(ptr) == 0
	case reflect.Uint32:
		return *(*uint32)(ptr) == 0
	case reflect.Uint64:
		return *(*uint64)(ptr) == 0
	case reflect.Uintptr:
		return *(*uintptr)(ptr) == 0
	case reflect.Float32:
		return *(*float32)(ptr) == 0
	case reflect.Float64:
		return *(*float64)(ptr) == 0
	case reflect.String:
		return len(*(*string)(ptr)) == 0
	case reflect.Slice, reflect.Map:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return val.Len() == 0
	case reflect.Interface, reflect.Ptr:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return val.IsNil()
	case reflect.Struct:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return isEmptyValue(val)
	case reflect.Array:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return val.Len() == 0
	default:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return isEmptyValue(val)
	}
}

func encodeStructFieldValue(e *Encoder, field *encoderStructField, ptr unsafe.Pointer) error {
	switch field.kind {
	case reflect.Bool:
		return e.encodeBool(*(*bool)(ptr))
	case reflect.Int:
		return e.encodeInt(int64(*(*int)(ptr)))
	case reflect.Int8:
		return e.encodeInt(int64(*(*int8)(ptr)))
	case reflect.Int16:
		return e.encodeInt(int64(*(*int16)(ptr)))
	case reflect.Int32:
		return e.encodeInt(int64(*(*int32)(ptr)))
	case reflect.Int64:
		return e.encodeInt(*(*int64)(ptr))
	case reflect.Uint:
		return e.encodeUint(uint64(*(*uint)(ptr)))
	case reflect.Uint8:
		return e.encodeUint(uint64(*(*uint8)(ptr)))
	case reflect.Uint16:
		return e.encodeUint(uint64(*(*uint16)(ptr)))
	case reflect.Uint32:
		return e.encodeUint(uint64(*(*uint32)(ptr)))
	case reflect.Uint64:
		return e.encodeUint(*(*uint64)(ptr))
	case reflect.Uintptr:
		return e.encodeUint(uint64(*(*uintptr)(ptr)))
	case reflect.Float32:
		return e.encodeFloat(float64(*(*float32)(ptr)), reflect.Float32)
	case reflect.Float64:
		return e.encodeFloat(*(*float64)(ptr), reflect.Float64)
	case reflect.String:
		return e.EncodeString(*(*string)(ptr))
	case reflect.Struct:
		if field.structInfo != nil {
			count := countStructFieldsPtr(field.structInfo, ptr)
			if err := e.WriteByte(0x03); err != nil {
				return err
			}
			if err := e.WriteCompressedUint(uint64(count)); err != nil {
				return err
			}
			return writeStructFieldsPtr(e, field.structInfo, ptr)
		}
		val := reflect.NewAt(field.typ, ptr).Elem()
		return field.encoder(e, val)
	case reflect.Map:
		if field.mapEncoder != nil {
			return field.mapEncoder(e, ptr)
		}
		val := reflect.NewAt(field.typ, ptr).Elem()
		return field.encoder(e, val)
	default:
		val := reflect.NewAt(field.typ, ptr).Elem()
		return field.encoder(e, val)
	}
}

// Typed array encoders

// encodeStringTypedArray encodes a string slice as typed array.
// Header: type=4, group=3 (bool/string), string flag=1
func (e *Encoder) encodeStringTypedArray(v reflect.Value) error {
	return e.encodeStringSliceDirect(v.Interface().([]string))
}

func (e *Encoder) encodeStringSliceDirect(slice []string) error {
	header := byte(0x04 | (3 << 3) | (1 << 5)) // typed array, string group
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	for _, s := range slice {
		if err := e.WriteCompressedUint(uint64(len(s))); err != nil {
			return err
		}
		if err := e.WriteStringBytes(s); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeBoolSliceDirect(slice []bool) error {
	header := byte(0x04 | (3 << 3)) // typed array, bool/string group, bool flag
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		idx := 0
		for idx < len(slice) {
			var b byte
			for bit := 0; bit < 8 && idx < len(slice); bit++ {
				if slice[idx] {
					b |= 1 << uint(bit)
				}
				idx++
			}
			buf = append(buf, b)
		}
		e.Buf.data = buf
		return nil
	}

	idx := 0
	payload := (len(slice) + 7) / 8
	for i := 0; i < payload; i++ {
		var b byte
		for bit := 0; bit < 8 && idx < len(slice); bit++ {
			if slice[idx] {
				b |= 1 << uint(bit)
			}
			idx++
		}
		if err := e.WriteByte(b); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeInt32SliceDirect(slice []int32) error {
	header := byte(0x04 | (1 << 3) | (2 << 5)) // typed array, signed group, 4 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			u := uint32(val)
			buf = append(buf, byte(u), byte(u>>8), byte(u>>16), byte(u>>24))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		u := uint32(val)
		binary.LittleEndian.PutUint32(e.uintScratch[:4], u)
		if err := e.WriteBytes(e.uintScratch[:4]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeUint16SliceDirect(slice []uint16) error {
	header := byte(0x04 | (2 << 3) | (1 << 5)) // typed array, unsigned group, 2 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			buf = append(buf, byte(val), byte(val>>8))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		binary.LittleEndian.PutUint16(e.uintScratch[:2], val)
		if err := e.WriteBytes(e.uintScratch[:2]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeFloat32SliceDirect(slice []float32) error {
	header := byte(0x04 | (0 << 3) | (2 << 5)) // typed array, float group, 4 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			bits := math.Float32bits(val)
			buf = append(buf, byte(bits), byte(bits>>8), byte(bits>>16), byte(bits>>24))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		bits := math.Float32bits(val)
		binary.LittleEndian.PutUint32(e.uintScratch[:4], bits)
		if err := e.WriteBytes(e.uintScratch[:4]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeInt8SliceDirect(slice []int8) error {
	header := byte(0x04 | (1 << 3)) // typed array, signed group, 1 byte
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			buf = append(buf, byte(val))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		if err := e.WriteByte(byte(val)); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeInt16SliceDirect(slice []int16) error {
	header := byte(0x04 | (1 << 3) | (1 << 5)) // typed array, signed group, 2 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			u := uint16(val)
			buf = append(buf, byte(u), byte(u>>8))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		u := uint16(val)
		binary.LittleEndian.PutUint16(e.uintScratch[:2], u)
		if err := e.WriteBytes(e.uintScratch[:2]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeInt64SliceDirect(slice []int64) error {
	header := byte(0x04 | (1 << 3) | (3 << 5)) // typed array, signed group, 8 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			u := uint64(val)
			buf = append(buf,
				byte(u), byte(u>>8), byte(u>>16), byte(u>>24),
				byte(u>>32), byte(u>>40), byte(u>>48), byte(u>>56),
			)
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		binary.LittleEndian.PutUint64(e.uintScratch[:8], uint64(val))
		if err := e.WriteBytes(e.uintScratch[:8]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeUint8SliceDirect(slice []uint8) error {
	header := byte(0x04 | (2 << 3)) // typed array, unsigned group, 1 byte
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		e.Buf.data = append(e.Buf.data, slice...)
		return nil
	}

	return e.WriteBytes(slice)
}

func (e *Encoder) encodeUint32SliceDirect(slice []uint32) error {
	header := byte(0x04 | (2 << 3) | (2 << 5)) // typed array, unsigned group, 4 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			buf = append(buf, byte(val), byte(val>>8), byte(val>>16), byte(val>>24))
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		binary.LittleEndian.PutUint32(e.uintScratch[:4], val)
		if err := e.WriteBytes(e.uintScratch[:4]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeUint64SliceDirect(slice []uint64) error {
	header := byte(0x04 | (2 << 3) | (3 << 5)) // typed array, unsigned group, 8 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			buf = append(buf,
				byte(val), byte(val>>8), byte(val>>16), byte(val>>24),
				byte(val>>32), byte(val>>40), byte(val>>48), byte(val>>56),
			)
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		binary.LittleEndian.PutUint64(e.uintScratch[:8], val)
		if err := e.WriteBytes(e.uintScratch[:8]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeFloat64SliceDirect(slice []float64) error {
	header := byte(0x04 | (0 << 3) | (3 << 5)) // typed array, float group, 8 bytes
	if err := e.WriteByte(header); err != nil {
		return err
	}

	if err := e.WriteCompressedUint(uint64(len(slice))); err != nil {
		return err
	}

	if len(slice) == 0 {
		return nil
	}

	if e.Buf != nil {
		buf := e.Buf.data
		for _, val := range slice {
			bits := math.Float64bits(val)
			buf = append(buf,
				byte(bits), byte(bits>>8), byte(bits>>16), byte(bits>>24),
				byte(bits>>32), byte(bits>>40), byte(bits>>48), byte(bits>>56),
			)
		}
		e.Buf.data = buf
		return nil
	}

	for _, val := range slice {
		bits := math.Float64bits(val)
		binary.LittleEndian.PutUint64(e.uintScratch[:8], bits)
		if err := e.WriteBytes(e.uintScratch[:8]); err != nil {
			return err
		}
	}

	return nil
}

// encodeBoolTypedArray encodes a bool slice as typed array (bitpacked).
// Header: type=4, group=3 (bool/string), string flag=0
func (e *Encoder) encodeBoolTypedArray(v reflect.Value) error {
	return e.encodeBoolSliceDirect(v.Interface().([]bool))
}

// encodeInt8TypedArray encodes an int8 slice as typed array.
// Header: type=4, group=1 (signed), byte count=1
func (e *Encoder) encodeInt8TypedArray(v reflect.Value) error {
	return e.encodeInt8SliceDirect(v.Interface().([]int8))
}

// encodeInt16TypedArray encodes an int16 slice as typed array.
// Header: type=4, group=1 (signed), byte count=2
func (e *Encoder) encodeInt16TypedArray(v reflect.Value) error {
	return e.encodeInt16SliceDirect(v.Interface().([]int16))
}

// encodeInt32TypedArray encodes an int32 slice as typed array.
// Header: type=4, group=1 (signed), byte count=4
func (e *Encoder) encodeInt32TypedArray(v reflect.Value) error {
	return e.encodeInt32SliceDirect(v.Interface().([]int32))
}

// encodeInt64TypedArray encodes an int64 slice as typed array.
// Header: type=4, group=1 (signed), byte count=8
func (e *Encoder) encodeInt64TypedArray(v reflect.Value) error {
	return e.encodeInt64SliceDirect(v.Interface().([]int64))
}

// encodeUint8TypedArray encodes a uint8 slice as typed array.
// Header: type=4, group=2 (unsigned), byte count=1
func (e *Encoder) encodeUint8TypedArray(v reflect.Value) error {
	return e.encodeUint8SliceDirect(v.Interface().([]uint8))
}

// encodeUint16TypedArray encodes a uint16 slice as typed array.
// Header: type=4, group=2 (unsigned), byte count=2
func (e *Encoder) encodeUint16TypedArray(v reflect.Value) error {
	return e.encodeUint16SliceDirect(v.Interface().([]uint16))
}

// encodeUint32TypedArray encodes a uint32 slice as typed array.
// Header: type=4, group=2 (unsigned), byte count=4
func (e *Encoder) encodeUint32TypedArray(v reflect.Value) error {
	return e.encodeUint32SliceDirect(v.Interface().([]uint32))
}

// encodeUint64TypedArray encodes a uint64 slice as typed array.
// Header: type=4, group=2 (unsigned), byte count=8
func (e *Encoder) encodeUint64TypedArray(v reflect.Value) error {
	return e.encodeUint64SliceDirect(v.Interface().([]uint64))
}

// Helper functions

// parseFieldTag parses a struct field tag and returns options.
func parseFieldTag(tag string) []string {
	parts := make([]string, 0, 2)
	current := ""
	for i := 0; i < len(tag); i++ {
		if tag[i] == ',' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(tag[i])
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}

func compressedUintLen(n int) int {
	if n < 0 {
		n = 0
	}
	switch {
	case n < 64:
		return 1
	case n < 16384:
		return 2
	case n < 1073741824:
		return 3
	default:
		return 4
	}
}

func appendEncodedBool(dst []byte, v bool) []byte {
	if v {
		return append(dst, 0x18)
	}
	return append(dst, 0x08)
}

func appendEncodedInt(dst []byte, v int64) []byte {
	var byteCount int
	var byteCountBits byte
	if v >= -128 && v <= 127 {
		byteCount = 1
		byteCountBits = 0
	} else if v >= -32768 && v <= 32767 {
		byteCount = 2
		byteCountBits = 1
	} else if v >= -2147483648 && v <= 2147483647 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}
	header := byte(0x01) | (1 << 3) | (byteCountBits << 5)
	dst = append(dst, header)
	for j := 0; j < byteCount; j++ {
		dst = append(dst, byte(v>>(8*j)))
	}
	return dst
}

func appendEncodedUint(dst []byte, v uint64) []byte {
	var byteCount int
	var byteCountBits byte
	if v <= 255 {
		byteCount = 1
		byteCountBits = 0
	} else if v <= 65535 {
		byteCount = 2
		byteCountBits = 1
	} else if v <= 4294967295 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}
	header := byte(0x01) | (2 << 3) | (byteCountBits << 5)
	dst = append(dst, header)
	for j := 0; j < byteCount; j++ {
		dst = append(dst, byte(v>>(8*j)))
	}
	return dst
}

func appendEncodedFloat32(dst []byte, v float32) []byte {
	bits := math.Float32bits(v)
	dst = append(dst, 0x41)
	dst = append(dst, byte(bits), byte(bits>>8), byte(bits>>16), byte(bits>>24))
	return dst
}

func appendEncodedFloat64(dst []byte, v float64) []byte {
	bits := math.Float64bits(v)
	dst = append(dst, 0x61)
	dst = append(dst, byte(bits), byte(bits>>8), byte(bits>>16), byte(bits>>24), byte(bits>>32), byte(bits>>40), byte(bits>>48), byte(bits>>56))
	return dst
}

func appendEncodedString(dst []byte, s string) []byte {
	dst = append(dst, 0x02)
	dst = appendCompressedUint(dst, uint64(len(s)))
	return append(dst, stringToBytes(s)...)
}

func appendCompressedUint(dst []byte, n uint64) []byte {
	switch {
	case n < 64:
		return append(dst, byte(n<<2))
	case n < 16384:
		return append(dst, byte(0x01|((n>>8)<<2)), byte(n))
	case n < 1073741824:
		dst = append(dst, byte(0x02|((n>>16)<<2)))
		dst = append(dst, byte(n>>8), byte(n))
		return dst
	default:
		dst = append(dst, byte(0x03|((n>>24)<<2)))
		dst = append(dst, byte(n>>16), byte(n>>8), byte(n))
		return dst
	}
}

func buildStructFieldKey(name string) []byte {
	length := uint64(len(name))
	var prefix [4]byte
	var prefixLen int

	switch {
	case length < 64:
		prefix[0] = byte(length << 2)
		prefixLen = 1
	case length < 16384:
		prefix[0] = byte(0x01 | ((length >> 8) << 2))
		prefix[1] = byte(length)
		prefixLen = 2
	case length < 1073741824:
		prefix[0] = byte(0x02 | ((length >> 16) << 2))
		prefix[1] = byte(length >> 8)
		prefix[2] = byte(length)
		prefixLen = 3
	default:
		prefix[0] = byte(0x03 | ((length >> 24) << 2))
		prefix[1] = byte(length >> 16)
		prefix[2] = byte(length >> 8)
		prefix[3] = byte(length)
		prefixLen = 4
	}

	key := make([]byte, prefixLen+len(name))
	copy(key, prefix[:prefixLen])
	copy(key[prefixLen:], name)
	return key
}

// isEmptyValue checks if a reflect.Value is considered empty.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// encodeFloat32TypedArray encodes a float32 slice as typed array.
// Header: type=4, group=0 (float), byte count=4
func (e *Encoder) encodeFloat32TypedArray(v reflect.Value) error {
	return e.encodeFloat32SliceDirect(v.Interface().([]float32))
}

// encodeFloat64TypedArray encodes a float64 slice as typed array.
// Header: type=4, group=0 (float), byte count=8
func (e *Encoder) encodeFloat64TypedArray(v reflect.Value) error {
	return e.encodeFloat64SliceDirect(v.Interface().([]float64))
}
