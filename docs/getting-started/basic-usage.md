# 📖 Basic Usage Guide

Comprehensive guide to common BEVE-Go usage patterns and best practices.

**Reading Time**: 10 minutes  
**Prerequisites**: [Quick Start](quick-start.md) completed

---

## Table of Contents

- [Type Support](#type-support)
- [Encoding Patterns](#encoding-patterns)
- [Decoding Patterns](#decoding-patterns)
- [Struct Tags](#struct-tags)
- [Working with Maps](#working-with-maps)
- [Slices and Arrays](#slices-and-arrays)
- [Pointers and Nil Values](#pointers-and-nil-values)
- [Interface Types](#interface-types)
- [Custom Types](#custom-types)
- [Best Practices](#best-practices)

---

## Type Support

### Primitive Types

```go
// Boolean
var b bool = true
data, _ := beve.Marshal(b)

// Integers (all sizes)
var i8 int8 = 127
var i16 int16 = 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// Unsigned integers
var u8 uint8 = 255
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615

// Floats
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// String
var s string = "Hello, BEVE!"
```

### Composite Types

```go
// Struct
type Person struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

// Slice
var numbers []int = []int{1, 2, 3, 4, 5}

// Array (fixed size)
var fixed [5]int = [5]int{1, 2, 3, 4, 5}

// Map (string keys)
var dict map[string]interface{} = map[string]interface{}{
    "key": "value",
}

// Map (int keys)
var indexed map[int]string = map[int]string{
    1: "one",
    2: "two",
}

// Pointer
var ptr *Person = &Person{Name: "Alice", Age: 30}

// Interface (any type)
var any interface{} = "can be anything"
```

---

## Encoding Patterns

### Pattern 1: Simple Values

```go
// Encode primitives
data, err := beve.Marshal(42)
data, err := beve.Marshal("hello")
data, err := beve.Marshal(true)
data, err := beve.Marshal(3.14)
```

### Pattern 2: Structs

```go
type User struct {
    ID       int64  `beve:"id"`
    Username string `beve:"username"`
    Email    string `beve:"email"`
}

user := User{
    ID:       123,
    Username: "alice",
    Email:    "alice@example.com",
}

data, err := beve.Marshal(user)
if err != nil {
    log.Fatal(err)
}
```

### Pattern 3: Nested Structs

```go
type Address struct {
    Street  string `beve:"street"`
    City    string `beve:"city"`
    Country string `beve:"country"`
}

type Person struct {
    Name    string  `beve:"name"`
    Age     int     `beve:"age"`
    Address Address `beve:"address"`
}

person := Person{
    Name: "Alice",
    Age:  30,
    Address: Address{
        Street:  "123 Main St",
        City:    "NYC",
        Country: "USA",
    },
}

data, _ := beve.Marshal(person)
```

### Pattern 4: Slices of Structs

```go
users := []User{
    {ID: 1, Username: "alice", Email: "alice@example.com"},
    {ID: 2, Username: "bob", Email: "bob@example.com"},
    {ID: 3, Username: "charlie", Email: "charlie@example.com"},
}

data, _ := beve.Marshal(users)
```

### Pattern 5: Maps with Mixed Values

```go
config := map[string]interface{}{
    "host":     "localhost",
    "port":     8080,
    "enabled":  true,
    "timeout":  30.0,
    "retries":  3,
    "features": []string{"auth", "logging"},
}

data, _ := beve.Marshal(config)
```

---

## Decoding Patterns

### Pattern 1: Decode to Known Type

```go
// You have the exact type
var user User
err := beve.Unmarshal(data, &user)
if err != nil {
    log.Fatal(err)
}

// Access fields
fmt.Println(user.Username) // "alice"
```

### Pattern 2: Decode to Interface

```go
// Type is unknown at compile time
var result interface{}
err := beve.Unmarshal(data, &result)

// Type assert to access
if user, ok := result.(map[string]interface{}); ok {
    fmt.Println(user["username"]) // "alice"
}
```

### Pattern 3: Decode Slice

```go
var users []User
err := beve.Unmarshal(data, &users)

// Iterate
for _, user := range users {
    fmt.Printf("%s <%s>\n", user.Username, user.Email)
}
```

### Pattern 4: Decode Map

```go
var config map[string]interface{}
err := beve.Unmarshal(data, &config)

// Access with type assertions
if host, ok := config["host"].(string); ok {
    fmt.Println("Host:", host)
}

if port, ok := config["port"].(int64); ok {
    fmt.Println("Port:", port)
}
```

### Pattern 5: Partial Decoding

```go
// Only decode fields you need
type UserSummary struct {
    ID       int64  `beve:"id"`
    Username string `beve:"username"`
    // Email is in data but not decoded
}

var summary UserSummary
err := beve.Unmarshal(userData, &summary)
```

---

## Struct Tags

### Basic Syntax

```go
type T struct {
    Field1 string `beve:"field1"`        // Field name in binary
    Field2 string `beve:"custom_name"`   // Custom name
    Field3 string `beve:"-"`             // Skip this field
    Field4 string `beve:"field4,omitempty"` // Omit if zero value
}
```

### Tag Options

**`omitempty`** - Omit field if zero value:

```go
type User struct {
    Name string `beve:"name"`
    Bio  string `beve:"bio,omitempty"` // Omitted if ""
    Age  int    `beve:"age,omitempty"` // Omitted if 0
}

user1 := User{Name: "Alice", Bio: "Engineer"}
// Encoded: {name: "Alice", bio: "Engineer"}

user2 := User{Name: "Bob"}
// Encoded: {name: "Bob"} (bio omitted)
```

**`-`** - Skip field entirely:

```go
type User struct {
    Name     string `beve:"name"`
    Password string `beve:"-"` // Never encoded or decoded
}

user := User{Name: "Alice", Password: "secret123"}
data, _ := beve.Marshal(user)
// Password is NOT in the binary data
```

**Custom names**:

```go
type APIResponse struct {
    UserID   int64  `beve:"user_id"`   // Snake case for API
    FullName string `beve:"full_name"` // Multiple words
}
```

### Tag Inheritance

```go
// Embedded structs inherit tags
type Base struct {
    ID        int64  `beve:"id"`
    CreatedAt int64  `beve:"created_at"`
}

type User struct {
    Base                // Inherits id and created_at
    Name string `beve:"name"`
}

// Encoded as: {id, created_at, name}
```

---

## Working with Maps

### String Keys (Most Common)

```go
// Create map
data := map[string]interface{}{
    "name":  "Alice",
    "age":   30,
    "roles": []string{"admin", "user"},
}

// Encode
binary, _ := beve.Marshal(data)

// Decode
var result map[string]interface{}
beve.Unmarshal(binary, &result)
```

### Integer Keys

```go
// Map with int keys
lookup := map[int]string{
    1: "one",
    2: "two",
    3: "three",
}

data, _ := beve.Marshal(lookup)

var decoded map[int]string
beve.Unmarshal(data, &decoded)
```

### Nested Maps

```go
// Deeply nested structure
config := map[string]interface{}{
    "database": map[string]interface{}{
        "host": "localhost",
        "port": 5432,
        "credentials": map[string]string{
            "username": "admin",
            "password": "secret",
        },
    },
    "cache": map[string]interface{}{
        "enabled": true,
        "ttl":     3600,
    },
}

data, _ := beve.Marshal(config)
```

### Safe Map Access

```go
var config map[string]interface{}
beve.Unmarshal(data, &config)

// Safe access with type assertions
if db, ok := config["database"].(map[string]interface{}); ok {
    if host, ok := db["host"].(string); ok {
        fmt.Println("Host:", host)
    }
}

// Helper function for deep access
func getStr(m map[string]interface{}, key string) (string, bool) {
    val, ok := m[key]
    if !ok {
        return "", false
    }
    str, ok := val.(string)
    return str, ok
}
```

---

## Slices and Arrays

### Dynamic Slices

```go
// Slice of primitives
numbers := []int{1, 2, 3, 4, 5}
data, _ := beve.Marshal(numbers)

var decoded []int
beve.Unmarshal(data, &decoded)
```

### Fixed Arrays

```go
// Array (fixed size)
var fixed [5]int = [5]int{1, 2, 3, 4, 5}
data, _ := beve.Marshal(fixed)

var decoded [5]int
beve.Unmarshal(data, &decoded)
```

### Slice of Structs

```go
type Point struct {
    X float64 `beve:"x"`
    Y float64 `beve:"y"`
}

points := []Point{
    {X: 0, Y: 0},
    {X: 10, Y: 20},
    {X: 30, Y: 40},
}

data, _ := beve.Marshal(points)
```

### Multidimensional Slices

```go
// 2D slice
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

data, _ := beve.Marshal(matrix)

var decoded [][]int
beve.Unmarshal(data, &decoded)
```

### Empty Slices

```go
// Empty slice vs nil
var empty []int = []int{}     // Empty slice
var null []int = nil          // Nil slice

data1, _ := beve.Marshal(empty) // Encoded as empty array []
data2, _ := beve.Marshal(null)  // Encoded as null

// Both decode to empty slice
var result1, result2 []int
beve.Unmarshal(data1, &result1) // []int{}
beve.Unmarshal(data2, &result2) // []int{}
```

---

## Pointers and Nil Values

### Pointer Fields

```go
type User struct {
    Name  string  `beve:"name"`
    Email *string `beve:"email,omitempty"` // Optional field
}

// With value
email := "alice@example.com"
user1 := User{Name: "Alice", Email: &email}

// Without value (nil)
user2 := User{Name: "Bob", Email: nil}

data1, _ := beve.Marshal(user1) // {name: "Alice", email: "alice@example.com"}
data2, _ := beve.Marshal(user2) // {name: "Bob"} (email omitted)
```

### Nil Detection

```go
var decoded User
beve.Unmarshal(data, &decoded)

if decoded.Email != nil {
    fmt.Println("Email:", *decoded.Email)
} else {
    fmt.Println("No email provided")
}
```

### Pointer to Struct

```go
// Encode pointer
user := &User{Name: "Alice"}
data, _ := beve.Marshal(user)

// Decode to pointer
var decoded *User
beve.Unmarshal(data, &decoded)
```

---

## Interface Types

### Empty Interface (any)

```go
// Can hold any value
var any interface{}

any = 42
data, _ := beve.Marshal(any)

any = "hello"
data, _ := beve.Marshal(any)

any = []int{1, 2, 3}
data, _ := beve.Marshal(any)
```

### Type Switching

```go
var result interface{}
beve.Unmarshal(data, &result)

switch v := result.(type) {
case string:
    fmt.Println("String:", v)
case int64:
    fmt.Println("Integer:", v)
case map[string]interface{}:
    fmt.Println("Object:", v)
case []interface{}:
    fmt.Println("Array:", v)
default:
    fmt.Println("Unknown type:", v)
}
```

---

## Custom Types

### Method 1: Type Alias

```go
// Define custom type
type UserID int64

type User struct {
    ID   UserID `beve:"id"`
    Name string `beve:"name"`
}

// Works automatically
user := User{ID: UserID(123), Name: "Alice"}
data, _ := beve.Marshal(user)
```

### Method 2: Custom Marshal/Unmarshal

```go
import "time"

// Custom time encoding
type Timestamp int64

func (t Timestamp) MarshalBEVE() ([]byte, error) {
    return beve.Marshal(int64(t))
}

func (t *Timestamp) UnmarshalBEVE(data []byte) error {
    var v int64
    err := beve.Unmarshal(data, &v)
    *t = Timestamp(v)
    return err
}

type Event struct {
    Name      string    `beve:"name"`
    Timestamp Timestamp `beve:"timestamp"`
}
```

### Method 3: Binary Marshaler

```go
// Implement encoding.BinaryMarshaler
type UUID [16]byte

func (u UUID) MarshalBinary() ([]byte, error) {
    return u[:], nil
}

func (u *UUID) UnmarshalBinary(data []byte) error {
    if len(data) != 16 {
        return errors.New("invalid UUID length")
    }
    copy(u[:], data)
    return nil
}

// BEVE automatically uses BinaryMarshaler
type User struct {
    ID   UUID   `beve:"id"`
    Name string `beve:"name"`
}
```

---

## Best Practices

### 1. Always Use Struct Tags

```go
// ❌ Bad: No tags
type User struct {
    Name string
    Age  int
}

// ✅ Good: Explicit tags
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}
```

**Why**: Explicit tags prevent breakage if you rename fields.

### 2. Use omitempty for Optional Fields

```go
// ✅ Good: Save space
type User struct {
    Name     string `beve:"name"`
    Email    string `beve:"email,omitempty"`
    Bio      string `beve:"bio,omitempty"`
    Website  string `beve:"website,omitempty"`
}
```

**Benefit**: Smaller payloads when fields are not set.

### 3. Validate After Unmarshal

```go
var user User
err := beve.Unmarshal(data, &user)
if err != nil {
    return err
}

// Validate
if user.Name == "" {
    return errors.New("name is required")
}
if user.Age < 0 || user.Age > 150 {
    return errors.New("invalid age")
}
```

### 4. Use Pointers for Large Structs

```go
// ✅ Good: Avoid copying
func ProcessUser(user *User) error {
    data, err := beve.Marshal(user) // No copy
    return err
}

// ❌ Bad: Copies entire struct
func ProcessUser(user User) error {
    data, err := beve.Marshal(user) // Struct copied
    return err
}
```

### 5. Reuse Buffers for Performance

```go
import "github.com/beve-org/beve-go/core"

// ✅ Good: Reuse encoder
enc := core.GetEncoderFromPool()
defer core.PutEncoderToPool(enc)

for _, item := range items {
    enc.Buf.Reset()
    enc.Encode(item)
    // Process enc.Buf.Bytes()
}
```

### 6. Handle Errors Properly

```go
// ✅ Good: Check all errors
data, err := beve.Marshal(user)
if err != nil {
    log.Printf("marshal failed: %v", err)
    return fmt.Errorf("encoding user: %w", err)
}

// ❌ Bad: Ignoring errors
data, _ := beve.Marshal(user) // Don't do this!
```

### 7. Version Your Data Structures

```go
// ✅ Good: Include version for evolution
type Message struct {
    Version int    `beve:"version"`
    Type    string `beve:"type"`
    Data    []byte `beve:"data"`
}

// Decode and check version
var msg Message
beve.Unmarshal(data, &msg)

if msg.Version != 1 {
    return fmt.Errorf("unsupported version: %d", msg.Version)
}
```

### 8. Use Type-Safe Enums

```go
// ✅ Good: Type-safe enum
type Status int

const (
    StatusPending Status = iota
    StatusActive
    StatusInactive
)

type User struct {
    Name   string `beve:"name"`
    Status Status `beve:"status"`
}
```

---

## Common Pitfalls

### Pitfall 1: Forgetting & when Unmarshaling

```go
// ❌ Wrong: Missing pointer
var user User
beve.Unmarshal(data, user) // Compile error!

// ✅ Right: Use pointer
var user User
beve.Unmarshal(data, &user)
```

### Pitfall 2: Modifying Data After Marshal

```go
data, _ := beve.Marshal(user)
// ❌ Don't modify the input after marshaling
user.Name = "Changed" // This won't affect 'data'
```

### Pitfall 3: Assuming Field Order

```go
// Field order in binary is not guaranteed
// Always use struct tags for decoding
```

### Pitfall 4: Large Allocations in Loop

```go
// ❌ Bad: Allocates every iteration
for _, item := range items {
    data, _ := beve.Marshal(item)
    // Process data
}

// ✅ Good: Reuse encoder
enc := core.GetEncoderFromPool()
defer core.PutEncoderToPool(enc)

for _, item := range items {
    enc.Buf.Reset()
    enc.Encode(item)
    // Process enc.Buf.Bytes()
}
```

---

## Complete Example: REST API

```go
package main

import (
    "log"
    "net/http"
    beve "github.com/beve-org/beve-go"
)

type User struct {
    ID       int64  `beve:"id"`
    Username string `beve:"username"`
    Email    string `beve:"email,omitempty"`
}

type ErrorResponse struct {
    Error string `beve:"error"`
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Mock user
    user := User{
        ID:       123,
        Username: "alice",
        Email:    "alice@example.com",
    }

    // Encode to BEVE
    data, err := beve.Marshal(user)
    if err != nil {
        sendError(w, err)
        return
    }

    // Send response
    w.Header().Set("Content-Type", "application/beve")
    w.Write(data)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
    // Read request body
    var user User
    err := beve.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        sendError(w, err)
        return
    }

    // Validate
    if user.Username == "" {
        sendError(w, fmt.Errorf("username required"))
        return
    }

    // Process user (save to DB, etc.)
    user.ID = 124 // Assign ID

    // Send response
    data, _ := beve.Marshal(user)
    w.Header().Set("Content-Type", "application/beve")
    w.WriteHeader(http.StatusCreated)
    w.Write(data)
}

func sendError(w http.ResponseWriter, err error) {
    resp := ErrorResponse{Error: err.Error()}
    data, _ := beve.Marshal(resp)
    w.Header().Set("Content-Type", "application/beve")
    w.WriteHeader(http.StatusBadRequest)
    w.Write(data)
}

func main() {
    http.HandleFunc("/user", handleGetUser)
    http.HandleFunc("/user/create", handleCreateUser)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## Next Steps

✅ **Basic Usage complete!** Continue with:

1. **[JSON Migration →](json-migration.md)** - Switch from encoding/json
2. **[User Guides →](../guides/encoding-decoding.md)** - Advanced features
3. **[Performance Tuning →](../guides/performance.md)** - Optimize for speed
4. **[Extensions →](../guides/extensions.md)** - Extension system

---

## Getting Help

- 📖 **Documentation**: [docs/INDEX.md](../INDEX.md)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/beve-org/beve-go/discussions)
- 🐛 **Issues**: [GitHub Issues](https://github.com/beve-org/beve-go/issues)

---

**Completion Time**: 10 minutes ✅  
**Next**: [JSON Migration →](json-migration.md)
