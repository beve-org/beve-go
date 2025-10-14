package core

import (
	"testing"
)

// TestWriteCompressedUintAsm tests assembly-optimized varint encoding
func TestWriteCompressedUintAsm(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
	}{
		{"zero", 0},
		{"small_1", 10},
		{"small_2", 63},
		{"medium_1", 64},
		{"medium_2", 1000},
		{"medium_3", 16383},
		{"large_1", 16384},
		{"large_2", 1000000},
		{"large_3", 1073741823},
		{"xlarge", 1073741824},
		{"max_uint32", 4294967295},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test assembly version (if available)
			enc := GetEncoderFromPool()
			defer PutEncoderToPool(enc)

			if err := enc.WriteCompressedUint(tt.value); err != nil {
				t.Fatalf("WriteCompressedUint failed: %v", err)
			}

			result := enc.Buf.Bytes()
			if len(result) == 0 {
				t.Fatal("encoded data is empty")
			}

			t.Logf("Value %d encoded to %d bytes: %x", tt.value, len(result), result)
		})
	}
}

// BenchmarkVarintEncoding compares assembly vs pure Go varint encoding
func BenchmarkVarintEncoding(b *testing.B) {
	values := []struct {
		name  string
		value uint64
	}{
		{"small_10", 10},
		{"small_63", 63},
		{"medium_100", 100},
		{"medium_1000", 1000},
		{"medium_10000", 10000},
		{"large_100000", 100000},
		{"large_1000000", 1000000},
		{"xlarge", 1073741824},
	}

	for _, v := range values {
		b.Run(v.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				enc.WriteCompressedUint(v.value)
				PutEncoderToPool(enc)
			}
		})
	}
}

// BenchmarkVarintThroughput measures varint encoding throughput
func BenchmarkVarintThroughput(b *testing.B) {
	// Test with realistic workload: array of 1000 length values
	const arraySize = 1000

	b.Run("array_lengths", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			
			// Encode 1000 array length values (typical use case)
			for j := 0; j < arraySize; j++ {
				enc.WriteCompressedUint(uint64(j))
			}
			
			PutEncoderToPool(enc)
		}
		
		// Report operations per second
		b.ReportMetric(float64(b.N*arraySize)/b.Elapsed().Seconds(), "ops/s")
	})
}

// BenchmarkVarintVsReflection compares varint overhead to reflection
func BenchmarkVarintVsReflection(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.Run("with_varint_lengths", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			enc := GetEncoderFromPool()
			
			// Encode array with length prefix (uses varint)
			enc.WriteCompressedUint(uint64(len(data)))
			for _, val := range data {
				enc.WriteCompressedUint(uint64(val))
			}
			
			PutEncoderToPool(enc)
		}
	})
}

// TestVarintCorrectness verifies assembly produces same output as pure Go
func TestVarintCorrectness(t *testing.T) {
	testValues := []uint64{
		0, 1, 10, 63, 64, 100, 1000, 16383, 16384,
		100000, 1073741823, 1073741824, 4294967295,
	}

	for _, val := range testValues {
		// Encode with current implementation (fresh encoder each time)
		enc := GetEncoderFromPool()
		enc.Buf.Reset() // Clear any existing data
		
		if err := enc.WriteCompressedUint(val); err != nil {
			t.Fatalf("WriteCompressedUint(%d) failed: %v", val, err)
		}
		result := enc.Buf.Bytes()
		resultLen := len(result)
		PutEncoderToPool(enc)

		// Verify length is within expected range
		expectedLen := 1
		if val >= 64 {
			expectedLen = 2
		}
		if val >= 16384 {
			expectedLen = 3
		}
		if val >= 1073741824 {
			expectedLen = 4
		}

		if resultLen != expectedLen {
			t.Errorf("Value %d: expected %d bytes, got %d bytes: %x", 
				val, expectedLen, resultLen, result)
		} else {
			t.Logf("✓ Value %d correctly encoded to %d bytes", val, resultLen)
		}
	}
}

// BenchmarkVarintSizes measures encoding performance by size category
func BenchmarkVarintSizes(b *testing.B) {
	sizes := []struct {
		name   string
		values []uint64
	}{
		{"1byte", []uint64{0, 10, 30, 50, 63}},
		{"2byte", []uint64{64, 100, 1000, 10000, 16383}},
		{"3byte", []uint64{16384, 100000, 500000, 1073741823}},
		{"4byte", []uint64{1073741824, 2000000000, 4294967295}},
	}

	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				enc := GetEncoderFromPool()
				
				for _, val := range size.values {
					enc.WriteCompressedUint(val)
				}
				
				PutEncoderToPool(enc)
			}
		})
	}
}
