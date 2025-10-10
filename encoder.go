package beve

import (
	"encoding/binary"
	"io"
	"math"
	"reflect"
	"strings"
	"sync"
)

// encoder handles the encoding of values to BEVE format
type encoder struct {
	w             io.Writer
	single        [1]byte
	uintScratch   [8]byte
	varintScratch [5]byte
	batchBuf      [256]byte // batch buffer for small writes
	batchLen      int       // current batch length
}

// newEncoder creates a new encoder
func newEncoder(w io.Writer) *encoder {
	return &encoder{w: w}
}

// encode encodes a reflect.Value to BEVE
func (e *encoder) encode(v reflect.Value) error {
	if !v.IsValid() {
		return e.encodeNull()
	}

	if isRawMessageType(v.Type()) {
		return e.encodeRawMessage(v.Bytes())
	}

	if v.CanInterface() {
		if bm, ok := v.Interface().(BinaryMarshaler); ok {
			return e.encodeBinaryMarshaler(bm)
		}
	}

	if v.Kind() != reflect.Ptr && v.CanAddr() {
		addr := v.Addr()
		if addr.CanInterface() {
			if bm, ok := addr.Interface().(BinaryMarshaler); ok {
				return e.encodeBinaryMarshaler(bm)
			}
		}
	}

	switch v.Kind() {
	case reflect.Invalid:
		return e.encodeNull()
	case reflect.Bool:
		return e.encodeBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return e.encodeInt(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return e.encodeUint(v.Uint())
	case reflect.Float32, reflect.Float64:
		return e.encodeFloat(v.Float(), v.Kind())
	case reflect.String:
		return e.encodeString(v.String())
	case reflect.Slice, reflect.Array:
		return e.encodeSlice(v)
	case reflect.Map:
		return e.encodeMap(v)
	case reflect.Struct:
		return e.encodeStruct(v)
	case reflect.Interface:
		if v.IsNil() {
			return e.encodeNull()
		}
		return e.encode(v.Elem())
	case reflect.Ptr:
		if v.IsNil() {
			return e.encodeNull()
		}
		return e.encode(v.Elem())
	default:
		return &UnsupportedError{"unsupported type: " + v.Type().String()}
	}
}

// encodeNull encodes a null value
func (e *encoder) encodeNull() error {
	return e.writeByte(0x00)
}

// encodeBool encodes a boolean
func (e *encoder) encodeBool(b bool) error {
	if b {
		return e.writeByte(0x18) // true
	}
	return e.writeByte(0x08) // false
}

// encodeInt encodes a signed integer
func (e *encoder) encodeInt(i int64) error {
	// Determine byte count based on value
	var byteCount int
	var byteCountBits byte

	if i >= -128 && i <= 127 {
		byteCount = 1
		byteCountBits = 0
	} else if i >= -32768 && i <= 32767 {
		byteCount = 2
		byteCountBits = 1
	} else if i >= -2147483648 && i <= 2147483647 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}

	header := byte(0x01) | (1 << 3) | (byteCountBits << 5) // type=1 (signed), base header
	if err := e.writeByte(header); err != nil {
		return err
	}

	// Write the integer in little-endian
	for j := 0; j < byteCount; j++ {
		b := byte(i >> (j * 8))
		if err := e.writeByte(b); err != nil {
			return err
		}
	}

	return nil
}

func (e *encoder) encodeRawMessage(data []byte) error {
	if len(data) == 0 {
		return &UnsupportedError{"RawMessage payload must contain a value"}
	}
	return e.writeBytes(data)
}

func (e *encoder) encodeBinaryMarshaler(m BinaryMarshaler) error {
	data, err := m.MarshalBEVE()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return &UnsupportedError{"BinaryMarshaler returned empty payload"}
	}
	return e.writeBytes(data)
}

// encodeUint encodes an unsigned integer
func (e *encoder) encodeUint(u uint64) error {
	var byteCount int
	var byteCountBits byte

	if u <= 255 {
		byteCount = 1
		byteCountBits = 0
	} else if u <= 65535 {
		byteCount = 2
		byteCountBits = 1
	} else if u <= 4294967295 {
		byteCount = 4
		byteCountBits = 2
	} else {
		byteCount = 8
		byteCountBits = 3
	}

	header := byte(0x01) | (2 << 3) | (byteCountBits << 5) // type=2 (unsigned)
	if err := e.writeByte(header); err != nil {
		return err
	}

	for j := 0; j < byteCount; j++ {
		b := byte(u >> (j * 8))
		if err := e.writeByte(b); err != nil {
			return err
		}
	}

	return nil
}

