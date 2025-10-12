# 🌐 BEVE WebAssembly Demo

High-performance binary serialization in the browser using Go + WebAssembly.

## 📦 What's Included

- **`beve.wasm`** - Compiled BEVE library (350KB, 106KB gzipped)
- **`wasm_exec.js`** - Go WebAssembly glue code
- **`index.html`** - Interactive demo with benchmarks

## 🚀 Features

- ✅ **Marshal/Unmarshal** JavaScript objects in the browser
- ✅ **Interactive UI** with real-time size comparison
- ✅ **Performance Benchmarks** (ops/sec, avg time)
- ✅ **Example Payloads** (small, medium, large structs)
- ✅ **Binary Hex Viewer** for encoded data
- ✅ **~50K ops/sec** on modern browsers

## 🔧 How to Build

```bash
# From project root
./scripts/build-wasm.sh wasm
```

This creates:
- `build/wasm/beve.wasm` (350KB raw, 106KB gzipped)
- `build/wasm/wasm_exec.js` (TinyGo glue code)

## 🌍 Running the Demo

### Option 1: Python HTTP Server

```bash
python3 -m http.server 8080
# Open: http://localhost:8080/build/wasm/
```

### Option 2: Node.js `http-server`

```bash
npm install -g http-server
http-server -p 8080
# Open: http://localhost:8080/build/wasm/
```

### Option 3: Go HTTP Server

```bash
go run -C build/wasm server.go
# Open: http://localhost:8080/
```

## 📖 JavaScript API

### Loading WASM Module

```javascript
const go = new Go();
WebAssembly.instantiateStreaming(fetch('beve.wasm'), go.importObject)
  .then(result => {
    go.run(result.instance);
    // beveWasm is now available globally
  });
```

### Marshal (Encode)

```javascript
const data = {
  id: 123,
  name: "Alice",
  email: "alice@example.com",
  active: true,
  tags: ["developer", "golang"]
};

const result = beveWasm.marshal(data);
if (result.error) {
  console.error('Marshal failed:', result.error);
} else {
  console.log('Encoded bytes:', result.data);
  console.log('Size:', result.data.length, 'bytes');
}
```

### Unmarshal (Decode)

```javascript
const encoded = result.data; // Uint8Array from marshal

const decoded = beveWasm.unmarshal(encoded);
if (decoded.error) {
  console.error('Unmarshal failed:', decoded.error);
} else {
  console.log('Decoded object:', decoded.data);
}
```

### Version

```javascript
const version = beveWasm.version();
console.log('BEVE Version:', version);
// Output: "1.2.0-wasm"
```

### Benchmark

```javascript
const testData = {id: 1, name: "Test", value: 42};
const iterations = 10000;

const bench = beveWasm.benchmark(testData, iterations);

console.log('Marshal Results:');
console.log('  Avg:', bench.marshal.avgMs.toFixed(4), 'ms');
console.log('  Throughput:', Math.round(bench.marshal.opsPerSec).toLocaleString(), 'ops/sec');

console.log('Unmarshal Results:');
console.log('  Avg:', bench.unmarshal.avgMs.toFixed(4), 'ms');
console.log('  Throughput:', Math.round(bench.unmarshal.opsPerSec).toLocaleString(), 'ops/sec');

console.log('Payload Size:', bench.payloadSize, 'bytes');
```

## 🎯 Performance

Tested on **Apple M2 Max** (Chrome 130):

| Operation | Throughput | Avg Time |
|-----------|-----------|----------|
| **Marshal** | ~55K ops/sec | ~18μs |
| **Unmarshal** | ~45K ops/sec | ~22μs |
| **Round-trip** | ~25K ops/sec | ~40μs |

### Size Comparison (Medium Struct)

| Format | Size | vs JSON |
|--------|------|---------|
| **JSON** | 170 bytes | 100% |
| **BEVE** | 120 bytes | **-29%** |
| **CBOR** | 115 bytes | -32% |

> BEVE achieves near-CBOR efficiency while maintaining Go-native types!

## 🔬 Technical Details

### Build Configuration

**TinyGo Compiler:**
- Version: 0.39.0+
- Target: `wasm`
- Optimization: `-opt=2`
- GC Mode: `-gc=leaking` (minimal GC for WASM)
- Debug: `-no-debug` (strip symbols)

**Output Size:**
- Uncompressed: 350 KB
- Gzipped: **106 KB** ✅
- Brotli: ~85 KB

### Browser Compatibility

✅ **Modern Browsers** (2020+):
- Chrome 90+
- Firefox 89+
- Safari 15+
- Edge 90+

⚠️ **Requirements**:
- WebAssembly support
- ES6+ (async/await, Promises)
- Fetch API

### Memory Usage

- **Initial**: ~2-4 MB (WASM runtime)
- **Per operation**: <1 KB (pooled buffers)
- **GC mode**: Leaking (manual management)

## 🛠️ Integration Examples

### React Component

```javascript
import { useEffect, useState } from 'react';

function BeveEncoder() {
  const [beve, setBeve] = useState(null);
  
  useEffect(() => {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch('/beve.wasm'), go.importObject)
      .then(result => {
        go.run(result.instance);
        setBeve(window.beveWasm);
      });
  }, []);
  
  const handleEncode = (data) => {
    if (!beve) return;
    const result = beve.marshal(data);
    console.log('Encoded:', result.data);
  };
  
  return <button onClick={() => handleEncode({test: 123})}>Encode</button>;
}
```

### Node.js (WASI)

```bash
# Build WASI target
./scripts/build-wasm.sh wasi

# Run with Node.js
node --experimental-wasi-unstable-preview1 test.js
```

### Edge Functions (Cloudflare Workers)

```javascript
// worker.js
import { instantiate } from './beve.wasm';

export default {
  async fetch(request) {
    const wasm = await instantiate();
    const data = {id: 1, name: "Edge Data"};
    const encoded = wasm.marshal(data);
    
    return new Response(encoded.data, {
      headers: {'Content-Type': 'application/beve'}
    });
  }
}
```

## 🐛 Troubleshooting

### "Failed to load WASM module"

**Solution**: Ensure server has correct MIME type:
```nginx
# nginx.conf
types {
    application/wasm wasm;
}
```

### "beveWasm is undefined"

**Solution**: Wait for WASM to load:
```javascript
await new Promise(resolve => setTimeout(resolve, 100));
console.log(typeof beveWasm); // "object"
```

### High memory usage

**Solution**: GC mode is "leaking" for performance. If memory is critical:
```bash
# Rebuild with conservative GC
tinygo build -target=wasm -gc=conservative -o beve.wasm ./wasm/main.go
```

## 📚 Resources

- **TinyGo Docs**: https://tinygo.org/docs/guides/webassembly/
- **Go WASM Docs**: https://github.com/golang/go/wiki/WebAssembly
- **BEVE Spec**: [SPECIFICATION_COMPLIANCE.md](../../SPECIFICATION_COMPLIANCE.md)

## 🤝 Contributing

Found a bug or have a feature request for WASM support? 

[Open an issue](https://github.com/beve-org/beve-go/issues) or submit a PR!

---

**Built with ❤️ for the modern web**  
🌐 **WebAssembly** | 🚀 **TinyGo** | ⚡ **BEVE**
