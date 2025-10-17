# 📦 Installation Guide

Get BEVE-Go up and running in 2 minutes.

---

## Requirements

**Go Version**: 1.21 or later

BEVE-Go requires Go 1.21+ for performance features:
- Per-P local caching in sync.Pool
- Improved reflection performance
- Better memory management

Check your Go version:
```bash
go version
# Should show: go version go1.21.0 or higher
```

**Upgrade Go** (if needed):
- **macOS**: `brew upgrade go`
- **Linux**: [Download from golang.org](https://golang.org/dl/)
- **Windows**: [Download installer](https://golang.org/dl/)

---

## Installation

### Option 1: Go Modules (Recommended)

```bash
go get github.com/beve-org/beve-go
```

This installs the latest stable version.

### Option 2: Specific Version

```bash
# Install specific version
go get github.com/beve-org/beve-go@v1.3.0

# Install latest from main branch (development)
go get github.com/beve-org/beve-go@main
```

### Option 3: Manual Installation

```bash
# Clone repository
git clone https://github.com/beve-org/beve-go.git
cd beve-go

# Install dependencies
go mod download

# Run tests to verify
go test ./...
```

---

## Verify Installation

Create a test file `test_beve.go`:

```go
package main

import (
    "fmt"
    beve "github.com/beve-org/beve-go"
)

func main() {
    // Test struct
    type User struct {
        Name string `beve:"name"`
        Age  int    `beve:"age"`
    }

    // Encode
    user := User{Name: "Alice", Age: 30}
    data, err := beve.Marshal(&user)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Encoded: %d bytes\n", len(data))

    // Decode
    var decoded User
    err = beve.Unmarshal(data, &decoded)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decoded: %+v\n", decoded)
    
    // Success!
    fmt.Println("✅ BEVE-Go is working!")
}
```

Run the test:
```bash
go run test_beve.go
```

Expected output:
```
Encoded: 23 bytes
Decoded: {Name:Alice Age:30}
✅ BEVE-Go is working!
```

---

## IDE Setup

### VS Code

**Recommended Extension**: [BEVE Syntax Highlighting](https://marketplace.visualstudio.com/items?itemName=beve.beve-vscode)

Install extension:
```bash
code --install-extension beve.beve-vscode
```

Features:
- Syntax highlighting for `.beve` files
- BEVE binary viewer
- Hover documentation
- Auto-completion for struct tags

### GoLand / IntelliJ IDEA

1. Open **Settings** → **Plugins**
2. Search for "BEVE"
3. Install **BEVE Support**
4. Restart IDE

### Vim / Neovim

Add to your config:
```vim
" Install vim-beve plugin
Plug 'beve-org/vim-beve'

" Enable syntax highlighting
autocmd BufRead,BufNewFile *.beve set filetype=beve
```

---

## Project Setup

### Initialize New Project

```bash
# Create project directory
mkdir my-beve-app
cd my-beve-app

# Initialize Go module
go mod init my-beve-app

# Install BEVE-Go
go get github.com/beve-org/beve-go

# Create main.go
cat > main.go << 'EOF'
package main

import (
    "fmt"
    beve "github.com/beve-org/beve-go"
)

type Message struct {
    Text      string `beve:"text"`
    Timestamp int64  `beve:"timestamp"`
}

func main() {
    msg := Message{Text: "Hello BEVE!", Timestamp: 1697500000}
    data, _ := beve.Marshal(msg)
    fmt.Printf("Message: %d bytes\n", len(data))
}
EOF

# Run
go run main.go
```

### Add to Existing Project

```bash
# In your project directory
go get github.com/beve-org/beve-go

# Update imports in your code
import beve "github.com/beve-org/beve-go"
```

---

## Dependencies

BEVE-Go has **zero external dependencies** for core functionality.

Optional dependencies (for extras):
- **Benchmarking**: `go test` built-in
- **Examples**: Standard library only
- **WASM**: Go 1.21+ WASM support

View dependencies:
```bash
go list -m all
```

Output:
```
github.com/beve-org/beve-go v1.3.0
# No external dependencies! 🎉
```

---

## Platform Support

BEVE-Go is tested on:

| Platform | Architecture | Status |
|----------|-------------|--------|
| **macOS** | ARM64 (M1/M2/M3) | ✅ Fully supported |
| **macOS** | AMD64 (Intel) | ✅ Fully supported |
| **Linux** | ARM64 (Graviton, Neoverse) | ✅ Fully supported |
| **Linux** | AMD64 (x86_64) | ✅ Fully supported |
| **Windows** | AMD64 | ✅ Fully supported |
| **Windows** | ARM64 | 🟡 Experimental |
| **FreeBSD** | AMD64 | 🟡 Community supported |
| **WASM** | wasm32 | ✅ Fully supported |

**SIMD Optimizations**:
- ✅ ARM64 NEON (M1, Graviton)
- ✅ AMD64 AVX2 (modern Intel/AMD)
- ✅ Automatic CPU detection
- ✅ Graceful fallback to scalar code

---

## Docker Setup

### Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o /app/server ./cmd/server

# Runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server /usr/local/bin/server

EXPOSE 8080
CMD ["server"]
```

Build and run:
```bash
docker build -t my-beve-app .
docker run -p 8080:8080 my-beve-app
```

### Docker Compose

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - BEVE_LOG_LEVEL=info
    volumes:
      - ./config:/config
```

---

## Troubleshooting

### Issue: "module not found"

**Solution**: Ensure Go modules are enabled:
```bash
export GO111MODULE=on
go mod tidy
```

### Issue: "go version too old"

**Solution**: Upgrade to Go 1.21+:
```bash
# macOS
brew upgrade go

# Linux (download latest)
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```

### Issue: Build fails on Windows

**Solution**: Use Go 1.21+ and ensure CGO is disabled:
```bash
set CGO_ENABLED=0
go build
```

### Issue: Import cycle detected

**Solution**: BEVE-Go has no import cycles. Check your own imports:
```bash
go list -f '{{.ImportPath}}: {{.Imports}}' ./...
```

---

## Uninstallation

Remove BEVE-Go from your project:

```bash
# Remove from go.mod
go mod edit -droprequire github.com/beve-org/beve-go

# Clean up
go mod tidy

# Remove from imports
# (manually edit your .go files)
```

---

## Next Steps

✅ **Installation complete!** Now learn:

1. **[Quick Start →](quick-start.md)** - 5-minute tutorial
2. **[Basic Usage →](basic-usage.md)** - Common patterns
3. **[Migrating from JSON →](json-migration.md)** - Switch from `encoding/json`

---

## Getting Help

- 📖 **Documentation**: [docs/INDEX.md](../INDEX.md)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/beve-org/beve-go/discussions)
- 🐛 **Issues**: [GitHub Issues](https://github.com/beve-org/beve-go/issues)
- 📧 **Email**: support@beve.org

---

**Installation Time**: ~2 minutes  
**Next**: [Quick Start Tutorial →](quick-start.md)
