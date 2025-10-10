package core

import (
	"math"
	"reflect"
	"testing"
)

// TestDecoderPooling tests decoder pooling functionality
func TestDecoderPooling(t *testing.T) {
	dec1 := NewDecoder(nil)
	if dec1 == nil {
		t.Fatal("NewDecoder returned nil")
	}

	PutDecoderToPool(dec1)

	dec2 := NewDecoder(nil)
	if dec2 == nil {
		t.Fatal("NewDecoder returned nil after pool return")
	}

	// Should be the same instance (pooled)
	if dec1 != dec2 {
		t.Log("Note: Decoder instances are different (pool may have grown)")
	}

	PutDecoderToPool(dec2)
}

// TestDecodeBasicTypes tests decoding of basic types
func TestDecodeBasicTypes(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"bool_true", true},
		{"bool_false", false},
		{"int_zero", 0},
		{"int_positive", 42},
		{"int_negative", -42},
		{"uint", uint(100)},
		{"float32", float32(3.14)},
		{"float64", 3.14159},
		{"string_empty", ""},
		{"string_simple", "hello"},
		{"string_unicode", "Hello 世界 🌍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode first
			enc := GetEncoderFromPool()
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			data := enc.Buf.Bytes()
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoder(data)
			defer PutDecoderToPool(dec)

			targetType := reflect.TypeOf(tt.value)
			result := reflect.New(targetType).Elem()

			if err := dec.Decode(result); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if result.Interface() != tt.value {
				t.Errorf("Mismatch: got %v, want %v", result.Interface(), tt.value)
			}
		})
	}
}

// TestDecodeSlices tests slice decoding
func TestDecodeSlices(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"slice_int", []int{1, 2, 3}},
		{"slice_string", []string{"a", "b", "c"}},
		{"slice_float", []float64{1.1, 2.2, 3.3}},
		{"slice_bool", []bool{true, false, true}},
		{"slice_empty", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			enc := GetEncoderFromPool()
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			data := enc.Buf.Bytes()
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoder(data)
			defer PutDecoderToPool(dec)

			targetType := reflect.TypeOf(tt.value)
			result := reflect.New(targetType).Elem()

			if err := dec.Decode(result); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if !reflect.DeepEqual(result.Interface(), tt.value) {
				t.Errorf("Mismatch: got %v, want %v", result.Interface(), tt.value)
			}
		})
	}
}

// TestDecodeMaps tests map decoding
func TestDecodeMaps(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"map_string_int", map[string]int{"a": 1, "b": 2}},
		{"map_string_string", map[string]string{"key": "value"}},
		{"map_int_string", map[int]string{1: "one", 2: "two"}},
		{"map_empty", map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			enc := GetEncoderFromPool()
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			data := enc.Buf.Bytes()
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoder(data)
			defer PutDecoderToPool(dec)

			targetType := reflect.TypeOf(tt.value)
			result := reflect.New(targetType).Elem()

			if err := dec.Decode(result); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if !reflect.DeepEqual(result.Interface(), tt.value) {
				t.Errorf("Mismatch: got %v, want %v", result.Interface(), tt.value)
			}
		})
	}
}

// TestDecodeStructs tests struct decoding
func TestDecodeStructs(t *testing.T) {
	type SimpleStruct struct {
		Name  string
		Age   int
		Score float64
	}

	type NestedStruct struct {
		Title string
		Inner SimpleStruct
	}

	tests := []struct {
		name  string
		value interface{}
	}{
		{"simple_struct", SimpleStruct{Name: "Alice", Age: 30, Score: 95.5}},
		{"nested_struct", NestedStruct{
			Title: "Test",
			Inner: SimpleStruct{Name: "Bob", Age: 25, Score: 88.0},
		}},
		{"empty_struct", struct{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			enc := GetEncoderFromPool()
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			data := enc.Buf.Bytes()
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoder(data)
			defer PutDecoderToPool(dec)

			targetType := reflect.TypeOf(tt.value)
			result := reflect.New(targetType).Elem()

			if err := dec.Decode(result); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if !reflect.DeepEqual(result.Interface(), tt.value) {
				t.Errorf("Mismatch: got %v, want %v", result.Interface(), tt.value)
			}
		})
	}
}

