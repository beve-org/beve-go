# 🧹 Week 1: Code Refactoring - Implementation Plan

## 📊 Current State Analysis

### encoder.go Structure (1,086 lines):
```
Lines 1-88:    Imports, types, Buffer, pool management
Lines 89-150:  encode() - main dispatcher
Lines 151-270: Primitive encoders (null, bool, int, uint, float)
Lines 271-339: String & slice encoders
Lines 340-391: Primitive slice fast path
Lines 392-421: Map encoder
Lines 422-458: Struct encoder
Lines 459-595: Typed array encoders (signed, unsigned, float)
Lines 596-665: Write helpers (writeByte, writeBytes, etc)
Lines 666-1086: Utility functions (varint, compression, etc)
```

### Files to Process:
```
✅ encoder.go          1,086 lines → Split into 4 files
✅ value_pool.go         115 lines → DELETE (unused)
✅ bulk_optimize.go      295 lines → MERGE or DELETE
✅ math_optimize.go      239 lines → Simplify to buffer utils
✅ lockfree_cache.go     244 lines → MERGE into encoder_cache.go
✅ encoder_cache.go      275 lines → Keep & enhance
✅ reflect_optimize.go   468 lines → Move to optimize/
```

---

## 🎯 Day 1: Directory Structure & Buffer Management

### Step 1.1: Create New Structure
```bash
mkdir -p core optimize
```

### Step 1.2: Create core/buffer.go (Buffer management)
**Extract from**: encoder.go (lines 1-88) + math_optimize.go

**Contents**:
- Buffer struct
- Buffer.Write, WriteByte, Grow, Reset
- Buffer pool management
- Growth strategy

**Target**: ~150 lines, clean buffer management

### Step 1.3: Create core/encoder_base.go (Core encoder)
**Extract from**: encoder.go (lines 1-150)

**Contents**:
- encoder struct definition
- encoder pool
- newEncoder()
- encode() dispatcher
- Type checking logic

**Target**: ~200 lines, core logic only

---

## 🎯 Day 2: Type-Specific Encoders

### Step 2.1: Create core/encoder_primitives.go
**Extract from**: encoder.go (lines 151-270)

**Contents**:
```go
// Primitive type encoders
func (e *encoder) encodeNull() error
func (e *encoder) encodeBool(b bool) error
func (e *encoder) encodeInt(i int64) error
func (e *encoder) encodeUint(u uint64) error
func (e *encoder) encodeFloat(f float64, kind reflect.Kind) error
func (e *encoder) encodeString(s string) error
func (e *encoder) encodeBinaryMarshaler(m BinaryMarshaler) error
func (e *encoder) encodeRawMessage(data []byte) error
```

**Target**: ~150 lines

### Step 2.2: Create core/encoder_collections.go
**Extract from**: encoder.go (lines 271-458)

**Contents**:
```go
// Collection encoders
func (e *encoder) encodeSlice(v reflect.Value) error
func (e *encoder) encodePrimitiveSlice(v reflect.Value, kind reflect.Kind) error
func (e *encoder) encodeMap(v reflect.Value) error
func (e *encoder) encodeStruct(v reflect.Value) error
```

**Target**: ~200 lines

### Step 2.3: Create core/encoder_arrays.go
**Extract from**: encoder.go (lines 459-595)

**Contents**:
```go
// Typed array encoders
func (e *encoder) encodeTypedArray(v reflect.Value, info typedArrayInfo) error
func (e *encoder) encodeSignedArray(v reflect.Value, length, byteCount int) error
func (e *encoder) encodeUnsignedArray(v reflect.Value, length, byteCount int) error
func (e *encoder) encodeFloatArray(v reflect.Value, length int, is32bit bool) error
```

**Target**: ~180 lines

---

## 🎯 Day 3: Write Helpers & Utilities

### Step 3.1: Create core/encoder_write.go
**Extract from**: encoder.go (lines 596-665)

**Contents**:
```go
// Low-level write operations
func (e *encoder) writeByte(b byte) error
func (e *encoder) writeBytes(data []byte) error
func (e *encoder) writeStringBytes(s string) error
func (e *encoder) writeIntBytes(value int64, count int) error
func (e *encoder) writeUintBytes(value uint64, count int) error
func (e *encoder) writeCompressedUint(n uint64) error
```

**Target**: ~120 lines

### Step 3.2: Create core/encoder_utils.go
**Extract from**: encoder.go (lines 666-1086)

**Contents**:
```go
// Utility functions
func varintSize(n uint64) int
func isPrimitive(kind reflect.Kind) bool
func getTypedArrayInfo(t reflect.Type) (typedArrayInfo, bool)
// Other helper functions
```

**Target**: ~150 lines

---

## 🎯 Day 4: Optimization Files

### Step 4.1: Move reflect_optimize.go → optimize/reflect.go
**Action**: Move + cleanup

