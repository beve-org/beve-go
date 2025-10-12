package core

import (
	"encoding/binary"
	"math"
	"reflect"
	"sync"
	"unsafe"
)

const intBitSize = 32 << (^uint(0) >> 63)

// decoder_collections.go - Decoders for collections (map, struct, array, slice)

// DecodeObject decodes an object (map or struct).
//
// BEVE object header format:
//   - Bits 0-2: Type (3 = object)
//   - Bits 3-4: Key type (0=string, 1=int, 2=uint)
//   - Bits 5-7: Reserved
func (d *Decoder) DecodeObject(v reflect.Value, header byte) error {
	keyType := (header >> 3) & 0x03

	size, err := d.ReadCompressedUint()
	if err != nil {
		return err
	}

	if v.Kind() == reflect.Map {
		return d.DecodeMap(v, keyType, int(size))
	} else if v.Kind() == reflect.Struct {
		return d.DecodeStruct(v, keyType, int(size))
	}

	return &UnsupportedError{"object type not supported"}
}

// DecodeMap decodes a map value.
//
// For each entry:
//  1. Read key (type determined by keyType in header)
//  2. Read value (full BEVE encoding)
func (d *Decoder) DecodeMap(v reflect.Value, keyType byte, size int) error {
	mapType := v.Type()
	keyTarget := mapType.Key()
	elemType := mapType.Elem()

	if v.IsNil() {
		if size > 0 {
			v.Set(reflect.MakeMapWithSize(mapType, size))
		} else {
			v.Set(reflect.MakeMap(mapType))
		}
	}

	valueDecoder := getMapValueDecoder(elemType)
	valueHolder := reflect.New(elemType)

	for i := 0; i < size; i++ {
		var (
			key reflect.Value
			err error
		)
		if keyType == 0 {
			keyStr, kerr := d.ReadKeyString()
			if kerr != nil {
				return kerr
			}
			key = reflect.ValueOf(keyStr)
		} else {
			key, err = d.ReadKey(keyType)
			if err != nil {
				return err
			}
		}

		convertedKey, err := convertMapKeyValue(key, keyTarget, keyType)
		if err != nil {
			return err
		}

		value := valueHolder.Elem()
		value.SetZero()
		if valueDecoder != nil {
			if err := valueDecoder(d, value); err != nil {
				return err
			}
		} else {
			if err := d.Decode(value); err != nil {
				return err
			}
		}

		v.SetMapIndex(convertedKey, value)
	}

	return nil
}

