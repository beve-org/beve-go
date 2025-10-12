package core

import (
	"reflect"
	"testing"
)

// BenchmarkDecodeStringTypedArray_Large tests large string arrays (1000 elements)
// This is where the 84MB allocation comes from in the original profiling
func BenchmarkDecodeStringTypedArray_Large(b *testing.B) {
	// Create large string array
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "This is a test string number " // ~29 chars each
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)
	enc.Encode(reflect.ValueOf(data))
	encoded := enc.Buf.Bytes()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dec := NewDecoder(encoded)
		result := reflect.New(reflect.TypeOf(data)).Elem()
		_ = dec.Decode(result)
		PutDecoderToPool(dec)
	}
}

// BenchmarkDecodeStringTypedArray_VeryLarge tests very large string arrays (10000 elements)
func BenchmarkDecodeStringTypedArray_VeryLarge(b *testing.B) {
	// Create very large string array
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "Test" // Small strings
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)
	enc.Encode(reflect.ValueOf(data))
	encoded := enc.Buf.Bytes()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dec := NewDecoder(encoded)
		result := reflect.New(reflect.TypeOf(data)).Elem()
		_ = dec.Decode(result)
		PutDecoderToPool(dec)
	}
}
