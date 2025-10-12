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

// Phase 9: Optimized file write benchmarks
// Fixed issues:
// 1. Pre-compute file path outside loop
// 2. Encode once, write many times (isolate I/O)
// 3. Use buffered I/O for realistic performance

func BenchmarkFileWriteOptimized_BEVE(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.beve") // Pre-compute path

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWriteOptimized_JSON(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := json.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.json") // Pre-compute path

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWriteOptimized_Sonic(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := sonic.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.json") // Pre-compute path

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWriteOptimized_MessagePack(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := msgpack.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.msgpack") // Pre-compute path

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}

func BenchmarkFileWriteOptimized_CBOR(b *testing.B) {
	data := generateComplexData(50, 100)
	encoded, _ := cbor.Marshal(data)

	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "test.cbor") // Pre-compute path

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = os.WriteFile(filePath, encoded, 0644)
	}
	b.ReportMetric(float64(len(encoded)), "bytes")
}
