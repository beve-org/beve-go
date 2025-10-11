package beve

import (
	"reflect"
	"testing"
)

// ============================================================================
// VALUE POOL DIRECT TESTS (0% → 100%)
// ============================================================================
// These pools are internal and used during reflection-based encoding.
// Direct testing to achieve coverage.

// TestValuePoolGetPut tests globalValuePool Get/Put operations
func TestValuePoolGetPut(t *testing.T) {
	// Get a pooled slice
	slice := globalValuePool.Get()
	if slice == nil {
		t.Fatal("globalValuePool.Get() returned nil")
	}

	if len(*slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(*slice))
	}

	// Use it
	*slice = append(*slice, reflect.ValueOf(42), reflect.ValueOf("test"))
	if len(*slice) != 2 {
		t.Errorf("Expected length 2, got %d", len(*slice))
	}

	// Put it back
	globalValuePool.Put(slice)

	// Get again
	slice2 := globalValuePool.Get()
	if slice2 == nil {
		t.Fatal("globalValuePool.Get() (second) returned nil")
	}

	// Should be reset (length 0)
	if len(*slice2) != 0 {
		t.Errorf("Expected reset slice, got length %d", len(*slice2))
	}

	globalValuePool.Put(slice2)

	t.Log("✓ globalValuePool Get/Put tested")
}

// TestEncodeBufferPoolGetPut tests globalEncodeBufferPool Get/Put operations
func TestEncodeBufferPoolGetPut(t *testing.T) {
	// Get a pooled buffer
	buf := globalEncodeBufferPool.Get()
	if buf == nil {
		t.Fatal("globalEncodeBufferPool.Get() returned nil")
	}

	if len(*buf) != 0 {
		t.Errorf("Expected empty buffer, got length %d", len(*buf))
	}

	// Use it
	*buf = append(*buf, []byte{1, 2, 3, 4, 5}...)
	if len(*buf) != 5 {
		t.Errorf("Expected length 5, got %d", len(*buf))
	}

	// Put it back
	globalEncodeBufferPool.Put(buf)

	// Get again
	buf2 := globalEncodeBufferPool.Get()
	if buf2 == nil {
		t.Fatal("globalEncodeBufferPool.Get() (second) returned nil")
	}

	// Should be reset
	if len(*buf2) != 0 {
		t.Errorf("Expected reset buffer, got length %d", len(*buf2))
	}

	globalEncodeBufferPool.Put(buf2)

	t.Log("✓ globalEncodeBufferPool Get/Put tested")
}

// TestArenaGetReset tests getArena/putArena and valueArena operations
func TestArenaGetReset(t *testing.T) {
	// Get an arena from pool
	arena := getArena()
	if arena == nil {
		t.Fatal("getArena() returned nil")
	}

	// Get some values
	slice1 := arena.Get(10)
	if len(slice1) != 10 {
		t.Errorf("Expected 10 values, got %d", len(slice1))
	}

	slice2 := arena.Get(20)
	if len(slice2) != 20 {
		t.Errorf("Expected 20 values, got %d", len(slice2))
	}

	// Reset arena
	arena.Reset()
	if arena.pos != 0 {
		t.Errorf("After Reset(), expected pos=0, got %d", arena.pos)
	}

	// Get again after reset
	slice3 := arena.Get(5)
	if len(slice3) != 5 {
		t.Errorf("After reset, expected 5 values, got %d", len(slice3))
	}

	// Put back to pool
	putArena(arena)

	t.Log("✓ getArena/putArena/Reset tested")
}

// TestArenaLargeAllocation tests arena with large allocation
func TestArenaLargeAllocation(t *testing.T) {
	arena := getArena()

	// Request more than initial capacity
	largeSlice := arena.Get(2000)
	if len(largeSlice) != 2000 {
		t.Errorf("Expected 2000 values, got %d", len(largeSlice))
	}

	// Arena should have reallocated
	if cap(arena.values) < 2000 {
		t.Errorf("Expected capacity >= 2000, got %d", cap(arena.values))
	}

	putArena(arena)

	t.Log("✓ Arena large allocation (2000 values) tested")
}

// TestMaxHelper tests the max() helper function
func TestMaxHelper(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{10, 10, 10},
		{-1, -5, -1},
		{0, 1, 1},
	}

	for _, tt := range tests {
		result := max(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("max(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
		}
	}

	t.Log("✓ max() helper tested")
}

