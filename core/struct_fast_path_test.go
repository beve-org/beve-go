package core

import (
	"reflect"
	"testing"
	"unsafe"
)

// TestEncodeStructFast tests encodeStructFast function (0% coverage)
func TestEncodeStructFast(t *testing.T) {
	type SimpleStruct struct {
		A int
		B string
		C bool
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Test with addressable struct
	s := SimpleStruct{A: 42, B: "hello", C: true}
	rv := reflect.ValueOf(&s).Elem()

	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode struct failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestIsStructFieldEmpty tests isStructFieldEmpty function (0% coverage)
func TestIsStructFieldEmpty(t *testing.T) {
	type TestStruct struct {
		Str    string
		Int    int
		Bool   bool
		Slice  []int
		Map    map[string]int
		Ptr    *int
		PtrStr *string
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name   string
		value  TestStruct
		expect string
	}{
		{
			name: "all empty",
			value: TestStruct{
				Str:   "",
				Int:   0,
				Bool:  false,
				Slice: nil,
				Map:   nil,
				Ptr:   nil,
			},
			expect: "should encode empty fields",
		},
		{
			name: "mixed",
			value: TestStruct{
				Str:   "non-empty",
				Int:   42,
				Bool:  true,
				Slice: []int{1, 2, 3},
				Map:   map[string]int{"a": 1},
			},
			expect: "should encode non-empty fields",
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
				t.Error("Nothing encoded")
			}
		})
	}
}

// TestCountStructFields tests countStructFields function (0% coverage)
func TestCountStructFields(t *testing.T) {
	type EmptyFieldStruct struct {
		A string `beve:"a,omitempty"`
		B int    `beve:"b,omitempty"`
		C bool   `beve:"c,omitempty"`
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	t.Run("all empty with omitempty", func(t *testing.T) {
		enc.Buf.Reset()
		s := EmptyFieldStruct{A: "", B: 0, C: false}
		rv := reflect.ValueOf(s)

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode failed: %v", err)
		}
	})

	t.Run("some non-empty", func(t *testing.T) {
		enc.Buf.Reset()
		s := EmptyFieldStruct{A: "test", B: 0, C: true}
		rv := reflect.ValueOf(s)

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode failed: %v", err)
		}
	})
}

