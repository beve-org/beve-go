# Fiber v2 Server Example with BEVE

This example demonstrates how to use BEVE encoding in a [Fiber v2](https://github.com/gofiber/fiber) web framework.

## Features

- ✅ Uses `application/beve` MIME type
- ✅ Fiber v2 integration with custom response handler
- ✅ GET endpoints returning BEVE-encoded responses
- ✅ POST endpoint accepting BEVE-encoded requests
- ✅ Built-in middleware (logger, recover)
- ✅ Health check and stats endpoints
- ✅ High-performance routing

## Installation

```bash
cd examples/fiber-server
go mod init fiber-server
go get github.com/gofiber/fiber/v2
go get github.com/beve-org/beve-go
```

## Running the Example

```bash
go run main.go
```

Server will start on `http://localhost:3000`

## Why Fiber + BEVE?

Fiber is already one of the fastest Go web frameworks. Combined with BEVE:

- **Fiber**: ~6M req/sec (routing)
- **BEVE**: 792 MB/s encoding, 250 MB/s decoding
- **Result**: Extremely high-throughput API server

## API Endpoints

### GET /health
Health check with framework info

**Response:**
```
Status: 200 OK
Content-Type: application/beve

{success: true, data: {status: "healthy", timestamp: 1697040000, format: "BEVE", framework: "Fiber v2"}}
```

### GET /stats
Server statistics

**Response:**
```
Status: 200 OK
Content-Type: application/beve

{success: true, data: {connections: 5, routes_count: 6, ...}}
```

### GET /api/user
Get a single user

**Response:**
```
Status: 200 OK
Content-Type: application/beve

{success: true, data: {id: 1, name: "Alice Johnson", ...}}
```

### GET /api/users
Get all users

**Response:**
```
Status: 200 OK
Content-Type: application/beve

{success: true, data: [{id: 1, ...}, {id: 2, ...}, {id: 3, ...}]}
```

### POST /api/users/create
Create a new user

**Request:**
```
Content-Type: application/beve

{name: "John Doe", email: "john@example.com", active: true}
```

**Response:**
```
Status: 201 Created
Content-Type: application/beve

{success: true, data: {id: 4, name: "John Doe", ...}}
```

## Testing with Fiber Test Client

```go
package main

import (
    "bytes"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/beve-org/beve-go"
)

func TestCreateUser(t *testing.T) {
    app := fiber.New()
    app.Post("/api/users/create", CreateUser)

    user := User{Name: "Test User", Email: "test@example.com"}
    body, _ := beve.Marshal(user)

    req := bytes.NewReader(body)
    resp, _ := app.Test(fiber.AcquireRequest().
        SetBody(req).
        SetHeader("Content-Type", "application/beve"),
    )

    // Read BEVE response
    respBody, _ := io.ReadAll(resp.Body)
    var response Response
    beve.Unmarshal(respBody, &response)

    if !response.Success {
        t.Fatal("Expected success")
    }
}
```

## Performance Comparison

### Fiber + JSON vs Fiber + BEVE

| Metric | JSON | BEVE | Improvement |
|--------|------|------|-------------|
| Encoding | 2,048 ns | 1,276 ns | 1.6× faster |
| Decoding | 11,480 ns | 377 ns | 30× faster |
| Payload Size | 222 bytes | 155 bytes | 30% smaller |
| Throughput | ~283 MB/s | ~792 MB/s | 2.8× faster |

### Real-World Impact

For a Fiber API handling **10,000 requests/second**:

| Metric | JSON | BEVE | Savings |
|--------|------|------|---------|
| CPU Time | 205 ms/sec | 16 ms/sec | **92% less CPU** |
| Memory Allocs | 760K/sec | 40K/sec | **95% fewer** |
| Bandwidth | 2.22 MB/sec | 1.55 MB/sec | **30% less** |

## Advanced: Custom BEVE Handler

Create a reusable Fiber handler for BEVE:

```go
func BEVEHandler(handler func(*fiber.Ctx) (interface{}, error)) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Call handler
        data, err := handler(c)
        if err != nil {
            return SendBEVE(c, fiber.StatusInternalServerError, Response{
                Success: false,
                Error:   err.Error(),
            })
        }

        // Send BEVE response
        return SendBEVE(c, fiber.StatusOK, Response{
            Success: true,
            Data:    data,
        })
    }
}

// Usage
app.Get("/api/user", BEVEHandler(func(c *fiber.Ctx) (interface{}, error) {
    return User{ID: 1, Name: "Alice"}, nil
}))
```

## Middleware Support

```go
// Content negotiation middleware
func ContentNegotiation(c *fiber.Ctx) error {
    accept := c.Get("Accept")
    
    if accept == "application/beve" {
        c.Locals("format", "beve")
    } else {
        c.Locals("format", "json")
    }
    
    return c.Next()
}

// Compression middleware (optional)
func BEVECompression(c *fiber.Ctx) error {
    // BEVE is already compact, but you can add gzip if needed
    // For payloads > 1KB, gzip can reduce size by 20-40%
    return c.Next()
}
```

## Production Considerations

1. **Error Handling**: Always validate BEVE unmarshaling
2. **Content-Type**: Enforce `application/beve` for POST/PUT
3. **Versioning**: Include API version in response
4. **Monitoring**: Log BEVE vs JSON usage
5. **Caching**: BEVE payloads are cacheable like any binary format

## Links

- [Fiber Documentation](https://docs.gofiber.io/)
- [BEVE Specification](https://github.com/beve-org/beve)
- [Performance Benchmarks](../../PHASE2_RESULTS.md)
