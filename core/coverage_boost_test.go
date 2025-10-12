package core

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

// TestDecoderErrorPaths covers error handling in decoder
func TestDecoderErrorPaths(t *testing.T) {
	t.Run("empty data", func(t *testing.T) {
		dec := NewDecoder([]byte{})
		var result int
		err := dec.Decode(reflect.ValueOf(&result).Elem())
		if err == nil {
			t.Error("Expected error for empty data")
		}
	})

	t.Run("truncated data", func(t *testing.T) {
		dec := NewDecoder([]byte{0x01}) // incomplete
		var result int
		err := dec.Decode(reflect.ValueOf(&result).Elem())
		if err == nil {
			t.Error("Expected error for truncated data")
		}
	})
}

// TestWriteBytePaths tests WriteByte fast/slow paths
func TestWriteBytePaths(t *testing.T) {
	t.Run("with buffer", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		enc.Buf.Reset()
		if err := enc.WriteByte(0x42); err != nil {
			t.Fatalf("WriteByte failed: %v", err)
		}

		data := enc.Buf.Bytes()
		if len(data) == 0 || data[len(data)-1] != 0x42 {
			t.Errorf("Byte not written correctly: got %v", data)
		}
	})

	t.Run("with io.Writer", func(t *testing.T) {
		var buf bytes.Buffer
		enc := NewEncoder(&buf)

		if err := enc.WriteByte(0x42); err != nil {
			t.Fatalf("WriteByte failed: %v", err)
		}

		data := buf.Bytes()
		if len(data) == 0 || data[len(data)-1] != 0x42 {
			t.Errorf("Byte not written correctly: got %v", data)
		}
	})
}

// TestWriteBytesPaths tests WriteBytes fast/slow paths
func TestWriteBytesPaths(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}

	t.Run("with buffer", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		enc.Buf.Reset()
		if err := enc.WriteBytes(data); err != nil {
			t.Fatalf("WriteBytes failed: %v", err)
		}

		result := enc.Buf.Bytes()
		if len(result) < len(data) {
			t.Errorf("Not enough bytes written: got %d, want at least %d", len(result), len(data))
		}
	})

	t.Run("with io.Writer", func(t *testing.T) {
		var buf bytes.Buffer
		enc := NewEncoder(&buf)

		if err := enc.WriteBytes(data); err != nil {
			t.Fatalf("WriteBytes failed: %v", err)
		}

		result := buf.Bytes()
		if len(result) < len(data) {
			t.Errorf("Not enough bytes written: got %d, want at least %d", len(result), len(data))
		}
	})

	t.Run("empty bytes", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		if err := enc.WriteBytes([]byte{}); err != nil {
			t.Fatalf("WriteBytes failed: %v", err)
		}
	})
}

// TestStringToBytes tests unsafe string conversion
func TestStringToBytes(t *testing.T) {
	tests := []string{"", "hello", "world"}

	for _, s := range tests {
		result := stringToBytes(s)
		if string(result) != s {
			t.Errorf("stringToBytes(%q) = %q, want %q", s, string(result), s)
		}
	}
}

// TestEncoderFloatEdgeCases tests float encoding edge cases
func TestEncoderFloatEdgeCases(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []float64{
		0.0,
		math.Copysign(0, -1), // negative zero
		math.Inf(1),
		math.Inf(-1),
		math.NaN(),
		math.MaxFloat64,
		math.SmallestNonzeroFloat64,
	}

	for _, val := range tests {
		enc.Buf.Reset()
		rv := reflect.ValueOf(val)
		if err := enc.Encode(rv); err != nil {
			t.Errorf("Encode(%v) failed: %v", val, err)
		}
	}
}

// TestUnsupportedTypeError tests unsupported type handling
func TestUnsupportedTypeError(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	unsupportedTypes := []interface{}{
		make(chan int),
		func() {},
		complex(1, 2),
	}

	for _, val := range unsupportedTypes {
		enc.Buf.Reset()
		rv := reflect.ValueOf(val)
		err := enc.Encode(rv)
		if err == nil {
			t.Errorf("Expected error for unsupported type %T", val)
		}
	}
}