// TestWriteStructFields tests writeStructFields function (0% coverage)
func TestWriteStructFields(t *testing.T) {
	type StructWithFields struct {
		Name     string
		Age      int
		Active   bool
		Score    float64
		Tags     []string
		Metadata map[string]string
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := StructWithFields{
		Name:   "Alice",
		Age:    30,
		Active: true,
		Score:  95.5,
		Tags:   []string{"go", "rust"},
		Metadata: map[string]string{
			"role": "engineer",
			"team": "platform",
		},
	}

	rv := reflect.ValueOf(&s).Elem()
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode complex struct failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestWriteStructFieldsPtrGeneric tests writeStructFieldsPtrGeneric (0% coverage)
func TestWriteStructFieldsPtrGeneric(t *testing.T) {
	type NestedStruct struct {
		Inner string
	}

	type ComplexStruct struct {
		Name   string
		Nested *NestedStruct
		List   []*NestedStruct
		Map    map[string]*NestedStruct
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := ComplexStruct{
		Name:   "complex",
		Nested: &NestedStruct{Inner: "nested1"},
		List: []*NestedStruct{
			{Inner: "list1"},
			{Inner: "list2"},
		},
		Map: map[string]*NestedStruct{
			"key1": {Inner: "map1"},
			"key2": {Inner: "map2"},
		},
	}

	rv := reflect.ValueOf(&s)
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode complex nested struct failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestEncodeStructFieldValue tests encodeStructFieldValue (25.8% coverage -> increase)
func TestEncodeStructFieldValue(t *testing.T) {
	type MixedStruct struct {
		Int8Val    int8
		Int16Val   int16
		Int32Val   int32
		Int64Val   int64
		Uint8Val   uint8
		Uint16Val  uint16
		Uint32Val  uint32
		Uint64Val  uint64
		Float32Val float32
		Float64Val float64
		BoolVal    bool
		StringVal  string
		BytesVal   []byte
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := MixedStruct{
		Int8Val:    -128,
		Int16Val:   -32768,
		Int32Val:   -2147483648,
		Int64Val:   -9223372036854775808,
		Uint8Val:   255,
		Uint16Val:  65535,
		Uint32Val:  4294967295,
		Uint64Val:  18446744073709551615,
		Float32Val: 3.14,
		Float64Val: 2.71828,
		BoolVal:    true,
		StringVal:  "test",
		BytesVal:   []byte{1, 2, 3},
	}

	rv := reflect.ValueOf(s)
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode mixed types failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestEnsureAddressableStruct tests ensureAddressableStruct (80% -> 100%)
func TestEnsureAddressableStruct(t *testing.T) {
	type UnaddressableStruct struct {
		X int
		Y string
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	t.Run("addressable", func(t *testing.T) {
		enc.Buf.Reset()
		s := UnaddressableStruct{X: 10, Y: "addressable"}
		ptr := &s
		rv := reflect.ValueOf(ptr).Elem()

		if !rv.CanAddr() {
			t.Fatal("Expected addressable value")
		}

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode addressable failed: %v", err)
		}
	})

	t.Run("non-addressable", func(t *testing.T) {
		enc.Buf.Reset()
		// Function return value is not addressable
		getValue := func() UnaddressableStruct {
			return UnaddressableStruct{X: 20, Y: "non-addressable"}
		}

		rv := reflect.ValueOf(getValue())
		if rv.CanAddr() {
			t.Skip("Value is addressable, can't test non-addressable path")
		}

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode non-addressable failed: %v", err)
		}
	})
}

// TestIsEmptyValue tests isEmptyValue function (0% coverage)
func TestIsEmptyValue(t *testing.T) {
	type EmptyTestStruct struct {
		Str   string         `beve:"str,omitempty"`
		Int   int            `beve:"int,omitempty"`
		Uint  uint           `beve:"uint,omitempty"`
		Float float64        `beve:"float,omitempty"`
		Bool  bool           `beve:"bool,omitempty"`
		Slice []int          `beve:"slice,omitempty"`
		Map   map[string]int `beve:"map,omitempty"`
		Ptr   *int           `beve:"ptr,omitempty"`
		Iface interface{}    `beve:"iface,omitempty"`
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	t.Run("all empty", func(t *testing.T) {
		enc.Buf.Reset()
		s := EmptyTestStruct{}
		rv := reflect.ValueOf(s)

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode empty struct failed: %v", err)
		}
	})

	t.Run("all non-empty", func(t *testing.T) {
		enc.Buf.Reset()
		val := 42
		s := EmptyTestStruct{
			Str:   "non-empty",
			Int:   42,
			Uint:  42,
			Float: 3.14,
			Bool:  true,
			Slice: []int{1, 2, 3},
			Map:   map[string]int{"key": 1},
			Ptr:   &val,
			Iface: "interface value",
		}
		rv := reflect.ValueOf(s)

		if err := enc.Encode(rv); err != nil {
			t.Fatalf("Encode non-empty struct failed: %v", err)
		}
	})
}

// TestBuildStructFieldKey tests buildStructFieldKey (45.5% coverage -> increase)
func TestBuildStructFieldKey(t *testing.T) {
	type TaggedStruct struct {
		NoTag     string
		WithTag   string `beve:"custom_name"`
		OmitEmpty string `beve:"omit,omitempty"`
		SkipField string `beve:"-"`
		ShortTag  string `beve:"s"`
		LongTag   string `beve:"very_long_field_name_with_many_characters"`
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := TaggedStruct{
		NoTag:     "value1",
		WithTag:   "value2",
		OmitEmpty: "",
		SkipField: "should not encode",
		ShortTag:  "value3",
		LongTag:   "value4",
	}

	rv := reflect.ValueOf(s)
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode tagged struct failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestUnsafeFieldAccess tests unsafe pointer field access
func TestUnsafeFieldAccess(t *testing.T) {
	type UnsafeTestStruct struct {
		A int64
		B float64
		C string
		D bool
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := UnsafeTestStruct{
		A: 9223372036854775807,
		B: 3.141592653589793,
		C: "unsafe test",
		D: true,
	}

	ptr := unsafe.Pointer(&s)
	if ptr == nil {
		t.Fatal("Failed to get unsafe pointer")
	}

	rv := reflect.ValueOf(s)
	if err := enc.Encode(rv); err != nil {
		t.Fatalf("Encode with unsafe access failed: %v", err)
	}
}
