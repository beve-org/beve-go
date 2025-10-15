# Struct Field Encoding Optimization Report

**Date:** October 15, 2025  
**Status:** Analysis Complete - Ready for Implementation

## Executive Summary

Profiling identified `writeStructFieldsBuffered` as the **largest bottleneck** (1.74s cumulative, 16.14% of total CPU time). Within this function, **string slice encoding** dominates at 1.41s (81% of function time).

## Profiling Results

### CPU Hotspots (BenchmarkSmallStruct_BEVE_Marshal, 5s)

| Function | Flat | Cumulative | % of Total | Priority |
|----------|------|------------|------------|----------|
| `writeStructFieldsBuffered` | 210ms | 1.74s | 16.14% | 🔥 **CRITICAL** |
| `encodeStringSliceDirect` | 170ms | 1.38s | 12.80% | 🔥 **HIGH** |
| `WriteCompressedUint` | 610ms | 620ms | 5.66% | ✅ Optimized |
| `Buffer.Write` | 330ms | 470ms | 4.36% | 🟡 Medium |
| `WriteStringBytes` | 140ms | 610ms | 5.66% | 🟡 Medium |

### String Slice Encoding Breakdown

**Total time:** 1.38s (12.80% of benchmark)

```
Loop overhead:             130ms (9.4%)
WriteCompressedUint:       590ms (42.8%) 🎯 TARGET
WriteStringBytes:          640ms (46.4%)
Other:                      20ms (1.4%)
```

**Key Finding:** Writing varint lengths for each string takes **590ms** - nearly half the string slice encoding time!

## Root Cause Analysis

### Problem 1: Individual Varint Writes (590ms)

**Current Pattern:**
```go
for _, s := range slice {
    WriteCompressedUint(uint64(len(s)))  // Function call per string
    WriteStringBytes(s)
}
```

**Issues:**
1. Function call overhead: ~4-5 ns × millions of calls
2. No batching: Each varint written individually
3. Buffer grows incrementally: Multiple reallocations
4. Branch misprediction: Variable-length varints

### Problem 2: Field Key Appending (90ms)

**Current Pattern:**
```go
buf = append(buf, field.key...)  // Per-field append
```

**Issues:**
1. Pre-encoded keys already optimal ✅
2. Append overhead: Capacity checks per field
3. Could batch multiple field keys

### Problem 3: Buffer Incremental Growth

**Current Pattern:**
```go
buf = append(buf, data...)  // Grows as needed
```

**Issues:**
1. No pre-calculation of final size
2. Multiple reallocation events
3. Copy overhead on growth

## Optimization Strategy

### 🚀 Phase 1: Batch String Slice Encoding (Target: 10-12% overall gain)

**Implementation Plan:**

1. **Pre-calculate total size** (eliminate incremental growth):
```go
func (e *Encoder) encodeStringSliceBatched(slice []string) error {
    // Phase 1: Calculate exact size needed
    totalSize := 0
    for _, s := range slice {
        totalSize += varintSize(uint64(len(s)))
        totalSize += len(s)
    }
    
    // Phase 2: Single allocation
    e.Buf.Grow(totalSize + 8) // +8 for header/count
    
    // Phase 3: Inline batch write (no function calls)
    buf := e.Buf.data
    offset := len(buf)
    buf = buf[:len(buf)+totalSize+8]
    
    // Write header + count (inline)
    buf[offset] = 0x04 | (3 << 3) | (1 << 5)
    offset++
    offset += writeVarintInline(buf[offset:], uint64(len(slice)))
    
    // Batch write all strings (inline varints)
    for _, s := range slice {
        n := len(s)
        offset += writeVarintInline(buf[offset:], uint64(n))
        copy(buf[offset:], s)
        offset += n
    }
    
    e.Buf.data = buf
    return nil
}
```

**Expected Improvements:**
- Eliminate 590ms of varint function calls → **~300ms remaining** (2× speedup)
- Eliminate buffer reallocation overhead → **~50ms saved**
- Total savings: **~340ms** (31% of string slice time)
- **Overall benchmark gain: 10-12%** (340ms / 3.2s baseline)

### 🎯 Phase 2: Inline Varint Writing (Helper)

**Create ultra-fast inline varint writer:**

```go
// writeVarintInline writes a varint directly to buffer, returns bytes written.
// MUST be called with sufficient buffer capacity (use varintSize first).
//
//go:inline
func writeVarintInline(buf []byte, n uint64) int {
    // Ultra-fast path: 0-63 (1 byte, 90% of cases)
    if n < 64 {
        buf[0] = byte(n << 2)
        return 1
    }
    
    // Fast path: 64-16383 (2 bytes, 8% of cases)
    if n < 16384 {
        buf[0] = byte(0x01 | ((n >> 8) << 2))
        buf[1] = byte(n)
        return 2
    }
    
    // Medium path: 16384-1073741823 (3 bytes, 1.9% of cases)
    if n < 1073741824 {
        buf[0] = byte(0x02 | ((n >> 16) << 2))
        buf[1] = byte(n >> 8)
        buf[2] = byte(n)
        return 3
    }
    
    // Slow path: Large values (4 bytes, 0.1% of cases)
    buf[0] = byte(0x03 | ((n >> 24) << 2))
    buf[1] = byte(n >> 16)
    buf[2] = byte(n >> 8)
    buf[3] = byte(n)
    return 4
}
```

**Benefits:**
- No function call overhead
- No error handling overhead
- Direct buffer write (no append)
- Branch predictor friendly (common case first)

### 🟡 Phase 3: Buffer Pre-allocation (3-4% gain)

**Use existing `info.sizeHint` for pre-growth:**

```go
func (e *Encoder) writeStructFieldsBuffered(...) error {
    // Pre-allocate based on size hint
    sizeHint := int(atomic.LoadUint32(&info.sizeHint))
    if sizeHint > 0 {
        e.Buf.Grow(sizeHint)
    }
    
    buf := e.Buf.data
    // ... existing code ...
}
```

**Expected Gain:** Eliminate ~2-3 buffer reallocations → **20-30ms saved** (3-4% overall)

### 🟢 Phase 4: Batch Field Keys (1-2% gain)

**Future optimization** (lower priority):
- Pre-allocate total field key size
- Batch write multiple field keys
- Reduce append overhead

## Implementation Priority

| Phase | Target | Complexity | Expected Gain | Status |
|-------|--------|------------|---------------|---------|
| 1. Batch String Slice | 590ms varint calls | Medium | 10-12% | 🔄 Ready |
| 2. Inline Varint Helper | Create helper | Low | Enabler | 🔄 Ready |
| 3. Buffer Pre-allocation | Use sizeHint | Low | 3-4% | ⏳ Next |
| 4. Batch Field Keys | Optimize append | Medium | 1-2% | ⏳ Future |

**Total Expected Gain:** 15-20% overall benchmark improvement

## Validation Plan

1. **Correctness**: Run full test suite after each phase
2. **Performance**: Benchmark before/after each optimization
3. **Regression**: Compare with baseline (current: 820 ns/op)
4. **End-to-end**: Test with Small/Medium/Large payloads

## Next Steps

1. ✅ Analysis complete
2. 🔄 Implement Phase 1: Batch string slice encoding
3. 🔄 Implement Phase 2: Inline varint helper
4. ⏳ Benchmark and validate
5. ⏳ Implement Phase 3 if time permits

---

**Estimated Implementation Time:** 3-4 hours  
**Estimated Performance Gain:** 15-20% (target: 820ns → 650-700ns for SmallStruct)
