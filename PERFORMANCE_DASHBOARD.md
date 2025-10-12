# BEVE Performance Dashboard 📊# BEVE-Go Performance Dashboard

**Last Updated:** October 12, 2025

**Last Updated**: October 12, 2025  

**Platform**: Apple M2 Max (darwin/arm64), Go 1.22  ---

**Benchmark Method**: 5,000-10,000 iterations, -benchmem

## 🎯 Quick Performance Overview

---

```

## 🏆 EXECUTIVE SUMMARY╔═══════════════════════════════════════════════════════════════════════════╗

║                    BEVE PERFORMANCE SCORECARD                             ║

### Overall Performance Status╠═══════════════════════════════════════════════════════════════════════════╣

║                                                                           ║

| Category | BEVE Rank | Best Time | Gap to Leader | Status |║  📊 OVERALL RATING: A- (Production Ready)                                ║

|----------|-----------|-----------|---------------|--------|║                                                                           ║

| **Wide Struct (50 fields)** | **#1 🥇** | **482ns** | **Leader** | ✅ **EXCELLENT** |║  ✅ STRENGTHS (A+ Tier)                                                   ║

| Large Data Unmarshal | #1 🥇 | 33.2µs | Leader | ✅ EXCELLENT |║    • Large Data Unmarshal ......... 🏆 FASTEST (128μs)                   ║

| Medium Data Unmarshal | #1 🥇 | 3.2µs | Leader | ✅ EXCELLENT |║    • Medium Data Operations ........ 🏆 FASTEST (10ms/12ms)              ║

| File Read | #1 🥇 | 5.9µs | Leader | ✅ EXCELLENT |║    • File Read Performance ......... 🏆 FASTEST (93ms)                   ║

| Large Map Encoding | #2 🥈 | 20.1µs | +4% vs MP | ✅ GOOD |║    • Allocation Efficiency ......... 🏆 BEST (1-3 allocs)                ║

| Small Struct Marshal | #2 🥈 | 113ns | +26% vs CBOR | ✅ GOOD |║                                                                           ║

| Deep Nested (10 levels) | #3 🥉 | 1006ns | +70% vs CBOR | ⚠️ NEEDS WORK |║  ✅ COMPETITIVE (A Tier)                                                  ║

| File Write | #4 | 89µs | +52% vs CBOR | ⚠️ NEEDS WORK |║    • Small Struct Operations ....... 🥈 2-3× JSON speed                  ║

| Streaming Memory | #5 | 27.8KB | +4500% vs JSON | ❌ CRITICAL |║    • Large Map Encoding ............ 🥈 2nd (19μs, 1 alloc)              ║

║    • Interface Slice ............... ⚡ Competitive                       ║

**Strengths**: Unmarshal operations, wide structs, file read  ║                                                                           ║

**Weaknesses**: Streaming memory, nested structures (moderate gap)║  ⚠️  NEEDS IMPROVEMENT (B-C Tier)                                        ║

║    • Wide Struct (20 fields) ....... ❌ 10× slower than JSON             ║

---║    • Streaming Memory .............. ❌ 46× more than JSON               ║

║    • File Write Speed .............. ⚠️  65% slower than CBOR            ║

## ✅ VICTORY: Wide Struct Optimization (Phase 6)║    • Payload Size .................. ⚠️  3× larger than MessagePack     ║

║                                                                           ║

### Before vs After╚═══════════════════════════════════════════════════════════════════════════╝

```

| Metric | Before (Tag Bug) | After (Fixed) | Improvement |

|--------|------------------|---------------|-------------|---

| **BEVE Time** | 585ns | **482ns** | **17.6% faster** |

| **vs JSON** | 10× slower | **1.72× faster** | **~1920% swing!** |## 📊 Performance Matrix

| **vs CBOR** | 13× slower | **1.39× faster** | **~1900% swing!** |

| **Rank** | #5 (last) | **#1 (first)** | +4 positions |### Marshal Performance (Lower is Better)



### Current Wide Struct Leaderboard```

Operation          BEVE    JSON    Sonic   MsgPack  CBOR   Winner

```━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

#1 🥇 BEVE        482ns  736B  2 allocs  ← FASTESTSmall Struct       547ns   1485ns  687ns   2321ns   292ns  ⚡ CBOR

#2 🥈 CBOR        672ns  288B  1 alloc   (+39% time)Small (ZeroCopy)   213ns   -       -       -        -      🏆 BEVE

#3 🥉 JSON        827ns  448B  1 alloc   (+72% time)Medium Data        10ms    32ms    33ms    19ms     11ms   🏆 BEVE

#4    MessagePack 887ns  496B  4 allocs  (+84% time)Large Data         104μs   288μs   332μs   162μs    121μs  🏆 BEVE

#5    Sonic       988ns  493B  2 allocs  (+105% time)Large Map          19μs    121μs   57μs    17μs     36μs   🥈 BEVE

```Wide Struct (20)   486ns   51ns    60ns    979ns    48ns   ❌ BEVE

