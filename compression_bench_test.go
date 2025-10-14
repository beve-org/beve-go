package beve

import (
	"testing"

	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
)

// Test data: Array of 100 user objects
type BenchUser struct {
	ID       string `json:"id" beve:"id"`
	Name     string `json:"name" beve:"name"`
	Email    string `json:"email" beve:"email"`
	Age      int    `json:"age" beve:"age"`
	Active   bool   `json:"active" beve:"active"`
	Balance  float64 `json:"balance" beve:"balance"`
	Address  string `json:"address" beve:"address"`
	Phone    string `json:"phone" beve:"phone"`
	Company  string `json:"company" beve:"company"`
	Country  string `json:"country" beve:"country"`
}

func generateBenchUsers(count int) []BenchUser {
	users := make([]BenchUser, count)
	for i := 0; i < count; i++ {
		users[i] = BenchUser{
			ID:      "user-" + string(rune(i)),
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Age:     30,
			Active:  true,
			Balance: 1234.56,
			Address: "123 Main St, Springfield, IL 62701",
			Phone:   "+1-555-123-4567",
			Company: "Acme Corporation",
			Country: "USA",
		}
	}
	return users
}

// Benchmark: Normal BEVE Marshal
func BenchmarkCompressionMarshal(b *testing.B) {
	users := generateBenchUsers(100)

	b.Run("Normal-BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		var totalSize int64
		for i := 0; i < b.N; i++ {
			data, err := Marshal(users)
			if err != nil {
				b.Fatal(err)
			}
			totalSize += int64(len(data))
		}
		b.ReportMetric(float64(totalSize)/float64(b.N), "bytes/op")
	})

	b.Run("BEVE+LZ4", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		var totalSize int64
		for i := 0; i < b.N; i++ {
			// Marshal to BEVE
			data, err := Marshal(users)
			if err != nil {
				b.Fatal(err)
			}

			// Compress with LZ4
			compressed := make([]byte, lz4.CompressBlockBound(len(data)))
			n, err := lz4.CompressBlock(data, compressed, nil)
			if err != nil {
				b.Fatal(err)
			}
			totalSize += int64(n)
		}
		b.ReportMetric(float64(totalSize)/float64(b.N), "bytes/op")
	})

	b.Run("BEVE+Zstd-Level1", func(b *testing.B) {
		b.ReportAllocs()
		encoder, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
		defer encoder.Close()

		b.ResetTimer()

		var totalSize int64
		for i := 0; i < b.N; i++ {
			// Marshal to BEVE
			data, err := Marshal(users)
			if err != nil {
				b.Fatal(err)
			}

			// Compress with Zstd
			compressed := encoder.EncodeAll(data, make([]byte, 0, len(data)))
			totalSize += int64(len(compressed))
		}
		b.ReportMetric(float64(totalSize)/float64(b.N), "bytes/op")
	})

	b.Run("BEVE+Zstd-Level3", func(b *testing.B) {
		b.ReportAllocs()
		encoder, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
		defer encoder.Close()

		b.ResetTimer()

		var totalSize int64
		for i := 0; i < b.N; i++ {
			// Marshal to BEVE
			data, err := Marshal(users)
			if err != nil {
				b.Fatal(err)
			}

			// Compress with Zstd
			compressed := encoder.EncodeAll(data, make([]byte, 0, len(data)))
			totalSize += int64(len(compressed))
		}
		b.ReportMetric(float64(totalSize)/float64(b.N), "bytes/op")
	})
}

