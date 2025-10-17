# BEVE Extensions for Go

**BEVE v1.3+ Extensions** - High-performance extensions for specialized data types and optimizations.

## Overview

This package implements the BEVE v1.0 extension system (Specification §6) for Go, providing:

- **Extension 0**: Field Index - O(1) field access in objects
- **Extension 1**: Typed Object Arrays - 48% size reduction for struct arrays
- **Extension 2**: Typed Nested Arrays - Exponential gains for nested structures
- **Extension 4**: Timestamp - Nanosecond precision with timezone support
- **Extension 5**: Duration - High-precision time intervals
- **Extension 6**: Interval - Start/end time pairs
- **Extension 8**: UUID - Binary UUID encoding (50% smaller than string)
- **Extension 9**: RegExp - Regular expressions with flags

## Performance Benefits

### Typed Object Arrays (Extension 1)

**Problem**: Standard BEVE repeats field names for every object in an array

```json
[
  {"name": "Alice", "age": 30},
  {"name": "Bob", "age": 25}
]
```

**Standard Encoding**: 112 bytes (field names repeated 2× = 48% waste)

**With Extension 1**: 58 bytes (field names stored once)

**Savings**: **48% smaller, 2-3× faster marshal**

### Nested Structures (Extension 2)

For deeply nested data (depth D, objects per level N):

| Depth | Objects | Standard Size | Extension Size | Savings |
|-------|---------|---------------|----------------|---------|
| D=1   | N=100   | 5.2 KB        | 2.7 KB         | 48%     |
| D=2   | N²=10K  | 520 KB        | 135 KB         | 74%     |
| D=3   | N³=1M   | 52 MB         | 6.8 MB         | 87%     |
| D=4   | N⁴=100M | 5.2 GB        | 337 MB         | **93%** |

**Formula**: `Savings = 1 - (1 / N^(D-1))`

### Field Index (Extension 0)

**Standard**: O(n) scan to find field by name  
**Extension 0**: O(1) lookup with offset table

**Use case**: Large objects with selective field access (database cursors, REST APIs)

## Quick Start

### Installation

```bash
go get github.com/meftunca/beve-go
```

### Basic Usage

#### Typed Object Arrays

```go
package main

import (
    "fmt"
    "github.com/meftunca/beve-go"
)

type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

func main() {
    users := []User{
        {"Alice", 30},
        {"Bob", 25},
    }

    // Encode with typed schema (48% smaller)
    data, _ := beve.MarshalTyped(users)
    fmt.Printf("Size: %d bytes\n", len(data))

    // Auto-detection (uses typed for N≥5)
    data, _ = beve.MarshalAuto(users)

    // Decode
    var decoded []map[string]interface{}
    beve.UnmarshalAuto(data, &decoded)
}
```

#### Timestamps

```go
import "time"

func main() {
    now := time.Now()
    
    // Encode timestamp (14 bytes UTC, 16 bytes with TZ)
    data, _ := beve.MarshalTimestamp(now)
    
    // Decode
    decoded, _ := beve.UnmarshalTimestamp(data)
    
    fmt.Println(decoded) // Nanosecond precision preserved
}
```

#### UUIDs

```go
func main() {
    uuid := [16]byte{0x55, 0x0e, 0x84, 0x00, /*...*/ }
    
    // Binary encoding (18 bytes vs 36 bytes string)
    data, _ := beve.MarshalUUID(uuid)
    
    // From string
    data, _ = beve.MarshalUUIDString("550e8400-e29b-41d4-a716-446655440000")
    
    // Decode
    decoded, _ := beve.UnmarshalUUIDString(data)
}
```

#### Field Index (Selective Access)

```go
func main() {
    obj := map[string]interface{}{
        "id":    123,
        "name":  "Alice",
        "email": "alice@example.com",
        "meta":  map[string]interface{}{"...": "large data"},
    }
    
    // Encode with index
    data, _ := beve.EncodeIndexedObject(obj)
    
    // Read single field (no full decode)
    name, _ := beve.ReadFieldByName(data, "name") // O(1) lookup
}
```

## API Reference

### High-Level API

#### Marshal Functions

