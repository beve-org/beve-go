# GitHub Copilot Instructions for BEVE-Go

## ⚠️ CRITICAL: Terminal Usage Guidelines

### DO NOT Send Large Commands to Terminal

**VSCode PTY Host crashes with large terminal outputs!**

❌ **NEVER DO THIS:**
- Long heredoc commands (`cat << 'EOF' ... EOF`) with >50 lines
- Large code blocks in terminal
- Multi-page reports/summaries via echo/cat
- Verbose benchmark outputs without filtering

✅ **ALWAYS DO THIS:**
- Create files for reports: `create_file` tool
- Use `tee` to save outputs: `go test ... | tee result.txt`
- Filter outputs: `| head -20`, `| grep pattern`
- Split large commands into smaller ones
- Use `>` or `>>` for file redirection instead of heredoc

**Example - Wrong:**
```bash
cat << 'EOF'
[500 lines of report...]
EOF
```

**Example - Correct:**
```bash
# Create file instead
create_file("REPORT.md", content)

# Or filter output
go test -bench=. | head -30 | tee results.txt
```

## 🎯 Project Overview

**BEVE (Binary Encoded Values)** is a high-performance binary serialization library for Go, optimized for speed, efficiency, and type safety. This project prioritizes **EXTREME PERFORMANCE**, **MEMORY EFFICIENCY**, and **DEVELOPER EXPERIENCE**.

### Key Performance Metrics

- **5.6× faster** than CBOR on small struct marshaling
- **64% smaller** payloads than JSON
- **22× fewer allocations** than CBOR on unmarshaling
- Target: <1μs for small struct encoding

## 🏗️ Architecture Principles

### 1. Performance-First Design

Every code change must consider:

- **CPU cache efficiency** (struct field layout matters!)
- **Allocation minimization** (use object pooling, pre-allocated buffers)
- **Branch prediction** (most common cases first in switch/if statements)
- **Unsafe operations** (when safe and measurably faster)

### 2. Memory Layout Optimization

```javascript
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

```javascript
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

## 🔥 Seviye 2: Agresif Optimizasyon Prensipleri

Bu bölümdeki teknikler, standart optimizasyonların yetersiz kaldığı, her nanosaniyenin kritik olduğu ve bunun için kod karmaşıklığı ile bakım maliyetinin artırılabileceği durumlar için tasarlanmıştır. Bu teknikleri uygulamadan önce **iki kez ölçüm yapın.**

### 1. Go Assembly (Plan 9) Direktifleri

Go derleyicisinin ürettiği kodun bile yavaş kaldığı en kritik (hot-path) fonksiyonlar için Go Assembly kullanın. Bu, optimizasyonun zirvesidir.

- **Ne zaman kullanılır:** `varint` kodlama/çözme, `byte` dizilerini toplu işleme, hashing gibi küçük ve sürekli tekrar eden, döngü-yoğun işlemler.
- **Kural:** Assembly implementasyonu, standart Go implementasyonuna göre en az **%20**'lik bir performans artışı sağlamalıdır.

```javascript
// ✅ AGRESİF: Kritik fonksiyon için Assembly implementasyonu
//go:noescape
func encodeVarintAsm(buf []byte, v uint64) int

// encode_varint_amd64.s dosyasında bu fonksiyonun
// assembly karşılığı bulunur.


```

### 2. `unsafe`'in Sınırlarını Zorlamak

`unsafe` paketi sadece sıfır kopyalı dönüşümler için değil, aynı zamanda reflection'ı tamamen ortadan kaldırmak için de kullanılabilir.

- **Sıfır Kopyalı&#32;`[]byte`&#32;→&#32;`string`&#32;Dönüşümü:** Okuma amaçlı `[]byte`'ları string'e dönüştürürken `string()` dönüşümünün yarattığı kopyalamadan kaçının.// ✅ AGRESİF: Bellek kopyalaması olmadan byte dizisini string'e çevirir.
// DİKKAT: Orijinal byte dizisi sonradan değiştirilmemelidir!
func bytesToString(b []byte) string {
    return \*(\*string)(unsafe.Pointer(&b))
}