// Benchmark: Unmarshal performance
func BenchmarkCompressionUnmarshal(b *testing.B) {
	users := generateBenchUsers(100)

	// Pre-encode data
	normalData, _ := Marshal(users)

	lz4Data := make([]byte, lz4.CompressBlockBound(len(normalData)))
	lz4Size, _ := lz4.CompressBlock(normalData, lz4Data, nil)
	lz4Data = lz4Data[:lz4Size]

	encoder, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
	zstdData1 := encoder.EncodeAll(normalData, make([]byte, 0, len(normalData)))
	encoder.Close()

	encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	zstdData3 := encoder.EncodeAll(normalData, make([]byte, 0, len(normalData)))
	encoder.Close()

	b.Run("Normal-BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var decoded []BenchUser
			if err := Unmarshal(normalData, &decoded); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("BEVE+LZ4", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Decompress LZ4
			decompressed := make([]byte, len(normalData))
			_, err := lz4.UncompressBlock(lz4Data, decompressed)
			if err != nil {
				b.Fatal(err)
			}

			// Unmarshal BEVE
			var decoded []BenchUser
			if err := Unmarshal(decompressed, &decoded); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("BEVE+Zstd-Level1", func(b *testing.B) {
		b.ReportAllocs()
		decoder, _ := zstd.NewReader(nil)
		defer decoder.Close()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Decompress Zstd
			decompressed, err := decoder.DecodeAll(zstdData1, nil)
			if err != nil {
				b.Fatal(err)
			}

			// Unmarshal BEVE
			var decoded []BenchUser
			if err := Unmarshal(decompressed, &decoded); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("BEVE+Zstd-Level3", func(b *testing.B) {
		b.ReportAllocs()
		decoder, _ := zstd.NewReader(nil)
		defer decoder.Close()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Decompress Zstd
			decompressed, err := decoder.DecodeAll(zstdData3, nil)
			if err != nil {
				b.Fatal(err)
			}

			// Unmarshal BEVE
			var decoded []BenchUser
			if err := Unmarshal(decompressed, &decoded); err != nil {
				b.Fatal(err)
			}
		}
	})
}

// Benchmark different payload sizes
func BenchmarkCompressionPayloadSizes(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		users := generateBenchUsers(size)
		normalData, _ := Marshal(users)

		b.Run("Size-"+string(rune(size))+"/Normal", func(b *testing.B) {
			b.SetBytes(int64(len(normalData)))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var decoded []BenchUser
				_ = Unmarshal(normalData, &decoded)
			}
		})

		b.Run("Size-"+string(rune(size))+"/LZ4", func(b *testing.B) {
			compressed := make([]byte, lz4.CompressBlockBound(len(normalData)))
			n, _ := lz4.CompressBlock(normalData, compressed, nil)
			compressed = compressed[:n]

			b.SetBytes(int64(len(compressed)))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				decompressed := make([]byte, len(normalData))
				_, _ = lz4.UncompressBlock(compressed, decompressed)

				var decoded []BenchUser
				_ = Unmarshal(decompressed, &decoded)
			}
		})
	}
}

// Test compression ratios
func TestCompressionRatios(t *testing.T) {
	users := generateBenchUsers(100)
	normalData, _ := Marshal(users)

	t.Logf("Normal BEVE size: %d bytes", len(normalData))

	// LZ4
	lz4Compressed := make([]byte, lz4.CompressBlockBound(len(normalData)))
	lz4Size, _ := lz4.CompressBlock(normalData, lz4Compressed, nil)
	lz4Ratio := float64(len(normalData)) / float64(lz4Size)
	t.Logf("LZ4 compressed: %d bytes (ratio: %.2fx, saved: %.1f%%)",
		lz4Size, lz4Ratio, (1-float64(lz4Size)/float64(len(normalData)))*100)

	// Zstd Level 1
	encoder1, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
	zstd1 := encoder1.EncodeAll(normalData, make([]byte, 0, len(normalData)))
	encoder1.Close()
	zstd1Ratio := float64(len(normalData)) / float64(len(zstd1))
	t.Logf("Zstd Level 1: %d bytes (ratio: %.2fx, saved: %.1f%%)",
		len(zstd1), zstd1Ratio, (1-float64(len(zstd1))/float64(len(normalData)))*100)

	// Zstd Level 3
	encoder3, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	zstd3 := encoder3.EncodeAll(normalData, make([]byte, 0, len(normalData)))
	encoder3.Close()
	zstd3Ratio := float64(len(normalData)) / float64(len(zstd3))
	t.Logf("Zstd Level 3: %d bytes (ratio: %.2fx, saved: %.1f%%)",
		len(zstd3), zstd3Ratio, (1-float64(len(zstd3))/float64(len(normalData)))*100)

	// Zstd Level 9 (max compression)
	encoder9, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	zstd9 := encoder9.EncodeAll(normalData, make([]byte, 0, len(normalData)))
	encoder9.Close()
	zstd9Ratio := float64(len(normalData)) / float64(len(zstd9))
	t.Logf("Zstd Level 9: %d bytes (ratio: %.2fx, saved: %.1f%%)",
		len(zstd9), zstd9Ratio, (1-float64(len(zstd9))/float64(len(normalData)))*100)
}
