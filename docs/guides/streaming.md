# 🌊 Streaming Guide

Master BEVE's streaming encoder/decoder for efficient batch processing and network communication.

**Reading Time**: 18 minutes  
**Level**: Intermediate  
**Prerequisites**: [Encoding/Decoding](encoding-decoding.md)

---

## Table of Contents

1. [Why Streaming?](#why-streaming)
2. [Stream Encoder](#stream-encoder)
3. [Stream Decoder](#stream-decoder)
4. [Buffering Strategies](#buffering-strategies)
5. [Network Streaming](#network-streaming)
6. [File Streaming](#file-streaming)
7. [Advanced Patterns](#advanced-patterns)
8. [Performance Optimization](#performance-optimization)

---

## Why Streaming?

### Problem: Memory Overhead

**❌ Bad: Load everything into memory**:
```go
// Load 1 million records into memory
users := make([]User, 1_000_000)
// Memory: ~500MB

// Marshal all at once
data, _ := beve.Marshal(users)
// Memory: ~1GB total (users + encoded data)
```

**✅ Good: Stream one at a time**:
```go
stream := beve.NewStreamEncoder(writer)

for _, user := range userIterator {
    stream.Encode(user)
    // Memory: ~100KB (only current user + buffer)
}
```

### Benefits

| Metric | Batch (Marshal) | Streaming |
|--------|----------------|-----------|
| **Memory** | O(N) - All data | O(1) - Constant |
| **Latency** | High (wait for all) | Low (progressive) |
| **Throughput** | Lower (single batch) | Higher (pipeline) |
| **Error Recovery** | All-or-nothing | Partial success |

### Use Cases

**Stream Encoder**:
- ✅ Large file generation
- ✅ Network responses (HTTP, gRPC)
- ✅ Batch ETL jobs
- ✅ Log aggregation
- ✅ Real-time data feeds

**Stream Decoder**:
- ✅ Large file parsing
- ✅ Network requests
- ✅ Database cursors
- ✅ Queue consumers
- ✅ Stream processing

---

## Stream Encoder

### Basic Usage

```go
// Create stream encoder
stream := beve.NewStreamEncoder(writer)
defer stream.Close()

// Encode multiple values
for _, user := range users {
    if err := stream.Encode(user); err != nil {
        return err
    }
}

// Flush remaining data
if err := stream.Flush(); err != nil {
    return err
}
```

### Constructor Options

#### 1. Default Buffer (8KB)

```go
stream := beve.NewStreamEncoder(writer)
// Buffer: 8KB (default)
// Auto-flush: When buffer full
```

#### 2. Custom Buffer Size

```go
stream := beve.NewStreamEncoderWithBuffer(writer, 64*1024)
// Buffer: 64KB
// Use for: High-throughput scenarios
```

#### 3. Unbuffered (Direct Write)

```go
stream := beve.NewStreamEncoderWithBuffer(writer, 0)
// Buffer: 0 (unbuffered)
// Use for: Low-latency, interactive responses
```

### Auto-Flush Behavior

Stream encoder automatically flushes when:

1. **Buffer full**: When encoded data exceeds buffer size
2. **Close called**: `stream.Close()` flushes remaining data
3. **Explicit flush**: `stream.Flush()` forces write

**Example**:
```go
stream := beve.NewStreamEncoderWithBuffer(writer, 1024) // 1KB buffer

// Encode 100 small structs (~50 bytes each)
for i := 0; i < 100; i++ {
    stream.Encode(User{Name: fmt.Sprintf("User%d", i)})
    // Auto-flushes every ~20 structs (when buffer > 1KB)
}

stream.Close() // Final flush
```

### HTTP Response Example

```go
func handleUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/beve")
    w.Header().Set("Transfer-Encoding", "chunked")
    
    stream := beve.NewStreamEncoder(w)
    defer stream.Close()
    
    // Query database (cursor)
    rows, _ := db.Query("SELECT * FROM users")
    defer rows.Close()
    
    for rows.Next() {
        var user User
        rows.Scan(&user.Name, &user.Age)
        
        // Encode and flush immediately
        if err := stream.Encode(user); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
    }
    
    // Final flush on Close
}
```

**Benefits**:
- **Progressive rendering**: Client receives data as it's generated
- **Low memory**: Server doesn't load all data at once
- **Fast TTFB**: Time to first byte reduced (no batching)

### File Generation Example

```go
func generateReport(filename string, records []Record) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    
    // Buffered stream (64KB)
    stream := beve.NewStreamEncoderWithBuffer(f, 64*1024)
    defer stream.Close()
    
    // Write header
    header := ReportHeader{
        Version:   "1.0",
        Generated: time.Now(),
        Count:     len(records),
    }
    stream.Encode(header)
    
    // Write records
    for _, record := range records {
        if err := stream.Encode(record); err != nil {
            return err
        }
    }
    
    return stream.Flush()
}
```

---

## Stream Decoder

### Basic Usage

```go
// Create stream decoder
stream := beve.NewStreamDecoder(reader)

// Decode multiple values
for {
    var user User
    err := stream.Decode(&user)
    if err == io.EOF {
        break // End of stream
    }
    if err != nil {
        return err
    }
    
    // Process user
    fmt.Println(user.Name)
}
```

### Reading from File

```go
func processFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    
    stream := beve.NewStreamDecoder(f)
    
    // Read header
    var header ReportHeader
    if err := stream.Decode(&header); err != nil {
        return err
    }
    
    fmt.Printf("Report version: %s\n", header.Version)
    
    // Read records
    count := 0
    for {
        var record Record
        err := stream.Decode(&record)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        processRecord(record)
        count++
    }
    
    fmt.Printf("Processed %d records\n", count)
    return nil
}
```

### HTTP Request Example

```go
func handleUpload(w http.ResponseWriter, r *http.Request) {
    stream := beve.NewStreamDecoder(r.Body)
    defer r.Body.Close()
    
    var users []User
    
    for {
        var user User
        err := stream.Decode(&user)
        if err == io.EOF {
            break
        }
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        
        // Validate user
        if err := validateUser(user); err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        
        users = append(users, user)
    }
    
    // Save to database
    db.SaveUsers(users)
    
    json.NewEncoder(w).Encode(map[string]int{
        "count": len(users),
    })
}
```

### Network Stream Example

```go
func receiveData(conn net.Conn) error {
    defer conn.Close()
    
    stream := beve.NewStreamDecoder(conn)
    
    for {
        var msg Message
        err := stream.Decode(&msg)
        if err == io.EOF {
            fmt.Println("Connection closed")
            break
        }
        if err != nil {
            return fmt.Errorf("decode error: %w", err)
        }
        
        // Handle message
        handleMessage(msg)
    }
    
    return nil
}
```

---

## Buffering Strategies

### Strategy 1: Small Buffer (Low Latency)

**Use Case**: Interactive applications, real-time feeds

```go
stream := beve.NewStreamEncoderWithBuffer(writer, 1024) // 1KB
// Flushes every 1KB (~20 small structs)
// Latency: Low (frequent flushes)
// Throughput: Lower (more syscalls)
```

**When to use**:
- WebSocket communication
- Real-time dashboards
- Interactive CLI tools

### Strategy 2: Medium Buffer (Balanced)

**Use Case**: General-purpose streaming

```go
stream := beve.NewStreamEncoder(writer) // 8KB default
// Flushes every 8KB (~160 small structs)
// Latency: Medium
// Throughput: Good
```

**When to use**:
- HTTP APIs
- File generation
- Log processing

### Strategy 3: Large Buffer (High Throughput)

**Use Case**: Batch processing, ETL jobs

```go
stream := beve.NewStreamEncoderWithBuffer(writer, 64*1024) // 64KB
// Flushes every 64KB (~1300 small structs)
// Latency: High (rare flushes)
// Throughput: Maximum (fewer syscalls)
```

**When to use**:
- Data exports
- Backup systems
- Analytics pipelines

### Strategy 4: Unbuffered (Immediate)

**Use Case**: Critical data, no buffering tolerated

```go
stream := beve.NewStreamEncoderWithBuffer(writer, 0) // 0 = unbuffered
// Flushes on every Encode()
// Latency: Minimal
// Throughput: Lowest (maximum syscalls)
```

**When to use**:
- Financial transactions
- Audit logs
- Critical alerts

### Buffer Size Comparison

| Buffer Size | Flushes/1000 | Latency | Throughput | Memory |
|-------------|--------------|---------|------------|--------|
| 0 (unbuffered) | 1000 | Minimal | Low | Minimal |
| 1KB | 50 | Low | Medium | Low |
| 8KB (default) | 6 | Medium | High | Medium |
| 64KB | 1 | High | Maximum | High |

---

## Network Streaming

### TCP Server Example

```go
func startServer(addr string) error {
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    defer ln.Close()
    
    fmt.Printf("Server listening on %s\n", addr)
    
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    // Bidirectional streaming
    decoder := beve.NewStreamDecoder(conn)
    encoder := beve.NewStreamEncoder(conn)
    
    for {
        // Read request
        var req Request
        if err := decoder.Decode(&req); err != nil {
            return
        }
        
        // Process request
        resp := processRequest(req)
        
        // Write response
        if err := encoder.Encode(resp); err != nil {
            return
        }
        encoder.Flush() // Immediate send
    }
}
```

### TCP Client Example

```go
func connectAndStream(addr string, messages []Message) error {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }
    defer conn.Close()
    
    stream := beve.NewStreamEncoder(conn)
    defer stream.Close()
    
    // Send messages
    for _, msg := range messages {
        if err := stream.Encode(msg); err != nil {
            return err
        }
    }
    
    // Read responses
    decoder := beve.NewStreamDecoder(conn)
    
    for range messages {
        var resp Response
        if err := decoder.Decode(&resp); err != nil {
            return err
        }
        fmt.Println("Response:", resp)
    }
    
    return nil
}
```

### WebSocket Streaming

```go
func handleWebSocket(ws *websocket.Conn) {
    defer ws.Close()
    
    // BEVE over WebSocket
    stream := beve.NewStreamEncoder(ws)
    
    // Send real-time updates
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        update := getLatestUpdate()
        if err := stream.Encode(update); err != nil {
            return
        }
        stream.Flush() // Push immediately
    }
}
```

---

## File Streaming

### Large File Processing

**Problem**: 10GB file with 100M records

**❌ Bad: Load all into memory**:
```go
data, _ := os.ReadFile("large.beve") // OOM!
var records []Record
beve.Unmarshal(data, &records) // Crash!
```

**✅ Good: Stream processing**:
```go
func processLargeFile(filename string) error {
    f, _ := os.Open(filename)
    defer f.Close()
    
    stream := beve.NewStreamDecoder(f)
    
    count := 0
    for {
        var record Record
        err := stream.Decode(&record)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        // Process one record at a time
        processRecord(record)
        count++
        
        if count%1000 == 0 {
            fmt.Printf("Processed %d records\n", count)
        }
    }
    
    return nil
}
// Memory: Constant (~1MB)
// Time: Linear (no memory bottleneck)
```

### Parallel File Processing

```go
func parallelProcess(filename string, workers int) error {
    f, _ := os.Open(filename)
    defer f.Close()
    
    stream := beve.NewStreamDecoder(f)
    
    // Worker pool
    jobs := make(chan Record, 100)
    done := make(chan bool)
    
    // Start workers
    for i := 0; i < workers; i++ {
        go func() {
            for record := range jobs {
                processRecord(record)
            }
            done <- true
        }()
    }
    
    // Feed jobs
    go func() {
        for {
            var record Record
            if err := stream.Decode(&record); err != nil {
                break
            }
            jobs <- record
        }
        close(jobs)
    }()
    
    // Wait for completion
    for i := 0; i < workers; i++ {
        <-done
    }
    
    return nil
}
```

### Compressed File Streaming

```go
func readCompressedFile(filename string) error {
    f, _ := os.Open(filename)
    defer f.Close()
    
    // Decompress on-the-fly
    gzReader, _ := gzip.NewReader(f)
    defer gzReader.Close()
    
    // Stream from decompressed data
    stream := beve.NewStreamDecoder(gzReader)
    
    for {
        var record Record
        if err := stream.Decode(&record); err == io.EOF {
            break
        } else if err != nil {
            return err
        }
        
        processRecord(record)
    }
    
    return nil
}
```

---

## Advanced Patterns

### Pattern 1: Streaming with Progress

```go
func streamWithProgress(filename string, total int) error {
    f, _ := os.Create(filename)
    defer f.Close()
    
    stream := beve.NewStreamEncoder(f)
    defer stream.Close()
    
    bar := progressbar.Default(int64(total))
    
    for i := 0; i < total; i++ {
        record := generateRecord(i)
        
        if err := stream.Encode(record); err != nil {
            return err
        }
        
        bar.Add(1)
    }
    
    return nil
}
```

### Pattern 2: Chunked Streaming

```go
func streamInChunks(records []Record, chunkSize int) error {
    stream := beve.NewStreamEncoder(writer)
    defer stream.Close()
    
    for i := 0; i < len(records); i += chunkSize {
        end := i + chunkSize
        if end > len(records) {
            end = len(records)
        }
        
        chunk := records[i:end]
        
        // Encode chunk
        for _, record := range chunk {
            stream.Encode(record)
        }
        
        // Flush after each chunk
        stream.Flush()
    }
    
    return nil
}
```

### Pattern 3: Multiplexed Streaming

```go
func multiplexStreams(sources []io.Reader, dest io.Writer) error {
    encoder := beve.NewStreamEncoder(dest)
    defer encoder.Close()
    
    for i, source := range sources {
        decoder := beve.NewStreamDecoder(source)
        
        for {
            var record Record
            err := decoder.Decode(&record)
            if err == io.EOF {
                break
            }
            if err != nil {
                return err
            }
            
            // Add source identifier
            record.SourceID = i
            
            encoder.Encode(record)
        }
    }
    
    return nil
}
```

### Pattern 4: Streaming with Validation

```go
func streamWithValidation(reader io.Reader) error {
    stream := beve.NewStreamDecoder(reader)
    
    valid := 0
    invalid := 0
    
    for {
        var record Record
        err := stream.Decode(&record)
        if err == io.EOF {
            break
        }
        if err != nil {
            invalid++
            continue // Skip invalid records
        }
        
        // Validate
        if err := validateRecord(record); err != nil {
            invalid++
            log.Printf("Invalid record: %v\n", err)
            continue
        }
        
        processRecord(record)
        valid++
    }
    
    log.Printf("Valid: %d, Invalid: %d\n", valid, invalid)
    return nil
}
```

---

## Performance Optimization

### 1. Buffer Size Tuning

**Benchmark different buffer sizes**:

```go
func BenchmarkBufferSize(b *testing.B) {
    sizes := []int{0, 1024, 8192, 65536}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("buffer_%d", size), func(b *testing.B) {
            buf := &bytes.Buffer{}
            stream := beve.NewStreamEncoderWithBuffer(buf, size)
            
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                stream.Encode(testData)
            }
            stream.Close()
        })
    }
}
```

**Results** (example):
```
buffer_0     10000 ns/op (unbuffered)
buffer_1024   2000 ns/op (5× faster)
buffer_8192   1500 ns/op (6.6× faster)
buffer_65536  1400 ns/op (7× faster)
```

### 2. Encoder Pooling with Streaming

```go
var encoderPool = sync.Pool{
    New: func() interface{} {
        return beve.NewStreamEncoder(nil)
    },
}

func encodeWithPool(writer io.Writer, data []Record) error {
    enc := encoderPool.Get().(*beve.StreamEncoder)
    defer encoderPool.Put(enc)
    
    enc.Reset(writer)
    
    for _, record := range data {
        if err := enc.Encode(record); err != nil {
            return err
        }
    }
    
    return enc.Flush()
}
```

### 3. Zero-Copy Streaming

```go
func zeroСopyStream(records []Record, writer io.Writer) error {
    buf := make([]byte, 0, 64*1024)
    
    for _, record := range records {
        // Encode with zero-copy
        data, _ := beve.MarshalZeroCopy(record, buf[:0])
        
        // Write directly (no intermediate buffer)
        writer.Write(data)
        
        // Reuse buffer
    }
    
    return nil
}
```

### 4. Batch Flushing

```go
func batchFlush(records []Record, batchSize int) error {
    stream := beve.NewStreamEncoder(writer)
    defer stream.Close()
    
    for i, record := range records {
        stream.Encode(record)
        
        // Flush every N records
        if (i+1)%batchSize == 0 {
            stream.Flush()
        }
    }
    
    return nil
}
```

### Performance Comparison

**Test**: Encode 100,000 records

| Method | Time | Memory | Syscalls |
|--------|------|--------|----------|
| Marshal (batch) | 50ms | 150MB | 1 |
| Stream (unbuffered) | 800ms | 1MB | 100,000 |
| Stream (1KB) | 120ms | 1MB | 5,000 |
| Stream (8KB) | 80ms | 8MB | 650 |
| Stream (64KB) | 65ms | 64MB | 80 |

**Optimal**: 8-64KB buffer (balanced latency + throughput)

---

## Summary

### Key Takeaways

1. **Use streaming** for large datasets (>1MB)
2. **Buffer size matters**: 8-64KB for most use cases
3. **Flush strategically**: Balance latency vs throughput
4. **Handle EOF gracefully**: `io.EOF` is not an error
5. **Pool encoders** for repeated operations
6. **Validate progressively**: Don't wait for full load

### Streaming Patterns

| Pattern | Use Case | Buffer Size |
|---------|----------|-------------|
| **Interactive** | Real-time feeds | 0-1KB |
| **API responses** | HTTP/gRPC | 8KB (default) |
| **File processing** | ETL jobs | 64KB |
| **Critical data** | Transactions | 0 (unbuffered) |

### Next Steps

- **[Performance →](performance.md)** - Deep performance tuning
- **[Extensions →](extensions.md)** - Advanced features
- **[Error Handling →](error-handling.md)** - Robust error patterns
- **[API Reference →](../api/core.md)** - Function docs

---

**Ready to optimize?** Check out the [Performance Guide](performance.md) for advanced techniques.
