# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Test Coverage - Major Improvement (October 2025) 🎯
- **Test Coverage: 62.2% → 68.0%** (+5.8 percentage points)
- **New Test Files**: 2 (extension_unmarshal_test.go, extension_utils_test.go)
- **New Test Functions**: +155 (178 total)
- **Coverage by Function**:
  - unmarshalExtension: 12% → 84% (+72%)
  - assignValue: 44% → 89% (+45%)
  - DetectEncoding: 38% → 83% (+45%)
  - readCompressedSize: 48% → 81% (+33%)
- **Test Improvements**:
  - Extension unmarshal tests (8 extension types + 9 error cases)
  - assignValue type conversion tests (int, uint, float, string, bool, slice, map)
  - DetectEncoding tests (8 extensions + 6 standard types)
  - readCompressedSize tests (round-trip, errors, offset handling)
- Files: `extension_unmarshal_test.go` (NEW), `extension_utils_test.go` (NEW)

### Performance - Slow Operations Optimization (October 2025) 🔥
- **RegExp Marshal: 173× faster** with LRU regex cache
  - Cache hit: 2,715ns → 15.7ns (0 allocs)
  - Cache miss: 34ns (1 alloc)
  - Thread-safe with `sync.RWMutex`
  - Files: `extension_regexp_cache.go` (NEW), `extension_regexp.go`
  
- **Field Index Encode: 95% fewer allocations**
  - 104 allocs → 5 allocs (single-buffer encoding)
  - 11μs → 9.3μs (15% faster)
  - Eliminated intermediate buffer copies
  - File: `extension_field_index.go`
  
- **Field Index Decode: 48% fewer allocations**
  - 204 allocs → 106 allocs (two-pass with pre-allocated arrays)
  - 4.4μs → 3.6μs (23% faster)
  - Zero-copy string references
  - File: `extension_field_index.go`

### Added - Performance Documentation
- Created comprehensive [SLOW_OPERATIONS_OPTIMIZATION.md](SLOW_OPERATIONS_OPTIMIZATION.md)
  - Detailed before/after analysis
  - Real-world impact scenarios
  - Benchmark comparisons with visualizations
  - 842 lines of technical documentation

### Changed - Documentation Policy
- **MANDATORY:** All code changes must update relevant documentation before commit
- Updated `.github/copilot-instructions.md` with strict documentation rules
- Documentation files must stay synchronized with code changes

### Performance - Pointer Optimization (January 2025) 🚀
- **67% allocation reduction** on small struct marshal (3 → 1 alloc)
- **1.14× faster marshal** (1,015ns → 889ns) by eliminating `reflect.New` heap copies
- **Zero-copy mode perfected:** 0 allocations, 0 bytes (277ns vs 889ns standard)
- **Root cause fixed:** Passing pointers avoids `ensureAddressableStruct` heap allocation (19.40% of total memory)

### Changed - Benchmark Suite Optimization
- **Updated all benchmarks to use pointers** (`Marshal(&user)` instead of `Marshal(user)`)
- Added performance best practices to README
- Created comprehensive [OPTIMIZATION_REPORT.md](OPTIMIZATION_REPORT.md) with BEVE vs CBOR analysis

### Performance vs CBOR (Apple M2 Max, ARM64)
- **Small unmarshal:** 3.2× faster (780ns vs 2,456ns)
- **Medium marshal:** 2.0× faster (7.5μs vs 15.5μs)
- **Medium unmarshal:** 3.7× faster (14.1μs vs 52.4μs)
- **Large marshal:** 1.8× faster (71μs vs 125μs)
- **Large unmarshal:** 2.8× faster (146μs vs 415μs), **93% fewer allocations** (416 vs 6,307)
- **Large map marshal:** 2.8× faster (12.4μs vs 35.0μs)
- **Overall:** BEVE wins **7 out of 8 benchmarks** against CBOR

