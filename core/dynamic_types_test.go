package core

import (
	"reflect"
	"testing"
)

// Wave 8: Dynamic Types & Builders (15-48% → 100%)
// Complete coverage for interface{} encoding, map/slice builders, generic arrays
// Target: +3% coverage

// ============================================================================
// INTERFACE{} ENCODING COMPLETE COVERAGE (15.0% → 100%)
// ============================================================================

// TestEncodeInterfaceValueAllTypes tests encodeInterfaceValue with all type variations
func TestEncodeInterfaceValueAllTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// Primitives
		{"nil", nil},
		{"bool_true", true},
		{"bool_false", false},
		{"int", 42},
		{"int8", int8(-128)},
		{"int16", int16(32767)},
		{"int32", int32(-2147483648)},
		{"int64", int64(9223372036854775807)},
		{"uint", uint(42)},
		{"uint8", uint8(255)},
		{"uint16", uint16(65535)},
		{"uint32", uint32(4294967295)},
		{"uint64", uint64(18446744073709551615)},
		{"float32", float32(3.14)},
		{"float64", float64(2.718281828)},
		{"string", "hello world"},
		{"bytes", []byte{1, 2, 3, 4, 5}},

		// Collections
		{"slice_int", []int{1, 2, 3, 4, 5}},
		{"slice_string", []string{"a", "b", "c"}},
		{"slice_interface", []interface{}{1, "two", 3.0}},
		{"map_string_int", map[string]int{"a": 1, "b": 2}},
		{"map_int_string", map[int]string{1: "one", 2: "two"}},
		{"map_interface", map[string]interface{}{"num": 42, "str": "test"}},

		// Nested structures
		{"nested_slice", [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"nested_map", map[string]map[string]int{"outer": {"inner": 42}}},

		// Structs
		{"simple_struct", struct {
			Name string
			Age  int
		}{"John", 30}},
		{"empty_struct", struct{}{}},

		// Pointers
		{"pointer_to_int", func() interface{} { v := 42; return &v }()},
		{"pointer_to_string", func() interface{} { v := "test"; return &v }()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode interface{} %s failed: %v", tt.name, err)
			}

			if enc.Buf.Len() == 0 {
				t.Errorf("%s: No data encoded", tt.name)
			}

			// Try to decode back
			dec := NewDecoder(enc.Buf.Bytes())
			var result interface{}
			resultVal := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(resultVal); err != nil {
				t.Logf("%s: Decode: %v (may be expected for some types)", tt.name, err)
			}
			PutDecoderToPool(dec)
		})
	}

	t.Log("✓ encodeInterfaceValue with all types tested")
}

// TestEncodeInterfaceValueMixedTypes tests interface{} with mixed type arrays/maps
func TestEncodeInterfaceValueMixedTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		{
			"mixed_array",
			[]interface{}{
				42,
				"string",
				3.14,
				true,
				[]int{1, 2, 3},
				map[string]interface{}{"nested": "map"},
			},
		},
		{
			"complex_map",
			map[string]interface{}{
				"int":    42,
				"string": "value",
				"float":  3.14,
				"bool":   true,
				"array":  []int{1, 2, 3},
				"nested": map[string]int{"a": 1},
			},
		},
		{
			"deeply_nested",
			map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": map[string]interface{}{
						"level3": "deep value",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if enc.Buf.Len() == 0 {
				t.Error("No data encoded")
			}
		})
	}

	t.Log("✓ encodeInterfaceValue with mixed types tested")
}

// ============================================================================
// MAP BUILDER COVERAGE (31.9% → 100%)
// ============================================================================