// encodeFloat encodes a floating point number
func (e *encoder) encodeFloat(f float64, kind reflect.Kind) error {
	var header byte
	var bytes []byte

	if kind == reflect.Float32 {
		val := float32(f)
		uintVal := math.Float32bits(val)
		bytes = make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, uintVal)
		header = 0x01 | (0 << 3) | (2 << 5) // float, 4 bytes
	} else {
		uintVal := math.Float64bits(f)
		bytes = make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uintVal)
		header = 0x01 | (0 << 3) | (3 << 5) // float, 8 bytes
	}

	if err := e.writeByte(header); err != nil {
		return err
	}

	// Write in little-endian
	for _, b := range bytes {
		if err := e.writeByte(b); err != nil {
			return err
		}
	}

	return nil
}

// encodeString encodes a string
func (e *encoder) encodeString(s string) error {
	header := byte(0x02) // string type
	if err := e.writeByte(header); err != nil {
		return err
	}

	// Write size as compressed unsigned integer
	size := uint64(len(s))
	if err := e.writeCompressedUint(size); err != nil {
		return err
	}

	return e.writeStringBytes(s)
}

// encodeSlice encodes a slice or array
func (e *encoder) encodeSlice(v reflect.Value) error {
	if info, ok := getTypedArrayInfo(v.Type().Elem()); ok {
		return e.encodeTypedArray(v, info)
	}

	length := v.Len()
	// Generic array header (type=5)
	header := byte(0x85)
	if err := e.writeByte(header); err != nil {
		return err
	}

	if err := e.writeCompressedUint(uint64(length)); err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		if err := e.encode(v.Index(i)); err != nil {
			return err
		}
	}

	return nil
}

// encodeMap encodes a map
func (e *encoder) encodeMap(v reflect.Value) error {
	keyType, err := determineMapKeyType(v.Type().Key())
	if err != nil {
		return err
	}

	header := byte(0x03) | (keyType << 3)
	if err := e.writeByte(header); err != nil {
		return err
	}

	size := uint64(v.Len())
	if err := e.writeCompressedUint(size); err != nil {
		return err
	}

	for _, key := range v.MapKeys() {
		if err := e.writeMapKey(key, keyType); err != nil {
			return err
		}

		if err := e.encode(v.MapIndex(key)); err != nil {
			return err
		}
	}

	return nil
}

// encodeStruct encodes a struct
func (e *encoder) encodeStruct(v reflect.Value) error {
	if err := e.writeByte(0x03); err != nil {
		return err
	}

	info := getStructInfo(v.Type())
	count := 0
	for _, fieldInfo := range info.fields {
		field := v.FieldByIndex(fieldInfo.index)
		if fieldInfo.omitEmpty && isEmptyValue(field) {
			continue
		}
		count++
	}

	if err := e.writeCompressedUint(uint64(count)); err != nil {
		return err
	}

	for _, fieldInfo := range info.fields {
		field := v.FieldByIndex(fieldInfo.index)
		if fieldInfo.omitEmpty && isEmptyValue(field) {
			continue
		}
		if err := e.writeCompressedUint(uint64(len(fieldInfo.name))); err != nil {
			return err
		}
		if err := e.writeStringBytes(fieldInfo.name); err != nil {
			return err
		}
		if err := e.encode(field); err != nil {
			return err
		}
	}

	return nil
}

// encodeSignedArray encodes signed integer array with bulk operations
func (e *encoder) encodeSignedArray(v reflect.Value, length, byteCount int) error {
	// For large arrays, use bulk buffer writing
	if length > 16 && byteCount <= 8 {
		totalBytes := length * byteCount
		buf := acquireBytes(totalBytes)
		defer releaseBytes(buf)

		offset := 0
		for i := 0; i < length; i++ {
			val := uint64(v.Index(i).Int())
			switch byteCount {
			case 1:
				buf[offset] = byte(val)
			case 2:
				binary.LittleEndian.PutUint16(buf[offset:], uint16(val))
			case 4:
				binary.LittleEndian.PutUint32(buf[offset:], uint32(val))
			case 8:
				binary.LittleEndian.PutUint64(buf[offset:], val)
			}
			offset += byteCount
		}
		return e.writeBytes(buf[:totalBytes])
	}

	// Small arrays: inline writes
	for i := 0; i < length; i++ {
		if err := e.writeIntBytes(v.Index(i).Int(), byteCount); err != nil {
			return err
		}
	}
	return nil
}

