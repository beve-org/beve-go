package beve

import (
	"reflect"
	"sync"
)

// Phase 1 optimization: Value pool to reduce allocations during slice/array encoding

// valuePool provides pooled reflect.Value slices
type valuePool struct {
	pool sync.Pool
}

var globalValuePool = &valuePool{
	pool: sync.Pool{
		New: func() interface{} {
			// Pre-allocate slice with reasonable capacity
			values := make([]reflect.Value, 0, 32)
			return &values
		},
	},
}

func (vp *valuePool) Get() *[]reflect.Value {
	slice := vp.pool.Get().(*[]reflect.Value)
	*slice = (*slice)[:0] // Reset length but keep capacity
	return slice
}

func (vp *valuePool) Put(slice *[]reflect.Value) {
	if cap(*slice) <= 256 { // Only pool reasonable sizes
		vp.pool.Put(slice)
	}
}

// Pre-allocated buffer pool for encoding operations
type encodeBufferPool struct {
	pool sync.Pool
}

var globalEncodeBufferPool = &encodeBufferPool{
	pool: sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 0, 128)
			return &buf
		},
	},
}

func (bp *encodeBufferPool) Get() *[]byte {
	buf := bp.pool.Get().(*[]byte)
	*buf = (*buf)[:0]
	return buf
}

func (bp *encodeBufferPool) Put(buf *[]byte) {
	if cap(*buf) <= 4096 { // Only pool up to 4KB
		bp.pool.Put(buf)
	}
}

// Arena allocator for batch value operations
type valueArena struct {
	values []reflect.Value
	pos    int
}

// Get returns a slice of reflect.Values from the arena
func (a *valueArena) Get(n int) []reflect.Value {
	if a.pos+n > len(a.values) {
		// Need more space, reallocate
		a.values = make([]reflect.Value, max(1024, n*2))
		a.pos = 0
	}

	slice := a.values[a.pos : a.pos+n]
	a.pos += n
	return slice
}

// Reset resets the arena for reuse
func (a *valueArena) Reset() {
	a.pos = 0
	// Keep the backing array for reuse
}

// Pool of arenas
var arenaPool = sync.Pool{
	New: func() interface{} {
		return &valueArena{
			values: make([]reflect.Value, 1024),
			pos:    0,
		}
	},
}

func getArena() *valueArena {
	arena := arenaPool.Get().(*valueArena)
	arena.Reset()
	return arena
}

func putArena(arena *valueArena) {
	if cap(arena.values) <= 2048 { // Only pool reasonable sizes
		arenaPool.Put(arena)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
