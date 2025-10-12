# 🚀 Assembly Optimization Potential Analysis

**Date:** October 12, 2025  
**Current Performance:** 719 ns/op, 1441 B/op, 3 allocs/op (SmallStruct Marshal)  
**Target:** <500 ns/op with Assembly optimizations

---

## 📊 Hot-Path Analysis (CPU Profiling)

### Top CPU Consumers (Flat Time)
Based on 50,000 benchmark iterations with CPU profiling:

| Function | Flat % | Cum % | Assembly Candidate | Expected Gain |
|----------|--------|-------|-------------------|---------------|
| `(*Buffer).WriteByte` | 7.69% | - | ✅ **HIGH** | 15-25% |
| `reflect.Value.Index` | 7.69% | - | ❌ (runtime) | - |
| `reflect.Value.SetUint` | 7.69% | - | ❌ (runtime) | - |
| `WriteCompressedUint` | ~5% | - | ✅ **MEDIUM** | 10-20% |
| `encodeInt` | ~3% | - | ✅ **MEDIUM** | 10-15% |
| `encodeUint` | ~3% | - | ✅ **MEDIUM** | 10-15% |

**Note:** Most CPU time (46.15%) is spent in runtime/reflection, which cannot be optimized with Assembly.

---

## 🎯 Assembly Optimization Candidates

### 1. ✅ **WriteCompressedUint** - HIGHEST PRIORITY

**Current Implementation:** Pure Go with branch prediction
```go
func (e *Encoder) WriteCompressedUint(n uint64) error {
    if n < 64 {
        return e.WriteByte(byte(n << 2))
    }
    if n < 16384 {
        e.varintScratch[0] = byte(0x01 | ((n >> 8) << 2))
        e.varintScratch[1] = byte(n)
        return e.WriteBytes(e.varintScratch[:2])
    }
    // ... more branches
}
```

**Assembly Potential:**
- **Branch elimination** using conditional moves (CMOV on x86-64)
- **SIMD instructions** for multi-byte encoding (SSE2/AVX)
- **Direct register manipulation** for bit shifting
- **Reduced function call overhead** (inline assembly)

**Expected Performance Gain:**
- **Best case:** 20-30% faster (branch-heavy workloads)
- **Typical case:** 10-20% faster
- **Worst case:** 5-10% faster (already well-optimized Go code)

**Benchmark to Create:**
```go
func BenchmarkWriteCompressedUint_Pure(b *testing.B)    // Current Go
func BenchmarkWriteCompressedUint_Assembly(b *testing.B) // Assembly version
```

**Implementation Complexity:** MEDIUM
- ~50-80 lines of Plan 9 Assembly
- Platform-specific versions (amd64, arm64)
- Requires careful testing for edge cases

---

### 2. ✅ **Buffer.WriteByte** - HIGH PRIORITY

**Current Implementation:**
```go
func (b *Buffer) WriteByte(c byte) error {
    if len(b.data) < cap(b.data) {
        b.data = b.data[:len(b.data)+1]
        b.data[len(b.data)-1] = c
        return nil
    }
    // Grow logic...
}
```

**Assembly Potential:**
- **Bounds check elimination** (already optimized by Go compiler)
- **Direct memory write** with no function call
- **Inline growth path** using assembly

**Expected Performance Gain:**
- **Best case:** 15-25% faster
- **Typical case:** 10-15% faster
- **Caveat:** Go compiler already does excellent job here

**Implementation Complexity:** LOW-MEDIUM
- ~30-40 lines of Assembly
- Simple memory operations
- Platform-specific (amd64, arm64)

---

### 3. ✅ **encodeInt / encodeUint** - MEDIUM PRIORITY

**Current Implementation:** Range checks + bit manipulation
```go
func (e *Encoder) encodeInt(i int64) error {
    var byteCount int
    var byteCountBits byte
    
    if i >= -128 && i <= 127 {
        byteCount = 1
        byteCountBits = 0
    } else if i >= -32768 && i <= 32767 {
        // ... more branches
    }
    // Write to scratch buffer
}
```

**Assembly Potential:**
- **BSR/CLZ instructions** (Bit Scan Reverse / Count Leading Zeros) for range detection
- **Conditional moves** instead of branches
- **Direct scratch buffer manipulation**
- **Batched byte writes**

**Expected Performance Gain:**
- **Best case:** 15-20% faster
- **Typical case:** 10-15% faster

**Implementation Complexity:** MEDIUM-HIGH
- ~80-120 lines per function
- Complex branching logic
- Need to maintain exact semantics

---

### 4. ⚠️ **Bulk Array Encoding (SIMD)** - FUTURE WORK

**Target:** `[]int32`, `[]float64`, `[]uint8` bulk encoding

**Assembly Potential:**
- **AVX2/AVX-512** on x86-64: Process 8 values at once
- **NEON** on ARM64: Process 4 values at once
- **Massive speedup** for large arrays (10-100× faster)

