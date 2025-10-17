# Extension API Reference

**BEVE-Go Extensions API Documentation**

Complete API reference for all 8 BEVE extensions (Ext 0-9, implemented: 0,1,4,5,6,8,9).

---

## Overview

**Implemented Extensions**:
- ✅ Extension 0: Field Index (O(1) access)
- ✅ Extension 1: Typed Object Array (48% smaller)
- ⏳ Extension 2: Typed Nested Array (planned v1.4)
- ✅ Extension 4: Timestamp (60× faster)
- ✅ Extension 5: Duration (4× faster)
- ✅ Extension 6: Interval (54× faster)
- ✅ Extension 8: UUID (400× faster)
- ✅ Extension 9: RegExp (4.9× faster)

---

## Extension 0: Field Index

### EncodeIndexedObject

```go
func EncodeIndexedObject(obj map[string]interface{}) ([]byte, error)
```

**Description**: Encode object with O(1) field access index.

**Use Case**: Large objects (10+ fields) with selective field access.

**Performance**: 67× faster field lookup (5.17μs → 77ns)

**Example**:

```go
obj := map[string]interface{}{
    "id":    123,
    "name":  "Alice",
    "email": "alice@example.com",
    // ... 10+ fields
}

data, _ := beve.EncodeIndexedObject(obj)
// Size: +24% overhead, but 67× faster access
```

### ReadFieldByName

```go
func ReadFieldByName(data []byte, fieldName string) (interface{}, error)
```

**Description**: O(1) field access without full decode.

**Example**:

```go
// Read single field (no full unmarshal)
name, _ := beve.ReadFieldByName(data, "name")
// 77ns vs 5.17μs for full decode
```

---

## Extension 1: Typed Object Array

### MarshalTyped

```go
func MarshalTyped(v interface{}) ([]byte, error)
```

**Description**: Encode array with shared schema.

**Use Case**: Homogeneous arrays (≥5 objects)

**Savings**: 48% smaller, 2.8× faster marshal

**Example**:

```go
users := []User{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    // ... 100 more
}

data, _ := beve.MarshalTyped(users)
// Size: 2.7 KB (vs 5.2 KB standard)
```

### UnmarshalTyped

```go
func UnmarshalTyped(data []byte, v interface{}) error
```

**Description**: Decode typed schema format.

**Example**:

```go
var users []User
err := beve.UnmarshalTyped(data, &users)
// 1.4× faster than standard unmarshal
```

---

## Extension 4: Timestamp

### MarshalTimestamp

```go
func MarshalTimestamp(t time.Time) ([]byte, error)
```

**Description**: Encode timestamp with nanosecond precision.

**Size**: 14-16 bytes (vs 31 bytes JSON)

**Performance**: 60× faster (1,200ns → 20ns)

**Example**:

```go
now := time.Now()
data, _ := beve.MarshalTimestamp(now)
// 14 bytes (UTC) or 16 bytes (with timezone)
```

### UnmarshalTimestamp

```go
func UnmarshalTimestamp(data []byte) (time.Time, error)
```

**Description**: Decode timestamp to `time.Time`.

**Performance**: 68× faster (2,400ns → 35ns)

**Example**:

```go
t, _ := beve.UnmarshalTimestamp(data)
fmt.Println(t)  // 2025-10-17 14:30:45.123456789 +0000 UTC
```

### EncodeTimestamp / DecodeTimestamp

```go
func EncodeTimestamp(ts Timestamp) ([]byte, error)
func DecodeTimestamp(data []byte) (Timestamp, error)

type Timestamp struct {
    Seconds        int64   // Unix epoch seconds
    Nanoseconds    uint32  // 0-999,999,999
    TimezoneOffset *int16  // Minutes from UTC (nil = UTC)
}
```

**Use Case**: Custom precision or timezone control.

**Example**:

```go
ts := beve.Timestamp{
    Seconds:     time.Now().Unix(),
    Nanoseconds: 123456789,
    TimezoneOffset: nil,  // UTC
}

data, _ := beve.EncodeTimestamp(ts)
```