// DecodeStruct decodes a struct value.
//
// Uses getStructInfo to get field mapping by name/tag.
// Unknown fields are skipped.
func (d *Decoder) DecodeStruct(v reflect.Value, keyType byte, size int) error {
	if keyType != 0 {
		return &UnsupportedError{"struct key type not supported"}
	}

	info := getStructInfo(v.Type())

	if !v.CanAddr() {
		return d.decodeStructSlow(v, info, size)
	}

	basePtr := unsafe.Pointer(v.UnsafeAddr())

	for i := 0; i < size; i++ {
		keyStr, err := d.ReadKeyString()
		if err != nil {
			return err
		}

		if fieldInfo, ok := info.fieldMap[keyStr]; ok {
			fieldPtr := unsafe.Add(basePtr, fieldInfo.offset)
			if fieldInfo.decoder != nil {
				if err := fieldInfo.decoder(d, fieldPtr); err != nil {
					return err
				}
			} else {
				if err := d.decodeStructFieldGeneric(fieldInfo, fieldPtr); err != nil {
					return err
				}
			}
		} else {
			if err := d.SkipValue(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Decoder) decodeStructSlow(v reflect.Value, info *structInfo, size int) error {
	for i := 0; i < size; i++ {
		keyStr, err := d.ReadKeyString()
		if err != nil {
			return err
		}

		if fieldInfo, ok := info.fieldMap[keyStr]; ok {
			field := v.FieldByIndex(fieldInfo.index)
			if err := d.Decode(field); err != nil {
				return err
			}
		} else {
			if err := d.SkipValue(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Decoder) decodeStructFieldGeneric(field *structField, fieldPtr unsafe.Pointer) error {
	value := reflect.NewAt(field.typ, fieldPtr).Elem()
	if field.kind == reflect.Ptr {
		if value.IsNil() {
			value.Set(reflect.New(field.typ.Elem()))
		}
		return d.Decode(value.Elem())
	}
	if field.kind == reflect.Interface {
		if value.IsNil() || !value.Elem().IsValid() {
			var dynamic interface{}
			if err := d.Decode(reflect.ValueOf(&dynamic).Elem()); err != nil {
				return err
			}
			value.Set(reflect.ValueOf(dynamic))
			return nil
		}
		return d.Decode(value.Elem())
	}
	return d.Decode(value)
}

// DecodeGenericArray decodes a generic array.
//
// BEVE generic array format:
//   - Compressed uint: array length
//   - N values: Each value fully encoded with type header
func (d *Decoder) DecodeGenericArray(v reflect.Value) error {
	size, err := d.ReadCompressedUint()
	if err != nil {
		return err
	}

	length, err := CheckedLength(size)
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			if err := d.Decode(v.Index(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() != length {
			return &UnsupportedError{"array length mismatch"}
		}
		for i := 0; i < length; i++ {
			if err := d.Decode(v.Index(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		// Decode as []interface{}
		slice := make([]interface{}, length)
		for i := 0; i < length; i++ {
			var elem interface{}
			if err := d.Decode(reflect.ValueOf(&elem).Elem()); err != nil {
				return err
			}
			slice[i] = elem
		}
		v.Set(reflect.ValueOf(slice))
		return nil
	default:
		return &UnsupportedError{"expected array or slice"}
	}
}

// ReadKey reads a map/struct key based on key type.
//
// Key types:
//   - 0: String (compressed length + UTF-8 data)
//   - 1: Signed int (8 bytes, little-endian)
//   - 2: Unsigned int (8 bytes, little-endian)
func (d *Decoder) ReadKey(keyType byte) (reflect.Value, error) {
	switch keyType {
	case 0: // string
		size, err := d.ReadCompressedUint()
		if err != nil {
			return reflect.Value{}, err
		}
		data, err := d.ReadBytes(int(size))
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(bytesToString(data)), nil
	case 1: // signed int
		data, err := d.ReadBytes(8)
		if err != nil {
			return reflect.Value{}, err
		}
		val := int64(binary.LittleEndian.Uint64(data))
		return reflect.ValueOf(val), nil
	case 2: // unsigned int
		data, err := d.ReadBytes(8)
		if err != nil {
			return reflect.Value{}, err
		}
		val := binary.LittleEndian.Uint64(data)
		return reflect.ValueOf(val), nil
	}
	return reflect.Value{}, &UnsupportedError{"unsupported key type"}
}

// ReadKeyString reads a string key without allocating a new string copy.
func (d *Decoder) ReadKeyString() (string, error) {
	size, err := d.ReadCompressedUint()
	if err != nil {
		return "", err
	}
	data, err := d.ReadBytes(int(size))
	if err != nil {
		return "", err
	}
	return bytesToString(data), nil
}

// SkipValue skips a value in the stream.
//
// This is used when encountering unknown struct fields or
// when we need to skip past values we don't care about.
func (d *Decoder) SkipValue() error {
	header, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch header & 0x07 {
	case 0: // null/bool - already consumed
		return nil
	case 1: // number
		byteCountBits := (header >> 5) & 0x07
		byteCount := d.GetByteCount(byteCountBits)
		d.Pos += byteCount
		return nil
	case 2: // string
		size, err := d.ReadCompressedUint()
		if err != nil {
			return err
		}
		d.Pos += int(size)
		return nil
	case 3: // object
		size, err := d.ReadCompressedUint()
		if err != nil {
			return err
		}
		for i := uint64(0); i < size; i++ {
			if err := d.SkipValue(); err != nil { // key
				return err
			}
			if err := d.SkipValue(); err != nil { // value
				return err
			}
		}
		return nil
	case 4: // typed array
		// Skip typed array - read size and skip data
		size, err := d.ReadCompressedUint()
		if err != nil {
			return err
		}
		length, err := CheckedLength(size)
		if err != nil {
			return err
		}
		group := (header >> 3) & 0x03
		switch group {
		case 0, 1, 2: // numeric arrays
			byteCountBits := (header >> 5) & 0x03
			byteCount := d.GetByteCount(byteCountBits)
			d.Pos += length * byteCount
		case 3: // bool or string arrays
			isString := ((header >> 5) & 0x01) == 1
			if isString {
				// String array - skip each string
				for i := 0; i < length; i++ {
					sz, err := d.ReadCompressedUint()
					if err != nil {
						return err
					}
					d.Pos += int(sz)
				}
			} else {
				// Bool array - bitpacked
				payload := (length + 7) / 8
				d.Pos += payload
			}
		}
		return nil
	case 5: // generic array
		size, err := d.ReadCompressedUint()
		if err != nil {
			return err
		}
		length, err := CheckedLength(size)
		if err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			if err := d.SkipValue(); err != nil {
				return err
			}
		}
		return nil
	case 6: // extension
		return &UnsupportedError{"extension type not supported"}
	default:
		return &UnsupportedError{"unknown type"}
	}
}

// Helper functions for struct and map decoding

type fieldDecoderFunc func(*Decoder, unsafe.Pointer) error
type mapValueDecoderFunc func(*Decoder, reflect.Value) error

// structField represents a struct field for encoding/decoding
type structField struct {
	name      string
	typ       reflect.Type
	index     []int
	tag       string
	omitEmpty bool
	inline    bool
	offset    uintptr
	kind      reflect.Kind
	decoder   fieldDecoderFunc
	bitSize   int
}

// structInfo contains cached struct field information
type structInfo struct {
	fields   []*structField
	fieldMap map[string]*structField
}

// structInfoCache caches struct field information for performance
var structInfoCache sync.Map      // map[reflect.Type]*structInfo
var mapValueDecoderCache sync.Map // map[reflect.Type]mapValueDecoderFunc

// getStructInfo returns cached struct field information
func getStructInfo(t reflect.Type) *structInfo {
	if t.Kind() != reflect.Struct {
		return &structInfo{
			fields:   []*structField{},
			fieldMap: make(map[string]*structField),
		}
	}

	// Check cache first
	if cached, ok := structInfoCache.Load(t); ok {
		return cached.(*structInfo)
	}

	fields := make([]*structField, 0, t.NumField())
	fieldMap := make(map[string]*structField)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" && !f.Anonymous {
			// Skip unexported fields (unless anonymous for embedding)
			continue
		}

		// Get field name (use tag if available)
		fieldName := f.Name
		// Get configured struct tag name (defaults to "beve")
		structTagName := "beve"
		if GetStructTag != nil {
			structTagName = GetStructTag()
		}

		// Try configured tag first, fallback to "json" if not found
		tag := f.Tag.Get(structTagName)
		if tag == "" && structTagName != "json" {
			tag = f.Tag.Get("json")
		}

		omitEmpty := false
		inline := false

		if tag == "-" {
			continue // Skip this field
		}

		// Parse tag options
		if tag != "" {
			parts := parseTag(tag)
			if len(parts) > 0 && parts[0] != "" {
				fieldName = parts[0]
			}
			for _, opt := range parts[1:] {
				if opt == "omitempty" {
					omitEmpty = true
				} else if opt == "inline" {
					inline = true
				}
			}
		}

		field := &structField{
			name:      fieldName,
			typ:       f.Type,
			index:     []int{i},
			tag:       string(f.Tag),
			omitEmpty: omitEmpty,
			inline:    inline,
		}

		field.offset = computeFieldOffset(t, field.index)
		field.kind = field.typ.Kind()
		field.bitSize = computeFieldBitSize(field.kind, field.typ)
		field.decoder = buildStructFieldDecoder(field)

		fields = append(fields, field)

		// Handle inline structs - add nested fields to fieldMap
		if inline && f.Type.Kind() == reflect.Struct {
			for j := 0; j < f.Type.NumField(); j++ {
				nestedField := f.Type.Field(j)
				if nestedField.PkgPath != "" {
					continue
				}

				nestedName := nestedField.Name

				// Get configured struct tag name for nested field
				nestedStructTagName := "beve"
				if GetStructTag != nil {
					nestedStructTagName = GetStructTag()
				}

				// Try configured tag first, fallback to "json"
				nestedTag := nestedField.Tag.Get(nestedStructTagName)
				if nestedTag == "" && nestedStructTagName != "json" {
					nestedTag = nestedField.Tag.Get("json")
				}

				if nestedTag != "" && nestedTag != "-" {
					parts := parseTag(nestedTag)
					if len(parts) > 0 && parts[0] != "" {
						nestedName = parts[0]
					}
				}

				// Map nested field with parent index
				nestedFieldInfo := &structField{
					name:  nestedName,
					typ:   nestedField.Type,
					index: []int{i, j}, // Path through parent field
					tag:   string(nestedField.Tag),
				}
				nestedFieldInfo.offset = computeFieldOffset(t, nestedFieldInfo.index)
				nestedFieldInfo.kind = nestedFieldInfo.typ.Kind()
				nestedFieldInfo.bitSize = computeFieldBitSize(nestedFieldInfo.kind, nestedFieldInfo.typ)
				nestedFieldInfo.decoder = buildStructFieldDecoder(nestedFieldInfo)
				fieldMap[nestedName] = nestedFieldInfo
				if nestedName != nestedField.Name {
					fieldMap[nestedField.Name] = nestedFieldInfo
				}
			}
		} else {
			// Map by both actual field name and tag name for decoder compatibility
			fieldMap[f.Name] = field
			if fieldName != f.Name {
				fieldMap[fieldName] = field
			}
		}
	}

	info := &structInfo{
		fields:   fields,
		fieldMap: fieldMap,
	}

	// Cache the result
	structInfoCache.Store(t, info)
	return info
}

// parseTag splits a struct tag string into parts
func parseTag(tag string) []string {
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

func computeFieldOffset(root reflect.Type, index []int) uintptr {
	offset := uintptr(0)
	current := root
	for _, idx := range index {
		field := current.Field(idx)
		offset += field.Offset
		current = field.Type
	}
	return offset
}

func computeFieldBitSize(kind reflect.Kind, typ reflect.Type) int {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		if bits := typ.Bits(); bits != 0 {
			return int(bits)
		}
		if kind == reflect.Int || kind == reflect.Uint {
			return intBitSize
		}
	}
	return 0
}

func getMapValueDecoder(t reflect.Type) mapValueDecoderFunc {
	if cached, ok := mapValueDecoderCache.Load(t); ok {
		if cached == nil {
			return nil
		}
		return cached.(mapValueDecoderFunc)
	}

	decoder := buildMapValueDecoder(t)
	if decoder == nil {
		mapValueDecoderCache.Store(t, nil)
		return nil
	}

	mapValueDecoderCache.Store(t, decoder)
	return decoder
}

func buildMapValueDecoder(t reflect.Type) mapValueDecoderFunc {
	kind := t.Kind()
	bitSize := computeFieldBitSize(kind, t)

	switch kind {
	case reflect.Bool:
		return func(d *Decoder, value reflect.Value) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			if header&0x07 != 0 {
				return d.Decode(value)
			}
			d.Pos++
			if header&0x08 == 0 {
				value.SetBool(false)
				return nil
			}
			value.SetBool(header&0x10 != 0)
			return nil
		}
	case reflect.String:
		return func(d *Decoder, value reflect.Value) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeBits := header & 0x07
			if typeBits == 0 {
				if header&0x08 != 0 {
					return d.Decode(value)
				}
				d.Pos++
				value.SetString("")
				return nil
			}
			if typeBits != 2 {
				return d.Decode(value)
			}
			d.Pos++
			size, err := d.ReadCompressedUint()
			if err != nil {
				return err
			}
			data, err := d.ReadBytes(int(size))
			if err != nil {
				return err
			}
			value.SetString(bytesToString(data))
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(d *Decoder, value reflect.Value) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.Decode(value)
				}
				d.Pos++
				value.SetInt(0)
				return nil
			}
			if typeID != 1 {
				return d.Decode(value)
			}
			numberType := (header >> 3) & 0x03
			if numberType != 1 {
				return d.Decode(value)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.Decode(value)
			}
			d.Pos++
			if d.Pos+byteCount > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+byteCount]
			d.Pos += byteCount
			var intVal int64
			for i, b := range data {
				intVal |= int64(b) << (8 * i)
			}
			if byteCount < 8 && (intVal&(1<<((byteCount*8)-1))) != 0 {
				intVal |= -1 << (byteCount * 8)
			}
			if !fitsSigned(intVal, bitSize) {
				return &UnsupportedError{"integer overflow"}
			}
			value.SetInt(intVal)
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return func(d *Decoder, value reflect.Value) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.Decode(value)
				}
				d.Pos++
				value.SetUint(0)
				return nil
			}
			if typeID != 1 {
				return d.Decode(value)
			}
			numberType := (header >> 3) & 0x03
			if numberType != 2 {
				return d.Decode(value)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.Decode(value)
			}
			d.Pos++
			if d.Pos+byteCount > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+byteCount]
			d.Pos += byteCount
			var uintVal uint64
			for i, b := range data {
				uintVal |= uint64(b) << (8 * i)
			}
			if !fitsUnsigned(uintVal, bitSize) {
				return &UnsupportedError{"integer overflow"}
			}
			value.SetUint(uintVal)
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(d *Decoder, value reflect.Value) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.Decode(value)
				}
				d.Pos++
				value.SetFloat(0)
				return nil
			}
			if typeID != 1 {
				return d.Decode(value)
			}
			numberType := (header >> 3) & 0x03
			if numberType != 0 {
				return d.Decode(value)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.Decode(value)
			}
			d.Pos++
			if d.Pos+byteCount > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+byteCount]
			d.Pos += byteCount
			var f64 float64
			switch byteCount {
			case 2:
				bits := binary.LittleEndian.Uint16(data)
				f64 = float64(math.Float32frombits(uint32(bits) << 16))
			case 4:
				bits := binary.LittleEndian.Uint32(data)
				f64 = float64(math.Float32frombits(bits))
			case 8:
				bits := binary.LittleEndian.Uint64(data)
				f64 = math.Float64frombits(bits)
			default:
				return d.Decode(value)
			}
			if kind == reflect.Float32 {
				value.SetFloat(float64(float32(f64)))
			} else {
				value.SetFloat(f64)
			}
			return nil
		}
	default:
		return nil
	}
}

