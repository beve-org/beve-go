package core

import (
	"reflect"
	"testing"
)

type benchmarkLevel3 struct {
	Z string `beve:"z"`
}

type benchmarkLevel2 struct {
	Y      string          `beve:"y"`
	Level3 benchmarkLevel3 `beve:",inline"`
}

type benchmarkStruct struct {
	ID      int               `beve:"id"`
	Name    string            `beve:"name"`
	Values  []int             `beve:"values"`
	Labels  []string          `beve:"labels"`
	Inline  benchmarkLevel2   `beve:",inline"`
	Meta    map[string]string `beve:"meta"`
	Enabled bool              `beve:"enabled"`
}

var (
	benchmarkStructInstance = benchmarkStruct{
		ID:     42,
		Name:   "benchmark",
		Values: []int{1, 2, 3, 4, 5, 6, 7, 8},
		Labels: []string{"alpha", "beta", "gamma"},
		Inline: benchmarkLevel2{
			Y: "inline-y",
			Level3: benchmarkLevel3{
				Z: "inline-z",
			},
		},
		Meta: map[string]string{
			"env":  "prod",
			"zone": "us-east-1",
		},
		Enabled: true,
	}
	benchmarkStructValue = reflect.ValueOf(&benchmarkStructInstance).Elem()
	benchmarkStructInfo  = getEncoderStructInfo(benchmarkStructValue.Type())
	benchmarkNumbers     = reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	benchmarkStrings     = reflect.ValueOf([]string{"one", "two", "three", "four", "five"})
)

func BenchmarkEncodeStructFast(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if err := enc.encodeStructFast(benchmarkStructValue); err != nil {
		b.Fatalf("encodeStructFast warmup failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		if err := enc.encodeStructFast(benchmarkStructValue); err != nil {
			b.Fatalf("encodeStructFast failed: %v", err)
		}
	}
}

func BenchmarkCountStructFields(b *testing.B) {
	b.ReportAllocs()

	// Warm caches
	countStructFields(benchmarkStructValue, benchmarkStructInfo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if countStructFields(benchmarkStructValue, benchmarkStructInfo) == 0 {
			b.Fatal("unexpected zero count")
		}
	}
}

func BenchmarkWriteStructFields(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Warm up encoder caches
	if err := writeStructFields(enc, benchmarkStructValue, benchmarkStructInfo); err != nil {
		b.Fatalf("writeStructFields warmup failed: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		if err := writeStructFields(enc, benchmarkStructValue, benchmarkStructInfo); err != nil {
			b.Fatalf("writeStructFields failed: %v", err)
		}
	}
}

func BenchmarkEncodePrimitiveSliceInts(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if err := enc.encodePrimitiveSlice(benchmarkNumbers, reflect.Int); err != nil {
		b.Fatalf("encodePrimitiveSlice warmup failed: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		if err := enc.encodePrimitiveSlice(benchmarkNumbers, reflect.Int); err != nil {
			b.Fatalf("encodePrimitiveSlice failed: %v", err)
		}
	}
}

func BenchmarkEncodePrimitiveSliceStrings(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if err := enc.encodePrimitiveSlice(benchmarkStrings, reflect.String); err != nil {
		b.Fatalf("encodePrimitiveSlice warmup failed: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		if err := enc.encodePrimitiveSlice(benchmarkStrings, reflect.String); err != nil {
			b.Fatalf("encodePrimitiveSlice failed: %v", err)
		}
	}
}

func BenchmarkWriteCompressedUint(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	values := []uint64{3, 64, 1024, 65535, 1 << 20}

	for _, v := range values {
		if err := enc.WriteCompressedUint(v); err != nil {
			b.Fatalf("WriteCompressedUint warmup failed: %v", err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		for _, v := range values {
			if err := enc.WriteCompressedUint(v); err != nil {
				b.Fatalf("WriteCompressedUint failed: %v", err)
			}
		}
	}
}

func BenchmarkBuildEncoderStructFields(b *testing.B) {
	b.ReportAllocs()

	t := benchmarkStructValue.Type()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		info := buildEncoderStructInfo(t)
		if len(info.fields) == 0 {
			b.Fatal("unexpected empty fields")
		}
	}
}