### Changed - Core Performance Optimization (Phase 11)
- **Profile-guided optimization** of core encoding hot paths
  - Buffer.Write: Eliminated append() overhead with manual slice management
  - WriteCompressedUint: Added fast path for small values (< 64, covers 80-90% of cases)
  - Files: `core/buffer.go`, `core/encoder_write_arm64.go`, `core/encoder_write_amd64.go`

### Performance - Phase 11 Results (January 2025)
- **Small struct marshal:** 709ns → 701ns (1% improvement, already optimal)
- **Medium struct marshal:** 14,863ns → 9,654ns (**35% faster** 🚀)
- **Large struct marshal:** 206,251ns → 83,759ns (**59% faster** 🔥)
- **Competitive position:** Now dominates ALL competitors (Sonic, JSON, CBOR, MessagePack) on medium/large workloads
- **Marshal leadership:**
  - Medium: 27-71% faster than all competitors
  - Large: 27-75% faster than all competitors
  - Maps: 11-87% faster than all competitors
- **Unmarshal leadership:** 9-91% faster across all workload sizes
- **Methodology:** CPU/memory profiling → hotspot identification → targeted optimization → validation

### Added - Advanced Optimizations (Phase 10)
- **SIMD-accelerated array encoding** for numeric arrays ([]int32, []int64, []float32, []float64)
  - Platform-specific implementations: AVX2 (AMD64), NEON (ARM64)
  - 4-8× speedup for large arrays (>64 elements)
  - CPU feature detection with automatic fallback
  - Files: `core/simd.go`, `core/simd_amd64.go`, `core/simd_arm64.go`, `core/simd_generic.go`

- **Assembly-optimized varint encoding** for hot paths
  - Hand-written assembly: `core/varint_amd64.s`, `core/varint_arm64.s`
  - 2-3× speedup for string/array length encoding
  - Platform-specific optimizations with pure Go fallback

- **Code generation tool** (`cmd/bevegen`) for zero-reflection struct marshaling
  - Generates optimized `MarshalBEVE()` and `UnmarshalBEVE()` methods
  - 10× faster than reflection-based encoding
  - Usage: `//go:generate bevegen -type=MyStruct`
  - Full documentation in `cmd/bevegen/README.md`

- **Experimental arena allocator** (`core/arena.go`)
  - Request-scoped memory allocation
  - 10-100× reduction in GC pressure for high-throughput scenarios
  - Arena pooling for reusable memory regions
  - HTTP handler integration examples

- **Fast-path optimization** for all primitive types and slices
  - Zero-reflection encoding for int, int8/16/32/64, uint, uint8/16/32/64, float32/64, bool, string
  - Zero-reflection encoding for all primitive slices ([]int32, []float64, []string, etc.)
  - 36 fast-path test cases ensuring correctness
  - File: `beve.go` (extended Marshal function)

### Changed
- Reorganized documentation structure:
  - Created `docs/benchmarks/` for performance data
  - Created `docs/development/` for phase summaries
  - Moved `PHASE_*.md` files to `docs/development/`
  - Moved benchmark result files to `docs/benchmarks/`

### Performance
- Small struct marshal: ~1569ns → ~100ns (15× faster with generated code)
- Large numeric arrays: 4-8× faster with SIMD encoding
- Varint encoding: ~15ns → ~6ns (2.5× faster with assembly)
- GC pressure: 10-100× reduction with arena allocator

### Documentation
- Added `docs/development/ADVANCED_OPTIMIZATIONS.md` - comprehensive optimization guide
- Added `cmd/bevegen/README.md` - code generation tool documentation
- Updated `.github/copilot-instructions.md` with Level 2 optimization guidelines

### Added - Multi-Platform Support
- Multi-platform benchmark automation via GitHub Actions
- Automated benchmark result commits to repository
- Visual benchmark charts (PNG) for each platform
- Multi-platform comparison report (`benchmarks/MULTI_PLATFORM.md`)

### Changed
- Improved `.gitignore` to allow platform-specific benchmark results
- Updated README with links to multi-platform benchmark results

