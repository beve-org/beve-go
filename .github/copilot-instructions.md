# BEVE Go - GitHub Copilot Instructions

## Project Overview
BEVE (Binary Efficient Versatile Encoding) is a high-performance binary serialization library for Go, designed to be faster than JSON, MessagePack, and CBOR while maintaining JSON compatibility.

Commit messages should be short!

**Status**: Production Ready (v1.3.0)  
**License**: MIT  
**MIME Type**: `application/beve`  
**Test Coverage**: 61.7% (23 test functions, 15 benchmarks)  
**Extensions**: 8/12 implemented (67% complete)

## Core Principles

1. **Performance First**: Optimized for modern CPUs with SIMD-friendly design
2. **Little Endian**: All multi-byte values use little-endian byte order
3. **Tagged Format**: Self-describing like JSON, no schema required
4. **Type Safety**: Full Go type system support with struct tags
5. **Zero Config**: Drop-in replacement for `encoding/json`

## Binary Format Spec (BEVE v1.0)

### Type System (3-bit header)
- `0b000`: null/boolean
- `0b001`: number (float/int/uint with byte count indicators)
- `0b010`: UTF-8 string
- `0b011`: object (string/int keys)
- `0b100`: typed array (homogeneous, compact)
- `0b101`: generic array (mixed types)
- `0b110`: extensions (matrices, complex numbers, variants)

### Key Features
- **Compressed integers**: 2-bit size indicator (1/2/4/8 bytes)
- **Typed arrays**: No per-element headers, SIMD-optimized
- **Boolean packing**: 8 booleans per byte (LSB-first)
- **IEEE-754 floats**: Support for bfloat16, float16/32/64/128
- **Strings**: UTF-8 with size prefix (no null termination)

## Architecture

### Core Components
- `encoder.go`: Marshal with buffer pooling, zero-copy mode
- `decoder.go`: Unmarshal with reflection caching
- `stream.go`: Streaming encoder (8KB buffer, auto-flush)
- `byte_pool.go`: Lock-free buffer pools
- `unsafe.go`: Performance optimizations

### Extension System (v1.3.0)
- `extension_api.go`: High-level extension API (MarshalAuto, UnmarshalAuto)
- `extension_field_index.go`: Extension 0 - O(1) field access (77ns)
- `extension_typed_array.go`: Extension 1 - Typed arrays (25-48% size reduction)
- `extension_timestamp.go`: Extension 4 - Timestamps (nanosecond precision)
- `extension_duration.go`: Extension 5 - Durations (14 bytes, 11ns encode)
- `extension_interval.go`: Extension 6 - Time intervals (29 bytes)
- `extension_uuid.go`: Extension 8 - UUIDs (50% smaller, 0.3ns marshal)
- `extension_regexp.go`: Extension 9 - RegExp patterns (7-51 bytes)

### Test Infrastructure
- `extension_test.go`: Core extension tests (6 test functions)
- `extension_advanced_test.go`: Advanced tests (17 test functions, 150+ cases)
- `extension_benchmark_test.go`: Performance benchmarks (12 benchmark functions, 40+ scenarios)

### Performance Patterns
- **Buffer Pooling**: Reuse with `GetEncoderFromPool()` / `PutEncoderToPool()`
- **ZeroCopy Mode**: Use `MarshalZeroCopy()` for 2-8× faster encoding
- **Streaming**: `NewStreamEncoder(w)` for batch operations
- **Type Caching**: Struct fields cached on first use

## Coding Guidelines

### When Writing Code
1. **Prefer little-endian operations**: Use `binary.LittleEndian`
2. **Avoid allocations**: Reuse buffers, use sync.Pool
3. **Benchmark everything**: Run `./scripts/bench.sh` for validation
4. **Maintain compatibility**: Test against JSON round-trips
5. **Document benchmarks**: Update `benchmarks/MULTI_PLATFORM.md`

### Common Patterns
- Varint encoding for sizes (2-bit + value bits)
- Pre-allocate buffers based on type size estimates
- Fast paths for primitive slices (no reflection)
- Unsafe optimizations for hot paths (with safety checks)

### Testing Requirements
- Unit tests for all marshal/unmarshal paths
- Benchmark tests vs JSON/CBOR/MessagePack/Sonic
- Integration tests for real-world structs
- Round-trip validation (BEVE → JSON → BEVE)
- Cross-platform CI (AMD64, ARM64, Windows)
- Extension tests with edge cases and error paths
- Coverage target: >60% (current: 68.0%)

**⚠️ Cross-Platform Compatibility Rules:**
- **NEVER use hardcoded paths** like `/tmp/` (Unix-only)
- **ALWAYS use `os.TempDir()`** for temp files
- **ALWAYS use `filepath.Join()`** for path construction
- Test on Windows CI before merging (GitHub Actions)

