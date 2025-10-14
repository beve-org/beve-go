package core

import (
	"testing"
)

// TestSIMDInt32Array tests SIMD-accelerated int32 array encoding
func TestSIMDInt32Array(t *testing.T) {
	tests := []struct {
		name string
		data []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"threshold", []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{"large", make([]int32, 100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize test data
			for i := range tt.data {
				tt.data[i] = int32(i * 10)
			}

			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			// Encode using SIMD path
			if err := enc.encodeSIMDInt32Array(tt.data); err != nil {
				t.Fatalf("encodeSIMDInt32Array failed: %v", err)
			}

			// Verify encoding succeeded
			if enc.Buf.Len() == 0 {
				t.Fatal("encoded data is empty")
			}

			t.Logf("Encoded %d elements to %d bytes", len(tt.data), enc.Buf.Len())
		})
	}
}

// TestSIMDFloat64Array tests SIMD-accelerated float64 array encoding
func TestSIMDFloat64Array(t *testing.T) {
	tests := []struct {
		name string
		data []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{3.14}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"threshold", []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8}},
		{"large", make([]float64, 100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize test data
			for i := range tt.data {
				tt.data[i] = float64(i) * 1.5
			}

			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			// Encode using SIMD path
			if err := enc.encodeSIMDFloat64Array(tt.data); err != nil {
				t.Fatalf("encodeSIMDFloat64Array failed: %v", err)
			}

			// Verify encoding succeeded
			if enc.Buf.Len() == 0 {
				t.Fatal("encoded data is empty")
			}

			t.Logf("Encoded %d elements to %d bytes", len(tt.data), enc.Buf.Len())
		})
	}
}

// BenchmarkSIMDInt32Array benchmarks SIMD vs scalar int32 encoding
func BenchmarkSIMDInt32Array(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 1024}

	for _, size := range sizes {
		data := make([]int32, size)
		for i := range data {
			data[i] = int32(i)
		}

		b.Run("SIMD/size="+string(rune(size)), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(size * 4)) // 4 bytes per int32

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				enc.encodeSIMDInt32Array(data)
				PutEncoderToPool(enc)
			}
		})

		b.Run("Scalar/size="+string(rune(size)), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(size * 4))

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				enc.encodeInt32ArrayScalar(data)
				PutEncoderToPool(enc)
			}
		})
	}
}

// BenchmarkSIMDFloat64Array benchmarks SIMD vs scalar float64 encoding
func BenchmarkSIMDFloat64Array(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 1024}

	for _, size := range sizes {
		data := make([]float64, size)
		for i := range data {
			data[i] = float64(i) * 1.5
		}

		b.Run("SIMD/size="+string(rune(size)), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(size * 8)) // 8 bytes per float64

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				enc.encodeSIMDFloat64Array(data)
				PutEncoderToPool(enc)
			}
		})

		b.Run("Scalar/size="+string(rune(size)), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(size * 8))

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				enc.encodeFloat64ArrayScalar(data)
				PutEncoderToPool(enc)
			}
		})
	}
}

// TestSIMDCapabilityDetection tests CPU feature detection
func TestSIMDCapabilityDetection(t *testing.T) {
	t.Logf("SIMD Enabled: %v", UseSIMD)
	t.Logf("Has AVX2: %v", HasAVX2)
	t.Logf("Has NEON: %v", HasNEON)

	// At least one should be true on modern systems
	if !UseSIMD && !HasAVX2 && !HasNEON {
		t.Log("SIMD not available on this platform (fallback to scalar)")
	}
}
