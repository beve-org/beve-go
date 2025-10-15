//go:build arm64 && !purego
// +build arm64,!purego

package core

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// Test data for UTF-8 validation
var (
	asciiShort    = []byte("Hello, World!")
	asciiLong     = []byte(strings.Repeat("The quick brown fox jumps over the lazy dog. ", 100))
	utf8Short     = []byte("Hello, 世界! 🌍")
	utf8Long      = []byte(strings.Repeat("Hello, 世界! 🌍 Γεια σου κόσμε! مرحبا بالعالم! ", 100))
	invalidUTF8   = []byte{0xFF, 0xFE, 0xFD}
	overlong      = []byte{0xC0, 0x80}             // Overlong encoding of NULL
	surrogatePair = []byte{0xED, 0xA0, 0x80}       // UTF-16 surrogate
	outOfRange    = []byte{0xF4, 0x90, 0x80, 0x80} // > U+10FFFF
)

// TestValidateUTF8SIMD tests the SIMD UTF-8 validator.
func TestValidateUTF8SIMD(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{"Empty", []byte{}, true},
		{"ASCII Short", asciiShort, true},
		{"ASCII Long", asciiLong, true},
		{"UTF-8 Short", utf8Short, true},
		{"UTF-8 Long", utf8Long, true},
		{"Invalid Bytes", invalidUTF8, false},
		{"Overlong Encoding", overlong, false},
		{"Surrogate Pair", surrogatePair, false},
		{"Out of Range", outOfRange, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateUTF8SIMD(tt.input)
			if result != tt.expected {
				t.Errorf("validateUTF8SIMD(%q) = %v, want %v", tt.input, result, tt.expected)
			}

			// Verify against standard library
			stdResult := utf8.Valid(tt.input)
			if result != stdResult {
				t.Errorf("validateUTF8SIMD(%q) = %v, but utf8.Valid = %v", tt.input, result, stdResult)
			}
		})
	}
}

// TestCountUTF8RunesSIMD tests the SIMD rune counter.
func TestCountUTF8RunesSIMD(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"Empty", []byte{}, 0},
		{"ASCII Short", asciiShort, 13},
		{"ASCII Long", asciiLong, 4600},
		{"UTF-8 Short", utf8Short, 13}, // "Hello, 世界! 🌍" = 13 runes
		{"UTF-8 Long", utf8Long, 4100}, // Approximate
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countUTF8RunesSIMD(tt.input)
			stdResult := utf8.RuneCount(tt.input)

			if result != stdResult {
				t.Errorf("countUTF8RunesSIMD(%q) = %d, but utf8.RuneCount = %d", tt.input, result, stdResult)
			}
		})
	}
}

// BenchmarkValidateUTF8_Short benchmarks UTF-8 validation for short strings.
func BenchmarkValidateUTF8_Short(b *testing.B) {
	b.Run("SIMD/ASCII", func(b *testing.B) {
		b.SetBytes(int64(len(asciiShort)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = validateUTF8SIMD(asciiShort)
		}
	})

	b.Run("Stdlib/ASCII", func(b *testing.B) {
		b.SetBytes(int64(len(asciiShort)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.Valid(asciiShort)
		}
	})

	b.Run("SIMD/UTF8", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Short)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = validateUTF8SIMD(utf8Short)
		}
	})

	b.Run("Stdlib/UTF8", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Short)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.Valid(utf8Short)
		}
	})
}

// BenchmarkValidateUTF8_Long benchmarks UTF-8 validation for long strings.
func BenchmarkValidateUTF8_Long(b *testing.B) {
	b.Run("SIMD/ASCII", func(b *testing.B) {
		b.SetBytes(int64(len(asciiLong)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = validateUTF8SIMD(asciiLong)
		}
	})

	b.Run("Stdlib/ASCII", func(b *testing.B) {
		b.SetBytes(int64(len(asciiLong)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.Valid(asciiLong)
		}
	})

	b.Run("SIMD/UTF8", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Long)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = validateUTF8SIMD(utf8Long)
		}
	})

	b.Run("Stdlib/UTF8", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Long)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.Valid(utf8Long)
		}
	})
}

// BenchmarkCountUTF8Runes benchmarks rune counting.
func BenchmarkCountUTF8Runes(b *testing.B) {
	b.Run("SIMD/Short", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Short)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = countUTF8RunesSIMD(utf8Short)
		}
	})

	b.Run("Stdlib/Short", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Short)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.RuneCount(utf8Short)
		}
	})

	b.Run("SIMD/Long", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Long)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = countUTF8RunesSIMD(utf8Long)
		}
	})

	b.Run("Stdlib/Long", func(b *testing.B) {
		b.SetBytes(int64(len(utf8Long)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = utf8.RuneCount(utf8Long)
		}
	})
}
