# BEVE Axios Interceptor

Automatic BEVE encoding/decoding for Axios HTTP client.

## 🚀 Features

- ✅ **Automatic encoding**: Request bodies automatically encoded to BEVE
- ✅ **Automatic decoding**: Response bodies automatically decoded from BEVE
- ✅ **Content-Type handling**: Sets `application/beve` headers automatically
- ✅ **JSON fallback**: Falls back to JSON on BEVE errors
- ✅ **Per-request control**: Enable/disable BEVE per request with `useBeve` option
- ✅ **Debug mode**: Log BEVE usage and compression stats
- ✅ **TypeScript support**: Full type safety with TypeScript
- ✅ **Framework agnostic**: Works with React, Vue, Angular, plain JS

## 📦 Installation

```bash
npm install axios @your-org/beve-ts
```

## 🔧 Quick Start

### 1. Basic Setup

```typescript
import axios from 'axios';
import { setupBeveInterceptor } from './axios-beve';
import { encode, decode } from '@your-org/beve-ts';

const api = axios.create({
  baseURL: 'https://api.example.com',
});

setupBeveInterceptor(api, {
  encode,
  decode,
  enableByDefault: true,
  debug: true,
});

// All requests now use BEVE automatically!
const response = await api.get('/users');
```

### 2. Per-Request Control

```typescript
// Enable BEVE for this request only
await api.post('/users', userData, { useBeve: true });

// Use JSON (default)
await api.get('/legacy-endpoint');
```

### 3. React Integration

```typescript
import { createContext, useContext, useState } from 'react';

const ApiContext = createContext(null);

export function ApiProvider({ children }) {
  const [api] = useState(() => {
    const instance = axios.create({
      baseURL: process.env.REACT_APP_API_URL,
    });
    setupBeveInterceptor(instance, { encode, decode, enableByDefault: true });
    return instance;
  });

  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
}

export function useApi() {
  return useContext(ApiContext);
}
```

## 🎯 Configuration Options

```typescript
interface BeveConfig {
  // BEVE encoder/decoder functions
  encode: (data: any) => Uint8Array;
  decode: (data: Uint8Array) => any;
  
  // Enable BEVE for all requests (default: false)
  enableByDefault?: boolean;
  
  // Fallback to JSON on BEVE errors (default: true)
  fallbackToJson?: boolean;
  
  // Log BEVE usage and stats (default: false)
  debug?: boolean;
}
```

## 📊 Performance Benefits

With the `debug: true` option, you'll see compression stats:

```
[BEVE Request] {
  url: '/api/users',
  originalSize: 1523,   // JSON size
  beveSize: 742,        // BEVE size
  compression: 51.3     // % saved
}
```

**Typical savings:**
- 40-60% smaller payloads than JSON
- 2-3× faster parsing than JSON
- Lower bandwidth costs
- Faster page loads

## 🔄 Backend Integration (Go)

### Fiber Framework

```go
package main

import (
    "github.com/beve-org/beve-go"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // BEVE middleware
    app.Use(func(c *fiber.Ctx) error {
        contentType := string(c.Request().Header.ContentType())
        
        if contentType == "application/beve" {
            // Decode BEVE request
            body := c.Body()
            var data interface{}
            if err := beve.Unmarshal(body, &data); err != nil {
                return c.Status(400).JSON(fiber.Map{"error": "Invalid BEVE"})
            }
            c.Locals("data", data)
        }
        
        return c.Next()
    })

    // Route handler
    app.Post("/users", func(c *fiber.Ctx) error {
        data := c.Locals("data")
        
        // Process data...
        result := map[string]interface{}{
            "id": 123,
            "status": "created",
        }
        
        // Encode BEVE response
        if c.Accepts("application/beve") != "" {
            encoded, _ := beve.Marshal(result)
            c.Set("Content-Type", "application/beve")
            return c.Send(encoded)
        }
        
        // Fallback to JSON
        return c.JSON(result)
    })

    app.Listen(":3000")
}
```

### Standard net/http

```go
package main

import (
    "encoding/json"
    "io"
    "net/http"
    "github.com/beve-org/beve-go"
)

func beveMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        contentType := r.Header.Get("Content-Type")
        
        if contentType == "application/beve" {
            body, _ := io.ReadAll(r.Body)
            var data interface{}
            if err := beve.Unmarshal(body, &data); err != nil {
                http.Error(w, "Invalid BEVE", http.StatusBadRequest)
                return
            }
            r.Body = io.NopCloser(bytes.NewReader(body))
            r.Header.Set("X-Decoded-Format", "beve")
        }
        
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        result := map[string]interface{}{"id": 123}
        
        // Check Accept header
        if r.Header.Get("Accept") == "application/beve" {
            encoded, _ := beve.Marshal(result)
            w.Header().Set("Content-Type", "application/beve")
            w.Write(encoded)
            return
        }
        
        // Fallback JSON
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
    })
    
    http.ListenAndServe(":3000", beveMiddleware(mux))
}
```

## 🧪 Testing

```typescript
import { setupBeveInterceptor } from './axios-beve';
import axios from 'axios';
import { encode, decode } from '@your-org/beve-ts';

describe('BEVE Interceptor', () => {
  it('encodes request body', async () => {
    const api = axios.create();
    setupBeveInterceptor(api, { encode, decode, enableByDefault: true });
    
    const data = { name: 'Test' };
    await api.post('/test', data);
    
    // Request should be BEVE encoded
    expect(lastRequest.headers['Content-Type']).toBe('application/beve');
  });

  it('decodes response', async () => {
    const api = axios.create();
    setupBeveInterceptor(api, { encode, decode, enableByDefault: true });
    
    const response = await api.get('/test');
    
    // Response should be decoded
    expect(response.data).toEqual({ id: 123 });
  });
});
```

## 🔍 Debugging

Enable debug mode to see detailed logs:

```typescript
setupBeveInterceptor(api, {
  encode,
  decode,
  enableByDefault: true,
  debug: true, // ← Enable debug logs
});
```

Output:
```
[BEVE] Interceptor installed { enableByDefault: true, fallbackToJson: true }
[BEVE Request] { url: '/users', originalSize: 1523, beveSize: 742, compression: 51.3 }
[BEVE Response] { url: '/users', beveSize: 856, decodedSize: 1842, expansion: 2.15x }
```

## 🚦 Browser Support

Requires WebAssembly support:
- ✅ Chrome/Edge 57+
- ✅ Firefox 52+
- ✅ Safari 11+
- ✅ Node.js 8+

Check support:
```typescript
import { isBEVESupported } from './axios-beve';

if (!isBEVESupported()) {
  console.warn('BEVE not supported, using JSON');
}
```

## 📝 License

MIT

## 🤝 Contributing

Contributions welcome! See [CONTRIBUTING.md](../../CONTRIBUTING.md)