**Changes**:
- Add clear documentation
- Remove unused functions
- Keep only safe reflection optimizations

**Target**: ~350 lines (from 468)

### Step 4.2: Create optimize/unsafe.go
**Extract from**: reflect_optimize.go

**Contents**:
- Unsafe struct field access
- Unsafe value extraction
- Clearly marked as UNSAFE operations

**Target**: ~100 lines

### Step 4.3: Consolidate encoder_cache.go
**Merge**: lockfree_cache.go → encoder_cache.go

**Actions**:
1. Keep lockfree_cache atomic implementation
2. Merge with existing encoder_cache
3. Single unified caching system

**Target**: ~350 lines (from 244 + 275 = 519)

---

## 🎯 Day 5: Cleanup & Documentation

### Step 5.1: Delete Unused Files
```bash
rm value_pool.go        # Not being used effectively
rm bulk_optimize.go     # Functionality covered by typed arrays
rm math_optimize.go     # Merged into core/buffer.go
rm lockfree_cache.go    # Merged into encoder_cache.go
```

### Step 5.2: Update Imports
- Fix all imports in new files
- Ensure all tests still compile
- No functionality changes

### Step 5.3: Add Package Documentation
Create core/doc.go:
```go
// Package core provides the core encoding functionality for BEVE.
//
// This package is organized into logical modules:
//   - buffer.go:           Buffer management and pooling
//   - encoder_base.go:     Core encoder structure and dispatch
//   - encoder_primitives.go: Primitive type encoders
//   - encoder_collections.go: Slice, map, struct encoders
//   - encoder_arrays.go:    Typed array encoders
//   - encoder_write.go:     Low-level write operations
//   - encoder_utils.go:     Helper functions
package core
```

---

## 📊 Expected Results After Day 5

### New Structure:
```
core/
├── doc.go                  (  30 lines) - Package docs
├── buffer.go               ( 150 lines) - Buffer management
├── encoder_base.go         ( 200 lines) - Core encoder
├── encoder_primitives.go   ( 150 lines) - Primitives
├── encoder_collections.go  ( 200 lines) - Collections
├── encoder_arrays.go       ( 180 lines) - Typed arrays
├── encoder_write.go        ( 120 lines) - Write helpers
└── encoder_utils.go        ( 150 lines) - Utilities

optimize/
├── reflect.go              ( 350 lines) - Safe reflection
└── unsafe.go               ( 100 lines) - Unsafe operations

Root files:
├── beve.go                 ( 180 lines) - Public API
├── decoder.go              (1238 lines) - Decoder (keep as-is)
├── encoder_cache.go        ( 350 lines) - Unified caching
├── unsafe.go               (  28 lines) - String conversion
└── rawmessage.go           (  36 lines) - RawMessage type

Test files: (no changes)
```

### Metrics:
```
Before:
  encoder.go:           1,086 lines
  + optimization files:  1,361 lines
  Total:                 2,447 lines

After:
  core/ package:         1,180 lines (split into 8 files)
  optimize/ package:       450 lines (split into 2 files)
  encoder_cache.go:        350 lines (consolidated)
  Total:                 1,980 lines

Reduction: 467 lines (19% smaller!)
Max file size: 350 lines (vs 1,086)
```

---

## ✅ Validation Steps

After each day:
```bash
# 1. Compile check
go build ./...

# 2. Run all tests
go test ./... -v

# 3. Run benchmarks
go test -bench=. -benchmem

# 4. Compare results
# Must match Phase 2 baseline exactly!
```

---

## 🎯 Success Criteria for Week 1

### Must Have:
- ✅ All files < 400 lines
- ✅ Clear separation of concerns
- ✅ All tests pass
- ✅ No performance regression
- ✅ No functionality changes

### Nice to Have:
- ⭐ Improved documentation
- ⭐ Clearer function names
- ⭐ Better code organization

---

## 📝 Implementation Notes

### Rules:
1. **No logic changes** - Pure refactoring only
2. **Test after each step** - Catch issues early
3. **Keep git history clean** - One commit per day
4. **Document as you go** - Add comments for complex parts

### Git Strategy:
```bash
# Day 1
git checkout -b week1-refactoring
git commit -m "Day 1: Create core/ structure, extract buffer management"

# Day 2
git commit -m "Day 2: Extract type-specific encoders"

# Day 3
git commit -m "Day 3: Extract write helpers and utilities"

# Day 4
git commit -m "Day 4: Reorganize optimization files"

# Day 5
git commit -m "Day 5: Cleanup unused files, add documentation"
git push origin week1-refactoring
```

---

## 🚀 Let's Start!

**Ready to begin Day 1?**
- Create core/ and optimize/ directories
- Extract buffer management
- Get everything compiling

**Time estimate**: 2-3 hours for Day 1
