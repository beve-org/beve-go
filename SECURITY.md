# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability, please follow these steps:

### 🔒 Private Disclosure

**DO NOT** create a public GitHub issue for security vulnerabilities.

Instead, please email us at:
- **Email**: security@beve.dev (or create a private security advisory on GitHub)

### 📧 What to Include

When reporting a vulnerability, please include:

1. **Description**: Clear description of the vulnerability
2. **Impact**: What could an attacker accomplish?
3. **Steps to Reproduce**: Detailed steps to reproduce the issue
4. **Proof of Concept**: Code sample demonstrating the vulnerability (if possible)
5. **Suggested Fix**: If you have ideas for how to fix it
6. **Your Contact Info**: How we can reach you for follow-up

### ⏱️ Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 5 business days
- **Fix Timeline**: Depends on severity
  - Critical: 1-7 days
  - High: 7-14 days
  - Medium: 14-30 days
  - Low: 30-90 days

### 🛡️ Security Update Process

1. **Acknowledgment**: We confirm receipt and validate the report
2. **Investigation**: We investigate and assess severity
3. **Fix Development**: We develop and test a fix
4. **Coordinated Disclosure**: We coordinate release timing with you
5. **Release**: We release the fix and publish a security advisory
6. **Credit**: We credit you in the advisory (unless you prefer anonymity)

## 🔐 Security Best Practices

When using BEVE:

### Input Validation
```go
// Always validate input size
const MaxInputSize = 10 * 1024 * 1024 // 10MB

func SafeUnmarshal(data []byte, v interface{}) error {
    if len(data) > MaxInputSize {
        return errors.New("input too large")
    }
    return beve.Unmarshal(data, v)
}
```

### Resource Limits
```go
// Set reasonable limits for collections
type Config struct {
    MaxArraySize int
    MaxMapSize   int
    MaxDepth     int
}

// Implement timeouts for decoding
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

done := make(chan error, 1)
go func() {
    done <- beve.Unmarshal(data, &result)
}()

select {
case err := <-done:
    return err
case <-ctx.Done():
    return errors.New("decode timeout")
}
```

### Untrusted Input
```go
// When decoding untrusted input, use type assertions carefully
var result interface{}
if err := beve.Unmarshal(untrustedData, &result); err != nil {
    return err
}

// Validate types before using
switch v := result.(type) {
case map[string]interface{}:
    // Process map safely
case []interface{}:
    // Process array safely
default:
    return errors.New("unexpected type")
}
```

### Memory Safety
```go
// BEVE uses reflection - be aware of memory implications
// For very large datasets, consider streaming or chunking

type StreamDecoder struct {
    decoder *beve.Decoder
    maxSize int
}

func (s *StreamDecoder) DecodeChunk(r io.Reader) error {
    // Process data in manageable chunks
    // instead of loading everything into memory
}
```

## 🚨 Known Security Considerations

### 1. Reflection-Based Decoding
BEVE uses Go's reflection package, which can have performance implications for very large inputs. Consider implementing size limits for untrusted input.

### 2. Recursive Data Structures
Deeply nested structures can cause stack overflow. We have built-in depth limits, but validate your data structure complexity.

### 3. Binary Format
BEVE is a binary format - always validate input from untrusted sources before decoding.

## 📊 Security Audits

We welcome security audits and will publicly acknowledge researchers who help improve BEVE's security.

### Hall of Fame
- None yet - be the first!

## 📞 Contact

For security concerns:
- **Email**: security@beve.dev
- **GitHub**: [Private Security Advisory](https://github.com/beve-org/beve-go/security/advisories/new)

For general questions:
- **GitHub Issues**: https://github.com/beve-org/beve-go/issues
- **GitHub Discussions**: https://github.com/beve-org/beve-go/discussions

---

Thank you for helping keep BEVE and its users safe! 🛡️
