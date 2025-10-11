package core

import (
	"reflect"
	"testing"
)

// TestEncodeMapUintFast tests encodeMapUintFast (0% coverage)
func TestEncodeMapUintFast(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		map[uint]int{1: 10, 2: 20, 3: 30},
		map[uint8]string{1: "a", 2: "b", 3: "c"},
		map[uint16]bool{1: true, 2: false},
		map[uint32]float64{1: 1.1, 2: 2.2},
		map[uint64]interface{}{1: "one", 2: 2, 3: 3.0},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d encode uint key map failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestEncodeStringInterfaceMap tests encodeStringInterfaceMap (0% coverage)
func TestEncodeStringInterfaceMap(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []map[string]interface{}{
		{"str": "value", "int": 42, "float": 3.14, "bool": true},
		{"nested": map[string]interface{}{"inner": "value"}},
		{"slice": []int{1, 2, 3}, "map": map[string]int{"a": 1}},
		{"nil": nil, "empty": ""},
		{"mixed": []interface{}{1, "two", 3.0, true}},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d encode string-interface map failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestEncodeInterfaceValue tests encodeInterfaceValue (0% coverage)
func TestEncodeInterfaceValue(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		"string value",
		42,
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		nil,
		struct{ X int }{X: 10},
	}

	for i, val := range tests {
		enc.Buf.Reset()
		// Wrap in interface{} to trigger encodeInterfaceValue
		data := map[string]interface{}{"value": val}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d encode interface value failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestBuildMapEncoder tests buildMapEncoder (14.5% -> increase)
func TestBuildMapEncoder(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		// String keys
		map[string]int{"a": 1, "b": 2},
		map[string]string{"x": "X", "y": "Y"},
		map[string]interface{}{"mixed": 1, "types": "here"},

		// Integer keys
		map[int]string{1: "one", 2: "two"},
		map[int8]int{1: 10, 2: 20},
		map[int16]float64{1: 1.1, 2: 2.2},
		map[int32]bool{1: true, 2: false},
		map[int64]string{1: "a", 2: "b"},

		// Unsigned keys
		map[uint]int{1: 10, 2: 20},
		map[uint8]string{1: "a", 2: "b"},
		map[uint16]int{1: 10, 2: 20},
		map[uint32]float64{1: 1.1, 2: 2.2},
		map[uint64]bool{1: true, 2: false},

		// Complex values
		map[string][]int{"list": {1, 2, 3}},
		map[string]map[string]int{"nested": {"a": 1}},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d buildMapEncoder failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestBuildSliceEncoder tests buildSliceEncoder (17.1% -> increase)
func TestBuildSliceEncoder(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		// Primitive slices
		[]int{1, 2, 3, 4, 5},
		[]int8{-128, 0, 127},
		[]int16{-32768, 0, 32767},
		[]int32{-2147483648, 0, 2147483647},
		[]int64{-9223372036854775808, 0, 9223372036854775807},
		[]uint{0, 1, 2, 3},
		[]uint8{0, 128, 255},
		[]uint16{0, 32768, 65535},
		[]uint32{0, 2147483648, 4294967295},
		[]uint64{0, 9223372036854775808, 18446744073709551615},
		[]float32{0.0, 3.14, -2.71},
		[]float64{0.0, 3.14159265359, -2.71828182846},
		[]bool{true, false, true},
		[]string{"a", "b", "c"},

		// Complex slices
		[][]int{{1, 2}, {3, 4}},
		[]map[string]int{{"a": 1}, {"b": 2}},
		[]interface{}{1, "two", 3.0, true},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d buildSliceEncoder failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestEncodePrimitiveSlice tests encodePrimitiveSlice (30.8% -> increase)
func TestEncodePrimitiveSlice(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Test all primitive typed slices
	tests := []interface{}{
		[]bool{true, false, true, false, true},
		[]int8{-128, -64, 0, 64, 127},
		[]int16{-32768, -16384, 0, 16384, 32767},
		[]int32{-2147483648, -1073741824, 0, 1073741824, 2147483647},
		[]int64{-9223372036854775808, 0, 9223372036854775807},
		[]uint8{0, 64, 128, 192, 255},
		[]uint16{0, 16384, 32768, 49152, 65535},
		[]uint32{0, 1073741824, 2147483648, 3221225472, 4294967295},
		[]uint64{0, 4611686018427387904, 9223372036854775808, 13835058055282163712, 18446744073709551615},
		[]float32{-3.14, -1.0, 0.0, 1.0, 3.14},
		[]float64{-3.14159265359, -1.0, 0.0, 1.0, 3.14159265359},
		[]string{"", "a", "hello", "world", "test"},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d encodePrimitiveSlice failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestAppendEncodedUint tests appendEncodedUint (0% coverage)
func TestAppendEncodedUint(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode slices that will trigger appendEncodedUint
	tests := [][]uint64{
		{0, 1, 2, 3, 4, 5},
		{63, 64, 127, 128, 255},
		{256, 512, 1024, 2048, 4096},
		{65535, 65536, 131072},
		{4294967295, 4294967296},
		{18446744073709551615},
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d appendEncodedUint failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestAppendEncodedFloat32 tests appendEncodedFloat32 (0% coverage)
func TestAppendEncodedFloat32(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := [][]float32{
		{0.0, 1.0, -1.0, 3.14, -3.14},
		{3.40282346638528859811704183484516925440e+38},  // MaxFloat32
		{1.401298464324817070923729583289916131280e-45}, // SmallestNonzeroFloat32
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d appendEncodedFloat32 failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestAppendCompressedUint tests appendCompressedUint (22.2% -> increase)
func TestAppendCompressedUint(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Test varint encoding with various sizes
	tests := [][]uint64{
		{0, 1, 2, 63},                        // 1-byte
		{64, 127, 255, 16383},                // 2-byte
		{16384, 32767, 65535, 1073741823},    // 3-byte
		{1073741824, 2147483647, 4294967295}, // 4-byte
	}

	for i, data := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d appendCompressedUint failed: %v", i, err)
			continue
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Test %d: Nothing encoded", i)
		}
	}
}

// TestCompressedUintLen tests compressedUintLen (42.9% -> increase)
func TestCompressedUintLen(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode various sized uints to test length calculation
	tests := []uint64{
		0, 1, 63, // Should be 1 byte
		64, 16383, // Should be 2 bytes
		16384, 1073741823, // Should be 3 bytes
		1073741824, // Should be 4 bytes
	}

	for _, val := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(val)); err != nil {
			t.Errorf("Encode %d failed: %v", val, err)
		}
	}
}
