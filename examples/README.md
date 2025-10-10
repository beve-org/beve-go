# BEVE Examples

This directory contains example programs demonstrating various features of the BEVE library.

## Examples

### Basic Usage (`basic/`)

Simple encode/decode example with a struct:

```bash
go run examples/basic/main.go
```

### Streaming (`streaming/`)

Demonstrates encoding/decoding multiple values in a stream:

```bash
go run examples/streaming/main.go
```

### Custom Types (`custom-types/`)

Shows how to use custom types with `encoding.BinaryMarshaler` interface:

```bash
go run examples/custom-types/main.go
```

## Running All Examples

```bash
for dir in examples/*/; do
    echo "Running $(basename $dir)..."
    go run $dir/main.go
    echo ""
done
```
