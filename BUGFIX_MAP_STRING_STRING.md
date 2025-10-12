# Bug Fix: map[string]string Encoding/Decoding Failure

**Date:** October 12, 2025  
**Type:** Pre-existing bug fix  
**Severity:** High (test failure)  
**Status:** ✅ FIXED

---

## 🐛 Problem

**Test Failure:**
```bash
=== RUN   TestDecodeMaps/map_string_string
    decoder_test.go:169: Decode failed: beve: expected slice or array
--- FAIL: TestDecodeMaps/map_string_string (0.00s)
```

**Impact:**
- `map[string]string` encoding/decoding was completely broken
- Pre-existing bug (existed before Phase 5 map[string]interface{} optimization)
- Other map types (map[string]int, map[string]interface{}) worked correctly

---

## 🔍 Root Cause Analysis

### Encoding Side

**File:** `core/encoder_map_zero_alloc.go`  
**Function:** `encodeMapStringString` (line 86)

**Incorrect Code:**
```go
for k, v := range m {
    // Encode key (raw: length + bytes)
    if err := e.WriteCompressedUint(uint64(len(k))); err != nil {
        return err
    }
    if err := e.WriteStringBytes(k); err != nil {
        return err
    }
    
    // ❌ BUG: Encode value WITHOUT header (only length + bytes)
    if err := e.WriteCompressedUint(uint64(len(v))); err != nil {
        return err
    }
    if err := e.WriteStringBytes(v); err != nil {  // ❌ Missing header!
        return err
    }
}
```

**What it produced:**
```
Hex: 03040c6b65791476616c7565
     ││││├─────┤├┴────────┤
     ││││key    │ value (NO HEADER!)
     ││││       └─ 14 (length 5)
     │││└─ "key" (6b 65 79)
     ││└─ 0c (length 3)
     │└─ 04 (map size 1)
     └─ 03 (map header: type 3, keyType 0)
```

### Decoding Side

**File:** `core/decoder_collections.go`  
**Function:** `buildMapValueDecoder` (line 638)

**Decoder expectation:**
```go
case reflect.String:
    return func(d *Decoder, value reflect.Value) error {
        header := d.Data[d.Pos]
        typeBits := header & 0x07
        
        if typeBits != 2 {  // ✅ Expects type 2 (string header 0x02)
            return d.Decode(value)  // Falls back to full decode
        }
        // Read string...
    }
```

**What happened:**
1. Encoder wrote: `14 76 61 6c 75 65` (length + bytes, NO header)
2. Decoder read first byte: `0x14` (20 decimal)
3. Type bits: `0x14 & 0x07 = 0x04` (type 4 = typed array!)
4. Decoder called `DecodeTypedArray` instead of string decoder
5. Error: "expected slice or array"

---

## ✅ Solution

**Changed:** Value encoding in `encodeMapStringString`

**File:** `core/encoder_map_zero_alloc.go`

**Before:**
```go
// Encode value (WRONG - no header)
if err := e.WriteCompressedUint(uint64(len(v))); err != nil {
    return err
}
if err := e.WriteStringBytes(v); err != nil {
    return err
}
```

**After:**
```go
// Encode value (CORRECT - full BEVE encoding with header)
if err := e.EncodeString(v); err != nil {
    return err
}
```

**What `EncodeString` does:**
```go
func (e *Encoder) EncodeString(s string) error {
    header := byte(0x02)  // ✅ String type header
    if err := e.WriteByte(header); err != nil {
        return err
    }
    
    // Write size
    if err := e.WriteCompressedUint(uint64(len(s))); err != nil {
        return err
    }
    
    // Write bytes
    return e.WriteStringBytes(s)
}
```

**New encoding:**
```
Hex: 03040c6b6579021476616c7565
     ││││├─────┤│├┴────────┤
     ││││key    ││value (WITH HEADER!)
     ││││       │└─ 14 (length 5)
     ││││       └─ 02 ✅ (string header)
     │││└─ "key"
     ││└─ 0c (length 3)
     │└─ 04 (map size 1)
     └─ 03 (map header)
```

---

## 🧪 Verification

### Test Results

**Before fix:**
```bash
=== RUN   TestDecodeMaps/map_string_string
    decoder_test.go:169: Decode failed: beve: expected slice or array
--- FAIL: TestDecodeMaps/map_string_string (0.00s)
```

**After fix:**
```bash
=== RUN   TestDecodeMaps
=== RUN   TestDecodeMaps/map_string_int
=== RUN   TestDecodeMaps/map_string_string
=== RUN   TestDecodeMaps/map_int_string
=== RUN   TestDecodeMaps/map_empty
--- PASS: TestDecodeMaps (0.00s)
    --- PASS: TestDecodeMaps/map_string_int (0.00s)
    --- PASS: TestDecodeMaps/map_string_string (0.00s)  ✅
    --- PASS: TestDecodeMaps/map_int_string (0.00s)
    --- PASS: TestDecodeMaps/map_empty (0.00s)
PASS
```

### Full Test Suite

```bash
$ go test ./...
ok      github.com/beve-org/beve-go        0.376s
ok      github.com/beve-org/beve-go/core   0.195s
```