func buildStructFieldDecoder(field *structField) fieldDecoderFunc {
	switch field.kind {
	case reflect.Bool:
		return func(d *Decoder, ptr unsafe.Pointer) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			if header&0x07 != 0 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			d.Pos++
			if header&0x08 == 0 {
				*(*bool)(ptr) = false
				return nil
			}
			*(*bool)(ptr) = header&0x10 != 0
			return nil
		}
	case reflect.String:
		return func(d *Decoder, ptr unsafe.Pointer) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeBits := header & 0x07
			if typeBits == 0 {
				if header&0x08 != 0 {
					return d.decodeStructFieldGeneric(field, ptr)
				}
				d.Pos++
				*(*string)(ptr) = ""
				return nil
			}
			if typeBits != 2 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			d.Pos++
			size, err := d.ReadCompressedUint()
			if err != nil {
				return err
			}
			data, err := d.ReadBytes(int(size))
			if err != nil {
				return err
			}
			*(*string)(ptr) = bytesToString(data)
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(d *Decoder, ptr unsafe.Pointer) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.decodeStructFieldGeneric(field, ptr)
				}
				d.Pos++
				switch field.kind {
				case reflect.Int:
					*(*int)(ptr) = 0
				case reflect.Int8:
					*(*int8)(ptr) = 0
				case reflect.Int16:
					*(*int16)(ptr) = 0
				case reflect.Int32:
					*(*int32)(ptr) = 0
				case reflect.Int64:
					*(*int64)(ptr) = 0
				}
				return nil
			}
			if typeID != 1 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			typeBits := (header >> 3) & 0x03
			if typeBits != 1 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			d.Pos++
			if d.Pos+byteCount > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+byteCount]
			d.Pos += byteCount
			var value int64
			for i, b := range data {
				value |= int64(b) << (8 * i)
			}
			if byteCount < 8 && (value&(1<<((byteCount*8)-1))) != 0 {
				value |= -1 << (byteCount * 8)
			}
			if !fitsSigned(value, field.bitSize) {
				return &UnsupportedError{"integer overflow"}
			}
			switch field.kind {
			case reflect.Int:
				*(*int)(ptr) = int(value)
			case reflect.Int8:
				*(*int8)(ptr) = int8(value)
			case reflect.Int16:
				*(*int16)(ptr) = int16(value)
			case reflect.Int32:
				*(*int32)(ptr) = int32(value)
			case reflect.Int64:
				*(*int64)(ptr) = value
			}
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return func(d *Decoder, ptr unsafe.Pointer) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.decodeStructFieldGeneric(field, ptr)
				}
				d.Pos++
				switch field.kind {
				case reflect.Uint:
					*(*uint)(ptr) = 0
				case reflect.Uint8:
					*(*uint8)(ptr) = 0
				case reflect.Uint16:
					*(*uint16)(ptr) = 0
				case reflect.Uint32:
					*(*uint32)(ptr) = 0
				case reflect.Uint64:
					*(*uint64)(ptr) = 0
				case reflect.Uintptr:
					*(*uintptr)(ptr) = 0
				}
				return nil
			}
			if typeID != 1 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			typeBits := (header >> 3) & 0x03
			if typeBits != 2 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			d.Pos++
			if d.Pos+byteCount > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+byteCount]
			d.Pos += byteCount
			var value uint64
			for i, b := range data {
				value |= uint64(b) << (8 * i)
			}
			if !fitsUnsigned(value, field.bitSize) {
				return &UnsupportedError{"integer overflow"}
			}
			switch field.kind {
			case reflect.Uint:
				*(*uint)(ptr) = uint(value)
			case reflect.Uint8:
				*(*uint8)(ptr) = uint8(value)
			case reflect.Uint16:
				*(*uint16)(ptr) = uint16(value)
			case reflect.Uint32:
				*(*uint32)(ptr) = uint32(value)
			case reflect.Uint64:
				*(*uint64)(ptr) = value
			case reflect.Uintptr:
				*(*uintptr)(ptr) = uintptr(value)
			}
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(d *Decoder, ptr unsafe.Pointer) error {
			if d.Pos >= len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			header := d.Data[d.Pos]
			typeID := header & 0x07
			if typeID == 0 {
				if header&0x08 != 0 {
					return d.decodeStructFieldGeneric(field, ptr)
				}
				d.Pos++
				if field.kind == reflect.Float32 {
					*(*float32)(ptr) = 0
				} else {
					*(*float64)(ptr) = 0
				}
				return nil
			}
			if typeID != 1 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			typeBits := (header >> 3) & 0x03
			if typeBits != 0 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			byteCount := d.GetByteCount((header >> 5) & 0x07)
			if byteCount == 0 {
				return d.decodeStructFieldGeneric(field, ptr)
			}
			d.Pos++
			var actual int
			switch byteCount {
			case 1:
				actual = 2
			case 2:
				actual = 2
			case 4:
				actual = 4
			case 8:
				actual = 8
			default:
				return d.decodeStructFieldGeneric(field, ptr)
			}
			if d.Pos+actual > len(d.Data) {
				return &UnsupportedError{"unexpected end of data"}
			}
			data := d.Data[d.Pos : d.Pos+actual]
			d.Pos += actual
			var f64 float64
			switch actual {
			case 2:
				uintVal := binary.LittleEndian.Uint16(data)
				f64 = float64(math.Float32frombits(uint32(uintVal) << 16))
			case 4:
				uintVal := binary.LittleEndian.Uint32(data)
				f64 = float64(math.Float32frombits(uintVal))
			case 8:
				uintVal := binary.LittleEndian.Uint64(data)
				f64 = math.Float64frombits(uintVal)
			default:
				return d.decodeStructFieldGeneric(field, ptr)
			}
			if field.kind == reflect.Float32 {
				*(*float32)(ptr) = float32(f64)
			} else {
				*(*float64)(ptr) = f64
			}
			return nil
		}
	default:
		return nil
	}
}

