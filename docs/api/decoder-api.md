# Decoder API Reference

**BEVE-Go Decoder API Documentation**

Complete API reference for BEVE decoding functions, options, and streaming decode.

---

## Table of Contents

1. [Core Functions](#core-functions)
2. [Decoder Options](#decoder-options)
3. [Streaming Decoder](#streaming-decoder)
4. [Type Conversion](#type-conversion)
5. [Error Handling](#error-handling)

---

## Core Functions

### Unmarshal

```go
func Unmarshal(data []byte, v interface{}) error
```

**Description**: Decode BEVE binary to Go value.

**Parameters**:
- `data []byte`: BEVE-encoded binary data
- `v interface{}`: Pointer to destination value

**Returns**:
- `error`: Decoding error (nil if successful)

**Example**:

```go
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

var user User
err := beve.Unmarshal(data, &user)
if err != nil {
    log.Fatal(err)
}
// user = {Name: "Alice", Age: 30}
```

**Performance**: 805ns (small struct, Neoverse-N2)

---

### UnmarshalAuto

```go
func UnmarshalAuto(data []byte, v interface{}) error
```

**Description**: Auto-detect extensions and decode.

**Features**:
- Detects Extension 1 (Typed Arrays)
- Detects Extension 4 (Timestamps)
- Detects Extension 8 (UUIDs)
- Falls back to standard decode

**Example**:

```go
var users []User
err := beve.UnmarshalAuto(data, &users)
// Automatically uses Extension 1 if encoded with typed schema
```

---

### UnmarshalTyped

```go
func UnmarshalTyped(data []byte, v interface{}) error
```

**Description**: Decode Extension 1 (Typed Object Array).

**When to Use**: Data encoded with `MarshalTyped()` or `UseTypedSchema: true`

**Example**:

```go
var users []User
err := beve.UnmarshalTyped(data, &users)
// Decodes shared schema format (48% smaller)
```

**Performance**: 1.4× faster than standard unmarshal

---

## Decoder Options

### UnmarshalOptions

```go
type UnmarshalOptions struct {
    // Validation
    MaxNestingDepth int
    MaxArraySize    int
    MaxMessageSize  int64
    
    // Type conversion
    AllowTypeCoercion bool
    StrictMode        bool
    
    // Performance
    DisableCache bool
}
```

**Field Details**:

#### MaxNestingDepth

**Type**: `int`  
**Default**: `16`  
**Description**: Maximum allowed nesting depth.

**Security**: Prevents stack overflow from deeply nested data.

```go
opts := beve.UnmarshalOptions{MaxNestingDepth: 8}
err := beve.UnmarshalWithOptions(data, &v, opts)
```

#### StrictMode

**Type**: `bool`  
**Default**: `false`  
**Description**: Fail on unknown fields or type mismatches.

```go
// BEVE data has field "extra" not in struct
opts := beve.UnmarshalOptions{StrictMode: true}
err := beve.UnmarshalWithOptions(data, &user, opts)
// err != nil (unknown field "extra")
```

---

## Streaming Decoder

### NewStreamDecoder

```go
func NewStreamDecoder(r io.Reader) *StreamDecoder
```

**Description**: Create streaming decoder for reading BEVE streams.

**Parameters**:
- `r io.Reader`: Input stream (file, network, etc.)

**Example**:

```go
f, _ := os.Open("input.beve")
defer f.Close()

dec := beve.NewStreamDecoder(f)
defer dec.Close()

for {
    var user User
    if err := dec.Decode(&user); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }
    
    processUser(user)
}
```

---

### StreamDecoder Methods

#### Decode

```go
func (d *StreamDecoder) Decode(v interface{}) error
```

**Description**: Decode next value from stream.

**Returns**:
- `io.EOF` when stream ends
- Error if decode fails
- `nil` if successful

**Example**:

```go
for {
    var msg Message
    err := dec.Decode(&msg)
    
    if err == io.EOF {
        break  // End of stream
    }
    if err != nil {
        return err  // Decode error
    }
    
    handleMessage(msg)
}
```

#### More

```go
func (d *StreamDecoder) More() bool
```

**Description**: Check if more data available.

**Example**:

```go
for dec.More() {
    var user User
    dec.Decode(&user)
}
```

---

## Type Conversion

### Supported Conversions

**Number Types**:

```go
// BEVE int → Go types
var i8  int8
var i16 int16
var i32 int32
var i64 int64

beve.Unmarshal(data, &i64)  // Auto-converts

// BEVE float → Go types
var f32 float32
var f64 float64

beve.Unmarshal(data, &f64)
```

**String/[]byte**:

```go
// BEVE string → []byte
var b []byte
beve.Unmarshal(data, &b)

// BEVE []byte → string
var s string
beve.Unmarshal(data, &s)
```

**Map/Struct**:

```go
// BEVE object → map
var m map[string]interface{}
beve.Unmarshal(data, &m)

// BEVE object → struct
var user User
beve.Unmarshal(data, &user)
```

---

### Type Coercion

**AllowTypeCoercion** (experimental):

```go
// BEVE: {age: "30"}  (string)
// Go:   {Age: int}

opts := beve.UnmarshalOptions{
    AllowTypeCoercion: true,
}

var user User
err := beve.UnmarshalWithOptions(data, &user, opts)
// user.Age = 30 (converted from string)
```

---

## Error Handling

### Error Types

```go
var (
    ErrInvalidHeader      = errors.New("invalid BEVE header")
    ErrUnexpectedEOF      = errors.New("unexpected end of data")
    ErrTypeMismatch       = errors.New("type mismatch")
    ErrUnknownField       = errors.New("unknown field")
    ErrNestingTooDeep     = errors.New("nesting exceeds max depth")
    ErrInvalidUTF8        = errors.New("invalid UTF-8 string")
)
```

### Error Examples

```go
// Invalid header
data := []byte{0xFF, 0x00}  // Invalid type ID
err := beve.Unmarshal(data, &v)
// err = ErrInvalidHeader

// Type mismatch
// BEVE: {age: "thirty"}  (string)
// Go:   {Age: int}
err := beve.Unmarshal(data, &user)
// err = ErrTypeMismatch

// Truncated data
data := data[:len(data)/2]  // Cut in half
err := beve.Unmarshal(data, &v)
// err = ErrUnexpectedEOF
```

---

## Validation

### IsValidBEVE

```go
func IsValidBEVE(data []byte) bool
```

**Description**: Quick validation of BEVE format.

**Example**:

```go
if !beve.IsValidBEVE(data) {
    return errors.New("invalid BEVE data")
}

err := beve.Unmarshal(data, &v)
```

---

## Best Practices

### 1. Always Use Pointers

```go
// ❌ BAD
var user User
beve.Unmarshal(data, user)  // Won't work!

// ✅ GOOD
var user User
beve.Unmarshal(data, &user)
```

### 2. Pre-allocate Slices

```go
// ❌ Allocates during unmarshal
var users []User
beve.Unmarshal(data, &users)

// ✅ Pre-allocate if size known
users := make([]User, 0, 100)
beve.Unmarshal(data, &users)
```

### 3. Validate Input

```go
func safeUnmarshal(data []byte, v interface{}) error {
    // 1. Size check
    if len(data) > 100*1024*1024 {  // 100 MB max
        return errors.New("data too large")
    }
    
    // 2. Format check
    if !beve.IsValidBEVE(data) {
        return errors.New("invalid format")
    }
    
    // 3. Unmarshal
    return beve.Unmarshal(data, v)
}
```

---

## Performance

### Benchmarks

| Operation | Standard | Typed (Ext 1) | Speedup |
|-----------|----------|---------------|---------|
| Small struct | 805 ns | - | - |
| Medium (100 objs) | 24 μs | 17 μs | 1.4× |
| Large (10K objs) | 230 μs | 100 μs | 2.3× |

### Memory Usage

```go
// Small struct: 600 bytes, 4 allocs
var user User
beve.Unmarshal(data, &user)

// Large array: 270 KB, 417 allocs
var users []User
beve.Unmarshal(data, &users)
```

---

**Next**: [Extension API](extension-api.md) · [Encoder API](encoder-api.md) · [Types API](types-api.md)
