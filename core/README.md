# BEVE Core - Internal Architecture Documentation

**Version:** 1.3.0  
**Last Updated:** 16 Ekim 2025  
**Status:** Production Ready

> ⚠️ **Internal Package**: This package is not part of the public API. Use the top-level `beve` package for application code.

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Module Reference](#module-reference)
4. [Performance Systems](#performance-systems)
5. [Memory Management](#memory-management)
6. [Thread Safety](#thread-safety)
7. [Optimization Phases](#optimization-phases)
8. [Development Guide](#development-guide)

---

## Overview

The `core` package contains the internal implementation of BEVE (Binary Efficient Versatile Encoding). It is organized into logical modules for maintainability, performance, and clarity.

### Design Principles

1. **Performance First**: Zero-copy where possible, minimal allocations
2. **Cache-Friendly**: Struct layout optimized for CPU cache lines
3. **Type Safety**: Full Go type system support with reflection fallback
4. **Modularity**: Clear separation of concerns across files

### Key Performance Characteristics

- **Small Structs**: 395-854ns (2-6× faster than JSON)
- **Medium Payloads**: 5.2-9.0μs (5-8× faster than JSON)
- **Large Payloads**: 46-67μs (4-6× faster than JSON)
- **Memory**: 78-99.9% reduction with ZeroCopy mode
- **Allocations**: 2-4 allocs for optimized paths, 17-20 for complex types

---

## Architecture

### High-Level Structure

```
core/
├── Encoder Pipeline
│   ├── encoder_base.go          # Core encoder, type dispatch
│   ├── encoder_primitives.go    # int, float, bool, string
│   ├── encoder_collections.go   # slice, map, struct
│   ├── encoder_write.go         # Low-level write operations
│   └── encoder_utils.go         # Helpers (varint, type checks)
│
├── Fast Paths (Phase 1)
│   ├── encoder_stack.go         # Stack-based encoding (143ns)
│   ├── encoder_cache.go         # Metadata cache (181-253ns)
│   └── encoder_fast_path.go     # Fast path routing
│
├── SIMD & Optimizations (Phase 2)
│   ├── simd.go                  # SIMD infrastructure
│   ├── simd_arm64.go            # NEON (ARM64)
│   ├── simd_amd64.go            # AVX2 (AMD64)
│   ├── simd_generic.go          # Fallback
│   ├── simd_string_*.go/.s      # String-specific SIMD
│   ├── prefetch.go              # Software prefetching (disabled)
│   └── prefetch_*.s/.go         # Platform-specific prefetch
│
├── Memory Management
│   ├── buffer.go                # Buffer pooling & growth
│   ├── buffer_platform.go       # Platform-specific buffer ops
│   └── arena.go                 # Arena allocator (batch ops)
│
├── Decoder Pipeline
│   ├── decoder_base.go          # Core decoder, type dispatch
│   ├── decoder_primitives.go    # Primitive decoding
│   ├── decoder_collections.go   # Collection decoding
│   ├── decoder_read.go          # Low-level read operations
│   ├── decoder_utils.go         # Helpers
│   └── decoder_fast_paths.go    # Decoder fast paths
│
├── Type System
│   ├── common.go                # Shared constants & types
│   └── doc.go                   # Package documentation, config
│
└── Documentation
    ├── README.md                # This file
    ├── BUILD_TAGS.md            # Build configuration
    └── PERFORMANCE_COMPARISON.md # Benchmark results
```

---

## Module Reference

### 1. Encoder Base (`encoder_base.go`)

**Purpose**: Core encoder structure and type dispatch logic.

**Key Components**:

```go
// Encoder - Main encoding structure (optimized for cache efficiency)
type Encoder struct {
    Buf           *Buffer      // Pooled buffer (hot path)
    w             io.Writer    // Target writer
    uintScratch   [9]byte      // Integer scratch buffer
    floatBuf      [9]byte      // Float scratch buffer
    varintScratch [5]byte      // Varint scratch buffer
    single        [1]byte      // Single byte writes
    batchLen      int          // Batch length
    batchBuf      [256]byte    // Batch buffer (cold path)
}

// encoderPool - sync.Pool for encoder reuse
var encoderPool = sync.Pool{...}

// GetEncoderFromPool() - Acquire encoder from pool
// PutEncoderToPool()   - Return encoder to pool
// NewEncoder(w)        - Create encoder with specific writer
```

**Memory Layout** (Cache Optimization):
- **First 64 bytes** (1 cache line): Hot path fields
  - Pointers: `Buf`, `w` (16 bytes)
  - Scratch buffers: 24 bytes
  - `batchLen`: 8 bytes
  - **Total: 48 bytes** (fits in L1 cache line)
- **Second cache line**: Cold path (`batchBuf`)

**Current Pool**: ⚠️ **sync.Pool** (mutex-based, global lock)

**Planned Upgrade** (Phase 3A):
- **Lock-Free Per-P Pools**: Eliminate mutex contention
- **Expected Improvement**: 1.3-1.4× faster (253ns → 180ns)
- **See**: [Phase 3A](#phase-3a-lock-free-pooling) below

---

### 2. Fast Path System (Phase 1)

#### 2.1 Stack Encoding (`encoder_stack.go`)

**Purpose**: Zero-allocation encoding for primitive-only structs.

**Features**:
- 128-byte stack buffer (1 cache line on M2 Max)
- No heap allocation for small structs
- Automatic fallback to heap if needed

**Performance**:
```
Small struct (4 primitive fields):
  - Time: 143ns/op
  - Memory: 112 B/op
  - Allocations: 2 allocs/op
  - Speedup: 4.2× vs baseline (600ns)
```

**Eligibility Criteria**:
- All fields must be primitives (int, float, bool, string)
- No slices, maps, nested structs, pointers
- Struct size < 80 bytes
- Coverage: ~20-30% of typical structs

**Code Example**:
```go
type SimpleStruct struct {
    ID   int32
    Name string
    Age  uint8
}

// Automatically uses stack encoding (no heap alloc!)
data, _ := beve.Marshal(SimpleStruct{1, "Alice", 30})
```

#### 2.2 Metadata Cache (`encoder_cache.go`)

**Purpose**: Pre-compute struct metadata to eliminate reflection overhead.

**Features**:
- 128-byte cache entries (aligned to cache line)
- Stores field offsets, types, tags
- Works for structs ≤12 fields
- Thread-safe with sync.Map

**Performance**:
```
Cached struct (8 fields, all primitives):
  - Time: 181ns/op
  - Speedup: 3.3× vs baseline (600ns)

Cached struct (8 primitives + 1 slice):
  - Time: 253ns/op
  - Speedup: 2.4× vs baseline (600ns)
```

**Cache Structure**:
```go
type structCache struct {
    fields       [12]fieldInfo  // Pre-computed field metadata
    fieldCount   int            // Active field count
    hasSlices    bool           // Contains slice fields
    totalSize    int            // Estimated encoding size
    _            [64]byte       // Cache-line padding
}

type fieldInfo struct {
    offset       uintptr        // Field memory offset
    typ          reflect.Type   // Field type
    tag          string         // BEVE tag name
    isPrimitive  bool           // Fast path eligible
}
```

**Coverage**: ~60-70% of structs (≤12 fields)

#### 2.3 Fast Path Routing (`encoder_fast_path.go`)

**Purpose**: Automatic optimal encoder selection.

**3-Tier System**:
```go
func (e *Encoder) EncodeAndDetach() ([]byte, error) {
    // Tier 1: Stack encoding (143ns)
    if canUseStackEncoding(v) {
        return e.encodeWithStack(v)
    }
    
    // Tier 2: Cached encoding (181-253ns)
    if cache := getCachedStructInfo(v.Type()); cache != nil {
        return e.encodeWithCache(v, cache)
    }
    
    // Tier 3: Standard reflection (600ns)
    return e.encodeReflect(v)
}
```

**Decision Tree**:
```
Input Struct
    │
    ├─ All primitives? ──Yes──> Stack Encoding (143ns)
    │
    ├─ ≤12 fields? ──Yes──> Check cache
    │                          │
    │                          ├─ Cached? ──Yes──> Cached Encoding (181-253ns)
    │                          └─ Not cached? ──> Build cache, use it
    │
    └─ Complex struct ──> Standard Reflection (600ns)
```

---

### 3. SIMD System (Phase 2)

#### 3.1 SIMD Infrastructure (`simd.go`)

**Purpose**: SIMD-accelerated array encoding with runtime CPU detection.

**Supported Platforms**:
- **ARM64**: NEON (128-bit vectors)
- **AMD64**: AVX2 (256-bit vectors)
- **Generic**: Scalar fallback

**Runtime Detection**:
```go
var (
    UseSIMD  bool  // SIMD enabled globally
    HasAVX2  bool  // AMD64 AVX2 support
    HasNEON  bool  // ARM64 NEON support
)

func init() {
    detectSIMDCapabilities()
    if os.Getenv("BEVE_DISABLE_SIMD") != "" {
        UseSIMD = false
    }
}
```

**Optimized Thresholds** (M2 Max tuned):
```go
// Reduced based on benchmarks showing 11× speedup for 8 elements
simdThresholdInt32   = 8   // Was 16, now 8
simdThresholdInt64   = 4   // Was 8, now 4
simdThresholdFloat32 = 8   // Was 16, now 8
simdThresholdFloat64 = 4   // Was 8, now 4
```

**Performance**:
```
8× int32 array:
  - SIMD:   15ns  (11× faster)
  - Scalar: 173ns

32× int32 array:
  - SIMD:   45ns  (8× faster)
  - Scalar: 360ns
```

#### 3.2 ARM64 NEON (`simd_arm64.go`, `.s`)

**Strategy**:
- Process 4× int32 or 2× int64 per iteration (128-bit Q registers)
- Zero-copy bulk write via `unsafe.Slice()`
- Hardware prefetcher handles memory access

**Implementation**:
```go
func (e *Encoder) encodeInt32ArraySIMD(data []int32) error {
    // Write BEVE header
    header := byte(0x04 | (1 << 3) | (2 << 5))
    e.WriteByte(header)
    e.WriteCompressedUint(uint64(len(data)))
    
    if len(data) > 0 {
        // Zero-copy: []int32 → []byte
        bytes := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), len(data)*4)
        
        // Bulk write (NEON auto-vectorized)
        e.WriteBytes(bytes)
    }
    
    return nil
}
```

**Assembly** (conceptual, not actual):
```arm
// VLD1.32 {q0}, [r0]!    // Load 4× int32 into Q0 register
// VST1.32 {q0}, [r1]!    // Store 4× int32 to output
```

#### 3.3 AMD64 AVX2 (`simd_amd64.go`, `.s`)

**Strategy**:
- Process 8× int32 or 4× int64 per iteration (256-bit YMM registers)
- Wider vectors = 2× throughput vs ARM64

**Performance** (Theoretical):
```
AMD64 (AVX2):
  - 8× int32 per cycle
  - ~8× faster than scalar

ARM64 (NEON):
  - 4× int32 per cycle
  - ~4× faster than scalar
```

#### 3.4 String SIMD (`simd_string_*.go`, `.s`)

**Purpose**: Vectorized string array encoding.

**Features**:
- Length-prefixed encoding
- SIMD copy for data bytes
- Handles variable-length strings

**Status**: Implemented, benchmarks pending.

#### 3.5 Software Prefetching (Phase 2A - **DISABLED**)

**Purpose**: Software hints to prefetch array data into cache.

**Result**: ⚠️ **Performance Regression** on M2 Max
- Medium: 4.9μs → 5.2μs (**6% slower**)
- Large: 49.6μs → 46.1μs (7% faster, inconsistent)

**Root Cause**: M2 Max hardware prefetcher is exceptionally strong. Software hints add overhead without benefit.

**Configuration**:
```bash
# Enable prefetching (disabled by default)
export BEVE_ENABLE_PREFETCH=true
```

**Files**:
- `prefetch.go` - Wrapper functions
- `prefetch_arm64.s` - ARM64 PRFM PLDL1KEEP
- `prefetch_amd64.s` - AMD64 PREFETCHT0
- `prefetch_generic.go` - No-op fallback

**Recommendation**: Keep disabled. May help on weaker CPUs (Intel, AMD Ryzen).

**See**: [PHASE_2A_PREFETCH_SIMD.md](../PHASE_2A_PREFETCH_SIMD.md)

---

### 4. Memory Management

#### 4.1 Buffer System (`buffer.go`)

**Purpose**: Pooled, growable byte buffers with platform-specific optimizations.

**Features**:
- Initial capacity: 512 bytes (configurable)
- Exponential growth: 2× on overflow
- Max pool capacity: 1MB (prevents unbounded growth)
- Zero-copy slice operations

**Buffer Structure**:
```go
type Buffer struct {
    data []byte   // Backing buffer
    pos  int      // Write position
}

// AcquireBuffer() - Get from pool
// ReleaseBuffer() - Return to pool
```

**Growth Strategy**:
```go
func (b *Buffer) ensureSpace(n int) {
    if b.pos+n > cap(b.data) {
        // Exponential growth: max(2×current, current+needed)
        newCap := max(cap(b.data)*2, cap(b.data)+n)
        newData := make([]byte, len(b.data), newCap)
        copy(newData, b.data)
        b.data = newData
    }
}
```

**Pool Hygiene**:
```go
// Large buffers not pooled to avoid memory bloat
if cap(buffer.data) > maxBufferPoolCapacity {
    // Release, don't pool
    buffer = nil  // GC reclaims
}
```

#### 4.2 Platform-Specific Buffers (`buffer_platform.go`)

**Purpose**: Architecture-specific buffer operations.

**Optimizations**:
- **ARM64**: NEON vector copies for large buffers
- **AMD64**: AVX2 vector copies
- **Generic**: Standard `copy()`

**Auto-Selection**:
```go
//go:build arm64 && !purego
// Use NEON intrinsics for buffer copy
```

#### 4.3 Arena Allocator (`arena.go`)

**Purpose**: Batch allocation for bulk encoding.

**Use Case**: Encode 1000s of structs with single allocation.

**Status**: 🚧 **Planned** (Phase 3B - Batch Arena)

**Design**:
```go
type Arena struct {
    buf    []byte   // Pre-allocated buffer (e.g., 5MB)
    offset int      // Current position
    marks  []int    // Slice boundaries
}

func (a *Arena) EncodeBatch(items []any) ([][]byte, error) {
    results := make([][]byte, 0, len(items))
    
    for _, item := range items {
        start := a.offset
        
        // Encode into arena (zero-copy)
        enc := a.borrowEncoder()
        enc.buf = a.buf[a.offset:]
        enc.Encode(item)
        
        // Slice result
        results = append(results, a.buf[start:a.offset])
    }
    
    return results, nil  // Single allocation!
}
```

**Expected Performance**:
```
100 structs without arena:
  - Time: 85μs
  - Allocations: 300+ (3 per struct)
  - Memory: 250KB

100 structs with arena:
  - Time: 30μs (2.8× faster)
  - Allocations: 3 (arena + results + marks)
  - Memory: 250KB (pre-allocated, reusable)
```

---

### 5. Decoder Pipeline

#### 5.1 Decoder Base (`decoder_base.go`)

**Purpose**: Core decoding structure and type dispatch.

**Key Components**:
```go
type Decoder struct {
    buf     []byte  // Input buffer
    pos     int     // Read position
    scratch [9]byte // Scratch buffer for primitives
}

// NewDecoder(data) - Create decoder
// Decode(v any)    - Decode into value
```

**Decoding Flow**:
```
Input Binary
    │
    ├─ Read type header (1 byte)
    │
    ├─ Dispatch to type-specific decoder
    │   ├─ Primitive? → decoder_primitives.go
    │   ├─ Collection? → decoder_collections.go
    │   └─ Complex? → decoder_utils.go
    │
    └─ Unmarshal into Go value
```

#### 5.2 Fast Path Decoder (`decoder_fast_paths.go`)

**Purpose**: Optimized decoding for common patterns.

**Features**:
- Direct field assignment (no reflection)
- Pre-computed struct layouts
- Type-specific fast paths

**Performance**:
```
Small struct decode:
  - Fast path: 547ns
  - Standard: 1.8μs
  - Speedup: 3.3×
```

---

### 6. Type System (`common.go`, `doc.go`)

#### 6.1 Type Constants (`common.go`)

**BEVE Type Tags**:
```go
const (
    TypeNull         = 0x00
    TypeFalse        = 0x08
    TypeTrue         = 0x18
    
    TypeInt32        = 0x01 | (1<<3) | (2<<5)  // Signed, 4 bytes
    TypeInt64        = 0x01 | (1<<3) | (3<<5)  // Signed, 8 bytes
    
    TypeFloat32      = 0x01 | (0<<3) | (2<<5)  // Float, 4 bytes
    TypeFloat64      = 0x01 | (0<<3) | (3<<5)  // Float, 8 bytes
    
    TypeString       = 0x02
    TypeObject       = 0x03
    TypeTypedArray   = 0x04
    TypeGenericArray = 0x05
)
```

**Header Format**:
```
Bit Layout (1 byte):
  [7:5] - Byte count indicator (0-7 → 1,2,4,8,16,32,64,128 bytes)
  [4:3] - Type group (0=float, 1=signed, 2=unsigned)
  [2:0] - Base type (0=null/bool, 1=number, 2=string, 3=object, 4=typed array, 5=generic array)
```

#### 6.2 Runtime Configuration (`doc.go`)

**Global Flags**:
```go
var (
    // EnablePrefetch - Software prefetching (default: false)
    EnablePrefetch = false
    
    // Set via: BEVE_ENABLE_PREFETCH=true
)

func init() {
    if val := os.Getenv("BEVE_ENABLE_PREFETCH"); val != "" {
        EnablePrefetch = parseBool(val)
    }
}
```

**Available Environment Variables**:
- `BEVE_DISABLE_SIMD=true` - Disable SIMD (fallback to scalar)
- `BEVE_ENABLE_PREFETCH=true` - Enable software prefetching
- `BEVE_USE_SYNC_POOL=true` - Force sync.Pool (Phase 3 fallback)

---

## Performance Systems

### Optimization Phases

#### ✅ Phase 1: Zero-Allocation Encoding (Complete)

**Scope**: Stack encoding + metadata cache

**Results**:
- Small struct: 600ns → **143-253ns** (2.4-4.2× faster)
- Memory: 78-99% reduction
- Coverage: 80-90% of structs

**Files**:
- `encoder_stack.go` - Stack-based encoding
- `encoder_cache.go` - Metadata cache
- `encoder_fast_path.go` - Fast path routing

**Documentation**: (Phase 1 docs deleted, integrate into this README)

#### ✅ Phase 2: SIMD Acceleration (Complete)

**Scope**: SIMD arrays + string optimizations

**Results**:
- Array encoding: 4-11× faster
- Thresholds optimized for M2 Max
- Platform-specific implementations

**Files**:
- `simd.go` - Infrastructure
- `simd_arm64.go/.s` - NEON
- `simd_amd64.go/.s` - AVX2
- `simd_string_*.go/.s` - String SIMD

**Status**: Production ready, well-tested

#### ⚠️ Phase 2A: Software Prefetching (Disabled)

**Scope**: Cache prefetch hints

**Result**: Performance regression on M2 Max
- Disabled by default
- Infrastructure kept for testing

**Files**:
- `prefetch.go` - Wrappers
- `prefetch_arm64.s` - ARM64 PRFM
- `prefetch_amd64.s` - AMD64 PREFETCHT0

**Documentation**: [PHASE_2A_PREFETCH_SIMD.md](../PHASE_2A_PREFETCH_SIMD.md)

#### 🚧 Phase 3A: Lock-Free Pooling (Planned)

**Scope**: Per-P encoder pools

**Goal**: Eliminate sync.Pool mutex contention

**Expected Results**:
- Small: 253ns → **180ns** (1.4× faster)
- Medium: 5.2μs → **4.0μs** (1.3× faster)
- Large: 46μs → **35μs** (1.3× faster)

**Design**:
```go
type encoderStack struct {
    _      [128]byte     // Cache-line padding
    head   *Encoder      // Stack head
    count  int32         // Pool depth
    _      [128]byte     // Cache-line padding
}

var perPPools []*encoderStack  // One per CPU core

func GetEncoderFromPoolLockFree() *Encoder {
    pid := runtime_procPin()    // Pin to current P
    defer runtime_procUnpin()
    
    stack := perPPools[pid]
    
    // Lock-free pop with atomic CAS
    for {
        head := stack.head
        if head == nil {
            return NewEncoder()
        }
        
        if atomic.CompareAndSwapPointer(&stack.head, head, head.next) {
            return head
        }
    }
}
```

**Status**: Design phase, implementation starting

**Complexity**: HIGH (runtime internals, atomic ops)

**Files** (planned):
- `encoder_pool_lockfree.go` - Per-P pool implementation
- `encoder_pool_lockfree_test.go` - Benchmarks & tests
- `PHASE_3A_LOCKFREE_POOL.md` - Documentation

#### 🚧 Phase 3B: Batch Arena (Planned)

**Scope**: Arena allocator for bulk encoding

**Goal**: Single allocation for 1000s of structs

**Expected Results**:
- 100 structs: 85μs → **30μs** (2.8× faster)
- Allocations: 300+ → **3** (99% reduction)

**Status**: Design phase

**Files** (planned):
- `arena.go` - Already exists, needs batch encode
- `arena_test.go` - Tests
- `PHASE_3B_BATCH_ARENA.md` - Documentation

#### 🚧 Phase 4: Parallel Encoding (Planned)

**Scope**: Multi-threaded array encoding with work stealing

**Goal**: 8-10× speedup for large arrays

**Expected Results**:
- 10K array: 50ms → **6ms** (8.3× faster)
- 100K array: 500ms → **55ms** (9.1× faster)

**Design**:
```go
type WorkerPool struct {
    workers   int
    taskQueue chan func()
    wg        sync.WaitGroup
}

func (pe *ParallelEncoder) EncodeArrayParallel(data []any) ([]byte, error) {
    // Split into chunks
    chunks := splitIntoChunks(data, chunkSize)
    
    // Parallel encode
    for i, chunk := range chunks {
        pool.Submit(func() {
            enc := GetEncoderFromPool()
            defer PutEncoderToPool(enc)
            
            for _, item := range chunk {
                enc.Encode(item)
            }
            
            results[i] = enc.Bytes()
        })
    }
    
    // Merge results
    return mergeChunks(results), nil
}
```

**Status**: Design phase

**Complexity**: HIGH (concurrency, work stealing)

**Files** (planned):
- `encoder_parallel.go` - Worker pool & parallel encoding
- `encoder_parallel_test.go` - Benchmarks
- `PHASE_4_PARALLEL_ENCODING.md` - Documentation

---

## Memory Management

### Current System (sync.Pool)

**Architecture**:
```
┌─────────────────────────────────────┐
│       Global encoderPool            │
│         (sync.Pool)                 │
│                                     │
│  ┌───────────────────────────────┐ │
│  │  Mutex Lock (contention!)     │ │
│  └───────────────────────────────┘ │
│                                     │
│  ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ │
│  │ Enc │ │ Enc │ │ Enc │ │ Enc │ │
│  └─────┘ └─────┘ └─────┘ └─────┘ │
└─────────────────────────────────────┘
```

**Characteristics**:
- ✅ Simple, battle-tested
- ✅ Automatic GC cleanup
- ❌ Mutex contention under high concurrency
- ❌ Encoders migrate across CPU cores (cache miss)

**Performance**:
- Single goroutine: Excellent
- 100+ goroutines: Degradation from lock contention

### Planned System (Phase 3A - Lock-Free Per-P)

**Architecture**:
```
┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐
│  CPU 0   │ │  CPU 1   │ │  CPU 2   │ │  CPU 3   │
│          │ │          │ │          │ │          │
│ ┌──────┐ │ │ ┌──────┐ │ │ ┌──────┐ │ │ ┌──────┐ │
│ │ Pool │ │ │ │ Pool │ │ │ │ Pool │ │ │ │ Pool │ │
│ └──────┘ │ │ └──────┘ │ │ └──────┘ │ │ └──────┘ │
│          │ │          │ │          │ │          │
│ Enc Enc  │ │ Enc Enc  │ │ Enc Enc  │ │ Enc Enc  │
└──────────┘ └──────────┘ └──────────┘ └──────────┘
   No lock!     No lock!     No lock!     No lock!
```

**Characteristics**:
- ✅ Zero contention (per-P, no sharing)
- ✅ CPU locality (encoder stays on same core)
- ✅ L1/L2 cache hot
- ⚠️ More complex (runtime internals)
- ⚠️ Higher memory (12 pools on 12-core)

**Expected Performance**:
- Single goroutine: Same or better
- 100+ goroutines: **1.3-1.4× faster**

---

## Thread Safety

### Thread-Safe Components

✅ **Pool Functions**:
- `GetEncoderFromPool()` - Safe (sync.Pool)
- `PutEncoderToPool()` - Safe (sync.Pool)

✅ **Buffer Pool**:
- `AcquireBuffer()` - Safe (sync.Pool)
- `ReleaseBuffer()` - Safe (sync.Pool)

✅ **SIMD Detection**:
- `UseSIMD`, `HasAVX2`, `HasNEON` - Read-only after init

✅ **Metadata Cache**:
- `structCache` - Safe (sync.Map)

### NOT Thread-Safe

❌ **Encoder**:
- Each goroutine must use separate encoder
- Sharing encoder = race condition

❌ **Buffer**:
- Must not share buffer across goroutines

❌ **Decoder**:
- Each goroutine must use separate decoder

### Best Practices

```go
// ✅ CORRECT: Each goroutine gets own encoder
func processItems(items []Item) {
    var wg sync.WaitGroup
    
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            
            enc := GetEncoderFromPool()
            defer PutEncoderToPool(enc)
            
            data, _ := enc.EncodeAndDetach(item)
            // Use data...
        }(item)
    }
    
    wg.Wait()
}

// ❌ WRONG: Sharing encoder across goroutines
var sharedEnc = GetEncoderFromPool()  // DON'T DO THIS!

func processItem(item Item) {
    data, _ := sharedEnc.EncodeAndDetach(item)  // RACE!
}
```

---

## Development Guide

### Adding a New Encoder

1. **Add encoder function to appropriate module**:
   ```go
   // encoder_primitives.go or encoder_collections.go
   func (e *Encoder) encodeMyType(v reflect.Value) error {
       // Implementation...
   }
   ```

2. **Add to type dispatch** (`encoder_base.go`):
   ```go
   func getEncoderFunc(t reflect.Type) encoderFunc {
       switch t.Kind() {
       // ... existing cases
       case reflect.MyType:
           return (*Encoder).encodeMyType
       }
   }
   ```

3. **Add tests**:
   ```go
   // encoder_test.go or dedicated test file
   func TestEncodeMyType(t *testing.T) {
       // Test cases...
   }
   ```

4. **Add benchmarks**:
   ```go
   func BenchmarkEncodeMyType(b *testing.B) {
       enc := GetEncoderFromPool()
       defer PutEncoderToPool(enc)
       
       data := MyType{...}
       
       b.ReportAllocs()
       for i := 0; i < b.N; i++ {
           enc.Encode(data)
           enc.Buf.Reset()
       }
   }
   ```

### Performance Testing

```bash
# Run benchmarks
go test -bench=. -benchmem ./core/

# CPU profiling
go test -bench=BenchmarkLarge -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profiling
go test -bench=BenchmarkLarge -memprofile=mem.prof
go tool pprof mem.prof

# Race detection
go test -race ./core/

# Coverage
go test -cover ./core/
```

### Code Style

**Performance Guidelines**:
1. Minimize allocations (use pooling)
2. Avoid reflection in hot paths
3. Use `unsafe` judiciously (with safety checks)
4. Profile before optimizing
5. Document performance characteristics

**Cache Optimization**:
1. Group hot fields together (first 64 bytes)
2. Align structs to cache lines (128 bytes on ARM64)
3. Pad between unrelated data structures
4. Avoid false sharing in concurrent code

**Example**:
```go
type HotStruct struct {
    // Hot path (first cache line)
    count    int
    ptr      *Buffer
    flags    uint32
    
    // Cold path (second cache line)
    _        [64]byte  // Padding
    scratch  [256]byte // Rarely accessed
}
```

---

## Summary

### Current State (v1.3.0)

✅ **Production Ready**:
- Phase 1: Stack + Cache (143-253ns)
- Phase 2: SIMD (4-11× faster arrays)
- Comprehensive test coverage
- Multi-platform support (ARM64, AMD64)

⚠️ **Experimental**:
- Phase 2A: Prefetching (disabled, available for testing)

🚧 **Planned**:
- Phase 3A: Lock-Free Pooling (1.3-1.4× improvement)
- Phase 3B: Batch Arena (2-3× for bulk ops)
- Phase 4: Parallel Encoding (8-10× for large arrays)

### Performance Profile

```
Small Structs (<10 fields):
  ├─ Stack:  143ns  ⚡ Best for primitives
  ├─ Cache:  181ns  ⚡ Best for mixed types
  └─ Std:    600ns  ⚪ Baseline

Medium Payloads (100-1000 items):
  ├─ ZeroCopy:  5.2μs   ⚡ Recommended
  ├─ Standard:  9.0μs   ⚪ Good
  └─ JSON:      31μs    🐌 6× slower

Large Payloads (10K+ items):
  ├─ ZeroCopy:  46μs    ⚡ Best
  ├─ Standard:  67μs    ⚪ Good
  ├─ JSON:      286μs   🐌 6× slower
  └─ Parallel:  6ms*    🚀 Future (8× faster)
  
* Phase 4 projected
```

### Key Takeaways

1. **Use the pool**: `GetEncoderFromPool()` / `PutEncoderToPool()`
2. **Prefer fast paths**: Small structs auto-optimized
3. **Trust SIMD**: Arrays ≥8 elements benefit
4. **Disable prefetch**: Enabled only for specific CPUs
5. **One encoder per goroutine**: Thread safety

---

**Questions? Issues?**  
See: [GitHub Issues](https://github.com/beve-org/beve-go/issues)

**Contributing?**  
See: [CONTRIBUTING.md](../../CONTRIBUTING.md)

**License**: MIT  
**Maintainer**: @meftunca
