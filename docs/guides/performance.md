# ⚡ Performance Optimization Guide

Master BEVE-Go's performance features for maximum throughput and minimal memory usage.

**Reading Time**: 25 minutes  
**Level**: Advanced  
**Prerequisites**: [Encoding/Decoding](encoding-decoding.md), [Streaming](streaming.md)

---

## Table of Contents

1. [Performance Overview](#performance-overview)
2. [Zero-Copy Mode](#zero-copy-mode)
3. [Buffer Pooling](#buffer-pooling)
4. [Arena Allocator](#arena-allocator)
5. [Fast Paths](#fast-paths)
6. [Memory Optimization](#memory-optimization)
7. [CPU Optimization](#cpu-optimization)
8. [Benchmarking](#benchmarking)
9. [Production Tuning](#production-tuning)

---

## Performance Overview

### BEVE vs Competitors

**Neoverse-N2 ARM64 Results**:

| Operation | BEVE | JSON | CBOR | MessagePack | Speedup |
|-----------|------|------|------|-------------|---------|
| **Small marshal** | 694ns | 4.78μs | 2.40μs | 2.29μs | **6.9× faster** |
| **Small unmarshal** | 805ns | 8.07μs | 7.93μs | 5.69μs | **10× faster** |
| **Large marshal** | 103μs | 380μs | 170μs | 274μs | **3.7× faster** |
| **Large unmarshal** | 230μs | 2.10ms | 637μs | 527μs | **9.1× faster** |

### Performance Goals

| Metric | Target | BEVE Actual | Status |
|--------|--------|-------------|--------|
| **Small struct marshal** | <1μs | 388-694ns | ✅ 2-3× better |
| **Large payload unmarshal** | <500μs | 230-267μs | ✅ 2× better |
| **Allocations** | <10 | 1-4 | ✅ 10× better |
| **Memory overhead** | <10% | ~5% | ✅ 2× better |
| **Zero-copy marshal** | <500ns | 388ns | ✅ Achieved |

---

## Zero-Copy Mode

### What is Zero-Copy?

**Standard Marshal** (2 allocations):
```go
data, _ := beve.Marshal(user)
// Allocation 1: Encoder buffer (64-4096 bytes)
// Allocation 2: Result slice (copy of buffer)
```

**Zero-Copy Marshal** (0 allocations):
```go
buf := make([]byte, 0, 1024) // Pre-allocated (reusable)
data, _ := beve.MarshalZeroCopy(user, buf)
// Allocation 0: Uses provided buffer
// Result: Slice of provided buffer (no copy)
```

### Performance Impact

**Benchmark Results** (Small struct):

```
BenchmarkMarshal-8           1,389 ns/op    1,344 B/op    1 allocs/op
BenchmarkMarshalZeroCopy-8     388 ns/op        0 B/op    0 allocs/op
```

**Improvement**: **3.6× faster, 0 allocations**

### Basic Usage

```go
// Create reusable buffer
buf := make([]byte, 0, 4096)

for _, user := range users {
    // Encode with zero-copy
    data, err := beve.MarshalZeroCopy(user, buf[:0])
    if err != nil {
        return err
    }
    
    // Use data immediately
    conn.Write(data)
    
    // Buffer reused in next iteration
}
```

### Buffer Sizing

**Rule of thumb**: Pre-allocate 2-4× expected size

```go
// Small structs (~50 bytes)
buf := make([]byte, 0, 256) // 256 bytes

// Medium structs (~500 bytes)
buf := make([]byte, 0, 2048) // 2KB

// Large structs (~5KB)
buf := make([]byte, 0, 16384) // 16KB

// Arrays/slices: estimate N × avg_size × 2
bufSize := len(users) * 100 * 2
buf := make([]byte, 0, bufSize)
```

### Estimation Function

```go
func estimateSize(v interface{}) int {
    // Conservative estimate
    t := reflect.TypeOf(v)
    
    switch t.Kind() {
    case reflect.Struct:
        return t.NumField() * 50 // ~50 bytes per field
    case reflect.Slice:
        rv := reflect.ValueOf(v)
        if rv.Len() == 0 {
            return 64
        }
        elemSize := estimateSize(rv.Index(0).Interface())
        return rv.Len() * elemSize
    case reflect.String:
        return len(v.(string)) + 8 // String + varint
    default:
        return 16 // Primitive types
    }
}

// Usage
size := estimateSize(user) * 2 // 2× safety margin
buf := make([]byte, 0, size)
```

### Buffer Reuse Patterns

#### Pattern 1: Loop Reuse

```go
buf := make([]byte, 0, 4096)

for _, item := range items {
    data, _ := beve.MarshalZeroCopy(item, buf[:0])
    process(data)
    // buf automatically reused
}
```

#### Pattern 2: Worker Pool

```go
type Worker struct {
    buf []byte
}

func NewWorker() *Worker {
    return &Worker{
        buf: make([]byte, 0, 8192),
    }
}

func (w *Worker) Process(item interface{}) {
    data, _ := beve.MarshalZeroCopy(item, w.buf[:0])
    send(data)
}

// Each worker has own buffer (no contention)
workers := make([]*Worker, runtime.NumCPU())
for i := range workers {
    workers[i] = NewWorker()
}
```

#### Pattern 3: sync.Pool Integration

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 4096)
    },
}

func encodeWithPool(v interface{}) ([]byte, error) {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf[:0])
    
    return beve.MarshalZeroCopy(v, buf[:0])
}
```

### Pitfalls and Solutions

#### Pitfall 1: Buffer Too Small

```go
buf := make([]byte, 0, 10) // Too small!
data, err := beve.MarshalZeroCopy(largeStruct, buf)
// Error: buffer too small
```

**Solution**: Grow buffer dynamically

```go
buf := make([]byte, 0, 256)
data, err := beve.MarshalZeroCopy(v, buf)
if err == beve.ErrBufferTooSmall {
    // Grow and retry
    buf = make([]byte, 0, len(buf)*2)
    data, err = beve.MarshalZeroCopy(v, buf)
}
```

#### Pitfall 2: Data Lifetime

```go
// ❌ Wrong: Data invalidated after buffer reuse
buf := make([]byte, 0, 1024)

for _, item := range items {
    data, _ := beve.MarshalZeroCopy(item, buf[:0])
    results = append(results, data) // BUG: data points to buf!
}
// All results point to same buffer (last item)!
```

**Solution**: Copy if data needed later

```go
for _, item := range items {
    data, _ := beve.MarshalZeroCopy(item, buf[:0])
    
    // Copy data for storage
    copied := make([]byte, len(data))
    copy(copied, data)
    results = append(results, copied)
}
```

---

## Buffer Pooling

### Encoder Pooling

**Standard Approach** (1000 allocations):
```go
for i := 0; i < 1000; i++ {
    data, _ := beve.Marshal(user) // New encoder each time
}
```

**Pooled Approach** (1 allocation):
```go
enc := beve.GetEncoderFromPool()
defer beve.PutEncoderToPool(enc)

for i := 0; i < 1000; i++ {
    data, _ := enc.Marshal(user)
    enc.Reset()
}
```

### Performance Impact

**Benchmark Results**:

```
BenchmarkNewEncoder-8        200 ns/op      encoder creation
BenchmarkPooledEncoder-8       8 ns/op      pool retrieval

Speedup: 25× faster (200ns → 8ns)
```

### Encoder Pool API

```go
// Get encoder from pool
enc := beve.GetEncoderFromPool()

// Use encoder
data, err := enc.Marshal(value)

// Reset for reuse
enc.Reset()

// Return to pool
beve.PutEncoderToPool(enc)
```

### Decoder Pooling

```go
// Get decoder from pool
dec := beve.GetDecoderFromPool(data)

// Decode
var user User
err := dec.Unmarshal(&user)

// Return to pool
beve.PutDecoderToPool(dec)
```

### Custom Buffer Pool

```go
type BufferPool struct {
    pool sync.Pool
}

func NewBufferPool(size int) *BufferPool {
    return &BufferPool{
        pool: sync.Pool{
            New: func() interface{} {
                return make([]byte, 0, size)
            },
        },
    }
}

func (p *BufferPool) Get() []byte {
    return p.pool.Get().([]byte)
}

func (p *BufferPool) Put(buf []byte) {
    p.pool.Put(buf[:0])
}

// Usage
var bufPool = NewBufferPool(4096)

func encode(v interface{}) []byte {
    buf := bufPool.Get()
    defer bufPool.Put(buf)
    
    data, _ := beve.MarshalZeroCopy(v, buf[:0])
    
    // Copy before returning
    result := make([]byte, len(data))
    copy(result, data)
    return result
}
```

---

## Arena Allocator

### What is Arena Allocation?

**Standard Allocation** (GC pressure):
```go
for i := 0; i < 1000; i++ {
    data, _ := beve.Marshal(user)
    // Each marshal allocates new memory
    // GC must track and free 1000 allocations
}
// GC overhead: High
```

**Arena Allocation** (bulk free):
```go
arena := beve.NewArenaPool(1024 * 1024) // 1MB arena

for i := 0; i < 1000; i++ {
    enc := beve.GetEncoderFromPoolWithArena(arena)
    data, _ := enc.Marshal(user)
    beve.PutEncoderToPool(enc)
}

arena.Reset() // Free all at once
// GC overhead: Minimal (1 allocation instead of 1000)
```

### Performance Impact

**Benchmark Results**:

```
BenchmarkStandard-8         1,389 ns/op    1,344 B/op    1 allocs/op
BenchmarkArena-8              599 ns/op      270 B/op    1 allocs/op
BenchmarkArenaReuse-8         270 ns/op        0 B/op    0 allocs/op

Arena: 2.3× faster
Arena (reuse): 5.1× faster
```

### Basic Usage

```go
// Create arena (1MB capacity)
arena := beve.NewArenaPool(1024 * 1024)

// Get encoder with arena
enc := beve.GetEncoderFromPoolWithArena(arena)

// Encode multiple values
for _, user := range users {
    data, _ := enc.Marshal(user)
    send(data)
    enc.Reset()
}

// Return encoder
beve.PutEncoderToPool(enc)

// Reset arena (frees all allocations)
arena.Reset()
```

### Decoder Arena

```go
arena := beve.NewArenaPool(1024 * 1024)

for _, data := range dataList {
    dec := beve.NewDecoderWithArena(data, arena)
    
    var user User
    dec.Unmarshal(&user)
    
    process(user)
}

// Free all arena allocations
arena.Reset()
```

### Arena Sizing

**Guidelines**:

```go
// Small batches (<100 items, ~50KB total)
arena := beve.NewArenaPool(128 * 1024) // 128KB

// Medium batches (100-1000 items, ~500KB total)
arena := beve.NewArenaPool(1024 * 1024) // 1MB

// Large batches (>1000 items, >5MB total)
arena := beve.NewArenaPool(10 * 1024 * 1024) // 10MB

// Estimate: batch_size × avg_item_size × 1.5
arenaSize := len(batch) * 500 * 1.5
```

### Arena Best Practices

#### 1. Reset After Batch

```go
arena := beve.NewArenaPool(1024 * 1024)

for _, batch := range batches {
    for _, item := range batch {
        enc := beve.GetEncoderFromPoolWithArena(arena)
        data, _ := enc.Marshal(item)
        send(data)
        beve.PutEncoderToPool(enc)
    }
    
    // Reset after each batch
    arena.Reset()
}
```

#### 2. Per-Worker Arenas

```go
type Worker struct {
    arena *beve.Arena
}

func NewWorker() *Worker {
    return &Worker{
        arena: beve.NewArenaPool(512 * 1024),
    }
}

func (w *Worker) Process(items []Item) {
    for _, item := range items {
        enc := beve.GetEncoderFromPoolWithArena(w.arena)
        data, _ := enc.Marshal(item)
        send(data)
        beve.PutEncoderToPool(enc)
    }
    
    w.arena.Reset()
}
```

#### 3. Arena Pool

```go
var arenaPool = sync.Pool{
    New: func() interface{} {
        return beve.NewArenaPool(1024 * 1024)
    },
}

func processWithArena(items []Item) {
    arena := arenaPool.Get().(*beve.Arena)
    defer func() {
        arena.Reset()
        arenaPool.Put(arena)
    }()
    
    for _, item := range items {
        enc := beve.GetEncoderFromPoolWithArena(arena)
        data, _ := enc.Marshal(item)
        process(data)
        beve.PutEncoderToPool(enc)
    }
}
```

---

## Fast Paths

### Primitive Fast Paths

BEVE detects and optimizes primitive types:

```go
// Fast path: Direct encoding (no reflection)
beve.Marshal(42)              // int → 2-9 bytes
beve.Marshal(3.14)            // float64 → 9 bytes
beve.Marshal("hello")         // string → 6 bytes
beve.Marshal(true)            // bool → 1 byte

// Slow path: Reflection
beve.Marshal(User{})          // struct → reflect fields
beve.Marshal(interface{}(42)) // interface → reflect type
```

### Typed Array Fast Paths

**Standard encoding** (slow):
```go
users := []User{{Name: "Alice"}, {Name: "Bob"}}
beve.Marshal(users)
// For each user:
//   Write field names (repeated!)
//   Write field values
// Total: ~200 bytes
```

**Typed array encoding** (fast):
```go
beve.MarshalTyped(users)
// Write field names once
// For each user:
//   Write field values only
// Total: ~120 bytes (40% smaller)
// Speed: 2-3× faster encoding
```

### Slice Fast Paths

**Optimized types** (memcpy):
```go
[]int32   → Header + memcpy (blazingly fast)
[]float64 → Header + memcpy
[]uint8   → Header + memcpy (bytes)

// Benchmark: 1M elements
[]int32:   50 μs (200 MB/s throughput)
```

**Non-optimized** (loop):
```go
[]User → Header + for each user (reflect)

// Benchmark: 1M elements  
[]User:   50 ms (1000× slower than primitives)
```

### String Interning (Future)

**Current**:
```go
users := []User{
    {Name: "Alice", City: "NYC"},
    {Name: "Alice", City: "NYC"}, // Duplicate strings
}
// "Alice" encoded twice (12 bytes)
// "NYC" encoded twice (8 bytes)
```

**With interning** (v1.4+):
```go
// "Alice" encoded once + reference (4 bytes)
// "NYC" encoded once + reference (4 bytes)
// Savings: ~50% for duplicate strings
```

---

## Memory Optimization

### 1. Pre-Allocate Slices

**❌ Bad: Dynamic growth**:
```go
var users []User
beve.Unmarshal(data, &users)
// Slice grows: 0→1→2→4→8→16→32→64→100
// Multiple allocations + copies
```

**✅ Good: Pre-allocate**:
```go
users := make([]User, 0, 100)
beve.Unmarshal(data, &users)
// Single allocation: 100 capacity
// Savings: 80% fewer allocations
```

### 2. Reuse Buffers

**❌ Bad: New buffer each time**:
```go
for _, user := range users {
    data, _ := beve.Marshal(user) // New allocation
    send(data)
}
// 1000 users = 1000 allocations
```

**✅ Good: Reuse buffer**:
```go
buf := make([]byte, 0, 1024)
for _, user := range users {
    data, _ := beve.MarshalZeroCopy(user, buf[:0])
    send(data)
}
// 1000 users = 1 allocation
```

### 3. Use Pointers for Large Structs

**❌ Bad: Copy entire struct**:
```go
func process(user User) { // 1KB struct copied
    data, _ := beve.Marshal(user)
}

for _, user := range users {
    process(user) // 1000× 1KB = 1MB copies
}
```

**✅ Good: Use pointer**:
```go
func process(user *User) { // 8 bytes pointer
    data, _ := beve.Marshal(user)
}

for i := range users {
    process(&users[i]) // 1000× 8 bytes = 8KB
}
```

### 4. Avoid interface{} Where Possible

**❌ Slow: Interface boxing**:
```go
values := []interface{}{1, 2, 3}
beve.Marshal(values)
// Each int boxed into interface (allocation + reflection)
```

**✅ Fast: Concrete type**:
```go
values := []int{1, 2, 3}
beve.Marshal(values)
// Direct encoding (no boxing, no reflection)
```

### 5. String vs []byte

**String** (immutable, safe):
```go
s := "hello"
beve.Marshal(s) // Safe, no surprises
```

**[]byte** (mutable, careful):
```go
b := []byte("hello")
data, _ := beve.Marshal(b)
// Modify b after marshal?
b[0] = 'H' // May corrupt data if zero-copy used!
```

**Rule**: If using zero-copy, don't modify source data!

---

## CPU Optimization

### 1. Little-Endian Optimization

BEVE uses little-endian (modern CPU native):

```go
// Fast: Native byte order (x86, ARM)
binary.LittleEndian.PutUint32(buf, value)
// 1 CPU instruction (MOV)

// Slow: Non-native byte order
binary.BigEndian.PutUint32(buf, value)
// Multiple instructions (BSWAP + MOV)
```

**Impact**: 10-20% faster on x86/ARM

### 2. Avoid Reflection in Hot Paths

**❌ Slow: Reflection every time**:
```go
for _, user := range users {
    t := reflect.TypeOf(user) // Expensive!
    // ... reflect operations
}
```

**✅ Fast: Cache type info**:
```go
// BEVE does this internally
var typeCache sync.Map

func marshalCached(v interface{}) {
    t := reflect.TypeOf(v)
    
    // Check cache
    if info, ok := typeCache.Load(t); ok {
        // Use cached info (fast path)
        return marshalWithInfo(v, info)
    }
    
    // Reflect once, cache result
    info := reflectType(t)
    typeCache.Store(t, info)
    return marshalWithInfo(v, info)
}
```

### 3. Inline Small Functions

**Compiler inlining** (automatic for small functions):
```go
//go:inline
func putVarint(buf []byte, v uint64) int {
    // Small function: Compiler inlines
    // No function call overhead
}
```

### 4. SIMD Opportunities (Future)

**Current**: Scalar operations
**Future (v1.5+)**: SIMD for typed arrays

```go
// Copy 1000 int32 values
// Current: ~200ns (scalar)
// SIMD (AVX-512): ~50ns (4× faster)
```

---

## Benchmarking

### Basic Benchmark

```go
func BenchmarkMarshal(b *testing.B) {
    user := User{Name: "Alice", Age: 30}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        beve.Marshal(user)
    }
}
```

**Run**:
```bash
go test -bench=BenchmarkMarshal -benchmem
```

### Comparative Benchmark

```go
func BenchmarkJSON(b *testing.B) {
    user := User{Name: "Alice", Age: 30}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        json.Marshal(user)
    }
}

func BenchmarkBEVE(b *testing.B) {
    user := User{Name: "Alice", Age: 30}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        beve.Marshal(user)
    }
}

func BenchmarkBEVEZeroCopy(b *testing.B) {
    user := User{Name: "Alice", Age: 30}
    buf := make([]byte, 0, 256)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        beve.MarshalZeroCopy(user, buf[:0])
    }
}
```

### Memory Profiling

```bash
# Run with memory profiling
go test -bench=. -benchmem -memprofile=mem.prof

# Analyze profile
go tool pprof mem.prof
```

### CPU Profiling

```bash
# Run with CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# Analyze profile
go tool pprof cpu.prof
```

### Benchmark Matrix

```go
func BenchmarkMatrix(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            users := generateUsers(size)
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                beve.Marshal(users)
            }
        })
    }
}
```

---

## Production Tuning

### 1. Profile in Production

**Enable profiling endpoint**:
```go
import _ "net/http/pprof"

