package beve

import (
	"encoding/json"
	"testing"
)

type benchStruct struct {
	Name    string
	Age     int
	Active  bool
	Scores  []float64
	Payload map[string]interface{}
}

var (
	benchBytesSink  []byte
	benchStructSink benchStruct
	benchSliceSink  []int32
)

func newBenchStruct() benchStruct {
	return benchStruct{
		Name:   "Ada Lovelace",
		Age:    36,
		Active: true,
		Scores: []float64{0.1, 1.25, 3.5, 42.0},
		Payload: map[string]interface{}{
			"projects": 3,
			"field":    "computing",
		},
	}
}

func newTypedArray() []int32 {
	arr := make([]int32, 1024)
	for i := range arr {
		arr[i] = int32(i)
	}
	return arr
}

func BenchmarkMarshalStruct(b *testing.B) {
	b.ReportAllocs()
	value := newBenchStruct()
	for i := 0; i < b.N; i++ {
		data, err := Marshal(value)
		if err != nil {
			b.Fatalf("Marshal failed: %v", err)
		}
		benchBytesSink = data
	}
}

func BenchmarkMarshalStructJSON(b *testing.B) {
	b.ReportAllocs()
	value := newBenchStruct()
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(value)
		if err != nil {
			b.Fatalf("json.Marshal failed: %v", err)
		}
		benchBytesSink = data
	}
}

func BenchmarkUnmarshalStruct(b *testing.B) {
	b.ReportAllocs()
	data, err := Marshal(newBenchStruct())
	if err != nil {
		b.Fatalf("Marshal setup failed: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var out benchStruct
		if err := Unmarshal(data, &out); err != nil {
			b.Fatalf("Unmarshal failed: %v", err)
		}
		benchStructSink = out
	}
}

func BenchmarkUnmarshalStructJSON(b *testing.B) {
	b.ReportAllocs()
	data, err := json.Marshal(newBenchStruct())
	if err != nil {
		b.Fatalf("json.Marshal setup failed: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var out benchStruct
		if err := json.Unmarshal(data, &out); err != nil {
			b.Fatalf("json.Unmarshal failed: %v", err)
		}
		benchStructSink = out
	}
}

func BenchmarkMarshalTypedArray(b *testing.B) {
	b.ReportAllocs()
	arr := newTypedArray()
	for i := 0; i < b.N; i++ {
		data, err := Marshal(arr)
		if err != nil {
			b.Fatalf("Marshal typed array failed: %v", err)
		}
		benchBytesSink = data
	}
}

func BenchmarkMarshalTypedArrayJSON(b *testing.B) {
	b.ReportAllocs()
	arr := newTypedArray()
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(arr)
		if err != nil {
			b.Fatalf("json.Marshal typed array failed: %v", err)
		}
		benchBytesSink = data
	}
}

func BenchmarkUnmarshalTypedArray(b *testing.B) {
	b.ReportAllocs()
	arr := newTypedArray()
	data, err := Marshal(arr)
	if err != nil {
		b.Fatalf("Marshal setup failed: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var out []int32
		if err := Unmarshal(data, &out); err != nil {
			b.Fatalf("Unmarshal typed array failed: %v", err)
		}
		benchSliceSink = out
	}
}

func BenchmarkUnmarshalTypedArrayJSON(b *testing.B) {
	b.ReportAllocs()
	arr := newTypedArray()
	data, err := json.Marshal(arr)
	if err != nil {
		b.Fatalf("json.Marshal setup failed: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var out []int32
		if err := json.Unmarshal(data, &out); err != nil {
			b.Fatalf("json.Unmarshal typed array failed: %v", err)
		}
		benchSliceSink = out
	}
}
