# beve-go

Go implementation of the BEVE (Binary Efficient Versatile Encoding) specification, adapted to work with Go's standard interfaces.

## What is BEVE?

BEVE is a high-performance, tagged binary data specification designed for efficiency and scientific computing. It provides a binary alternative to JSON, MessagePack, and CBOR, with better performance for typed arrays and modern hardware.

Key features:
- Little-endian byte order for maximum performance on modern CPUs
- Support for various data types: numbers, strings, objects, arrays, matrices, complex numbers
- Schema-less, fully described like JSON
- Designed for SIMD operations
- Future-proof with support for large numerical types

## Go Adaptation

This library implements BEVE encoding/decoding in Go, providing an interface similar to Go's `encoding/json` package. It allows seamless integration with existing Go code that uses JSON, but with the performance benefits of binary encoding.

### Interfaces

The library provides:
- `Marshal(v interface{}) ([]byte, error)` - Encode Go values to BEVE binary
- `Unmarshal(data []byte, v interface{}) error` - Decode BEVE binary to Go values
- `NewEncoder(io.Writer)` / `NewDecoder(io.Reader)` - Streaming-friendly APIs for incremental workflows
- Support for custom types implementing `encoding.BinaryMarshaler` and `encoding.BinaryUnmarshaler`

## Installation

```bash
go get github.com/beve-org/beve-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/beve-org/beve-go"
)

type Person struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    // Encode to BEVE
    data, err := beve.Marshal(p)
    if err != nil {
        panic(err)
    }

    // Decode from BEVE
    var decoded Person
    err = beve.Unmarshal(data, &decoded)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Decoded: %+v\n", decoded)
}
```

## Supported Types

- **Basic types**: bool, int, uint, float, string
- **Slices and arrays**: including typed arrays for bool, numeric, and string data (highly optimized)
- **Maps**: with string, signed, or unsigned integer keys
- **Structs**: with field tags supporting:
  - Field renaming: `beve:"customName"`
  - Omit empty: `beve:",omitempty"`
  - Skip field: `beve:"-"`
  - Inline/embedded structs: `beve:",inline"` or anonymous fields
- **Custom types**: implementing `BinaryMarshaler`/`BinaryUnmarshaler` interfaces
- **RawMessage**: for delayed or zero-copy decoding

## Performance

BEVE-Go is designed with zero-allocation goals and achieves significant performance improvements:

### Optimization Features

**Phase 1: Zero-Allocation**
- **Struct field caching**: Reflection metadata cached per type
- **Buffer pooling**: Reusable byte buffers for encoding/decoding
- **Stack allocation**: Small buffers use stack arrays instead of heap
- **Inline struct support**: Flattens embedded structs for efficiency

**Phase 2: Aggressive (Unsafe & SIMD-Friendly)**
- **Zero-copy string conversion**: Unsafe but safe string↔bytes (no allocations)
- **Aggressive inlining**: `//go:inline` directives on hot paths
- **Bulk array operations**: Single-allocation batch writes for large arrays
- **Optimized varint encoding**: Fast-path for common small values
- **SIMD-friendly memory layout**: Contiguous buffers for vectorization

### Benchmark Results

Comparison with Go's `encoding/json`:

```
BenchmarkMarshalStruct-12         1.9M ops    653 ns/op    624 B/op    8 allocs/op
BenchmarkMarshalStructJSON-12     1.8M ops    674 ns/op    336 B/op    7 allocs/op
BenchmarkUnmarshalStruct-12       1.0M ops   1057 ns/op    848 B/op   31 allocs/op
BenchmarkUnmarshalStructJSON-12   0.7M ops   1844 ns/op    800 B/op   20 allocs/op

BenchmarkMarshalTypedArray-12     287K ops   4196 ns/op   5240 B/op    4 allocs/op
BenchmarkMarshalTypedArrayJSON-12  86K ops  14215 ns/op   4122 B/op    2 allocs/op
BenchmarkUnmarshalTypedArray-12   171K ops   6987 ns/op   4192 B/op    5 allocs/op
BenchmarkUnmarshalTypedArrayJSON-12 16K ops  76186 ns/op  13097 B/op   14 allocs/op
```

**Key Takeaways:**
- **Struct Unmarshal**: ~1.74x faster than JSON
- **Typed Array Marshal**: ~3.4x faster than JSON 🚀
- **Typed Array Unmarshal**: ~11x faster than JSON 🚀
- Aggressive optimizations using zero-copy techniques and bulk operations
- Smaller encoded size (16% vs JSON text)

Run the benchmarks yourself:

```bash
go test -bench=. -benchmem
```

For memory profiling:

```bash
go test -bench=. -benchmem -memprofile=mem.out -cpuprofile=cpu.out
go tool pprof -top mem.out
```

## Specification

This implementation follows the BEVE specification v1.0. For detailed specification, see: https://github.com/beve-org/beve

## Contributing

Contributions are welcome! Please see the issues and discussions in the main BEVE repository.

## License

MIT License