## [1.3.1] - 2025-10-11

### Added
- LICENSE file (MIT License)
- CONTRIBUTING.md with contribution guidelines
- SECURITY.md with security policy and best practices
- CHANGELOG.md for tracking changes

### Fixed
- GitHub Actions aggregate job now includes Python setup
- Corrected artifact glob pattern for multi-platform aggregation

## [1.3.0] - 2025-10-10

### Added
- Medium and large payload unmarshal benchmarks
- Zero-copy encoding with `MarshalZeroCopy()` API
- Comprehensive benchmark automation script (`scripts/bench.sh`)
- JSON-based benchmark reporting with sorted results
- Python visualization script (`scripts/plot_benchmarks.py`)
- Multi-platform CI workflow (Linux/macOS/Windows, x86/ARM)

### Changed
- Benchmark suite now uses size-specific iteration counts:
  - Small: 10,000x
  - Medium: 5,000x
  - Large: 3,000x
- Improved benchmark reporting with environment metadata
- Updated `PERFORMANCE.md` with latest results

### Performance
- Small struct marshal: **5.6× faster** than CBOR
- Small struct unmarshal: **6.4× faster** than CBOR, **22× fewer allocations**
- Medium payload: **27% faster** than CBOR
- Large payload: **1.4× faster** than CBOR

## [1.2.0] - 2025-10-08

### Added
- Core package reorganization for better modularity
- Typed array support for all primitive slice types (int8-int64, uint8-uint64, float32/64, bool, string)
- Direct slice encoding functions (`encodeInt32SliceDirect`, etc.)
- Slice encoder caching in struct field metadata
- Buffered struct field writing with pre-computed empty field masks

### Changed
- Moved encoder logic to `core/` package
- Split encoder into focused modules:
  - `encoder_base.go` - Core encoder and pooling
  - `encoder_primitives.go` - Primitive type encoding
  - `encoder_collections.go` - Collections (struct, map, array)
  - `encoder_utils.go` - Utility functions
  - `encoder_write.go` - Low-level write operations
- Improved struct encoding performance with better caching

### Performance
- **20% faster** struct encoding via optimized field iteration
- **17% less memory** usage in collection encoding
- Reduced allocations in slice encoding paths

## [1.1.0] - 2025-10-05

### Added
- Buffer pooling for encoder instances
- Type info caching to avoid repeated reflection
- Fast paths for primitive slice encoding
- Struct field cache warmup
- Batch array encoding (16-item chunks)
- Lock-free encoder cache

### Performance
- **95% allocation reduction** (362 → 17 allocs for typical workloads)
- **Phase 2 optimizations**: 20% faster, 17% less memory
- Encoder cache provides 30-40% performance improvement

### Fixed
- Struct and map field name encoding (removed incorrect type headers)
- Buffer pre-growth logic for better memory utilization

## [1.0.0] - 2025-10-01

### Added
- Initial stable release
- Full Go type system support
- Binary format with varint encoding
- Typed arrays for homogeneous data
- Struct tags (`beve:"name,omitempty"`)
- Custom marshaling via `BinaryMarshaler` interface
- Streaming encoder/decoder support
- Inline struct support
- `RawMessage` type for delayed decoding

### Features
- **64% smaller** than JSON
- **Type-safe** binary encoding
- **Zero configuration** required
- **Production ready**

### Supported Types
- Primitives: bool, int/uint (all sizes), float32/64, string
- Collections: slice, array, map, struct
- Special: nil, pointers, interfaces, embedded structs
- Custom: BinaryMarshaler/BinaryUnmarshaler interfaces

---

## Version History

- **1.3.x**: Multi-platform benchmarks, security & contribution docs
- **1.2.x**: Core package refactor, typed arrays, performance improvements
- **1.1.x**: Optimization phase - buffer pooling, caching, fast paths
- **1.0.x**: Initial stable release with full feature set

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute to BEVE.

## Security

See [SECURITY.md](SECURITY.md) for security policy and reporting vulnerabilities.