func main() {
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Your app
    startServer()
}
```

**Collect profiles**:
```bash
# CPU profile (30 seconds)
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof

# Heap profile
curl http://localhost:6060/debug/pprof/heap > heap.prof

# Analyze
go tool pprof cpu.prof
```

### 2. Monitor GC Pressure

```go
import "runtime"

func monitorGC() {
    var m runtime.MemStats
    
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        runtime.ReadMemStats(&m)
        
        fmt.Printf("Alloc: %d MB\n", m.Alloc/1024/1024)
        fmt.Printf("NumGC: %d\n", m.NumGC)
        fmt.Printf("PauseNs: %d ms\n", m.PauseNs[(m.NumGC+255)%256]/1e6)
    }
}
```

### 3. Tune GOMAXPROCS

```go
import "runtime"

func init() {
    // Set to number of physical cores
    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)
    
    fmt.Printf("GOMAXPROCS: %d\n", numCPU)
}
```

### 4. Worker Pool Pattern

```go
type WorkerPool struct {
    workers   []*Worker
    jobs      chan Job
    results   chan Result
}

func NewWorkerPool(size int) *WorkerPool {
    pool := &WorkerPool{
        workers: make([]*Worker, size),
        jobs:    make(chan Job, size*10),
        results: make(chan Result, size*10),
    }
    
    for i := 0; i < size; i++ {
        pool.workers[i] = NewWorker(pool.jobs, pool.results)
        go pool.workers[i].Start()
    }
    
    return pool
}

