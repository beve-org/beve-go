# Types API Reference

**BEVE-Go Type System Documentation**

Complete reference for supported Go types, struct tags, and type mappings.

---

## Supported Types

### Primitives

| Go Type | BEVE Type | Example |
|---------|-----------|---------|
| `bool` | Boolean | `true`, `false` |
| `int`, `int8-64` | Signed Integer | `-42`, `123` |
| `uint`, `uint8-64` | Unsigned Integer | `0`, `255` |
| `float32`, `float64` | Floating Point | `3.14`, `2.5e10` |
| `string` | UTF-8 String | `"hello"` |
| `[]byte` | String | `[]byte("data")` |
| `nil` | Null | `nil` |

### Composite Types

| Go Type | BEVE Type | Example |
|---------|-----------|---------|
| `[]T` | Array | `[]int{1,2,3}` |
| `[N]T` | Array | `[3]int{1,2,3}` |
| `map[K]V` | Object | `map[string]int{"a":1}` |
| `struct` | Object | `User{Name: "Alice"}` |

### Special Types

| Go Type | Extension | Notes |
|---------|-----------|-------|
| `time.Time` | Ext 4 | Nanosecond precision |
| `time.Duration` | Ext 5 | Time intervals |
| `[16]byte` (UUID) | Ext 8 | Binary UUID |
| `*regexp.Regexp` | Ext 9 | Regex patterns |

---

## Struct Tags

### Tag Format

```go
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age,omitempty"`
    ID   string `beve:"-"`  // Ignore field
}
```

**Tag Options**:
- `beve:"fieldname"` - Custom field name
- `beve:",omitempty"` - Omit if zero value
- `beve:"-"` - Skip field (not encoded)

### Multiple Tag Support

```go
type User struct {
    Name string `json:"username" beve:"name"`
}

// Use json tags
opts := beve.MarshalOptions{StructTag: "json"}
data, _ := beve.MarshalWithOptions(user, opts)
```

---

## Type Mapping

### BEVE → Go

| BEVE Type | Go Types (Compatible) |
|-----------|----------------------|
| Null | `nil`, `interface{}` |
| Boolean | `bool` |
| Number | `int*`, `uint*`, `float*` |
| String | `string`, `[]byte` |
| Object | `map`, `struct`, `interface{}` |
| Array | `[]T`, `[N]T`, `interface{}` |

### Conversion Rules

**Number Conversions**:

```go
// BEVE int64 → Go types
var i8  int8   // Truncates if > 127
var i32 int32  // Truncates if > 2^31-1
var f64 float64 // Converts to float
```

**String/Bytes**:

```go
// BEVE string → []byte (copy)
var b []byte

// BEVE string → string (no copy)
var s string
```

---

## Interfaces

### interface{}

**Unmarshal to `interface{}`**:

```go
var v interface{}
beve.Unmarshal(data, &v)

// Type assertions
switch val := v.(type) {
case int:
    fmt.Println("Number:", val)
case string:
    fmt.Println("String:", val)
case map[string]interface{}:
    fmt.Println("Object:", val)
case []interface{}:
    fmt.Println("Array:", val)
}
```

---

## Custom Types

### BinaryMarshaler/Unmarshaler

```go
type BinaryMarshaler interface {
    MarshalBEVE() ([]byte, error)
}

type BinaryUnmarshaler interface {
    UnmarshalBEVE([]byte) error
}
```

**Example**:

```go
type Color struct {
    R, G, B uint8
}

func (c Color) MarshalBEVE() ([]byte, error) {
    return []byte{c.R, c.G, c.B}, nil
}

func (c *Color) UnmarshalBEVE(data []byte) error {
    c.R, c.G, c.B = data[0], data[1], data[2]
    return nil
}
```

---

## Type Aliases

```go
type UserID int64
type Username string

user := struct {
    ID   UserID
    Name Username
}{
    ID:   12345,
    Name: "alice",
}

data, _ := beve.Marshal(user)
// Encoded as int64 and string
```

---

## Unsupported Types

❌ **Not Supported**:
- `chan` (channels)
- `func` (functions)
- `complex64`, `complex128` (use Extension 3 - future)
- Circular references (panics)

---

**Next**: [Encoder API](encoder-api.md) · [Decoder API](decoder-api.md) · [Extension API](extension-api.md)
