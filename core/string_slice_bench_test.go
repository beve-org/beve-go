package core

import (
	"testing"
)

// BenchmarkStringSliceEncoding tests string slice encoding performance
// with the new batched optimization (Phase 13).
func BenchmarkStringSliceEncoding(b *testing.B) {
	// Small strings (typical field names, short text)
	b.Run("small_strings_10", func(b *testing.B) {
		slice := []string{
			"id", "name", "email", "age", "city",
			"country", "zip", "phone", "status", "role",
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})

	// Medium strings (typical user data)
	b.Run("medium_strings_20", func(b *testing.B) {
		slice := make([]string, 20)
		for i := range slice {
			// ~30 chars each
			slice[i] = "user_" + string(rune('a'+i%26)) + "@example.com_test_data_medium_length"
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})

	// Large strings (typical text fields)
	b.Run("large_strings_100", func(b *testing.B) {
		slice := make([]string, 100)
		for i := range slice {
			// ~100 chars each
			slice[i] = "This is a longer text field that represents typical user-generated content like descriptions, comments, or other text data stored in databases and transmitted over networks."
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})

	// Mixed sizes (realistic workload)
	b.Run("mixed_sizes_50", func(b *testing.B) {
		slice := make([]string, 50)
		for i := range slice {
			switch i % 3 {
			case 0:
				slice[i] = "short" // 5 chars
			case 1:
				slice[i] = "medium_length_string_value" // 26 chars
			case 2:
				slice[i] = "This is a much longer string that represents more complex data like descriptions or other verbose fields" // ~100 chars
			}
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})

	// Stress test (1000 strings)
	b.Run("stress_test_1000", func(b *testing.B) {
		slice := make([]string, 1000)
		for i := range slice {
			// Variable length strings (5-50 chars)
			length := 5 + (i % 46)
			slice[i] = string(make([]byte, length))
			for j := range slice[i] {
				slice[i] = slice[i][:j] + string(rune('a'+(i+j)%26)) + slice[i][j+1:]
			}
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})
}

// BenchmarkStringSliceReuse tests encoder reuse pattern
func BenchmarkStringSliceReuse(b *testing.B) {
	slice := []string{
		"id", "name", "email", "age", "city",
		"country", "zip", "phone", "status", "role",
	}

	b.Run("with_pooling", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			enc.encodeStringSliceDirect(slice)
			PutEncoderToPool(enc)
		}
	})

	b.Run("with_reuse", func(b *testing.B) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc.Reset()
			enc.encodeStringSliceDirect(slice)
		}
	})
}
