package beve

import (
	"testing"
)

// Comparison benchmarks with other binary formats
// Run: go test -bench=BenchmarkComparison -benchmem

type BenchStruct struct {
	Name  string
	Age   int
	Score float64
	Tags  []string
}

func getBenchStruct() BenchStruct {
	return BenchStruct{
		Name:  "TestUser",
		Age:   30,
		Score: 95.5,
		Tags:  []string{"go", "rust", "python"},
	}
}

func BenchmarkComparisonBEVE_Marshal(b *testing.B) {
	data := getBenchStruct()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComparisonBEVE_Unmarshal(b *testing.B) {
	data := getBenchStruct()
	encoded, _ := Marshal(data)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var out BenchStruct
		if err := Unmarshal(encoded, &out); err != nil {
			b.Fatal(err)
		}
	}
}

// Large array benchmark
func BenchmarkComparisonBEVE_LargeArray(b *testing.B) {
	data := make([]int64, 1000)
	for i := range data {
		data[i] = int64(i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, err := Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
		var out []int64
		if err := Unmarshal(encoded, &out); err != nil {
			b.Fatal(err)
		}
	}
}

// Memory efficiency test
func BenchmarkComparisonBEVE_Memory(b *testing.B) {
	data := getBenchStruct()
	b.ResetTimer()
	b.ReportAllocs()

	var totalBytes int64
	for i := 0; i < b.N; i++ {
		encoded, err := Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
		totalBytes += int64(len(encoded))
	}
	b.ReportMetric(float64(totalBytes)/float64(b.N), "bytes/op")
}
