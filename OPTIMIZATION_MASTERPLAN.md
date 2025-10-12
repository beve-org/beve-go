# BEVE Optimization Masterplan
## From Good to #1 in All Categories

**Date**: October 12, 2025  
**Current Status**: Market leader in typical use cases, bottlenecks in edge cases  
**Goal**: Achieve #1 ranking in ALL scenarios

---

## 🔍 Bottleneck Analysis (CPU Profiling Results)

### Critical Findings

#### 1. **Wide Struct Encoding: 10× SLOWER than JSON/CBOR** 🔴 CRITICAL
```
BEVE: 507.6 ns/op    737 B/op    2 allocs/op
JSON:  50.3 ns/op      8 B/op    1 allocs/op  (10× FASTER)
```

**Root Causes Identified:**
- `encodeByKind` overhead: Reflection-based type dispatch
- `buildEncoderStructFieldsRecursive`: Building metadata for 50 fields
- Field key encoding: Each field name gets varint length prefix
- No struct layout optimization

**Profile Hotspots:**
- 40ms (10%) in `(*Encoder).encodeInt` - varint encoding for each field
- Reflection operations dominate CPU time

---

#### 2. **Large Map Encoding: 521 allocations** 🟡 HIGH PRIORITY
```
BEVE:        19,086 ns/op    8,063 B/op    521 allocs/op
MessagePack: 23,474 ns/op   10,328 B/op    527 allocs/op
```

**Actually FASTER than MessagePack! But can be optimized further:**

