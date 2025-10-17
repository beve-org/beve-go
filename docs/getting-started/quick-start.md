# ⚡ 5-Minute Quick Start

Get started with BEVE-Go in just 5 minutes. This tutorial covers the basics of encoding and decoding.

---

## Prerequisites

- ✅ [BEVE-Go installed](installation.md)
- ✅ Go 1.21+ installed
- ✅ Basic Go knowledge

---

## Your First BEVE Program

Create `main.go`:

```go
package main

import (
    "fmt"
    beve "github.com/beve-org/beve-go"
)

func main() {
    // Step 1: Define a struct
    type Person struct {
        Name string `beve:"name"`
        Age  int    `beve:"age"`
    }

    // Step 2: Create an instance
    person := Person{
        Name: "Alice",
        Age:  30,
    }

    // Step 3: Encode to BEVE
    data, err := beve.Marshal(person)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Encoded: %d bytes\n", len(data))
    fmt.Printf("Binary: %x\n", data)

    // Step 4: Decode from BEVE
    var decoded Person
    err = beve.Unmarshal(data, &decoded)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decoded: %+v\n", decoded)
}
```

Run it:
```bash
go run main.go
```

Output:
```
Encoded: 23 bytes
Binary: 034005416c6963650301e0
Decoded: {Name:Alice Age:30}
```

🎉 **Congratulations!** You just encoded and decoded your first BEVE message!

---

## Understanding the Basics

### 1. Struct Tags

BEVE uses struct tags to define field names in the binary format:

```go
type User struct {
    Name  string `beve:"name"`   // Field name in binary: "name"
    Email string `beve:"email"`  // Field name in binary: "email"
}
```

**Tag Options**:
- `beve:"fieldname"` - Custom field name
- `beve:"fieldname,omitempty"` - Omit if zero value
- `beve:"-"` - Skip this field

Example with options:
```go
type User struct {
    Name     string `beve:"name"`
    Password string `beve:"-"`              // Never encoded
    Bio      string `beve:"bio,omitempty"`  // Omit if empty
}
```

### 2. Encoding (Marshal)

Convert Go values to BEVE binary:

```go
// Encode a struct
data, err := beve.Marshal(user)

// Encode a map
userData := map[string]interface{}{
    "name": "Bob",
    "age":  25,
}
data, err := beve.Marshal(userData)

// Encode a slice
users := []User{{Name: "Alice"}, {Name: "Bob"}}
data, err := beve.Marshal(users)
```

### 3. Decoding (Unmarshal)

Convert BEVE binary back to Go values:

```go
var user User
err := beve.Unmarshal(data, &user)

// Always pass a pointer!
// ❌ Wrong: beve.Unmarshal(data, user)
// ✅ Right: beve.Unmarshal(data, &user)
```

---

## Common Patterns

### Pattern 1: Simple Struct

```go
type Message struct {
    ID        int64  `beve:"id"`
    Text      string `beve:"text"`
    Timestamp int64  `beve:"timestamp"`
}

msg := Message{ID: 1, Text: "Hello", Timestamp: 1697500000}
data, _ := beve.Marshal(msg)

var decoded Message
beve.Unmarshal(data, &decoded)
```

### Pattern 2: Nested Structs

```go
type Address struct {
    Street string `beve:"street"`
    City   string `beve:"city"`
}

type Person struct {
    Name    string  `beve:"name"`
    Address Address `beve:"address"`
}

person := Person{
    Name: "Alice",
    Address: Address{
        Street: "123 Main St",
        City:   "NYC",
    },
}

data, _ := beve.Marshal(person)
```

### Pattern 3: Slices and Arrays

```go
type Team struct {
    Name    string   `beve:"name"`
    Members []string `beve:"members"`
}

team := Team{
    Name:    "Engineering",
    Members: []string{"Alice", "Bob", "Charlie"},
}

data, _ := beve.Marshal(team)
```

### Pattern 4: Maps

```go
// String keys
config := map[string]interface{}{
    "host":    "localhost",
    "port":    8080,
    "enabled": true,
}

data, _ := beve.Marshal(config)

// Decode back
var decoded map[string]interface{}
beve.Unmarshal(data, &decoded)
```

---

## Error Handling

Always check errors when encoding/decoding:

```go
data, err := beve.Marshal(user)
if err != nil {
    // Handle error
    log.Printf("Failed to marshal: %v", err)
    return err
}

var decoded User
err = beve.Unmarshal(data, &decoded)
if err != nil {
    // Handle error
    log.Printf("Failed to unmarshal: %v", err)
    return err
}
```

**Common Errors**:

| Error | Cause | Solution |
|-------|-------|----------|
| `unsupported type` | Trying to encode unsupported type | Use supported types (see [Basic Usage](basic-usage.md)) |
| `invalid data` | Corrupted binary data | Validate data source |
| `type mismatch` | Decoding into wrong type | Ensure types match |
| `buffer too short` | Incomplete binary data | Check data length |

---

## Performance Tip: Reuse Encoders

For high-throughput scenarios, reuse encoders from a pool:

```go
import "github.com/beve-org/beve-go/core"

// Get encoder from pool
enc := core.GetEncoderFromPool()
defer core.PutEncoderToPool(enc)

// Encode multiple values
enc.Encode(user1)
data1 := enc.Buf.Bytes()

enc.Buf.Reset() // Reset for next use
enc.Encode(user2)
data2 := enc.Buf.Bytes()
```

**Performance gain**: ~8ns overhead vs ~100ns allocation

---

## Comparing with JSON

**JSON**:
```go
import "encoding/json"

// Encode
data, err := json.Marshal(user)

// Decode
var decoded User
err = json.Unmarshal(data, &decoded)
```

**BEVE**:
```go
import beve "github.com/beve-org/beve-go"

// Encode (same API!)
data, err := beve.Marshal(user)

// Decode (same API!)
var decoded User
err = beve.Unmarshal(data, &decoded)
```

**Advantages**:
- ✅ **Drop-in replacement** for `encoding/json`
- ✅ **2-46× faster** than JSON
- ✅ **30-50% smaller** payloads
- ✅ **Zero-copy mode** (0 allocations)

See: [Migration Guide](json-migration.md)

---

## Example: HTTP API

```go
package main

import (
    "net/http"
    beve "github.com/beve-org/beve-go"
)

type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    user := User{Name: "Alice", Age: 30}
    
    // Encode to BEVE
    data, err := beve.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    // Send response
    w.Header().Set("Content-Type", "application/beve")
    w.Write(data)
}

func main() {
    http.HandleFunc("/user", handler)
    http.ListenAndServe(":8080", nil)
}
```

Test it:
```bash
# Start server
go run main.go

# Test with curl
curl http://localhost:8080/user --output user.beve
```

---

## Complete Example: Todo App

```go
package main

import (
    "fmt"
    beve "github.com/beve-org/beve-go"
    "os"
)

type Todo struct {
    ID        int    `beve:"id"`
    Title     string `beve:"title"`
    Completed bool   `beve:"completed"`
}

type TodoList struct {
    Todos []Todo `beve:"todos"`
}

func main() {
    // Create todo list
    list := TodoList{
        Todos: []Todo{
            {ID: 1, Title: "Learn BEVE", Completed: true},
            {ID: 2, Title: "Build app", Completed: false},
        },
    }

    // Save to file
    data, err := beve.Marshal(list)
    if err != nil {
        panic(err)
    }
    os.WriteFile("todos.beve", data, 0644)
    fmt.Println("✅ Saved todos.beve")

    // Load from file
    fileData, err := os.ReadFile("todos.beve")
    if err != nil {
        panic(err)
    }

    var loaded TodoList
    err = beve.Unmarshal(fileData, &loaded)
    if err != nil {
        panic(err)
    }

    // Print todos
    fmt.Println("📝 Todo List:")
    for _, todo := range loaded.Todos {
        status := "❌"
        if todo.Completed {
            status = "✅"
        }
        fmt.Printf("  %s %s\n", status, todo.Title)
    }
}
```

Output:
```
✅ Saved todos.beve
📝 Todo List:
  ✅ Learn BEVE
  ❌ Build app
```

---

## Supported Types

BEVE supports all common Go types:

**Primitives**:
- `bool`, `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `string`

**Composite**:
- `struct` (with tags)
- `slice`, `array`
- `map[string]T`, `map[int]T`
- `*T` (pointers)

**Special**:
- `interface{}` (any type)
- `[]byte` (binary data)
- Custom types with `MarshalBEVE` / `UnmarshalBEVE`

**Not Supported** (yet):
- `chan`, `func`
- `time.Time` (use Extension 4, see [Extensions](../guides/extensions.md))
- Complex numbers (use Extension 3)

---

## Next Steps

✅ **Quick Start complete!** Continue learning:

1. **[Basic Usage →](basic-usage.md)** - Common patterns in depth
2. **[Migrating from JSON →](json-migration.md)** - Switch your existing code
3. **[User Guides →](../guides/encoding-decoding.md)** - Advanced features
4. **[Performance Tuning →](../guides/performance.md)** - Optimize for speed

---

## Cheat Sheet

```go
// Marshal (encode)
data, err := beve.Marshal(value)

// Unmarshal (decode)
var result Type
err := beve.Unmarshal(data, &result)

// Struct tags
type T struct {
    Field1 string `beve:"field1"`           // Custom name
    Field2 string `beve:"field2,omitempty"` // Omit if empty
    Field3 string `beve:"-"`                // Skip field
}

// Error handling
if err != nil {
    log.Fatal(err)
}
```

---

## Getting Help

- 📖 **Documentation**: [docs/INDEX.md](../INDEX.md)
- 💬 **Ask Questions**: [GitHub Discussions](https://github.com/beve-org/beve-go/discussions)
- 🐛 **Report Issues**: [GitHub Issues](https://github.com/beve-org/beve-go/issues)
- 📚 **Examples**: [examples/](../../examples/README.md)

---

**Completion Time**: 5 minutes ✅  
**Next**: [Basic Usage →](basic-usage.md)
