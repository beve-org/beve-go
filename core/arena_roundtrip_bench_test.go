package core

import (
	"reflect"
	"testing"
)

// BenchmarkArenaRoundtrip benchmarks full encode→decode cycle with arena.
//
// This measures the total allocation reduction when using arena for both
// encoding and decoding operations.
func BenchmarkArenaRoundtrip(b *testing.B) {
	type TestStruct struct {
		ID     int64     `beve:"id"`
		Name   string    `beve:"name"`
		Values []float64 `beve:"values"`
		Flags  []bool    `beve:"flags"`
	}

	original := TestStruct{
		ID:     12345,
		Name:   "Test Data",
		Values: make([]float64, 100),
		Flags:  make([]bool, 100),
	}
	for i := range original.Values {
		original.Values[i] = float64(i) * 3.14
		original.Flags[i] = i%2 == 0
	}

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Encode
			enc := GetEncoderFromPool()
			err := enc.Encode(reflect.ValueOf(original))
			if err != nil {
				b.Fatal(err)
			}
			
			// Get encoded data
			data := make([]byte, len(enc.Buf.Bytes()))
			copy(data, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			// Decode
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

		arena := NewArena(16 * 1024) // 16KB arena for both encode+decode
		defer arena.Free()

		for i := 0; i < b.N; i++ {
			arena.Reset()

			// Encode with arena
			enc := GetEncoderFromPoolWithArena(arena)
			err := enc.Encode(reflect.ValueOf(original))
			if err != nil {
				b.Fatal(err)
			}

			// Get encoded data (copy to avoid arena lifetime issues)
			data := make([]byte, len(enc.Buf.Bytes()))
			copy(data, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			// Decode with arena
			dec := NewDecoderWithArena(data, arena)
			var result TestStruct
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})

	b.Run("with_separate_arenas", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		encArena := NewArena(8 * 1024)
		decArena := NewArena(8 * 1024)
		defer encArena.Free()
		defer decArena.Free()

		for i := 0; i < b.N; i++ {
			encArena.Reset()
			decArena.Reset()

			// Encode
			enc := GetEncoderFromPoolWithArena(encArena)
			err := enc.Encode(reflect.ValueOf(original))
			if err != nil {
				b.Fatal(err)
			}

			data := make([]byte, len(enc.Buf.Bytes()))
			copy(data, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			// Decode
			dec := NewDecoderWithArena(data, decArena)
			var result TestStruct
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkArenaRoundtrip_LargePayload tests arena benefits on large data.
func BenchmarkArenaRoundtrip_LargePayload(b *testing.B) {
	type Record struct {
		ID     int64   `beve:"id"`
		Data   []int32 `beve:"data"`
		Scores []float64 `beve:"scores"`
	}

	type Dataset struct {
		Records []Record `beve:"records"`
		Count   int64    `beve:"count"`
	}

	// Create dataset with 100 records
	dataset := Dataset{
		Records: make([]Record, 100),
		Count:   100,
	}
	for i := range dataset.Records {
		dataset.Records[i] = Record{
			ID:     int64(i),
			Data:   make([]int32, 50),
			Scores: make([]float64, 50),
		}
		for j := range dataset.Records[i].Data {
			dataset.Records[i].Data[j] = int32(j)
			dataset.Records[i].Scores[j] = float64(j) * 1.5
		}
	}

	b.Run("without_arena", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			err := enc.Encode(reflect.ValueOf(dataset))
			if err != nil {
				b.Fatal(err)
			}

			data := make([]byte, len(enc.Buf.Bytes()))
			copy(data, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			dec := NewDecoder(data)
			var result Dataset
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

		arena := NewArena(128 * 1024) // 128KB arena for large payload
		defer arena.Free()

		for i := 0; i < b.N; i++ {
			arena.Reset()

			enc := GetEncoderFromPoolWithArena(arena)
			err := enc.Encode(reflect.ValueOf(dataset))
			if err != nil {
				b.Fatal(err)
			}

			data := make([]byte, len(enc.Buf.Bytes()))
			copy(data, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			dec := NewDecoderWithArena(data, arena)
			var result Dataset
			v := reflect.ValueOf(&result).Elem()
			if err := dec.Decode(v); err != nil {
				b.Fatal(err)
			}
			PutDecoderToPool(dec)
		}
	})
}

// BenchmarkArenaRoundtrip_ArenaPoolComparison tests arena pool vs create/destroy.
func BenchmarkArenaRoundtrip_ArenaPoolComparison(b *testing.B) {
	type SimpleData struct {
		ID    int64  `beve:"id"`
		Value string `beve:"value"`
	}

	data := SimpleData{
		ID:    42,
		Value: "test data",
	}

	b.Run("arena_create_destroy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			arena := NewArena(1024)
			
			enc := GetEncoderFromPoolWithArena(arena)
			enc.Encode(reflect.ValueOf(data))
			encoded := make([]byte, len(enc.Buf.Bytes()))
			copy(encoded, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			dec := NewDecoderWithArena(encoded, arena)
			var result SimpleData
			dec.Decode(reflect.ValueOf(&result).Elem())
			PutDecoderToPool(dec)

			arena.Free()
		}
	})

	b.Run("arena_pool_reuse", func(b *testing.B) {
		pool := NewArenaPool(1024)
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			arena := pool.Get()
			
			enc := GetEncoderFromPoolWithArena(arena)
			enc.Encode(reflect.ValueOf(data))
			encoded := make([]byte, len(enc.Buf.Bytes()))
			copy(encoded, enc.Buf.Bytes())
			PutEncoderToPool(enc)

			dec := NewDecoderWithArena(encoded, arena)
			var result SimpleData
			dec.Decode(reflect.ValueOf(&result).Elem())
			PutDecoderToPool(dec)

			pool.Put(arena)
		}
	})
}