func fitsSigned(value int64, bits int) bool {
	if bits == 0 {
		bits = intBitSize
	}
	if bits >= 64 {
		return true
	}
	min := -(int64(1) << (bits - 1))
	max := (int64(1) << (bits - 1)) - 1
	return value >= min && value <= max
}

func fitsUnsigned(value uint64, bits int) bool {
	if bits == 0 {
		bits = intBitSize
	}
	if bits >= 64 {
		return true
	}
	max := (uint64(1) << bits) - 1
	return value <= max
}

// convertMapKeyValue converts a key value to the target map key type
func convertMapKeyValue(key reflect.Value, targetType reflect.Type, keyType byte) (reflect.Value, error) {
	result := reflect.New(targetType).Elem()

	switch targetType.Kind() {
	case reflect.String:
		if key.Kind() != reflect.String {
			if !key.Type().ConvertibleTo(reflect.TypeOf("")) {
				return reflect.Value{}, &UnsupportedError{"map key not convertible to string"}
			}
			key = key.Convert(reflect.TypeOf(""))
		}
		result.SetString(key.String())
		return result, nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var value int64
		if keyType == 1 {
			value = key.Int()
		} else if keyType == 2 {
			unsigned := key.Uint()
			if unsigned > (1<<63 - 1) {
				return reflect.Value{}, &UnsupportedError{"map key overflow"}
			}
			value = int64(unsigned)
		} else {
			return reflect.Value{}, &UnsupportedError{"map key type mismatch"}
		}
		if result.OverflowInt(value) {
			return reflect.Value{}, &UnsupportedError{"map key overflow"}
		}
		result.SetInt(value)
		return result, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var value uint64
		if keyType == 2 {
			value = key.Uint()
		} else if keyType == 1 {
			signed := key.Int()
			if signed < 0 {
				return reflect.Value{}, &UnsupportedError{"map key overflow"}
			}
			value = uint64(signed)
		} else {
			return reflect.Value{}, &UnsupportedError{"map key type mismatch"}
		}
		if result.OverflowUint(value) {
			return reflect.Value{}, &UnsupportedError{"map key overflow"}
		}
		result.SetUint(value)
		return result, nil

	default:
		// Try direct conversion
		if key.Type() == targetType {
			return key, nil
		}
		if key.Type().ConvertibleTo(targetType) {
			return key.Convert(targetType), nil
		}
		return reflect.Value{}, &UnsupportedError{"unsupported map key type: " + targetType.String()}
	}
}

