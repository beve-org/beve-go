# 🔌 Extensions System Guide

Master BEVE's 8 extensions for specialized data types and optimizations.

**Reading Time**: 20 minutes  
**Level**: Advanced  
**Prerequisites**: [Basic Usage](../getting-started/basic-usage.md)

---

## Table of Contents

1. [Extension Overview](#extension-overview)
2. [Extension 0: Field Index](#extension-0-field-index)
3. [Extension 1: Typed Arrays](#extension-1-typed-arrays)
4. [Extension 4: Timestamps](#extension-4-timestamps)
5. [Extension 5: Durations](#extension-5-durations)
6. [Extension 6: Intervals](#extension-6-intervals)
7. [Extension 8: UUIDs](#extension-8-uuids)
8. [Extension 9: RegExp](#extension-9-regexp)
9. [Auto-Detection](#auto-detection)

---

## Extension Overview

### What Are Extensions?

BEVE extensions provide optimized encoding for specialized data types:

| Extension | Purpose | Size Benefit | Speed Benefit |
|-----------|---------|--------------|---------------|
| **0** | Field Index | N/A | O(1) field access |
| **1** | Typed Arrays | 35-48% smaller | 2-3× faster |
| **4** | Timestamps | 14-16 bytes | Nanosecond precision |
| **5** | Durations | 14 bytes | 11ns encode |
| **6** | Intervals | 29 bytes | Time ranges |
| **8** | UUIDs | 50% smaller | 400× faster |
| **9** | RegExp | 7-51 bytes | Pattern + flags |

### Extension Status

✅ **Implemented** (8/12):
- Extension 0: Field Index
- Extension 1: Typed Object Arrays
- Extension 4: Timestamps
- Extension 5: Durations
- Extension 6: Intervals
- Extension 8: UUIDs
- Extension 9: RegExp

⏳ **Planned** (4/12):
- Extension 2: Typed Nested Arrays
- Extension 3: Complex Numbers
- Extension 7: Matrices
- Extension 10-11: Reserved

---

## Extension 0: Field Index

### Purpose

O(1) field access without full unmarshal.

### Problem

**Standard**: Read single field requires full unmarshal

```go
// Want: Only read "email" field
// Must: Unmarshal entire object
var user User
beve.Unmarshal(data, &user)
email := user.Email
// Cost: O(N) where N = number of fields
```

### Solution

**Field Index**: Direct field access

```go
// Encode with field index
data, _ := beve.EncodeIndexedObject(user)

// Read single field (no full unmarshal)
email, _ := beve.ReadFieldByName(data, "email")
// Cost: O(1) lookup
```

### Format

```
[Extension 0 Header]
[Field Count]
--- Index Table ---
[Field 0: name, offset, size, flags]
[Field 1: name, offset, size, flags]
...
--- Data Section ---
[Field 0 value bytes]
[Field 1 value bytes]
...
```

### Usage

```go
type User struct {
    ID       int64  `beve:"id"`
    Name     string `beve:"name"`
    Email    string `beve:"email"`
    Metadata []byte `beve:"metadata"` // Large field
}

// Encode with index
user := User{
    ID:       123,
    Name:     "Alice",
    Email:    "alice@example.com",
    Metadata: make([]byte, 100_000), // 100KB
}

data, _ := beve.EncodeIndexedObject(user)

// Read single field (fast!)
email, _ := beve.ReadFieldByName(data, "email")
// No need to decode 100KB metadata!

// Decode entire object (if needed)
decoded, _ := beve.DecodeIndexedObject(data)
```

### Performance

**Benchmark** (Large object with 20 fields):

```
FullUnmarshal:     5,200 ns/op
ReadSingleField:      77 ns/op

Speedup: 67× faster for single field access
```

### Use Cases

- ✅ Database cursors (select specific columns)
- ✅ REST APIs (partial responses)
- ✅ Large objects with selective access
- ✅ Configuration files (read specific settings)

---

## Extension 1: Typed Arrays

### Purpose

Store struct array field names once (not per element).

### Problem

**Standard**: Field names repeated for every element

```go
users := []User{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
}

// Standard encoding:
// {name: "Alice", age: 30}
// {name: "Bob", age: 25}      <- "name" and "age" repeated!
// {name: "Charlie", age: 35}  <- "name" and "age" repeated!
// Total: 112 bytes
```

### Solution

**Typed Array**: Field names stored once

```go
// Typed encoding:
// Schema: [name, age]
// Values: ["Alice", 30], ["Bob", 25], ["Charlie", 35]
// Total: 58 bytes (48% smaller!)
```

### Format

```
[Extension 1 Header]
[Field Count]
[Field Names: name, age, ...]
[Array Size]
[Values only: no keys!]
```

### Usage

```go
users := []User{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    // ... 100 users
}

// Manual typed encoding
data, _ := beve.MarshalTyped(users)

// Auto-detection (uses typed if N≥5)
data, _ := beve.MarshalAuto(users)

// Decode
var decoded []User
beve.UnmarshalAuto(data, &decoded)
```

### Performance

**Benchmark** (100 users):

| Method | Size | Encode Time | Speedup |
|--------|------|-------------|---------|
| Standard | 15,000 bytes | 30 μs | 1× |
| Typed | 9,750 bytes | 12 μs | 2.5× faster |

**Savings**: 35-48% smaller, 2-3× faster

### Threshold

Auto-detection uses typed arrays when:
- Array size ≥ 5 elements (default)
- Configurable with `MarshalOptions.MinArraySize`

```go
opts := beve.MarshalOptions{
    UseTypedSchema: true,
    MinArraySize:   10, // Use typed for N≥10
}
```

---

## Extension 4: Timestamps

### Purpose

Encode `time.Time` with nanosecond precision and timezone.

### Problem

**Standard**: Convert `time.Time` to `int64`

```go
type Event struct {
    Name      string `beve:"name"`
    Timestamp int64  `beve:"timestamp"` // Unix seconds (lose nanoseconds!)
}

event := Event{
    Name:      "Meeting",
    Timestamp: time.Now().Unix(), // Lose nanoseconds!
}
```

### Solution

**Extension 4**: Full `time.Time` support

```go
type Event struct {
    Name      string    `beve:"name"`
    Timestamp time.Time `beve:"timestamp"` // Full precision!
}

// Encode (14-16 bytes)
data, _ := beve.MarshalTimestamp(time.Now())

// Decode
decoded, _ := beve.UnmarshalTimestamp(data)
```

### Format

```
[Extension 4 Header]
[Precision Byte: 0x01 = nanosecond, has_tz bit]
[Seconds: int64 (8 bytes)]
[Nanoseconds: uint32 (4 bytes)]
[Timezone Offset: int16 (2 bytes, optional)]
```

**Size**:
- UTC: 14 bytes (seconds + nanoseconds)
- With timezone: 16 bytes (+2 bytes offset)

### Usage

```go
// Current time (UTC)
now := time.Now().UTC()
data, _ := beve.MarshalTimestamp(now)
// Size: 14 bytes

// With timezone
nowTZ := time.Now().In(time.FixedZone("EST", -5*3600))
data, _ := beve.MarshalTimestamp(nowTZ)
// Size: 16 bytes

// Decode
decoded, _ := beve.UnmarshalTimestamp(data)

// Custom timestamp
ts := beve.Timestamp{
    Seconds:     1697500000,
    Nanoseconds: 123456789,
    TimezoneOffset: nil, // UTC
}
data, _ := beve.EncodeTimestamp(ts)
```

### Performance

```
EncodeTimestamp:    2.1 ns/op
DecodeTimestamp:    3.8 ns/op
```

**Comparison with JSON**:
- JSON: `"2025-10-17T14:39:04.123456789Z"` (30 bytes)
- BEVE: `[0xA6 + 14 bytes]` (15 bytes) = **50% smaller**

---

## Extension 5: Durations

### Purpose

Encode `time.Duration` with nanosecond precision.

### Format

```
[Extension 5 Header]
[Sign: 1 byte (0x00 = positive, 0x01 = negative)]
[Seconds: int64 (8 bytes)]
[Nanoseconds: uint32 (4 bytes)]
[Precision: 1 byte (0x01 = nanosecond)]
```

**Size**: 14 bytes total

### Usage

```go
// Encode duration
duration := 5*time.Hour + 30*time.Minute + 15*time.Second + 123456*time.Nanosecond

data, _ := beve.EncodeDuration(duration)
// Size: 14 bytes

// Decode
decoded, _ := beve.DecodeDuration(data)

// Negative duration
negativeDur := -10 * time.Minute
data, _ := beve.EncodeDuration(negativeDur)
```

### Performance

```
EncodeDuration:    11 ns/op
DecodeDuration:    18 ns/op
```

### Use Cases

- ✅ API latency tracking
- ✅ Cache TTL values
- ✅ Timeout configurations
- ✅ Performance metrics

---

## Extension 6: Intervals

### Purpose

Encode time ranges (start + end).

### Format

```
[Extension 6 Header]
[Start Timestamp: Extension 4 format (14-16 bytes)]
[End Timestamp: Extension 4 format (14-16 bytes)]
```

**Size**: 29-33 bytes (14+14 or 16+16 + 1 header)

### Usage

```go
start := time.Now()
end := start.Add(1 * time.Hour)

// Encode interval
data, _ := beve.EncodeInterval(start, end)
// Size: 29 bytes (both UTC)

// Decode
decodedStart, decodedEnd, _ := beve.DecodeInterval(data)

// Check if time in interval
now := time.Now()
if now.After(decodedStart) && now.Before(decodedEnd) {
    fmt.Println("Within interval")
}
```

### Performance

```
EncodeInterval:    44 ns/op
DecodeInterval:    67 ns/op
```

### Use Cases

- ✅ Event scheduling
- ✅ Reservation systems
- ✅ Time-based access control
- ✅ Analytics time ranges

---

## Extension 8: UUIDs

### Purpose

Encode UUIDs as binary (16 bytes) instead of string (36 bytes).

### Problem

**String UUID**: 36 bytes + overhead

```go
type User struct {
    ID string `beve:"id"` // "550e8400-e29b-41d4-a716-446655440000"
}
// Size: 36 bytes + 1 header + varint = ~39 bytes
```

### Solution

**Binary UUID**: 16 bytes + 2 overhead

```go
type User struct {
    ID [16]byte `beve:"id"` // Binary UUID
}
// Size: 16 bytes + 2 header = 18 bytes (54% smaller!)
```

### Format

```
[Extension 8 Header]
[Version: 1 byte (1-5)]
[UUID Bytes: 16 bytes]
```

**Size**: 18 bytes total (vs 36-39 bytes for string)

### Usage

```go
// From string
uuidStr := "550e8400-e29b-41d4-a716-446655440000"
data, _ := beve.MarshalUUIDString(uuidStr)
// Size: 18 bytes

// From binary
var uuidBytes [16]byte
copy(uuidBytes[:], []byte{0x55, 0x0e, 0x84, 0x00, /*...*/})
data, _ := beve.MarshalUUID(uuidBytes)

// Decode to string
decodedStr, _ := beve.UnmarshalUUIDString(data)
// "550e8400-e29b-41d4-a716-446655440000"

// Decode to binary
decodedBytes, _ := beve.UnmarshalUUID(data)
```

### Performance

```
MarshalUUID:         0.3 ns/op (400× faster than string)
UnmarshalUUID:       2.1 ns/op
MarshalUUIDString:   5.8 ns/op
UnmarshalUUIDString: 12.4 ns/op
```

**Savings**: 50% smaller, 400× faster (binary vs string)

### Use Cases

- ✅ Primary keys (databases)
- ✅ Distributed IDs
- ✅ Session tokens
- ✅ Message IDs (queues, logs)

---

## Extension 9: RegExp

### Purpose

Encode regular expression patterns with flags.

### Format

```
[Extension 9 Header]
[Flags: 1 byte]
[Pattern Length: varint]
[Pattern: UTF-8 bytes]
```

**Flags**:
```go
const (
    FlagCaseInsensitive byte = 0x01 // (?i)
    FlagMultiline       byte = 0x02 // (?m)
    FlagDotAll          byte = 0x04 // (?s)
    FlagUnicode         byte = 0x08 // Unicode mode
    FlagGlobal          byte = 0x10 // Global search
)
```

**Size**: 7-51 bytes (varies with pattern length)

### Usage

```go
// Encode pattern + flags
pattern := "^[a-z]+$"
flags := beve.FlagCaseInsensitive | beve.FlagMultiline

data, _ := beve.EncodeRegExp(pattern, flags)
// Size: ~15 bytes

// Decode
regexData, _ := beve.DecodeRegExp(data)
fmt.Printf("Pattern: %s, Flags: 0x%02x\n", regexData.Pattern, regexData.Flags)

// From *regexp.Regexp
re := regexp.MustCompile("(?i)hello")
data, _ := beve.MarshalRegExp(re)

// To *regexp.Regexp
decoded, _ := beve.UnmarshalRegExp(data)
if decoded.MatchString("HELLO") {
    fmt.Println("Match!")
}
```

### Performance

```
EncodeRegExp (simple):    1.4 μs/op
EncodeRegExp (complex):   6.8 μs/op
DecodeRegExp:             2.3 μs/op
```

### Use Cases

- ✅ Validation rules
- ✅ Search patterns
- ✅ Text processing configs
- ✅ Log parsers

---

## Auto-Detection

### How It Works

All `UnmarshalAuto` functions automatically detect extension headers:

```go
// Works with ANY BEVE format
var result interface{}
beve.UnmarshalAuto(data, &result)

// Detects:
// - Standard BEVE
// - Extension 0 (Field Index)
// - Extension 1 (Typed Arrays)
// - Extension 4 (Timestamps)
// - Extension 5 (Durations)
// - Extension 6 (Intervals)
// - Extension 8 (UUIDs)
// - Extension 9 (RegExp)
```

### Detection Logic

```go
func DetectEncoding(data []byte) string {
    if len(data) == 0 {
        return "empty"
    }
    
    header := data[0]
    
    switch header {
    case 0x86:
        return "field_index"      // Extension 0
    case 0x8E:
        return "typed_array"      // Extension 1
    case 0xA6:
        return "timestamp"        // Extension 4
    case 0xAE:
        return "duration"         // Extension 5
    case 0xB6:
        return "interval"         // Extension 6
    case 0xC6:
        return "uuid"             // Extension 8
    case 0xCE:
        return "regexp"           // Extension 9
    default:
        return "standard"         // Standard BEVE
    }
}
```

### Manual Detection

```go
encoding := beve.DetectEncoding(data)

switch encoding {
case "timestamp":
    ts, _ := beve.UnmarshalTimestamp(data)
    fmt.Println("Timestamp:", ts)
case "uuid":
    uuid, _ := beve.UnmarshalUUIDString(data)
    fmt.Println("UUID:", uuid)
case "typed_array":
    var users []User
    beve.UnmarshalTyped(data, &users)
default:
    var result interface{}
    beve.Unmarshal(data, &result)
}
```

---

## Extension Comparison

### Size Comparison

| Type | Standard | Extension | Savings |
|------|----------|-----------|---------|
| **100 users** | 15,000 bytes | 9,750 bytes (Ext 1) | 35% |
| **Timestamp** | 30 bytes (JSON) | 14 bytes (Ext 4) | 53% |
| **UUID** | 36 bytes | 18 bytes (Ext 8) | 50% |
| **Duration** | 20 bytes | 14 bytes (Ext 5) | 30% |
| **Interval** | 60 bytes | 29 bytes (Ext 6) | 52% |

### Performance Comparison

| Operation | Standard | Extension | Speedup |
|-----------|----------|-----------|---------|
| **Array encode** | 30 μs | 12 μs (Ext 1) | 2.5× |
| **Field access** | 5,200 ns | 77 ns (Ext 0) | 67× |
| **UUID encode** | 120 ns | 0.3 ns (Ext 8) | 400× |
| **Timestamp** | 50 ns | 2.1 ns (Ext 4) | 24× |

---

## Best Practices

### 1. Use Auto-Detection

```go
// ✅ Good: Works with all formats
data, _ := beve.MarshalAuto(users)
var decoded []User
beve.UnmarshalAuto(data, &decoded)

// ❌ Bad: Manual format selection
data, _ := beve.MarshalTyped(users)
beve.UnmarshalTyped(data, &decoded)
// Breaks if format changes
```

### 2. Threshold Tuning

```go
// Default: Use typed arrays for N≥5
opts := beve.MarshalOptions{
    MinArraySize: 5,
}

// High-throughput: Lower threshold
opts.MinArraySize = 3 // Use typed for N≥3

// Low-latency: Higher threshold
opts.MinArraySize = 10 // Use typed for N≥10
```

### 3. Extension Negotiation

```go
// Producer capabilities
producer := beve.Capabilities{
    SupportsTypedArray: true,
    SupportsTimestamp:  true,
    SupportsUUID:       true,
}

// Consumer capabilities (old version)
consumer := beve.Capabilities{
    SupportsTypedArray: false, // Old parser
}

// Negotiate
opts := beve.NegotiateFormat(producer, consumer)
// Result: opts.UseTypedSchema = false (fallback)
```

### 4. Hybrid Encoding

```go
// Support old + new parsers
opts := beve.MarshalOptions{
    UseTypedSchema:  true,
    IncludeFallback: true, // Include generic encoding
}

data, _ := beve.MarshalWithOptions(users, opts)

// Format: [0xEE] [typed] [0xFF] [generic]
// New parsers: Use typed (fast)
// Old parsers: Use generic (compatible)
```

---

## Summary

### Extension Quick Reference

| Extension | Size | Speed | Use Case |
|-----------|------|-------|----------|
| **0** | N/A | 67× faster | Selective field access |
| **1** | 35-48% smaller | 2-3× faster | Struct arrays (N≥5) |
| **4** | 53% smaller | 24× faster | Timestamps |
| **5** | 30% smaller | Fast | Durations |
| **6** | 52% smaller | Fast | Time ranges |
| **8** | 50% smaller | 400× faster | UUIDs |
| **9** | Variable | Fast | RegExp patterns |

### When to Use Each Extension

- **Field Index (0)**: Large objects, selective reads
- **Typed Arrays (1)**: Arrays of structs (N≥5)
- **Timestamps (4)**: All `time.Time` usage
- **Durations (5)**: All `time.Duration` usage
- **Intervals (6)**: Time ranges, scheduling
- **UUIDs (8)**: All UUID storage
- **RegExp (9)**: Pattern validation, text processing

### Next Steps

- **[Arena Allocator →](arena-allocator.md)** - Advanced memory management
- **[Error Handling →](error-handling.md)** - Robust error patterns
- **[API Reference →](../api/extensions.md)** - Detailed API docs
- **[EXTENSIONS_README →](../../EXTENSIONS_README.md)** - Full specification

---

**Want more details?** See the [Extensions README](../../EXTENSIONS_README.md) for complete specification and benchmarks.