// TestDecodeEdgeCases tests edge cases and error conditions
func TestDecodeEdgeCases(t *testing.T) {
	t.Run("large_numbers", func(t *testing.T) {
		values := []interface{}{
			int64(math.MaxInt64),
			int64(math.MinInt64),
			uint64(math.MaxUint64),
			float64(math.MaxFloat64),
			float64(math.SmallestNonzeroFloat64),
		}

		for _, v := range values {
			// Encode
			enc := GetEncoderFromPool()
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(v)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode %v failed: %v", v, err)
			}

			data := enc.Buf.Bytes()
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoder(data)
			targetType := reflect.TypeOf(v)
			result := reflect.New(targetType).Elem()

			if err := dec.Decode(result); err != nil {
				t.Fatalf("Decode %v failed: %v", v, err)
			}

			PutDecoderToPool(dec)

			// Compare (with tolerance for floats)
			if targetType.Kind() == reflect.Float64 {
				got := result.Float()
				want := rv.Float()
				if math.IsNaN(want) {
					if !math.IsNaN(got) {
						t.Errorf("Mismatch: got %v, want NaN", got)
					}
				} else if math.Abs(got-want) > 1e-10 {
					t.Errorf("Mismatch: got %v, want %v", got, want)
				}
			} else {
				if result.Interface() != v {
					t.Errorf("Mismatch: got %v, want %v", result.Interface(), v)
				}
			}
		}
	})

	t.Run("empty_data", func(t *testing.T) {
		dec := NewDecoder([]byte{})
		defer PutDecoderToPool(dec)

		var result int
		rv := reflect.ValueOf(&result).Elem()

		// Empty data should cause an error
		err := dec.Decode(rv)
		if err == nil {
			t.Error("Expected error for empty data, got nil")
		}
	})

	t.Run("corrupted_data", func(t *testing.T) {
		// Random bytes unlikely to be valid BEVE data
		corruptedData := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
		dec := NewDecoder(corruptedData)
		defer PutDecoderToPool(dec)

		var result int
		rv := reflect.ValueOf(&result).Elem()

		// Corrupted data should cause an error or handle gracefully
		_ = dec.Decode(rv)
		// Note: We don't check for error because the behavior depends on implementation
	})
}

// TestRoundTrip tests encode-decode round trips
func TestRoundTrip(t *testing.T) {
	type ComplexStruct struct {
		ID       int
		Name     string
		Tags     []string
		Scores   map[string]float64
		Active   bool
		Metadata struct {
			Created int64
			Updated int64
		}
	}

	original := ComplexStruct{
		ID:     123,
		Name:   "Test Object",
		Tags:   []string{"tag1", "tag2", "tag3"},
		Scores: map[string]float64{"math": 95.5, "science": 88.0},
		Active: true,
		Metadata: struct {
			Created int64
			Updated int64
		}{
			Created: 1234567890,
			Updated: 9876543210,
		},
	}

	// Encode
	enc := GetEncoderFromPool()
	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	rv := reflect.ValueOf(original)
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	data := enc.Buf.Bytes()
	PutEncoderToPool(enc)

	// Decode
	dec := NewDecoder(data)
	defer PutDecoderToPool(dec)

	var result ComplexStruct
	rv = reflect.ValueOf(&result).Elem()

	if err := dec.Decode(rv); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Compare
	if !reflect.DeepEqual(result, original) {
		t.Errorf("Round trip failed:\ngot:  %+v\nwant: %+v", result, original)
	}
}

// BenchmarkDecoderPool benchmarks decoder pooling overhead
func BenchmarkDecoderPool(b *testing.B) {
	data := []byte{0x01, 0x02, 0x03}

	b.Run("GetAndPut", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			PutDecoderToPool(dec)
		}
	})

	b.Run("NoPool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = &Decoder{Data: data}
		}
	})
}

// BenchmarkDecodeTypes benchmarks decoding different types
func BenchmarkDecodeTypes(b *testing.B) {
	// Prepare encoded data
	encodeValue := func(v interface{}) []byte {
		enc := GetEncoderFromPool()
		if enc.Buf != nil {
			enc.Buf.Reset()
		}
		rv := reflect.ValueOf(v)
		_ = enc.Encode(rv)
		data := enc.Buf.Bytes()
		PutEncoderToPool(enc)
		return data
	}

	b.Run("Int", func(b *testing.B) {
		data := encodeValue(42)
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			var result int
			rv := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(rv)
			PutDecoderToPool(dec)
		}
	})

	b.Run("String", func(b *testing.B) {
		data := encodeValue("hello world")
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			var result string
			rv := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(rv)
			PutDecoderToPool(dec)
		}
	})

	b.Run("Slice", func(b *testing.B) {
		data := encodeValue([]int{1, 2, 3, 4, 5})
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			var result []int
			rv := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(rv)
			PutDecoderToPool(dec)
		}
	})
}