// DecodeTypedArray decodes a typed array (optimized format for homogeneous arrays).
//
// BEVE typed array header:
//   - Bits 0-2: Type (4 = typed array)
//   - Bits 3-4: Group (0=float, 1=signed, 2=unsigned, 3=bool/string)
//   - Bits 5-7: Element size or flags
func (d *Decoder) DecodeTypedArray(v reflect.Value, header byte) error {
	group := (header >> 3) & 0x03
	size, err := d.ReadCompressedUint()
	if err != nil {
		return err
	}
	length, err := CheckedLength(size)
	if err != nil {
		return err
	}

	switch group {
	case 3:
		// boolean or string
		isString := ((header >> 5) & 0x01) == 1
		if isString {
			return d.decodeStringTypedArray(v, length)
		}
		return d.decodeBoolTypedArray(v, length)
	case 0:
		byteCount := d.GetByteCount((header >> 5) & 0x07)
		return d.decodeFloatTypedArray(v, length, byteCount)
	case 1:
		byteCount := d.GetByteCount((header >> 5) & 0x07)
		return d.decodeSignedTypedArray(v, length, byteCount)
	case 2:
		byteCount := d.GetByteCount((header >> 5) & 0x07)
		return d.decodeUnsignedTypedArray(v, length, byteCount)
	default:
		return &UnsupportedError{"unknown typed array"}
	}
}

