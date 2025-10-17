# 🚨 Error Handling Guide

Comprehensive guide to handling errors in BEVE encoding/decoding operations.

**Reading Time**: 12 minutes  
**Level**: Intermediate  
**Prerequisites**: [Basic Usage](../getting-started/basic-usage.md)

---

## Table of Contents

1. [Error Types](#error-types)
2. [Marshal Errors](#marshal-errors)
3. [Unmarshal Errors](#unmarshal-errors)
4. [Error Handling Patterns](#error-handling-patterns)
5. [Validation](#validation)
6. [Recovery Strategies](#recovery-strategies)
7. [Production Best Practices](#production-best-practices)

---

## Error Types

### Common Error Categories

BEVE errors fall into 5 categories:

| Category | Examples | Recoverable? |
|----------|----------|--------------|
| **Type Errors** | Type mismatch, unsupported type | ❌ No |
| **Data Errors** | Invalid BEVE header, corrupted data | ❌ No |
| **Size Errors** | Buffer too small, data too large | ✅ Yes |
| **Schema Errors** | Missing fields, wrong struct | ⚠️ Partial |
| **System Errors** | Out of memory, I/O errors | ⚠️ Maybe |

### Error Variables

```go
import "github.com/beve-org/beve-go"

// Pre-defined errors
beve.ErrInvalidHeader       // Invalid BEVE header byte
beve.ErrTypeMismatch        // Type doesn't match expected
beve.ErrBufferTooSmall      // Buffer insufficient for encoding
beve.ErrUnsupportedType     // Type not supported by BEVE
beve.ErrDataCorrupted       // Data integrity check failed
beve.ErrInvalidSize         // Size field invalid or corrupted
```

### Error Checking

```go
import "errors"

// Check specific error
if errors.Is(err, beve.ErrTypeMismatch) {
    // Handle type mismatch
}

// Check error type
var sizeErr *beve.BufferTooSmallError
if errors.As(err, &sizeErr) {
    // Access error details
    fmt.Printf("Need %d bytes, have %d\n", sizeErr.Need, sizeErr.Have)
}
```

---

## Marshal Errors

### Type Mismatch

**Error**: Unsupported type

```go
// ❌ Error: Function type not supported
fn := func() {}
data, err := beve.Marshal(fn)
// err: unsupported type: func()
```

**Solution**: Use supported types

```go
// ✅ Use supported types
type Handler struct {
    Name string `beve:"name"`
    // Store function name, not function
}

handler := Handler{Name: "handleRequest"}
data, _ := beve.Marshal(handler)
```

### Unsupported Types

BEVE cannot encode:
- Functions
- Channels
- Unsafe pointers
- Cyclic references (recursive structs)

```go
// ❌ Unsupported
type BadStruct struct {
    Fn   func()     // Function
    Ch   chan int   // Channel
    Self *BadStruct // Cyclic reference
}
```

### Buffer Too Small

**Error**: Zero-copy buffer insufficient

```go
buf := make([]byte, 0, 10) // Too small!
data, err := beve.MarshalZeroCopy(largeStruct, buf)
// err: buffer too small: need 1000 bytes, have 10
```

**Solution 1**: Pre-allocate larger buffer

```go
// Estimate size
size := beve.EstimateSize(largeStruct)
buf := make([]byte, 0, size*2)

data, err := beve.MarshalZeroCopy(largeStruct, buf)
```

**Solution 2**: Grow and retry

```go
buf := make([]byte, 0, 256)
data, err := beve.MarshalZeroCopy(v, buf)

if errors.Is(err, beve.ErrBufferTooSmall) {
    // Grow buffer and retry
    buf = make([]byte, 0, len(buf)*2)
    data, err = beve.MarshalZeroCopy(v, buf)
}
```

**Solution 3**: Fallback to standard marshal

```go
data, err := beve.MarshalZeroCopy(v, buf)
if errors.Is(err, beve.ErrBufferTooSmall) {
    // Fallback to standard marshal (auto-grows)
    data, err = beve.Marshal(v)
}
```

---

## Unmarshal Errors

### Invalid Header

**Error**: Data is not valid BEVE

```go
data := []byte{0xFF, 0xFF, 0xFF} // Random bytes
var user User
err := beve.Unmarshal(data, &user)
// err: invalid BEVE header: 0xFF
```

**Solution**: Validate before unmarshal

```go
// Check if data is valid BEVE
if !beve.Valid(data) {
    return errors.New("invalid BEVE data")
}

var user User
err := beve.Unmarshal(data, &user)
```

### Type Mismatch

**Error**: Data type doesn't match target

```go
// Data contains string
data := beve.Marshal("hello")

// Try to unmarshal as int
var num int
err := beve.Unmarshal(data, &num)
// err: type mismatch: expected int, got string
```

**Solution**: Check type before unmarshal

```go
// Detect type
encoding := beve.DetectEncoding(data)

switch encoding {
case "string":
    var s string
    beve.Unmarshal(data, &s)
case "int":
    var i int
    beve.Unmarshal(data, &i)
default:
    return errors.New("unexpected type")
}
```

### Nil Pointer

**Error**: Unmarshal to nil pointer

```go
var user *User // nil pointer!
err := beve.Unmarshal(data, user)
// err: cannot unmarshal to nil pointer
```

**Solution**: Initialize pointer

```go
user := &User{}
err := beve.Unmarshal(data, user)
```

### Missing Fields

**Warning**: Not an error, but important behavior

```go
type User struct {
    Name  string `beve:"name"`
    Email string `beve:"email"`
}

// Data only has "name"
data := beve.Marshal(map[string]interface{}{
    "name": "Alice",
})

var user User
beve.Unmarshal(data, &user)
// user.Name = "Alice"
// user.Email = "" (zero value, not error)
```

**Solution**: Validate after unmarshal

```go
var user User
err := beve.Unmarshal(data, &user)
if err != nil {
    return err
}

// Validate required fields
if user.Email == "" {
    return errors.New("email is required")
}
```

---

## Error Handling Patterns

### Pattern 1: Wrap Errors

```go
func loadUser(data []byte) (*User, error) {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        return nil, fmt.Errorf("failed to load user: %w", err)
    }
    return &user, nil
}

// Caller can check wrapped error
user, err := loadUser(data)
if err != nil {
    if errors.Is(err, beve.ErrTypeMismatch) {
        // Handle type error specifically
    }
    return err
}
```

### Pattern 2: Graceful Degradation

```go
func decodeUserOrDefault(data []byte) User {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        // Return default on error
        log.Printf("Decode error: %v, using default", err)
        return User{Name: "Unknown"}
    }
    return user
}
```

### Pattern 3: Partial Success

```go
func decodeUsers(data []byte) ([]User, []error) {
    var rawUsers []map[string]interface{}
    if err := beve.Unmarshal(data, &rawUsers); err != nil {
        return nil, []error{err}
    }
    
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

// Usage
users, errs := decodeUsers(data)
if len(errs) > 0 {
    log.Printf("Failed to decode %d users", len(errs))
}
fmt.Printf("Successfully decoded %d users\n", len(users))
```

### Pattern 4: Retry with Fallback

```go
func encode(v interface{}) ([]byte, error) {
    // Try zero-copy first
    buf := make([]byte, 0, 1024)
    data, err := beve.MarshalZeroCopy(v, buf)
    if err == nil {
        return data, nil
    }
    
    // Fallback to standard marshal
    if errors.Is(err, beve.ErrBufferTooSmall) {
        return beve.Marshal(v)
    }
    
    // Other errors
    return nil, err
}
```

### Pattern 5: Validate and Recover

```go
func decodeWithValidation(data []byte) (*User, error) {
    // Validate BEVE format
    if !beve.Valid(data) {
        // Try to recover (maybe it's JSON?)
        var user User
        if err := json.Unmarshal(data, &user); err == nil {
            log.Println("Recovered JSON data")
            return &user, nil
        }
        return nil, errors.New("invalid data format")
    }
    
    // Decode BEVE
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        return nil, err
    }
    
    // Validate business rules
    if err := validateUser(user); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    return &user, nil
}
```

---

## Validation

### Pre-Marshal Validation

```go
func validateAndMarshal(user User) ([]byte, error) {
    // Validate before encoding
    if user.Name == "" {
        return nil, errors.New("name is required")
    }
    if user.Age < 0 || user.Age > 150 {
        return nil, errors.New("age must be 0-150")
    }
    if !isValidEmail(user.Email) {
        return nil, errors.New("invalid email format")
    }
    
    // Marshal
    return beve.Marshal(user)
}
```

### Post-Unmarshal Validation

```go
func unmarshalAndValidate(data []byte) (*User, error) {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        return nil, fmt.Errorf("unmarshal failed: %w", err)
    }
    
    // Validate after decoding
    if err := validateUser(user); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    return &user, nil
}

func validateUser(user User) error {
    if user.Name == "" {
        return errors.New("name is required")
    }
    if user.Age < 0 {
        return errors.New("age must be non-negative")
    }
    if user.Email == "" {
        return errors.New("email is required")
    }
    return nil
}
```

### Validation with go-validator

```go
import "github.com/go-playground/validator/v10"

type User struct {
    Name  string `beve:"name" validate:"required,min=2,max=50"`
    Age   int    `beve:"age" validate:"gte=0,lte=150"`
    Email string `beve:"email" validate:"required,email"`
}

var validate = validator.New()

func unmarshalAndValidate(data []byte) (*User, error) {
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        return nil, err
    }
    
    // Validate with go-validator
    if err := validate.Struct(user); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    return &user, nil
}
```

---

## Recovery Strategies

### Strategy 1: Fallback Format

```go
func decode(data []byte) (*User, error) {
    // Try BEVE first
    if beve.Valid(data) {
        var user User
        if err := beve.Unmarshal(data, &user); err == nil {
            return &user, nil
        }
    }
    
    // Fallback to JSON
    var user User
    if err := json.Unmarshal(data, &user); err == nil {
        log.Println("Decoded as JSON (fallback)")
        return &user, nil
    }
    
    return nil, errors.New("unsupported format")
}
```

### Strategy 2: Default Values

```go
func decodeWithDefaults(data []byte) User {
    defaults := User{
        Name:  "Unknown",
        Age:   0,
        Email: "no-reply@example.com",
    }
    
    var user User
    if err := beve.Unmarshal(data, &user); err != nil {
        log.Printf("Decode error: %v, using defaults", err)
        return defaults
    }
    
    // Merge with defaults (fill missing fields)
    if user.Name == "" {
        user.Name = defaults.Name
    }
    if user.Email == "" {
        user.Email = defaults.Email
    }
    
    return user
}
```

### Strategy 3: Skip Invalid Items

```go
func decodeArray(data []byte) []User {
    var rawData [][]byte
    beve.Unmarshal(data, &rawData)
    
    var users []User
    
    for i, item := range rawData {
        var user User
        if err := beve.Unmarshal(item, &user); err != nil {
            log.Printf("Skipping item %d: %v", i, err)
            continue
        }
        users = append(users, user)
    }
    
    return users
}
```

### Strategy 4: Circuit Breaker

```go
type CircuitBreaker struct {
    failures   int
    threshold  int
    isOpen     bool
    lastFailed time.Time
}

func (cb *CircuitBreaker) Decode(data []byte) (*User, error) {
    // Check if circuit open
    if cb.isOpen {
        if time.Since(cb.lastFailed) < 30*time.Second {
            return nil, errors.New("circuit breaker open")
        }
        cb.isOpen = false
        cb.failures = 0
    }
    
    // Try decode
    var user User
    err := beve.Unmarshal(data, &user)
    
    if err != nil {
        cb.failures++
        cb.lastFailed = time.Now()
        
        if cb.failures >= cb.threshold {
            cb.isOpen = true
            log.Println("Circuit breaker opened")
        }
        
        return nil, err
    }
    
    // Success: reset failures
    cb.failures = 0
    return &user, nil
}
```

---

## Production Best Practices

### 1. Always Check Errors

```go
// ❌ Bad: Ignore errors
data, _ := beve.Marshal(user)

// ✅ Good: Handle errors
data, err := beve.Marshal(user)
if err != nil {
    log.Printf("Marshal error: %v", err)
    return err
}
```

### 2. Wrap with Context

```go
// ✅ Good: Add context
func saveUser(user User) error {
    data, err := beve.Marshal(user)
    if err != nil {
        return fmt.Errorf("failed to marshal user %s: %w", user.Name, err)
    }
    
    if err := db.Save(data); err != nil {
        return fmt.Errorf("failed to save user %s: %w", user.Name, err)
    }
    
    return nil
}
```

### 3. Log for Debugging

```go
// ✅ Good: Log errors with details
func process(data []byte) error {
    var user User
    err := beve.Unmarshal(data, &user)
    if err != nil {
        log.Printf("Unmarshal error: %v, data size: %d bytes, header: 0x%02x",
            err, len(data), data[0])
        return err
    }
    return nil
}
```

### 4. Metrics and Monitoring

```go
var (
    marshalErrors   = prometheus.NewCounter(...)
    unmarshalErrors = prometheus.NewCounter(...)
)

func marshal(v interface{}) ([]byte, error) {
    data, err := beve.Marshal(v)
    if err != nil {
        marshalErrors.Inc()
        log.Printf("Marshal error: %v", err)
    }
    return data, err
}
```

### 5. Panic Recovery

```go
func safeMarshal(v interface{}) (data []byte, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("marshal panic: %v", r)
            log.Printf("Panic in marshal: %v\nStack: %s", r, debug.Stack())
        }
    }()
    
    return beve.Marshal(v)
}
```

---

## Summary

### Error Handling Checklist

**Marshal**:
- ✅ Check for unsupported types
- ✅ Handle buffer too small (zero-copy)
- ✅ Validate input before encoding
- ✅ Wrap errors with context

**Unmarshal**:
- ✅ Validate BEVE format first
- ✅ Check for type mismatches
- ✅ Initialize pointers before unmarshal
- ✅ Validate output after decoding

**Production**:
- ✅ Log errors with details
- ✅ Add metrics/monitoring
- ✅ Implement retry logic
- ✅ Use circuit breakers
- ✅ Recover from panics

### Common Errors Quick Reference

| Error | Cause | Solution |
|-------|-------|----------|
| `invalid header` | Invalid BEVE data | Validate with `beve.Valid()` |
| `type mismatch` | Wrong target type | Check type with `DetectEncoding()` |
| `buffer too small` | Zero-copy buffer insufficient | Pre-allocate or grow buffer |
| `nil pointer` | Unmarshal to nil | Initialize pointer first |
| `unsupported type` | Function/channel/etc. | Use supported types only |

### Next Steps

- **[Production Deployment →](../production/deployment.md)** - Deploy with confidence
- **[Monitoring →](../production/monitoring.md)** - Track errors in production
- **[Troubleshooting →](../production/troubleshooting.md)** - Debug issues
- **[API Reference →](../api/core.md)** - Error types reference

---

**Ready for production?** Check the [Production Deployment Guide](../production/deployment.md) for best practices.