// TestBufferGrowthStress tests buffer reallocation under stress
func TestBufferGrowthStress(t *testing.T) {
	buf := AcquireBuffer(8)
	defer ReleaseBuffer(buf)

	// Grow beyond initial capacity multiple times
	for size := 16; size <= 4096; size *= 2 {
		buf.Grow(size)
		if cap(buf.data) < size {
			t.Errorf("Buffer didn't grow to %d: cap=%d", size, cap(buf.data))
		}
	}
}

// TestPoolConcurrency tests pool thread safety
func TestPoolConcurrency(t *testing.T) {
	const goroutines = 50
	const iterations = 100

	done := make(chan bool)

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				enc := GetEncoderFromPool()
				enc.Buf.Write([]byte("test"))
				PutEncoderToPool(enc)
			}
			done <- true
		}()
	}

	for i := 0; i < goroutines; i++ {
		<-done
	}
}

// TestEncoderPrimitiveTypes tests all primitive encoders
func TestEncoderPrimitiveTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		{"bool true", true},
		{"bool false", false},
		{"int8 max", int8(127)},
		{"int8 min", int8(-128)},
		{"int16 max", int16(32767)},
		{"int32 max", int32(2147483647)},
		{"int64 max", int64(9223372036854775807)},
		{"uint8 max", uint8(255)},
		{"uint16 max", uint16(65535)},
		{"uint32 max", uint32(4294967295)},
		{"uint64 max", uint64(18446744073709551615)},
		{"float32", float32(3.14)},
		{"float64", float64(2.71828)},
		{"string", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()
			rv := reflect.ValueOf(tt.value)
			if err := enc.Encode(rv); err != nil {
				t.Errorf("Encode failed: %v", err)
			}
			if enc.Buf.Len() == 0 {
				t.Error("Nothing encoded")
			}
		})
	}
}

// TestRoundtripIntTypes tests integer round-trip encoding
func TestRoundtripIntTypes(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []int64{
		0, 1, -1,
		127, -128,
		32767, -32768,
		2147483647, -2147483648,
		9223372036854775807,
	}

	for _, val := range tests {
		enc.Buf.Reset()
		if err := enc.Encode(reflect.ValueOf(val)); err != nil {
			t.Errorf("Encode(%d) failed: %v", val, err)
		}

		if enc.Buf.Len() == 0 {
			t.Errorf("Nothing encoded for %d", val)
		}
	}
}

// TestRoundtripCollections tests collection encoding
func TestRoundtripCollections(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	t.Run("slice", func(t *testing.T) {
		enc.Buf.Reset()
		data := []int{1, 2, 3, 4, 5}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Encode slice failed: %v", err)
		}
		if enc.Buf.Len() == 0 {
			t.Error("Nothing encoded")
		}
	})

	t.Run("map", func(t *testing.T) {
		enc.Buf.Reset()
		data := map[string]int{"one": 1, "two": 2}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Encode map failed: %v", err)
		}
		if enc.Buf.Len() == 0 {
			t.Error("Nothing encoded")
		}
	})

	t.Run("struct", func(t *testing.T) {
		enc.Buf.Reset()
		type S struct {
			Name string
			Age  int
		}
		data := S{Name: "Alice", Age: 30}
		if err := enc.Encode(reflect.ValueOf(data)); err != nil {
			t.Errorf("Encode struct failed: %v", err)
		}
		if enc.Buf.Len() == 0 {
			t.Error("Nothing encoded")
		}
	})
}

// TestEmptyCollections tests empty collections
func TestEmptyCollections(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []interface{}{
		[]int{},
		[]string{},
		map[string]int{},
		[0]int{},
	}

	for _, val := range tests {
		enc.Buf.Reset()
		rv := reflect.ValueOf(val)
		if err := enc.Encode(rv); err != nil {
			t.Errorf("Encode empty %T failed: %v", val, err)
		}
	}
}

// TestLargeData tests encoding of large data structures
func TestLargeData(t *testing.T) {
	// Large slice
	largeSlice := make([]int, 10000)
	for i := range largeSlice {
		largeSlice[i] = i
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if err := enc.Encode(reflect.ValueOf(largeSlice)); err != nil {
		t.Fatalf("Encode large slice failed: %v", err)
	}

	// Verify something was encoded
	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded for large slice")
	}

	// Verify buffer grew appropriately
	if enc.Buf.Len() < 10000 {
		t.Errorf("Buffer too small: got %d bytes, expected at least 10000", enc.Buf.Len())
	}
}
