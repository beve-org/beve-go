# Extension 2: Typed Nested Array (Exponential Gains)

**Extension ID**: 2  
**Status**: ⚠️ Planned (v1.4)  
**Version**: BEVE v1.4+ (Future)  
**Expected Performance**: **87-93% smaller** for deeply nested data  

## Overview

### What is Typed Nested Array?

Extension 2 extends Extension 1 (Typed Object Array) to **nested structures**, providing exponential space savings for hierarchical data.

**Problem**: Standard BEVE repeats field names at **every nesting level**:

```json
[
  {
    "user": {
      "name": "Alice",
      "age": 30
    },
    "posts": [
      {"title": "Post 1", "views": 100},
      {"title": "Post 2", "views": 200}
    ]
  },
  // ... 99 more users with posts
]
```

**Repetition**:
- Level 0: `"user"`, `"posts"` (repeated 100×)
- Level 1: `"name"`, `"age"` (repeated 100×)
- Level 2: `"title"`, `"views"` (repeated 200×)

**Extension 2**: Hierarchical schema (stored once per level):

```
[Schema Level 0: "user", "posts"]
  [Schema Level 1 (user): "name", "age"]
  [Schema Level 2 (posts): "title", "views"]
[Values: nested data without field names]
```

### Benefits (Theoretical)

| Depth | Objects | Standard BEVE | Extension 2 | Savings |
|-------|---------|---------------|-------------|---------|
| D=1   | N=100   | 5.2 KB        | 2.7 KB      | **48%** |
| D=2   | N²=10K  | 520 KB        | 135 KB      | **74%** |
| D=3   | N³=1M   | 52 MB         | 6.8 MB      | **87%** |
| D=4   | N⁴=100M | 5.2 GB        | 337 MB      | **93%** |

**Formula**: `Savings = 1 - (1 / N^(D-1))`

**Result**: Exponential gains for deeply nested structures!

---

## Binary Format (Proposed)

### Structure

```
┌────────────────────────────────────────────────────────────┐
│ [0x96]          Extension 2 Header (1 byte)                │
├────────────────────────────────────────────────────────────┤
│ [Depth]         Nesting Depth (varint)                     │
│                 Number of schema levels                    │
├────────────────────────────────────────────────────────────┤
│                 SCHEMA SECTION                             │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Level 0 Schema:                                     │  │
│  │   [field_count] (varint)                           │  │
│  │   [field_0_name] [field_1_name] ...                │  │
│  │   [field_0_type] [field_1_type] ...                │  │
│  ├─────────────────────────────────────────────────────┤  │
│  │ Level 1 Schema (for nested objects):                │  │
│  │   [nested_field_index] (which field is nested)     │  │
│  │   [field_count] (varint)                           │  │
│  │   [field_0_name] [field_1_name] ...                │  │
│  │   [field_0_type] [field_1_type] ...                │  │
│  ├─────────────────────────────────────────────────────┤  │
│  │ Level D Schema: ...                                 │  │
│  └─────────────────────────────────────────────────────┘  │
├────────────────────────────────────────────────────────────┤
│ [Count]         Object Count (varint)                      │
├────────────────────────────────────────────────────────────┤
│                 VALUES SECTION                             │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Object 0:                                           │  │
│  │   [value_0] [value_1] ... [nested_array]           │  │
│  │   (Nested array uses Level 1 schema)               │  │
│  ├─────────────────────────────────────────────────────┤  │
│  │ Object N: ...                                       │  │
│  └─────────────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────────────┘
```

### Type Indicators

**Field Type Byte**:
```
Bits 0-2:  Base type (0=null, 1=number, 2=string, 3=object, 4=array, 5=bool)
Bit 3:     Is nested (uses child schema)
Bits 4-7:  Nested schema index (if bit 3 set)
```

### Example Encoding (Depth 2)

**Input**:

```json
[
  {
    "user": {"name": "Alice", "age": 30},
    "posts": [
      {"title": "Post 1", "views": 100}
    ]
  },
  {
    "user": {"name": "Bob", "age": 25},
    "posts": [
      {"title": "Post 2", "views": 200}
    ]
  }
]
```

**Binary Layout**:

