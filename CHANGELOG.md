# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