- **Reflection'sız Struct Erişimi:** Struct tipleri için ilk başlangıçta (initialization) alanların offset'lerini (`unsafe.Offsetof`) hesaplayarak bir "erişim planı" oluşturun. Sonraki işlemlerde `reflect.Value.Field()` yerine doğrudan `unsafe.Pointer` ve offset aritmetiği ile alanlara erişin. Bu, `go-json` gibi kütüphanelerin kullandığı bir tekniktir.

### 3. Gelişmiş Derleyici İpuçları (Compiler Directives)

`//go:inline` dışında, derleyiciye bellek yönetimi konusunda daha fazla ipucu vererek heap alokasyonlarını engelleyin.

- **`//go:noescape`&#32;Kullanımı:** Bir fonksiyonun argüman olarak aldığı pointer'ların fonksiyon dışına "kaçmadığını" derleyiciye bildirin. Bu, derleyicinin bu pointer'lar için heap alokasyonu yapmasını engeller ve stack üzerinde kalmalarını sağlar.// ✅ AGRESİF: "b" pointer'ının heap'e kaçmasını engeller.
//go:noescape
func getBufferPtr(b \*Buffer) unsafe.Pointer {
    return unsafe.Pointer(&b.buf[0])
}



### 4. Kod Üretimi (Code Generation) Odaklı Yaklaşım

Reflection'ı yavaş yavaş azaltmak yerine, onu tamamen ortadan kaldırmayı hedefleyin.

- **Strateji:** Projenin temel bir parçası olarak, kullanıcıların struct tanımlarını analiz edip onlara özel, reflection kullanmayan `MarshalBEVE` ve `UnmarshalBEVE` metodları üreten bir kod üretici (`go generate`) aracı geliştirin.
- **Çıktı:** Üretilen kod, doğrudan alan erişimi yapar, `type switch`'lere gerek duymaz ve inline edilebilir hale gelir. Bu, `protobuf` veya `cap'n proto` gibi kütüphanelerin performansının arkasındaki anahtar stratejidir.

### 5. SIMD (Single Instruction, Multiple Data) Optimizasyonları

`[]int32`, `[]float64` gibi ilkel tipteki dizilerin toplu kodlanması için SIMD komut setlerinden faydalanın.

- **Strateji:** `golang.org/x/sys/cpu` paketi ile çalışma anında `AVX2` gibi SIMD setlerinin varlığını kontrol edin. Destek varsa, Go'nun içsel (intrinsic) fonksiyonlarını veya assembly'yi kullanarak bir döngüde 4 veya 8 elemanı aynı anda işleyin.
- **Kullanım Alanı:** Büyük veri setleri, bilimsel hesaplama veya makine öğrenmesi modellerinin serileştirilmesi için idealdir.

### 6. Özel Bellek Yönetimi (Arena/Slab Allocators)

`sync.Pool`, genel amaçlı bir çözümdür. Daha da ileri giderek, belirli boyutlardaki nesneler için özel bellek havuzları (allocator) oluşturun.

- **Arena Allocator:** Go 1.20+ ile gelen deneysel `arena` paketini araştırın. Tek bir isteğin ömrü boyunca yapılan tüm alokasyonları tek bir büyük bellek bloğunda toplayarak Garbage Collector (GC) yükünü dramatik şekilde azaltır.
- **Slab Allocator:** Sık sık aynı boyutta (örneğin 1KB'lık buffer'lar) alokasyon yapılıyorsa, bu boyuta özel bir slab allocator tasarlayın. Bu, bellek parçalanmasını (fragmentation) azaltır ve alokasyon/de-alokasyon hızını artırır.

## 📝 Coding Standards

### Naming Conventions

- **Exported functions**: PascalCase (`Marshal`, `Encode`, `GetEncoderFromPool`)
- **Internal functions**: camelCase (`encodeInt`, `writeVarint`, `ensureAddressable`)
- **Constants**: UPPER\_SNAKE\_CASE (`MAX_BUFFER_SIZE`, `TYPE_INT32`)
- **Struct tags**: lowercase with commas (`beve:"name,omitempty"`)

### Function Design

```javascript
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

