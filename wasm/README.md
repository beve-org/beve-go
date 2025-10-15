# BEVE-Go WASM - JSON Translation API

## Overview

BEVE-Go now includes comprehensive JSON ↔ BEVE translation support in WebAssembly, enabling seamless format conversion in browser and Node.js environments.

## Build

```bash
# Build with TinyGo (recommended - produces 377KB binary, 149KB gzipped)
./scripts/build-wasm.sh wasm

# Output: build/wasm/beve.wasm
```

## Installation

### Browser

```html
<script src="wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("beve.wasm"), go.importObject)
        .then((result) => {
            go.run(result.instance);
            // beveWasm is now available globally
        });
</script>
```

### Node.js

```javascript
const fs = require('fs');
require('./wasm_exec.js');

const go = new Go();
const wasmBuffer = fs.readFileSync('./beve.wasm');

WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    go.run(result.instance);
    // beveWasm is now available
});
```

## API Reference

### Core Serialization

#### `beveWasm.marshal(jsObject)`

Encodes a JavaScript object to BEVE binary format.

```javascript
const data = { name: "Alice", age: 30 };
const result = beveWasm.marshal(data);

if (result.error) {
    console.error("Marshal error:", result.error);
} else {
    const beveData = result.data; // Uint8Array
}
```

#### `beveWasm.unmarshal(uint8Array)`

Decodes BEVE binary to JavaScript object.

```javascript
const result = beveWasm.unmarshal(beveData);

if (result.error) {
    console.error("Unmarshal error:", result.error);
} else {
    const jsObject = result.data;
}
```

---

### JSON Translation (NEW in v1.3.0)

#### `beveWasm.fromJson(jsonString | uint8Array)`

Converts JSON string or bytes to BEVE binary format.

**Parameters:**
- `jsonString` (string) - JSON string
- OR `uint8Array` (Uint8Array) - JSON bytes

**Returns:**
```javascript
{
    data: Uint8Array,  // BEVE binary
    size: number       // Size in bytes
}
// OR on error:
{ error: string }
```

**Example:**
```javascript
const json = JSON.stringify({ name: "Bob", age: 25 });
const result = beveWasm.fromJson(json);

if (result.error) {
    console.error(result.error);
} else {
    console.log(`JSON (${json.length} bytes) → BEVE (${result.size} bytes)`);
    const savings = ((json.length - result.size) / json.length * 100).toFixed(1);
    console.log(`Space savings: ${savings}%`);
}
```

#### `beveWasm.toJson(uint8Array, prettyPrint?)`

Converts BEVE binary to JSON string.

**Parameters:**
- `uint8Array` (Uint8Array) - BEVE binary data
- `prettyPrint` (boolean, optional) - Enable pretty printing (default: false)

**Returns:**
```javascript
{
    data: string,  // JSON string
    size: number   // Size in bytes
}
// OR on error:
{ error: string }
```

**Example:**
```javascript
// Compact JSON
const result = beveWasm.toJson(beveData, false);
console.log(result.data); // {"name":"Bob","age":25}

// Pretty-printed JSON
const prettyResult = beveWasm.toJson(beveData, true);
console.log(prettyResult.data);
// {
//   "name": "Bob",
//   "age": 25
// }
```

#### `beveWasm.fromJsonWithStats(jsonString | uint8Array)`

Converts JSON to BEVE with detailed conversion statistics.

**Returns:**
```javascript
{
    data: Uint8Array,
    stats: {
        originalSize: number,   // Original JSON size
        convertedSize: number,  // BEVE size
        ratio: number,          // Compression ratio (0-1)
        savings: number         // Space savings (0-1)
    }
}
// OR on error:
{ error: string }
```

**Example:**
```javascript
const json = JSON.stringify({ users: [...] }); // Large dataset
const result = beveWasm.fromJsonWithStats(json);

if (!result.error) {
    console.log(`Original: ${result.stats.originalSize} bytes`);
    console.log(`BEVE: ${result.stats.convertedSize} bytes`);
    console.log(`Ratio: ${result.stats.ratio.toFixed(3)}`);
    console.log(`Savings: ${(result.stats.savings * 100).toFixed(1)}%`);
}
```

---

### Validation