---

## Extension 5: Duration

### EncodeDuration

```go
func EncodeDuration(d time.Duration) ([]byte, error)
```

**Description**: Encode time interval with nanosecond precision.

**Size**: 14 bytes fixed

**Performance**: 4× faster (45ns → 11ns)

**Example**:

```go
timeout := 30 * time.Second
data, _ := beve.EncodeDuration(timeout)
// 14 bytes
```

### DecodeDuration

```go
func DecodeDuration(data []byte) (time.Duration, error)
```

**Performance**: 4.4× faster (80ns → 18ns)

**Example**:

```go
d, _ := beve.DecodeDuration(data)
fmt.Println(d)  // 30s
```

---

## Extension 6: Interval

### EncodeInterval

```go
func EncodeInterval(start, end time.Time) ([]byte, error)
```

**Description**: Encode time range (start/end pair).

**Size**: 29-33 bytes (vs 62 bytes for 2 timestamps)

**Performance**: 54× faster (2,400ns → 44ns)

**Example**:

```go
start := time.Now()
end := start.Add(1 * time.Hour)

data, _ := beve.EncodeInterval(start, end)
// 29 bytes (UTC) or 33 bytes (with timezone)
```

### DecodeInterval

```go
func DecodeInterval(data []byte) (start, end time.Time, err error)
```

**Performance**: 68× faster (4,800ns → 70ns)

**Example**:

```go
start, end, _ := beve.DecodeInterval(data)
fmt.Printf("Range: %v to %v\n", start, end)
```

### Advanced: Inclusiveness Flags

```go
type Interval struct {
    Start           time.Time
    End             time.Time
    InclusiveStart  bool  // [  or (
    InclusiveEnd    bool  // ]  or )
}

func EncodeIntervalWithFlags(iv Interval) ([]byte, error)
func DecodeIntervalWithFlags(data []byte) (Interval, error)
```

**Example**:

```go
iv := beve.Interval{
    Start:          start,
    End:            end,
    InclusiveStart: true,   // [
    InclusiveEnd:   false,  // )
}

data, _ := beve.EncodeIntervalWithFlags(iv)
// Represents [start, end) interval
```

---

## Extension 8: UUID

### MarshalUUID

```go
func MarshalUUID(u [16]byte) ([]byte, error)
```

**Description**: Encode UUID in binary format.

**Size**: 18 bytes (vs 38 bytes JSON hyphenated string)

**Performance**: 400× faster (1,200ns → 0.3ns)

**Example**:

```go
uuid := [16]byte{
    0x55, 0x0e, 0x84, 0x00,
    0xe2, 0x9b, 0x41, 0xd4,
    0xa7, 0x16, 0x44, 0x66,
    0x55, 0x44, 0x00, 0x00,
}

data, _ := beve.MarshalUUID(uuid)
// 18 bytes (vs 36-38 bytes string)
```

### UnmarshalUUID

```go
func UnmarshalUUID(data []byte) ([16]byte, error)
```

**Performance**: 166× faster (2,500ns → 15ns)

**Example**:

```go
uuid, _ := beve.UnmarshalUUID(data)
fmt.Printf("%x\n", uuid)
```

### String Helpers

```go
func MarshalUUIDString(s string) ([]byte, error)
func UnmarshalUUIDString(data []byte) (string, error)
```

**Example**:

```go
// From string
data, _ := beve.MarshalUUIDString("550e8400-e29b-41d4-a716-446655440000")

// To string
s, _ := beve.UnmarshalUUIDString(data)
fmt.Println(s)  // "550e8400-e29b-41d4-a716-446655440000"
```

---

## Extension 9: RegExp

### EncodeRegExp

```go
func EncodeRegExp(pattern string, flags byte) ([]byte, error)
```

**Description**: Encode regex pattern with flags.

**Size**: 51% smaller (51 bytes → 25 bytes medium pattern)

