# Production Security Guide

**BEVE-Go Security Best Practices**

Comprehensive security guidelines for deploying BEVE-Go applications in production environments.

---

## Table of Contents

1. [Input Validation](#input-validation)
2. [DoS Attack Prevention](#dos-attack-prevention)
3. [Memory Safety](#memory-safety)
4. [Data Privacy](#data-privacy)
5. [Common Vulnerabilities](#common-vulnerabilities)
6. [Security Checklist](#security-checklist)

---

## Input Validation

### Size Limits

**Enforce strict size limits** to prevent resource exhaustion:

```go
// security/limits.go
package security

import "errors"

const (
    MaxMessageSize  = 100 * 1024 * 1024  // 100 MB
    MaxNestingDepth = 16                 // Prevent stack overflow
    MaxArraySize    = 1_000_000          // 1M elements max
    MaxStringLength = 10 * 1024 * 1024   // 10 MB max string
    MaxObjectKeys   = 10_000             // Max keys in object
)

var (
    ErrMessageTooLarge   = errors.New("message exceeds max size")
    ErrNestingTooDeep    = errors.New("nesting depth exceeds limit")
    ErrArrayTooLarge     = errors.New("array size exceeds limit")
    ErrStringTooLong     = errors.New("string exceeds max length")
    ErrTooManyKeys       = errors.New("object has too many keys")
)

func ValidateMessageSize(data []byte) error {
    if len(data) > MaxMessageSize {
        return ErrMessageTooLarge
    }
    return nil
}
```

### Safe Unmarshal Wrapper

```go
// security/unmarshal.go
package security

import (
    "bytes"
    "context"
    "time"
)

type UnmarshalOptions struct {
    MaxSize         int
    MaxNestingDepth int
    Timeout         time.Duration
}

func SafeUnmarshal(data []byte, v interface{}, opts UnmarshalOptions) error {
    // Size check
    if len(data) > opts.MaxSize {
        return ErrMessageTooLarge
    }
    
    // Validate BEVE header
    if !beve.IsValidBEVE(data) {
        return errors.New("invalid BEVE format")
    }
    
    // Create timeout context
    ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
    defer cancel()
    
    // Unmarshal with timeout
    done := make(chan error, 1)
    go func() {
        done <- beve.Unmarshal(data, v)
    }()
    
    select {
    case err := <-done:
        return err
    case <-ctx.Done():
        return errors.New("unmarshal timeout")
    }
}
```

### Nesting Depth Validation

```go
// security/depth.go
package security

func ValidateNestingDepth(data []byte, maxDepth int) error {
    depth := 0
    maxSeen := 0
    
    for i := 0; i < len(data); i++ {
        header := data[i]
        typeID := header & 0x07
        
        switch typeID {
        case 0x03, 0x04, 0x05:  // Object, TypedArray, GenericArray
            depth++
            if depth > maxSeen {
                maxSeen = depth
            }
            if depth > maxDepth {
                return ErrNestingTooDeep
            }
        }
        
        // Simplification: actual impl would parse full structure
    }
    
    return nil
}
```

---

## DoS Attack Prevention

### Rate Limiting

**Per-client rate limiting**:

```go
// security/ratelimit.go
package security

import (
    "golang.org/x/time/rate"
    "sync"
)

type RateLimiter struct {
    limiters sync.Map  // map[string]*rate.Limiter
    rate     rate.Limit
    burst    int
}

func NewRateLimiter(r rate.Limit, burst int) *RateLimiter {
    return &RateLimiter{
        rate:  r,
        burst: burst,
    }
}

func (rl *RateLimiter) GetLimiter(clientID string) *rate.Limiter {
    if limiter, ok := rl.limiters.Load(clientID); ok {
        return limiter.(*rate.Limiter)
    }
    
    limiter := rate.NewLimiter(rl.rate, rl.burst)
    rl.limiters.Store(clientID, limiter)
    return limiter
}

func (rl *RateLimiter) Allow(clientID string) bool {
    limiter := rl.GetLimiter(clientID)
    return limiter.Allow()
}
```

**HTTP Middleware**:

```go
// middleware/ratelimit.go
package middleware

func RateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            clientID := r.Header.Get("X-Client-ID")
            if clientID == "" {
                clientID = r.RemoteAddr
            }
            
            if !rl.Allow(clientID) {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

### Timeout Configuration

```go
// server/server.go
package server

import (
    "net/http"
    "time"
)

func NewServer() *http.Server {
    return &http.Server{
        Addr:              ":8080",
        ReadTimeout:       30 * time.Second,  // Prevent slow read attacks
        WriteTimeout:      30 * time.Second,  // Prevent slow write attacks
        IdleTimeout:       120 * time.Second, // Close idle connections
        ReadHeaderTimeout: 10 * time.Second,  // Header read timeout
        MaxHeaderBytes:    1 << 20,           // 1 MB max headers
    }
}
```

### Resource Limits

```go
// security/limits.go

type ResourceLimits struct {
    MaxConcurrentOps  int           // Max parallel operations
    MaxMemoryPerOp    int64         // Max memory per operation
    OperationTimeout  time.Duration // Max operation duration
}

var DefaultLimits = ResourceLimits{
    MaxConcurrentOps: 10000,
    MaxMemoryPerOp:   100 * 1024 * 1024, // 100 MB
    OperationTimeout: 30 * time.Second,
}

// Semaphore for concurrent ops
type OpSemaphore struct {
    sem chan struct{}
}

func NewOpSemaphore(max int) *OpSemaphore {
    return &OpSemaphore{
        sem: make(chan struct{}, max),
    }
}

func (s *OpSemaphore) Acquire() {
    s.sem <- struct{}{}
}

func (s *OpSemaphore) Release() {
    <-s.sem
}

func (s *OpSemaphore) Do(fn func() error) error {
    s.Acquire()
    defer s.Release()
    return fn()
}
```

---

## Memory Safety

### Bounds Checking

**BEVE decoder includes automatic bounds checking**:

```go
// Internal decoder checks (built-in)
func (d *Decoder) readBytes(n int) ([]byte, error) {
    if d.pos+n > len(d.data) {
        return nil, ErrUnexpectedEOF
    }
    
    result := d.data[d.pos : d.pos+n]
    d.pos += n
    return result, nil
}
```

**Custom validation layer**:

```go
// security/validation.go
package security

func ValidateBounds(data []byte, offset, length int) error {
    if offset < 0 || length < 0 {
        return errors.New("negative offset or length")
    }
    
    if offset+length > len(data) {
        return errors.New("bounds check failed")
    }
    
    return nil
}
```

### Integer Overflow Protection

```go
// security/overflow.go
package security

import "math"

func SafeAdd(a, b int) (int, error) {
    if a > 0 && b > math.MaxInt-a {
        return 0, errors.New("integer overflow")
    }
    if a < 0 && b < math.MinInt-a {
        return 0, errors.New("integer underflow")
    }
    return a + b, nil
}

func SafeMultiply(a, b int) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }
    
    result := a * b
    if result/b != a {
        return 0, errors.New("integer overflow")
    }
    
    return result, nil
}
```

### Unsafe Pointer Usage

**BEVE uses `unsafe` for performance** - guidelines:

```go
// ✅ SAFE: Bounded string conversion
func unsafeBytesToString(b []byte) string {
    if len(b) == 0 {
        return ""
    }
    return unsafe.String(&b[0], len(b))
}

// ❌ UNSAFE: No bounds check
func unsafeBadConversion(ptr uintptr, len int) string {
    return unsafe.String((*byte)(unsafe.Pointer(ptr)), len)  // DANGEROUS!
}
```

**Security rules**:
1. Always validate `len(b) > 0` before `unsafe.Pointer(&b[0])`
2. Never use `uintptr` directly
3. Keep `unsafe` in separate files with `// +build !safe` tag

---

## Data Privacy

### Sensitive Field Handling

**Redact sensitive data in logs**:

```go
// logging/redact.go
package logging

import "reflect"

type Redactor struct {
    sensitiveFields map[string]bool
}

func NewRedactor() *Redactor {
    return &Redactor{
        sensitiveFields: map[string]bool{
            "password": true,
            "token":    true,
            "api_key":  true,
            "secret":   true,
            "ssn":      true,
            "credit_card": true,
        },
    }
}

func (r *Redactor) Redact(data interface{}) interface{} {
    v := reflect.ValueOf(data)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return data
    }
    
    result := make(map[string]interface{})
    t := v.Type()
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        fieldName := field.Tag.Get("json")
        if fieldName == "" {
            fieldName = field.Name
        }
        
        if r.sensitiveFields[fieldName] {
            result[fieldName] = "[REDACTED]"
        } else {
            result[fieldName] = v.Field(i).Interface()
        }
    }
    
    return result
}
```

### Field-Level Encryption

```go
// crypto/encrypt.go
package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
)

type FieldEncryptor struct {
    key []byte  // 32 bytes for AES-256
}

func NewFieldEncryptor(key []byte) (*FieldEncryptor, error) {
    if len(key) != 32 {
        return nil, errors.New("key must be 32 bytes for AES-256")
    }
    return &FieldEncryptor{key: key}, nil
}

func (e *FieldEncryptor) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        return "", err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *FieldEncryptor) Decrypt(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("invalid ciphertext")
    }
    
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    
    return string(plaintext), nil
}
```

### Audit Logging

```go
// audit/logger.go
package audit

import "encoding/json"

type AuditEvent struct {
    Timestamp   string `json:"timestamp"`
    UserID      string `json:"user_id"`
    Action      string `json:"action"`
    Resource    string `json:"resource"`
    IPAddress   string `json:"ip_address"`
    Success     bool   `json:"success"`
    Error       string `json:"error,omitempty"`
}

func LogDataAccess(userID, resourceID, ipAddr string, success bool, err error) {
    event := AuditEvent{
        Timestamp: time.Now().UTC().Format(time.RFC3339),
        UserID:    userID,
        Action:    "data_access",
        Resource:  resourceID,
        IPAddress: ipAddr,
        Success:   success,
    }
    
    if err != nil {
        event.Error = err.Error()
    }
    
    data, _ := json.Marshal(event)
    auditLog.Println(string(data))
}
```

---

## Common Vulnerabilities

### 1. Malformed BEVE Payloads

**Attack**: Send invalid BEVE headers to crash decoder

**Mitigation**:

```go
func ValidateBEVEHeader(data []byte) error {
    if len(data) < 1 {
        return errors.New("empty payload")
    }
    
    header := data[0]
    typeID := header & 0x07
    
    // Valid type IDs: 0-6
    if typeID > 6 {
        return errors.New("invalid type ID")
    }
    
    return nil
}
```

### 2. Billion Laughs Attack (Nested Arrays)

**Attack**: Exponential memory explosion with deeply nested arrays

```
[[[[[[[[[[...]]]]]]]]]]  // 1000 levels deep
```

**Mitigation**: Enforce `MaxNestingDepth = 16`

```go
func countNestingDepth(data []byte) int {
    depth := 0
    maxDepth := 0
    
    for _, b := range data {
        if isOpeningDelimiter(b) {
            depth++
            if depth > maxDepth {
                maxDepth = depth
            }
        } else if isClosingDelimiter(b) {
            depth--
        }
    }
    
    return maxDepth
}

func validateNesting(data []byte) error {
    if countNestingDepth(data) > MaxNestingDepth {
        return ErrNestingTooDeep
    }
    return nil
}
```

### 3. Hash Collision Attacks (Object Keys)

**Attack**: Craft object keys to cause hash collisions, degrading performance

**Mitigation**: Use cryptographic hash for object keys

```go
// Go's map uses SipHash (collision-resistant)
// No additional mitigation needed in BEVE
```

### 4. Zip Bomb Equivalent

**Attack**: Highly compressible BEVE payload that expands massively

```
// Compressed: 1 KB
// Expands to: 1 GB after unmarshal
```

**Mitigation**:

```go
func ValidateExpansionRatio(compressedSize, expandedSize int) error {
    const MaxExpansionRatio = 1000  // 1000:1 max
    
    ratio := expandedSize / compressedSize
    if ratio > MaxExpansionRatio {
        return errors.New("expansion ratio too high")
    }
    
    return nil
}
```

### 5. Resource Exhaustion

**Attack**: Send many large messages to exhaust memory/CPU

**Mitigation**:

```go
// Rate limiting (see DoS section)
// Memory limits (GOMEMLIMIT)
// Timeouts (ReadTimeout, WriteTimeout)
// Concurrent operation limits (OpSemaphore)
```

---

## Security Checklist

### Pre-Production

- [ ] **Input Validation**
  - [ ] Max message size enforced (`100 MB`)
  - [ ] Max nesting depth enforced (`16 levels`)
  - [ ] Max array size enforced (`1M elements`)
  - [ ] Max string length enforced (`10 MB`)
  - [ ] BEVE header validation enabled

- [ ] **Rate Limiting**
  - [ ] Per-client rate limiting configured
  - [ ] Burst limits set (`100 req/min`)
  - [ ] IP-based blocking for abuse

- [ ] **Timeouts**
  - [ ] Read timeout: `30s`
  - [ ] Write timeout: `30s`
  - [ ] Operation timeout: `30s`
  - [ ] Idle connection timeout: `120s`

- [ ] **Resource Limits**
  - [ ] Max concurrent operations: `10,000`
  - [ ] Memory limit (`GOMEMLIMIT`): `8 GB`
  - [ ] CPU limit (`GOMAXPROCS`): Auto-detect

- [ ] **Data Privacy**
  - [ ] Sensitive fields redacted in logs
  - [ ] Encryption enabled for PII
  - [ ] Audit logging configured
  - [ ] Access controls implemented

- [ ] **Memory Safety**
  - [ ] Bounds checking enabled
  - [ ] Integer overflow protection
  - [ ] Unsafe pointer usage audited

- [ ] **Monitoring**
  - [ ] Error rate alerts configured
  - [ ] Memory usage alerts configured
  - [ ] Latency alerts configured
  - [ ] Security events logged

### Production Hardening

```go
// config/production.go
package config

var ProductionConfig = Config{
    MaxMessageSize:     100 * 1024 * 1024,
    MaxNestingDepth:    16,
    MaxArraySize:       1_000_000,
    MaxConcurrentOps:   10_000,
    ReadTimeout:        30 * time.Second,
    WriteTimeout:       30 * time.Second,
    EnableRateLimit:    true,
    RateLimit:          100,  // req/min
    RateBurst:          10,
    EnableAuditLog:     true,
    RedactSensitive:    true,
}
```

---

## Security Incident Response

### Detection

**Monitor for**:
- Spike in error rate
- Unusual traffic patterns
- High memory usage
- Slow response times
- Invalid BEVE payloads

### Response Plan

1. **Identify**: Alert triggered
2. **Contain**: Rate limit offending IPs
3. **Analyze**: Review logs, payloads
4. **Remediate**: Apply fixes, update rules
5. **Report**: Document incident

### Example: Handling DoS Attack

```go
// incident/dos.go
package incident

func HandleDoSAttack(clientIP string) {
    // 1. Block IP
    rateLimiter.Block(clientIP, 1*time.Hour)
    
    // 2. Log incident
    audit.LogSecurityEvent("dos_attack", map[string]interface{}{
        "client_ip": clientIP,
        "action":    "blocked",
        "duration":  "1h",
    })
    
    // 3. Alert team
    alertManager.Send("DoS attack detected from " + clientIP)
    
    // 4. Scale up (if needed)
    if metricsExceedThreshold() {
        scaler.ScaleUp(2)  // 2× replicas
    }
}
```

---

## Summary

### Security Layers

1. **Input Validation**: Size, nesting, format checks
2. **Rate Limiting**: Per-client, burst protection
3. **Timeouts**: Read, write, operation timeouts
4. **Resource Limits**: Concurrent ops, memory, CPU
5. **Memory Safety**: Bounds checks, overflow protection
6. **Data Privacy**: Encryption, redaction, audit logs
7. **Monitoring**: Error rates, anomalies, incidents

### Key Metrics

| Security Metric | Target | Alert |
|-----------------|--------|-------|
| **Invalid Payloads** | < 0.1% | > 1% |
| **Rate Limit Hits** | < 5% | > 10% |
| **Blocked IPs** | < 10/day | > 100/day |
| **Memory Violations** | 0 | > 0 |

### Compliance

- ✅ **OWASP Top 10**: Covered
- ✅ **CWE-20**: Input validation
- ✅ **CWE-400**: Resource exhaustion
- ✅ **CWE-770**: Allocation without limits
- ✅ **GDPR**: PII encryption/redaction

---

**Next**: [Troubleshooting](troubleshooting.md) · [Monitoring](monitoring.md) · [Deployment](deployment.md)
