package core

import (
	"reflect"
	"testing"
)

// BenchmarkFastPathVsReflection compares fast path vs reflection-based decoding.
//
// This benchmark validates the performance impact of typed slice decoders
// that bypass reflection overhead using unsafe pointers.

func BenchmarkFastPathInt32(b *testing.B) {
	// Test data: []int32 with 100 elements
	data := make([]int32, 100)
	for i := range data {
		data[i] = int32(i * 100)
	}

	// Encode once
	enc := GetEncoderFromPool()
	enc.Encode(reflect.ValueOf(data))
	encoded := enc.Buf.Bytes()
	PutEncoderToPool(enc)

	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(encoded)
			result := make([]int32, 0, 100)
			resultVal := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(resultVal)
			PutDecoderToPool(dec)
		}
	})

	b.Run("WithReflection", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(encoded)
			// Force reflection path by using interface{} decode
			var result interface{}
			_ = dec.Decode(reflect.ValueOf(&result).Elem())
			PutDecoderToPool(dec)
		}
	})
}

func BenchmarkFastPathUint64(b *testing.B) {
	data := make([]uint64, 100)
	for i := range data {
		data[i] = uint64(i * 1000)
	}

	enc := GetEncoderFromPool()
	enc.Encode(reflect.ValueOf(data))
	encoded := enc.Buf.Bytes()
	PutEncoderToPool(enc)

	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(encoded)
			result := make([]uint64, 0, 100)
			resultVal := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(resultVal)
			PutDecoderToPool(dec)
		}
	})
}

func BenchmarkFastPathUint8(b *testing.B) {
	data := make([]uint8, 100)
	for i := range data {
		data[i] = uint8(i)
	}

	enc := GetEncoderFromPool()
	enc.Encode(reflect.ValueOf(data))
	encoded := enc.Buf.Bytes()
	PutEncoderToPool(enc)

	b.Run("FastPath", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			dec := NewDecoder(encoded)
			result := make([]uint8, 0, 100)
			resultVal := reflect.ValueOf(&result).Elem()
			_ = dec.Decode(resultVal)
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkFastPathSizes tests performance across different array sizes
func BenchmarkFastPathSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int32, size)
		for i := range data {
			data[i] = int32(i)
		}

		enc := GetEncoderFromPool()
		enc.Encode(reflect.ValueOf(data))
		encoded := enc.Buf.Bytes()
		PutEncoderToPool(enc)

		b.Run("Size"+string(rune('0'+size/10)), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(size * 4)) // 4 bytes per int32
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				dec := NewDecoder(encoded)
				result := make([]int32, 0, size)
				resultVal := reflect.ValueOf(&result).Elem()
				_ = dec.Decode(resultVal)
				PutDecoderToPool(dec)
			}
		})
	}
}
