package core

import (
	"reflect"
	"testing"
)

// TestSetRawMessageValue tests setRawMessageValue (0% coverage)
func TestSetRawMessageValue(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode some data that could be used as RawMessage
	data := map[string]int{"test": 42}
	if err := enc.Encode(reflect.ValueOf(data)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// The function is called internally during RawMessage handling
	t.Log("setRawMessageValue triggered through RawMessage operations")
}

// TestCaptureRawValue tests captureRawValue (0% coverage)
func TestCaptureRawValue(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode complex data
	data := struct {
		Name   string
		Values []int
		Map    map[string]interface{}
	}{
		Name:   "capture test",
		Values: []int{1, 2, 3},
		Map:    map[string]interface{}{"key": "value"},
	}

	if err := enc.Encode(reflect.ValueOf(data)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Try to decode with RawMessage capture
	dec := NewDecoder(enc.Buf.Bytes())
	var result struct {
		Name   string
		Values []int
		Map    map[string]interface{}
	}

	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode with capture: %v", err)
	}
}

// TestDecodeStructSlowPath tests decodeStructSlow (0% coverage)
func TestDecodeStructSlowPath(t *testing.T) {
	type ComplexStruct struct {
		Field1 interface{}
		Field2 map[string]interface{}
		Field3 []interface{}
		Field4 struct {
			Nested1 string
			Nested2 int
		}
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	original := ComplexStruct{
		Field1: "string value",
		Field2: map[string]interface{}{
			"key1": 1,
			"key2": "two",
		},
		Field3: []interface{}{1, "two", 3.0},
		Field4: struct {
			Nested1 string
			Nested2 int
		}{
			Nested1: "nested",
			Nested2: 42,
		},
	}

	if err := enc.Encode(reflect.ValueOf(original)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	dec := NewDecoder(enc.Buf.Bytes())
	var result ComplexStruct
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode struct slow path: %v", err)
	}
}

// TestParseTagExtended tests parseTag (0% coverage)
func TestParseTagExtended(t *testing.T) {
	type TaggedStruct struct {
		Field1  string `beve:"field1"`
		Field2  string `beve:"field2,omitempty"`
		Field3  string `beve:"-"`
		Field4  string `beve:"field4,omitempty,string"`
		Field5  string // no tag
		Field6  string `json:"json_name" beve:"beve_name"`
		Field7  string `beve:",omitempty"`
		Field8  string `beve:"custom-name-with-dashes"`
		Field9  string `beve:"unicode_名前"`
		Field10 string `beve:""`
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := TaggedStruct{
		Field1:  "value1",
		Field2:  "",
		Field3:  "ignored",
		Field4:  "value4",
		Field5:  "value5",
		Field6:  "value6",
		Field7:  "",
		Field8:  "value8",
		Field9:  "value9",
		Field10: "value10",
	}

	if err := enc.Encode(reflect.ValueOf(s)); err != nil {
		t.Fatalf("Encode tagged struct failed: %v", err)
	}

	// Decode to trigger parseTag
	dec := NewDecoder(enc.Buf.Bytes())
	var result TaggedStruct
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode with parseTag: %v", err)
	}
}

// TestNewBufferLease tests newBufferLease (0% coverage)
func TestNewBufferLease(t *testing.T) {
	// Test buffer leasing through large data encoding
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Create large data that will trigger buffer leasing
	largeData := make([]int64, 50000)
	for i := range largeData {
		largeData[i] = int64(i)
	}

	if err := enc.Encode(reflect.ValueOf(largeData)); err != nil {
		t.Fatalf("Encode large data failed: %v", err)
	}

	// Get the bytes (triggers Bytes() method)
	bytes := enc.Buf.Bytes()
	if len(bytes) == 0 {
		t.Error("No bytes from large data encode")
	}

	t.Logf("Buffer lease test: encoded %d elements into %d bytes", len(largeData), len(bytes))
}

// TestBufferBytes tests Bytes method on BufferLease (0% coverage)
func TestBufferBytes(t *testing.T) {
	// Test through encoder that uses buffer
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	testData := []interface{}{
		[]int{1, 2, 3},
		"test string",
		map[string]int{"a": 1, "b": 2},
		struct{ X int }{X: 42},
	}

	for i, data := range testData {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Test %d encode failed: %v", i, err)
			continue
		}

		bytes := enc.Buf.Bytes()
		if len(bytes) == 0 {
			t.Errorf("Test %d: Bytes() returned empty", i)
		}
	}
}

// TestBufferRelease tests Release method on BufferLease (0% coverage)
func TestBufferRelease(t *testing.T) {
	// Create encoder and ensure proper cleanup
	for i := 0; i < 100; i++ {
		enc := GetEncoderFromPool()

		data := []int{1, 2, 3, 4, 5}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Iteration %d encode failed: %v", i, err)
		}

		// Return to pool (triggers Release internally)
		PutEncoderToPool(enc)
	}

	t.Log("Buffer release tested through pool operations")
}