// decodeBoolTypedArray decodes a bitpacked boolean array.
func (d *Decoder) decodeBoolTypedArray(v reflect.Value, length int) error {
	payload := (length + 7) / 8
	data, err := d.ReadBytes(payload)
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			if err := SetBoolElement(v.Index(i), bit == 1); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() < length {
			return &UnsupportedError{"array too small"}
		}
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			if err := SetBoolElement(v.Index(i), bit == 1); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		slice := make([]bool, length)
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			slice[i] = bit == 1
		}
		v.Set(reflect.ValueOf(slice))
		return nil
	default:
		return &UnsupportedError{"expected slice or array"}
	}
}

// decodeStringTypedArray decodes an array of strings.
func (d *Decoder) decodeStringTypedArray(v reflect.Value, length int) error {
	switch v.Kind() {
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.String && v.CanAddr() {
			slicePtr := (*[]string)(unsafe.Pointer(v.UnsafeAddr()))
			slice := *slicePtr
			if cap(slice) < length {
				slice = make([]string, length)
			} else {
				slice = slice[:length]
			}
			for i := 0; i < length; i++ {
				size, err := d.ReadCompressedUint()
				if err != nil {
					return err
				}
				data, err := d.ReadBytes(int(size))
				if err != nil {
					return err
				}
				slice[i] = bytesToString(data)
			}
			*slicePtr = slice
			return nil
		}
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			size, err := d.ReadCompressedUint()
			if err != nil {
				return err
			}
			data, err := d.ReadBytes(int(size))
			if err != nil {
				return err
			}
			v.Index(i).SetString(bytesToString(data))
		}
		return nil
	case reflect.Array:
		if v.Len() < length {
			return &UnsupportedError{"array too small"}
		}
		for i := 0; i < length; i++ {
			size, err := d.ReadCompressedUint()
			if err != nil {
				return err
			}
			data, err := d.ReadBytes(int(size))
			if err != nil {
				return err
			}
			v.Index(i).SetString(bytesToString(data))
		}
		return nil
	case reflect.Interface:
		slice := make([]string, length)
		for i := 0; i < length; i++ {
			size, err := d.ReadCompressedUint()
			if err != nil {
				return err
			}
			data, err := d.ReadBytes(int(size))
			if err != nil {
				return err
			}
			slice[i] = bytesToString(data)
		}
		v.Set(reflect.ValueOf(slice))
		return nil
	default:
		return &UnsupportedError{"expected slice or array"}
	}
}