**Result:** ✅ All tests passing (except unrelated assembly-poc build failure)

---

## 📊 Performance Impact

**Phase 5 optimization (map[string]interface{}) preserved:**

| Metric | Before Fix | After Fix | Change |
|--------|-----------|----------|---------|
| **Time** | 20,084 ns/op | 20,075 ns/op | -9 ns (0.04% improvement) |
| **Memory** | 4,105 B/op | 4,107 B/op | +2 bytes (0.05% increase) |
| **Allocations** | 1 allocs/op | 1 allocs/op | ✅ No change |

**Analysis:**
- ✅ Phase 5 optimization still works perfectly
- ✅ 99.93% allocation reduction maintained (1353 → 1)
- ✅ Performance virtually identical (±0.05%)
- ✅ +2 bytes is acceptable (string header overhead)

---

## 🎓 Why This Happened

### Original Design Intent

The fast path optimizations in `encoder_map_zero_alloc.go` were designed to **skip redundant encoding** for primitive map values:

**Theory (incorrect assumption):**
- "Map values don't need headers since decoder knows the expected type"
- "We can save bytes by writing raw data"

**Reality (correct BEVE protocol):**
- Map values MUST be full BEVE-encoded
- Decoder uses generic value decoding (doesn't know type ahead of time)
- Even in maps, each value needs its type header

### Comparison with Other Types

**map[string]int** - Works correctly:
```go
if err := e.encodeInt(int64(v)); err != nil {  // ✅ encodeInt writes header
    return err
}
```

**map[string]string** - Was broken:
```go
if err := e.WriteCompressedUint(uint64(len(v))); err != nil {  // ❌ No header
    return err
}
if err := e.WriteStringBytes(v); err != nil {  // ❌ Just bytes
    return err
}
```

**Key difference:** `encodeInt` includes header logic, but the code was manually writing length + bytes for strings.

---

## 🔧 Similar Issues Checked

Verified that other specialized map encoders DON'T have this bug:

### ✅ map[string]int (encoder_map_zero_alloc.go line 70)
```go
if err := e.encodeInt(int64(v)); err != nil {  // ✅ Uses encodeInt (includes header)
    return err
}
```

### ✅ map[string]float64 (encoder_map_zero_alloc.go line 141)
```go
if err := e.encodeFloat(v, reflect.Float64); err != nil {  // ✅ Uses encodeFloat (includes header)
    return err
}
```

### ✅ map[string]bool (encoder_map_zero_alloc.go line 188)
```go
if err := e.encodeBool(v); err != nil {  // ✅ Uses encodeBool (includes header)
    return err
}
```

### ✅ map[string]interface{} (encoder_collections.go line 494)
```go
if err := encodeInterfaceValue(e, v); err != nil {  // ✅ Full encoding with headers
    return err
}
```

**Conclusion:** Only `map[string]string` had this bug. All other map types correctly encode values with headers.

---

## 🎯 Lessons Learned

### 1. Protocol Consistency is Critical

**Principle:** All BEVE values must be fully encoded with headers, regardless of context.

**Why:**
- Decoder doesn't always know the expected type
- Self-describing format is more robust
- Enables forward compatibility

### 2. Test Coverage Gaps

**How this slipped through:**
- `map[string]string` test existed but was already failing
- Pre-existing failure was not blocking development
- Need better CI to fail-fast on ANY test failure

**Action:** Ensure CI blocks all merges if ANY test fails.

### 3. Fast Path Pitfalls

**Optimization trap:**
- "Let's skip the header to save bytes!"
- Works for encoding, breaks decoding
- Premature optimization without full protocol understanding

**Correct approach:**
- Optimize WITHIN the protocol constraints
- Don't break encoding/decoding symmetry
- Always verify roundtrip tests

### 4. Manual vs Helper Functions

**Pattern:**
```go
// ❌ BAD: Manual encoding (easy to forget headers)
e.WriteCompressedUint(len(s))
e.WriteStringBytes(s)

// ✅ GOOD: Use encoding helpers (headers included)
e.EncodeString(s)
```

**Guideline:** Always use high-level encoding functions for map values.

---

## 📝 Related Files

**Modified:**
- `core/encoder_map_zero_alloc.go` - Fixed `encodeMapStringString` (line 108-110)

**Tested:**
- `core/decoder_test.go` - `TestDecodeMaps` now passes
- All map-related tests verified

**Related Optimizations:**
- Phase 5 `map[string]interface{}` optimization (separate, working correctly)

---

## ✅ Completion Checklist

- [x] Bug identified and root cause analyzed
- [x] Fix implemented (use `EncodeString` instead of manual encoding)
- [x] All tests passing (`TestDecodeMaps` ✓)
- [x] Performance validated (no regression)
- [x] Phase 5 optimization preserved (99.93% allocation reduction maintained)
- [x] Similar issues checked (no other map types affected)
- [x] Documentation created (this file)

**Status:** ✅ **BUG FIXED AND VERIFIED**

---

*Last updated: October 12, 2025*  
*Fixed by: BEVE Development Team*  
*Go version: 1.22+ (darwin/arm64)*
