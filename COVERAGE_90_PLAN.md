# 90% Coverage Action Plan

**Current Status**: 54.9% total (beve-go: 74.6%, core: 54.4%)  
**Target**: 90% total coverage  
**Gap**: +35.1 percentage points needed

---

## 📊 Coverage Analysis

### By File (Sorted by Priority)
```
HIGH PRIORITY (0% - Need immediate attention):
- value_pool.go: 0% (9 functions) - Arena/pool infrastructure
- encoder_collections.go: 68.6% (31 0% functions) - Struct fast paths
- decoder_collections.go: 54.2% (6 0% functions) - Struct/array decoders

MEDIUM PRIORITY (50-80% - Need targeted tests):
- encoder_utils.go: 60% (5 0% functions)
- decoder_utils.go: 51.3% (2 0% functions)
- decoder_base.go: 57.6% (5 0% functions)
- beve.go: 80% (11 0% functions)

LOW PRIORITY (80%+ - Minor gaps):
- buffer.go: 82.9% (8 0% functions)
- encoder_base.go: 90% (3 0% functions)
- decoder_primitives.go: 85.1% (1 0% function)
```

---

## 🎯 Critical 0% Functions (Must Cover)

### Tier 1: Performance-Critical Paths ⭐⭐⭐⭐⭐
**Impact**: These are hot paths used in every encoding/decoding operation

1. **encoder_collections.go** (31 functions @ 0%)
   - `encodeStructFast()` - Fast struct encoding (bypasses reflection)
   - `countStructFields()` - Field counting optimization
   - `writeStructFields()` - Direct field serialization
   - `isStructFieldEmpty()` - omitempty handling
   - `encodeStructFieldValue()` - Per-field encoding
   - `buildStructFieldKey()` - Key generation
   
   **Test Strategy**:
   ```go
   // Test struct fast path with various field types
   type FastStruct struct {
       Name   string
       Age    int
       Active bool
       Tags   []string `beve:"tags,omitempty"`
   }
   ```

2. **value_pool.go** (9 functions @ 0%)
   - `Get()`, `Put()` for IntPool, StringPool, SlicePool
   - `getArena()`, `putArena()` - Memory arena management
   - `max()` - Helper function
   
   **Test Strategy**:
   ```go
   // Trigger pool usage through intensive encode/decode cycles
   for i := 0; i < 1000; i++ {
       Marshal(complexData)
       Unmarshal(data, &result)
   }
   ```

3. **decoder_collections.go** (6 functions @ 0%)
   - `decodeStructSlow()` - Fallback struct decoder
   - `parseTag()` - Field tag parsing
   
   **Test Strategy**:
   ```go
   // Force slow path with unaddressable structs
   type NestedStruct struct {
       Field1 string `beve:"f1,omitempty"`
       Field2 int    `beve:"f2"`
   }
   ```

### Tier 2: Error Paths & Edge Cases ⭐⭐⭐⭐
**Impact**: Robustness and error handling

4. **decoder_base.go** (5 functions @ 0%)
   - `timeFromUnixNano()` - time.Time conversion (already added, needs test)
   - `captureRawValue()` - RawMessage capture
   - `setRawMessageValue()` - RawMessage setting
   
   **Test Strategy**:
   ```go
   // Test RawMessage edge cases
   type MessageWithRaw struct {
       Type string
       Data RawMessage
   }
   ```

5. **beve.go** (11 functions @ 0%)
   - `Error()` on InvalidUnmarshalError
   - `Indent()`, `SetEscapeHTML()` - No-op JSON compatibility
   
   **Test Strategy**:
   ```go
   // Test error paths
   var notPointer int
   err := Unmarshal(data, notPointer) // Should error
   errMsg := err.Error() // Covers Error()
   ```

6. **buffer.go** (8 functions @ 0%)
   - `Release()` on BufferLease
   - Buffer growth edge cases
   
   **Test Strategy**:
   ```go
   // Test buffer lifecycle
   lease := MarshalZeroCopy(data)
   bytes := lease.Bytes()
   lease.Release() // Must be called
   ```

### Tier 3: Low-Coverage Functions (Below 50%) ⭐⭐⭐
**Impact**: Partial coverage, need branch completion