### CI/CD Automation
- **Multi-platform benchmarks**: Automatic testing on M1, Neoverse-N2, EPYC, Windows
- **Extension tracking**: Dedicated benchmark step for all 8 extensions
- **Coverage reports**: HTML reports with function-level analysis generated on every run
- **Artifact preservation**: Platform-specific results with JSON aggregation
- **Visualization**: Automatic chart generation (PNG) for performance comparison

## Performance Targets

**Current Benchmarks** (Neoverse-N2 ARM64):
- Small struct marshal: **1,389 ns** (2-8× faster than competitors)
- Small struct unmarshal: **1,796 ns** (46× faster than JSON)
- Large payload marshal: **121 μs** ZeroCopy (87% faster than JSON)
- Large payload unmarshal: **543 μs** (4.5× faster than JSON)

**Extension Performance**:
- Field Index (O(1) access): **77 ns** per field
- UUID Binary Marshal: **0.3 ns** (400× faster than string)
- Duration Marshal: **11 ns** (nanosecond precision)
- Interval Marshal: **44 ns** (29 bytes total)
- Extension Detection: **~2 ns** (essentially free)
- RegExp Marshal: **1.4-6.8 μs** (pattern complexity dependent)

**Allocation Targets**:
- Small: 2-4 allocs per operation
- Medium: <60 allocs (vs 600+ for competitors)
- Large: <500 allocs (vs 6000+ for competitors)
- Extension operations: 0-3 allocs (most zero-allocation)

## Common Issues

1. **Time.Time in structs**: Known limitation, use `int64` UnixNano() workaround
2. **Empty interfaces**: Less efficient, prefer concrete types
3. **Encoder reuse**: Always call `Close()` to return to pool
4. **Tag names**: Default is `beve`, configure with `SetStructTag("json")`
5. **Windows paths**: Use `os.TempDir()` and `filepath.Join()`, never hardcode `/tmp/`

## References

- **Spec**: [SPECIFICATION.md](../SPECIFICATION.md)
- **Benchmarks**: [benchmarks/MULTI_PLATFORM.md](../benchmarks/MULTI_PLATFORM.md)
- **Examples**: [examples/](../examples/)
- **Extensions Guide**: [EXTENSIONS_README.md](../EXTENSIONS_README.md)
- **Test Coverage**: [COVERAGE_IMPROVEMENT_REPORT.md](../COVERAGE_IMPROVEMENT_REPORT.md)
- **Test Enhancement**: [TEST_ENHANCEMENT_SUMMARY.md](../TEST_ENHANCEMENT_SUMMARY.md)
- **Documentation Update**: [DOCUMENTATION_UPDATE_SUMMARY.md](../DOCUMENTATION_UPDATE_SUMMARY.md)
- **Performance Optimization**: [SLOW_OPERATIONS_OPTIMIZATION.md](../SLOW_OPERATIONS_OPTIMIZATION.md)
- **C++ Reference**: [Glaze](https://github.com/stephenberry/glaze) (original BEVE impl)

---

## 🚨 CRITICAL: Documentation Update Policy

**⚠️ MANDATORY RULE: Every code change MUST update relevant documentation immediately.**

When making ANY change (optimization, feature, bug fix):

1. **BEFORE committing code**:
   - ✅ Update README.md if user-facing behavior changes
   - ✅ Update relevant .md files (benchmarks, coverage, optimization reports)
   - ✅ Update copilot-instructions.md with new metrics/files
   - ✅ Update CHANGELOG.md with changes
   - ✅ Add/update code comments for complex logic

2. **Documentation Files That Must Stay Current**:
   - `README.md` - Main documentation, examples, benchmarks
   - `COVERAGE_IMPROVEMENT_REPORT.md` - After test changes
   - `SLOW_OPERATIONS_OPTIMIZATION.md` - After performance work
   - `benchmarks/MULTI_PLATFORM.md` - After CI/CD benchmark runs
   - `EXTENSIONS_README.md` - After extension changes
   - `.github/copilot-instructions.md` - After project structure changes

3. **Commit Message Must Reference Docs**:
   ```
   ✅ Good: "⚡ perf: Optimize X (2× faster) - Updated SLOW_OPERATIONS_OPTIMIZATION.md"
   ❌ Bad:  "⚡ perf: Optimize X (2× faster)"
   ```

4. **CI/CD Rule**: If docs are outdated, workflow should warn (future enhancement)

**THIS IS NON-NEGOTIABLE. Outdated documentation is worse than no documentation.**

---

**When helping with this project:**
- Focus on performance and memory efficiency
- Maintain backward compatibility with v1.0 spec
- Write benchmarks for any new features
- Keep allocations minimal (use profiling tools)
- Document edge cases and limitations
- Test extensions thoroughly with edge cases and error paths
- **UPDATE ALL RELEVANT DOCUMENTATION BEFORE COMMITTING** 🚨
- Run `./scripts/bench.sh` to validate performance improvements
- Check CI/CD artifacts for multi-platform benchmark results