```
Offset | Hex          | Description
-------|--------------|------------------------------------------
0x00   | 96           | Extension 2 header
0x01   | 02           | Depth: 2 levels
       |              |
       | --- SCHEMA LEVEL 0 ---
0x02   | 02           | Field count: 2 ("user", "posts")
0x03   | 04 'user'    | Field 0: "user"
0x08   | 05 'posts'   | Field 1: "posts"
0x0E   | 18           | Field 0 type: 0b00011000 (object, nested, schema #1)
0x0F   | 24           | Field 1 type: 0b00100100 (array, nested, schema #2)
       |              |
       | --- SCHEMA LEVEL 1 (user) ---
0x10   | 02           | Field count: 2 ("name", "age")
0x11   | 04 'name'    | Field 0: "name"
0x16   | 03 'age'     | Field 1: "age"
0x1A   | 02           | Field 0 type: string
0x1B   | 01           | Field 1 type: number
       |              |
       | --- SCHEMA LEVEL 2 (posts) ---
0x1C   | 02           | Field count: 2 ("title", "views")
0x1D   | 05 'title'   | Field 0: "title"
0x23   | 05 'views'   | Field 1: "views"
0x29   | 02           | Field 0 type: string
0x2A   | 01           | Field 1 type: number
       |              |
0x2B   | 02           | Object count: 2
       |              |
       | --- VALUES ---
0x2C   | [VALUES: "Alice", 30, [["Post 1", 100]]]
       | [VALUES: "Bob", 25, [["Post 2", 200]]]
```

**Size Comparison**:

| Format | Size | vs Extension 2 |
|--------|------|----------------|
| **Extension 2** | **~120 bytes** | Baseline |
| Extension 1 (flat) | ~150 bytes | +25% |
| Standard BEVE | ~280 bytes | +133% |
| JSON | ~220 bytes | +83% |

---

## API Usage (Proposed v1.4)

### Encoding Nested Arrays

**Automatic Detection**:

```go
type UserWithPosts struct {
    User  User   `beve:"user"`
    Posts []Post `beve:"posts"`
}

type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

type Post struct {
    Title string `beve:"title"`
    Views int    `beve:"views"`
}

users := []UserWithPosts{
    {
        User:  User{Name: "Alice", Age: 30},
        Posts: []Post{{Title: "Post 1", Views: 100}},
    },
    // ... 99 more
}

// Auto-detect nested structure (uses Extension 2)
data, err := beve.MarshalAuto(users)
```

**Explicit Nested Encoding**:

```go
// Force nested typed schema
opts := beve.MarshalOptions{
    UseTypedSchema:  true,
    UseNestedSchema: true, // Enable Extension 2
    MaxDepth:        16,   // Limit nesting depth
}

data, err := beve.MarshalWithOptions(users, opts)
```

### Decoding Nested Arrays

**Automatic Detection**:

```go
var users []UserWithPosts
err := beve.UnmarshalAuto(data, &users)
// Detects Extension 2 header, decodes nested structure
```

---

## Performance (Projected)

### Expected Benchmarks

**Depth 1** (100 objects, Extension 1 equivalent):

| Operation | Standard | Extension 2 | Improvement |
|-----------|----------|-------------|-------------|
| **Marshal** | 17,420 ns | ~6,500 ns | **2.7× faster** |
| **Unmarshal** | 24,150 ns | ~17,000 ns | **1.4× faster** |
| **Size** | 5,200 bytes | 2,700 bytes | **48% smaller** |

**Depth 2** (100 parent × 10 child = 1,000 nested objects):

| Operation | Standard | Extension 2 | Improvement |
|-----------|----------|-------------|-------------|
| **Marshal** | ~200 μs | ~50 μs | **4× faster** |
| **Unmarshal** | ~280 μs | ~100 μs | **2.8× faster** |
| **Size** | 52 KB | 13.5 KB | **74% smaller** |

**Depth 3** (100 × 10 × 10 = 10,000 nested objects):

| Operation | Standard | Extension 2 | Improvement |
|-----------|----------|-------------|-------------|
| **Size** | 520 KB | 68 KB | **87% smaller** |

---

## Use Cases

### When to Use

✅ **Use Extension 2 When**:
- Deeply nested data (depth ≥ 2)
- Large hierarchical structures (1000+ nested objects)
- API responses with nested lists
- Tree-like data (org charts, file systems)

❌ **Don't Use When**:
- Shallow nesting (depth = 1) - Use Extension 1 instead
- Heterogeneous nested objects (varying schemas)
- Small datasets (< 100 objects)

### Real-World Scenarios

**Scenario 1: Organization Chart**

```go
type Employee struct {
    Name       string     `beve:"name"`
    Title      string     `beve:"title"`
    Directs    []Employee `beve:"directs"` // Recursive nesting
}

// CEO → 10 VPs → 100 Directors → 1000 Employees
// Depth 4, 1,110 total employees

// Standard BEVE: ~5.2 MB (field names repeated at each level)
// Extension 2: ~337 KB (93% savings!)
```

**Scenario 2: File System Tree**

```go
type Directory struct {
    Name    string      `beve:"name"`
    Files   []File      `beve:"files"`
    SubDirs []Directory `beve:"subdirs"` // Nested directories
}

type File struct {
    Name string `beve:"name"`
    Size int64  `beve:"size"`
}

// Root → 100 dirs → 1000 files
// Depth 3

// Standard: 52 MB
// Extension 2: 6.8 MB (87% savings)
```

**Scenario 3: E-commerce Orders**

