# Production Troubleshooting Guide

**BEVE-Go Debugging & Issue Resolution**

Comprehensive troubleshooting guide for diagnosing and fixing production issues with BEVE-Go applications.

---

## Table of Contents

1. [Common Issues](#common-issues)
2. [Performance Debugging](#performance-debugging)
3. [Memory Issues](#memory-issues)
4. [Decode Errors](#decode-errors)
5. [Production Runbooks](#production-runbooks)
6. [Debugging Tools](#debugging-tools)

---

## Common Issues

### Issue 1: High Memory Usage

**Symptoms**:
- Memory steadily increasing
- OOM kills in Kubernetes
- GC pauses increasing

**Causes**:
1. Buffer pool not returning buffers
2. Memory leaks in user code
3. Large messages accumulating

**Solutions**:

```go
// ❌ BAD: Buffer leak
func processData(data interface{}) []byte {
    enc := beve.NewEncoder()
    result := enc.Marshal(data)
    // LEAK! Buffer never returned
    return result
}

// ✅ GOOD: Proper buffer management
func processData(data interface{}) []byte {
    enc := beve.GetEncoderFromPool()
    defer beve.PutEncoderToPool(enc)
    
    result := enc.Marshal(data)
    return result
}
```

**Diagnosis**:

```bash
# Check memory usage
go tool pprof http://localhost:6060/debug/pprof/heap

# Top memory allocators
(pprof) top10

# Find leak source
(pprof) list processData
```

**Fix**:

```go
// Add finalizer to detect leaks
func NewEncoder() *Encoder {
    enc := &Encoder{}
    
    runtime.SetFinalizer(enc, func(e *Encoder) {
        if e.buf != nil {
            log.Println("WARNING: Encoder not returned to pool")
        }
    })
    
    return enc
}
```

---

### Issue 2: Slow Performance

**Symptoms**:
- p99 latency > 100ms
- Throughput drop
- CPU usage high

**Causes**:
1. Reflection cache miss
2. Large allocations
3. GC pressure

**Solutions**:

```go
// ❌ BAD: Cache miss on every call
func processUsers(users []User) {
    for _, user := range users {
        beve.Marshal(user)  // Reflection on every iteration
    }
}

// ✅ GOOD: Pre-warm cache
func init() {
    // Warm up reflection cache at startup
    beve.Marshal(User{})
}

func processUsers(users []User) {
    for _, user := range users {
        beve.Marshal(user)  // Cache hit
    }
}
```

**Diagnosis**:

```bash
# CPU profiling
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Find hot spots
(pprof) top10
(pprof) list Marshal
```

**Optimization**:

```go
// Use zero-copy mode for large payloads
opts := beve.MarshalOptions{
    ZeroCopy: true,  // 2-8× faster
}
data, _ := beve.MarshalWithOptions(largeData, opts)
```

---

### Issue 3: Decode Errors

**Symptoms**:
- Random unmarshal failures
- `invalid BEVE header` errors
- Data corruption

**Causes**:
1. Corrupted data
2. Version mismatch
3. Truncated messages

**Solutions**:

```go
// Validate before unmarshal
func safeUnmarshal(data []byte, v interface{}) error {
    // 1. Check minimum size
    if len(data) < 1 {
        return errors.New("empty payload")
    }
    
    // 2. Validate header
    if !beve.IsValidBEVE(data) {
        return errors.New("invalid BEVE format")
    }
    
    // 3. Checksum validation (if available)
    if hasChecksum(data) {
        if !validateChecksum(data) {
            return errors.New("checksum mismatch")
        }
    }
    
    // 4. Unmarshal
    return beve.Unmarshal(data, v)
}
```

**Hex Dump Analysis**:

```go
// Debug corrupted data
func debugBEVE(data []byte) {
    fmt.Println("=== BEVE Hex Dump ===")
    fmt.Printf("Size: %d bytes\n", len(data))
    fmt.Printf("Header: 0x%02x (type: %d)\n", data[0], data[0]&0x07)
    
    // Dump first 64 bytes
    for i := 0; i < min(64, len(data)); i += 16 {
        end := min(i+16, len(data))
        fmt.Printf("%04x: ", i)
        
        for j := i; j < end; j++ {
            fmt.Printf("%02x ", data[j])
        }
        
        fmt.Printf("\n")
    }
}
```

---

### Issue 4: Buffer Pool Exhaustion

**Symptoms**:
- `beve_buffer_pool_hit_rate` < 0.8
- High allocation rate
- GC thrashing

**Causes**:
1. Pool size too small
2. Buffers not returned
3. Burst traffic

**Solutions**:

```go
// Increase pool size
beve.SetBufferPoolSize(50000)  // Was 10000

// Monitor pool stats
stats := beve.GetPoolStats()
fmt.Printf("Hit rate: %.2f%%\n", stats.HitRate()*100)
fmt.Printf("Available: %d/%d\n", stats.AvailableBuffers, stats.TotalBuffers)
```

**Auto-scaling Pool**:

```go
// auto_pool.go
package pool

import "sync/atomic"

type AutoScalingPool struct {
    currentSize  int32
    minSize      int32
    maxSize      int32
    missThreshold float64
}

func (p *AutoScalingPool) Monitor() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        stats := beve.GetPoolStats()
        hitRate := stats.HitRate()
        
        if hitRate < p.missThreshold {
            // Scale up
            newSize := int32(float64(p.currentSize) * 1.5)
            if newSize > p.maxSize {
                newSize = p.maxSize
            }
            
            atomic.StoreInt32(&p.currentSize, newSize)
            beve.SetBufferPoolSize(int(newSize))
            
            log.Printf("Scaled pool up to %d (hit rate: %.2f%%)", newSize, hitRate*100)
        }
    }
}
```

---

### Issue 5: High CPU Usage

**Symptoms**:
- CPU consistently > 80%
- Slow response times
- Increased latency

**Causes**:
1. Inefficient encoding/decoding
2. GC overhead
3. Reflection overhead

**Solutions**:

```go
// Use struct tags to avoid reflection
type User struct {
    Name string `beve:"name"`  // Explicit field name
    Age  int    `beve:"age"`
}

// Use typed arrays for homogeneous data
users := []User{...}
data, _ := beve.MarshalTyped(users)  // Extension 1
```

**GOMAXPROCS Tuning**:

```go
import "runtime"

func init() {
    // Limit CPU cores if needed
    numCPU := runtime.NumCPU()
    
    if numCPU > 8 {
        runtime.GOMAXPROCS(8)  // Cap at 8 cores
    }
}
```

---

## Performance Debugging

### Enable pprof

```go
// main.go
import _ "net/http/pprof"

func main() {
    // pprof server (localhost only!)
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your app
    startServer()
}
```

### CPU Profiling

```bash
# Collect 30s CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Analyze
(pprof) top10
(pprof) list beve.Marshal
(pprof) web  # Visualize (requires graphviz)
```

### Memory Profiling

```bash
# Heap profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Allocation profile
go tool pprof http://localhost:6060/debug/pprof/allocs

# Find allocations
(pprof) top10 -alloc_space
(pprof) list NewEncoder
```

### Trace Analysis

```bash
# Collect trace
curl http://localhost:6060/debug/pprof/trace?seconds=5 > trace.out

# Analyze
go tool trace trace.out
```

---

## Memory Issues

### Detecting Memory Leaks

```go
// memory_monitor.go
package monitor

import (
    "runtime"
    "time"
)

func MonitorMemory() {
    var lastAlloc uint64
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        // Check for leak (steady increase)
        if m.Alloc > lastAlloc*2 && lastAlloc > 0 {
            log.Printf("WARNING: Possible memory leak (was: %d MB, now: %d MB)",
                lastAlloc/1024/1024, m.Alloc/1024/1024)
            
            // Force GC
            runtime.GC()
            
            // Re-check after GC
            runtime.ReadMemStats(&m)
            
            if m.Alloc > lastAlloc*1.5 {
                log.Printf("CRITICAL: Memory leak confirmed")
                // Trigger heap dump
                takeHeapDump()
            }
        }
        
        lastAlloc = m.Alloc
    }
}

func takeHeapDump() {
    f, _ := os.Create(fmt.Sprintf("heap-%d.prof", time.Now().Unix()))
    defer f.Close()
    pprof.WriteHeapProfile(f)
}
```

### Analyzing Heap Dumps

```bash
# Compare two heap dumps
go tool pprof -base heap-1697452800.prof heap-1697453100.prof

# Find growth
(pprof) top10 -alloc_space
(pprof) list NewEncoder
```

### GC Tuning

```bash
# Set memory limit
export GOMEMLIMIT=8GiB

# Increase GC target (less frequent GC)
export GOGC=200  # Default: 100

# Disable GC (debugging only!)
export GOGC=off
```

---

## Decode Errors

### Error Categories

**1. Invalid Header**:

```go
// Error: invalid BEVE header
// Cause: Corrupted data or wrong format

func diagnoseHeader(data []byte) {
    if len(data) == 0 {
        log.Println("Empty payload")
        return
    }
    
    header := data[0]
    typeID := header & 0x07
    
    log.Printf("Header: 0x%02x, Type ID: %d", header, typeID)
    
    if typeID > 6 {
        log.Println("INVALID: Type ID out of range (0-6)")
    }
}
```

**2. Truncated Message**:

```go
// Error: unexpected EOF
// Cause: Incomplete message

func validateMessageComplete(data []byte) error {
    // Simple check: does size match expected?
    expectedSize := calculateExpectedSize(data)
    
    if len(data) < expectedSize {
        return fmt.Errorf("truncated: have %d bytes, need %d", len(data), expectedSize)
    }
    
    return nil
}
```

**3. Type Mismatch**:

```go
// Error: cannot unmarshal X into Go value of type Y
// Cause: Schema mismatch

type User struct {
    Name string
    Age  int
}

var data []byte  // BEVE encoded as {name: "Alice", age: "30"}  (string age!)

var user User
err := beve.Unmarshal(data, &user)
// Error: cannot unmarshal string into int

// Solution: Use interface{} for flexible schemas
var flexUser map[string]interface{}
beve.Unmarshal(data, &flexUser)
```

### Comparative Testing

```go
// Compare BEVE vs JSON
func compareBEVEJSON(data interface{}) {
    // Marshal both
    beveData, beveErr := beve.Marshal(data)
    jsonData, jsonErr := json.Marshal(data)
    
    if beveErr != nil && jsonErr == nil {
        log.Printf("BEVE failed, JSON succeeded: %v", beveErr)
        
        // Try unmarshal
        var v1, v2 interface{}
        beve.Unmarshal(beveData, &v1)
        json.Unmarshal(jsonData, &v2)
        
        // Deep compare
        if !reflect.DeepEqual(v1, v2) {
            log.Println("Data mismatch between BEVE and JSON")
        }
    }
}
```

---

## Production Runbooks

### Runbook 1: High Memory Alert

**Trigger**: `beve_allocated_bytes > 2GB` for 5 min

**Steps**:

1. **Check pool stats**:
   ```bash
   curl http://localhost:8080/metrics | grep beve_buffer_pool
   ```

2. **Force GC**:
   ```bash
   curl -X POST http://localhost:8080/debug/gc
   ```

3. **If memory still high**, take heap dump:
   ```bash
   curl http://localhost:6060/debug/pprof/heap > heap.prof
   ```

4. **Analyze**:
   ```bash
   go tool pprof heap.prof
   (pprof) top10
   ```

5. **If leak confirmed**, restart pod:
   ```bash
   kubectl delete pod beve-app-xyz
   ```

---

### Runbook 2: High Latency

**Trigger**: `p99 > 10ms` for 2 min

**Steps**:

1. **Check throughput**:
   ```bash
   curl http://localhost:8080/metrics | grep beve_operations_per_second
   ```

2. **CPU profile**:
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
   ```

3. **Check buffer pool**:
   ```bash
   # Hit rate should be > 95%
   curl http://localhost:8080/metrics | grep beve_buffer_pool_hit_rate
   ```

4. **If pool exhausted**, scale pool:
   ```go
   beve.SetBufferPoolSize(50000)
   ```

5. **If CPU-bound**, scale horizontally:
   ```bash
   kubectl scale deployment beve-app --replicas=6
   ```

---

### Runbook 3: Decode Error Spike

**Trigger**: `beve_errors_total{type="unmarshal"} > 10/sec`

**Steps**:

1. **Sample error logs**:
   ```bash
   kubectl logs beve-app-xyz | grep "unmarshal failed" | head -20
   ```

2. **Check for pattern**:
   ```bash
   # Are all errors from same client?
   kubectl logs beve-app-xyz | grep "unmarshal failed" | grep -o "client_id=[^ ]*" | sort | uniq -c
   ```

3. **If malicious**, block client:
   ```go
   rateLimiter.Block(clientID, 1*time.Hour)
   ```

4. **If data corruption**, check upstream:
   ```bash
   # Inspect producer logs
   kubectl logs producer-service | grep ERROR
   ```

5. **Enable verbose logging**:
   ```go
   log.SetLevel(log.DEBUG)
   ```

---

## Debugging Tools

### 1. BEVE Inspector

```go
// tools/inspector.go
package tools

import "fmt"

func InspectBEVE(data []byte) {
    fmt.Println("=== BEVE Inspector ===")
    fmt.Printf("Total size: %d bytes\n\n", len(data))
    
    offset := 0
    depth := 0
    
    for offset < len(data) {
        header := data[offset]
        typeID := header & 0x07
        
        indent := strings.Repeat("  ", depth)
        fmt.Printf("%s[%04x] Type: %s (0x%02x)\n", indent, offset, typeName(typeID), header)
        
        // Parse based on type
        switch typeID {
        case 0x03:  // Object
            size, n := readVarint(data[offset+1:])
            fmt.Printf("%s  Keys: %d\n", indent, size)
            offset += n + 1
            depth++
        case 0x04:  // Typed Array
            size, n := readVarint(data[offset+1:])
            fmt.Printf("%s  Elements: %d\n", indent, size)
            offset += n + 1
        default:
            offset++
        }
    }
}

func typeName(id byte) string {
    names := []string{"Null/Bool", "Number", "String", "Object", "TypedArray", "GenericArray", "Extension"}
    if id < 7 {
        return names[id]
    }
    return "Invalid"
}
```

### 2. Diff Tool

```go
// tools/diff.go
package tools

func DiffBEVE(data1, data2 []byte) {
    fmt.Println("=== BEVE Diff ===")
    
    if bytes.Equal(data1, data2) {
        fmt.Println("Identical")
        return
    }
    
    // Find first difference
    for i := 0; i < min(len(data1), len(data2)); i++ {
        if data1[i] != data2[i] {
            fmt.Printf("First diff at offset %d:\n", i)
            fmt.Printf("  Data1: 0x%02x\n", data1[i])
            fmt.Printf("  Data2: 0x%02x\n", data2[i])
            
            // Context (16 bytes before/after)
            start := max(0, i-16)
            end := min(len(data1), i+16)
            
            fmt.Printf("\nContext (data1):\n")
            dumpHex(data1[start:end], start)
            
            fmt.Printf("\nContext (data2):\n")
            dumpHex(data2[start:end], start)
            
            return
        }
    }
    
    fmt.Printf("Size diff: %d vs %d bytes\n", len(data1), len(data2))
}
```

### 3. Performance Harness

```go
// tools/harness.go
package tools

func BenchmarkOperation(fn func() error, iterations int) {
    var durations []time.Duration
    
    for i := 0; i < iterations; i++ {
        start := time.Now()
        err := fn()
        duration := time.Since(start)
        
        if err != nil {
            log.Printf("Error at iteration %d: %v", i, err)
            continue
        }
        
        durations = append(durations, duration)
    }
    
    // Statistics
    sort.Slice(durations, func(i, j int) bool {
        return durations[i] < durations[j]
    })
    
    p50 := durations[len(durations)/2]
    p95 := durations[int(float64(len(durations))*0.95)]
    p99 := durations[int(float64(len(durations))*0.99)]
    
    fmt.Printf("Results (%d iterations):\n", len(durations))
    fmt.Printf("  p50: %v\n", p50)
    fmt.Printf("  p95: %v\n", p95)
    fmt.Printf("  p99: %v\n", p99)
}
```

---

## Summary

### Quick Diagnosis

| Symptom | Likely Cause | Quick Fix |
|---------|-------------|-----------|
| High memory | Buffer leak | Check `defer PutEncoderToPool()` |
| Slow performance | Cache miss | Pre-warm reflection cache |
| Decode errors | Corrupted data | Validate with `IsValidBEVE()` |
| Pool exhaustion | Pool too small | Increase `SetBufferPoolSize()` |
| High CPU | Inefficient code | Use zero-copy mode |

### Debug Checklist

- [ ] Enable pprof server (`localhost:6060`)
- [ ] Configure Prometheus metrics
- [ ] Set up structured logging
- [ ] Add request tracing
- [ ] Implement health checks
- [ ] Create runbooks for common issues
- [ ] Test error scenarios
- [ ] Monitor buffer pool hit rate
- [ ] Profile CPU/memory regularly
- [ ] Keep heap dumps for analysis

### Tools

- **pprof**: CPU/memory profiling
- **trace**: Goroutine/blocking analysis
- **BEVE Inspector**: Hex dump analysis
- **Diff Tool**: Compare BEVE payloads
- **Performance Harness**: Benchmark operations

---

**Next**: [Deployment](deployment.md) · [Monitoring](monitoring.md) · [Security](security.md)