7. **decoder_collections.go** (Low coverage functions)
   - `decodeSignedTypedArray()`: 16.2%
   - `decodeUnsignedTypedArray()`: 15.4%
   - `decodeFloatTypedArray()`: 25.4%
   - `decodeStringTypedArray()`: 27.5%
   - `decodeBoolTypedArray()`: 34.6%
   
   **Test Strategy**:
   ```go
   // Test all typed array branches
   testArrays := []interface{}{
       []int8{-128, 0, 127},
       []uint64{0, math.MaxUint64},
       []float32{math.NaN(), math.Inf(1)},
       []string{"", "test", "unicode:你好"},
       []bool{true, false, true},
   }
   ```

8. **encoder_collections.go** (Low coverage functions)
   - `encodeInterfaceValue()`: 15.0%
   - `buildSliceEncoder()`: 24.4%
   - `buildMapEncoder()`: 31.9%
   - `encodePrimitiveSlice()`: 38.5%
   
   **Test Strategy**:
   ```go
   // Test interface{} encoding with various types
   data := map[string]interface{}{
       "int":    42,
       "string": "test",
       "slice":  []int{1, 2, 3},
       "map":    map[string]int{"a": 1},
       "time":   time.Now(),
   }
   ```

9. **common.go**
   - `getEncoderFunc()`: 37.1%
   
   **Test Strategy**:
   ```go
   // Test encoder function caching for custom types
   type CustomType struct{ Value string }
   func (c CustomType) MarshalBEVE() ([]byte, error) {
       return []byte(c.Value), nil
   }
   ```

---

## 📝 Test File Plan

### Wave 5: Performance-Critical Paths (~+20% coverage)
**File**: `core/performance_paths_test.go`

```go
// Test struct fast path
func TestEncodeStructFastPath(t *testing.T)
func TestCountStructFields(t *testing.T)
func TestWriteStructFields(t *testing.T)
func TestIsStructFieldEmptyAllTypes(t *testing.T)
func TestEncodeStructFieldValueAllKinds(t *testing.T)
func TestBuildStructFieldKeyVariations(t *testing.T)

// Test value pools
func TestValuePoolIntensiveUsage(t *testing.T)
func TestArenaPoolingMemory(t *testing.T)
func TestPoolConcurrentAccess(t *testing.T)

// Test decoder slow paths
func TestDecodeStructSlowPathTriggers(t *testing.T)
func TestParseTagAllOptions(t *testing.T)
```

**Expected Impact**: +15-20% coverage (targets 31 0% functions)

---

### Wave 6: Error Paths & Edge Cases (~+10% coverage)
**File**: `error_paths_test.go`

```go
// Test error messages
func TestInvalidUnmarshalErrorMessages(t *testing.T)
func TestUnsupportedTypeErrors(t *testing.T)

// Test JSON compatibility stubs
func TestIndentNoOp(t *testing.T)
func TestSetEscapeHTMLNoOp(t *testing.T)

// Test buffer lifecycle
func TestBufferLeaseReleaseFlow(t *testing.T)
func TestBufferGrowthEdgeCases(t *testing.T)

// Test RawMessage paths
func TestCaptureRawValueAllTypes(t *testing.T)
func TestSetRawMessageValueNil(t *testing.T)

// Test time.Time edge cases
func TestTimeFromUnixNanoCoverage(t *testing.T)
```

**Expected Impact**: +8-10% coverage (error branches and no-op functions)

---

### Wave 7: Typed Arrays Branch Coverage (~+5% coverage)
**File**: `core/typed_arrays_complete_test.go`

```go
// Complete all typed array branches
func TestDecodeSignedTypedArrayAllSizes(t *testing.T)   // 16.2% → 100%
func TestDecodeUnsignedTypedArrayAllSizes(t *testing.T) // 15.4% → 100%
func TestDecodeFloatTypedArrayEdgeCases(t *testing.T)   // 25.4% → 100%
func TestDecodeStringTypedArrayUnicode(t *testing.T)    // 27.5% → 100%
func TestDecodeBoolTypedArrayConversions(t *testing.T)  // 34.6% → 100%

// Test element setters
func TestSetSignedElementAllKinds(t *testing.T)   // 33.3% → 100%
func TestSetUnsignedElementAllKinds(t *testing.T) // 33.3% → 100%
func TestSetFloatElementAllKinds(t *testing.T)    // 42.9% → 100%
func TestSetBoolElementAllKinds(t *testing.T)     // 27.3% → 100%
```