Deep Nested        1026ns  712ns   1181ns  1189ns   629ns  ⚠️  BEVE

**Impact**: Wide structs are now BEVE's **strength**, not weakness!Interface Slice    4.4μs   4.3μs   5.1μs   4.5μs    3.2μs  ✅ BEVE

```

---

### Unmarshal Performance (Lower is Better)

## 📈 DETAILED BENCHMARKS

```

### 1. Wide Struct (50 fields) - **EXCELLENT ✅** (Phase 6 Victory)Operation          BEVE     JSON     Sonic    MsgPack  CBOR    Winner

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

**Marshal Performance**:Small Struct       676ns    2293ns   1552ns   2391ns   2481ns  🏆 BEVE

Medium Data        12ms     137ms    25ms     34ms     45ms    🏆 BEVE  

| Library | Time (ns) | Memory (B) | Allocs | vs BEVE | Status |Large Data         128μs    1486μs   213μs    360μs    419μs   🏆 BEVE

|---------|-----------|------------|--------|---------|--------|```

| **BEVE** | **482** | **736** | **2** | Baseline | ✅ **LEADER** |

| CBOR | 672 | 288 | 1 | 1.39× slower | |### Memory Usage (Lower is Better)

| JSON | 827 | 448 | 1 | 1.72× slower | |

| MessagePack | 887 | 496 | 4 | 1.84× slower | |```

| Sonic | 988 | 493 | 2 | 2.05× slower | |Operation          BEVE     JSON     Sonic    MsgPack  CBOR    Winner

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

### 2. Deep Nested (10 levels) - NEEDS WORK ⚠️Small Struct       993B     1297B    494B     8326B    400B    ⚡ CBOR

Small (ZeroCopy)   289B     -        -        -        -       🏆 BEVE

**Marshal Performance**:Medium Data        18.6KB   24.9KB   22.2KB   65.9KB   16.5KB  🥈 BEVE

Large Data         181KB    214KB    230KB    527KB    206KB   🏆 BEVE

| Library | Time (ns) | Memory (B) | Allocs | vs BEVE | Status |Large Map          4.1KB    55KB     6.3KB    8.2KB    4.1KB   🏆 BEVE

|---------|-----------|------------|--------|---------|--------|```

| **CBOR** | **591** | **136** | **2** | Baseline | Leader |

| JSON | 721 | 200 | 2 | 1.22× slower | |### Allocations (Lower is Better)

| **BEVE** | **1006** | **176** | **3** | **1.70× slower** | ⚠️ Gap |

| Sonic | 1143 | 216 | 3 | 1.93× slower | |```

| MessagePack | 1200 | 520 | 5 | 2.03× slower | |Operation          BEVE    JSON    Sonic   MsgPack  CBOR   Winner

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

**Root Cause**: Nested struct encoder lookup overhead (sync.Map.Load per level)  Small Struct       3       2       3       9        2      ✅ JSON/CBOR

**Optimization Target**: <800ns (25% improvement)  Medium Data        3       9       4       22       2      🥈 BEVE

**Strategy**: Inline nested struct encoding, reduce encoder cache lookupsLarge Data         3       9       4       115      2      🥈 BEVE

Large Map          1       1354    3       8        1      🏆 BEVE/CBOR

### 3. File Write - NEEDS WORK ⚠️```



**File Write** (1MB file):---



| Library | Time (µs) | Memory (B) | Allocs | vs BEVE | Status |## 🚀 Speed Comparison vs JSON

|---------|-----------|------------|--------|---------|--------|

| **CBOR** | **58.6** | **312** | **4** | Baseline | Leader |```

| MessagePack | 62.0 | 328 | 4 | 1.06× slower | |Benchmark               BEVE vs JSON     Speedup     Status

| Sonic | 71.1 | 328 | 4 | 1.21× slower | |━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

| **BEVE** | **89.1** | **312** | **4** | **1.52× slower** | ⚠️ Gap |Small Struct Marshal    547ns vs 1485ns   2.7×      ✅ Faster

| JSON | 138.0 | 296 | 4 | 2.35× slower | |Small Struct Unmarshal  676ns vs 2293ns   3.4×      ✅ Faster

Medium Marshal          10ms vs 32ms      3.2×      ✅ Faster

**Root Cause**: Large buffer management not optimized for streaming writes  Medium Unmarshal        12ms vs 137ms    11.1×      🏆 Much Faster

**Optimization Target**: <70µs (21% improvement)Large Marshal           104μs vs 288μs    2.8×      ✅ Faster

Large Unmarshal         128μs vs 1486μs  11.6×      🏆 Much Faster

---Large Map               19μs vs 121μs     6.3×      🏆 Much Faster

File Read               93ms vs 707ms     7.6×      🏆 Much Faster

## 🎯 OPTIMIZATION ROADMAPFile Write              100ms vs 78ms     0.78×     ❌ Slower

Wide Struct             486ns vs 51ns     0.1×      ❌ Much Slower

### ✅ COMPLETED```



