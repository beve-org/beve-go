package beve

import (
	"encoding/binary"
	"math"
	"reflect"
)

// decoder handles the decoding of BEVE format to values
type decoder struct {
	data []byte
	pos  int
}

// newDecoder creates a new decoder
func newDecoder(data []byte) *decoder {
	return &decoder{data: data, pos: 0}
}

// decode decodes BEVE data into a reflect.Value
func (d *decoder) decode(v reflect.Value) error {
	header, err := d.readByte()
	if err != nil {
		return err
	}

	switch header & 0x07 { // type bits
	case 0: // null or bool
		if header&0x08 != 0 { // bool bit
			b := header&0x10 != 0 // true/false bit
			return d.setBool(v, b)
		}
		return d.setNil(v)
	case 1: // number
		return d.decodeNumber(v, header)
	case 2: // string
		return d.decodeString(v)
	case 3: // object
		return d.decodeObject(v, header)
	case 4: // typed array
		return d.decodeTypedArray(v, header)
	case 5: // generic array
		return d.decodeGenericArray(v)
	case 6: // extension
		return d.decodeExtension(v, header)
	default:
		return &UnsupportedError{"unknown type"}
	}
}

// decodeNumber decodes a number
func (d *decoder) decodeNumber(v reflect.Value, header byte) error {
	typeBits := (header >> 3) & 0x03
	byteCountBits := (header >> 5) & 0x07

	byteCount := d.getByteCount(byteCountBits)

	switch typeBits {
	case 0: // float
		return d.decodeFloat(v, byteCount)
	case 1: // signed int
		return d.decodeInt(v, byteCount)
	case 2: // unsigned int
		return d.decodeUint(v, byteCount)
	}
	return &UnsupportedError{"invalid number type"}
}

// decodeFloat decodes a float
func (d *decoder) decodeFloat(v reflect.Value, byteCount int) error {
	var val interface{}
	switch byteCount {
	case 1: // bfloat16 - not fully supported, treat as float32
		data, err := d.readBytes(2)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint16(data)
		// Convert bfloat16 to float32 approximation
		val = math.Float32frombits(uint32(uintVal) << 16)
	case 4: // float32
		data, err := d.readBytes(4)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint32(data)
		val = math.Float32frombits(uintVal)
	case 8: // float64
		data, err := d.readBytes(8)
		if err != nil {
			return err
		}
		uintVal := binary.LittleEndian.Uint64(data)
		val = math.Float64frombits(uintVal)
	default:
		return &UnsupportedError{"unsupported float size"}
	}

	return d.setValue(v, val)
}

// decodeInt decodes a signed integer
func (d *decoder) decodeInt(v reflect.Value, byteCount int) error {
	data, err := d.readBytes(byteCount)
	if err != nil {
		return err
	}

	var val int64
	for i, b := range data {
		val |= int64(b) << (i * 8)
	}

	// Sign extend if necessary
	if byteCount < 8 && (val&(1<<((byteCount*8)-1))) != 0 {
		val |= -1 << (byteCount * 8)
	}

	return d.setValue(v, val)
}

// decodeUint decodes an unsigned integer
func (d *decoder) decodeUint(v reflect.Value, byteCount int) error {
	data, err := d.readBytes(byteCount)
	if err != nil {
		return err
	}

	var val uint64
	for i, b := range data {
		val |= uint64(b) << (i * 8)
	}

	return d.setValue(v, val)
}

// decodeString decodes a string
func (d *decoder) decodeString(v reflect.Value) error {
	size, err := d.readCompressedUint()
	if err != nil {
		return err
	}

	data, err := d.readBytes(int(size))
	if err != nil {
		return err
	}

	str := string(data)
	return d.setValue(v, str)
}

// decodeObject decodes an object/map
func (d *decoder) decodeObject(v reflect.Value, header byte) error {
	keyType := (header >> 3) & 0x03

	size, err := d.readCompressedUint()
	if err != nil {
		return err
	}

	if v.Kind() == reflect.Map {
		return d.decodeMap(v, keyType, int(size))
	} else if v.Kind() == reflect.Struct {
		return d.decodeStruct(v, keyType, int(size))
	}

	return &UnsupportedError{"object type not supported"}
}

