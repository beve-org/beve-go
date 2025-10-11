# GitHub Copilot Instructions for BEVE-Go

## 🎯 Project Overview

**BEVE (Binary Encoded Values)** is a high-performance binary serialization library for Go, optimized for speed, efficiency, and type safety. This project prioritizes **performance**, **memory efficiency**, and **developer experience**.

### Key Performance Metrics
- **5.6× faster** than CBOR on small struct marshaling
- **64% smaller** payloads than JSON
- **22× fewer allocations** than CBOR on unmarshaling
- Target: <1μs for small struct encoding

---

## 🏗️ Architecture Principles

### 1. Performance-First Design
Every code change must consider:
- **CPU cache efficiency** (struct field layout matters!)
- **Allocation minimization** (use object pooling, pre-allocated buffers)
- **Branch prediction** (most common cases first in switch/if statements)
- **Unsafe operations** (when safe and measurably faster)

### 2. Memory Layout Optimization
```go
// ✅ GOOD: Hot fields in first cache line (64 bytes)
type Encoder struct {
    Buf *Buffer      // 8 bytes - most accessed
    w   io.Writer    // 8 bytes - second most accessed
    scratch [24]byte // 24 bytes - frequently used
    counter int      // 8 bytes
    // Total: 48 bytes (fits in one cache line)
    
    coldData [256]byte // rarely accessed, second cache line
}

// ❌ BAD: Hot and cold fields mixed
type Encoder struct {
    coldData [256]byte // pushes hot fields to other cache lines!
    Buf *Buffer
    w   io.Writer
}
```

### 3. Pooling Strategy
- Use `sync.Pool` for frequently allocated objects (Encoder, Decoder, Buffer)
- Cap pooled buffer sizes at 1MB to prevent unbounded memory growth
- Always reset state before returning to pool

### 4. Reflection Optimization
```go
// ✅ GOOD: Fast path for primitives (no reflection)
func Marshal(v interface{}) ([]byte, error) {
    switch val := v.(type) {
    case int:
        return marshalInt(val) // direct, no reflect.ValueOf
    case string:
        return marshalString(val)
    default:
        return marshalReflect(reflect.ValueOf(v)) // fallback
    }
}

// ❌ BAD: Always using reflection
func Marshal(v interface{}) ([]byte, error) {
    return marshalReflect(reflect.ValueOf(v)) // slow for primitives
}
```

---

## 📝 Coding Standards

### Naming Conventions
- **Exported functions**: PascalCase (`Marshal`, `Encode`, `GetEncoderFromPool`)
- **Internal functions**: camelCase (`encodeInt`, `writeVarint`, `ensureAddressable`)
- **Constants**: UPPER_SNAKE_CASE (`MAX_BUFFER_SIZE`, `TYPE_INT32`)
- **Struct tags**: lowercase with commas (`beve:"name,omitempty"`)

### Function Design
```go
// ✅ GOOD: Clear documentation, performance notes
// encodeInt encodes a signed integer using varint encoding.
//
// Performance: ~5ns for small values (<128), up to 20ns for large values.
// Uses pre-allocated scratch buffer to avoid allocations.
//
// Format: First byte contains value | length indicator
//go:inline
func (e *Encoder) encodeInt(value int64) error {
    // implementation
}

// ❌ BAD: No docs, unclear purpose
func (e *Encoder) doInt(v int64) error {
    // implementation
}
```

### Error Handling
```go
// ✅ GOOD: Specific error types
type UnsupportedError struct {
    Message string
}

func (e *UnsupportedError) Error() string {
    return "beve: " + e.Message
}

// Usage
return &UnsupportedError{"cannot encode channel type"}

// ❌ BAD: Generic errors
return errors.New("bad type")
```

### Unsafe Code Guidelines
**When to use `unsafe`:**
- ✅ Zero-copy string → []byte conversion (read-only)
- ✅ Struct field access with known offsets
- ✅ Performance-critical paths with benchmarks proving >10% improvement

**When NOT to use `unsafe`:**
- ❌ Pointer arithmetic without proper bounds checking
- ❌ Type conversions that violate Go memory model
- ❌ Any case that causes `checkptr` errors

```go
// ✅ GOOD: Safe string→[]byte (read-only)
func stringToBytes(s string) []byte {
    if len(s) == 0 {
        return nil
    }
    return unsafe.Slice(unsafe.StringData(s), len(s))
}

// ❌ BAD: Unsafe pointer arithmetic
func getFieldPtr(base unsafe.Pointer, offset uintptr) unsafe.Pointer {
    return unsafe.Pointer(uintptr(base) + offset) // potential invalid pointer!
}

// ✅ GOOD: Safe offset arithmetic
func getFieldPtr(base unsafe.Pointer, offset uintptr) unsafe.Pointer {
    return unsafe.Add(base, offset) // Go 1.17+ safe version
}
```

---

## 🧪 Testing Requirements

