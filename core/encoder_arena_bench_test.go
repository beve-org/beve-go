package core

import (
	"reflect"
	"testing"
)

// BenchmarkEncoderArena_SmallStruct benchmarks encoding small structs with/without arena.
func BenchmarkEncoderArena_SmallStruct(b *testing.B) {
	type SmallStruct struct {
		ID   int64  `beve:"id"`
		Name string `beve:"name"`
		Age  int32  `beve:"age"`
	}

	data := SmallStruct{
		ID:   12345,
		Name: "John Doe",
		Age:  30,
	}

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			err := enc.Encode(reflect.ValueOf(data))
			if err != nil {
				b.Fatal(err)
			}
			PutEncoderToPool(enc)
		}
	})

	b.Run("with_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		arena := NewArena(1024) // 1KB arena
		defer arena.Free()

		for i := 0; i < b.N; i++ {
			arena.Reset()
			enc := GetEncoderFromPoolWithArena(arena)
			err := enc.Encode(reflect.ValueOf(data))
			if err != nil {
				b.Fatal(err)
			}
			PutEncoderToPool(enc)
		}
	})
}

// BenchmarkEncoderArena_TypedArray benchmarks encoding typed arrays with/without arena.
func BenchmarkEncoderArena_TypedArray(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		testData := make([]int32, size)
		for i := range testData {
			testData[i] = int32(i)
		}

		b.Run("without_arena_size_"+string(rune('0'+size/100)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				err := enc.Encode(reflect.ValueOf(testData))
				if err != nil {
					b.Fatal(err)
				}
				PutEncoderToPool(enc)
			}
		})

		b.Run("with_arena_size_"+string(rune('0'+size/100)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			arena := NewArena(size * 8)
			defer arena.Free()

			for i := 0; i < b.N; i++ {
				arena.Reset()
				enc := GetEncoderFromPoolWithArena(arena)
				err := enc.Encode(reflect.ValueOf(testData))
				if err != nil {
					b.Fatal(err)
				}
				PutEncoderToPool(enc)
			}
		})
	}
}

// BenchmarkEncoderArena_MixedWorkload benchmarks encoding complex structs.
func BenchmarkEncoderArena_MixedWorkload(b *testing.B) {
	type ComplexStruct struct {
		IDs    []int64   `beve:"ids"`
		Scores []float64 `beve:"scores"`
		Flags  []bool    `beve:"flags"`
		Name   string    `beve:"name"`
	}

	data := ComplexStruct{
		IDs:    make([]int64, 50),
		Scores: make([]float64, 50),
		Flags:  make([]bool, 50),
		Name:   "Complex Data Structure",
	}
	for i := range data.IDs {
		data.IDs[i] = int64(i)
		data.Scores[i] = float64(i) * 1.5
		data.Flags[i] = i%2 == 0
	}

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			err := enc.Encode(reflect.ValueOf(data))
			if err != nil {
				b.Fatal(err)
			}
			PutEncoderToPool(enc)
		}
	})

	b.Run("with_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		arena := NewArena(8 * 1024) // 8KB arena
		defer arena.Free()

		for i := 0; i < b.N; i++ {
			arena.Reset()
			enc := GetEncoderFromPoolWithArena(arena)
			err := enc.Encode(reflect.ValueOf(data))
			if err != nil {
				b.Fatal(err)
			}
			PutEncoderToPool(enc)
		}
	})
}

// BenchmarkEncoderArena_Overhead measures arena overhead for encoder.
func BenchmarkEncoderArena_Overhead(b *testing.B) {
	b.Run("get_encoder_from_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			PutEncoderToPool(enc)
		}
	})

	b.Run("get_encoder_with_arena", func(b *testing.B) {
		arena := NewArena(1024)
		defer arena.Free()

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			arena.Reset()
			enc := GetEncoderFromPoolWithArena(arena)
			PutEncoderToPool(enc)
		}
	})
}