**Performance**: 4.9× faster (6,800ns → 1,400ns)

**Flags**:

```go
const (
    FlagCaseInsensitive byte = 0x01  // (?i) or /i
    FlagMultiline       byte = 0x02  // (?m) or /m
    FlagDotAll          byte = 0x04  // (?s) or /s
    FlagUnicode         byte = 0x08  // Unicode mode
    FlagGlobal          byte = 0x10  // /g (global search)
)
```

**Example**:

```go
pattern := `^\w+@\w+\.\w+$`
flags := beve.FlagCaseInsensitive | beve.FlagMultiline

data, _ := beve.EncodeRegExp(pattern, flags)
// 25 bytes
```

### DecodeRegExp

```go
func DecodeRegExp(data []byte) (RegExpData, error)

type RegExpData struct {
    Pattern string
    Flags   byte
}
```

**Performance**: 3.9× faster (8,200ns → 2,100ns)

**Example**:

```go
re, _ := beve.DecodeRegExp(data)
fmt.Printf("Pattern: %s, Flags: 0x%02x\n", re.Pattern, re.Flags)
```

### regexp.Regexp Helpers

```go
func MarshalRegExp(r *regexp.Regexp) ([]byte, error)
func UnmarshalRegExp(data []byte) (*regexp.Regexp, error)
```

**Example**:

```go
// From regexp.Regexp
re := regexp.MustCompile(`(?i)hello`)
data, _ := beve.MarshalRegExp(re)

// To regexp.Regexp
re2, _ := beve.UnmarshalRegExp(data)
fmt.Println(re2.MatchString("HELLO"))  // true
```

---

## Auto-Detection

### MarshalAuto

```go
func MarshalAuto(v interface{}) ([]byte, error)
```

**Description**: Automatically choose best encoding.

**Logic**:
- Arrays ≥5 objects → Extension 1 (Typed)
- `time.Time` → Extension 4 (Timestamp)
- `[16]byte` UUID → Extension 8
- `*regexp.Regexp` → Extension 9
- Default → Standard BEVE

**Example**:

```go
users := []User{ /* 100 users */ }
data, _ := beve.MarshalAuto(users)
// Automatically uses Extension 1 (typed schema)
```

### UnmarshalAuto

```go
func UnmarshalAuto(data []byte, v interface{}) error
```

**Description**: Auto-detect extension headers.

**Example**:

```go
var result interface{}
err := beve.UnmarshalAuto(data, &result)
// Works with any BEVE format (standard or extensions)
```

---

## Utility Functions

### DetectEncoding

```go
func DetectEncoding(data []byte) string
```

**Description**: Identify encoding type.

**Returns**: `"standard"`, `"typed_array"`, `"timestamp"`, `"uuid"`, etc.

**Example**:

```go
encoding := beve.DetectEncoding(data)
fmt.Println(encoding)  // "typed_array"
```

### IsExtension

```go
func IsExtension(data []byte) bool
```

**Description**: Check if data uses extensions.

**Example**:

```go
if beve.IsExtension(data) {
    extID, _ := beve.GetExtensionID(data)
    fmt.Printf("Extension %d detected\n", extID)
}
```

---

## Performance Summary

| Extension | Marshal Speedup | Unmarshal Speedup | Size Savings |
|-----------|----------------|-------------------|--------------|
| Ext 0 (Field Index) | - | 67× (O(1) access) | +24% |
| Ext 1 (Typed Array) | 2.8× | 1.4× | 48% |
| Ext 4 (Timestamp) | 60× | 68× | 53% |
| Ext 5 (Duration) | 4× | 4.4× | 30% |
| Ext 6 (Interval) | 54× | 68× | 53% |
| Ext 8 (UUID) | 400× | 166× | 53% |
| Ext 9 (RegExp) | 4.9× | 3.9× | 51% |

---

**Next**: [Types API](types-api.md) · [Encoder API](encoder-api.md) · [Decoder API](decoder-api.md)
