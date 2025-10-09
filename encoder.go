package beve

import (
	"encoding/binary"
	"io"
	"math"
	"reflect"
	"strings"
)

// encoder handles the encoding of values to BEVE format
type encoder struct {
	w io.Writer
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

	// Write string bytes
	for _, b := range []byte(s) {
		if err := e.writeByte(b); err != nil {
			return err
		}
	}

	return nil
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

	t := v.Type()
	type fieldInfo struct {
		name  string
		value reflect.Value
	}
	var fields []fieldInfo

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}

		name, opts, skip := parseBeveTag(field.Tag.Get("beve"))
		if skip {
			continue
		}
		if name == "" {
			name = field.Name
		}

		value := v.Field(i)
		if opts != nil && opts.Contains("omitempty") && isEmptyValue(value) {
			continue
		}

		fields = append(fields, fieldInfo{name: name, value: value})
	}

	if err := e.writeCompressedUint(uint64(len(fields))); err != nil {
		return err
	}

	for _, field := range fields {
		if err := e.writeCompressedUint(uint64(len(field.name))); err != nil {
			return err
		}
		if err := e.writeBytes([]byte(field.name)); err != nil {
			return err
		}
		if err := e.encode(field.value); err != nil {
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
		buf := make([]byte, payloadLen)
		for i := 0; i < length; i++ {
			if v.Index(i).Bool() {
				buf[i>>3] |= 1 << (uint(i) & 7)
			}
		}
		return e.writeBytes(buf)
	case typedArraySigned:
		for i := 0; i < length; i++ {
			if err := e.writeIntBytes(v.Index(i).Int(), info.byteCount); err != nil {
				return err
			}
		}
	case typedArrayUnsigned:
		for i := 0; i < length; i++ {
			if err := e.writeUintBytes(v.Index(i).Uint(), info.byteCount); err != nil {
				return err
			}
		}
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
			if err := e.writeBytes([]byte(s)); err != nil {
				return err
			}
		}
	default:
		return &UnsupportedError{"unsupported typed array category"}
	}

	return nil
}

// writeByte writes a single byte
func (e *encoder) writeByte(b byte) error {
	if bw, ok := e.w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}
	_, err := e.w.Write([]byte{b})
	return err
}

func (e *encoder) writeBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	_, err := e.w.Write(data)
	return err
}

func (e *encoder) writeIntBytes(value int64, count int) error {
	u := uint64(value)
	for i := 0; i < count; i++ {
		if err := e.writeByte(byte(u >> (i * 8))); err != nil {
			return err
		}
	}
	return nil
}

func (e *encoder) writeUintBytes(value uint64, count int) error {
	for i := 0; i < count; i++ {
		if err := e.writeByte(byte(value >> (i * 8))); err != nil {
			return err
		}
	}
	return nil
}

// writeCompressedUint writes a compressed unsigned integer
func (e *encoder) writeCompressedUint(n uint64) error {
	if n < 64 {
		return e.writeByte(byte(n << 2))
	} else if n < 16384 {
		if err := e.writeByte(byte(0x01 | ((n >> 8) << 2))); err != nil {
			return err
		}
		return e.writeByte(byte(n & 0xFF))
	} else if n < 1073741824 {
		if err := e.writeByte(byte(0x02 | ((n >> 16) << 2))); err != nil {
			return err
		}
		if err := e.writeByte(byte((n >> 8) & 0xFF)); err != nil {
			return err
		}
		return e.writeByte(byte(n & 0xFF))
	} else {
		if err := e.writeByte(byte(0x03 | ((n >> 24) << 2))); err != nil {
			return err
		}
		if err := e.writeByte(byte((n >> 16) & 0xFF)); err != nil {
			return err
		}
		if err := e.writeByte(byte((n >> 8) & 0xFF)); err != nil {
			return err
		}
		return e.writeByte(byte(n & 0xFF))
	}
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
		return e.writeBytes([]byte(s))
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