// encodeUnsignedArray encodes unsigned integer array with bulk operations
func (e *encoder) encodeUnsignedArray(v reflect.Value, length, byteCount int) error {
	// For large arrays, use bulk buffer writing
	if length > 16 && byteCount <= 8 {
		totalBytes := length * byteCount
		buf := acquireBytes(totalBytes)
		defer releaseBytes(buf)

		offset := 0
		for i := 0; i < length; i++ {
			val := v.Index(i).Uint()
			switch byteCount {
			case 1:
				buf[offset] = byte(val)
			case 2:
				binary.LittleEndian.PutUint16(buf[offset:], uint16(val))
			case 4:
				binary.LittleEndian.PutUint32(buf[offset:], uint32(val))
			case 8:
				binary.LittleEndian.PutUint64(buf[offset:], val)
			}
			offset += byteCount
		}
		return e.writeBytes(buf[:totalBytes])
	}

	// Small arrays: inline writes
	for i := 0; i < length; i++ {
		if err := e.writeUintBytes(v.Index(i).Uint(), byteCount); err != nil {
			return err
		}
	}
	return nil
}

func (e *encoder) encodeTypedArray(v reflect.Value, info typedArrayInfo) error {
	length := v.Len()
	if err := e.writeByte(info.header); err != nil {
		return err
	}
	if err := e.writeCompressedUint(uint64(length)); err != nil {
		return err
	}

	switch info.category {
	case typedArrayBool:
		if length == 0 {
			return nil
		}
		payloadLen := (length + 7) / 8
		buf := acquireBytes(payloadLen)
		for i := 0; i < payloadLen; i++ {
			buf[i] = 0
		}
		for i := 0; i < length; i++ {
			if v.Index(i).Bool() {
				buf[i>>3] |= 1 << (uint(i) & 7)
			}
		}
		err := e.writeBytes(buf[:payloadLen])
		releaseBytes(buf)
		return err
	case typedArraySigned:
		return e.encodeSignedArray(v, length, info.byteCount)
	case typedArrayUnsigned:
		return e.encodeUnsignedArray(v, length, info.byteCount)
	case typedArrayFloat:
		for i := 0; i < length; i++ {
			fv := v.Index(i).Float()
			if info.byteCount == 4 {
				bits := math.Float32bits(float32(fv))
				if err := e.writeUintBytes(uint64(bits), info.byteCount); err != nil {
					return err
				}
			} else if info.byteCount == 8 {
				bits := math.Float64bits(fv)
				if err := e.writeUintBytes(bits, info.byteCount); err != nil {
					return err
				}
			} else {
				return &UnsupportedError{"unsupported float byte count"}
			}
		}
	case typedArrayString:
		for i := 0; i < length; i++ {
			s := v.Index(i).String()
			if err := e.writeCompressedUint(uint64(len(s))); err != nil {
				return err
			}
			if err := e.writeStringBytes(s); err != nil {
				return err
			}
		}
	default:
		return &UnsupportedError{"unsupported typed array category"}
	}

	return nil
}

// writeByte writes a single byte
//
//go:inline
func (e *encoder) writeByte(b byte) error {
	if bw, ok := e.w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}
	e.single[0] = b
	_, err := e.w.Write(e.single[:])
	return err
}

//go:inline
func (e *encoder) writeBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	_, err := e.w.Write(data)
	return err
}

func (e *encoder) writeStringBytes(s string) error {
	if len(s) == 0 {
		return nil
	}
	// For string writer optimization
	if sw, ok := e.w.(io.StringWriter); ok {
		_, err := sw.WriteString(s)
		return err
	}
	// Use zero-copy conversion (unsafe but safe in our context)
	// The data is immediately written and not retained
	_, err := e.w.Write(stringToBytes(s))
	return err
}

//go:inline
func (e *encoder) writeIntBytes(value int64, count int) error {
	binary.LittleEndian.PutUint64(e.uintScratch[:], uint64(value))
	return e.writeBytes(e.uintScratch[:count])
}

//go:inline
func (e *encoder) writeUintBytes(value uint64, count int) error {
	binary.LittleEndian.PutUint64(e.uintScratch[:], value)
	return e.writeBytes(e.uintScratch[:count])
}

