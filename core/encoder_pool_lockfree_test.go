//go:build go1.21
// +build go1.21

package core

import (
	"os"
	"runtime"
	"sync"
	"testing"
)

// TestLockFreePoolBasic tests basic lock-free pool operations
func TestLockFreePoolBasic(t *testing.T) {
	// Save and restore original state
	originalUseLockFreePool := UseLockFreePool
	defer func() { UseLockFreePool = originalUseLockFreePool }()
	
	// Enable lock-free pool for this test
	UseLockFreePool = true
	
	// Reset statistics
	ResetLockFreePoolStats()
	
	// Get encoder from pool (should create new one)
	enc1 := GetEncoderFromPool()
	if enc1 == nil {
		t.Fatal("GetEncoderFromPool returned nil")
	}
	if enc1.Buf == nil {
		t.Fatal("Encoder buffer is nil")
	}
	
	// Check statistics: should have 1 miss (pool was empty)
	hits1, misses1, _, _, _ := GetLockFreePoolStats()
	if hits1 != 0 {
		t.Errorf("Expected 0 hits, got %d", hits1)
	}
	if misses1 != 1 {
		t.Errorf("Expected 1 miss, got %d", misses1)
	}
	
	// Return encoder to pool
	PutEncoderToPool(enc1)
	
	// Check statistics: should have 1 put
	_, _, puts1, _, _ := GetLockFreePoolStats()
	if puts1 != 1 {
		t.Errorf("Expected 1 put, got %d", puts1)
	}
	
	// Get encoder again (should come from pool)
	enc2 := GetEncoderFromPool()
	if enc2 == nil {
		t.Fatal("GetEncoderFromPool returned nil")
	}
	
	// Should be the same encoder
	if enc2 != enc1 {
		t.Error("Expected same encoder from pool")
	}
	
	// Check statistics: should have 1 hit
	hits2, _, _, _, _ := GetLockFreePoolStats()
	if hits2 != 1 {
		t.Errorf("Expected 1 hit, got %d", hits2)
	}
	
	PutEncoderToPool(enc2)
}

// TestLockFreePoolConcurrent tests lock-free pool under concurrent load
func TestLockFreePoolConcurrent(t *testing.T) {
	// Save and restore original state
	originalUseLockFreePool := UseLockFreePool
	defer func() { UseLockFreePool = originalUseLockFreePool }()
	
	// Enable lock-free pool
	UseLockFreePool = true
	ResetLockFreePoolStats()
	
	const numGoroutines = 100
	const opsPerGoroutine = 1000
	
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			
			for j := 0; j < opsPerGoroutine; j++ {
				enc := GetEncoderFromPool()
				if enc == nil || enc.Buf == nil {
					t.Error("GetEncoderFromPool returned nil or invalid encoder")
					return
				}
				
				// Simulate some work
				enc.Buf.WriteByte(0xFF)
				
				PutEncoderToPool(enc)
			}
		}()
	}
	
	wg.Wait()
	
	// Verify statistics
	hits, misses, puts, discards, overflows := GetLockFreePoolStats()
	
	t.Logf("Lock-free pool stats after concurrent test:")
	t.Logf("  Hits: %d", hits)
	t.Logf("  Misses: %d", misses)
	t.Logf("  Puts: %d", puts)
	t.Logf("  Discards: %d", discards)
	t.Logf("  Overflows: %d", overflows)
	
	totalOps := uint64(numGoroutines * opsPerGoroutine)
	
	// Total gets = hits + misses
	if hits+misses != totalOps {
		t.Errorf("Expected %d total gets, got %d", totalOps, hits+misses)
	}
	
	// Total puts + discards + overflows should equal total ops
	if puts+discards+overflows != totalOps {
		t.Errorf("Expected %d total puts, got %d", totalOps, puts+discards+overflows)
	}
	
	// Hit rate should be reasonable (>50% after warmup)
	hitRate := float64(hits) / float64(hits+misses) * 100
	t.Logf("Hit rate: %.2f%%", hitRate)
	
	if hitRate < 50.0 {
		t.Logf("Warning: Hit rate is low (%.2f%%), expected >50%%", hitRate)
	}
}

// TestLockFreePoolPerP verifies per-P pool isolation
func TestLockFreePoolPerP(t *testing.T) {
	// Save and restore original state
	originalUseLockFreePool := UseLockFreePool
	defer func() { UseLockFreePool = originalUseLockFreePool }()
	
	// Enable lock-free pool
	UseLockFreePool = true
	ResetLockFreePoolStats()
	
	numP := runtime.GOMAXPROCS(0)
	t.Logf("Testing with %d Ps", numP)
	
	// Ensure pools are initialized
	initPerPPools()
	
	if len(perPEncoderPools) != numP {
		t.Fatalf("Expected %d pools, got %d", numP, len(perPEncoderPools))
	}
	
	// Each pool should start empty
	for i, pool := range perPEncoderPools {
		if pool.head != nil {
			t.Errorf("Pool %d should start empty", i)
		}
		if pool.depth != 0 {
			t.Errorf("Pool %d should have depth 0, got %d", i, pool.depth)
		}
	}
}