**Root Causes:**
- One allocation per map entry (521 allocs for 1000 entries - wait, that's only 521, not 1000!)
- Actually not bad - likely pooling working
- But still room for batch allocation

**Optimization Opportunity:** Batch allocate map entries

---

#### 3. **Deep Nested Encoding: 1.7× slower than CBOR** 🟡 MEDIUM PRIORITY
```
BEVE: 774.7 ns/op    145 B/op    3 allocs/op
CBOR: 466.1 ns/op    104 B/op    2 allocs/op  (1.7× FASTER)
```

**Root Causes:**
- Cache lookup for each nested type
- Pointer dereferencing overhead
- No inline optimization for common patterns

---

#### 4. **Interface Slice: 1.4× slower than CBOR** 🟢 LOW PRIORITY
```
BEVE: 3,563 ns/op    408 B/op    2 allocs/op
CBOR: 2,469 ns/op    376 B/op    2 allocs/op  (1.4× FASTER)
```

**Root Causes:**
- `encodeInterfaceValue` has large type switch
- `reflect.ValueOf` for each element
- No homogeneous slice detection

---

## 🎯 Optimization Strategy (Priority Order)

### Phase 4.1: Wide Struct Optimization (Week 1) 🔴 CRITICAL
**Goal**: Match or beat JSON/CBOR performance (target: < 100 ns/op)

#### Option A: Code Generation (Recommended)
**Implementation:**
```go
// Generate specialized encoder for each struct type
//go:generate go run github.com/beve-org/beve-go/cmd/bevegen

// Generated code (example):
func (e *Encoder) encodeWideStructProfile_Fast(v WideStructProfile) error {
    // Pre-computed field count and keys
    e.WriteByte(TYPE_STRUCT | (50 << 3))
    
    // Inline field encoding (no reflection!)
    e.WriteFieldKey(precomputedKey_F1)
    e.encodeInt_Inline(int64(v.F1))
    
    e.WriteFieldKey(precomputedKey_F2)
    e.encodeInt_Inline(int64(v.F2))
    
    // ... repeat for all 50 fields
    
    return nil
}
```

**Benefits:**
- ✅ Eliminates reflection completely
- ✅ Pre-computed field keys (no varint encoding at runtime)
- ✅ Inline field encoding (no function calls)
- ✅ Compiler optimizations (loop unrolling, constant folding)

**Expected Results:**
- Speed: 50-100 ns/op (5-10× faster)
- Memory: < 50 B/op (15× less)
- Allocations: 1 alloc (2× fewer)

**Implementation Steps:**
1. Create `cmd/bevegen` tool
2. Parse struct definitions from Go files
3. Generate `*_beve.go` files with fast encoders
4. Register generated encoders in type cache
5. Fallback to reflection for non-generated types

**Timeline:** 2-3 days
**Complexity:** Medium
**Impact:** 🔥 MASSIVE

---

#### Option B: Struct Layout Optimization (Quick Win)
**Implementation:**
```go
// Pre-compute field keys at cache build time
type encoderStructField struct {
    precomputedKey []byte  // NEW: Store encoded key
    offset         uintptr
    encoder        encoderFunc
    // ...
}

func buildEncoderStructInfo(t reflect.Type) *encoderStructInfo {
    // Build field key once, reuse for all instances
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        name := getFieldName(field)
        
        // Pre-encode the field key
        keyBuf := make([]byte, 0, len(name)+5)
        keyBuf = appendVarint(keyBuf, len(name))
        keyBuf = append(keyBuf, name...)
        
        fields[i].precomputedKey = keyBuf  // Store for reuse
    }
}
```

**Benefits:**
- ✅ No varint encoding per marshal call
- ✅ No string → []byte conversion
- ✅ Cache-friendly access pattern

**Expected Results:**
- Speed: 200-300 ns/op (2× faster)
- Memory: 400 B/op (2× less)
- Allocations: 2 allocs (same)

**Timeline:** 1 day
**Complexity:** Low
**Impact:** 🔥 HIGH (quick win)

---

### Phase 4.2: Map Batch Allocation (Week 1) 🟡 HIGH PRIORITY
**Goal**: Reduce allocations by 50%

**Implementation:**
```go
type mapEntryBatch struct {
    entries []mapEntry
    used    int
}

type mapEntry struct {
    keyBuf   [64]byte  // Inline buffer for small keys
    valueBuf [64]byte  // Inline buffer for small values
}

var mapBatchPool = sync.Pool{
    New: func() interface{} {
        return &mapEntryBatch{
            entries: make([]mapEntry, 100),  // Batch of 100
        }
    },
}

func (e *Encoder) encodeMap(v reflect.Value) error {
    size := v.Len()
    
    if size > 50 {
        // Use batch allocation for large maps
        batch := mapBatchPool.Get().(*mapEntryBatch)
        defer func() {
            batch.used = 0
            mapBatchPool.Put(batch)
        }()
        
        // Encode using batch
        // ...
    } else {
        // Small map: use existing code
    }
}
```

**Expected Results:**
- Speed: 15,000 ns/op (1.3× faster)
- Allocations: 250 allocs (2× fewer)

**Timeline:** 1-2 days
**Complexity:** Medium
**Impact:** 🔥 MEDIUM

---

### Phase 4.3: Inline Nested Struct (Week 2) 🟡 MEDIUM PRIORITY
**Goal**: Match CBOR performance for deep nesting

**Implementation:**
```go
// Detect common nested patterns
func (e *Encoder) encodeStructFast(v reflect.Value) error {
    info := getStructInfo(v.Type())
    
    // Check if struct has nested structs
    if info.hasNestedStructs && info.depth < 5 {
        // Use inline encoding (no cache lookup)
        return e.encodeStructInline(v, info)
    }
    
    // Normal path
    return e.encodeStructNormal(v, info)
}

func (e *Encoder) encodeStructInline(v reflect.Value, info *encoderStructInfo) error {
    // Inline nested struct encoding
    for _, field := range info.fields {
        if field.kind == reflect.Struct {
            // Encode directly without cache lookup
            e.encodeFieldKey(field.key)
            e.encodeStructFieldsInline(v, field)
        } else {
            // Normal field encoding
        }
    }
}
```

**Expected Results:**
- Speed: 500 ns/op (1.5× faster)
- Memory: 100 B/op (1.5× less)

**Timeline:** 2 days
**Complexity:** Medium
**Impact:** 🔥 MEDIUM

---

### Phase 4.4: Homogeneous Interface Slice (Week 2) 🟢 LOW PRIORITY
**Goal**: Match CBOR performance for interface slices

**Implementation:**
```go
func (e *Encoder) encodeInterfaceSlice(v reflect.Value) error {
    size := v.Len()
    
    if size > 10 {
        // Check if homogeneous (all elements same type)
        firstType := v.Index(0).Elem().Type()
        isHomogeneous := true
        
        for i := 1; i < min(size, 5); i++ {  // Sample first 5 elements
            if v.Index(i).Elem().Type() != firstType {
                isHomogeneous = false
                break
            }
        }
        
        if isHomogeneous {
            // Use typed slice encoding (much faster!)
            return e.encodeTypedInterfaceSlice(v, firstType)
        }
    }
    
    // Fallback to element-by-element encoding
    return e.encodeInterfaceSliceGeneric(v)
}
```

**Expected Results:**
- Speed: 2,500 ns/op (1.4× faster)
- Memory: 350 B/op (1.2× less)

**Timeline:** 1 day
**Complexity:** Low
**Impact:** 🔥 LOW

---

### Phase 4.5: SIMD Optimizations (Week 3) 🚀 ADVANCED
**Goal**: Leverage CPU vector instructions for bulk operations

**Implementation:**
```go
// Use SIMD for primitive slice encoding
import "golang.org/x/sys/cpu"

func (e *Encoder) encodePrimitiveSlice(v reflect.Value) error {
    if cpu.X86.HasAVX2 || cpu.ARM64.HasNEON {
        // SIMD path for int32/int64/float64 slices
        return e.encodePrimitiveSliceSIMD(v)
    }
    
    // Fallback
    return e.encodePrimitiveSliceGeneric(v)
}

func (e *Encoder) encodePrimitiveSliceSIMD(v reflect.Value) error {
    // Process 8 int64s at once using SIMD
    // 8× faster for large arrays
}
```

**Expected Results:**
- Speed: 1,200 ns/op (8× faster for 1000-element slice)

**Timeline:** 3-4 days
**Complexity:** High (assembly required)
**Impact:** 🔥 HIGH (for large arrays)

---

## 📊 Expected Final Results (After All Optimizations)

### Wide Struct (50 fields)
```
Before: 507.6 ns/op    737 B/op    2 allocs/op
After:   50.0 ns/op     50 B/op    1 allocs/op
Result: 10× FASTER, 15× LESS MEMORY ✅ BEATS JSON/CBOR
```

### Large Map (1000 entries)
```
Before: 19,086 ns/op    8,063 B/op    521 allocs/op
After:  15,000 ns/op    4,000 B/op    250 allocs/op
Result: 1.3× FASTER, 2× FEWER ALLOCS ✅ BEATS MESSAGEPACK
```

### Deep Nested (10 levels)
```
Before: 774.7 ns/op    145 B/op    3 allocs/op
After:  500.0 ns/op    100 B/op    2 allocs/op
Result: 1.5× FASTER ✅ MATCHES CBOR
```

### Interface Slice (100 elements)
```
Before: 3,563 ns/op    408 B/op    2 allocs/op
After:  2,500 ns/op    350 B/op    2 allocs/op
Result: 1.4× FASTER ✅ MATCHES CBOR
```

### Large Int Slice (1000 elements) - BONUS
```
Before: 9,953 ns/op    3,118 B/op    2 allocs/op
After:  1,200 ns/op    3,000 B/op    2 allocs/op
Result: 8× FASTER with SIMD ✅ MARKET LEADING
```

---

## 🎯 Final Performance Matrix (Post-Optimization)

| Scenario | Current | Target | Status |
|----------|---------|--------|--------|
| **Small structs** | 🥇 1st | 🥇 1st | ✅ Already #1 |
| **Wide structs** | 5th | 🥇 1st | 🔴 10× improvement needed |
| **Large maps** | 3rd | 🥇 1st | 🟡 1.3× improvement |
| **Deep nesting** | 3rd | 🥇 1st | 🟡 1.5× improvement |
| **Interface slices** | 2nd | 🥇 1st | 🟢 1.4× improvement |
| **String-heavy** | 🥈 2nd | 🥇 1st | 🟢 Minor tweaks |
| **Sequential I/O** | 🥇 1st | 🥇 1st | ✅ Already #1 |
| **Large arrays** | ? | 🥇 1st | 🚀 SIMD bonus |

---

## 🛠️ Implementation Timeline

### Week 1: Critical Fixes
- **Day 1-2**: Struct field key pre-computation (quick win)
- **Day 3-5**: Code generation tool (`bevegen`)
- **Day 6-7**: Map batch allocation

### Week 2: Medium Priority
- **Day 8-10**: Inline nested struct encoding
- **Day 11-12**: Homogeneous interface slice detection
- **Day 13-14**: Testing & benchmarking

### Week 3: Advanced Optimizations
- **Day 15-18**: SIMD for primitive slices
- **Day 19-20**: Performance tuning
- **Day 21**: Documentation & blog post

---

## 📝 Success Metrics

### Before Optimization
- **Market Position**: #1 in 2/8 scenarios
- **Average Speedup**: Baseline
- **Edge Case Performance**: 3-10× slower

### After Optimization (Target)
- **Market Position**: #1 in 8/8 scenarios ✅
- **Average Speedup**: 2-10× across all scenarios ✅
- **Edge Case Performance**: Matches or beats all competitors ✅

---

## 🚀 Phase 4.1 Action Items (This Week)

### Immediate (Next 2 hours)
1. ✅ Create profiling benchmarks
2. ✅ Analyze CPU profile
3. ✅ Document bottlenecks
4. ⏳ Implement field key pre-computation

### Short-term (Next 3 days)
1. ⏳ Build code generation tool
2. ⏳ Test generated code
3. ⏳ Benchmark improvements
4. ⏳ Implement map batch allocation

### This Week
1. ⏳ Complete critical optimizations
2. ⏳ Achieve 5× speedup on wide structs
3. ⏳ Update benchmarks
4. ⏳ Write optimization blog post

---

## 💡 Key Insights

### What We Learned
1. **Wide structs**: Reflection + varint overhead kills performance
   - Solution: Code generation eliminates both
   
2. **Maps**: Not as bad as we thought (already competitive with MessagePack)
   - Solution: Batch allocation for marginal gain
   
3. **Deep nesting**: Cache lookup overhead
   - Solution: Inline encoding for common patterns
   
4. **Interface slices**: Type checking every element
   - Solution: Detect homogeneous slices

### Surprising Findings
- ✅ Cache-friendly vs unfriendly struct layout: **NO DIFFERENCE** (146.7 vs 150.0 ns)
  - CPU cache is good enough, no need to optimize layout
  
- ✅ String vs []byte: **String is faster** (124.2 vs 159.1 ns)
  - Keep current string handling
  
- ✅ Encoder reuse: **10% faster** (137.0 vs 150.8 ns)
  - Already documented, working well
  
- ✅ Small struct: **Already faster than Sonic** (92.4 vs 94.5 ns)
  - We're beating the fastest JSON library!

---

## 🔧 Technical Decisions

### Code Generation vs Runtime Optimization
**Decision**: Both!
- Code generation for maximum performance (wide structs)
- Runtime optimization for flexibility (other cases)

### SIMD: Worth It?
**Decision**: Yes, but Phase 3
- 8× speedup for large arrays is significant
- Only for ARM64/AMD64 (90%+ of deployments)
- Graceful fallback for other architectures

### Breaking Changes?
**Decision**: No
- All optimizations internal
- Fully backward compatible
- Users opt-in to code generation

---

## 📚 References

- Profiling results: `profile_results.txt`
- CPU profile: `cpu.prof`
- Benchmark code: `profiling_bench_test.go`
- Weakness analysis: `WEAKNESS_REPORT.md`

---

**Next Step**: Implement Phase 4.1 - Field Key Pre-computation (2 hours) 🚀
