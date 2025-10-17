# Encoder API Reference

**BEVE-Go Encoder API Documentation**

Complete API reference for BEVE encoding functions, options, and best practices.

---

## Table of Contents

1. [Core Functions](#core-functions)
2. [Encoder Options](#encoder-options)
3. [Streaming Encoder](#streaming-encoder)
4. [Buffer Pool API](#buffer-pool-api)
5. [Zero-Copy Mode](#zero-copy-mode)
6. [Extension Encoding](#extension-encoding)

---

## Core Functions

### Marshal

```go
func Marshal(v interface{}) ([]byte, error)
```

**Description**: Encode Go value to BEVE binary format.

**Parameters**:
- `v interface{}`: Any Go value (struct, map, slice, primitive)

**Returns**:
- `[]byte`: BEVE-encoded binary data
- `error`: Encoding error (nil if successful)

**Example**:

```go
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

user := User{Name: "Alice", Age: 30}
data, err := beve.Marshal(user)
if err != nil {
    log.Fatal(err)
}
// data = BEVE binary (52 bytes)
```

**Performance**: 694ns (small struct, Neoverse-N2)

---

### MarshalZeroCopy

```go
func MarshalZeroCopy(v interface{}) ([]byte, error)
```

**Description**: High-performance encoding without buffer allocation.

**When to Use**:
- Large payloads (> 1 KB)
- Performance-critical paths
- When output buffer is short-lived

**Performance**: 388ns (2-8× faster than standard Marshal)

**Example**:

```go
data, err := beve.MarshalZeroCopy(largeStruct)
// data references internal buffer - do not modify!

// Copy if you need to keep data
safeCopy := make([]byte, len(data))
copy(safeCopy, data)
```

**⚠️ Warning**: Returned slice references internal buffer. Do not modify or keep long-term.

---

### MarshalWithOptions

```go
func MarshalWithOptions(v interface{}, opts MarshalOptions) ([]byte, error)
```

**Description**: Marshal with custom encoding options.

**Parameters**:
- `v interface{}`: Value to encode
- `opts MarshalOptions`: Encoding configuration

**Example**:

```go
opts := beve.MarshalOptions{
    UseTypedSchema:  true,  // Extension 1
    UseFieldIndex:   false,
    ZeroCopy:        true,
    StructTag:       "json",  // Use json tags instead of beve
}

data, err := beve.MarshalWithOptions(users, opts)
```

---

## Encoder Options

### MarshalOptions

```go
type MarshalOptions struct {
    // Extension 1: Typed Object Arrays
    UseTypedSchema  bool
    
    // Extension 0: Field Index
    UseFieldIndex   bool
    
    // Performance
    ZeroCopy        bool
    
    // Struct tag to use (default: "beve")
    StructTag       string
    
    // Auto-detection settings
    AutoDetect      bool
    MinArraySize    int  // Min size for typed schema (default: 5)
    
    // Validation
    MaxNestingDepth int  // Default: 16
    MaxArraySize    int  // Default: 1M
}
```

**Field Details**:

#### UseTypedSchema

**Type**: `bool`  
**Default**: `false`  
**Description**: Enable Extension 1 (Typed Object Array) for homogeneous arrays.

**When to Use**:
- Arrays of structs (≥ 5 objects)
- API responses with repeated objects
- Database result sets

**Savings**: 48% smaller, 2-3× faster

```go
users := []User{ /* 100 users */ }
opts := beve.MarshalOptions{UseTypedSchema: true}
data, _ := beve.MarshalWithOptions(users, opts)
// Size: 2.7 KB (vs 5.2 KB standard)
```

#### ZeroCopy

**Type**: `bool`  
**Default**: `false`  
**Description**: Return slice referencing internal buffer (no allocation).

**Performance**: 2-8× faster, 0 allocations

**⚠️ Limitation**: Returned buffer is ephemeral - do not modify or keep.

#### StructTag

**Type**: `string`  
**Default**: `"beve"`  
**Description**: Struct tag to use for field names.

**Example**:

```go
type User struct {
    Name string `json:"username" beve:"name"`
}

// Use json tags
opts := beve.MarshalOptions{StructTag: "json"}
data, _ := beve.MarshalWithOptions(user, opts)
// Encoded with field name "username"
```

---

## Streaming Encoder

### NewStreamEncoder

```go
func NewStreamEncoder(w io.Writer) *StreamEncoder
```

**Description**: Create streaming encoder for batch operations.

**Parameters**:
- `w io.Writer`: Output stream (file, network, etc.)

**Returns**:
- `*StreamEncoder`: Encoder instance

**Features**:
- 8 KB internal buffer
- Auto-flush on buffer full
- Configurable flush interval

**Example**:

```go
f, _ := os.Create("output.beve")
defer f.Close()

enc := beve.NewStreamEncoder(f)
defer enc.Close()

for _, user := range users {
    enc.Encode(user)  // Buffered write
}
// Final flush on Close()
```

**Performance**: 6-8× faster than repeated `Marshal()` + `Write()`

---

### StreamEncoder Methods

#### Encode

```go
func (e *StreamEncoder) Encode(v interface{}) error
```

**Description**: Encode value to stream (buffered).

**Example**:

```go
for i := 0; i < 10000; i++ {
    user := User{Name: fmt.Sprintf("User%d", i), Age: i}
    if err := enc.Encode(user); err != nil {
        return err
    }
}
```

#### Flush

```go
func (e *StreamEncoder) Flush() error
```

**Description**: Write buffered data to underlying writer.

**When to Call**:
- After batch of writes
- Before long idle period
- For explicit synchronization

**Example**:

```go
enc.Encode(user1)
enc.Encode(user2)
enc.Flush()  // Ensure data written
```

#### Close

```go
func (e *StreamEncoder) Close() error
```

**Description**: Flush buffer and close encoder.

**⚠️ Important**: Always call `Close()` or use `defer enc.Close()`

#### SetBufferSize

```go
func (e *StreamEncoder) SetBufferSize(size int)
```

**Description**: Configure internal buffer size.

**Default**: 8192 bytes (8 KB)

**Example**:

```go
enc := beve.NewStreamEncoder(w)
enc.SetBufferSize(65536)  // 64 KB buffer
```

---

## Buffer Pool API

### GetEncoderFromPool

```go
func GetEncoderFromPool() *Encoder
```

**Description**: Acquire encoder from buffer pool (reuse).

**Performance**: 2-4× faster than `NewEncoder()` (allocation savings)

**⚠️ Must Return**: Call `PutEncoderToPool()` when done.

**Example**:

```go
enc := beve.GetEncoderFromPool()
defer beve.PutEncoderToPool(enc)

data := enc.Marshal(user)
// Use data...
```

---

### PutEncoderToPool

```go
func PutEncoderToPool(e *Encoder)
```

**Description**: Return encoder to pool for reuse.

**⚠️ Critical**: Always call in `defer` to prevent leaks.

**Example**:

```go
func processUser(user User) []byte {
    enc := beve.GetEncoderFromPool()
    defer beve.PutEncoderToPool(enc)  // IMPORTANT!
    
    return enc.Marshal(user)
}
```

---

### SetBufferPoolSize

```go
func SetBufferPoolSize(size int)
```

**Description**: Configure buffer pool capacity.

**Default**: 10,000 buffers

**Rule of Thumb**: `pool_size = max_concurrent_requests × 1.5`

**Example**:

```go
// For 10K concurrent requests
beve.SetBufferPoolSize(15000)
```

---

### GetPoolStats

```go
func GetPoolStats() PoolStats
```

**Description**: Query buffer pool statistics.

**Returns**:

```go
type PoolStats struct {
    TotalBuffers     int     // Pool capacity
    AvailableBuffers int     // Currently available
    Hits             uint64  // Cache hits
    Misses           uint64  // Cache misses
}

func (s PoolStats) HitRate() float64 {
    total := s.Hits + s.Misses
    if total == 0 {
        return 0
    }
    return float64(s.Hits) / float64(total)
}
```

**Example**:

```go
stats := beve.GetPoolStats()
fmt.Printf("Hit rate: %.2f%%\n", stats.HitRate()*100)
fmt.Printf("Available: %d/%d\n", stats.AvailableBuffers, stats.TotalBuffers)

if stats.HitRate() < 0.8 {
    log.Println("Pool exhausted - consider increasing size")
}
```

---

## Zero-Copy Mode

### When to Use

✅ **Use Zero-Copy When**:
- Encoding large payloads (> 1 KB)
- Performance is critical
- Output is consumed immediately
- Short-lived data

❌ **Don't Use When**:
- Need to keep encoded data
- Will modify output buffer
- Data lifetime unclear

### Performance Comparison

| Payload Size | Standard | Zero-Copy | Speedup |
|--------------|----------|-----------|---------|
| Small (52 B) | 694 ns   | 388 ns    | 1.8×    |
| Medium (18 KB) | 9 μs   | 6 μs      | 1.5×    |
| Large (180 KB) | 103 μs | 68 μs     | 1.5×    |

### Safe Usage Pattern

```go
func safeZeroCopy(v interface{}) []byte {
    // 1. Encode (zero-copy)
    temp, _ := beve.MarshalZeroCopy(v)
    
    // 2. Copy if keeping data
    result := make([]byte, len(temp))
    copy(result, temp)
    
    return result
}
```

---

## Extension Encoding

### Typed Arrays (Extension 1)

```go
users := []User{ /* homogeneous array */ }
data, _ := beve.MarshalTyped(users)
```

**Size**: 48% smaller for 100+ objects

### Field Index (Extension 0)

```go
obj := map[string]interface{}{
    "id":   123,
    "name": "Alice",
    // ... 10+ fields
}

data, _ := beve.EncodeIndexedObject(obj)
```

**Access**: O(1) field lookup (67× faster)

### Timestamps (Extension 4)

```go
now := time.Now()
data, _ := beve.MarshalTimestamp(now)
```

**Size**: 14-16 bytes (vs 31 bytes JSON)  
**Performance**: 60× faster marshal

### UUIDs (Extension 8)

```go
uuid := [16]byte{ /* UUID bytes */ }
data, _ := beve.MarshalUUID(uuid)
```

**Size**: 18 bytes (vs 38 bytes JSON)  
**Performance**: 400× faster (0.3ns marshal)

---

## Best Practices

### 1. Always Use Buffer Pool

```go
// ❌ BAD: Allocates every time
func process() {
    enc := beve.NewEncoder()
    data := enc.Marshal(user)
    // Leak!
}

// ✅ GOOD: Reuse from pool
func process() {
    enc := beve.GetEncoderFromPool()
    defer beve.PutEncoderToPool(enc)
    
    data := enc.Marshal(user)
}
```

### 2. Pre-warm Reflection Cache

```go
func init() {
    // Warm up cache at startup
    beve.Marshal(User{})
    beve.Marshal(Post{})
    // ... other types
}
```

### 3. Use Typed Schema for Arrays

```go
// Array of 100+ objects
users := []User{ /* ... */ }

// Automatically use typed schema
opts := beve.MarshalOptions{
    AutoDetect:   true,
    MinArraySize: 5,
}
data, _ := beve.MarshalWithOptions(users, opts)
```

### 4. Stream Large Datasets

```go
// ❌ BAD: Load all in memory
users := loadAllUsers()  // 1M users
data, _ := beve.Marshal(users)

// ✅ GOOD: Stream
enc := beve.NewStreamEncoder(w)
defer enc.Close()

for user := range streamUsers() {
    enc.Encode(user)
}
```

---

## Error Handling

### Common Errors

```go
var (
    ErrUnsupportedType  = errors.New("unsupported type")
    ErrNestingTooDeep   = errors.New("nesting depth exceeds limit")
    ErrCircularRef      = errors.New("circular reference detected")
    ErrBufferTooSmall   = errors.New("buffer too small")
)
```

### Error Examples

```go
// Unsupported type
func() {
    ch := make(chan int)
    _, err := beve.Marshal(ch)
    // err = ErrUnsupportedType (channels not supported)
}

// Nesting too deep
func() {
    type Node struct {
        Child *Node
    }
    // 20 levels deep
    _, err := beve.Marshal(deepNode)
    // err = ErrNestingTooDeep (max: 16)
}
```

---

## Performance Tuning

### Encoder Configuration

```go
// Production settings
beve.SetBufferPoolSize(50000)  // Large pool
beve.SetStructTag("beve")      // Explicit tags
beve.SetMaxNestingDepth(16)    // Safety limit
```

### Benchmarking

```go
func BenchmarkMarshal(b *testing.B) {
    user := User{Name: "Alice", Age: 30}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        beve.Marshal(user)
    }
}

// Result: 694 ns/op (Neoverse-N2)
```

---

## API Summary

| Function | Use Case | Performance |
|----------|----------|-------------|
| `Marshal()` | Standard encoding | 694 ns |
| `MarshalZeroCopy()` | High performance | 388 ns (1.8×) |
| `MarshalTyped()` | Homogeneous arrays | 48% smaller |
| `NewStreamEncoder()` | Batch operations | 6-8× faster |
| `GetEncoderFromPool()` | Reuse buffers | 2-4× faster |

---

**Next**: [Decoder API](decoder-api.md) · [Extension API](extension-api.md) · [Types API](types-api.md)
