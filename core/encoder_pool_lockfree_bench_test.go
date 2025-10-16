//go:build go1.21
// +build go1.21

package core

import (
	"reflect"
	"sync"
	"testing"
)

// testStruct is a realistic test structure
type testStruct struct {
	ID      int32
	Name    string
	Age     uint8
	Score   float64
	Active  bool
	Tags    []string
	Metrics map[string]int
}

// BenchmarkRealWorldPoolComparison compares pools with actual encoding workload
func BenchmarkRealWorldPoolComparison(b *testing.B) {
	testData := testStruct{
		ID:      12345,
		Name:    "John Doe",
		Age:     30,
		Score:   98.5,
		Active:  true,
		Tags:    []string{"go", "rust", "python"},
		Metrics: map[string]int{"views": 1000, "likes": 250},
	}
	
	// Pre-compute reflect.Value to avoid reflection overhead in benchmark
	v := reflect.ValueOf(testData)
	
	benchmarks := []struct {
		name          string
		useLockFree   bool
		numGoroutines int
	}{
		{"SyncPool_1G_Encode", false, 1},
		{"LockFree_1G_Encode", true, 1},
		{"SyncPool_10G_Encode", false, 10},
		{"LockFree_10G_Encode", true, 10},
		{"SyncPool_100G_Encode", false, 100},
		{"LockFree_100G_Encode", true, 100},
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
					
					// Actual encoding work
					enc.Buf.Reset()
					_ = enc.encodeStructFast(v)
					
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
							
							// Actual encoding work
							enc.Buf.Reset()
							_ = enc.encodeStructFast(v)
							
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
				var hitRate float64
				if total > 0 {
					hitRate = float64(hits) / float64(total) * 100
				}
				
				b.ReportMetric(hitRate, "%_hit_rate")
				b.ReportMetric(float64(puts), "puts")
				b.ReportMetric(float64(discards), "discards")
				b.ReportMetric(float64(overflows), "overflows")
			}
		})
	}
}

// BenchmarkPoolContentionScaling tests scaling under increasing contention
func BenchmarkPoolContentionScaling(b *testing.B) {
	goroutineCounts := []int{1, 2, 4, 8, 12, 24, 48, 96}
	
	for _, numG := range goroutineCounts {
		b.Run("SyncPool_"+string(rune('0'+numG))+"G", func(b *testing.B) {
			// Save and restore
			originalUseLockFreePool := UseLockFreePool
			defer func() { UseLockFreePool = originalUseLockFreePool }()
			UseLockFreePool = false
			
			b.ReportAllocs()
			b.ResetTimer()
			
			opsPerG := b.N / numG
			if opsPerG == 0 {
				opsPerG = 1
			}
			
			var wg sync.WaitGroup
			wg.Add(numG)
			
			for g := 0; g < numG; g++ {
				go func() {
					defer wg.Done()
					for i := 0; i < opsPerG; i++ {
						enc := GetEncoderFromPool()
						enc.Buf.WriteByte(0xFF)
						PutEncoderToPool(enc)
					}
				}()
			}
			
			wg.Wait()
		})
		
		b.Run("LockFree_"+string(rune('0'+numG))+"G", func(b *testing.B) {
			// Save and restore
			originalUseLockFreePool := UseLockFreePool
			defer func() { UseLockFreePool = originalUseLockFreePool }()
			UseLockFreePool = true
			ResetLockFreePoolStats()
			
			b.ReportAllocs()
			b.ResetTimer()
			
			opsPerG := b.N / numG
			if opsPerG == 0 {
				opsPerG = 1
			}
			
			var wg sync.WaitGroup
			wg.Add(numG)
			
			for g := 0; g < numG; g++ {
				go func() {
					defer wg.Done()
					for i := 0; i < opsPerG; i++ {
						enc := GetEncoderFromPool()
						enc.Buf.WriteByte(0xFF)
						PutEncoderToPool(enc)
					}
				}()
			}
			
			wg.Wait()
			
			b.StopTimer()
			hits, misses, _, _, _ := GetLockFreePoolStats()
			total := hits + misses
			if total > 0 {
				hitRate := float64(hits) / float64(total) * 100
				b.ReportMetric(hitRate, "%_hit_rate")
			}
		})
	}
}