type Worker struct {
    buf   []byte
    arena *beve.Arena
}

func (w *Worker) Start() {
    for job := range jobs {
        enc := beve.GetEncoderFromPoolWithArena(w.arena)
        data, _ := enc.Marshal(job.Data)
        results <- Result{Data: data}
        beve.PutEncoderToPool(enc)
    }
}
```

### 5. Batch Processing

```go
func processBatch(items []Item, batchSize int) {
    arena := beve.NewArenaPool(1024 * 1024)
    buf := make([]byte, 0, 8192)
    
    for i := 0; i < len(items); i += batchSize {
        end := i + batchSize
        if end > len(items) {
            end = len(items)
        }
        
        batch := items[i:end]
        
        for _, item := range batch {
            data, _ := beve.MarshalZeroCopy(item, buf[:0])
            send(data)
        }
        
        // Reset arena after batch
        arena.Reset()
    }
}
```

---

## Summary

### Performance Checklist

**Encoding**:
- ✅ Use `MarshalZeroCopy` for hot paths (2-8× faster)
- ✅ Pool encoders with `GetEncoderFromPool()` (25× faster)
- ✅ Use arena allocator for batches (2-5× faster)
- ✅ Pre-allocate buffers (estimate × 2)
- ✅ Use typed arrays for struct slices (35-48% smaller)

**Decoding**:
- ✅ Pre-allocate slices with capacity
- ✅ Pool decoders
- ✅ Use arena for temporary data
- ✅ Avoid interface{} where possible

**Memory**:
- ✅ Reuse buffers across iterations
- ✅ Use pointers for large structs
- ✅ Batch reset arenas
- ✅ Monitor GC pressure

**CPU**:
- ✅ Cache type information (automatic)
- ✅ Use fast paths (primitives, typed arrays)
- ✅ Profile production workloads
- ✅ Tune GOMAXPROCS

### Performance Gains Summary

| Technique | Speedup | Use Case |
|-----------|---------|----------|
| **Zero-Copy** | 2-8× | High-throughput systems |
| **Encoder Pooling** | 25× | Hot paths (8ns vs 200ns) |
| **Arena Allocator** | 2-5× | Batch processing |
| **Typed Arrays** | 2-3× | Struct arrays (N≥5) |
| **Buffer Reuse** | 10-100× | Loop encoding |
| **Pre-Allocation** | 2-5× | Large slices |

### Next Steps

- **[Arena Allocator Guide →](arena-allocator.md)** - Deep dive into arenas
- **[Extensions →](extensions.md)** - Advanced features
- **[Production Deployment →](../production/deployment.md)** - Deploy optimized
- **[Benchmarks →](../../benchmarks/MULTI_PLATFORM.md)** - Full results

---

**Ready to optimize further?** Check the [Arena Allocator Guide](arena-allocator.md) for advanced memory management.