**Expected Performance Gain:**
- **Array encoding:** 500-1000% faster for large arrays
- **SmallStruct impact:** Minimal (no large arrays)

**Implementation Complexity:** HIGH
- Requires SIMD expertise
- CPU feature detection at runtime
- Multiple fallback paths

---

## 📈 Overall Performance Projection

### Conservative Estimate (Assembly for Top 3 Candidates)

**Current Baseline:**
- SmallStruct Marshal: **719 ns/op**
- Breakdown estimate:
  - Reflection: ~350 ns (48.7%) - ❌ Cannot optimize
  - Encoding primitives: ~200 ns (27.8%) - ✅ **Assembly target**
  - Buffer operations: ~100 ns (13.9%) - ✅ **Assembly target**
  - Overhead: ~69 ns (9.6%)

**Assembly Optimization Impact:**

| Component | Current (ns) | Assembly (ns) | Gain |
|-----------|--------------|---------------|------|
| Reflection | 350 | 350 | 0% (not optimizable) |
| Encoding primitives | 200 | **140-160** | 20-30% |
| Buffer operations | 100 | **70-85** | 15-30% |
| Overhead | 69 | 69 | 0% |
| **TOTAL** | **719** | **629-664** | **7.6-12.5%** |

### Aggressive Estimate (Perfect Assembly + SIMD)

With perfect Assembly implementation and future SIMD work:

- **SmallStruct Marshal:** 550-600 ns/op (**15-23% faster**)
- **Large array encoding:** 5-10× faster
- **Memory (B/op):** No change (allocation-bound)

---

## 🎯 Recommended Implementation Strategy

### Phase 1: Proof of Concept (1-2 weeks)
**Goal:** Validate assembly gains with minimal investment

1. **Implement:** `WriteCompressedUint` in Assembly (amd64 only)
2. **Benchmark:** Compare Pure Go vs Assembly
3. **Validate:** Ensure correctness with extensive tests
4. **Decision:** If <10% gain, abort. If >15% gain, continue.

**Files to Create:**
```
core/
  encoder_write_amd64.s     # Assembly implementation
  encoder_write_amd64.go    # Go wrapper
  encoder_write_generic.go  # Fallback (build tag)
  encoder_write_test.go     # Validation tests
  encoder_write_bench_test.go # Performance comparison
```

**Expected Outcome:**
- **Best case:** 20-30% faster WriteCompressedUint → 5-8% overall improvement
- **Worst case:** 5-10% faster → 1-2% overall improvement

### Phase 2: Core Functions (2-4 weeks)
**If Phase 1 successful (>10% gain):**

1. **Implement:**
   - `Buffer.WriteByte` (Assembly)
   - `encodeInt` (Assembly)
   - `encodeUint` (Assembly)

2. **Multi-platform:**
   - amd64 (primary)
   - arm64 (secondary - Apple Silicon)
   - Fallback to Go (other platforms)

**Expected Outcome:**
- **10-15% overall improvement** on SmallStruct Marshal
- **Target: <640 ns/op**

### Phase 3: SIMD Optimizations (Future - 4-8 weeks)
**Long-term research project:**

1. **Bulk array encoding** with AVX2/NEON
2. **Runtime CPU feature detection**
3. **Massive gains for large payloads**

**Expected Outcome:**
- **No impact on SmallStruct** (no arrays)
- **5-10× faster** for `[]int32`, `[]float64` encoding
- **Competitive edge** for scientific/ML workloads

---

## ⚠️ Risks & Considerations

### Technical Risks

1. **Maintenance Burden** (HIGH)
   - Assembly code is harder to read/debug
   - Need expertise for both amd64 and arm64
   - Compiler improvements may make Assembly obsolete

2. **Platform Fragmentation** (MEDIUM)
   - Must maintain fallback Go implementations
   - Testing matrix grows (amd64, arm64, 386, etc.)
   - CI/CD complexity increases

3. **Correctness** (HIGH)
   - Assembly bugs are harder to catch
   - Need extensive fuzzing and property-based testing
   - Edge cases (overflow, negative numbers) critical

### Business Risks

1. **Diminishing Returns** (HIGH)
   - **48% of time is in reflection** - cannot optimize with Assembly
   - Best case: **12-15% overall gain**
   - Effort vs reward ratio may not justify

2. **Code Generation Alternative** (MEDIUM)
   - **Code generation** (like protobuf) eliminates reflection entirely
   - Potentially **2-3× faster** than Assembly optimizations
   - Less maintenance burden than Assembly

---

## 🔬 Benchmarking Plan

### Baseline Benchmarks (Before Assembly)

```bash
# Current performance
go test -bench=BenchmarkSmallStruct_BEVE_Marshal -benchmem -benchtime=50000x

# Detailed profiling
go test -bench=BenchmarkWriteCompressedUint -cpuprofile=cpu.prof -memprofile=mem.prof
go tool pprof -top cpu.prof
```