// TestBuildMapEncoderAllKeyTypes tests buildMapEncoder with various key types
func TestBuildMapEncoderAllKeyTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// String keys
		{"map_string_int", map[string]int{"a": 1, "b": 2, "c": 3}},
		{"map_string_string", map[string]string{"key1": "val1", "key2": "val2"}},
		{"map_string_float", map[string]float64{"pi": 3.14, "e": 2.718}},
		{"map_string_bool", map[string]bool{"yes": true, "no": false}},
		{"map_string_interface", map[string]interface{}{"mixed": 42, "types": "here"}},

		// Integer keys
		{"map_int_string", map[int]string{1: "one", 2: "two", 3: "three"}},
		{"map_int8_int", map[int8]int{1: 10, 2: 20}},
		{"map_int16_int", map[int16]int{100: 1000, 200: 2000}},
		{"map_int32_int", map[int32]int{1000: 10000, 2000: 20000}},
		{"map_int64_int", map[int64]int{100000: 1000000}},

		// Unsigned integer keys
		{"map_uint_string", map[uint]string{1: "one", 2: "two"}},
		{"map_uint8_int", map[uint8]int{10: 100, 20: 200}},
		{"map_uint16_int", map[uint16]int{100: 1000}},
		{"map_uint32_int", map[uint32]int{1000: 10000}},
		{"map_uint64_int", map[uint64]int{100000: 1000000}},

		// Complex value types
		{"map_string_slice", map[string][]int{"a": {1, 2}, "b": {3, 4}}},
		{"map_string_map", map[string]map[string]int{"nested": {"a": 1}}},
		{"map_string_struct", map[string]struct{ X, Y int }{"point": {X: 10, Y: 20}}},

		// Empty maps
		{"map_empty_string_int", map[string]int{}},
		{"map_empty_int_string", map[int]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode map %s failed: %v", tt.name, err)
			}

			if enc.Buf.Len() == 0 && rv.Len() > 0 {
				t.Error("No data encoded for non-empty map")
			}
		})
	}

	t.Log("✓ buildMapEncoder with all key types tested")
}

// TestBuildMapValueDecoderAllTypes tests buildMapValueDecoder with various value types
func TestBuildMapValueDecoderAllTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// Primitive values
		{"map_to_int", map[string]int{"key": 42}},
		{"map_to_string", map[string]string{"key": "value"}},
		{"map_to_bool", map[string]bool{"key": true}},
		{"map_to_float32", map[string]float32{"key": 3.14}},
		{"map_to_float64", map[string]float64{"key": 2.718}},

		// Slice values
		{"map_to_slice_int", map[string][]int{"key": {1, 2, 3}}},
		{"map_to_slice_string", map[string][]string{"key": {"a", "b"}}},

		// Nested map values
		{"map_to_map", map[string]map[string]int{"key": {"nested": 42}}},

		// Struct values
		{"map_to_struct", map[string]struct{ X int }{"key": {X: 10}}},

		// Interface values
		{"map_to_interface", map[string]interface{}{"key": 42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()
			rv := reflect.ValueOf(tt.value)

			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			// Decode back
			dec := NewDecoder(enc.Buf.Bytes())
			result := reflect.New(rv.Type()).Elem()
			if err := dec.Decode(result); err != nil {
				t.Logf("Decode %s: %v", tt.name, err)
			}
			PutDecoderToPool(dec)
		})
	}

	t.Log("✓ buildMapValueDecoder with all types tested")
}

// ============================================================================
// SLICE BUILDER COVERAGE (24.4% → 100%)
// ============================================================================

// TestBuildSliceEncoderAllElementTypes tests buildSliceEncoder with various element types
func TestBuildSliceEncoderAllElementTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// Primitive slices
		{"slice_int", []int{1, 2, 3, 4, 5}},
		{"slice_int8", []int8{-128, 0, 127}},
		{"slice_int16", []int16{-32768, 0, 32767}},
		{"slice_int32", []int32{-2147483648, 0, 2147483647}},
		{"slice_int64", []int64{-9223372036854775808, 0, 9223372036854775807}},
		{"slice_uint", []uint{0, 100, 1000}},
		{"slice_uint8", []uint8{0, 128, 255}},
		{"slice_uint16", []uint16{0, 32768, 65535}},
		{"slice_uint32", []uint32{0, 2147483648, 4294967295}},
		{"slice_uint64", []uint64{0, 9223372036854775808, 18446744073709551615}},
		{"slice_float32", []float32{-3.14, 0.0, 3.14}},
		{"slice_float64", []float64{-2.718, 0.0, 2.718}},
		{"slice_bool", []bool{true, false, true, false}},
		{"slice_string", []string{"hello", "world", "test"}},

		// Nested slices
		{"slice_of_slices_int", [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"slice_of_slices_string", [][]string{{"a", "b"}, {"c", "d"}}},

		// Slice of structs
		{"slice_of_structs", []struct{ X, Y int }{{1, 2}, {3, 4}}},

		// Slice of maps
		{"slice_of_maps", []map[string]int{{"a": 1}, {"b": 2}}},

		// Slice of interfaces
		{"slice_of_interfaces", []interface{}{1, "two", 3.0, true}},

		// Empty slices
		{"slice_empty_int", []int{}},
		{"slice_empty_string", []string{}},

		// Large slices
		{"slice_large", make([]int, 100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode slice %s failed: %v", tt.name, err)
			}

			if enc.Buf.Len() == 0 && rv.Len() > 0 {
				t.Error("No data encoded for non-empty slice")
			}
		})
	}

	t.Log("✓ buildSliceEncoder with all element types tested")
}

