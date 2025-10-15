# BEVE-Go Build Tags Strategy

## Overview

The `core` package uses conditional compilation (build tags) to provide optimized implementations for different CPU architectures while maintaining fallback support for all platforms.

## File Structure

### Optimized Version (AMD64/ARM64)
**File:** `encoder_write_common.go`  
**Build Tag:** `//go:build (amd64 || arm64) && !purego`

This file contains the **optimized implementation** for modern 64-bit architectures:
- AMD64 (Intel/AMD x86-64)
- ARM64 (Apple Silicon, AWS Graviton, etc.)

**Features:**
- Phase 11 optimizations (6.2% faster overall)
- 35% improvement in string-heavy workloads
- Direct buffer writes to avoid overhead
- Advanced inline optimizations

### Fallback Version (Other Architectures)
**File:** `encoder_write.go`  
**Build Tag:** `//go:build (!amd64 && !arm64) || purego`

This file provides **universal compatibility** for:
- 32-bit architectures (386, ARM)
- RISC-V, MIPS, s390x, etc.
- Pure Go mode (when `-tags=purego` is specified)

**Features:**
- Clean, readable implementation
- No architecture-specific optimizations
- Guaranteed compatibility

## Functions Provided

Both files implement the same interface:

```go
func (e *Encoder) WriteByte(b byte) error
func (e *Encoder) WriteBytes(data []byte) error
func (e *Encoder) WriteStringBytes(s string) error
func (e *Encoder) WriteCompressedUint(n uint64) error
```

## Build Examples

### Default Build (Uses Optimized Version on AMD64/ARM64)
```bash
go build .
```

### Force Fallback Version (Pure Go)
```bash
go build -tags=purego .
```

### Cross-Compile for Different Architectures
```bash
# For 32-bit x86 (uses fallback)
GOARCH=386 go build .

# For RISC-V (uses fallback)
GOARCH=riscv64 go build .

# For ARM64 (uses optimized)
GOARCH=arm64 go build .
```

## Performance Comparison

Benchmark on Apple M2 Max (ARM64):

| Mode | ns/op | Performance |
|------|-------|-------------|
| **Optimized** (encoder_write_common.go) | 24.18 | Baseline |
| **Fallback** (encoder_write.go) | 25.14 | -4% slower |

The optimized version provides measurable performance improvements while maintaining identical functionality.

## Testing

Both implementations are tested automatically:

```bash
# Test optimized version
go test ./core -v

# Test fallback version
go test -tags=purego ./core -v

# Benchmark comparison
go test -bench=WriteCompressedUint -benchmem ./core
go test -tags=purego -bench=WriteCompressedUint -benchmem ./core
```

## Migration Notes

**Phase 11 Migration (2025):**
- Replaced assembly implementations with pure Go
- Optimized version shows 6.2% overall improvement
- String-heavy workloads improved by 35%
- Both files now use pure Go (no assembly)

## When to Use Each Version

### Optimized Version (`encoder_write_common.go`)
✅ AMD64/ARM64 production deployments  
✅ High-performance requirements  
✅ Cloud platforms (AWS, GCP, Azure)  
✅ Apple Silicon Macs

### Fallback Version (`encoder_write.go`)
✅ Cross-platform compatibility  
✅ Embedded systems (MIPS, RISC-V)  
✅ Legacy 32-bit systems  
✅ Build verification with `-tags=purego`

## Contributing

When modifying encoder write functions:

1. **Update both files** to maintain API compatibility
2. **Keep implementations consistent** in functionality
3. **Benchmark both versions** before committing
4. **Test with and without** `-tags=purego`

Example test workflow:
```bash
# Test both versions
go test ./core -v
go test -tags=purego ./core -v

# Benchmark both
go test -bench=. ./core -benchmem
go test -tags=purego -bench=. ./core -benchmem
```

## Architecture Support Matrix

| Architecture | File Used | Status |
|-------------|-----------|--------|
| AMD64 | encoder_write_common.go | ✅ Optimized |
| ARM64 | encoder_write_common.go | ✅ Optimized |
| 386 | encoder_write.go | ✅ Fallback |
| ARM | encoder_write.go | ✅ Fallback |
| RISC-V | encoder_write.go | ✅ Fallback |
| MIPS | encoder_write.go | ✅ Fallback |
| s390x | encoder_write.go | ✅ Fallback |
| PPC64 | encoder_write.go | ✅ Fallback |
| WASM | encoder_write.go | ✅ Fallback |

## FAQ

**Q: Why two files with the same functions?**  
A: Build tags allow Go to select the right file at compile time, enabling architecture-specific optimizations without runtime overhead.

**Q: Can I force the fallback version on AMD64/ARM64?**  
A: Yes, use `-tags=purego` to compile the fallback version on any platform.

**Q: Will this affect my code?**  
A: No, both implementations provide identical API and behavior. The selection is transparent.

**Q: How do I know which version is compiled?**  
A: Run `go list -f '{{.GoFiles}}' github.com/beve-org/beve-go/core` to see which files are included.

## References

- [Go Build Constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [BEVE Specification](../../SPECIFICATION.md)
- [Phase 11 Migration Report](../../FAST_PATH_OPTIMIZATION_REPORT.md)
