package core

import (
	"math"
	"reflect"
	"testing"
)

// TestEncoderPooling tests encoder pooling functionality
func TestEncoderPooling(t *testing.T) {
	enc1 := GetEncoderFromPool()
	if enc1 == nil {
		t.Fatal("GetEncoderFromPool returned nil")
	}

	PutEncoderToPool(enc1)

	enc2 := GetEncoderFromPool()
	if enc2 == nil {
		t.Fatal("GetEncoderFromPool returned nil after pool return")
	}

	// Should be the same instance (pooled)
	if enc1 != enc2 {
		t.Log("Note: Encoder instances are different (pool may have grown)")
	}

	PutEncoderToPool(enc2)
}

// TestEncoderBasicTypes tests encoding of basic Go types
func TestEncoderBasicTypes(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"nil", nil},
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
			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if !rv.IsValid() {
				// Handle nil case
				if err := enc.EncodeNull(); err != nil {
					t.Fatalf("EncodeNull failed: %v", err)
				}
				return
			}

			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if enc.Buf == nil || enc.Buf.Len() == 0 {
				t.Error("Encoder produced no output")
			}
		})
	}
}

// TestEncoderSlices tests slice encoding
func TestEncoderSlices(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"slice_int", []int{1, 2, 3}},
		{"slice_string", []string{"a", "b", "c"}},
		{"slice_float", []float64{1.1, 2.2, 3.3}},
		{"slice_bool", []bool{true, false, true}},
		{"slice_empty", []int{}},
		{"array_int", [3]int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if enc.Buf == nil || enc.Buf.Len() == 0 {
				t.Error("Encoder produced no output")
			}
		})
	}
}

// TestEncoderMaps tests map encoding
func TestEncoderMaps(t *testing.T) {
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
			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if enc.Buf == nil || enc.Buf.Len() == 0 {
				t.Error("Encoder produced no output")
			}
		})
	}
}

// TestEncoderStructs tests struct encoding
func TestEncoderStructs(t *testing.T) {
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
			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if enc.Buf == nil || enc.Buf.Len() == 0 {
				t.Error("Encoder produced no output")
			}
		})
	}
}

// TestEncoderTypedArrays tests typed array encoding
func TestEncoderTypedArrays(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		slice interface{}
	}{
		{"int8", []int8{1, 2, 3}},
		{"int16", []int16{100, 200, 300}},
		{"int32", []int32{1000, 2000, 3000}},
		{"int64", []int64{10000, 20000, 30000}},
		{"uint8", []uint8{1, 2, 3}},
		{"uint16", []uint16{100, 200, 300}},
		{"uint32", []uint32{1000, 2000, 3000}},
		{"uint64", []uint64{10000, 20000, 30000}},
		{"float32", []float32{1.1, 2.2, 3.3}},
		{"float64", []float64{1.1, 2.2, 3.3}},
		{"bool", []bool{true, false, true}},
		{"string", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(tt.slice)
			if err := enc.encodeSlice(rv); err != nil {
				t.Fatalf("encodeSlice failed: %v", err)
			}

			if enc.Buf == nil || enc.Buf.Len() == 0 {
				t.Error("Encoder produced no output")
			}
		})
	}
}

// TestEncoderEdgeCases tests edge cases and error conditions
func TestEncoderEdgeCases(t *testing.T) {
	t.Run("nil_pointer", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		if enc.Buf != nil {
			enc.Buf.Reset()
		}

		var ptr *int
		rv := reflect.ValueOf(ptr)
		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encoding nil pointer failed: %v", err)
		}
	})

	t.Run("nil_interface", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		if enc.Buf != nil {
			enc.Buf.Reset()
		}

		var iface interface{}
		rv := reflect.ValueOf(&iface).Elem()
		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encoding nil interface failed: %v", err)
		}
	})

	t.Run("large_numbers", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		values := []interface{}{
			int64(math.MaxInt64),
			int64(math.MinInt64),
			uint64(math.MaxUint64),
			float64(math.MaxFloat64),
			float64(math.SmallestNonzeroFloat64),
		}

		for _, v := range values {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(v)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encoding %v failed: %v", v, err)
			}
		}
	})

	t.Run("special_floats", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		values := []float64{
			math.NaN(),
			math.Inf(1),
			math.Inf(-1),
		}

		for _, v := range values {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}

			rv := reflect.ValueOf(v)
			if err := enc.Encode(rv); err != nil {
				t.Fatalf("Encoding %v failed: %v", v, err)
			}
		}
	})
}

// TestBufferGrowth tests buffer growth behavior
func TestBufferGrowth(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if enc.Buf == nil {
		t.Skip("Encoder has no buffer")
	}

	initialLen := enc.Buf.Len()

	// Write enough data to trigger growth
	largeData := make([]byte, 4096)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	if err := enc.WriteBytes(largeData); err != nil {
		t.Fatalf("WriteBytes failed: %v", err)
	}

	if enc.Buf.Len() <= initialLen {
		t.Errorf("Buffer did not grow: initial=%d, current=%d", initialLen, enc.Buf.Len())
	}
}

// BenchmarkEncoderPool benchmarks encoder pooling overhead
func BenchmarkEncoderPool(b *testing.B) {
	b.Run("GetAndPut", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			PutEncoderToPool(enc)
		}
	})

	b.Run("NoPool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = &Encoder{}
		}
	})
}

// BenchmarkEncoderTypes benchmarks encoding different types
func BenchmarkEncoderTypes(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	b.Run("Int", func(b *testing.B) {
		b.ReportAllocs()
		rv := reflect.ValueOf(42)
		for i := 0; i < b.N; i++ {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}
			_ = enc.Encode(rv)
		}
	})

	b.Run("String", func(b *testing.B) {
		b.ReportAllocs()
		rv := reflect.ValueOf("hello world")
		for i := 0; i < b.N; i++ {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}
			_ = enc.Encode(rv)
		}
	})

	b.Run("Slice", func(b *testing.B) {
		b.ReportAllocs()
		rv := reflect.ValueOf([]int{1, 2, 3, 4, 5})
		for i := 0; i < b.N; i++ {
			if enc.Buf != nil {
				enc.Buf.Reset()
			}
			_ = enc.Encode(rv)
		}
	})
}