### Must-Have Tests for New Features
1. **Unit tests**: Basic functionality
2. **Edge case tests**: nil, empty, zero values
3. **Race tests**: `go test -race` must pass
4. **Benchmarks**: Compare with baseline

### Benchmark Standards
```go
// ✅ GOOD: Complete benchmark with memory tracking
func BenchmarkEncodeSmallStruct(b *testing.B) {
    data := SmallStruct{ID: 123, Name: "test"}
    b.ReportAllocs()
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _, err := Marshal(data)
        if err != nil {
            b.Fatal(err)
        }
    }
}

// Run with: go test -bench=. -benchmem -benchtime=10000x
```

### Performance Regression Tests
Before committing optimizations:
```bash
# 1. Baseline
go test -bench=BenchmarkSmallStruct_BEVE_Marshal -benchmem -benchtime=10000x > bench_before.txt

# 2. Make changes

# 3. Compare
go test -bench=BenchmarkSmallStruct_BEVE_Marshal -benchmem -benchtime=10000x > bench_after.txt

# 4. Verify improvement (or revert if regression)
```

**Acceptance criteria:**
- ✅ Time: Same or faster (no >5% regression)
- ✅ Memory: Same or less (no >10% increase)
- ✅ Allocations: Same or fewer

---

## 🚀 Common Optimization Patterns

### Pattern 1: Pre-allocated Scratch Buffers
```go
type Encoder struct {
    scratch [32]byte // reuse for temporary data
}

func (e *Encoder) encodeVarint(n uint64) error {
    // Use scratch instead of allocating []byte
    e.scratch[0] = byte(n << 2)
    return e.WriteByte(e.scratch[0])
}
```

### Pattern 2: Inline Hints for Hot Paths
```go
//go:inline
func (e *Encoder) WriteByte(b byte) error {
    if e.Buf != nil {
        return e.Buf.WriteByte(b) // fast path
    }
    return e.w.WriteByte(b) // slow path
}
```

### Pattern 3: Type Switch for Common Cases
```go
// ✅ GOOD: Common types first
switch v.Kind() {
case reflect.Int, reflect.Int64: // most common
    return e.encodeInt(v.Int())
case reflect.String: // second most common
    return e.EncodeString(v.String())
case reflect.Slice: // third most common
    return e.encodeSlice(v)
default: // rare types
    return e.encodeGeneric(v)
}
```

### Pattern 4: Lock-Free Caching
```go
// ✅ GOOD: sync.Map for concurrent reads, rare writes
var typeCache sync.Map

func getEncoder(t reflect.Type) encoderFunc {
    if enc, ok := typeCache.Load(t); ok {
        return enc.(encoderFunc)
    }
    
    enc := buildEncoder(t)
    typeCache.Store(t, enc)
    return enc
}
```

---

## 🐛 Debugging Tips

### Performance Issues
1. **Profile first**: `go test -cpuprofile=cpu.out -bench=.`
2. **Analyze**: `go tool pprof cpu.out`
3. **Check allocations**: `go test -memprofile=mem.out -bench=.`
4. **Look for**:
   - Excessive `runtime.newobject` calls
   - `reflect.Value` operations in hot paths
   - Unnecessary type conversions

### Memory Leaks
```bash
# Heap profile
go test -memprofile=mem.out -bench=.
go tool pprof -alloc_space mem.out

# Look for:
# - Large buffer pooling (>1MB retained)
# - Circular references preventing GC
# - Leaked goroutines holding memory
```

### Race Conditions
```bash
go test -race ./...

# Common issues:
# - sync.Map vs regular map confusion
# - Encoder/Decoder used across goroutines
# - Shared buffers without synchronization
```

---

## 📋 Pre-Commit Checklist

Before submitting PR:
- [ ] `go test ./...` passes
- [ ] `go test -race ./...` passes
- [ ] `golangci-lint run` passes (no errors)
- [ ] Benchmarks show no regression (or justified)
- [ ] Documentation updated (GoDoc, README if needed)
- [ ] Commit message follows convention: `<type>(<scope>): <subject>`
- [ ] Added benchmark files if performance-sensitive change

---

## 🎨 Code Generation Guidelines

### When suggesting struct definitions:
```go
// ✅ GOOD: Annotated with performance notes
type MyStruct struct {
    // Hot path fields first (cache efficiency)
    ID     int64  `beve:"id"`
    Status byte   `beve:"status"`
    
    // Cold path fields last
    Metadata map[string]string `beve:"metadata,omitempty"`
}
```

### When suggesting encoder functions:
```go
// ✅ GOOD: Include inline hint if hot path
//go:inline
func (e *Encoder) encodeSmallInt(n int) error {
    if n < 64 {
        return e.WriteByte(byte(n << 2))
    }
    return e.encodeInt(int64(n))
}
```