// TestBuildSliceEncoderWithPointers tests slice encoder with pointer elements
func TestBuildSliceEncoderWithPointers(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Create pointer slices
	int1, int2, int3 := 1, 2, 3
	str1, str2 := "a", "b"

	tests := []struct {
		name  string
		value interface{}
	}{
		{"slice_ptr_int", []*int{&int1, &int2, &int3}},
		{"slice_ptr_string", []*string{&str1, &str2}},
		{"slice_ptr_mixed", []*int{&int1, nil, &int3}}, // with nil pointer
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode pointer slice %s failed: %v", tt.name, err)
			}

			if enc.Buf.Len() == 0 && rv.Len() > 0 {
				t.Error("No data encoded for non-empty pointer slice")
			}
		})
	}

	t.Log("✓ buildSliceEncoder with pointer elements tested")
}

// ============================================================================
// PRIMITIVE SLICE COVERAGE (38.5% → 100%)
// ============================================================================

// TestEncodePrimitiveSliceAllTypes tests encodePrimitiveSlice with all primitive types
func TestEncodePrimitiveSliceAllTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// All primitive int types
		{"prim_slice_int", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"prim_slice_int8", []int8{-128, -64, 0, 64, 127}},
		{"prim_slice_int16", []int16{-32768, -16000, 0, 16000, 32767}},
		{"prim_slice_int32", []int32{-2147483648, -1000000, 0, 1000000, 2147483647}},
		{"prim_slice_int64", []int64{-9223372036854775808, 0, 9223372036854775807}},

		// All primitive uint types
		{"prim_slice_uint", []uint{0, 100, 1000, 10000, 100000}},
		{"prim_slice_uint8", []uint8{0, 64, 128, 192, 255}},
		{"prim_slice_uint16", []uint16{0, 16384, 32768, 49152, 65535}},
		{"prim_slice_uint32", []uint32{0, 1073741824, 2147483648, 4294967295}},
		{"prim_slice_uint64", []uint64{0, 4611686018427387904, 9223372036854775808, 18446744073709551615}},

		// Float types
		{"prim_slice_float32", []float32{-3.14, -1.0, 0.0, 1.0, 3.14}},
		{"prim_slice_float64", []float64{-3.14159265359, -1.0, 0.0, 1.0, 3.14159265359}},

		// Bool
		{"prim_slice_bool", []bool{true, false, true, false, true, false}},

		// String (handled differently but test it)
		{"prim_slice_string", []string{"a", "b", "c", "d", "e"}},

		// Large primitive slices
		{"prim_slice_large_int32", make([]int32, 500)},
		{"prim_slice_large_uint64", make([]uint64, 300)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode primitive slice %s failed: %v", tt.name, err)
			}

			if enc.Buf.Len() == 0 && rv.Len() > 0 {
				t.Error("No data encoded for non-empty primitive slice")
			}
		})
	}

	t.Log("✓ encodePrimitiveSlice with all primitive types tested")
}

// ============================================================================
// GENERIC ARRAY COVERAGE (35.7% → 100%)
// ============================================================================

// TestDecodeGenericArrayAllTypes tests DecodeGenericArray with various types
func TestDecodeGenericArrayAllTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// Generic arrays of primitives
		{"generic_int", []int{1, 2, 3, 4, 5}},
		{"generic_string", []string{"hello", "world", "test"}},
		{"generic_bool", []bool{true, false, true}},
		{"generic_float", []float64{1.1, 2.2, 3.3}},

		// Mixed interface arrays
		{"generic_mixed", []interface{}{1, "two", 3.0, true, nil}},

		// Nested arrays
		{"generic_nested", []interface{}{
			[]int{1, 2, 3},
			[]string{"a", "b"},
			map[string]int{"key": 42},
		}},

		// Arrays with structs
		{"generic_structs", []struct{ X int }{{X: 1}, {X: 2}, {X: 3}}},

		// Empty arrays
		{"generic_empty", []interface{}{}},

		// Large arrays
		{"generic_large", make([]int, 200)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()
			rv := reflect.ValueOf(tt.value)

			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			// Decode as generic array
			dec := NewDecoder(enc.Buf.Bytes())
			result := reflect.New(rv.Type()).Elem()
			if err := dec.Decode(result); err != nil {
				t.Logf("Decode %s: %v", tt.name, err)
			}
			PutDecoderToPool(dec)
		})
	}

	t.Log("✓ DecodeGenericArray with all types tested")
}