// decodeMap decodes into a map
func (d *decoder) decodeMap(v reflect.Value, keyType byte, size int) error {
	if v.IsNil() {
		// Create map if nil
		mapType := v.Type()
		v.Set(reflect.MakeMap(mapType))
	}

	for i := 0; i < size; i++ {
		key, err := d.readKey(keyType)
		if err != nil {
			return err
		}

		elemType := v.Type().Elem()
		elemValue := reflect.New(elemType).Elem()
		if err := d.decode(elemValue); err != nil {
			return err
		}

		v.SetMapIndex(key, elemValue)
	}

	return nil
}

// decodeStruct decodes into a struct
func (d *decoder) decodeStruct(v reflect.Value, keyType byte, size int) error {
	t := v.Type()
	fieldMap := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		fieldName := field.Name
		if tag := field.Tag.Get("beve"); tag != "" {
			fieldName = tag
		}
		fieldMap[fieldName] = i
	}

	for i := 0; i < size; i++ {
		key, err := d.readKey(keyType)
		if err != nil {
			return err
		}

		keyStr := key.String()
		if fieldIndex, ok := fieldMap[keyStr]; ok {
			field := v.Field(fieldIndex)
			if err := d.decode(field); err != nil {
				return err
			}
		} else {
			// Skip unknown fields
			if err := d.skipValue(); err != nil {
				return err
			}
		}
	}

	return nil
}

// decodeTypedArray decodes a typed array
func (d *decoder) decodeTypedArray(v reflect.Value, header byte) error {
	group := (header >> 3) & 0x03
	size, err := d.readCompressedUint()
	if err != nil {
		return err
	}
	length, err := checkedLength(size)
	if err != nil {
		return err
	}

	switch group {
	case 3:
		// boolean or string
		isString := ((header >> 5) & 0x01) == 1
		if isString {
			return &UnsupportedError{"typed string arrays not implemented"}
		}
		return d.decodeBoolTypedArray(v, length)
	case 0:
		byteCount := d.getByteCount((header >> 5) & 0x07)
		return d.decodeFloatTypedArray(v, length, byteCount)
	case 1:
		byteCount := d.getByteCount((header >> 5) & 0x07)
		return d.decodeSignedTypedArray(v, length, byteCount)
	case 2:
		byteCount := d.getByteCount((header >> 5) & 0x07)
		return d.decodeUnsignedTypedArray(v, length, byteCount)
	default:
		return &UnsupportedError{"unknown typed array"}
	}
}

