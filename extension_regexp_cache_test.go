package beve

import (
	"regexp"
	"sync"
	"testing"
)

func TestRegexpCache(t *testing.T) {
	cache := &regexpCache{
		cache: make(map[string]*regexp.Regexp, 4),
		size:  4,
	}

	// Test cache miss (first compile)
	pattern := "^test$"
	r1, err := cache.get(pattern)
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}
	if r1 == nil {
		t.Fatal("Expected compiled regexp, got nil")
	}

	// Test cache hit (should return same instance)
	r2, err := cache.get(pattern)
	if err != nil {
		t.Fatalf("Failed to get cached pattern: %v", err)
	}
	if r1 != r2 {
		t.Error("Expected same regexp instance from cache")
	}

	// Test invalid pattern
	_, err = cache.get("[invalid")
	if err == nil {
		t.Error("Expected error for invalid pattern")
	}

	// Test cache eviction
	patterns := []string{
		"pattern1", "pattern2", "pattern3", "pattern4", "pattern5",
	}
	for _, p := range patterns {
		_, err := cache.get(p)
		if err != nil {
			t.Fatalf("Failed to compile %s: %v", p, err)
		}
	}

	// After eviction, cache should be smaller
	cache.mu.RLock()
	cacheSize := len(cache.cache)
	cache.mu.RUnlock()

	if cacheSize > cache.size {
		t.Errorf("Cache size %d exceeds limit %d", cacheSize, cache.size)
	}
}

func TestRegexpCacheConcurrency(t *testing.T) {
	cache := &regexpCache{
		cache: make(map[string]*regexp.Regexp, 64),
		size:  64,
	}

	var wg sync.WaitGroup
	patterns := []string{"^test$", "[0-9]+", "[a-z]+", "\\d{3}"}

	// Concurrent reads and writes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, pattern := range patterns {
				_, err := cache.get(pattern)
				if err != nil {
					t.Errorf("Failed to compile %s: %v", pattern, err)
				}
			}
		}()
	}

	wg.Wait()
}

func BenchmarkRegexpCacheHit(b *testing.B) {
	cache := &regexpCache{
		cache: make(map[string]*regexp.Regexp, 64),
		size:  64,
	}

	pattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

	// Prime cache
	_, _ = cache.get(pattern)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cache.get(pattern)
	}
}

func BenchmarkRegexpCacheMiss(b *testing.B) {
	cache := &regexpCache{
		cache: make(map[string]*regexp.Regexp, 64),
		size:  64,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pattern := "^test" + string(rune(i%26+'a')) + "$"
		_, _ = cache.get(pattern)
	}
}

func BenchmarkRegexpCompileDirect(b *testing.B) {
	pattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = regexp.Compile(pattern)
	}
}