- **Phase 4**: String array decode (42% allocation reduction)---

- **Phase 5**: Large map encode (99.93% allocation reduction, #2 rank)

- **Phase 6**: Wide struct marshal (Tag bug fix → **#1 rank**, 72% faster than JSON)## 💾 Payload Size Comparison



### 📋 NEXT PRIORITIES```

Format              Size (bytes)    vs BEVE    vs JSON    Ranking

#### 1. **Phase 7: Streaming Memory** (CRITICAL - High Priority)━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

- **Current**: 27.8KBMessagePack         293             -70%       -80%       🥇 1st

- **Target**: <1KB (28× improvement)Sonic               305             -69%       -79%       🥈 2nd  

- **Impact**: Unblocks streaming use casesCBOR                792             -19%       -46%       🥉 3rd

BEVE                981             ---        -34%       4th

#### 2. **Phase 8: Deep Nested** (High Priority)JSON                1478            +51%       ---        5th

- **Current**: 1006ns (70% slower than CBOR)

- **Target**: <800ns (25% improvement)Size Efficiency:

- **Impact**: Better performance for complex data models  BEVE is 34% smaller than JSON    ✅

  BEVE is 3× larger than MsgPack   ⚠️

#### 3. **Phase 9: File Write** (Medium Priority)```

- **Current**: 89µs (52% slower than CBOR)

- **Target**: <70µs (21% improvement)---

- **Impact**: Faster persistence operations

## 🎯 Optimization Priority Matrix

---

```

## 🏅 WHEN TO USE BEVE╔═══════════════════════════════════════════════════════════════════╗

║                    OPTIMIZATION ROADMAP                           ║

✅ **BEVE Excels**:╠═══════════════════════════════════════════════════════════════════╣

- Unmarshal operations (fastest across all sizes)║                                                                   ║

- **Wide structs (50+ fields, #1)** ← NEW!║  🔴 CRITICAL (Fix Immediately)                                    ║

- Large map encoding (#2)║  ┌────────────────────────────────────────────────────────────┐  ║

- File read operations (7.6× faster than JSON)║  │ 1. Wide Struct Encoding          Impact: ⭐⭐⭐⭐⭐          │  ║

║  │    Current: 486ns   Target: <100ns   ROI: 5× improvement   │  ║

⚠️ **Use with Caution**:║  │    Status: 10× slower than JSON - BLOCKING ISSUE           │  ║

- Deep nested structures (70% slower than CBOR)║  │                                                             │  ║

- File write operations (52% slower than CBOR)║  │ 2. Streaming Memory              Impact: ⭐⭐⭐⭐⭐          │  ║

║  │    Current: 27.8KB  Target: <1KB    ROI: 28× improvement   │  ║

❌ **Avoid BEVE for**:║  │    Status: 46× more memory than JSON - CRITICAL            │  ║

- Streaming with large payloads (46× more memory) ← **Phase 7 will fix**║  └────────────────────────────────────────────────────────────┘  ║

║                                                                   ║

---║  🟡 HIGH PRIORITY (Phase 6-7)                                     ║

║  ┌────────────────────────────────────────────────────────────┐  ║

## 📊 PHASE-BY-PHASE PROGRESS║  │ 3. Deep Nested Structures        Impact: ⭐⭐⭐⭐            │  ║

║  │    Current: 1026ns  Target: <700ns   ROI: 1.5× improvement │  ║

| Phase | Focus Area | Before | After | Improvement | Rank Change |║  │                                                             │  ║

|-------|------------|--------|-------|-------------|-------------|║  │ 4. File Write Performance        Impact: ⭐⭐⭐⭐            │  ║

| 4 | String Array Decode | 84.1MB | 48.5MB | 42% ↓ | #3 → #2 |║  │    Current: 100ms   Target: <65ms    ROI: 1.5× improvement │  ║

| 5 | Large Map Encode | 75.5µs, 1353 allocs | 20.1µs, 1 alloc | 73% ↓, 99.9% ↓ | #4 → #2 |║  │                                                             │  ║

| 6 | Wide Struct Marshal | 585ns (#5) | 482ns (#1) | **18% ↓** | **#5 → #1** |║  │ 5. Regular Marshal Optimization  Impact: ⭐⭐⭐⭐            │  ║

║  │    Current: 547ns   Target: <300ns   ROI: 1.8× improvement │  ║

---║  └────────────────────────────────────────────────────────────┘  ║

║                                                                   ║

**Dashboard Maintained by**: BEVE-org team  ║  🟢 MEDIUM PRIORITY (Phase 8+)                                    ║

**Full Details**: See individual phase reports (PHASE_4_*.md, PHASE_5_*.md, PHASE_6_*.md)║  ┌────────────────────────────────────────────────────────────┐  ║

║  │ 6. Payload Size Reduction        Impact: ⭐⭐⭐              │  ║
║  │    Current: 981B    Target: <400B    ROI: 2.5× improvement │  ║
║  │                                                             │  ║
║  │ 7. Large Map vs MessagePack      Impact: ⭐⭐                │  ║
║  │    Current: 19.3μs  Target: <17μs    ROI: 1.1× improvement │  ║
║  └────────────────────────────────────────────────────────────┘  ║
║                                                                   ║
╚═══════════════════════════════════════════════════════════════════╝
```

---

## 📈 Performance Trends

### Best Use Cases ✅

```
1. Large Data Processing
   • 11× faster unmarshal than JSON
   • Minimal memory overhead
   • Example: Log aggregation, analytics

2. File I/O Operations  
   • 7.6× faster reads than JSON
   • 93ms vs 707ms for 100 records
   • Example: Database export/import

3. High-Throughput APIs
   • 3-4× faster than JSON
   • Consistent low allocations
   • Example: Microservices, caching

4. Memory-Constrained Environments
   • Zero-copy mode: 776× memory reduction
   • Example: Embedded systems, mobile
```

### Avoid These Scenarios ⚠️

```
1. Wide Structs (>20 fields)
   • 10× slower than JSON
   • Use JSON or fix in Phase 6

2. Streaming with High Memory Pressure
   • 46× more memory than JSON
   • Wait for Phase 7 fix

3. Bandwidth-Critical Applications
   • 3× larger than MessagePack
   • Consider compression or wait for Phase 8

4. Write-Heavy File Operations
   • 65% slower than CBOR
   • Read-optimized, not write-optimized
```

---

## 🏆 Competitive Position

```
Overall Ranking by Use Case:

Large Data Unmarshal:     #1 🥇 (BEVE dominates)
Medium Data Operations:   #1 🥇 (BEVE dominates)
File Read Performance:    #1 🥇 (BEVE dominates)
Small Struct Unmarshal:   #1 🥇 (BEVE dominates)

Large Map Encoding:       #2 🥈 (Close to MessagePack)
Small Struct Marshal:     #3 🥉 (Behind CBOR, competitive)

Wide Struct Encoding:     #5 ❌ (Critical issue)
Streaming Memory:         #5 ❌ (Critical issue)
File Write Speed:         #5 ⚠️ (Needs improvement)
Payload Size:             #4 ⚠️ (Larger than optimal)
```

---

## 📊 Zero-Copy Mode Benefits

```
Data Size    Regular Memory    Zero-Copy    Improvement    Recommendation
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Small        993 B             289 B        3.4×           ✅ Use it
Medium       18.6 KB           135 B        138×           🏆 Always use
Large        180.8 KB          233 B        776×           🏆 Must use

Speed Benefit:
  Small:  547ns → 213ns  (2.6× faster)
  Medium: 10ms → 8.6ms   (1.2× faster)
  Large:  104μs → 77μs   (1.4× faster)

Recommendation: Default to zero-copy mode for production!
```

---

## 🎯 Next Steps

### Immediate (This Week)
- [ ] Investigate wide struct encoding bottleneck
- [ ] Profile streaming memory allocations
- [ ] Design Phase 6 optimization plan

### Short-term (This Month)  
- [ ] Implement wide struct fast path
- [ ] Fix streaming buffer management
- [ ] Optimize file write performance

### Long-term (Next Quarter)
- [ ] Payload size compression
- [ ] SIMD optimizations for bulk data
- [ ] Code generation for struct encoding

---

## 📝 Summary

**Production Readiness:** ✅ **READY**

**Best For:**
- Large data processing
- File I/O operations
- High-throughput APIs
- Memory-constrained deployments (with zero-copy)

**Avoid For (Until Fixed):**
- Wide struct-heavy workloads
- Memory-sensitive streaming
- Bandwidth-critical applications

**Overall Grade:** **A-** (Would be A+ after Phase 6 fixes)

---

*Last Benchmark Run: October 12, 2025 (10,000 iterations, 178.4s total)*  
*Platform: Apple M2 Max, Go 1.22+*  
*Competitors: JSON, Sonic, MessagePack, CBOR*