```go
type Order struct {
    ID      int           `beve:"id"`
    User    User          `beve:"user"`    // Nested user
    Items   []OrderItem   `beve:"items"`   // Nested items
}

type OrderItem struct {
    Product  Product `beve:"product"`  // Nested product
    Quantity int     `beve:"quantity"`
}

// 10,000 orders × 5 items avg = 50,000 nested objects
// Depth 3

// Standard: ~5.2 MB
// Extension 2: ~680 KB (87% savings)
```

---

## Implementation Challenges

### Schema Inference

**Challenge**: Detect nested schema automatically

```go
// Must analyze struct tags recursively
func inferNestedSchema(t reflect.Type) *NestedSchema {
    schema := &NestedSchema{Level: 0}
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        
        if field.Type.Kind() == reflect.Struct {
            // Nested object - create child schema
            childSchema := inferNestedSchema(field.Type)
            childSchema.Level = schema.Level + 1
            schema.Children = append(schema.Children, childSchema)
        } else if field.Type.Kind() == reflect.Slice {
            elemType := field.Type.Elem()
            if elemType.Kind() == reflect.Struct {
                // Nested array - create child schema
                childSchema := inferNestedSchema(elemType)
                childSchema.Level = schema.Level + 1
                schema.Children = append(schema.Children, childSchema)
            }
        }
    }
    
    return schema
}
```

### Recursive Encoding

**Challenge**: Encode nested values efficiently

```go
func encodeNestedValues(schema *NestedSchema, values []interface{}) []byte {
    var buf bytes.Buffer
    
    for _, value := range values {
        for i, field := range schema.Fields {
            fieldValue := getFieldValue(value, i)
            
            if schema.Children[i] != nil {
                // Nested field - recurse
                nestedData := encodeNestedValues(
                    schema.Children[i],
                    fieldValue.([]interface{}),
                )
                buf.Write(nestedData)
            } else {
                // Primitive field - encode directly
                beve.Marshal(fieldValue, &buf)
            }
        }
    }
    
    return buf.Bytes()
}
```

### Max Depth Limit

**Limitation**: Stack overflow risk for very deep nesting

```go
const MaxNestingDepth = 16 // Safety limit

func validateDepth(schema *NestedSchema) error {
    if schema.Level > MaxNestingDepth {
        return errors.New("nesting depth exceeds maximum (16)")
    }
    
    for _, child := range schema.Children {
        if err := validateDepth(child); err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## Best Practices (Future)

### Depth Selection

```go
// Analyze depth before encoding
depth := calculateDepth(data)

if depth >= 2 {
    // Use Extension 2 (exponential gains)
    opts.UseNestedSchema = true
} else if isHomogeneous(data) {
    // Use Extension 1 (linear gains)
    opts.UseTypedSchema = true
} else {
    // Use standard BEVE
    opts.UseTypedSchema = false
}
```

### Schema Caching

```go
// Cache nested schemas for reuse
var schemaCache sync.Map

func getOrCreateSchema(t reflect.Type) *NestedSchema {
    if cached, ok := schemaCache.Load(t); ok {
        return cached.(*NestedSchema)
    }
    
    schema := inferNestedSchema(t)
    schemaCache.Store(t, schema)
    return schema
}
```

---

## Migration from Extension 1

**Before** (Extension 1 - flat arrays):

```go
// Only top-level array uses typed schema
users := []User{...}
data, _ := beve.MarshalTyped(users)
// Nested fields still repeat names
```

**After** (Extension 2 - nested arrays):

```go
// Nested arrays also use typed schema
usersWithPosts := []UserWithPosts{...}
data, _ := beve.MarshalTyped(usersWithPosts)
// Field names stored once per nesting level (exponential savings)
```

---

## Roadmap

### v1.4 (Planned - Q1 2026)

- [ ] Extension 2 implementation
- [ ] Nested schema inference
- [ ] Recursive encoding/decoding
- [ ] Depth limit (16 levels)
- [ ] Schema caching

### v1.5 (Future)

- [ ] Schema compression (deduplicate common field names)
- [ ] Sparse object handling (null value optimization)
- [ ] Circular reference detection

---

## Summary

**Extension 2 will provide** (when implemented):
- ✅ **74-93% smaller** for deeply nested data (depth 2-4)
- ✅ **2-4× faster** marshal (less data to write)
- ✅ **1.4-2.8× faster** unmarshal (schema cached)
- ✅ **Exponential gains** with nesting depth
- ⚠️ **Depth limit** (16 levels max)
- ⚠️ **Homogeneous only** (same schema per level)

**Best for** (future):
- Organization charts (depth 3-4)
- File systems (depth 3-5)
- E-commerce orders (depth 2-3)
- Nested API responses (depth 2-3)

**Status**: Planned for BEVE v1.4 (Q1 2026)

---

**Last Updated**: October 17, 2025  
**BEVE Version**: v1.3.0 (Extension 2 not yet implemented)  
**Target Release**: v1.4.0 (Q1 2026)
