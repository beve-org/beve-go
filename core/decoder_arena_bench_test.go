package core

import (
	"reflect"
	"testing"
)

// BenchmarkDecoderArena_CaptureRawValue benchmarks raw value capture with/without arena.
//
// This measures the allocation overhead for BinaryUnmarshaler and RawMessage types.
func BenchmarkDecoderArena_CaptureRawValue(b *testing.B) {
	// Create sample BEVE data (100-byte string)
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)
	
	testString := string(make([]byte, 100))
	enc.EncodeString(testString)
	data := make([]byte, len(enc.Buf.Bytes()))
	copy(data, enc.Buf.Bytes())

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			dec.Pos = 0 // Start at string header
			_, err := dec.captureRawValue()
			if err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})

	b.Run("with_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		arena := NewArena(1024) // 1KB arena (enough for 100-byte string)
		defer arena.Free()
		
		for i := 0; i < b.N; i++ {
			arena.Reset() // Reuse arena for each iteration
			dec := NewDecoderWithArena(data, arena)
			dec.Pos = 0
			_, err := dec.captureRawValue()
			if err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkDecoderArena_TypedArray benchmarks typed array decoding with/without arena.
//
// This measures the allocation overhead for []int32, []float64, etc.
func BenchmarkDecoderArena_TypedArray(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		// Prepare encoded data
		enc := GetEncoderFromPool()
		
		// Create []int32 array
		testData := make([]int32, size)
		for i := range testData {
			testData[i] = int32(i)
		}
		err := enc.Encode(reflect.ValueOf(testData))
		if err != nil {
			b.Fatal(err)
		}
		encodedData := make([]byte, len(enc.Buf.Bytes()))
		copy(encodedData, enc.Buf.Bytes())
		PutEncoderToPool(enc)

		b.Run("without_arena_int32_size_"+string(rune('0'+size/100)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				dec := NewDecoder(encodedData)
				var result []int32
				v := reflect.ValueOf(&result).Elem()
				if err := dec.Decode(v); err != nil {
					b.Fatal(err)
				}
				PutDecoderToPool(dec)
			}
		})

		b.Run("with_arena_int32_size_"+string(rune('0'+size/100)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			
			arena := NewArena(size * 8) // Allocate enough for array
			defer arena.Free()
			
			for i := 0; i < b.N; i++ {
				arena.Reset()
				dec := NewDecoderWithArena(encodedData, arena)
				var result []int32
				v := reflect.ValueOf(&result).Elem()
				if err := dec.Decode(v); err != nil {
					b.Fatal(err)
				}
				PutDecoderToPool(dec)
			}
		})
	}

	// Float64 benchmark
	size := 100
	enc := GetEncoderFromPool()
	
	testFloats := make([]float64, size)
	for i := range testFloats {
		testFloats[i] = float64(i) * 3.14
	}
	enc.Encode(reflect.ValueOf(testFloats))
	floatData := make([]byte, len(enc.Buf.Bytes()))
	copy(floatData, enc.Buf.Bytes())
	PutEncoderToPool(enc)

	b.Run("without_arena_float64", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			dec := NewDecoder(floatData)
			var result []float64
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})

	b.Run("with_arena_float64", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		arena := NewArena(size * 16)
		defer arena.Free()
		
		for i := 0; i < b.N; i++ {
			arena.Reset()
			dec := NewDecoderWithArena(floatData, arena)
			var result []float64
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkDecoderArena_MixedWorkload benchmarks realistic mixed workload.
//
// Simulates decoding a struct with multiple array fields.
func BenchmarkDecoderArena_MixedWorkload(b *testing.B) {
	// Create test struct with multiple arrays
	type TestStruct struct {
		IDs    []int64   `beve:"ids"`
		Scores []float64 `beve:"scores"`
		Flags  []bool    `beve:"flags"`
	}

	original := TestStruct{
		IDs:    make([]int64, 50),
		Scores: make([]float64, 50),
		Flags:  make([]bool, 50),
	}
	for i := range original.IDs {
		original.IDs[i] = int64(i)
		original.Scores[i] = float64(i) * 1.5
		original.Flags[i] = i%2 == 0
	}

	// Encode once
	enc := GetEncoderFromPool()
	enc.Encode(reflect.ValueOf(original))
	data := make([]byte, len(enc.Buf.Bytes()))
	copy(data, enc.Buf.Bytes())
	PutEncoderToPool(enc)

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			dec := NewDecoder(data)
			var result TestStruct
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})

	b.Run("with_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		
		arena := NewArena(8 * 1024) // 8KB arena
		defer arena.Free()
		
		for i := 0; i < b.N; i++ {
			arena.Reset()
			dec := NewDecoderWithArena(data, arena)
			var result TestStruct
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkDecoderArena_Overhead measures arena allocation overhead.
func BenchmarkDecoderArena_Overhead(b *testing.B) {
	b.Run("create_arena", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			arena := NewArena(64 * 1024)
			arena.Free()
		}
	})

	b.Run("arena_pool_get_put", func(b *testing.B) {
		pool := NewArenaPool(64 * 1024)
		b.ReportAllocs()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			arena := pool.Get()
			pool.Put(arena)
		}
	})
}
