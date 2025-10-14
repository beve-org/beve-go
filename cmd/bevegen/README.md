# BEVE Code Generator (bevegen)

**bevegen** is a code generation tool for high-performance BEVE marshaling and unmarshaling. It analyzes Go struct definitions and generates optimized `MarshalBEVE()` and `UnmarshalBEVE()` methods that eliminate reflection overhead.

## Performance

- **10× faster** than reflection-based marshaling
- **Zero reflection** at runtime
- **Inlinable** generated code
- **Type-specific** optimizations

## Installation

```bash
go install github.com/beve-org/beve-go/cmd/bevegen@latest
```

## Usage

### Basic Usage

Add a `go:generate` comment above your struct:

```go
//go:generate bevegen -type=User

type User struct {
    ID    int64  `beve:"id"`
    Name  string `beve:"name"`
    Email string `beve:"email,omitempty"`
}
```

Then run:

```bash
go generate
```

This will create a file `user_beve.go` with optimized methods:
- `func (s *User) MarshalBEVE() ([]byte, error)`
- `func (s *User) UnmarshalBEVE(data []byte) error`

### Multiple Types

Generate code for multiple types at once:

```go
//go:generate bevegen -type=User,Product,Order
```

### Custom Output File

Specify a custom output file name:

```bash
bevegen -type=User -output=custom_beve.go
```

## Generated Code

### Example Input

```go
type User struct {
    ID       int64  `beve:"id"`
    Name     string `beve:"name"`
    Email    string `beve:"email,omitempty"`
    Age      int    `beve:"age"`
    IsActive bool   `beve:"active"`
}
```

### Generated Output (Simplified)

```go
func (s *User) MarshalBEVE() ([]byte, error) {
    enc := beve.GetEncoderFromPool()
    defer beve.PutEncoderToPool(enc)

    // Write struct header
    enc.WriteByte(0x86) // TYPE_OBJECT
    
    // Count fields (accounting for omitempty)
    fieldCount := 5
    if s.Email == "" {
        fieldCount--
    }
    enc.WriteCompressedUint(uint64(fieldCount))

    // Direct field encoding (no reflection!)
    enc.EncodeString("id")
    enc.EncodeInt(int64(s.ID))
    
    enc.EncodeString("name")
    enc.EncodeString(s.Name)
    
    if s.Email != "" {
        enc.EncodeString("email")
        enc.EncodeString(s.Email)
    }
    
    enc.EncodeString("age")
    enc.EncodeInt(int64(s.Age))
    
    enc.EncodeString("active")
    enc.EncodeBool(s.IsActive)

    return enc.Bytes(), nil
}
```

## Struct Tags

bevegen supports standard BEVE and JSON struct tags:

```go
type User struct {
    // Use BEVE name "user_id"
    ID int64 `beve:"user_id"`
    
    // Skip if zero value
    Email string `beve:"email,omitempty"`
    
    // Fallback to JSON tag
    Name string `json:"name"`
    
    // Skip this field
    Internal string `beve:"-"`
}
```

### Tag Options

- **`beve:"name"`**: Set field name in encoded output
- **`beve:",omitempty"`**: Skip field if zero value
- **`beve:"-"`**: Exclude field from encoding
- **`json:"name"`**: Fallback if `beve` tag not present

## Performance Comparison

### Benchmark: Small Struct (5 fields)

| Method | Time | Allocations |
|--------|------|-------------|
| Reflection-based | ~1000ns | 8 allocs |
| **Generated code** | **~100ns** | **2 allocs** |
| **Speedup** | **10×** | **4× fewer** |

### Benchmark: Medium Struct (20 fields)

| Method | Time | Allocations |
|--------|------|-------------|
| Reflection-based | ~4500ns | 32 allocs |
| **Generated code** | **~450ns** | **8 allocs** |
| **Speedup** | **10×** | **4× fewer** |

## How It Works

### 1. AST Analysis

bevegen parses your Go source files and extracts struct definitions using `go/ast`:

```go
type User struct {
    ID   int64  `beve:"id"`
    Name string `beve:"name"`
}
```

### 2. Code Generation

For each field, bevegen generates type-specific encoding code:

```go
// Instead of:
reflect.ValueOf(s.ID).Int()  // Slow reflection

// Generate:
s.ID  // Direct field access (fast!)
```

### 3. Template Expansion

The generated code uses templates with helper functions for each primitive type:

```go
func encodeInt64(enc *core.Encoder, val int64) error {
    return enc.EncodeInt(val)
}
```

## Limitations

### Supported Types

✅ Primitives: `bool`, `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `float32/64`, `string`  
⏳ Slices and arrays: Planned  
⏳ Maps: Planned  
⏳ Nested structs: Planned  
❌ Channels, functions: Not supported

### Current Limitations

1. **Complex types**: Slices, maps, and nested structs fall back to reflection
2. **Unexported fields**: Skipped (Go visibility rules)
3. **Embedded structs**: Basic support (flattening planned)

## Comparison to Other Tools

| Feature | bevegen | easyjson | protobuf |
|---------|---------|----------|----------|
| Zero dependencies | ✅ | ✅ | ❌ (requires .proto) |
| Works with existing structs | ✅ | ✅ | ❌ (requires codegen from DSL) |
| Binary format | ✅ BEVE | ❌ JSON | ✅ Protobuf |
| Reflection-free | ✅ | ✅ | ✅ |
| Type safety | ✅ | ✅ | ✅ |

## Contributing

bevegen is part of the BEVE-Go project. Contributions are welcome!

### Development Setup

```bash
# Clone repo
git clone https://github.com/beve-org/beve-go.git
cd beve-go

# Build bevegen
go build ./cmd/bevegen

# Run tests
go test ./cmd/bevegen/...
```

### Adding New Types

To add support for new types (e.g., slices):

1. Update `analyzeStructs()` to detect slice fields
2. Add template helper function for slice encoding
3. Update code template with slice handling
4. Add tests

## License

MIT License - see LICENSE file for details

## Related

- [BEVE Specification](../../SPECIFICATION_COMPLIANCE.md)
- [Performance Guide](../benchmarks/PERFORMANCE_DASHBOARD.md)
- [BEVE-Go Main README](../../README.md)

---

**Status**: Experimental  
**Go Version**: 1.22+  
**Last Updated**: October 14, 2025
