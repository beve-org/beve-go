# BEVE Performance Optimization Plan

## 🔍 Current Performance Analysis

### Platform-Specific Issues

| Platform | Small Struct Marshal | Memory | Issue |
|----------|---------------------|--------|-------|
| **Windows AMD64** | 3014 ns/op | 2839 B | 3x slower than macOS |
| **Linux AMD64** | 978.5 ns/op | 723 B | Good speed, but memory varies |
| **macOS ARM64** | 1150 ns/op | 1938 B | Baseline performance |

### Root Causes

1. **Excessive allocations on Windows** - 2839 bytes vs 723 on Linux
2. **Buffer growth overhead** - Different OS memory allocators behave differently
3. **Reflection overhead** - `reflect.Value` operations not optimized
4. **Type switching cost** - Multiple type checks per encode

---

## 🎯 Optimization Strategies

### Phase 1: Buffer Optimizations (High Impact)

#### 1.1 Pre-allocate buffers based on type size hints
**Problem**: Buffer grows dynamically, causing multiple allocations
**Solution**: Add size hints for common patterns

```go
// Estimate size before encoding
func estimateSize(v reflect.Value) int {
    switch v.Kind() {
    case reflect.Struct:
        return 256 // typical small struct
    case reflect.Slice:
        return v.Len() * 8 + 16
    case reflect.Map:
        return v.Len() * 16 + 16
    default:
        return 64
    }
}

// In GetEncoderFromPool:
enc.Buf.Grow(estimateSize(v)) // pre-grow before encoding
```

**Expected gain**: -30% allocations, -15% time

#### 1.2 Platform-specific buffer initial capacity
**Problem**: defaultBufferCapacity = 512 too small for Windows
**Solution**: Adjust based on OS

```go
const (
    defaultBufferCapacityWindows = 1024 // Windows needs more
    defaultBufferCapacityUnix    = 512
)

func getBufferInitialCapacity() int {
    if runtime.GOOS == "windows" {
        return defaultBufferCapacityWindows
    }
    return defaultBufferCapacityUnix
}
```

**Expected gain**: -25% Windows allocations

#### 1.3 Inline small writes
**Problem**: WriteByte/WriteBytes function call overhead
**Solution**: Use direct buffer append for hot paths

```go
// Instead of:
e.WriteByte(b)

// Use:
e.Buf.data = append(e.Buf.data, b)

// Batch small writes:
e.Buf.data = append(e.Buf.data, b1, b2, b3, b4)
```

**Expected gain**: -10% time for small operations

---

### Phase 2: Reflection Optimizations (Medium Impact)

#### 2.1 Type cache with fast path
**Problem**: Type info computed on every encode
**Solution**: Cache encoder functions per type

```go
var encoderCache sync.Map // map[reflect.Type]encoderFunc

func (e *Encoder) Encode(v reflect.Value) error {
    typ := v.Type()
    
    // Fast path: check cache first
    if fn, ok := encoderCache.Load(typ); ok {
        return fn.(encoderFunc)(e, v)
    }
    
    // Slow path: build and cache
    fn := buildEncoder(typ)
    encoderCache.Store(typ, fn)
    return fn(e, v)
}
```

**Expected gain**: -20% time, -30% allocs for repeated types

#### 2.2 Eliminate reflect.ValueOf in hot paths
**Problem**: `reflect.ValueOf` allocation overhead
**Solution**: Pass `reflect.Value` directly when possible

```go
// Public API keeps interface{}
func Marshal(v interface{}) ([]byte, error) {
    rv := reflect.ValueOf(v)
    return marshalValue(rv) // internal uses Value
}

// Internal functions use Value
func marshalValue(v reflect.Value) ([]byte, error)
```

**Expected gain**: -1 alloc per marshal

---

### Phase 3: CPU-Specific Optimizations (Low Impact, High Complexity)

#### 3.1 SIMD for typed arrays
**Problem**: Byte-by-byte copy in typed array encoding
**Solution**: Use assembly for bulk operations on x86/ARM

```go
//go:build amd64
func encodeInt32SliceFast(dst []byte, src []int32)

//go:build arm64
func encodeInt32SliceFast(dst []byte, src []int32)
```

**Expected gain**: -40% time for large typed arrays

#### 3.2 Branch prediction hints
**Problem**: Unpredictable branches in type switch
**Solution**: Reorder cases by frequency

```go
// Most common types first
switch v.Kind() {
case reflect.Int, reflect.Int64:  // Most common
    return e.encodeInt(v)
case reflect.String:               // Second most common
    return e.encodeString(v)
case reflect.Struct:               // Third
    return e.encodeStruct(v)
// ... less common types
}
```

**Expected gain**: -5% time from better branch prediction

---

### Phase 4: Memory Layout Optimizations

#### 4.1 Reduce Encoder struct size
**Problem**: Large encoder struct causes cache misses
**Solution**: Move rarely-used fields to separate struct

```go
type Encoder struct {
    Buf         *Buffer    // Hot: always used
    Writer      io.Writer  // Hot: often used
    
    // Cold fields - move to separate config
    config      *EncoderConfig
}

type EncoderConfig struct {
    uintScratch   [10]byte
    varintScratch [10]byte
    // ... other scratch buffers
}
```

**Expected gain**: Better CPU cache utilization

#### 4.2 Align critical fields
**Problem**: Misaligned struct fields cause extra memory loads
**Solution**: Explicit padding for cache line alignment

```go
type Encoder struct {
    Buf    *Buffer   // 8 bytes
    Writer io.Writer // 16 bytes
    _      [8]byte   // padding to 32 bytes (cache line aligned)
    // ...
}
```

---

## 📊 Expected Overall Improvements

| Optimization Phase | Time Improvement | Alloc Reduction | Complexity |
|-------------------|------------------|-----------------|------------|
| Phase 1: Buffers | -25% | -40% | Low |
| Phase 2: Reflection | -20% | -30% | Medium |
| Phase 3: CPU-specific | -15% | 0% | High |
| Phase 4: Memory layout | -10% | 0% | Medium |
| **Total Estimated** | **-50%** | **-55%** | - |

### Platform-Specific Targets

| Platform | Current | Target | Strategy |
|----------|---------|--------|----------|
| Windows | 3014 ns | <1500 ns | Phase 1 + 2 |
| Linux | 978 ns | <700 ns | Phase 2 + 3 |
| macOS | 1150 ns | <800 ns | Phase 1 + 2 |

---

## 🚀 Implementation Priority

### Quick Wins (1-2 hours)
1. ✅ Pre-allocate buffers (1.1)
2. ✅ Platform-specific capacity (1.2)
3. ✅ Inline small writes (1.3)

### Medium Effort (1 day)
4. ✅ Type cache (2.1)
5. ✅ Reflect.Value API (2.2)
6. ✅ Branch reordering (3.2)

### Long Term (1 week)
7. ⏳ SIMD optimizations (3.1)
8. ⏳ Memory layout (4.1, 4.2)

---

## 🧪 Benchmarking Strategy

After each optimization:
```bash
# Before
go test -bench=. -benchmem -benchtime=10000x > before.txt

# Apply optimization

# After
go test -bench=. -benchmem -benchtime=10000x > after.txt

# Compare
benchstat before.txt after.txt
```

Track regression on all platforms via GitHub Actions.

---

## 📝 Notes

- Focus on Phase 1 first - highest ROI, lowest risk
- Phase 3 requires platform-specific testing
- Keep backward compatibility
- Add benchmarks for each optimization
- Document performance characteristics in godoc