### Assembly Benchmarks (After Implementation)

```bash
# Compare Pure Go vs Assembly
go test -bench=BenchmarkWriteCompressedUint_Pure -benchmem -benchtime=100000x
go test -bench=BenchmarkWriteCompressedUint_Assembly -benchmem -benchtime=100000x

# End-to-end impact
go test -bench=BenchmarkSmallStruct_BEVE -benchmem -benchtime=50000x
```

### Success Criteria

| Metric | Current | Target (Assembly) | Stretch Goal |
|--------|---------|-------------------|--------------|
| SmallStruct ns/op | 719 | <640 (11% faster) | <600 (16% faster) |
| WriteCompressedUint ns/op | ~5ns | <4ns (20% faster) | <3.5ns (30% faster) |
| Code maintainability | ✅ Good | ⚠️ Fair | ❌ Poor |

**Abort Conditions:**
- If Phase 1 shows <8% improvement → Stop
- If maintenance burden exceeds 2× current → Stop
- If code generation proves superior → Pivot

---

## 💡 Alternative: Code Generation (Recommended First)

**Instead of Assembly, consider:**

### Approach: Type-Specific Marshaler Generation

```go
//go:generate beve-gen -type=User -output=user_beve.go

type User struct {
    ID   int    `beve:"id"`
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

// Generated code (no reflection!)
func (u *User) MarshalBEVE() ([]byte, error) {
    enc := core.GetEncoderFromPool()
    defer core.PutEncoderToPool(enc)
    
    enc.Buf.Reset()
    enc.encodeInt(u.ID)      // Direct call, no reflect.ValueOf
    enc.EncodeString(u.Name) // Direct call
    enc.encodeInt(u.Age)     // Direct call
    
    return enc.Buf.Bytes(), nil
}
```

**Benefits:**
- **2-3× faster** than reflection (eliminates 48% overhead!)
- **Easier to maintain** than Assembly
- **Type-safe** - compiler checks
- **Cross-platform** - no Assembly needed

**Implementation Effort:**
- **2-3 weeks** for basic code generator
- **Similar to protobuf/msgpack codegen**

---

## 📊 Final Recommendation

### Priority Order (ROI-based)

1. **🥇 Code Generation** (Highest ROI)
   - **Effort:** 2-3 weeks
   - **Gain:** 2-3× faster (200-240% improvement!)
   - **Maintenance:** LOW
   - **Risk:** LOW

2. **🥈 Assembly (WriteCompressedUint only)** (Medium ROI)
   - **Effort:** 1-2 weeks
   - **Gain:** 10-15% faster overall
   - **Maintenance:** MEDIUM
   - **Risk:** MEDIUM

3. **🥉 Full Assembly Suite** (Low ROI)
   - **Effort:** 4-8 weeks
   - **Gain:** 15-20% faster overall
   - **Maintenance:** HIGH
   - **Risk:** HIGH

### Suggested Action Plan

**Week 1-3:** Implement **Code Generation** for structs
- Bigger impact than Assembly
- Eliminates reflection overhead (48% of total time!)
- Easier to maintain

**Week 4-5:** If still needed, implement **Assembly for WriteCompressedUint**
- Only if benchmarks show >15% gain potential
- Focus on amd64 (most common platform)

**Week 6+:** Evaluate SIMD for bulk arrays (future work)

---

## 🎓 Educational Resources

### Learning Plan 9 Assembly (Go)

1. **Official Documentation:**
   - [A Quick Guide to Go's Assembler](https://go.dev/doc/asm)
   - [The Design of the Go Assembler](https://www.youtube.com/watch?v=KINIAgRpkDA)

2. **Real-world Examples:**
   - `crypto/aes` - AES-NI instructions
   - `math/bits` - Bit manipulation
   - `encoding/json` - String escaping (older versions)
   - `go-json` by goccy - Modern JSON with Assembly

3. **Tools:**
   ```bash
   # Disassemble Go code to see compiler output
   go tool compile -S encoder_write.go > encoder_write.s
   
   # Inspect generated assembly
   go build -gcflags="-S" 2>&1 | less
   ```

---

## 📝 Conclusion

**Assembly can provide 10-20% performance gains** for BEVE-Go, but **Code Generation offers 200-300% gains** with less effort.

**Recommendation:** Start with **Code Generation**, then evaluate Assembly if sub-500ns/op is critical.

**Current Performance (719 ns/op) is already excellent** - only 20% slower than JSON with 64% smaller payloads.

---

**Next Steps:**
1. ✅ Review this analysis
2. ⏳ Decide: Code Generation vs Assembly
3. ⏳ Create POC for chosen approach
4. ⏳ Benchmark and validate
5. ⏳ Implement if ROI is positive