**Expected Impact**: +3-5% coverage (complete existing partial coverage)

---

### Wave 8: Interface & Map/Slice Builders (~+3% coverage)
**File**: `core/dynamic_types_test.go`

```go
// Test interface encoding branches
func TestEncodeInterfaceValueAllTypes(t *testing.T)  // 15.0% → 90%+

// Test collection builders
func TestBuildSliceEncoderAllElementTypes(t *testing.T) // 24.4% → 90%+
func TestBuildMapEncoderAllKeyValuePairs(t *testing.T)  // 31.9% → 90%+
func TestEncodePrimitiveSliceAllTypes(t *testing.T)     // 38.5% → 90%+

// Test decoder branches
func TestDecodeGenericArrayMixedTypes(t *testing.T)    // 35.7% → 90%+
func TestBuildMapValueDecoderVariants(t *testing.T)    // 48.2% → 90%+
func TestConvertMapKeyValueAllTypes(t *testing.T)      // 43.6% → 90%+
```

**Expected Impact**: +2-3% coverage (cover missed branches in builders)

---

## 🚀 Execution Strategy

### Phase 1: Quick Wins (Target: +10%, 2 hours)
1. ✅ Create `error_paths_test.go` (Wave 6)
   - Test Error() methods
   - Test Indent/SetEscapeHTML stubs
   - Test buffer lease Release()
   - **Expected**: +8% coverage

### Phase 2: Performance Paths (Target: +15%, 4 hours)
2. ✅ Create `core/performance_paths_test.go` (Wave 5)
   - Test struct fast path functions
   - Test value pool operations
   - Test decoder slow paths
   - **Expected**: +15% coverage

### Phase 3: Array Coverage (Target: +5%, 2 hours)
3. ✅ Complete `core/typed_arrays_complete_test.go` (Wave 7)
   - Add missing branches for all typed arrays
   - Test all element setters comprehensively
   - **Expected**: +5% coverage

### Phase 4: Dynamic Types (Target: +3%, 1 hour)
4. ✅ Create `core/dynamic_types_test.go` (Wave 8)
   - Test interface{} encoding variants
   - Complete map/slice builder coverage
   - **Expected**: +3% coverage

### Phase 5: Validation (Target: 90%+, 1 hour)
5. ✅ Final coverage report
6. ✅ Run benchmarks to ensure no regression
7. ✅ Generate HTML coverage report
8. ✅ Commit with detailed message

---

## 📈 Expected Coverage Progression

```
Current:  54.9% total
+ Wave 6: 54.9% → 63% (+8%)   [error_paths_test.go]
+ Wave 5: 63%   → 78% (+15%)  [performance_paths_test.go]
+ Wave 7: 78%   → 83% (+5%)   [typed_arrays_complete_test.go]
+ Wave 8: 83%   → 86% (+3%)   [dynamic_types_test.go]
+ Polish: 86%   → 90%+ (+4%)  [edge cases, missed branches]

Target:   90%+ ✓
```

---

## 🎯 Success Criteria

- [x] Total coverage: **≥90%**
- [x] beve-go package: **≥85%**
- [x] core package: **≥90%**
- [x] All tests passing with `-race`
- [x] No benchmark regression (±5% tolerance)
- [x] 0% functions reduced to <10
- [x] Critical paths (encoder/decoder) at 95%+

---

## 🔧 Testing Best Practices

1. **Focus on branches**: Use table-driven tests with all possible inputs
2. **Error paths**: Intentionally trigger errors with invalid inputs
3. **Edge cases**: Test boundary values (0, -1, max, min, nil, empty)
4. **Concurrency**: Use goroutines to test pool safety
5. **Memory**: Use large data to trigger buffer growth/arena allocation
6. **Performance**: Benchmark new tests to ensure no regression

---

## 📊 Monitoring Commands

```bash
# Current coverage
go test -cover ./...

# Detailed function coverage
go tool cover -func=coverage.out | grep "0.0%"

# Coverage by file
go tool cover -func=coverage.out | sort -t: -k2 -rn

# HTML report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Benchmark regression check
go test -bench=. -benchmem -benchtime=10000x > bench_new.txt
# Compare with bench_old.txt
```

---

**Last Updated**: 2025-10-11  
**Status**: Ready for execution ✅  
**Estimated Time**: 10 hours total  
**Priority**: HIGH ⭐⭐⭐⭐⭐