#### `beveWasm.validateJson(jsonString | uint8Array)`

Validates JSON syntax.

**Returns:**
```javascript
{ valid: boolean }
```

**Example:**
```javascript
const valid = beveWasm.validateJson('{"valid": true}');
console.log(valid.valid); // true

const invalid = beveWasm.validateJson('{invalid json}');
console.log(invalid.valid); // false
```

#### `beveWasm.validateBeve(uint8Array)`

Validates BEVE binary format.

**Returns:**
```javascript
{ valid: boolean }
```

**Example:**
```javascript
const result = beveWasm.validateBeve(beveData);
if (result.valid) {
    console.log("✅ Valid BEVE format");
} else {
    console.log("❌ Invalid BEVE format");
}
```

---

### Utilities

#### `beveWasm.version()`

Returns BEVE library version.

**Returns:** `string` (e.g., "1.3.0-wasm")

#### `beveWasm.benchmark(data, iterations)`

Runs marshal/unmarshal benchmark.

**Parameters:**
- `data` (object) - JavaScript object to benchmark
- `iterations` (number) - Number of iterations

**Returns:**
```javascript
{
    marshal: {
        totalMs: number,
        avgMs: number,
        opsPerSec: number
    },
    unmarshal: {
        totalMs: number,
        avgMs: number,
        opsPerSec: number
    },
    payloadSize: number
}
```

## Usage Examples

### Example 1: API Gateway Translation

```javascript
// Frontend sends JSON, backend expects BEVE
async function sendToBackend(data) {
    const json = JSON.stringify(data);
    const result = beveWasm.fromJson(json);
    
    if (result.error) {
        throw new Error(result.error);
    }
    
    await fetch('/api/endpoint', {
        method: 'POST',
        headers: { 'Content-Type': 'application/beve' },
        body: result.data
    });
}
```

### Example 2: Storage Optimization

```javascript
// Save data in BEVE format (30% smaller)
function saveToLocalStorage(key, data) {
    const json = JSON.stringify(data);
    const result = beveWasm.fromJson(json);
    
    if (!result.error) {
        // Store as base64
        const base64 = btoa(String.fromCharCode(...result.data));
        localStorage.setItem(key, base64);
        
        console.log(`Saved ${result.size} bytes (${
            ((json.length - result.size) / json.length * 100).toFixed(1)
        }% savings)`);
    }
}

function loadFromLocalStorage(key) {
    const base64 = localStorage.getItem(key);
    if (!base64) return null;
    
    // Decode base64 to Uint8Array
    const beveData = Uint8Array.from(atob(base64), c => c.charCodeAt(0));
    
    // Convert BEVE to JSON
    const result = beveWasm.toJson(beveData);
    return result.error ? null : JSON.parse(result.data);
}
```

### Example 3: Real-time Analytics

```javascript
// Collect metrics with compression
class MetricsCollector {
    constructor() {
        this.buffer = [];
    }
    
    addMetric(metric) {
        this.buffer.push(metric);
        
        if (this.buffer.length >= 100) {
            this.flush();
        }
    }
    
    flush() {
        const json = JSON.stringify(this.buffer);
        const result = beveWasm.fromJsonWithStats(json);
        
        if (!result.error) {
            console.log(`Sending ${result.stats.convertedSize} bytes` +
                       ` (${(result.stats.savings * 100).toFixed(1)}% savings)`);
            
            this.send(result.data);
            this.buffer = [];
        }
    }
    
    send(beveData) {
        navigator.sendBeacon('/metrics', beveData);
    }
}
```

### Example 4: Format Migration

```javascript
// Migrate JSON files to BEVE
async function migrateDatabase() {
    const jsonFiles = await fetchFileList();
    const stats = { total: 0, converted: 0, errors: 0, savedBytes: 0 };
    
    for (const file of jsonFiles) {
        try {
            const json = await fetch(file).then(r => r.text());
            const result = beveWasm.fromJsonWithStats(json);
            
            if (result.error) {
                stats.errors++;
                continue;
            }
            
            await saveFile(file.replace('.json', '.beve'), result.data);
            
            stats.total++;
            stats.converted++;
            stats.savedBytes += (result.stats.originalSize - result.stats.convertedSize);
            
        } catch (err) {
            console.error(`Failed to migrate ${file}:`, err);
            stats.errors++;
        }
    }
    
    console.log(`Migrated ${stats.converted}/${stats.total} files`);
    console.log(`Saved ${(stats.savedBytes / 1024).toFixed(2)} KB`);
}
```

