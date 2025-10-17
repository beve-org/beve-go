# 🏷️ Struct Tags Complete Reference

Everything you need to know about BEVE struct tags for precise control over encoding.

**Reading Time**: 15 minutes  
**Level**: Intermediate  
**Prerequisites**: [Basic Usage](../getting-started/basic-usage.md)

---

## Table of Contents

1. [Tag Basics](#tag-basics)
2. [Tag Options](#tag-options)
3. [Field Visibility](#field-visibility)
4. [Tag Inheritance](#tag-inheritance)
5. [Advanced Patterns](#advanced-patterns)
6. [Tag Validation](#tag-validation)
7. [Best Practices](#best-practices)
8. [Common Issues](#common-issues)

---

## Tag Basics

### Basic Syntax

BEVE uses struct tags to control field encoding:

```go
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}
```

**Encoded as**:
```json
{
  "name": "Alice",
  "age": 30
}
```

### Default Behavior

Without tags, BEVE uses field names as-is:

```go
type User struct {
    Name string  // Encoded as "Name" (capital N)
    Age  int     // Encoded as "Age"
}
```

**Why use tags?**
- Control field names (camelCase, snake_case, etc.)
- Match external APIs (JSON compatibility)
- Skip private fields
- Mark optional fields

### Tag Format

```go
`beve:"name,option1,option2"`
```

- **name**: Field name in encoded output
- **options**: Comma-separated encoding options

**Examples**:
```go
`beve:"username"`              // Name only
`beve:"email,omitempty"`       // Name + option
`beve:",omitempty"`            // Default name + option
`beve:"-"`                     // Skip field
```

---

## Tag Options

### 1. omitempty

Skip field if it has zero value:

```go
type User struct {
    Name  string `beve:"name"`
    Email string `beve:"email,omitempty"`
    Age   int    `beve:"age,omitempty"`
}
```

**Encoding**:
```go
user1 := User{Name: "Alice", Email: "alice@example.com", Age: 30}
// Output: {name: "Alice", email: "alice@example.com", age: 30}

user2 := User{Name: "Bob"}
// Output: {name: "Bob"}
// email and age omitted (zero values)
```

**Zero Values**:
| Type | Zero Value |
|------|------------|
| `string` | `""` (empty string) |
| `int`, `int32`, `int64` | `0` |
| `uint`, `uint32`, `uint64` | `0` |
| `float32`, `float64` | `0.0` |
| `bool` | `false` |
| `pointer` | `nil` |
| `slice` | `nil` |
| `map` | `nil` |
| `struct` | All fields zero |

**Benefits**:
- **Smaller payloads**: Skip unnecessary fields
- **Bandwidth savings**: 10-30% smaller for sparse data
- **Optional fields**: Natural optional field pattern

### 2. `-` (Skip Field)

Always skip field (never encoded):

```go
type User struct {
    Name     string `beve:"name"`
    Password string `beve:"-"`  // Never encoded!
    token    string             // Private field (also skipped)
}
```

**Encoding**:
```go
user := User{Name: "Alice", Password: "secret123"}
// Output: {name: "Alice"}
// Password is never encoded
```

**Use Cases**:
- Sensitive data (passwords, tokens)
- Computed fields (not persisted)
- Internal state (caches, locks)

### 3. string (String Encoding)

Force numeric types to encode as strings:

```go
type Product struct {
    ID    int64  `beve:"id,string"`
    Price float64 `beve:"price,string"`
}
```

**Encoding**:
```go
product := Product{ID: 123456789, Price: 19.99}
// Output: {id: "123456789", price: "19.99"}
// Instead of: {id: 123456789, price: 19.99}
```

**Use Cases**:
- JavaScript number limits (safe integers: ±2^53)
- Financial data (avoid float precision loss)
- External API requirements

**⚠️ Warning**: Slightly slower (string conversion overhead)

### 4. inline (Flatten Embedded Struct)

Promote embedded struct fields to parent level:

```go
type Address struct {
    Street string `beve:"street"`
    City   string `beve:"city"`
}

type User struct {
    Name    string  `beve:"name"`
    Address Address `beve:"address,inline"`
}
```

**Without inline**:
```json
{
  "name": "Alice",
  "address": {
    "street": "123 Main St",
    "city": "NYC"
  }
}
```

**With inline**:
```json
{
  "name": "Alice",
  "street": "123 Main St",
  "city": "NYC"
}
```

**Use Cases**:
- Flatten nested structures
- Match flat database schemas
- Simplify API responses

---

## Field Visibility

### Exported vs Unexported Fields

Go visibility rules apply:

```go
type User struct {
    Name     string `beve:"name"`      // ✅ Exported: Encoded
    age      int    `beve:"age"`       // ❌ Unexported: Skipped
    Email    string `beve:"email"`     // ✅ Exported: Encoded
    password string `beve:"-"`         // ❌ Unexported: Skipped
}
```

**Rule**: Only exported fields (capital first letter) are encoded.

### Private Field Encoding

**Problem**: Can't encode private fields from other packages

```go
// Package A
package models

type User struct {
    name string // Private (can't encode outside package)
}
```

**Solution 1**: Export field

```go
type User struct {
    Name string `beve:"name"`
}
```

**Solution 2**: Custom marshaler

```go
func (u User) MarshalBEVE() ([]byte, error) {
    return beve.Marshal(map[string]interface{}{
        "name": u.name, // Access private field inside package
    })
}
```

---

## Tag Inheritance

### Embedded Structs

Embedded struct fields are promoted:

```go
type Person struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

type Employee struct {
    Person                       // Embedded
    Company string `beve:"company"`
}
```

**Encoding**:
```go
emp := Employee{
    Person:  Person{Name: "Alice", Age: 30},
    Company: "ACME Corp",
}
// Output: {name: "Alice", age: 30, company: "ACME Corp"}
// Person fields promoted to Employee level
```

### Embedded Struct Tags

Embedded struct can override parent tags:

```go
type Person struct {
    Name string `beve:"name"`
}

type Employee struct {
    Person `beve:"person"` // Override: Don't promote
    ID     int `beve:"id"`
}
```

**Encoding**:
```go
emp := Employee{Person: Person{Name: "Alice"}, ID: 123}
// Output: {person: {name: "Alice"}, id: 123}
// Person not promoted (has tag)
```

### Tag Precedence

When fields conflict, first declaration wins:

```go
type A struct {
    Name string `beve:"name"`
}

type B struct {
    Name string `beve:"name"`
}

type C struct {
    A           // First embedded
    B           // Second embedded
    Name string `beve:"name"` // Direct field
}
```

**Encoding**:
```go
c := C{
    A:    A{Name: "A"},
    B:    B{Name: "B"},
    Name: "C",
}
// Output: {name: "C"}
// Direct field takes precedence
```

**Precedence Order**:
1. Direct fields
2. First embedded struct
3. Second embedded struct
4. ... (declaration order)

---

## Advanced Patterns

### Pattern 1: Optional Fields with Pointers

Use pointers to distinguish between "missing" and "zero value":

```go
type User struct {
    Name  string `beve:"name"`
    Age   *int   `beve:"age,omitempty"`
    Email *string `beve:"email,omitempty"`
}
```

**Encoding**:
```go
// Age provided (zero is valid)
age := 0
user1 := User{Name: "Alice", Age: &age}
// Output: {name: "Alice", age: 0}

// Age not provided (nil)
user2 := User{Name: "Bob", Age: nil}
// Output: {name: "Bob"}
// age omitted (nil pointer)
```

**Use Cases**:
- Distinguish `0` vs "not set"
- Partial updates (PATCH requests)
- Optional configuration

### Pattern 2: Multiple Tag Sets

Support both JSON and BEVE:

```go
type User struct {
    Name  string `json:"username" beve:"name"`
    Email string `json:"email"    beve:"email"`
}
```

**Usage**:
```go
// JSON encoding
jsonData, _ := json.Marshal(user)
// Output: {"username": "Alice", "email": "..."}

// BEVE encoding
beveData, _ := beve.Marshal(user)
// Output: {name: "Alice", email: "..."}
```

**Benefits**:
- Gradual migration (JSON → BEVE)
- Multiple API versions
- Client compatibility

### Pattern 3: Versioned Structs

Handle schema evolution:

```go
// V1 (initial version)
type UserV1 struct {
    Name string `beve:"name"`
}

// V2 (add email)
type UserV2 struct {
    Name  string `beve:"name"`
    Email string `beve:"email,omitempty"` // Optional for backward compat
}

// V3 (rename email)
type UserV3 struct {
    Name       string `beve:"name"`
    Email      string `beve:"email,omitempty"`
    EmailAddr  string `beve:"email_address,omitempty"` // New field
}
```

**Backward Compatibility**:
```go
// V1 data can unmarshal to V2
v1Data := beve.Marshal(UserV1{Name: "Alice"})

var v2 User
beve.Unmarshal(v1Data, &v2)
// v2.Name = "Alice"
// v2.Email = "" (missing field → zero value)
```

### Pattern 4: Conditional Encoding

Encode different fields based on context:

```go
type User struct {
    Name     string `beve:"name"`
    Email    string `beve:"email"`
    Password string `beve:"-"` // Never encoded
}

// Custom marshaler for public API
func (u User) MarshalPublic() ([]byte, error) {
    return beve.Marshal(struct {
        Name  string `beve:"name"`
        Email string `beve:"email"`
    }{
        Name:  u.Name,
        Email: u.Email,
    })
}

// Custom marshaler for internal API
func (u User) MarshalInternal() ([]byte, error) {
    return beve.Marshal(u) // All fields except Password
}
```

### Pattern 5: Nested Tags

Complex nested structures:

```go
type Company struct {
    Name      string   `beve:"name"`
    Employees []Employee `beve:"employees,omitempty"`
}

type Employee struct {
    Person                      // Embedded with promotion
    Department string `beve:"department"`
}

type Person struct {
    Name string `beve:"name"`
    Age  int    `beve:"age,omitempty"`
}
```

**Encoding**:
```go
company := Company{
    Name: "ACME",
    Employees: []Employee{
        {Person: Person{Name: "Alice", Age: 30}, Department: "Eng"},
        {Person: Person{Name: "Bob"}, Department: "Sales"},
    },
}
// Output:
// {
//   name: "ACME",
//   employees: [
//     {name: "Alice", age: 30, department: "Eng"},
//     {name: "Bob", department: "Sales"}
//   ]
// }
```

---

## Tag Validation

### Common Tag Errors

#### 1. Invalid Tag Syntax

```go
// ❌ Wrong: Space after comma
type User struct {
    Name string `beve:"name, omitempty"`
}

// ✅ Correct: No spaces
type User struct {
    Name string `beve:"name,omitempty"`
}
```

#### 2. Duplicate Field Names

```go
// ❌ Wrong: Both fields encode as "id"
type User struct {
    UserID   int `beve:"id"`
    PersonID int `beve:"id"` // Duplicate!
}

// ✅ Correct: Unique names
type User struct {
    UserID   int `beve:"user_id"`
    PersonID int `beve:"person_id"`
}
```

#### 3. Reserved Names

```go
// ❌ Avoid: May conflict with internal fields
type User struct {
    Type string `beve:"type"` // Reserved in some contexts
}

// ✅ Better: Prefix or rename
type User struct {
    UserType string `beve:"user_type"`
}
```

### Validation Tools

**Static Analysis** (go vet):

```bash
go vet ./...
# Checks for:
# - Invalid tag syntax
# - Duplicate field names
# - Type mismatches
```

**Runtime Validation**:

```go
func validateStructTags(v interface{}) error {
    t := reflect.TypeOf(v)
    if t.Kind() != reflect.Struct {
        return errors.New("not a struct")
    }
    
    seen := make(map[string]bool)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("beve")
        
        if tag == "" || tag == "-" {
            continue
        }
        
        // Extract field name (before first comma)
        name := strings.Split(tag, ",")[0]
        
        if seen[name] {
            return fmt.Errorf("duplicate field name: %s", name)
        }
        seen[name] = true
    }
    
    return nil
}
```

---

## Best Practices

### 1. Always Use Tags

**❌ Avoid**:
```go
type User struct {
    Name string  // Encodes as "Name" (capital N)
    Age  int     // Encodes as "Age"
}
```

**✅ Prefer**:
```go
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}
```

**Why?**
- Explicit field names (less brittle)
- API consistency (lowercase convention)
- Easier migration (change struct without breaking API)

### 2. Use omitempty for Optional Fields

**❌ Avoid**:
```go
type User struct {
    Name  string `beve:"name"`
    Email string `beve:"email"` // Always encoded (even if empty)
}
```

**✅ Prefer**:
```go
type User struct {
    Name  string `beve:"name"`
    Email string `beve:"email,omitempty"` // Omit if empty
}
```

**Benefits**:
- **Smaller payloads**: Skip empty fields
- **Backward compat**: Add fields without breaking old clients
- **Optional pattern**: Natural way to express optional data

### 3. Use Pointers for Nullable Fields

**❌ Avoid**:
```go
type User struct {
    Age int `beve:"age"` // 0 = zero or missing?
}
```

**✅ Prefer**:
```go
type User struct {
    Age *int `beve:"age,omitempty"` // nil = missing, 0 = zero
}
```

**Use Case**:
```go
// Age = 0 (valid)
age := 0
user1 := User{Age: &age}

// Age = missing
user2 := User{Age: nil}
```

### 4. Skip Sensitive Fields

**❌ Never**:
```go
type User struct {
    Name     string `beve:"name"`
    Password string `beve:"password"` // DON'T!
}
```

**✅ Always**:
```go
type User struct {
    Name     string `beve:"name"`
    Password string `beve:"-"` // Skip encoding
}
```

**Sensitive Fields**:
- Passwords
- API keys
- Tokens
- Private keys
- Internal state

### 5. Document Tag Meanings

**✅ Good**:
```go
type User struct {
    // Name is the user's full name
    Name string `beve:"name"`
    
    // Email is optional; omitted if empty
    Email string `beve:"email,omitempty"`
    
    // InternalID is never exposed in API
    InternalID int `beve:"-"`
}
```

### 6. Consistent Naming Convention

**❌ Inconsistent**:
```go
type User struct {
    Name  string `beve:"name"`        // snake_case
    Age   int    `beve:"age"`         // snake_case
    Email string `beve:"emailAddress"` // camelCase
}
```

**✅ Consistent**:
```go
type User struct {
    Name  string `beve:"name"`
    Age   int    `beve:"age"`
    Email string `beve:"email_address"` // All snake_case
}
```

**Common Conventions**:
- `snake_case`: `user_name`, `email_address`
- `camelCase`: `userName`, `emailAddress`
- `kebab-case`: `user-name`, `email-address`

**Pick one and stick with it!**

### 7. Version Your Schemas

**✅ Good**:
```go
type UserV1 struct {
    Name string `beve:"name"`
}

type UserV2 struct {
    Name  string `beve:"name"`
    Email string `beve:"email,omitempty"` // New field
}

// Migration function
func (v1 UserV1) ToV2() UserV2 {
    return UserV2{Name: v1.Name}
}
```

---

## Common Issues

### Issue 1: Tag Not Applied

**Problem**:
```go
type User struct {
    name string `beve:"name"` // Won't work!
}
```

**Reason**: Field is unexported (lowercase)

**Solution**:
```go
type User struct {
    Name string `beve:"name"` // Exported (capital N)
}
```

### Issue 2: omitempty Not Working

**Problem**:
```go
type User struct {
    Age int `beve:"age,omitempty"` // 0 is still encoded!
}
```

**Reason**: `0` is the zero value, but field is not a pointer

**Solution**:
```go
type User struct {
    Age *int `beve:"age,omitempty"` // nil is omitted
}
```

### Issue 3: Field Order Changes

**Problem**:
```go
// V1
type User struct {
    Name string `beve:"name"`
    Age  int    `beve:"age"`
}

// V2 (reordered fields)
type User struct {
    Age  int    `beve:"age"`
    Name string `beve:"name"`
}
```

**Impact**: None! BEVE uses field names (like JSON), not order.

**Safe Operations**:
- ✅ Reorder fields
- ✅ Add fields with `omitempty`
- ✅ Remove unused fields
- ❌ Rename fields without migration

### Issue 4: Tag Typo

**Problem**:
```go
type User struct {
    Name string `bevee:"name"` // Typo: "bevee" not "beve"
}
```

**Result**: Tag ignored, field encodes as "Name"

**Solution**: Use linter to catch typos

```bash
go vet ./...
```

### Issue 5: Circular References

**Problem**:
```go
type User struct {
    Name    string `beve:"name"`
    Manager *User  `beve:"manager"` // Circular!
}
```

**Result**: Stack overflow (infinite recursion)

**Solution**: Break cycle with ID reference

```go
type User struct {
    Name      string `beve:"name"`
    ManagerID *int64 `beve:"manager_id,omitempty"` // ID instead
}
```

---

## Tag Cheat Sheet

| Tag | Example | Effect |
|-----|---------|--------|
| `beve:"name"` | `Name string \`beve:"name"\`` | Field encodes as "name" |
| `beve:",omitempty"` | `Email string \`beve:",omitempty"\`` | Skip if zero value |
| `beve:"-"` | `Password string \`beve:"-"\`` | Never encode |
| `beve:"id,string"` | `ID int64 \`beve:"id,string"\`` | Encode as string |
| `beve:"addr,inline"` | `Address Addr \`beve:"addr,inline"\`` | Flatten embedded |
| No tag | `Name string` | Encode as "Name" |

---

## Summary

### Key Takeaways

1. **Always use tags** for explicit field names
2. **Use `omitempty`** for optional fields (10-30% smaller payloads)
3. **Skip sensitive fields** with `beve:"-"`
4. **Use pointers** to distinguish nil vs zero
5. **Version schemas** for backward compatibility
6. **Validate tags** with `go vet` and runtime checks
7. **Document tags** for maintainability

### Tag Options Quick Reference

```go
type Example struct {
    // Basic field name
    Name string `beve:"name"`
    
    // Optional field (omit if empty)
    Email string `beve:"email,omitempty"`
    
    // Skip field (never encoded)
    Password string `beve:"-"`
    
    // Force string encoding
    ID int64 `beve:"id,string"`
    
    // Inline embedded struct
    Address Address `beve:",inline"`
    
    // Pointer for nullable
    Age *int `beve:"age,omitempty"`
    
    // Multiple tags (JSON + BEVE)
    Username string `json:"username" beve:"name"`
}
```

### Next Steps

- **[Encoding/Decoding →](encoding-decoding.md)** - Marshal/Unmarshal deep dive
- **[Streaming →](streaming.md)** - Stream encoding patterns
- **[Performance →](performance.md)** - Optimize with tags
- **[API Reference →](../api/core.md)** - Function documentation

---

**Questions?** Check the [FAQ](../getting-started/faq.md) or open a [Discussion](https://github.com/beve-org/beve-go/discussions).
