package beve

import (
	"encoding/json"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

type SimpleStruct struct {
	Name string
	Age  int
}

// BEVE Benchmarks
func BenchmarkSimple_BEVE_Marshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkSimple_BEVE_MarshalZeroCopy(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lease, err := MarshalZeroCopy(data)
		if err != nil {
			b.Fatalf("MarshalZeroCopy failed: %v", err)
		}
		benchBytesSink = lease.Bytes()
		lease.Release()
	}
}

func BenchmarkSimple_BEVE_Unmarshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	encoded, _ := Marshal(data)
	var result SimpleStruct
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Unmarshal(encoded, &result)
	}
}

// JSON Benchmarks
func BenchmarkSimple_JSON_Marshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := json.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkSimple_JSON_Unmarshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	encoded, _ := json.Marshal(data)
	var result SimpleStruct
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(encoded, &result)
	}
}

// MessagePack Benchmarks
func BenchmarkSimple_MsgPack_Marshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := msgpack.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkSimple_MsgPack_Unmarshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	encoded, _ := msgpack.Marshal(data)
	var result SimpleStruct
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(encoded, &result)
	}
}

// CBOR Benchmarks
func BenchmarkSimple_CBOR_Marshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := cbor.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkSimple_CBOR_Unmarshal(b *testing.B) {
	data := SimpleStruct{Name: "John", Age: 30}
	encoded, _ := cbor.Marshal(data)
	var result SimpleStruct
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cbor.Unmarshal(encoded, &result)
	}
}
