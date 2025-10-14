# BEVE Go - GitHub Copilot Instructions

## Project Overview
BEVE (Binary Efficient Versatile Encoding) is a high-performance binary serialization library for Go, designed to be faster than JSON, MessagePack, and CBOR while maintaining JSON compatibility.

**Status**: Production Ready (v1.3.0)  
**License**: MIT  
**MIME Type**: `application/beve`

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

## Performance Targets

**Current Benchmarks** (Neoverse-N2 ARM64):
- Small struct marshal: **1,389 ns** (2-8× faster than competitors)
- Small struct unmarshal: **1,796 ns** (46× faster than JSON)
- Large payload marshal: **121 μs** ZeroCopy (87% faster than JSON)
- Large payload unmarshal: **543 μs** (4.5× faster than JSON)

**Allocation Targets**:
- Small: 2-4 allocs per operation
- Medium: <60 allocs (vs 600+ for competitors)
- Large: <500 allocs (vs 6000+ for competitors)

## Common Issues

1. **Time.Time in structs**: Known limitation, use `int64` UnixNano() workaround
2. **Empty interfaces**: Less efficient, prefer concrete types
3. **Encoder reuse**: Always call `Close()` to return to pool
4. **Tag names**: Default is `beve`, configure with `SetStructTag("json")`

## References

- **Spec**: [SPECIFICATION.md](../SPECIFICATION.md)
- **Benchmarks**: [benchmarks/MULTI_PLATFORM.md](../benchmarks/MULTI_PLATFORM.md)
- **Examples**: [examples/](../examples/)
- **C++ Reference**: [Glaze](https://github.com/stephenberry/glaze) (original BEVE impl)

---

**When helping with this project:**
- Focus on performance and memory efficiency
- Maintain backward compatibility with v1.0 spec
- Write benchmarks for any new features
- Keep allocations minimal (use profiling tools)
- Document edge cases and limitations
