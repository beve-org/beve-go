# Struct Tag Configuration Example

This example demonstrates BEVE's flexible struct tag configuration system, allowing seamless integration with existing codebases.

## Features Demonstrated

1. **Default BEVE Tags** - Using native `beve:"..."` tags
2. **JSON Tag Compatibility** - Drop-in replacement for `encoding/json`
3. **Custom Tags** - Using `msgpack`, `proto`, or any other tag name
4. **Multi-Tag Structs** - Same struct with different tag configurations
5. **Automatic Fallback** - Falls back to `json` tags when configured tag not found
6. **Field Skipping** - Using `-` to skip sensitive fields
7. **Zero Performance Overhead** - Tag resolution at cache build time

## Running the Example

```bash
go run main.go
```

## Example Output

```
🏷️  BEVE Configurable Struct Tag Demo
================================================

📌 Scenario 1: Default BEVE Tags
Current tag: beve
Encoded size: 59 bytes
Decoded: {ID:1 Username:alice Email:alice@example.com IsActive:true}

📌 Scenario 2: Switch to JSON Tags
Current tag: json
Encoded size: 55 bytes
Decoded: {ID:2 Username:bob Email:bob@example.com IsActive:false}
...
```

## Key Concepts

### 1. Tag Configuration

```go
import beve "github.com/beve-org/beve-go"

// Set custom tag name
beve.SetStructTag("json")

// Get current tag name
currentTag := beve.GetStructTag() // Returns "json"
```

### 2. Struct Definitions

```go
// Option A: BEVE tags (default)
type User struct {
    ID   int    `beve:"id"`
    Name string `beve:"name,omitempty"`
}

// Option B: JSON tags (compatibility mode)
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name,omitempty"`
}

// Option C: Custom tags
type User struct {
    ID   int    `msgpack:"id"`
    Name string `msgpack:"name,omitempty"`
}

// Option D: Multiple tags (maximum flexibility)
type User struct {
    ID   int    `beve:"id" json:"user_id" msgpack:"uid"`
    Name string `beve:"name" json:"username" msgpack:"uname,omitempty"`
}
```

### 3. Migration from JSON

```go
// Before: using encoding/json
import "encoding/json"
data, _ := json.Marshal(user)

// After: using BEVE with json tags (ONE line change)
import beve "github.com/beve-org/beve-go"
beve.SetStructTag("json") // Add this once at startup
data, _ := beve.Marshal(user) // Same API!
```

## Supported Tag Options

All standard struct tag options work with any configured tag name:

- **Field naming**: `json:"custom_name"`
- **Omit empty**: `json:"field,omitempty"`
- **Skip field**: `json:"-"`
- **Inline structs**: `json:",inline"`

## Performance

Tag configuration has **ZERO runtime overhead**:

```
BenchmarkStructTag_BeveTag-12    370.8 ns/op    153 B/op    5 allocs/op
BenchmarkStructTag_JSONTag-12    357.9 ns/op    153 B/op    5 allocs/op
```

Tag resolution happens during type cache construction, not during encoding/decoding.

## Use Cases

### 1. **Migrating from JSON**
Existing projects with json tags can adopt BEVE without changing struct definitions.

### 2. **Multi-Format Support**
Projects using multiple serialization formats can share struct definitions.

### 3. **Legacy Compatibility**
Support existing APIs/protocols without code duplication.

### 4. **Microservices**
Different services can use different tag conventions while sharing data models.

## Best Practices

1. **Set tag once at startup** - Avoid changing tags at runtime unless necessary
2. **Use default (beve) for new projects** - Better semantics and clarity
3. **Use json compatibility for migrations** - Smooth transition path
4. **Document tag usage** - Make it clear which tag convention is used

## Notes

- Changing tag name at runtime clears internal caches (encoder/decoder)
- Empty tag name defaults to "beve"
- Always falls back to "json" tags if configured tag not found
- Thread-safe implementation using `sync.RWMutex`

## Related Documentation

- [Main README](../../README.md) - Full BEVE documentation
- [Struct Tag Tests](../../struct_tag_test.go) - Comprehensive test suite
