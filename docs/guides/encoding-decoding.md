# 🔄 Encoding & Decoding Deep Dive

Complete guide to BEVE's Marshal/Unmarshal system with advanced patterns and optimizations.

**Reading Time**: 20 minutes  
**Level**: Intermediate  
**Prerequisites**: [Basic Usage](../getting-started/basic-usage.md)

---

## Table of Contents

1. [Marshal System](#marshal-system)
2. [Unmarshal System](#unmarshal-system)
3. [Encoder/Decoder Lifecycle](#encoderdecoder-lifecycle)
4. [Advanced Patterns](#advanced-patterns)
5. [Type System](#type-system)
6. [Performance Optimization](#performance-optimization)
7. [Error Handling](#error-handling)
8. [Reflection Deep Dive](#reflection-deep-dive)

---

## Marshal System

### Basic Marshal

The simplest way to encode data:

```go
data, err := beve.Marshal(value)
if err != nil {
    return fmt.Errorf("marshal failed: %w", err)
}
```

**How it works:**

1. **Type inspection**: Reflect on value type
2. **Buffer allocation**: Allocate output buffer (default: 64 bytes)
3. **Encoding**: Write BEVE binary format
4. **Return**: Return encoded bytes

### Marshal Variants

BEVE provides multiple marshal functions for different use cases:

#### 1. Standard Marshal

```go
// Marshal: General purpose encoding
data, err := beve.Marshal(user)
// Use: Most common, balanced performance
// Allocations: 1-2 allocations
// Speed: Fast (1-5 μs for small structs)
```

#### 2. Zero-Copy Marshal

```go
// MarshalZeroCopy: No allocations, uses provided buffer
buf := make([]byte, 0, 1024)
data, err := beve.MarshalZeroCopy(user, buf)
// Use: Hot paths, high-throughput systems
// Allocations: 0 allocations
// Speed: Blazingly fast (0.3-2 μs, 2-8× faster)
// Warning: Caller manages buffer lifecycle
```

**Zero-Copy Example**:
```go
// Reuse buffer across multiple encodes
buf := make([]byte, 0, 4096)

for _, user := range users {
    // Encode into same buffer (no allocations)
    data, err := beve.MarshalZeroCopy(user, buf[:0])
    if err != nil {
        return err
    }
    
    // Use data immediately
    conn.Write(data)
    
    // Buffer reused in next iteration
}
```

#### 3. Marshal with Extensions

```go
// MarshalAuto: Automatically use best extension
data, err := beve.MarshalAuto(users)
// Use: Arrays of structs (N≥5 uses typed arrays)
// Benefit: 48% smaller for struct arrays
// Speed: Same as Marshal

// MarshalTyped: Force typed array encoding
data, err := beve.MarshalTyped(users)
// Use: Always use Extension 1 (typed arrays)
// Benefit: Field names stored once
```

**Extension Example**:
```go
// Array of 100 users
users := make([]User, 100)

// Standard marshal: 15,000 bytes
data1, _ := beve.Marshal(users)

// Typed marshal: 9,750 bytes (35% smaller!)
data2, _ := beve.MarshalTyped(users)

// Auto: Uses typed if N≥5
data3, _ := beve.MarshalAuto(users)
```

### Marshal Options

Control encoding behavior with options:

```go
opts := beve.MarshalOptions{
    UseTypedSchema:  true,  // Enable Extension 1
    UseFieldIndex:   true,  // Enable Extension 0
    IncludeFallback: false, // Hybrid encoding
    AutoDetect:      true,  // Smart format selection
    MinArraySize:    5,     // Threshold for typed arrays
}

data, err := beve.MarshalWithOptions(value, opts)
```

**Option Details**:

| Option | Default | Effect |
|--------|---------|--------|
| `UseTypedSchema` | `false` | Store field names once for struct arrays |
| `UseFieldIndex` | `false` | Add O(1) field lookup index |
| `IncludeFallback` | `false` | Include generic encoding for old parsers |
| `AutoDetect` | `true` | Choose best format automatically |
| `MinArraySize` | `5` | Typed array threshold (N≥5) |

### Marshal Internals

**Encoding Pipeline**:

```
Value → Type Check → Encoder Selection → Buffer Write → Return Bytes
         ↓              ↓                     ↓
      Primitive?    Fast Path          Write Header
      Struct?       Reflection         Write Data
      Slice?        Fast Path          Write Footer
```

**Fast Paths** (no reflection):

```go
// Fast paths for primitive slices
[]int32   → Typed array header + memcpy
[]float64 → Typed array header + memcpy
[]string  → Typed string array + varint sizes
[]bool    → Packed bits (8 bools per byte)

// Slow path (reflection)
[]User → Reflect each field → Encode values
```

**Performance Tips**:

```go
// ❌ Slow: Marshal inside loop
for _, user := range users {
    data, _ := beve.Marshal(user)
    send(data)
}
// 1000 allocations, 5ms total

// ✅ Fast: Marshal slice once
data, _ := beve.Marshal(users)
send(data)
// 1 allocation, 0.5ms total (10× faster!)
```

---

## Unmarshal System

### Basic Unmarshal

Decode BEVE data into a value:

```go
var user User
err := beve.Unmarshal(data, &user)
if err != nil {
    return fmt.Errorf("unmarshal failed: %w", err)
}
```

**⚠️ Important**: Always pass pointer (`&user`), not value (`user`)

### Unmarshal Variants

#### 1. Standard Unmarshal

```go
var user User
err := beve.Unmarshal(data, &user)
// Use: General purpose decoding
// Speed: Fast (0.8-2 μs for small structs)
// Memory: 2-4 allocations
```

#### 2. Unmarshal to Interface

```go
var result interface{}
err := beve.Unmarshal(data, &result)
// Use: Unknown data structure
// Result: map[string]interface{} or []interface{}
// Speed: Slower (dynamic allocation)
```

**Interface Decoding Example**:
```go
data := []byte{/* BEVE data */}

var result interface{}
beve.Unmarshal(data, &result)

// Type switch on result
switch v := result.(type) {
case map[string]interface{}:
    fmt.Println("Object:", v["name"])
case []interface{}:
    fmt.Println("Array length:", len(v))
case string:
    fmt.Println("String:", v)
case int64:
    fmt.Println("Integer:", v)
}
```

#### 3. Unmarshal with Type Hint

```go
// Pre-allocate slice capacity
users := make([]User, 0, 100)
err := beve.Unmarshal(data, &users)
// Benefit: Reduces allocations (no slice growth)
```

#### 4. Unmarshal with Extensions

```go
var users []User
err := beve.UnmarshalAuto(data, &users)
// Auto-detects extension headers
// Works with Extension 0, 1, 2, 4, 5, 6, 8, 9
```

### Unmarshal Behavior

**Type Matching**:

```go
// Strict type matching
data := beve.Marshal(int32(42))

var i int32
beve.Unmarshal(data, &i) // ✅ OK: Exact match

var j int64
beve.Unmarshal(data, &j) // ✅ OK: Compatible type

var s string
beve.Unmarshal(data, &s) // ❌ Error: Type mismatch
```

**Partial Unmarshal**:

```go
// Struct with extra fields in data
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

// Data has: {name: "Alice", age: 30, email: "alice@example.com"}
// Unmarshal to User: Only name and age extracted
// email is ignored (not in struct)

var user User
beve.Unmarshal(data, &user)
// user.Name = "Alice"
// user.Age = 30
// email is ignored
```

**Missing Fields**:

```go
type User struct {
    Name  string `beve:"name"`
    Age   int    `beve:"age"`
    Email string `beve:"email"`
}

// Data has: {name: "Alice", age: 30}
// Missing: email

var user User
beve.Unmarshal(data, &user)
// user.Name = "Alice"
// user.Age = 30
// user.Email = "" (zero value)
```

### Unmarshal Internals

**Decoding Pipeline**:

```
Bytes → Header Check → Type Dispatch → Value Creation → Field Assignment
          ↓               ↓                ↓                ↓
       Validate     Fast/Slow Path    Allocate         Set Fields
```

**Fast Paths**:

```go
// Primitive types: Direct read
int32   → Read 4 bytes → Little-endian conversion
float64 → Read 8 bytes → IEEE-754 conversion
string  → Read varint size + bytes

// Typed arrays: Bulk read
[]int32   → Read header + memcpy
[]float64 → Read header + memcpy

// Slow path: Reflection
User → Read field count → For each field:
         Read key → Lookup struct field → Unmarshal value
```

---

## Encoder/Decoder Lifecycle

### Encoder Pooling

Reuse encoders for better performance:

```go
// Get encoder from pool
enc := beve.GetEncoderFromPool()
defer beve.PutEncoderToPool(enc)

// Encode multiple values
for _, user := range users {
    data, err := enc.Marshal(user)
    if err != nil {
        return err
    }
    send(data)
    
    // Reset for next value
    enc.Reset()
}

// Return to pool
```

**Performance**:
- **Pool hit**: ~8ns overhead
- **Pool miss**: ~200ns (new encoder)
- **Savings**: 95-97% faster than creating new encoder

**With Arena**:

```go
// Arena for zero-allocation encoding
arena := beve.NewArenaPool(1024 * 1024) // 1MB arena

enc := beve.GetEncoderFromPoolWithArena(arena)
defer beve.PutEncoderToPool(enc)

// All allocations use arena (55% faster reuse)
for _, user := range users {
    data, _ := enc.Marshal(user)
    send(data)
    enc.Reset()
}

// Arena reset after batch
arena.Reset()
```

### Decoder Pooling

Similar pattern for decoders:

```go
// Get decoder from pool
dec := beve.GetDecoderFromPool(data)
defer beve.PutDecoderToPool(dec)

// Decode
var user User
err := dec.Unmarshal(&user)
```

### Streaming Encoder

For batch encoding with auto-flush:

```go
// Create stream encoder (8KB buffer)
stream := beve.NewStreamEncoder(writer)
defer stream.Close()

// Encode multiple values
for _, user := range users {
    err := stream.Encode(user)
    if err != nil {
        return err
    }
    // Auto-flushes when buffer full
}

// Final flush
stream.Flush()
```

**Benefits**:
- Buffered writes (reduces syscalls)
- Auto-flush on buffer full
- Configurable buffer size

**Example: Write to file**:

```go
f, _ := os.Create("users.beve")
defer f.Close()

stream := beve.NewStreamEncoder(f)
defer stream.Close()

for _, user := range users {
    stream.Encode(user)
}
```

### Streaming Decoder

For reading multiple values:

```go
stream := beve.NewStreamDecoder(reader)

for {
    var user User
    err := stream.Decode(&user)
    if err == io.EOF {
        break
    }
    if err != nil {
        return err
    }
    
    // Process user
    fmt.Println(user.Name)
}
```

---

## Advanced Patterns

### 1. Batch Encoding

**Pattern**: Encode multiple values efficiently

```go
// ❌ Inefficient: Marshal each separately
for _, user := range users {
    data, _ := beve.Marshal(user)
    results = append(results, data)
}
// N allocations, N marshal calls

// ✅ Efficient: Marshal as slice
data, _ := beve.Marshal(users)
// 1 allocation, 1 marshal call
```

**With Zero-Copy**:

```go
buf := make([]byte, 0, 64*1024) // 64KB buffer

for _, batch := range batches {
    // Encode batch (zero-copy)
    data, _ := beve.MarshalZeroCopy(batch, buf[:0])
    
    // Write to network
    conn.Write(data)
}
```

### 2. Conditional Encoding

**Pattern**: Encode different types based on runtime condition

```go
type Message struct {
    Type string
    Data interface{}
}

func encodeMessage(msg Message) ([]byte, error) {
    switch msg.Type {
    case "user":
        return beve.Marshal(msg.Data.(User))
    case "order":
        return beve.Marshal(msg.Data.(Order))
    default:
        return nil, fmt.Errorf("unknown type: %s", msg.Type)
    }
}
```

**With Extensions** (variant type):

```go
// Use Extension 1 (Type Tag)
data, _ := beve.MarshalVariant(0, user)  // Tag 0 = User
data, _ := beve.MarshalVariant(1, order) // Tag 1 = Order

// Decode with tag
tag, value, _ := beve.UnmarshalVariant(data)
switch tag {
case 0:
    var user User
    beve.Unmarshal(value, &user)
case 1:
    var order Order
    beve.Unmarshal(value, &order)
}
```

### 3. Incremental Decoding

**Pattern**: Decode large data in chunks

```go
// Large array [1000 objects]
data := []byte{/* 1MB of BEVE data */}

// Read header to get array size
dec := beve.NewDecoder(data)
header, _ := dec.ReadArrayHeader()
fmt.Printf("Array size: %d\n", header.Size)

// Decode one object at a time
for i := 0; i < header.Size; i++ {
    var user User
    dec.DecodeNext(&user)
    
    // Process immediately (low memory)
    processUser(user)
}
```

### 4. Selective Field Decoding

**Pattern**: Decode only needed fields (with Extension 0)

```go
// Data has 20 fields, but we only need 2
data := []byte{/* BEVE object with field index */}

// Read specific fields (O(1) lookup)
name, _ := beve.ReadFieldByName(data, "name")
age, _ := beve.ReadFieldByName(data, "age")

// No full unmarshal needed!
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

### 5. Nested Encoding

**Pattern**: Encode complex nested structures

```go
type Company struct {
    Name       string
    Departments []Department
}

type Department struct {
    Name      string
    Employees []Employee
}

// Standard: Field names repeated at each level
data1, _ := beve.Marshal(company)
// Size: Large (field names repeated 100+ times)

// With Extension 2 (Typed Nested Arrays)
data2, _ := beve.MarshalTyped(company)
// Size: 87% smaller (field names stored once per level)
```

---

## Type System

### Supported Types

BEVE supports all Go types:

#### Primitives

```go
bool          → BEVE null/boolean (1 byte)
int8/int16    → BEVE signed integer (2-3 bytes)
int32/int64   → BEVE signed integer (3-9 bytes)
uint8/uint16  → BEVE unsigned integer (2-3 bytes)
uint32/uint64 → BEVE unsigned integer (3-9 bytes)
float32/float64 → BEVE float (5-9 bytes)
string        → BEVE string (varint + bytes)
```

#### Composites

```go
struct        → BEVE object
[]T           → BEVE array (typed or generic)
[N]T          → BEVE array (fixed size)
map[K]V       → BEVE object (string/int keys)
*T            → BEVE value or null
interface{}   → BEVE any type
```

#### Special Types

```go
time.Time     → Extension 4 (timestamp)
time.Duration → Extension 5 (duration)
[16]byte      → Extension 8 (UUID)
*regexp.Regexp → Extension 9 (regexp)
```

### Type Coercion

BEVE performs safe type coercion:

```go
// Integer widening
int32 → int64   // ✅ OK: Safe widening
int64 → int32   // ⚠️ OK: Truncation possible
uint32 → int64  // ✅ OK: Safe conversion
int32 → uint32  // ⚠️ OK: Sign loss possible

// Float conversion
float32 → float64 // ✅ OK: Safe widening
float64 → float32 // ⚠️ OK: Precision loss possible

// String conversion
string → []byte // ✅ OK: UTF-8 bytes
[]byte → string // ✅ OK: UTF-8 string
```

### Custom Types

Implement custom marshaling:

#### Method 1: MarshalBEVE/UnmarshalBEVE

```go
type CustomType struct {
    Value int
}

func (c CustomType) MarshalBEVE() ([]byte, error) {
    // Custom encoding logic
    return beve.Marshal(c.Value * 2)
}

func (c *CustomType) UnmarshalBEVE(data []byte) error {
    var v int
    err := beve.Unmarshal(data, &v)
    c.Value = v / 2
    return err
}
```

#### Method 2: encoding.BinaryMarshaler

```go
func (c CustomType) MarshalBinary() ([]byte, error) {
    // Works with both BEVE and other formats
    return beve.Marshal(c.Value)
}

func (c *CustomType) UnmarshalBinary(data []byte) error {
    return beve.Unmarshal(data, &c.Value)
}
```

---

## Performance Optimization

### Optimization Checklist

#### 1. Use Zero-Copy for Hot Paths

```go
// Before: Standard marshal
for _, item := range items {
    data, _ := beve.Marshal(item)
    send(data)
}
// 1000 items = 1000 allocations

// After: Zero-copy marshal
buf := make([]byte, 0, 4096)
for _, item := range items {
    data, _ := beve.MarshalZeroCopy(item, buf[:0])
    send(data)
}
// 1000 items = 1 allocation (buffer)
```

**Speedup**: 2-8× faster

#### 2. Pool Encoders/Decoders

```go
// Before: New encoder each time
for _, item := range items {
    data, _ := beve.Marshal(item)
}
// 1000 items = 1000 encoder creations

// After: Pooled encoder
enc := beve.GetEncoderFromPool()
defer beve.PutEncoderToPool(enc)

for _, item := range items {
    data, _ := enc.Marshal(item)
    enc.Reset()
}
// 1000 items = 1 encoder (reused)
```

**Speedup**: ~8ns overhead vs ~200ns creation

#### 3. Use Typed Arrays

```go
// Before: Generic array
users := []User{...}
data, _ := beve.Marshal(users)
// 15,000 bytes (field names repeated 100×)

// After: Typed array
data, _ := beve.MarshalTyped(users)
// 9,750 bytes (field names stored once)
```

**Savings**: 35-48% smaller

#### 4. Pre-Allocate Slices

```go
// Before: Dynamic growth
var users []User
beve.Unmarshal(data, &users)
// Slice grows: 0→1→2→4→8→16→32→64→100
// Multiple allocations + copies

// After: Pre-allocated capacity
users := make([]User, 0, 100)
beve.Unmarshal(data, &users)
// Slice allocated once: 0→100
// Single allocation
```

**Savings**: 50-80% fewer allocations

#### 5. Reuse Buffers

```go
// Before: New buffer each iteration
for _, batch := range batches {
    buf := make([]byte, 0, 4096)
    data, _ := beve.MarshalZeroCopy(batch, buf)
    send(data)
}
// 100 batches = 100 buffer allocations

// After: Reuse buffer
buf := make([]byte, 0, 4096)
for _, batch := range batches {
    data, _ := beve.MarshalZeroCopy(batch, buf[:0])
    send(data)
    // buf reused
}
// 100 batches = 1 buffer allocation
```

**Savings**: 99% fewer allocations

### Performance Benchmarks

**Small Struct** (3 fields):

| Operation | Time | Memory | Allocations |
|-----------|------|--------|-------------|
| Marshal | 889ns | 1.3KB | 1 |
| MarshalZeroCopy | 388ns | 0 | 0 |
| Unmarshal | 780ns | 600B | 4 |

**Large Payload** (100 records):

| Operation | Time | Memory | Allocations |
|-----------|------|--------|-------------|
| Marshal | 121μs | 196KB | 1 |
| MarshalZeroCopy | 71μs | 39B | 0 |
| MarshalTyped | 103μs | 128KB | 1 |
| Unmarshal | 146μs | 266KB | 417 |

---

## Error Handling

### Common Errors

#### 1. Type Mismatch

```go
var i int
data := beve.Marshal("hello")

err := beve.Unmarshal(data, &i)
// Error: type mismatch: expected int, got string
```

**Solution**: Check types before unmarshal

```go
if beve.GetType(data) != beve.TypeInt {
    return errors.New("expected integer")
}
```

#### 2. Buffer Too Small

```go
buf := make([]byte, 0, 10) // Too small!
data, err := beve.MarshalZeroCopy(largeStruct, buf)
// Error: buffer too small: need 1000 bytes, have 10
```

**Solution**: Estimate size or use dynamic buffer

```go
// Estimate size
size := beve.EstimateSize(largeStruct)
buf := make([]byte, 0, size*2) // 2× safety margin

// Or use dynamic buffer
data, _ := beve.Marshal(largeStruct) // Auto-grows
```

#### 3. Invalid BEVE Data

```go
data := []byte{0xFF, 0xFF, 0xFF} // Invalid BEVE
var user User
err := beve.Unmarshal(data, &user)
// Error: invalid BEVE header
```

**Solution**: Validate before unmarshal

```go
if !beve.Valid(data) {
    return errors.New("invalid BEVE data")
}
```

#### 4. Nil Pointer

```go
var user *User // nil pointer!
err := beve.Unmarshal(data, user)
// Error: cannot unmarshal to nil pointer
```

**Solution**: Initialize pointer

```go
user := &User{}
err := beve.Unmarshal(data, user)
```

### Error Handling Patterns

#### Pattern 1: Wrap Errors

```go
func loadUser(data []byte) (*User, error) {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        return nil, fmt.Errorf("failed to load user: %w", err)
    }
    return &user, nil
}

// Caller can unwrap
err := loadUser(data)
if errors.Is(err, beve.ErrTypeMismatch) {
    // Handle type error
}
```

#### Pattern 2: Graceful Degradation

```go
func decodeUserOrDefault(data []byte) User {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        // Return default on error
        return User{Name: "Unknown"}
    }
    return user
}
```

#### Pattern 3: Partial Success

```go
func decodeUsers(data []byte) ([]User, []error) {
    var rawUsers []map[string]interface{}
    beve.Unmarshal(data, &rawUsers)
    
    var users []User
    var errors []error
    
    for i, raw := range rawUsers {
        var user User
        if err := mapToStruct(raw, &user); err != nil {
            errors = append(errors, fmt.Errorf("user %d: %w", i, err))
            continue
        }
        users = append(users, user)
    }
    
    return users, errors
}
```

---

## Reflection Deep Dive

### How BEVE Uses Reflection

**Type Cache**:

```go
// First marshal: Reflect on struct
data, _ := beve.Marshal(user)
// - Inspect User struct
// - Cache field info
// - Store field offsets

// Subsequent marshals: Use cache
data, _ := beve.Marshal(user2)
// - Lookup User in cache
// - Skip reflection
// - Direct field access
```

**Cache Hit Rate**: 99.9% (only first use reflects)

### Fast Path Detection

BEVE detects types that don't need reflection:

```go
// Fast paths (no reflection)
beve.Marshal(42)              // int → Direct encode
beve.Marshal("hello")         // string → Direct encode
beve.Marshal([]int{1,2,3})   // []int → Memcpy
beve.Marshal(true)            // bool → Direct encode

// Slow paths (reflection)
beve.Marshal(User{})          // struct → Reflect fields
beve.Marshal([]User{})        // []struct → Reflect each
beve.Marshal(interface{}(42)) // interface → Reflect type
```

### Struct Field Traversal

**Encoding Order**:

```go
type User struct {
    Name  string `beve:"name"`
    Age   int    `beve:"age"`
    Email string `beve:"email"`
}

// Marshal order: Field declaration order
// Output: {name: "Alice", age: 30, email: "alice@example.com"}
```

**Embedded Structs**:

```go
type Person struct {
    Name string `beve:"name"`
}

type Employee struct {
    Person          // Embedded
    ID     int `beve:"id"`
}

// Marshals as: {name: "Alice", id: 123}
// Embedded fields promoted to parent
```

### Performance Impact

**Reflection Cost**:

```go
// First marshal (with reflection)
user := User{Name: "Alice"}
data, _ := beve.Marshal(user)
// Time: 1,200ns (includes reflection)

// Cached marshal (no reflection)
user2 := User{Name: "Bob"}
data, _ := beve.Marshal(user2)
// Time: 889ns (cache hit, 26% faster)
```

**Optimization**: Cache is global and concurrent-safe

---

## Summary

### Key Takeaways

1. **Marshal Variants**: Use `MarshalZeroCopy` for hot paths (2-8× faster)
2. **Pool Resources**: Reuse encoders/decoders (~8ns overhead)
3. **Typed Arrays**: Use `MarshalTyped` for struct arrays (35-48% smaller)
4. **Pre-Allocate**: Pre-allocate slices and buffers (50-80% fewer allocations)
5. **Error Handling**: Always check errors, wrap for context
6. **Reflection**: First use reflects, subsequent uses use cache (26% faster)

### Performance Summary

| Pattern | Speedup | Use Case |
|---------|---------|----------|
| Zero-Copy | 2-8× | High-throughput systems |
| Encoder Pooling | 25× (8ns vs 200ns) | Hot paths |
| Typed Arrays | 35-48% smaller | Struct arrays (N≥5) |
| Buffer Reuse | 99% fewer allocs | Batch processing |
| Arena Allocator | 55% faster | Pool reuse scenarios |

### Next Steps

- **[Struct Tags →](struct-tags.md)** - Complete tag reference
- **[Streaming →](streaming.md)** - Stream encoding/decoding
- **[Performance →](performance.md)** - Deep performance tuning
- **[Extensions →](extensions.md)** - Extension system guide

---

**Want to learn more?** Check out the [API Reference](../api/core.md) for detailed function documentation.