// decodeSignedTypedArray decodes an array of signed integers.
func (d *Decoder) decodeSignedTypedArray(v reflect.Value, length, byteCount int) error {
	switch v.Kind() {
	case reflect.Slice:
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			data, err := d.ReadBytes(byteCount)
			if err != nil {
				return err
			}
			var val int64
			for j, b := range data {
				val |= int64(b) << (j * 8)
			}
			// Sign extend
			if byteCount < 8 && (val&(1<<((byteCount*8)-1))) != 0 {
				val |= -1 << (byteCount * 8)
			}
			if err := SetSignedElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() < length {
			return &UnsupportedError{"array too small"}
		}
		for i := 0; i < length; i++ {
			data, err := d.ReadBytes(byteCount)
			if err != nil {
				return err
			}
			var val int64
			for j, b := range data {
				val |= int64(b) << (j * 8)
			}
			// Sign extend
			if byteCount < 8 && (val&(1<<((byteCount*8)-1))) != 0 {
				val |= -1 << (byteCount * 8)
			}
			if err := SetSignedElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		// Create typed slice based on byte count
		switch byteCount {
		case 1:
			slice := make([]int8, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				val := int8(data[0])
				slice[i] = val
			}
			v.Set(reflect.ValueOf(slice))
		case 2:
			slice := make([]int16, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val int64
				for j, b := range data {
					val |= int64(b) << (j * 8)
				}
				if val&(1<<15) != 0 {
					val |= -1 << 16
				}
				slice[i] = int16(val)
			}
			v.Set(reflect.ValueOf(slice))
		case 4:
			slice := make([]int32, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val int64
				for j, b := range data {
					val |= int64(b) << (j * 8)
				}
				if val&(1<<31) != 0 {
					val |= -1 << 32
				}
				slice[i] = int32(val)
			}
			v.Set(reflect.ValueOf(slice))
		default:
			slice := make([]int64, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val int64
				for j, b := range data {
					val |= int64(b) << (j * 8)
				}
				slice[i] = val
			}
			v.Set(reflect.ValueOf(slice))
		}
		return nil
	default:
		return &UnsupportedError{"expected slice or array"}
	}
}

// decodeUnsignedTypedArray decodes an array of unsigned integers.
func (d *Decoder) decodeUnsignedTypedArray(v reflect.Value, length, byteCount int) error {
	switch v.Kind() {
	case reflect.Slice:
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			data, err := d.ReadBytes(byteCount)
			if err != nil {
				return err
			}
			var val uint64
			for j, b := range data {
				val |= uint64(b) << (j * 8)
			}
			if err := SetUnsignedElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() < length {
			return &UnsupportedError{"array too small"}
		}
		for i := 0; i < length; i++ {
			data, err := d.ReadBytes(byteCount)
			if err != nil {
				return err
			}
			var val uint64
			for j, b := range data {
				val |= uint64(b) << (j * 8)
			}
			if err := SetUnsignedElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		// Create typed slice based on byte count
		switch byteCount {
		case 1:
			slice := make([]uint8, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				slice[i] = data[0]
			}
			v.Set(reflect.ValueOf(slice))
		case 2:
			slice := make([]uint16, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val uint64
				for j, b := range data {
					val |= uint64(b) << (j * 8)
				}
				slice[i] = uint16(val)
			}
			v.Set(reflect.ValueOf(slice))
		case 4:
			slice := make([]uint32, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val uint64
				for j, b := range data {
					val |= uint64(b) << (j * 8)
				}
				slice[i] = uint32(val)
			}
			v.Set(reflect.ValueOf(slice))
		default:
			slice := make([]uint64, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(byteCount)
				if err != nil {
					return err
				}
				var val uint64
				for j, b := range data {
					val |= uint64(b) << (j * 8)
				}
				slice[i] = val
			}
			v.Set(reflect.ValueOf(slice))
		}
		return nil
	default:
		return &UnsupportedError{"expected slice or array"}
	}
}

// decodeFloatTypedArray decodes an array of floating-point numbers.
func (d *Decoder) decodeFloatTypedArray(v reflect.Value, length, byteCount int) error {
	switch v.Kind() {
	case reflect.Slice:
		if err := EnsureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			var val float64
			switch byteCount {
			case 4:
				data, err := d.ReadBytes(4)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint32(data)
				val = float64(math.Float32frombits(uintVal))
			case 8:
				data, err := d.ReadBytes(8)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint64(data)
				val = math.Float64frombits(uintVal)
			default:
				return &UnsupportedError{"unsupported float size"}
			}
			if err := SetFloatElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() < length {
			return &UnsupportedError{"array too small"}
		}
		for i := 0; i < length; i++ {
			var val float64
			switch byteCount {
			case 4:
				data, err := d.ReadBytes(4)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint32(data)
				val = float64(math.Float32frombits(uintVal))
			case 8:
				data, err := d.ReadBytes(8)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint64(data)
				val = math.Float64frombits(uintVal)
			default:
				return &UnsupportedError{"unsupported float size"}
			}
			if err := SetFloatElement(v.Index(i), val, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		// Create typed slice based on byte count
		switch byteCount {
		case 4:
			slice := make([]float32, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(4)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint32(data)
				slice[i] = math.Float32frombits(uintVal)
			}
			v.Set(reflect.ValueOf(slice))
		case 8:
			slice := make([]float64, length)
			for i := 0; i < length; i++ {
				data, err := d.ReadBytes(8)
				if err != nil {
					return err
				}
				uintVal := binary.LittleEndian.Uint64(data)
				slice[i] = math.Float64frombits(uintVal)
			}
			v.Set(reflect.ValueOf(slice))
		default:
			return &UnsupportedError{"unsupported float size"}
		}
		return nil
	default:
		return &UnsupportedError{"expected slice or array"}
	}
}

// DecodeExtension decodes an extension type (placeholder).
func (d *Decoder) DecodeExtension(v reflect.Value, header byte) error {
	return &UnsupportedError{"extensions not implemented"}
}