```go
// MarshalTyped: Always use typed schema (Extension 1)
func MarshalTyped(v interface{}) ([]byte, error)

// MarshalAuto: Choose best encoding (typed if N≥5)
func MarshalAuto(v interface{}) ([]byte, error)

// MarshalWithOptions: Full control
func MarshalWithOptions(v interface{}, opts MarshalOptions) ([]byte, error)

type MarshalOptions struct {
    UseTypedSchema  bool // Enable Extension 1/2
    UseFieldIndex   bool // Enable Extension 0
    IncludeFallback bool // Hybrid mode (backward compat)
    AutoDetect      bool // Automatic format selection
    MinArraySize    int  // Threshold for typed (default: 5)
}
```

#### Unmarshal Functions

```go
// UnmarshalAuto: Auto-detects extension headers
func UnmarshalAuto(data []byte, v interface{}) error

// UnmarshalTyped: For Extension 1/2 data
func UnmarshalTyped(data []byte, v interface{}) error
```

### Extension-Specific APIs

#### Extension 0: Field Index

```go
func EncodeIndexedObject(obj map[string]interface{}) ([]byte, error)
func DecodeIndexedObject(data []byte) (map[string]interface{}, error)
func ReadFieldByName(data []byte, fieldName string) (interface{}, error)
```

#### Extension 1: Typed Object Array

```go
func EncodeTypedArray(v interface{}) ([]byte, error)
func DecodeTypedArray(data []byte) ([]map[string]interface{}, error)
```

#### Extension 2: Typed Nested Array

```go
func EncodeTypedNestedArray(v interface{}) ([]byte, error)
func DecodeTypedNestedArray(data []byte) ([]map[string]interface{}, error)
```

#### Extension 4: Timestamp

```go
func MarshalTimestamp(t time.Time) ([]byte, error)
func UnmarshalTimestamp(data []byte) (time.Time, error)

func EncodeTimestamp(ts Timestamp) ([]byte, error)
func DecodeTimestamp(data []byte) (Timestamp, error)

type Timestamp struct {
    Seconds        int64
    Nanoseconds    uint32
    TimezoneOffset *int16 // Minutes from UTC (nil = UTC)
}
```

#### Extension 5: Duration

```go
func EncodeDuration(d time.Duration) ([]byte, error)
func DecodeDuration(data []byte) (time.Duration, error)
```

#### Extension 6: Interval

```go
func EncodeInterval(start, end time.Time) ([]byte, error)
func DecodeInterval(data []byte) (start, end time.Time, err error)
```

#### Extension 8: UUID

```go
func MarshalUUID(u [16]byte) ([]byte, error)
func UnmarshalUUID(data []byte) ([16]byte, error)

func MarshalUUIDString(s string) ([]byte, error)
func UnmarshalUUIDString(data []byte) (string, error)
```

#### Extension 9: RegExp

```go
func EncodeRegExp(pattern string, flags byte) ([]byte, error)
func DecodeRegExp(data []byte) (RegExpData, error)

func MarshalRegExp(r *regexp.Regexp) ([]byte, error)
func UnmarshalRegExp(data []byte) (*regexp.Regexp, error)

const (
    FlagCaseInsensitive byte = 0x01 // (?i)
    FlagMultiline       byte = 0x02 // (?m)
    FlagDotAll          byte = 0x04 // (?s)
    FlagUnicode         byte = 0x08 // Unicode mode
    FlagGlobal          byte = 0x10 // Global search
)
```

## Backward Compatibility

### Hybrid Encoding

Support old parsers while using extensions:

```go
opts := beve.MarshalOptions{
    UseTypedSchema:  true,
    IncludeFallback: true, // Include generic encoding
}

data, _ := beve.MarshalWithOptions(users, opts)

// Format: [0xEE] [typed_data] [0xFF] [generic_data]
// New parsers use typed, old parsers use generic
```

### Capability Negotiation

```go
// Producer capabilities
producer := beve.Capabilities{
    SupportsTypedArray: true,
    SupportsTimestamp:  true,
}

// Consumer capabilities
consumer := beve.Capabilities{
    SupportsTypedArray: false, // Old parser
}

// Negotiate format
opts := beve.NegotiateFormat(producer, consumer)
// Result: opts.UseTypedSchema = false (fallback to generic)
```