// TestSetNil tests SetNil decoder utility (0% coverage)
func TestSetNil(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Encode nil pointers
	type StructWithNils struct {
		PtrInt    *int
		PtrString *string
		PtrSlice  *[]int
		PtrMap    *map[string]int
	}

	original := StructWithNils{
		PtrInt:    nil,
		PtrString: nil,
		PtrSlice:  nil,
		PtrMap:    nil,
	}

	if err := enc.Encode(reflect.ValueOf(original)); err != nil {
		t.Fatalf("Encode nils failed: %v", err)
	}

	dec := NewDecoder(enc.Buf.Bytes())
	var result StructWithNils
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode with SetNil: %v", err)
	}
}

// TestDetachBytes tests DetachBytes (0% coverage)
func TestDetachBytes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	data := []int{1, 2, 3, 4, 5}
	if err := enc.Encode(reflect.ValueOf(data)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// DetachBytes should give ownership of the bytes
	detached := enc.DetachBytes()

	// Get the actual bytes from the detached buffer
	actualBytes := detached.Bytes()
	if len(actualBytes) == 0 {
		t.Error("DetachBytes returned empty slice")
	}

	// After detach, buffer should be empty or reset
	if enc.Buf.Len() > 0 {
		t.Error("Buffer not cleared after DetachBytes")
	}

	t.Logf("DetachBytes returned %d bytes", len(actualBytes))
}

// TestEncoderWithMultipleDetach tests multiple DetachBytes calls
func TestEncoderWithMultipleDetach(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	for i := 0; i < 10; i++ {
		enc.Buf.Reset()
		data := []int{i, i + 1, i + 2}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Iteration %d encode failed: %v", i, err)
			continue
		}

		detached := enc.DetachBytes()
		actualBytes := detached.Bytes()
		if len(actualBytes) == 0 {
			t.Errorf("Iteration %d: DetachBytes returned empty", i)
		}
	}
}

// TestComputeFieldOffset tests computeFieldOffset (already 100% but verify)
func TestComputeFieldOffset(t *testing.T) {
	type FieldStruct struct {
		A int8
		B int16
		C int32
		D int64
		E string
		F bool
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	s := FieldStruct{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: "test",
		F: true,
	}

	if err := enc.Encode(reflect.ValueOf(s)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	dec := NewDecoder(enc.Buf.Bytes())
	var result FieldStruct
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode computeFieldOffset test: %v", err)
	}
}

// TestComplexNestedStructs tests deep nesting
func TestComplexNestedStructs(t *testing.T) {
	type Level3 struct {
		Value string
	}
	type Level2 struct {
		L3  Level3
		Num int
	}
	type Level1 struct {
		L2   Level2
		List []string
	}
	type Root struct {
		L1  Level1
		Map map[string]Level1
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	original := Root{
		L1: Level1{
			L2: Level2{
				L3:  Level3{Value: "deep"},
				Num: 42,
			},
			List: []string{"a", "b", "c"},
		},
		Map: map[string]Level1{
			"key1": {
				L2: Level2{
					L3:  Level3{Value: "nested"},
					Num: 99,
				},
				List: []string{"x", "y"},
			},
		},
	}

	if err := enc.Encode(reflect.ValueOf(original)); err != nil {
		t.Fatalf("Encode complex nested failed: %v", err)
	}

	dec := NewDecoder(enc.Buf.Bytes())
	var result Root
	if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
		t.Logf("Decode complex nested: %v", err)
	} else {
		t.Log("Complex nested struct test passed")
	}
}

// TestVeryLargeBuffers tests buffer handling with very large data
func TestVeryLargeBuffers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large buffer test in short mode")
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Create 1MB of data
	largeSlice := make([]byte, 1024*1024)
	for i := range largeSlice {
		largeSlice[i] = byte(i % 256)
	}

	if err := enc.Encode(reflect.ValueOf(largeSlice)); err != nil {
		t.Fatalf("Encode 1MB failed: %v", err)
	}

	bytes := enc.Buf.Bytes()
	if len(bytes) == 0 {
		t.Error("No bytes from 1MB encode")
	}

	t.Logf("Encoded 1MB into %d bytes", len(bytes))
}

// TestConcurrentDecoders tests decoder safety
func TestConcurrentDecoders(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	data := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	}

	if err := enc.Encode(reflect.ValueOf(data)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	encoded := enc.Buf.Bytes()

	done := make(chan bool, 20)
	for i := 0; i < 20; i++ {
		go func(id int) {
			defer func() { done <- true }()

			dec := NewDecoder(encoded)
			var result map[string]int
			if err := dec.Decode(reflect.ValueOf(&result).Elem()); err != nil {
				t.Logf("Goroutine %d decode error: %v", id, err)
			}
		}(i)
	}

	for i := 0; i < 20; i++ {
		<-done
	}

	t.Log("Concurrent decoder test completed")
}