```javascript
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

**When to use&#32;`unsafe`:**

- ✅ Zero-copy string → []byte conversion (read-only)
- ✅ Struct field access with known offsets
- ✅ Reflection'ı tamamen ortadan kaldırmak için struct field offset'lerini önbelleğe alırken.
- ✅ Performance-critical paths with benchmarks proving >10% improvement

**When NOT to use&#32;`unsafe`:**

- ❌ Pointer arithmetic without proper bounds checking
- ❌ Type conversions that violate Go memory model
- ❌ Any case that causes `checkptr` errors

```javascript
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

## 🧪 Testing Requirements

### Must-Have Tests for New Features

1. **Unit tests**: Basic functionality
2. **Edge case tests**: nil, empty, zero values
3. **Race tests**: `go test -race` must pass
4. **Benchmarks**: Compare with baseline

### Benchmark Standards

```javascript
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

```javascript
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

## 🚀 Common Optimization Patterns

### Pattern 1: Pre-allocated Scratch Buffers

```javascript
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

```javascript
//go:inline
func (e *Encoder) WriteByte(b byte) error {
    if e.Buf != nil {
        return e.Buf.WriteByte(b) // fast path
    }
    return e.w.WriteByte(b) // slow path
}


```

### Pattern 3: Type Switch for Common Cases

```javascript
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

```javascript
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

```javascript
# Heap profile
go test -memprofile=mem.out -bench=.
go tool pprof -alloc_space mem.out

# Look for:
# - Large buffer pooling (>1MB retained)
# - Circular references preventing GC
# - Leaked goroutines holding memory


```

### Race Conditions

```javascript
go test -race ./...

# Common issues:
# - sync.Map vs regular map confusion
# - Encoder/Decoder used across goroutines
# - Shared buffers without synchronization


```

## 📋 Pre-Commit Checklist

Before submitting PR:

- [ ] `go test ./...` passes
- [ ] `go test -race ./...` passes
- [ ] `golangci-lint run` passes (no errors)
- [ ] Benchmarks show no regression (or justified)
- [ ] Documentation updated (GoDoc, README if needed)
- [ ] Commit message follows convention: `<type>(<scope>): <subject>`
- [ ] Added benchmark files if performance-sensitive change

## 🎨 Code Generation Guidelines

### When suggesting struct definitions:

```javascript
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

```javascript
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

```javascript
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

## 🔍 Common Pitfalls to Avoid

### ❌ Don't: Modify pooled objects after returning

```javascript
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

```javascript
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

```javascript
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

```javascript
// BAD: Requires Go 1.20+
func example() {
    m := map[string]int{} 
    clear(m) // only in Go 1.21+
}

// GOOD: Check go.mod version (currently 1.22)
// If using features from 1.22+, ensure go.mod matches


```

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

## 🎯 Project Goals

### Short-term (Current Phase)

- ✅ Achieve <1μs small struct marshal
- ✅ Reduce allocations to <5 per operation
- ⏳ SIMD optimizations for bulk array encoding
- ⏳ Code generation for struct marshaling
- ⏳ Multi-platform benchmarks (Linux, Windows, macOS, ARM)
- ⏳ Improve Windows performance (currently 3× slower)

### Medium-term

- [ ] Streaming API for large payloads
- [ ] WebAssembly support

### Long-term

- [ ] Become fastest binary serialization in Go ecosystem
- [ ] Sub-500ns for small structs
- [ ] Zero-allocation mode for pre-allocated buffers
- [ ] Cross-language specification (BEVE format standard)

## 🏆 Success Criteria

A contribution is successful when:

1. **Tests pass**: `go test -race ./...` ✅
2. **Lint passes**: `golangci-lint run` ✅
3. **Benchmarks stable**: No >5% regression ✅
4. **Documentation complete**: GoDoc + README updates ✅
5. **Performance justified**: Benchmarks prove claims ✅

## 💡 Final Tips

- **Measure, don't guess**: Always benchmark before claiming performance improvements
- **Cache-friendly is fast**: Think about CPU cache lines in struct layout
- **Unsafe is sharp**: Use only when proven necessary and safe
- **Tests are documentation**: Write clear, comprehensive tests
- **Performance is paramount**: Maintainability is managed through rigorous testing and documentation, not by avoiding faster techniques.

**Remember**: This is a performance-critical library. Every nanosecond counts.

*Last updated: October 2025* *Go version: 1.22+* *Maintained by: BEVE-org team*