### Auto-Detection

All unmarshal functions auto-detect extensions:

```go
var result interface{}

// Works with any BEVE format
beve.UnmarshalAuto(data, &result)
```

## Utility Functions

```go
// DetectEncoding: Identify encoding type
encoding := beve.DetectEncoding(data)
// Returns: "typed_array", "timestamp", "uuid", etc.

// IsExtension: Check if data uses extensions
if beve.IsExtension(data) {
    extID, _ := beve.GetExtensionID(data)
    fmt.Printf("Extension %d detected\n", extID)
}

// SupportsExtension: Check parser capabilities
if beve.SupportsExtension(beve.ExtTypedArray) {
    // Use typed encoding
}

// GetCapabilities: Query current capabilities
caps := beve.GetCapabilities()
fmt.Printf("Supports %d extensions\n", len(caps))
```

## Binary Format Details

### Extension 0: Field Index

```
[0x86]                  // Extension 0 header
[0x03]                  // Object type
[field_count: varint]   // Number of fields
--- Index Table ---
[name_size: varint]     // Field 0 name length
[name: bytes]           // Field 0 name
[offset: uint32]        // Offset to value (from data start)
[size: uint16]          // Value size in bytes
[flags: byte]           // Type flags
... repeat for N fields ...
--- Data Section ---
[value_0: bytes]        // Field 0 value
[value_1: bytes]        // Field 1 value
...
```

### Extension 1: Typed Object Array

```
[0x8E]                  // Extension 1 header
[field_count: varint]   // Schema size
[field_0_name]          // Field names (once!)
[field_1_name]
...
[array_size: varint]    // Object count
[obj_0_value_0]         // Values only (no keys!)
[obj_0_value_1]
...
[obj_N_value_M]
```

### Extension 4: Timestamp

```
[0xA6]                  // Extension 4 header
[precision: byte]       // Bits 1-3=precision, bit 0=has_tz
[seconds: int64]        // Little-endian epoch seconds
[nanos: uint32]         // Little-endian nanoseconds
[tz_offset: int16]      // Optional timezone (minutes from UTC)
```

### Extension 8: UUID

```
[0xC6]                  // Extension 8 header
[version: byte]         // UUID version (1-5)
[uuid: 16 bytes]        // Binary UUID
```

## Performance Tips

1. **Use MarshalAuto()** for arrays - automatic threshold detection
2. **Use Extension 0** for large objects with selective access
3. **Use Extension 2** for nested structures (depth ≥ 2)
4. **Struct tags**: Use `beve:"fieldname"` for explicit field names
5. **Pool encoders**: Extensions automatically use buffer pools

## Limitations

- **Extension 2**: Max nesting depth = 16 levels
- **Timestamps**: No leap second support (POSIX time)
- **RegExp**: Go regexp flags are limited (no lookahead, etc.)
- **UUID**: Version extracted from byte 6 (RFC 4122)

## Migration Guide

### From Standard BEVE

```go
// Before
data, _ := beve.Marshal(users)

// After (same API, better performance)
data, _ := beve.MarshalAuto(users)
```

### From JSON

```go
// Before
data, _ := json.Marshal(users)

// After (2-8× faster, 30-50% smaller)
data, _ := beve.MarshalAuto(users)
```

## Benchmarks

See [benchmarks/MULTI_PLATFORM.md](../benchmarks/MULTI_PLATFORM.md) for full results.

**Apple M2 Max ARM64:**
- Small struct marshal: **1.39 μs** (BEVE) vs **2.57 μs** (JSON) = **1.8× faster**
- Large payload: **121 μs** (BEVE ZeroCopy) vs **943 μs** (JSON) = **7.8× faster**
- Memory: **2-4 allocs** (BEVE) vs **600+ allocs** (JSON) = **150-300× fewer**

## Contributing

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## License

MIT License - See [LICENSE](../LICENSE)

## References

- [BEVE Specification](../SPECIFICATION.md)
- [Extension Design Rationale](../docs/EXTENSION_DESIGN.md)
- [Performance Analysis](../benchmarks/MULTI_PLATFORM.md)
- [C++ Reference Implementation](https://github.com/stephenberry/glaze)
