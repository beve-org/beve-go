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

- Basic types: bool, int, uint, float, string
- Slices and arrays
- Maps
- Structs with field tags
- Custom types implementing the binary interfaces

## Performance

BEVE is designed to be faster than MessagePack, CBOR, and BSON, especially for typed arrays. Benchmarks show significant improvements in both speed and size for numerical data.

## Specification

This implementation follows the BEVE specification v1.0. For detailed specification, see: https://github.com/beve-org/beve

## Contributing

Contributions are welcome! Please see the issues and discussions in the main BEVE repository.

## License

MIT License