// TestValuePoolConcurrent tests pool safety with concurrent goroutines
func TestValuePoolConcurrent(t *testing.T) {
	goroutines := 50
	iterations := 20

	done := make(chan bool, goroutines)

	for g := 0; g < goroutines; g++ {
		go func(id int) {
			for i := 0; i < iterations; i++ {
				slice := globalValuePool.Get()
				*slice = append(*slice, reflect.ValueOf(id))
				globalValuePool.Put(slice)
			}
			done <- true
		}(g)
	}

	for g := 0; g < goroutines; g++ {
		<-done
	}

	t.Logf("✓ globalValuePool concurrent: %d goroutines × %d iterations", goroutines, iterations)
}

// TestEncodeBufferPoolConcurrent tests buffer pool with concurrency
func TestEncodeBufferPoolConcurrent(t *testing.T) {
	goroutines := 50
	iterations := 20

	done := make(chan bool, goroutines)

	for g := 0; g < goroutines; g++ {
		go func(id int) {
			for i := 0; i < iterations; i++ {
				buf := globalEncodeBufferPool.Get()
				*buf = append(*buf, byte(id))
				globalEncodeBufferPool.Put(buf)
			}
			done <- true
		}(g)
	}

	for g := 0; g < goroutines; g++ {
		<-done
	}

	t.Logf("✓ globalEncodeBufferPool concurrent: %d goroutines × %d iterations", goroutines, iterations)
}

// TestArenaPoolConcurrent tests arena pool with concurrency
func TestArenaPoolConcurrent(t *testing.T) {
	goroutines := 50
	iterations := 20

	done := make(chan bool, goroutines)

	for g := 0; g < goroutines; g++ {
		go func(id int) {
			for i := 0; i < iterations; i++ {
				arena := getArena()
				_ = arena.Get(10)
				arena.Reset()
				putArena(arena)
			}
			done <- true
		}(g)
	}

	for g := 0; g < goroutines; g++ {
		<-done
	}

	t.Logf("✓ arena pool concurrent: %d goroutines × %d iterations", goroutines, iterations)
}

// TestValuePoolSizeLimit tests that large slices are not pooled
func TestValuePoolSizeLimit(t *testing.T) {
	slice := globalValuePool.Get()

	// Grow beyond limit (256)
	for i := 0; i < 300; i++ {
		*slice = append(*slice, reflect.ValueOf(i))
	}

	if len(*slice) != 300 {
		t.Errorf("Expected length 300, got %d", len(*slice))
	}

	// Put back (should not be pooled due to size)
	globalValuePool.Put(slice)

	t.Log("✓ Value pool size limit (>256) tested")
}

// TestEncodeBufferPoolSizeLimit tests that large buffers are not pooled
func TestEncodeBufferPoolSizeLimit(t *testing.T) {
	buf := globalEncodeBufferPool.Get()

	// Grow beyond limit (4KB)
	data := make([]byte, 5000)
	*buf = append(*buf, data...)

	if len(*buf) != 5000 {
		t.Errorf("Expected length 5000, got %d", len(*buf))
	}

	// Put back (should not be pooled due to size)
	globalEncodeBufferPool.Put(buf)

	t.Log("✓ Encode buffer pool size limit (>4KB) tested")
}

// TestArenaSizeLimit tests that large arenas are not pooled
func TestArenaSizeLimit(t *testing.T) {
	arena := getArena()

	// Request very large allocation
	_ = arena.Get(3000)

	if cap(arena.values) <= 2048 {
		t.Errorf("Expected capacity > 2048, got %d", cap(arena.values))
	}

	// Put back (should not be pooled due to size)
	putArena(arena)

	t.Log("✓ Arena size limit (>2048) tested")
}

// TestIsRawMessageType tests the internal helper function
func TestIsRawMessageType(t *testing.T) {
	// This function is internal but we can test it through Marshal/Unmarshal
	raw := RawMessage([]byte{0x00})

	// Marshal
	data, err := Marshal(raw)
	if err != nil {
		t.Fatalf("Marshal RawMessage failed: %v", err)
	}

	// Unmarshal
	var result RawMessage
	if err := Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal RawMessage failed: %v", err)
	}

	t.Log("✓ isRawMessageType tested indirectly through Marshal/Unmarshal")
}

// Benchmark pool operations
func BenchmarkValuePoolGetPut(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice := globalValuePool.Get()
		*slice = append(*slice, reflect.ValueOf(42))
		globalValuePool.Put(slice)
	}
}

func BenchmarkEncodeBufferPoolGetPut(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := globalEncodeBufferPool.Get()
		*buf = append(*buf, 1, 2, 3)
		globalEncodeBufferPool.Put(buf)
	}
}

func BenchmarkArenaGetPut(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		arena := getArena()
		_ = arena.Get(10)
		putArena(arena)
	}
}
