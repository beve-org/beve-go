# 🔧 Core Package Integration Plan

**Date**: 10 Ekim 2025  
**Goal**: Migrate encoder.go to use core/ package WITHOUT breaking tests  
**Strategy**: Gradual, test-driven migration  

---

## Current Status

✅ core/ package: 1,205 lines, 7 files, compiles cleanly  
✅ All tests passing: 22/22 PASS  
❌ encoder.go: Still using old monolithic code (1,087 lines)  

---

## Migration Strategy: GRADUAL + TEST-DRIVEN

### Phase 1: Setup (5 min)
```
1. Import core package in encoder.go
2. Keep old code intact (don't delete anything yet)
3. Run tests → should still pass
```

### Phase 2: Replace Low-Level Helpers (10 min)
Start with safest functions (no dependencies):
```
✅ writeBytes() → use core.writeBytes()
✅ writeByte() → use core.writeByte()
✅ writeCompressedUint() → use core.writeCompressedUint()
```
**Test after each change!**

### Phase 3: Replace Primitive Encoders (15 min)
```
✅ encodeNull() → use core.encodeNull()
✅ encodeBool() → use core.encodeBool()
✅ encodeInt() → use core.encodeInt()
✅ encodeUint() → use core.encodeUint()
✅ encodeFloat() → use core.encodeFloat()
✅ encodeString() → use core.encodeString()
```
**Test after each change!**

### Phase 4: Replace Collection Encoders (15 min)
```
✅ encodeSlice() → use core.encodeSlice()
✅ encodeMap() → use core.encodeMapFast()
✅ encodeStruct() → use core.encodeStructFast()
```
**Test after each change!**

### Phase 5: Replace Main Dispatcher (10 min)
```
✅ encode() → use core.encode()
```
**Full test suite run!**

### Phase 6: Cleanup (10 min)
```
1. Remove old code from encoder.go
2. Keep only public API (Marshal, Unmarshal, etc.)
3. encoder struct should wrap core.encoder
4. Final test run
```

---

## Expected Results

### Before:
```
encoder.go: 1,087 lines (everything)
core/: 1,205 lines (not used)
```

### After:
```
encoder.go: ~200 lines (public API only)
core/: 1,205 lines (main implementation)

Reduction: 82% smaller encoder.go! 🎉
```

---

## Safety Checks

After EACH change:
```bash
go test -v
```

If test fails:
1. Revert last change
2. Analyze difference
3. Fix core/ or adapter code
4. Try again

---

## Rollback Plan

If things go wrong:
```bash
git checkout encoder.go
```

We have core/ package as backup - nothing will be lost!

---

## Timeline

- Phase 1: 5 min
- Phase 2: 10 min
- Phase 3: 15 min
- Phase 4: 15 min
- Phase 5: 10 min
- Phase 6: 10 min

**Total: ~65 minutes** (with testing at each step)

---

## Success Criteria

✅ All 22 tests still pass  
✅ encoder.go < 300 lines  
✅ All functionality moved to core/  
✅ No performance regressions  
✅ Clean imports and structure  

---

Ready to start? 🚀
