// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Benchmarks comparing Assembly vs Pure Go implementations.

package core

import (
	"testing"
)

// Benchmark WriteCompressedUint with small values (<64)
// This is the most common case (array lengths, string lengths, etc.)
func BenchmarkWriteCompressedUint_Small_Assembly(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.WriteCompressedUint(uint64(i % 64))
	}
}

// Benchmark WriteCompressedUint with medium values (64-16383)
func BenchmarkWriteCompressedUint_Medium_Assembly(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.WriteCompressedUint(uint64(1000 + (i % 15000)))
	}
}

// Benchmark WriteCompressedUint with large values (>16384)
func BenchmarkWriteCompressedUint_Large_Assembly(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.WriteCompressedUint(uint64(100000 + i))
	}
}

// Benchmark mixed workload (realistic distribution)
// Most values are small (<64), some medium, few large
func BenchmarkWriteCompressedUint_Mixed_Assembly(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	values := []uint64{
		10, 20, 30, 5, 15, 63, // Small (80%)
		100, 500, 1000, 5000, // Medium (15%)
		50000, 100000, // Large (5%)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		enc.Buf.Reset()
		_ = enc.WriteCompressedUint(values[i%len(values)])
	}
}

// Test correctness of assembly implementation
func TestWriteCompressedUintAsm_Correctness(t *testing.T) {
	testCases := []struct {
		name     string
		value    uint64
		expected []byte
	}{
		{"zero", 0, []byte{0}},
		{"small_10", 10, []byte{40}},
		{"boundary_63", 63, []byte{252}},
		{"boundary_64", 64, []byte{1, 64}},
		{"medium_100", 100, []byte{1, 100}},
		{"medium_1000", 1000, []byte{13, 232}},
		{"boundary_16383", 16383, []byte{253, 255}},
		{"boundary_16384", 16384, []byte{2, 64, 0}},
		{"large_100000", 100000, []byte{6, 134, 160}},
		{"boundary_1073741823", 1073741823, []byte{254, 255, 255}},
		{"boundary_1073741824", 1073741824, []byte{3, 0, 0, 0}},
		{"max_uint32", 4294967295, []byte{255, 255, 255, 255}},
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			enc.Buf.Reset()

			err := enc.WriteCompressedUint(tc.value)
			if err != nil {
				t.Fatalf("WriteCompressedUint(%d) failed: %v", tc.value, err)
			}

			got := enc.Buf.Bytes()
			if len(got) != len(tc.expected) {
				t.Errorf("Length mismatch for %d: got %d bytes, want %d bytes", tc.value, len(got), len(tc.expected))
				t.Errorf("  Got:      %v", got)
				t.Errorf("  Expected: %v", tc.expected)
				return
			}

			for i := range got {
				if got[i] != tc.expected[i] {
					t.Errorf("Byte %d mismatch for %d: got 0x%02X, want 0x%02X", i, tc.value, got[i], tc.expected[i])
					t.Errorf("  Got:      %v", got)
					t.Errorf("  Expected: %v", tc.expected)
					return
				}
			}
		})
	}
}

// Fuzz test to ensure assembly matches expected behavior
func FuzzWriteCompressedUintAsm(f *testing.F) {
	// Seed corpus with interesting values
	seeds := []uint64{
		0, 1, 10, 63, 64, 100, 255, 256,
		16383, 16384, 65535, 65536,
		1073741823, 1073741824, 4294967295,
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	f.Fuzz(func(t *testing.T, n uint64) {
		// Limit to 32-bit values for this encoding
		if n > 4294967295 {
			return
		}

		enc.Buf.Reset()
		err := enc.WriteCompressedUint(n)
		if err != nil {
			t.Fatalf("WriteCompressedUint(%d) failed: %v", n, err)
		}

		result := enc.Buf.Bytes()

		// Validate length constraints based on value ranges
		var expectedLen int
		if n < 64 {
			expectedLen = 1
		} else if n < 16384 {
			expectedLen = 2
		} else if n < 1073741824 {
			expectedLen = 3
		} else {
			expectedLen = 4
		}

		if len(result) != expectedLen {
			t.Errorf("n=%d should produce %d byte(s), got %d", n, expectedLen, len(result))
		}
	})
}