### When suggesting tests:
```go
// ✅ GOOD: Table-driven with edge cases
func TestEncodeInt(t *testing.T) {
    tests := []struct{
        name  string
        input int64
        want  []byte
    }{
        {"zero", 0, []byte{0x00}},
        {"small positive", 10, []byte{0x28}},
        {"negative", -1, []byte{0x01, 0xFF}},
        {"max int64", math.MaxInt64, []byte{...}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

---

## 🔍 Common Pitfalls to Avoid

### ❌ Don't: Modify pooled objects after returning
```go
// BAD
enc := GetEncoderFromPool()
data := enc.Buf.Bytes()
PutEncoderToPool(enc) // enc.Buf may be reused!
return data // data may be corrupted!

// GOOD
enc := GetEncoderFromPool()
data := append([]byte(nil), enc.Buf.Bytes()...) // copy
PutEncoderToPool(enc)
return data
```

### ❌ Don't: Use reflection when type is known
```go
// BAD
func Marshal(v interface{}) ([]byte, error) {
    rv := reflect.ValueOf(v)
    return marshalValue(rv) // slow for all types
}

// GOOD
func Marshal(v interface{}) ([]byte, error) {
    // Fast path for common types
    switch val := v.(type) {
    case int:
        return marshalInt(val)
    case string:
        return marshalString(val)
    default:
        return marshalValue(reflect.ValueOf(v))
    }
}
```

### ❌ Don't: Allocate in loops
```go
// BAD
for _, item := range items {
    buf := make([]byte, 100) // allocates every iteration!
    // use buf
}

// GOOD
buf := make([]byte, 100)
for _, item := range items {
    buf = buf[:0] // reset without allocating
    // use buf
}
```

### ❌ Don't: Ignore Go version compatibility
```go
// BAD: Requires Go 1.20+
func example() {
    m := map[string]int{} 
    clear(m) // only in Go 1.21+
}

// GOOD: Check go.mod version (currently 1.22)
// If using features from 1.22+, ensure go.mod matches
```

---

## 📚 Key Files Reference

### Core Encoding/Decoding
- `beve.go` - Public API (Marshal, Unmarshal)
- `core/encoder_base.go` - Encoder struct and pooling
- `core/encoder_primitives.go` - Primitive type encoding (int, float, bool, string)
- `core/encoder_collections.go` - Complex types (struct, slice, map, array)
- `core/encoder_write.go` - Low-level write operations (byte, varint, buffer)
- `core/decoder.go` - Decoder implementation

### Performance
- `core/buffer.go` - Buffer pooling and management
- `core/buffer_platform.go` - Platform-specific optimizations (Windows vs Unix)
- `core/common.go` - Type dispatch and caching
- `PERFORMANCE_SUMMARY.md` - Optimization journey documentation
- `OPTIMIZATION_PLAN.md` - Future optimization strategies

### Testing & Benchmarks
- `*_test.go` - Unit tests
- `comparison_test.go` - Benchmarks vs JSON/CBOR/MessagePack
- `bench_phase*.txt` - Historical benchmark results

### CI/CD
- `.github/workflows/ci.yml` - Lint, test, coverage
- `.github/workflows/benchmarks.yml` - Multi-platform benchmarks

---

## 🤝 Collaboration Style

### When reviewing code:
- Focus on **performance impact** first
- Suggest **concrete improvements** with code examples
- Provide **benchmark evidence** for optimization claims
- Be **constructive** and **specific**

### When writing code:
- **Document why**, not just what
- Include **performance notes** in comments
- Add **inline hints** (`//go:inline`) where proven beneficial
- Run **benchmarks** before claiming improvements

---

## 🎯 Project Goals

### Short-term (Current Phase)
- ✅ Achieve <1μs small struct marshal
- ✅ Reduce allocations to <5 per operation
- ⏳ Multi-platform benchmarks (Linux, Windows, macOS, ARM)
- ⏳ Improve Windows performance (currently 3× slower)

### Medium-term
- [ ] SIMD optimizations for bulk array encoding
- [ ] Streaming API for large payloads
- [ ] Code generation for struct marshaling
- [ ] WebAssembly support

### Long-term
- [ ] Become fastest binary serialization in Go ecosystem
- [ ] Sub-500ns for small structs
- [ ] Zero-allocation mode for pre-allocated buffers
- [ ] Cross-language specification (BEVE format standard)

---

## 🏆 Success Criteria

A contribution is successful when:
1. **Tests pass**: `go test -race ./...` ✅
2. **Lint passes**: `golangci-lint run` ✅
3. **Benchmarks stable**: No >5% regression ✅
4. **Documentation complete**: GoDoc + README updates ✅
5. **Performance justified**: Benchmarks prove claims ✅

---

## 💡 Final Tips

- **Measure, don't guess**: Always benchmark before claiming performance improvements
- **Cache-friendly is fast**: Think about CPU cache lines in struct layout
- **Unsafe is sharp**: Use only when proven necessary and safe
- **Tests are documentation**: Write clear, comprehensive tests
- **Readability matters**: Performance code should still be maintainable

**Remember**: This is a performance-critical library. Every nanosecond counts, but so does code clarity and safety. Balance is key.

---

_Last updated: October 2025_
_Go version: 1.22+_
_Maintained by: BEVE-org team_