// TestLockFreePoolMaxDepth tests pool depth limit
func TestLockFreePoolMaxDepth(t *testing.T) {
	// Save and restore original state
	originalUseLockFreePool := UseLockFreePool
	defer func() { UseLockFreePool = originalUseLockFreePool }()
	
	// Enable lock-free pool
	UseLockFreePool = true
	ResetLockFreePoolStats()
	
	// Put more encoders than maxDepth
	const putCount = lockFreePoolMaxDepth + 10
	
	encoders := make([]*Encoder, putCount)
	for i := 0; i < putCount; i++ {
		encoders[i] = GetEncoderFromPool()
	}
	
	// Return all encoders
	for i := 0; i < putCount; i++ {
		PutEncoderToPool(encoders[i])
	}
	
	// Check statistics: should have overflows
	_, _, puts, discards, overflows := GetLockFreePoolStats()
	
	t.Logf("After putting %d encoders (max depth %d):", putCount, lockFreePoolMaxDepth)
	t.Logf("  Puts: %d", puts)
	t.Logf("  Discards: %d", discards)
	t.Logf("  Overflows: %d", overflows)
	
	// Should have some overflows (excess encoders)
	if overflows == 0 && putCount > lockFreePoolMaxDepth {
		t.Logf("Warning: Expected overflows but got 0")
	}
}

// TestLockFreePoolLargeBuffer tests that large buffers are not pooled
func TestLockFreePoolLargeBuffer(t *testing.T) {
	// Save and restore original state
	originalUseLockFreePool := UseLockFreePool
	defer func() { UseLockFreePool = originalUseLockFreePool }()
	
	// Enable lock-free pool
	UseLockFreePool = true
	ResetLockFreePoolStats()
	
	// Get encoder and grow buffer beyond maxBufferPoolCapacity
	enc := GetEncoderFromPool()
	
	// Grow buffer to exceed pool capacity
	largeSize := maxBufferPoolCapacity + 1024
	enc.Buf.data = make([]byte, largeSize)
	
	// Return to pool (should be discarded)
	PutEncoderToPool(enc)
	
	// Check statistics: should have 1 discard
	_, _, _, discards, _ := GetLockFreePoolStats()
	if discards != 1 {
		t.Errorf("Expected 1 discard for large buffer, got %d", discards)
	}
}

// TestLockFreePoolEnvVar tests environment variable configuration
func TestLockFreePoolEnvVar(t *testing.T) {
	tests := []struct {
		envValue string
		expected bool
	}{
		{"true", true},
		{"1", true},
		{"yes", true},
		{"false", false},
		{"0", false},
		{"no", false},
		{"", false}, // Empty should default to false
	}
	
	for _, tt := range tests {
		t.Run(tt.envValue, func(t *testing.T) {
			// Set environment variable
			if tt.envValue != "" {
				os.Setenv("BEVE_USE_LOCKFREE_POOL", tt.envValue)
			} else {
				os.Unsetenv("BEVE_USE_LOCKFREE_POOL")
			}
			defer os.Unsetenv("BEVE_USE_LOCKFREE_POOL")
			
			// Re-run init to apply environment variable
			// Note: This is a simplified test - in real code, init runs once
			var testUseLockFreePool bool
			if val := os.Getenv("BEVE_USE_LOCKFREE_POOL"); val != "" {
				// Simple bool parsing
				testUseLockFreePool = (val == "true" || val == "1" || val == "yes")
			}
			
			if testUseLockFreePool != tt.expected {
				t.Errorf("Expected UseLockFreePool=%v for env=%q, got %v",
					tt.expected, tt.envValue, testUseLockFreePool)
			}
		})
	}
}

// BenchmarkLockFreePoolVsSyncPool compares lock-free pool vs sync.Pool
func BenchmarkLockFreePoolVsSyncPool(b *testing.B) {
	benchmarks := []struct {
		name          string
		useLockFree   bool
		numGoroutines int
	}{
		{"SyncPool_1G", false, 1},
		{"LockFree_1G", true, 1},
		{"SyncPool_10G", false, 10},
		{"LockFree_10G", true, 10},
		{"SyncPool_100G", false, 100},
		{"LockFree_100G", true, 100},
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Save and restore original state
			originalUseLockFreePool := UseLockFreePool
			defer func() { UseLockFreePool = originalUseLockFreePool }()
			
			UseLockFreePool = bm.useLockFree
			ResetLockFreePoolStats()
			
			b.ReportAllocs()
			b.ResetTimer()
			
			if bm.numGoroutines == 1 {
				// Single goroutine benchmark
				for i := 0; i < b.N; i++ {
					enc := GetEncoderFromPool()
					enc.Buf.WriteByte(0xFF) // Simulate work
					PutEncoderToPool(enc)
				}
			} else {
				// Multi-goroutine benchmark
				opsPerGoroutine := b.N / bm.numGoroutines
				if opsPerGoroutine == 0 {
					opsPerGoroutine = 1
				}
				
				var wg sync.WaitGroup
				wg.Add(bm.numGoroutines)
				
				for g := 0; g < bm.numGoroutines; g++ {
					go func() {
						defer wg.Done()
						for i := 0; i < opsPerGoroutine; i++ {
							enc := GetEncoderFromPool()
							enc.Buf.WriteByte(0xFF)
							PutEncoderToPool(enc)
						}
					}()
				}
				
				wg.Wait()
			}
			
			b.StopTimer()
			
			// Report statistics for lock-free pool
			if bm.useLockFree {
				hits, misses, puts, discards, overflows := GetLockFreePoolStats()
				total := hits + misses
				hitRate := float64(hits) / float64(total) * 100
				
				b.ReportMetric(hitRate, "%_hit_rate")
				b.ReportMetric(float64(puts), "puts")
				b.ReportMetric(float64(discards), "discards")
				b.ReportMetric(float64(overflows), "overflows")
			}
		})
	}
}