// writeCompressedUint writes a compressed unsigned integer
//
//go:inline
func (e *encoder) writeCompressedUint(n uint64) error {
	// Fast path for small numbers (most common case)
	if n < 64 {
		return e.writeByte(byte(n << 2))
	}
	if n < 16384 {
		e.varintScratch[0] = byte(0x01 | ((n >> 8) << 2))
		e.varintScratch[1] = byte(n)
		return e.writeBytes(e.varintScratch[:2])
	}
	if n < 1073741824 {
		e.varintScratch[0] = byte(0x02 | ((n >> 16) << 2))
		e.varintScratch[1] = byte(n >> 8)
		e.varintScratch[2] = byte(n)
		return e.writeBytes(e.varintScratch[:3])
	}
	e.varintScratch[0] = byte(0x03 | ((n >> 24) << 2))
	e.varintScratch[1] = byte(n >> 16)
	e.varintScratch[2] = byte(n >> 8)
	e.varintScratch[3] = byte(n)
	return e.writeBytes(e.varintScratch[:4])
}

type typedArrayCategory int

const (
	typedArrayBool typedArrayCategory = iota
	typedArraySigned
	typedArrayUnsigned
	typedArrayFloat
	typedArrayString
)

type typedArrayInfo struct {
	category  typedArrayCategory
	header    byte
	byteCount int
	elemKind  reflect.Kind
}

const maxScratchSize = 1 << 16 // 64KiB per pooled buffer

var byteSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024)
	},
}

type structField struct {
	name      string
	index     []int
	omitEmpty bool
}

type structInfo struct {
	fields   []structField
	fieldMap map[string]structField
}

var structFieldCache sync.Map // map[reflect.Type]*structInfo

func acquireBytes(size int) []byte {
	if size <= 0 {
		return nil
	}
	if size > maxScratchSize {
		return make([]byte, size)
	}
	buf := byteSlicePool.Get().([]byte)
	if cap(buf) < size {
		buf = make([]byte, size)
	}
	return buf[:size]
}

func releaseBytes(buf []byte) {
	if buf == nil {
		return
	}
	if cap(buf) > maxScratchSize {
		return
	}
	full := buf[:cap(buf)]
	for i := range full {
		full[i] = 0
	}
	byteSlicePool.Put(full)
}

func gatherStructFields(t reflect.Type, parentIndex []int, visited map[reflect.Type]bool, fields *[]structField, fieldMap map[string]structField) {
	if visited[t] {
		return
	}
	visited[t] = true
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" && !field.Anonymous {
			continue
		}
		name, opts, skip := parseBeveTag(field.Tag.Get("beve"))
		if skip {
			continue
		}
		inline := opts.Contains("inline") || (field.Anonymous && name == "")
		fieldType := field.Type
		index := append(parentIndex[:len(parentIndex):len(parentIndex)], field.Index...)
		if inline {
			base := fieldType
			if base.Kind() == reflect.Ptr {
				base = base.Elem()
			}
			if base.Kind() == reflect.Struct {
				gatherStructFields(base, index, visited, fields, fieldMap)
			}
			continue
		}
		if name == "" {
			name = field.Name
		}
		info := structField{
			name:      name,
			index:     index,
			omitEmpty: opts.Contains("omitempty"),
		}
		*fields = append(*fields, info)
		if _, exists := fieldMap[name]; !exists {
			fieldMap[name] = info
		}
	}
}

func getStructInfo(t reflect.Type) *structInfo {
	if cached, ok := structFieldCache.Load(t); ok {
		return cached.(*structInfo)
	}
	fields := make([]structField, 0, t.NumField())
	fieldMap := make(map[string]structField, t.NumField())
	visited := make(map[reflect.Type]bool)
	gatherStructFields(t, nil, visited, &fields, fieldMap)
	res := &structInfo{fields: fields, fieldMap: fieldMap}
	structFieldCache.Store(t, res)
	return res
}

const typedArrayStringHeader = byte(0x04 | (3 << 3) | (1 << 5))

