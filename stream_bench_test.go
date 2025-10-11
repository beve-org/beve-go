package beve

import (
	"bytes"
	"encoding/json"
	"testing"
)

// BenchmarkStreamEncoder_BEVE tests optimized streaming encoder
func BenchmarkStreamEncoder_BEVE(b *testing.B) {
	type Record struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	records := make([]Record, 100)
	for i := range records {
		records[i] = Record{ID: i, Name: "User", Age: 25 + i%50}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		stream := NewStreamEncoder(&buf)

		for j := range records {
			if err := stream.Encode(records[j]); err != nil {
				b.Fatal(err)
			}
		}

		if err := stream.Close(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStreamEncoder_JSON tests JSON streaming for comparison
func BenchmarkStreamEncoder_JSON(b *testing.B) {
	type Record struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	records := make([]Record, 100)
	for i := range records {
		records[i] = Record{ID: i, Name: "User", Age: 25 + i%50}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)

		for j := range records {
			if err := enc.Encode(records[j]); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkStreamEncoder_SingleRecord tests single record encoding
func BenchmarkStreamEncoder_SingleRecord(b *testing.B) {
	type Record struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	record := Record{ID: 1, Name: "User", Age: 30}

	b.Run("BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			stream := NewStreamEncoder(&buf)
			_ = stream.Encode(record)
			_ = stream.Close()
		}
	})

	b.Run("JSON", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			_ = enc.Encode(record)
		}
	})
}

// BenchmarkStreamEncoder_LargeRecords tests large record streaming
func BenchmarkStreamEncoder_LargeRecords(b *testing.B) {
	type Record struct {
		ID       int               `beve:"id"`
		Name     string            `beve:"name"`
		Email    string            `beve:"email"`
		Tags     []string          `beve:"tags"`
		Metadata map[string]string `beve:"metadata"`
	}

	records := make([]Record, 1000)
	for i := range records {
		records[i] = Record{
			ID:    i,
			Name:  "User Name Long String",
			Email: "user@example.com",
			Tags:  []string{"tag1", "tag2", "tag3"},
			Metadata: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		}
	}

	b.Run("BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			stream := NewStreamEncoder(&buf)

			for j := range records {
				if err := stream.Encode(records[j]); err != nil {
					b.Fatal(err)
				}
			}

			if err := stream.Close(); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("JSON", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)

			for j := range records {
				if err := enc.Encode(records[j]); err != nil {
					b.Fatal(err)
				}
			}
		}
	})
}

// BenchmarkStreamEncoder_BufferSizes tests different buffer sizes
func BenchmarkStreamEncoder_BufferSizes(b *testing.B) {
	type Record struct {
		ID   int    `beve:"id"`
		Name string `beve:"name"`
	}

	records := make([]Record, 100)
	for i := range records {
		records[i] = Record{ID: i, Name: "User"}
	}

	bufferSizes := []int{1024, 4096, 8192, 16384, 32768}

	for _, size := range bufferSizes {
		b.Run(string(rune(size)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var buf bytes.Buffer
				stream := NewStreamEncoderSize(&buf, size)

				for j := range records {
					_ = stream.Encode(records[j])
				}

				_ = stream.Close()
			}
		})
	}
}

// BenchmarkStreamEncoder_Primitives tests primitive type streaming
func BenchmarkStreamEncoder_Primitives(b *testing.B) {
	ints := make([]int, 100)
	for i := range ints {
		ints[i] = i
	}

	b.Run("Integers", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			stream := NewStreamEncoder(&buf)

			for j := range ints {
				_ = stream.Encode(ints[j])
			}

			_ = stream.Close()
		}
	})

	strings := make([]string, 100)
	for i := range strings {
		strings[i] = "test string"
	}

	b.Run("Strings", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			stream := NewStreamEncoder(&buf)

			for j := range strings {
				_ = stream.Encode(strings[j])
			}

			_ = stream.Close()
		}
	})
}
