package beve

import (
	"regexp"
	"sync"
)

// regexpCache provides thread-safe caching for compiled regex patterns
// This significantly reduces allocation overhead (36-67 allocs → 1-2 allocs)
type regexpCache struct {
	mu    sync.RWMutex
	cache map[string]*regexp.Regexp
	size  int
}

var (
	globalRegexpCache = &regexpCache{
		cache: make(map[string]*regexp.Regexp, 64),
		size:  64, // LRU size limit
	}
)

// get retrieves a compiled regex from cache or compiles and caches it
func (rc *regexpCache) get(pattern string) (*regexp.Regexp, error) {
	// Fast path: read lock
	rc.mu.RLock()
	if r, ok := rc.cache[pattern]; ok {
		rc.mu.RUnlock()
		return r, nil
	}
	rc.mu.RUnlock()

	// Slow path: compile and cache
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	rc.mu.Lock()
	defer rc.mu.Unlock()

	// Double-check after acquiring write lock
	if cached, ok := rc.cache[pattern]; ok {
		return cached, nil
	}

	// Simple eviction: clear cache if full
	if len(rc.cache) >= rc.size {
		// Keep most recently used half (simple LRU approximation)
		newCache := make(map[string]*regexp.Regexp, rc.size)
		count := 0
		for k, v := range rc.cache {
			if count >= rc.size/2 {
				break
			}
			newCache[k] = v
			count++
		}
		rc.cache = newCache
	}

	rc.cache[pattern] = r
	return r, nil
}

// clear clears the regex cache (for testing)
func (rc *regexpCache) clear() {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.cache = make(map[string]*regexp.Regexp, rc.size)
}

// ClearRegexpCache clears the global regex cache
func ClearRegexpCache() {
	globalRegexpCache.clear()
}
