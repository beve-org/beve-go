package beve

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

// File I/O benchmarks with realistic data

func BenchmarkFileWrite_BEVE(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := Marshal(data)

	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "test.beve")
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWrite_JSON(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := json.Marshal(data)

	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "test.json")
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWrite_Sonic(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := sonic.Marshal(data)

	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "test.json")
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWrite_MessagePack(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := msgpack.Marshal(data)

	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "test.msgpack")
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWrite_CBOR(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := cbor.Marshal(data)

	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "test.cbor")
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

// ==================== FILE READ BENCHMARKS ====================

func BenchmarkFileRead_BEVE(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.beve")
	_ = os.WriteFile(filePath, encoded, 0644)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fileData, _ := os.ReadFile(filePath)
		var result ComplexData
		_ = Unmarshal(fileData, &result)
	}
}

func BenchmarkFileRead_JSON(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := json.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.json")
	_ = os.WriteFile(filePath, encoded, 0644)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fileData, _ := os.ReadFile(filePath)
		var result ComplexData
		_ = json.Unmarshal(fileData, &result)
	}
}

func BenchmarkFileRead_Sonic(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := sonic.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.json")
	_ = os.WriteFile(filePath, encoded, 0644)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fileData, _ := os.ReadFile(filePath)
		var result ComplexData
		_ = sonic.Unmarshal(fileData, &result)
	}
}

func BenchmarkFileRead_MessagePack(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := msgpack.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.msgpack")
	_ = os.WriteFile(filePath, encoded, 0644)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fileData, _ := os.ReadFile(filePath)
		var result ComplexData
		_ = msgpack.Unmarshal(fileData, &result)
	}
}

func BenchmarkFileRead_CBOR(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := cbor.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.cbor")
	_ = os.WriteFile(filePath, encoded, 0644)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fileData, _ := os.ReadFile(filePath)
		var result ComplexData
		_ = cbor.Unmarshal(fileData, &result)
	}
}

// ==================== ROUND-TRIP BENCHMARKS ====================

func BenchmarkRoundTrip_BEVE(b *testing.B) {
	data := generateComplexData(20, 40)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, _ := Marshal(data)
		var result ComplexData
		_ = Unmarshal(encoded, &result)
	}
}

func BenchmarkRoundTrip_JSON(b *testing.B) {
	data := generateComplexData(20, 40)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, _ := json.Marshal(data)
		var result ComplexData
		_ = json.Unmarshal(encoded, &result)
	}
}

func BenchmarkRoundTrip_Sonic(b *testing.B) {
	data := generateComplexData(20, 40)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, _ := sonic.Marshal(data)
		var result ComplexData
		_ = sonic.Unmarshal(encoded, &result)
	}
}

func BenchmarkRoundTrip_MessagePack(b *testing.B) {
	data := generateComplexData(20, 40)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, _ := msgpack.Marshal(data)
		var result ComplexData
		_ = msgpack.Unmarshal(encoded, &result)
	}
}

func BenchmarkRoundTrip_CBOR(b *testing.B) {
	data := generateComplexData(20, 40)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoded, _ := cbor.Marshal(data)
		var result ComplexData
		_ = cbor.Unmarshal(encoded, &result)
	}
}

// ==================== STREAMING BENCHMARKS ====================
// Note: These test file I/O + encoding together. For pure streaming performance,
// see stream_bench_test.go which uses in-memory buffers.

func BenchmarkStream_BEVE(b *testing.B) {
	data := generateComplexData(10, 20)
	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "stream.beve")
		f, _ := os.Create(filePath)

		// Use StreamEncoder for better performance
		stream := NewStreamEncoder(f)
		_ = stream.Encode(data)
		_ = stream.Close()

		f.Close()
	}
}

func BenchmarkStream_JSON(b *testing.B) {
	data := generateComplexData(10, 20)
	tmpDir := b.TempDir()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		filePath := filepath.Join(tmpDir, "stream.json")
		f, _ := os.Create(filePath)
		enc := json.NewEncoder(f)
		_ = enc.Encode(data)
		f.Close()
	}
}
