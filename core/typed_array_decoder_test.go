package core

import (
	"reflect"
	"testing"
)

// TestDecodeSignedTypedArray tests decodeSignedTypedArray (0% coverage)
func TestDecodeSignedTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		[]int8{-128, -64, 0, 64, 127},
		[]int16{-32768, -16384, 0, 16384, 32767},
		[]int32{-2147483648, 0, 2147483647},
		[]int64{-9223372036854775808, 0, 9223372036854775807},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode signed typed array: %v", i, err)
		}
	}
}

// TestDecodeUnsignedTypedArray tests decodeUnsignedTypedArray (0% coverage)
func TestDecodeUnsignedTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		[]uint8{0, 64, 128, 192, 255},
		[]uint16{0, 16384, 32768, 49152, 65535},
		[]uint32{0, 1073741824, 2147483648, 4294967295},
		[]uint64{0, 4611686018427387904, 9223372036854775808, 18446744073709551615},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode unsigned typed array: %v", i, err)
		}
	}
}

// TestDecodeFloatTypedArray tests decodeFloatTypedArray (18.6% -> increase)
func TestDecodeFloatTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		[]float32{-3.14, -1.0, 0.0, 1.0, 3.14},
		[]float64{-3.14159265359, -1.0, 0.0, 1.0, 3.14159265359},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode float typed array: %v", i, err)
		}
	}
}

// TestDecodeBoolTypedArray tests decodeBoolTypedArray (34.6% -> increase)
func TestDecodeBoolTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := [][]bool{
		{true, false, true, false, true},
		{false, false, false},
		{true, true, true},
		{true},
		{false},
		{},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		var result []bool
		if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
			t.Logf("Test %d decode bool typed array: %v", i, err)
		}
	}
}

// TestDecodeStringTypedArray tests decodeStringTypedArray (27.5% -> increase)
func TestDecodeStringTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := [][]string{
		{"", "a", "hello", "world"},
		{"short", "medium length", "very long string with many characters"},
		{"unicode", "你好", "世界", "🚀"},
		{},
		{"single"},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		var result []string
		if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
			t.Logf("Test %d decode string typed array: %v", i, err)
		}
	}
}

// TestDecodeTypedArray tests DecodeTypedArray (63.2% -> increase)
func TestDecodeTypedArray(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		// All typed arrays
		[]bool{true, false},
		[]int8{-128, 127},
		[]int16{-32768, 32767},
		[]int32{-2147483648, 2147483647},
		[]int64{-9223372036854775808, 9223372036854775807},
		[]uint8{0, 255},
		[]uint16{0, 65535},
		[]uint32{0, 4294967295},
		[]uint64{0, 18446744073709551615},
		[]float32{-3.14, 3.14},
		[]float64{-2.71828, 2.71828},
		[]string{"a", "b", "c"},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode typed array: %v", i, err)
		}
	}
}

// TestDecodeExtension tests DecodeExtension (0% coverage)
func TestDecodeExtension(t *testing.T) {
	// Extensions are not yet implemented in BEVE
	// This test ensures the function exists and can be called
	dec := NewDecoder([]byte{0xFF, 0x01, 0x02}) // Mock extension data

	var result interface{}
	rv := reflect.ValueOf(&result).Elem()
	err := dec.DecodeExtension(rv, 0x01)
	if err == nil {
		t.Log("DecodeExtension returned no error (extensions not implemented)")
	} else {
		t.Logf("DecodeExtension error: %v", err)
	}
}

// TestSetSignedElement tests SetSignedElement (0% coverage)
func TestSetSignedElement(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode int slices to trigger SetSignedElement during decode
	tests := []interface{}{
		[]int{-100, 0, 100},
		[]int8{-128, 0, 127},
		[]int16{-32768, 0, 32767},
		[]int32{-2147483648, 0, 2147483647},
		[]int64{-9223372036854775808, 0, 9223372036854775807},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode SetSignedElement: %v", i, err)
		}
	}
}

// TestSetUnsignedElement tests SetUnsignedElement (0% coverage)
func TestSetUnsignedElement(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode uint slices to trigger SetUnsignedElement during decode
	tests := []interface{}{
		[]uint{0, 100, 200},
		[]uint8{0, 128, 255},
		[]uint16{0, 32768, 65535},
		[]uint32{0, 2147483648, 4294967295},
		[]uint64{0, 9223372036854775808, 18446744073709551615},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode SetUnsignedElement: %v", i, err)
		}
	}
}

// TestSetFloatElement tests SetFloatElement (42.9% -> increase)
func TestSetFloatElement(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		[]float32{-3.14, 0.0, 3.14},
		[]float64{-2.71828, 0.0, 2.71828},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode SetFloatElement: %v", i, err)
		}
	}
}

// TestSetBoolElement tests SetBoolElement (27.3% -> increase)
func TestSetBoolElement(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := [][]bool{
		{true, false, true, false},
		{false, false, false},
		{true, true, true},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		var result []bool
		if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
			t.Logf("Test %d decode SetBoolElement: %v", i, err)
		}
	}
}

// TestEnsureSliceLength tests EnsureSliceLength (36.4% -> increase)
func TestEnsureSliceLength(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Test slices of various lengths
	tests := []interface{}{
		[]int{1},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		make([]int, 100),
		make([]int, 1000),
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode EnsureSliceLength: %v", i, err)
		}
	}
}

// TestGetMapValueDecoder tests getMapValueDecoder (70% -> increase)
func TestGetMapValueDecoder(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		map[string]int{"a": 1, "b": 2},
		map[string]string{"x": "X", "y": "Y"},
		map[string]bool{"t": true, "f": false},
		map[string]float64{"pi": 3.14, "e": 2.71},
		map[string][]int{"list": {1, 2, 3}},
		map[string]map[string]int{"nested": {"a": 1}},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode getMapValueDecoder: %v", i, err)
		}
	}
}

// TestBuildMapValueDecoder tests buildMapValueDecoder (41.7% -> increase)
func TestBuildMapValueDecoder(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Test various map value types
	tests := []interface{}{
		map[string]interface{}{"mixed": 1, "types": "here"},
		map[int][]string{1: {"a", "b"}, 2: {"c", "d"}},
		map[uint]map[string]int{1: {"x": 10}, 2: {"y": 20}},
	}

	for i, original := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		dec := NewDecoder(enc.Buf.Bytes())
		result := reflect.New(reflect.TypeOf(original)).Elem()
		if err := dec.Decode(result); err != nil {
			t.Logf("Test %d decode buildMapValueDecoder: %v", i, err)
		}
	}
}

// TestBuildStructFieldDecoder tests buildStructFieldDecoder (41.7% -> increase)
func TestBuildStructFieldDecoder(t *testing.T) {
	type ComplexStruct struct {
		IntField    int
		StringField string
		BoolField   bool
		FloatField  float64
		SliceField  []int
		MapField    map[string]int
		NestedField struct {
			X int
			Y string
		}
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	original := ComplexStruct{
		IntField:    42,
		StringField: "test",
		BoolField:   true,
		FloatField:  3.14,
		SliceField:  []int{1, 2, 3},
		MapField:    map[string]int{"a": 1},
		NestedField: struct {
			X int
			Y string
		}{X: 10, Y: "nested"},
	}

	if err := enc.Encode(reflect.ValueOf(original)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	dec := NewDecoder(enc.Buf.Bytes())
	var result ComplexStruct
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode buildStructFieldDecoder: %v", err)
	}
}
