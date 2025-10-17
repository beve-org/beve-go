# Production Monitoring Guide

**BEVE-Go Observability & Monitoring Best Practices**

Comprehensive guide to monitoring BEVE-Go applications in production with Prometheus, Grafana, OpenTelemetry, and structured logging.

---

## Table of Contents

1. [Key Metrics](#key-metrics)
2. [Prometheus Integration](#prometheus-integration)
3. [Grafana Dashboards](#grafana-dashboards)
4. [Distributed Tracing](#distributed-tracing)
5. [Structured Logging](#structured-logging)
6. [Alerting Rules](#alerting-rules)
7. [Performance Monitoring](#performance-monitoring)

---

## Key Metrics

### Core BEVE Metrics

**Throughput**:
- `beve_marshal_total` - Total marshal operations
- `beve_unmarshal_total` - Total unmarshal operations
- `beve_operations_per_second` - Current ops/sec

**Latency**:
- `beve_marshal_duration_seconds` - Marshal latency histogram
- `beve_unmarshal_duration_seconds` - Unmarshal latency histogram

**Memory**:
- `beve_buffer_pool_size` - Total buffer pool size
- `beve_buffer_pool_hit_rate` - Pool hit/miss ratio
- `beve_allocated_bytes` - Current allocated memory

**Errors**:
- `beve_errors_total` - Total errors (labeled by type)
- `beve_decode_errors_total` - Decode-specific errors

---

## Prometheus Integration

### Metrics Implementation

```go
// metrics/metrics.go
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // Counters
    MarshalTotal = promauto.NewCounter(prometheus.CounterOpts{
        Name: "beve_marshal_total",
        Help: "Total number of marshal operations",
    })
    
    UnmarshalTotal = promauto.NewCounter(prometheus.CounterOpts{
        Name: "beve_unmarshal_total",
        Help: "Total number of unmarshal operations",
    })
    
    ErrorsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "beve_errors_total",
        Help: "Total errors by type",
    }, []string{"type"})
    
    // Histograms
    MarshalDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name:    "beve_marshal_duration_seconds",
        Help:    "Marshal operation duration",
        Buckets: []float64{0.000001, 0.000005, 0.00001, 0.00005, 0.0001, 0.0005, 0.001, 0.005, 0.01},
    })
    
    UnmarshalDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name:    "beve_unmarshal_duration_seconds",
        Help:    "Unmarshal operation duration",
        Buckets: []float64{0.000001, 0.000005, 0.00001, 0.00005, 0.0001, 0.0005, 0.001, 0.005, 0.01},
    })
    
    MessageSize = promauto.NewHistogram(prometheus.HistogramOpts{
        Name:    "beve_message_size_bytes",
        Help:    "BEVE message size distribution",
        Buckets: prometheus.ExponentialBuckets(100, 10, 8), // 100B to 10MB
    })
    
    // Gauges
    BufferPoolSize = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "beve_buffer_pool_size",
        Help: "Current buffer pool size",
    })
    
    BufferPoolHitRate = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "beve_buffer_pool_hit_rate",
        Help: "Buffer pool hit rate (0-1)",
    })
    
    AllocatedBytes = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "beve_allocated_bytes",
        Help: "Current allocated memory in bytes",
    })
    
    GoroutineCount = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "beve_goroutines",
        Help: "Number of goroutines",
    })
)
```

### Instrumentation

```go
// middleware/metrics.go
package middleware

import (
    "time"
    "your-app/metrics"
)

func InstrumentedMarshal(data interface{}) ([]byte, error) {
    start := time.Now()
    
    result, err := beve.Marshal(data)
    
    // Record duration
    metrics.MarshalDuration.Observe(time.Since(start).Seconds())
    metrics.MarshalTotal.Inc()
    
    if err != nil {
        metrics.ErrorsTotal.WithLabelValues("marshal").Inc()
        return nil, err
    }
    
    // Record size
    metrics.MessageSize.Observe(float64(len(result)))
    
    return result, nil
}

func InstrumentedUnmarshal(data []byte, v interface{}) error {
    start := time.Now()
    
    err := beve.Unmarshal(data, v)
    
    // Record duration
    metrics.UnmarshalDuration.Observe(time.Since(start).Seconds())
    metrics.UnmarshalTotal.Inc()
    
    if err != nil {
        metrics.ErrorsTotal.WithLabelValues("unmarshal").Inc()
        return err
    }
    
    return nil
}
```

### Buffer Pool Monitoring

```go
// monitoring/pool.go
package monitoring

import (
    "runtime"
    "time"
    "your-app/metrics"
)

func MonitorBufferPool(interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    
    for range ticker.C {
        stats := beve.GetPoolStats()
        
        // Update metrics
        metrics.BufferPoolSize.Set(float64(stats.TotalBuffers))
        
        hitRate := float64(stats.Hits) / float64(stats.Hits + stats.Misses)
        metrics.BufferPoolHitRate.Set(hitRate)
        
        // Memory stats
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        metrics.AllocatedBytes.Set(float64(m.Alloc))
        metrics.GoroutineCount.Set(float64(runtime.NumGoroutine()))
    }
}

func init() {
    go MonitorBufferPool(5 * time.Second)
}
```

### Prometheus Configuration

```yaml
# prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'beve-app'
    static_configs:
      - targets: ['beve-app:8080']
    metrics_path: /metrics
    scrape_timeout: 10s
```

---

## Grafana Dashboards

### Dashboard JSON

```json
{
  "dashboard": {
    "title": "BEVE Performance Dashboard",
    "panels": [
      {
        "title": "Operations Per Second",
        "targets": [
          {
            "expr": "rate(beve_marshal_total[1m])",
            "legendFormat": "Marshal"
          },
          {
            "expr": "rate(beve_unmarshal_total[1m])",
            "legendFormat": "Unmarshal"
          }
        ],
        "type": "graph"
      },
      {
        "title": "Latency (p50, p95, p99)",
        "targets": [
          {
            "expr": "histogram_quantile(0.50, rate(beve_marshal_duration_seconds_bucket[1m]))",
            "legendFormat": "p50"
          },
          {
            "expr": "histogram_quantile(0.95, rate(beve_marshal_duration_seconds_bucket[1m]))",
            "legendFormat": "p95"
          },
          {
            "expr": "histogram_quantile(0.99, rate(beve_marshal_duration_seconds_bucket[1m]))",
            "legendFormat": "p99"
          }
        ],
        "type": "graph"
      },
      {
        "title": "Buffer Pool Hit Rate",
        "targets": [
          {
            "expr": "beve_buffer_pool_hit_rate",
            "legendFormat": "Hit Rate"
          }
        ],
        "type": "gauge",
        "thresholds": [
          {"value": 0.8, "color": "red"},
          {"value": 0.9, "color": "yellow"},
          {"value": 0.95, "color": "green"}
        ]
      },
      {
        "title": "Memory Usage",
        "targets": [
          {
            "expr": "beve_allocated_bytes / 1024 / 1024",
            "legendFormat": "Allocated (MB)"
          }
        ],
        "type": "graph"
      },
      {
        "title": "Error Rate",
        "targets": [
          {
            "expr": "rate(beve_errors_total[1m])",
            "legendFormat": "{{type}}"
          }
        ],
        "type": "graph"
      }
    ]
  }
}
```

### Key Panels

**1. Throughput Panel**:
```promql
# Requests per second
rate(beve_marshal_total[1m]) + rate(beve_unmarshal_total[1m])
```

**2. Latency Panel** (percentiles):
```promql
# p50
histogram_quantile(0.50, rate(beve_marshal_duration_seconds_bucket[1m]))

# p95
histogram_quantile(0.95, rate(beve_marshal_duration_seconds_bucket[1m]))

# p99
histogram_quantile(0.99, rate(beve_marshal_duration_seconds_bucket[1m]))
```

**3. Error Rate Panel**:
```promql
# Errors per second by type
rate(beve_errors_total[1m])
```

---

## Distributed Tracing

### OpenTelemetry Integration

```go
// tracing/tracing.go
package tracing

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("beve")

func TracedMarshal(ctx context.Context, data interface{}) ([]byte, error) {
    ctx, span := tracer.Start(ctx, "beve.marshal",
        trace.WithSpanKind(trace.SpanKindInternal),
    )
    defer span.End()
    
    result, err := beve.Marshal(data)
    
    if err != nil {
        span.RecordError(err)
        span.SetAttributes(attribute.String("error", err.Error()))
        return nil, err
    }
    
    span.SetAttributes(
        attribute.Int("message.size", len(result)),
        attribute.String("data.type", fmt.Sprintf("%T", data)),
    )
    
    return result, nil
}

func TracedUnmarshal(ctx context.Context, data []byte, v interface{}) error {
    ctx, span := tracer.Start(ctx, "beve.unmarshal",
        trace.WithSpanKind(trace.SpanKindInternal),
    )
    defer span.End()
    
    span.SetAttributes(
        attribute.Int("input.size", len(data)),
    )
    
    err := beve.Unmarshal(data, v)
    
    if err != nil {
        span.RecordError(err)
        span.SetAttributes(attribute.String("error", err.Error()))
    }
    
    return err
}
```

### Jaeger Configuration

```go
// tracing/jaeger.go
package tracing

import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitJaeger(serviceName string) (*sdktrace.TracerProvider, error) {
    exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
        jaeger.WithEndpoint("http://jaeger:14268/api/traces"),
    ))
    if err != nil {
        return nil, err
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
        sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.1)), // 10% sampling
    )
    
    otel.SetTracerProvider(tp)
    
    return tp, nil
}
```

---

## Structured Logging

### Logger Implementation

```go
// logging/logger.go
package logging

import (
    "encoding/json"
    "log"
    "os"
    "time"
)

type Level string

const (
    DEBUG Level = "debug"
    INFO  Level = "info"
    WARN  Level = "warn"
    ERROR Level = "error"
)

type Logger struct {
    level  Level
    logger *log.Logger
}

type LogEntry struct {
    Timestamp string                 `json:"timestamp"`
    Level     Level                  `json:"level"`
    Message   string                 `json:"message"`
    Fields    map[string]interface{} `json:"fields,omitempty"`
}

func New(level Level) *Logger {
    return &Logger{
        level:  level,
        logger: log.New(os.Stdout, "", 0),
    }
}

func (l *Logger) log(level Level, msg string, fields map[string]interface{}) {
    entry := LogEntry{
        Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
        Level:     level,
        Message:   msg,
        Fields:    fields,
    }
    
    data, _ := json.Marshal(entry)
    l.logger.Println(string(data))
}

func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
    var f map[string]interface{}
    if len(fields) > 0 {
        f = fields[0]
    }
    l.log(INFO, msg, f)
}

func (l *Logger) Error(msg string, err error, fields ...map[string]interface{}) {
    f := map[string]interface{}{"error": err.Error()}
    if len(fields) > 0 {
        for k, v := range fields[0] {
            f[k] = v
        }
    }
    l.log(ERROR, msg, f)
}
```

### BEVE Operation Logging

```go
// logging/beve.go
package logging

import "time"

func LogMarshal(start time.Time, size int, err error) {
    duration := time.Since(start)
    
    if err != nil {
        logger.Error("Marshal failed", err, map[string]interface{}{
            "duration_ns": duration.Nanoseconds(),
        })
        return
    }
    
    logger.Info("Marshal complete", map[string]interface{}{
        "duration_ns": duration.Nanoseconds(),
        "size_bytes":  size,
    })
}

func LogUnmarshal(start time.Time, inputSize int, err error) {
    duration := time.Since(start)
    
    if err != nil {
        logger.Error("Unmarshal failed", err, map[string]interface{}{
            "duration_ns":   duration.Nanoseconds(),
            "input_size":    inputSize,
        })
        return
    }
    
    logger.Info("Unmarshal complete", map[string]interface{}{
        "duration_ns":   duration.Nanoseconds(),
        "input_size":    inputSize,
    })
}
```

### Example Log Output

```json
{
  "timestamp": "2025-10-17T14:30:45.123456789Z",
  "level": "info",
  "message": "Marshal complete",
  "fields": {
    "duration_ns": 694000,
    "size_bytes": 1024
  }
}

{
  "timestamp": "2025-10-17T14:30:45.456789012Z",
  "level": "error",
  "message": "Unmarshal failed",
  "fields": {
    "error": "invalid BEVE header",
    "duration_ns": 12000,
    "input_size": 2048
  }
}
```

---

## Alerting Rules

### Prometheus Alerts

```yaml
# alerts.yml
groups:
  - name: beve_alerts
    interval: 30s
    rules:
      # High Latency Alert
      - alert: BeveHighLatency
        expr: histogram_quantile(0.99, rate(beve_marshal_duration_seconds_bucket[5m])) > 0.01
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "BEVE p99 latency above 10ms"
          description: "p99 latency is {{ $value }}s (threshold: 0.01s)"
      
      # High Error Rate
      - alert: BeveHighErrorRate
        expr: rate(beve_errors_total[5m]) > 10
        for: 2m
        labels:
          severity: critical
        annotations:
          summary: "High BEVE error rate"
          description: "Error rate is {{ $value }} errors/sec"
      
      # Low Buffer Pool Hit Rate
      - alert: BeveLowPoolHitRate
        expr: beve_buffer_pool_hit_rate < 0.8
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "Low buffer pool hit rate"
          description: "Hit rate is {{ $value }} (threshold: 0.8)"
      
      # High Memory Usage
      - alert: BeveHighMemory
        expr: beve_allocated_bytes > 2147483648  # 2GB
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High memory usage"
          description: "Allocated {{ $value | humanize }}B"
      
      # Service Down
      - alert: BeveServiceDown
        expr: up{job="beve-app"} == 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "BEVE service is down"
          description: "Instance {{ $labels.instance }} is down"
```

### Alert Manager Configuration

```yaml
# alertmanager.yml
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname', 'cluster']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 12h
  receiver: 'default'
  
  routes:
    - match:
        severity: critical
      receiver: pagerduty
      continue: true
    
    - match:
        severity: warning
      receiver: slack

receivers:
  - name: 'default'
    email_configs:
      - to: 'team@example.com'
  
  - name: 'slack'
    slack_configs:
      - api_url: 'https://hooks.slack.com/services/XXX'
        channel: '#alerts'
        text: '{{ range .Alerts }}{{ .Annotations.summary }}\n{{ end }}'
  
  - name: 'pagerduty'
    pagerduty_configs:
      - service_key: 'YOUR_PAGERDUTY_KEY'
```

---

## Performance Monitoring

### Continuous Profiling

```go
// profiling/continuous.go
package profiling

import (
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

func StartContinuousProfiling() {
    // Enable mutex profiling
    runtime.SetMutexProfileFraction(5)
    
    // Enable block profiling
    runtime.SetBlockProfileRate(5)
    
    // Start pprof server (internal only!)
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Periodic heap snapshots
    go func() {
        ticker := time.NewTicker(5 * time.Minute)
        defer ticker.Stop()
        
        for range ticker.C {
            takeHeapSnapshot()
        }
    }()
}

func takeHeapSnapshot() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    // Log if memory is high
    if m.Alloc > 2*1024*1024*1024 {  // > 2GB
        log.Printf("High memory: %d MB", m.Alloc/1024/1024)
        
        // Could trigger heap dump here
        // pprof.WriteHeapProfile(f)
    }
}
```

### Custom Metrics

```go
// metrics/custom.go
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // Extension-specific metrics
    ExtensionUsage = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "beve_extension_usage_total",
        Help: "Extension usage count by type",
    }, []string{"extension"})
    
    // Zero-copy mode
    ZeroCopyOps = promauto.NewCounter(prometheus.CounterOpts{
        Name: "beve_zero_copy_ops_total",
        Help: "Zero-copy operations count",
    })
    
    // Arena allocator
    ArenaAllocations = promauto.NewCounter(prometheus.CounterOpts{
        Name: "beve_arena_allocations_total",
        Help: "Arena allocator usage",
    })
)

func RecordExtensionUse(extID int) {
    ExtensionUsage.WithLabelValues(fmt.Sprintf("ext_%d", extID)).Inc()
}
```

---

## Summary

### Monitoring Stack

**Components**:
- ✅ Prometheus (metrics collection)
- ✅ Grafana (visualization)
- ✅ Jaeger (distributed tracing)
- ✅ AlertManager (alerting)
- ✅ Structured logging (JSON)

### Key Metrics to Track

| Metric | Target | Alert Threshold |
|--------|--------|-----------------|
| **p99 Latency** | < 2 ms | > 10 ms |
| **Throughput** | > 10K ops/sec | < 1K ops/sec |
| **Error Rate** | < 0.1% | > 1% |
| **Pool Hit Rate** | > 95% | < 80% |
| **Memory Usage** | < 1 GB | > 2 GB |

### Dashboard URLs

- **Grafana**: `http://grafana:3000/d/beve-dashboard`
- **Prometheus**: `http://prometheus:9090`
- **Jaeger**: `http://jaeger:16686`
- **pprof**: `http://localhost:6060/debug/pprof`

---

**Next**: [Security Guide](security.md) · [Troubleshooting](troubleshooting.md) · [Deployment](deployment.md)
