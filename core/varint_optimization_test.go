package core

import (
	"testing"
)

func TestVarintSizeLookup(t *testing.T) {
	tests := []struct {
		value    uint64
		expected int
	}{
		{0, 1},
		{63, 1},
		{64, 2},
		{16383, 2},
		{16384, 3},
		{65535, 3},
		{65536, 3},
		{1073741823, 3},
		{1073741824, 4},
		{^uint64(0), 4}, // Max uint64
	}

	for _, tt := range tests {
		got := varintSize(tt.value)
		if got != tt.expected {
			t.Errorf("varintSize(%d) = %d, want %d", tt.value, got, tt.expected)
		}
	}
}

func TestVarintCaching(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Cache some varints
	values := []uint64{10, 100, 1000, 10000, 100000}
	sizes := make([]int, len(values))

	for i, v := range values {
		sizes[i] = enc.cacheVarintSize(v)
	}

	// Verify cache count
	if enc.varintCacheCount != len(values) {
		t.Errorf("cache count = %d, want %d", enc.varintCacheCount, len(values))
	}

	// Verify cached values
	for i, v := range values {
		if enc.varintCache[i] != v {
			t.Errorf("cache[%d] = %d, want %d", i, enc.varintCache[i], v)
		}
	}

	// Verify sizes
	for i, v := range values {
		expected := varintSize(v)
		if sizes[i] != expected {
			t.Errorf("size[%d] = %d, want %d", i, sizes[i], expected)
		}
	}

	// Clear and verify
	enc.clearVarintCache()
	if enc.varintCacheCount != 0 {
		t.Errorf("after clear, cache count = %d, want 0", enc.varintCacheCount)
	}
}

func TestVarintCacheOverflow(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	// Try to cache more than capacity (8)
	for i := 0; i < 20; i++ {
		_ = enc.cacheVarintSize(uint64(i))
	}

	// Should max out at 8
	if enc.varintCacheCount != 8 {
		t.Errorf("cache count = %d, want 8 (max capacity)", enc.varintCacheCount)
	}

	// First 8 values should be cached
	for i := 0; i < 8; i++ {
		if enc.varintCache[i] != uint64(i) {
			t.Errorf("cache[%d] = %d, want %d", i, enc.varintCache[i], i)
		}
	}
}

// Benchmark varintSize lookup table vs branching
func BenchmarkVarintSize_Small(b *testing.B) {
	values := []uint64{10, 20, 30, 40, 50}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range values {
			_ = varintSize(v)
		}
	}
}

func BenchmarkVarintSize_Medium(b *testing.B) {
	values := []uint64{100, 1000, 10000, 15000, 16000}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range values {
			_ = varintSize(v)
		}
	}
}

func BenchmarkVarintSize_Large(b *testing.B) {
	values := []uint64{100000, 1000000, 10000000, 100000000, 1000000000}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range values {
			_ = varintSize(v)
		}
	}
}

// Benchmark caching vs double encoding
func BenchmarkVarintCaching_WithCache(b *testing.B) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	values := []uint64{10, 100, 1000, 10000, 100000}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.clearVarintCache()
		
		// Phase 1: Calculate sizes (with caching)
		totalSize := 0
		for _, v := range values {
			totalSize += enc.cacheVarintSize(v)
		}
		
		// Phase 2: Write (using cache - not measured here, just clearing)
		enc.clearVarintCache()
	}
}

func BenchmarkVarintCaching_WithoutCache(b *testing.B) {
	values := []uint64{10, 100, 1000, 10000, 100000}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Phase 1: Calculate sizes
		totalSize := 0
		for _, v := range values {
			totalSize += varintSize(v)
		}
		
		// Phase 2: Would re-calculate sizes for encoding (simulated)
		for _, v := range values {
			_ = varintSize(v)
		}
	}
}

// End-to-end benchmark: Struct encoding with varint caching
type BenchStruct struct {
	Name    string
	Age     int
	Email   string
	Address string
	Phone   string
}

func BenchmarkStructEncoding_Varints(b *testing.B) {
	s := BenchStruct{
		Name:    "John Doe",
		Age:     30,
		Email:   "john@example.com",
		Address: "123 Main St",
		Phone:   "555-1234",
	}
	
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		enc.clearVarintCache()
		enc.Buf.data = enc.Buf.data[:0]
		
		// Simulate struct encoding (field count + field name lengths)
		_ = enc.cacheVarintSize(5) // Field count
		_ = enc.cacheVarintSize(uint64(len(s.Name)))
		_ = enc.cacheVarintSize(uint64(len(s.Email)))
		_ = enc.cacheVarintSize(uint64(len(s.Address)))
		_ = enc.cacheVarintSize(uint64(len(s.Phone)))
	}
}