func getTypedArrayInfo(t reflect.Type) (typedArrayInfo, bool) {
	switch t.Kind() {
	case reflect.Bool:
		return typedArrayInfo{
			category: typedArrayBool,
			header:   0x04 | (3 << 3),
			elemKind: reflect.Bool,
		}, true
	case reflect.Int8:
		return typedArrayInfo{
			category:  typedArraySigned,
			header:    typedArrayHeader(1, 1),
			byteCount: 1,
			elemKind:  reflect.Int8,
		}, true
	case reflect.Int16:
		return typedArrayInfo{
			category:  typedArraySigned,
			header:    typedArrayHeader(1, 2),
			byteCount: 2,
			elemKind:  reflect.Int16,
		}, true
	case reflect.Int32:
		return typedArrayInfo{
			category:  typedArraySigned,
			header:    typedArrayHeader(1, 4),
			byteCount: 4,
			elemKind:  reflect.Int32,
		}, true
	case reflect.Int64:
		return typedArrayInfo{
			category:  typedArraySigned,
			header:    typedArrayHeader(1, 8),
			byteCount: 8,
			elemKind:  reflect.Int64,
		}, true
	case reflect.Uint8:
		return typedArrayInfo{
			category:  typedArrayUnsigned,
			header:    typedArrayHeader(2, 1),
			byteCount: 1,
			elemKind:  reflect.Uint8,
		}, true
	case reflect.Uint16:
		return typedArrayInfo{
			category:  typedArrayUnsigned,
			header:    typedArrayHeader(2, 2),
			byteCount: 2,
			elemKind:  reflect.Uint16,
		}, true
	case reflect.Uint32:
		return typedArrayInfo{
			category:  typedArrayUnsigned,
			header:    typedArrayHeader(2, 4),
			byteCount: 4,
			elemKind:  reflect.Uint32,
		}, true
	case reflect.Uint64:
		return typedArrayInfo{
			category:  typedArrayUnsigned,
			header:    typedArrayHeader(2, 8),
			byteCount: 8,
			elemKind:  reflect.Uint64,
		}, true
	case reflect.Float32:
		return typedArrayInfo{
			category:  typedArrayFloat,
			header:    typedArrayHeader(0, 4),
			byteCount: 4,
			elemKind:  reflect.Float32,
		}, true
	case reflect.Float64:
		return typedArrayInfo{
			category:  typedArrayFloat,
			header:    typedArrayHeader(0, 8),
			byteCount: 8,
			elemKind:  reflect.Float64,
		}, true
	case reflect.String:
		return typedArrayInfo{
			category: typedArrayString,
			header:   typedArrayStringHeader,
			elemKind: reflect.String,
		}, true
	default:
		return typedArrayInfo{}, false
	}
}

func typedArrayHeader(typeGroup byte, byteCount int) byte {
	indicator, ok := byteCountIndicator(byteCount)
	if !ok {
		return 0
	}
	return byte(0x04) | (typeGroup << 3) | (indicator << 5)
}

func byteCountIndicator(n int) (byte, bool) {
	switch n {
	case 1:
		return 0, true
	case 2:
		return 1, true
	case 4:
		return 2, true
	case 8:
		return 3, true
	case 16:
		return 4, true
	default:
		return 0, false
	}
}

func determineMapKeyType(t reflect.Type) (byte, error) {
	switch t.Kind() {
	case reflect.String:
		return 0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return 1, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return 2, nil
	default:
		return 0, &UnsupportedError{"unsupported map key type: " + t.String()}
	}
}

func (e *encoder) writeMapKey(key reflect.Value, keyType byte) error {
	switch keyType {
	case 0:
		if key.Kind() != reflect.String {
			if !key.Type().ConvertibleTo(reflect.TypeOf("")) {
				return &UnsupportedError{"map key not convertible to string"}
			}
			key = key.Convert(reflect.TypeOf(""))
		}
		s := key.String()
		if err := e.writeCompressedUint(uint64(len(s))); err != nil {
			return err
		}
		return e.writeStringBytes(s)
	case 1:
		if !isSignedIntegerKind(key.Kind()) {
			if !key.Type().ConvertibleTo(reflect.TypeOf(int64(0))) {
				return &UnsupportedError{"map key not convertible to int64"}
			}
			key = key.Convert(reflect.TypeOf(int64(0)))
		}
		return e.writeIntBytes(key.Int(), 8)
	case 2:
		if !isUnsignedIntegerKind(key.Kind()) {
			if !key.Type().ConvertibleTo(reflect.TypeOf(uint64(0))) {
				return &UnsupportedError{"map key not convertible to uint64"}
			}
			key = key.Convert(reflect.TypeOf(uint64(0)))
		}
		return e.writeUintBytes(key.Uint(), 8)
	default:
		return &UnsupportedError{"unsupported map key type"}
	}
}

type tagOptions map[string]struct{}

func (o tagOptions) Contains(opt string) bool {
	if o == nil {
		return false
	}
	_, ok := o[opt]
	return ok
}

func parseBeveTag(tag string) (string, tagOptions, bool) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return "", nil, false
	}
	if tag == "-" {
		return "", nil, true
	}
	parts := strings.Split(tag, ",")
	name := strings.TrimSpace(parts[0])
	if name == "-" {
		return "", nil, true
	}
	opts := make(tagOptions)
	for _, opt := range parts[1:] {
		opt = strings.TrimSpace(opt)
		if opt != "" {
			opts[opt] = struct{}{}
		}
	}
	return name, opts, false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	case reflect.Invalid:
		return true
	default:
		return false
	}
}

func isSignedIntegerKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func isUnsignedIntegerKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	default:
		return false
	}
}