// decodeGenericArray decodes a generic array
func (d *decoder) decodeGenericArray(v reflect.Value) error {
	size, err := d.readCompressedUint()
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		length := int(size)
		if v.IsNil() {
			v.Set(reflect.MakeSlice(v.Type(), length, length))
		} else if v.Len() != length {
			if v.Cap() >= length {
				v.Set(v.Slice(0, length))
			} else {
				v.Set(reflect.MakeSlice(v.Type(), length, length))
			}
		}
		for i := 0; i < length; i++ {
			if err := d.decode(v.Index(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		total := int(size)
		limit := total
		if limit > v.Len() {
			limit = v.Len()
		}
		for i := 0; i < limit; i++ {
			if err := d.decode(v.Index(i)); err != nil {
				return err
			}
		}
		for i := limit; i < total; i++ {
			if err := d.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		length := int(size)
		slice := make([]interface{}, length)
		for i := 0; i < length; i++ {
			var elem interface{}
			elemValue := reflect.ValueOf(&elem).Elem()
			if err := d.decode(elemValue); err != nil {
				return err
			}
			slice[i] = elem
		}
		v.Set(reflect.ValueOf(slice))
		return nil
	default:
		return &UnsupportedError{"array decode target must be slice, array, or interface"}
	}
}

// decodeExtension decodes an extension
func (d *decoder) decodeExtension(v reflect.Value, header byte) error {
	// Extensions not implemented yet
	return &UnsupportedError{"extensions not implemented"}
}

// readKey reads a key based on key type
func (d *decoder) readKey(keyType byte) (reflect.Value, error) {
	switch keyType {
	case 0: // string
		size, err := d.readCompressedUint()
		if err != nil {
			return reflect.Value{}, err
		}
		data, err := d.readBytes(int(size))
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(string(data)), nil
	case 1: // signed int
		// Assume int64 for simplicity
		data, err := d.readBytes(8)
		if err != nil {
			return reflect.Value{}, err
		}
		var val int64
		for i, b := range data {
			val |= int64(b) << (i * 8)
		}
		return reflect.ValueOf(val), nil
	case 2: // unsigned int
		data, err := d.readBytes(8)
		if err != nil {
			return reflect.Value{}, err
		}
		var val uint64
		for i, b := range data {
			val |= uint64(b) << (i * 8)
		}
		return reflect.ValueOf(val), nil
	}
	return reflect.Value{}, &UnsupportedError{"unsupported key type"}
}

// skipValue skips a value in the stream
func (d *decoder) skipValue() error {
	header, err := d.readByte()
	if err != nil {
		return err
	}

	// Simple skip - just advance position
	// This is not complete, but for basic functionality
	switch header & 0x07 {
	case 0: // null/bool - already consumed
	case 1: // number
		byteCountBits := (header >> 5) & 0x07
		byteCount := d.getByteCount(byteCountBits)
		d.pos += byteCount
	case 2: // string
		size, err := d.readCompressedUint()
		if err != nil {
			return err
		}
		d.pos += int(size)
	case 3: // object
		size, err := d.readCompressedUint()
		if err != nil {
			return err
		}
		for i := uint64(0); i < size; i++ {
			if err := d.skipValue(); err != nil { // key
				return err
			}
			if err := d.skipValue(); err != nil { // value
				return err
			}
		}
	case 4: // typed array
		size, err := d.readCompressedUint()
		if err != nil {
			return err
		}
		length, err := checkedLength(size)
		if err != nil {
			return err
		}
		group := (header >> 3) & 0x03
		switch group {
		case 3:
			isString := ((header >> 5) & 0x01) == 1
			if isString {
				return &UnsupportedError{"typed string arrays not implemented"}
			}
			payload := (length + 7) / 8
			if _, err := d.readBytes(payload); err != nil {
				return err
			}
		case 0, 1, 2:
			byteCount := d.getByteCount((header >> 5) & 0x07)
			total, err := totalByteCount(length, byteCount)
			if err != nil {
				return err
			}
			if _, err := d.readBytes(total); err != nil {
				return err
			}
		default:
			return &UnsupportedError{"unknown typed array"}
		}
	case 5: // generic arrays
		size, err := d.readCompressedUint()
		if err != nil {
			return err
		}
		for i := uint64(0); i < size; i++ {
			if err := d.skipValue(); err != nil {
				return err
			}
		}
	}
	return nil
}

// readByte reads a single byte
func (d *decoder) readByte() (byte, error) {
	if d.pos >= len(d.data) {
		return 0, &UnsupportedError{"unexpected end of data"}
	}
	b := d.data[d.pos]
	d.pos++
	return b, nil
}

// readBytes reads multiple bytes
func (d *decoder) readBytes(n int) ([]byte, error) {
	if d.pos+n > len(d.data) {
		return nil, &UnsupportedError{"unexpected end of data"}
	}
	result := d.data[d.pos : d.pos+n]
	d.pos += n
	return result, nil
}

// readCompressedUint reads a compressed unsigned integer
func (d *decoder) readCompressedUint() (uint64, error) {
	b, err := d.readByte()
	if err != nil {
		return 0, err
	}

	sizeIndicator := b & 0x03
	value := uint64(b >> 2)

	switch sizeIndicator {
	case 0:
		// value already set
	case 1:
		next, err := d.readByte()
		if err != nil {
			return 0, err
		}
		value = (value << 8) | uint64(next)
	case 2:
		for i := 0; i < 3; i++ {
			next, err := d.readByte()
			if err != nil {
				return 0, err
			}
			value = (value << 8) | uint64(next)
		}
	case 3:
		for i := 0; i < 7; i++ {
			next, err := d.readByte()
			if err != nil {
				return 0, err
			}
			value = (value << 8) | uint64(next)
		}
	}

	return value, nil
}

// getByteCount converts byte count bits to actual count
func (d *decoder) getByteCount(bits byte) int {
	switch bits {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 8
	case 4:
		return 16
	default:
		return 8 // fallback
	}
}

// setValue sets a value using reflection
func (d *decoder) setValue(v reflect.Value, val interface{}) error {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(val.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(val.(int64))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(val.(uint64))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(val.(float64))
	case reflect.String:
		v.SetString(val.(string))
	case reflect.Interface:
		v.Set(reflect.ValueOf(val))
	default:
		return &UnsupportedError{"cannot set value of type " + v.Type().String()}
	}
	return nil
}

// setBool sets a boolean value
func (d *decoder) setBool(v reflect.Value, b bool) error {
	if v.Kind() == reflect.Bool {
		v.SetBool(b)
		return nil
	}
	return &UnsupportedError{"expected bool"}
}

// setNil sets a nil value
func (d *decoder) setNil(v reflect.Value) error {
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	return &UnsupportedError{"cannot set nil"}
}

func (d *decoder) decodeBoolTypedArray(v reflect.Value, length int) error {
	payload := (length + 7) / 8
	data, err := d.readBytes(payload)
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := ensureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			if err := setBoolElement(v.Index(i), bit == 1); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() != length {
			return &UnsupportedError{"typed array length mismatch"}
		}
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			if err := setBoolElement(v.Index(i), bit == 1); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		values := make([]bool, length)
		for i := 0; i < length; i++ {
			bit := (data[i>>3] >> (uint(i) & 7)) & 0x01
			values[i] = bit == 1
		}
		v.Set(reflect.ValueOf(values))
		return nil
	default:
		return &UnsupportedError{"unsupported bool typed array target"}
	}
}

func (d *decoder) decodeSignedTypedArray(v reflect.Value, length, byteCount int) error {
	if byteCount != 1 && byteCount != 2 && byteCount != 4 && byteCount != 8 {
		return &UnsupportedError{"unsupported signed typed array size"}
	}
	total, err := totalByteCount(length, byteCount)
	if err != nil {
		return err
	}
	data, err := d.readBytes(total)
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := ensureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			value, err := readSignedValue(data[offset:offset+byteCount], byteCount)
			if err != nil {
				return err
			}
			if err := setSignedElement(v.Index(i), value, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() != length {
			return &UnsupportedError{"typed array length mismatch"}
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			value, err := readSignedValue(data[offset:offset+byteCount], byteCount)
			if err != nil {
				return err
			}
			if err := setSignedElement(v.Index(i), value, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		switch byteCount {
		case 1:
			values := make([]int8, length)
			for i := 0; i < length; i++ {
				values[i] = int8(data[i])
			}
			v.Set(reflect.ValueOf(values))
		case 2:
			values := make([]int16, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = int16(binary.LittleEndian.Uint16(data[offset : offset+byteCount]))
			}
			v.Set(reflect.ValueOf(values))
		case 4:
			values := make([]int32, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = int32(binary.LittleEndian.Uint32(data[offset : offset+byteCount]))
			}
			v.Set(reflect.ValueOf(values))
		case 8:
			values := make([]int64, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = int64(binary.LittleEndian.Uint64(data[offset : offset+byteCount]))
			}
			v.Set(reflect.ValueOf(values))
		default:
			return &UnsupportedError{"unsupported signed typed array size"}
		}
		return nil
	default:
		return &UnsupportedError{"unsupported signed typed array target"}
	}
}

func (d *decoder) decodeUnsignedTypedArray(v reflect.Value, length, byteCount int) error {
	if byteCount != 1 && byteCount != 2 && byteCount != 4 && byteCount != 8 {
		return &UnsupportedError{"unsupported unsigned typed array size"}
	}
	total, err := totalByteCount(length, byteCount)
	if err != nil {
		return err
	}
	data, err := d.readBytes(total)
	if err != nil {
		return err
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := ensureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			value, err := readUnsignedValue(data[offset:offset+byteCount], byteCount)
			if err != nil {
				return err
			}
			if err := setUnsignedElement(v.Index(i), value, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if v.Len() != length {
			return &UnsupportedError{"typed array length mismatch"}
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			value, err := readUnsignedValue(data[offset:offset+byteCount], byteCount)
			if err != nil {
				return err
			}
			if err := setUnsignedElement(v.Index(i), value, byteCount); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		switch byteCount {
		case 1:
			values := make([]uint8, length)
			copy(values, data[:length])
			v.Set(reflect.ValueOf(values))
		case 2:
			values := make([]uint16, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = binary.LittleEndian.Uint16(data[offset : offset+byteCount])
			}
			v.Set(reflect.ValueOf(values))
		case 4:
			values := make([]uint32, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = binary.LittleEndian.Uint32(data[offset : offset+byteCount])
			}
			v.Set(reflect.ValueOf(values))
		case 8:
			values := make([]uint64, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				values[i] = binary.LittleEndian.Uint64(data[offset : offset+byteCount])
			}
			v.Set(reflect.ValueOf(values))
		default:
			return &UnsupportedError{"unsupported unsigned typed array size"}
		}
		return nil
	default:
		return &UnsupportedError{"unsupported unsigned typed array target"}
	}
}

func (d *decoder) decodeFloatTypedArray(v reflect.Value, length, byteCount int) error {
	total, err := totalByteCount(length, byteCount)
	if err != nil {
		return err
	}
	data, err := d.readBytes(total)
	if err != nil {
		return err
	}

	switch byteCount {
	case 4, 8:
		// supported sizes
	default:
		return &UnsupportedError{"unsupported float typed array size"}
	}

	switch v.Kind() {
	case reflect.Slice:
		if err := ensureSliceLength(v, length); err != nil {
			return err
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			if byteCount == 4 {
				bits := binary.LittleEndian.Uint32(data[offset : offset+byteCount])
				value := math.Float32frombits(bits)
				if err := setFloatElement(v.Index(i), value, 0, byteCount); err != nil {
					return err
				}
			} else {
				bits := binary.LittleEndian.Uint64(data[offset : offset+byteCount])
				value := math.Float64frombits(bits)
				if err := setFloatElement(v.Index(i), 0, value, byteCount); err != nil {
					return err
				}
			}
		}
		return nil
	case reflect.Array:
		if v.Len() != length {
			return &UnsupportedError{"typed array length mismatch"}
		}
		for i := 0; i < length; i++ {
			offset := i * byteCount
			if byteCount == 4 {
				bits := binary.LittleEndian.Uint32(data[offset : offset+byteCount])
				value := math.Float32frombits(bits)
				if err := setFloatElement(v.Index(i), value, 0, byteCount); err != nil {
					return err
				}
			} else {
				bits := binary.LittleEndian.Uint64(data[offset : offset+byteCount])
				value := math.Float64frombits(bits)
				if err := setFloatElement(v.Index(i), 0, value, byteCount); err != nil {
					return err
				}
			}
		}
		return nil
	case reflect.Interface:
		if byteCount == 4 {
			values := make([]float32, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				bits := binary.LittleEndian.Uint32(data[offset : offset+byteCount])
				values[i] = math.Float32frombits(bits)
			}
			v.Set(reflect.ValueOf(values))
		} else {
			values := make([]float64, length)
			for i := 0; i < length; i++ {
				offset := i * byteCount
				bits := binary.LittleEndian.Uint64(data[offset : offset+byteCount])
				values[i] = math.Float64frombits(bits)
			}
			v.Set(reflect.ValueOf(values))
		}
		return nil
	default:
		return &UnsupportedError{"unsupported float typed array target"}
	}
}

func ensureSliceLength(v reflect.Value, length int) error {
	if v.Kind() != reflect.Slice {
		return &UnsupportedError{"target must be slice"}
	}
	if v.IsNil() {
		v.Set(reflect.MakeSlice(v.Type(), length, length))
		return nil
	}
	if v.Len() != length {
		if v.Cap() >= length {
			v.Set(v.Slice(0, length))
		} else {
			v.Set(reflect.MakeSlice(v.Type(), length, length))
		}
	}
	return nil
}

func setBoolElement(elem reflect.Value, value bool) error {
	switch elem.Kind() {
	case reflect.Bool:
		elem.SetBool(value)
		return nil
	case reflect.Interface:
		elem.Set(reflect.ValueOf(value))
		return nil
	default:
		return &UnsupportedError{"cannot assign bool to " + elem.Type().String()}
	}
}

func setSignedElement(elem reflect.Value, value int64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if elem.OverflowInt(value) {
			return &UnsupportedError{"signed integer overflow"}
		}
		elem.SetInt(value)
		return nil
	case reflect.Interface:
		switch byteCount {
		case 1:
			elem.Set(reflect.ValueOf(int8(value)))
		case 2:
			elem.Set(reflect.ValueOf(int16(value)))
		case 4:
			elem.Set(reflect.ValueOf(int32(value)))
		case 8:
			elem.Set(reflect.ValueOf(value))
		default:
			elem.Set(reflect.ValueOf(value))
		}
		return nil
	default:
		return &UnsupportedError{"cannot assign signed integer to " + elem.Type().String()}
	}
}

func setUnsignedElement(elem reflect.Value, value uint64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if elem.OverflowUint(value) {
			return &UnsupportedError{"unsigned integer overflow"}
		}
		elem.SetUint(value)
		return nil
	case reflect.Interface:
		switch byteCount {
		case 1:
			elem.Set(reflect.ValueOf(uint8(value)))
		case 2:
			elem.Set(reflect.ValueOf(uint16(value)))
		case 4:
			elem.Set(reflect.ValueOf(uint32(value)))
		case 8:
			elem.Set(reflect.ValueOf(value))
		default:
			elem.Set(reflect.ValueOf(value))
		}
		return nil
	default:
		return &UnsupportedError{"cannot assign unsigned integer to " + elem.Type().String()}
	}
}

func setFloatElement(elem reflect.Value, value32 float32, value64 float64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Float32:
		if byteCount == 4 {
			elem.SetFloat(float64(value32))
		} else {
			elem.SetFloat(value64)
		}
		return nil
	case reflect.Float64:
		if byteCount == 4 {
			elem.SetFloat(float64(value32))
		} else {
			elem.SetFloat(value64)
		}
		return nil
	case reflect.Interface:
		if byteCount == 4 {
			elem.Set(reflect.ValueOf(value32))
		} else {
			elem.Set(reflect.ValueOf(value64))
		}
		return nil
	default:
		return &UnsupportedError{"cannot assign float to " + elem.Type().String()}
	}
}

func readSignedValue(data []byte, byteCount int) (int64, error) {
	switch byteCount {
	case 1:
		return int64(int8(data[0])), nil
	case 2:
		return int64(int16(binary.LittleEndian.Uint16(data))), nil
	case 4:
		return int64(int32(binary.LittleEndian.Uint32(data))), nil
	case 8:
		return int64(binary.LittleEndian.Uint64(data)), nil
	default:
		return 0, &UnsupportedError{"unsupported signed typed array size"}
	}
}

func readUnsignedValue(data []byte, byteCount int) (uint64, error) {
	switch byteCount {
	case 1:
		return uint64(data[0]), nil
	case 2:
		return uint64(binary.LittleEndian.Uint16(data)), nil
	case 4:
		return uint64(binary.LittleEndian.Uint32(data)), nil
	case 8:
		return binary.LittleEndian.Uint64(data), nil
	default:
		return 0, &UnsupportedError{"unsupported unsigned typed array size"}
	}
}

func checkedLength(size uint64) (int, error) {
	maxInt := ^uint(0) >> 1
	if size > uint64(maxInt) {
		return 0, &UnsupportedError{"typed array too large"}
	}
	return int(size), nil
}

func totalByteCount(length, byteCount int) (int, error) {
	if length == 0 || byteCount == 0 {
		return 0, nil
	}
	total := length * byteCount
	if total/byteCount != length {
		return 0, &UnsupportedError{"typed array size overflow"}
	}
	return total, nil
}