## Performance

### Typical Results (Browser - Apple M2 Max)

| Operation | Input Size | Time | Throughput | Savings |
|-----------|-----------|------|------------|---------|
| fromJson (small) | 38 bytes | ~700 ns | 54 MB/s | 13% |
| toJson (small) | 33 bytes | ~1 μs | 33 MB/s | - |
| fromJson (medium) | 383 bytes | ~3.8 μs | 100 MB/s | 34% |
| toJson (medium) | 254 bytes | ~4.7 μs | 54 MB/s | - |
| fromJson (large) | 100 KB | ~400 μs | 250 MB/s | 20% |

### Space Savings

- Small objects (< 100 bytes): **10-25% smaller**
- Medium payloads (100 B - 10 KB): **20-35% smaller**
- Large datasets (> 10 KB): **15-30% smaller**
- Arrays of numbers: **40-50% smaller** (typed arrays)

## Testing

```bash
# Start a local server
cd build/wasm
python3 -m http.server 8080

# Open in browser
open http://localhost:8080/test-json.html
```

The test page runs 8 comprehensive tests:
1. Simple object conversion
2. Round-trip verification
3. Pretty-print formatting
4. Statistics tracking
5. JSON validation
6. BEVE validation
7. Large array performance
8. Complex nested structures

## Type Support

| JSON Type | BEVE Type | Notes |
|-----------|-----------|-------|
| null | null | Direct mapping |
| boolean | boolean | Direct mapping |
| number (integer) | int64 | Whole numbers |
| number (float) | float64 | Decimals |
| string | UTF-8 string | Full Unicode support |
| array | typed/generic array | Optimized for homogeneous arrays |
| object | object | String keys only |

## Error Handling

All functions return either a result object or an error object:

```javascript
const result = beveWasm.fromJson(json);

if (result.error) {
    // Handle error
    console.error("Conversion failed:", result.error);
} else {
    // Use result.data
    console.log("Success:", result.data);
}
```

Common errors:
- `"fromJson requires exactly 1 argument"` - Missing parameter
- `"invalid JSON"` - Malformed JSON syntax
- `"marshal failed"` - Internal encoding error
- `"unmarshal failed"` - Invalid BEVE format

## Browser Compatibility

- ✅ Chrome 57+
- ✅ Firefox 52+
- ✅ Safari 11+
- ✅ Edge 16+
- ✅ Node.js 14+ (with `--experimental-wasm-bigint`)

## Size Comparison

| Build Method | Binary Size | Gzipped | Notes |
|-------------|-------------|---------|-------|
| TinyGo | 377 KB | 149 KB | Recommended |
| Go (standard) | 2.1 MB | 567 KB | Not recommended for WASM |

**Why TinyGo?**
- 5.5× smaller binary
- 3.8× smaller gzipped
- Optimized for WASM
- No runtime overhead

## Limitations

1. **No streaming support** - Entire payload must fit in memory
2. **Time.Time encoding** - Use Unix timestamps (int64)
3. **Large integers** - JavaScript numbers are float64 (safe range: ±2^53)
4. **Binary data** - Encode as base64 strings or use Uint8Array

## Changelog

### v1.3.0 (2025-10-15)
- ✨ Added `fromJson` - JSON → BEVE conversion
- ✨ Added `toJson` - BEVE → JSON conversion
- ✨ Added `fromJsonWithStats` - Conversion with statistics
- ✨ Added `validateJson` - JSON syntax validation
- ✨ Added `validateBeve` - BEVE format validation
- 📦 Updated version to 1.3.0-wasm
- 🎨 Added comprehensive test HTML page

### v1.2.0
- Core marshal/unmarshal support
- Benchmark utilities
- TinyGo build optimization

## License

MIT License - See [LICENSE](../../LICENSE) for details.

## Links

- [BEVE Specification](../../SPECIFICATION.md)
- [Go Package Documentation](../../README.md)
- [Translator Package](../../translator/README.md)
- [Examples](../../examples/)
