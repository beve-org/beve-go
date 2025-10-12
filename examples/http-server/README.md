# HTTP Server Example with BEVE

This example demonstrates how to use BEVE encoding in a standard `net/http` server.

## Features

- ✅ Uses `application/beve` MIME type
- ✅ GET endpoints returning BEVE-encoded responses
- ✅ POST endpoint accepting BEVE-encoded requests
- ✅ Middleware for logging
- ✅ Health check endpoint
- ✅ Error handling

## Running the Example

```bash
cd examples/http-server
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### GET /health
Health check endpoint

**Response:**
```
Status: 200 OK
Content-Type: application/beve

{success: true, data: {status: "healthy", timestamp: 1697040000, format: "BEVE"}}
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

{success: true, data: {id: 4, name: "John Doe", email: "john@example.com", created_at: "...", active: true}}
```

## Testing with curl

### Get user (if you have a BEVE client)
```bash
curl -H "Accept: application/beve" http://localhost:8080/api/user
```

### Create user with BEVE payload
First, encode your data to BEVE format using a Go client or tool, then:

```bash
# This is conceptual - you'd need to encode to BEVE first
curl -X POST \
  -H "Content-Type: application/beve" \
  --data-binary @user.beve \
  http://localhost:8080/api/users/create
```

## Performance Benefits

Using BEVE in HTTP APIs provides:

- **30% smaller payloads** compared to JSON
- **30× faster unmarshaling** for incoming requests
- **792 MB/s throughput** for response encoding
- **Lower CPU usage** for high-traffic APIs
- **Reduced bandwidth** costs

## Integration with Go Clients

```go
package main

import (
    "bytes"
    "fmt"
    "net/http"
    "github.com/beve-org/beve-go"
)

type User struct {
    Name  string `beve:"name"`
    Email string `beve:"email"`
    Active bool  `beve:"active"`
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com", Active: true}
    
    // Marshal to BEVE
    data, _ := beve.Marshal(user)
    
    // Send request
    resp, _ := http.Post(
        "http://localhost:8080/api/users/create",
        "application/beve",
        bytes.NewReader(data),
    )
    defer resp.Body.Close()
    
    // Read BEVE response
    body, _ := io.ReadAll(resp.Body)
    var response Response
    beve.Unmarshal(body, &response)
    
    fmt.Printf("Created user: %+v\n", response.Data)
}
```

## Comparison: BEVE vs JSON

For a typical user response:

| Format | Size | Encoding Time | Decoding Time |
|--------|------|---------------|---------------|
| BEVE | 155 bytes | 1,276 ns | 377 ns |
| JSON | 222 bytes | 2,048 ns | 11,480 ns |

**BEVE Advantages:**
- 30% smaller
- 1.6× faster encoding
- 30× faster decoding