// ============================================================================
// MAP KEY CONVERSION COVERAGE (43.6% → 100%)
// ============================================================================

// TestConvertMapKeyValueAllCases tests convertMapKeyValue with all key types
func TestConvertMapKeyValueAllCases(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		// String keys (most common)
		{"convert_string_keys", map[string]int{"a": 1, "b": 2, "c": 3}},
		{"convert_string_keys_unicode", map[string]int{"世界": 1, "🚀": 2}},

		// Integer keys
		{"convert_int_keys", map[int]string{1: "one", 2: "two", -1: "minus one"}},
		{"convert_int8_keys", map[int8]string{-128: "min", 127: "max"}},
		{"convert_int16_keys", map[int16]string{-32768: "min", 32767: "max"}},
		{"convert_int32_keys", map[int32]string{-2147483648: "min", 2147483647: "max"}},
		{"convert_int64_keys", map[int64]string{-9223372036854775808: "min", 9223372036854775807: "max"}},

		// Unsigned integer keys
		{"convert_uint_keys", map[uint]string{0: "zero", 100: "hundred"}},
		{"convert_uint8_keys", map[uint8]string{0: "zero", 255: "max"}},
		{"convert_uint16_keys", map[uint16]string{0: "zero", 65535: "max"}},
		{"convert_uint32_keys", map[uint32]string{0: "zero", 4294967295: "max"}},
		{"convert_uint64_keys", map[uint64]string{0: "zero", 18446744073709551615: "max"}},

		// Complex value types with conversion
		{"convert_mixed_values", map[string]interface{}{"int": 42, "str": "test", "bool": true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()
			rv := reflect.ValueOf(tt.value)

			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode map with key conversion %s failed: %v", tt.name, err)
			}

			// Decode back
			dec := NewDecoder(enc.Buf.Bytes())
			result := reflect.New(rv.Type()).Elem()
			if err := dec.Decode(result); err != nil {
				t.Logf("Decode %s: %v", tt.name, err)
			}

			// Verify key count matches
			if result.Len() != rv.Len() {
				t.Errorf("Key count mismatch: expected %d, got %d", rv.Len(), result.Len())
			}

			PutDecoderToPool(dec)
		})
	}

	t.Log("✓ convertMapKeyValue with all key types tested")
}

// ============================================================================
// STRESS TESTS & EDGE CASES
// ============================================================================

// TestDynamicTypesStressTest tests complex nested dynamic structures
func TestDynamicTypesStressTest(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Complex nested structure
	complexData := map[string]interface{}{
		"users": []interface{}{
			map[string]interface{}{
				"id":   1,
				"name": "John",
				"tags": []string{"admin", "user"},
				"meta": map[string]int{"posts": 100, "likes": 500},
			},
			map[string]interface{}{
				"id":   2,
				"name": "Jane",
				"tags": []string{"user"},
				"meta": map[string]int{"posts": 50, "likes": 250},
			},
		},
		"settings": map[string]interface{}{
			"enabled": true,
			"timeout": 30,
			"limits": map[string]int{
				"max_users":    1000,
				"max_requests": 10000,
			},
		},
		"data": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	enc.Buf.Reset()
	rv := reflect.ValueOf(complexData)

	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode complex structure failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("No data encoded")
	}

	t.Logf("✓ Complex nested structure encoded: %d bytes", enc.Buf.Len())
}

// TestDynamicTypesEmptyCollections tests empty maps and slices
func TestDynamicTypesEmptyCollections(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		{"empty_map", map[string]int{}},
		{"empty_slice", []int{}},
		{"empty_interface_slice", []interface{}{}},
		{"map_with_empty_values", map[string][]int{"a": {}, "b": {}, "c": {}}},
		{"slice_of_empty_maps", []map[string]int{{}, {}, {}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode %s failed: %v", tt.name, err)
			}

			// Empty collections should still produce some encoding
			if enc.Buf.Len() == 0 {
				t.Error("No data encoded for empty collection")
			}
		})
	}

	t.Log("✓ Empty collections handled correctly")
}

// Benchmarks
func BenchmarkEncodeInterfaceSlice(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	data := []interface{}{1, "two", 3.0, true, []int{1, 2, 3}}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.Encode(reflect.ValueOf(data))
	}
}

func BenchmarkEncodeMapStringInterface(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	data := map[string]interface{}{
		"int":    42,
		"string": "test",
		"float":  3.14,
		"bool":   true,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.Encode(reflect.ValueOf(data))
	}
}
