# Production Deployment Guide

**BEVE-Go Production Deployment Best Practices**

This guide covers deploying BEVE-Go applications to production environments with optimal performance, reliability, and scalability.

---

## Table of Contents

1. [Environment Setup](#environment-setup)
2. [Containerization](#containerization)
3. [Kubernetes Deployment](#kubernetes-deployment)
4. [Performance Tuning](#performance-tuning)
5. [Zero-Downtime Deployment](#zero-downtime-deployment)
6. [Load Balancing](#load-balancing)
7. [Configuration Management](#configuration-management)
8. [Health Checks](#health-checks)

---

## Environment Setup

### Development vs Production Configs

**Development Environment**:

```go
// config/dev.go
package config

type Config struct {
    // Smaller buffers for faster iteration
    BufferPoolSize      int  `env:"BUFFER_POOL_SIZE" envDefault:"100"`
    MaxBufferSize       int  `env:"MAX_BUFFER_SIZE" envDefault:"1048576"` // 1MB
    
    // Verbose logging
    LogLevel            string `env:"LOG_LEVEL" envDefault:"debug"`
    
    // No limits for testing
    MaxMessageSize      int64  `env:"MAX_MESSAGE_SIZE" envDefault:"0"` // unlimited
}
```

**Production Environment**:

```go
// config/prod.go
package config

type Config struct {
    // Large buffer pools for high throughput
    BufferPoolSize      int  `env:"BUFFER_POOL_SIZE" envDefault:"10000"`
    MaxBufferSize       int  `env:"MAX_BUFFER_SIZE" envDefault:"104857600"` // 100MB
    
    // Minimal logging
    LogLevel            string `env:"LOG_LEVEL" envDefault:"warn"`
    
    // Strict limits for security
    MaxMessageSize      int64  `env:"MAX_MESSAGE_SIZE" envDefault:"104857600"` // 100MB
    MaxNestingDepth     int    `env:"MAX_NESTING_DEPTH" envDefault:"16"`
    MaxArraySize        int    `env:"MAX_ARRAY_SIZE" envDefault:"1000000"`
    
    // Resource limits
    MaxConcurrentOps    int    `env:"MAX_CONCURRENT_OPS" envDefault:"10000"`
    ReadTimeout         int    `env:"READ_TIMEOUT_SEC" envDefault:"30"`
    WriteTimeout        int    `env:"WRITE_TIMEOUT_SEC" envDefault:"30"`
}
```

### Environment Variables

**Required**:

```bash
# .env.production
BUFFER_POOL_SIZE=10000
MAX_BUFFER_SIZE=104857600
LOG_LEVEL=warn
MAX_MESSAGE_SIZE=104857600
```

**Optional Performance Tuning**:

```bash
# Go runtime
GOMAXPROCS=0  # Auto-detect (use all CPUs)
GOMEMLIMIT=8GiB  # Soft memory limit for GC

# BEVE specific
BEVE_ZERO_COPY=true  # Enable zero-copy mode
BEVE_USE_ARENA=true  # Enable arena allocator
BEVE_ARENA_SIZE=4096  # Arena chunk size
```

---

## Containerization

### Dockerfile (Multi-Stage Build)

**Optimized for Size & Performance**:

```dockerfile
# Stage 1: Builder
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

WORKDIR /build

# Copy go.mod/go.sum first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with production flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=${VERSION:-dev}" \
    -tags=production \
    -trimpath \
    -o /app/server \
    ./cmd/server

# Stage 2: Runtime
FROM alpine:3.18

# Add ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 app && \
    adduser -D -u 1000 -G app app

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Copy configs (if any)
COPY --from=builder /build/configs ./configs

# Set ownership
RUN chown -R app:app /app

# Switch to non-root
USER app

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run
ENTRYPOINT ["/app/server"]
```

**Build & Push**:

```bash
#!/bin/bash
# build.sh

VERSION=$(git describe --tags --always)
IMAGE="myregistry/beve-app:${VERSION}"

# Build
docker build -t ${IMAGE} \
    --build-arg VERSION=${VERSION} \
    --target=production \
    .

# Tag as latest
docker tag ${IMAGE} myregistry/beve-app:latest

# Push
docker push ${IMAGE}
docker push myregistry/beve-app:latest
```

### Docker Compose (Development)

```yaml
# docker-compose.yml
version: '3.8'

services:
  app:
    build:
      context: .
      target: builder
    ports:
      - "8080:8080"
    environment:
      - BUFFER_POOL_SIZE=100
      - LOG_LEVEL=debug
      - GOMAXPROCS=4
    volumes:
      - ./configs:/app/configs:ro
    networks:
      - beve-net
    restart: unless-stopped
    
  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml:ro
    networks:
      - beve-net

networks:
  beve-net:
    driver: bridge
```

---

## Kubernetes Deployment

### Deployment Manifest

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: beve-app
  labels:
    app: beve-app
    version: v1.3.0
spec:
  replicas: 3
  selector:
    matchLabels:
      app: beve-app
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0  # Zero-downtime
  template:
    metadata:
      labels:
        app: beve-app
        version: v1.3.0
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
        prometheus.io/path: "/metrics"
    spec:
      containers:
      - name: beve-app
        image: myregistry/beve-app:v1.3.0
        imagePullPolicy: IfNotPresent
        
        ports:
        - containerPort: 8080
          name: http
          protocol: TCP
        
        env:
        - name: BUFFER_POOL_SIZE
          value: "10000"
        - name: LOG_LEVEL
          value: "warn"
        - name: GOMAXPROCS
          valueFrom:
            resourceFieldRef:
              resource: limits.cpu
              divisor: "1"
        
        resources:
          requests:
            cpu: "500m"      # 0.5 CPU core
            memory: "512Mi"  # 512 MB
          limits:
            cpu: "2000m"     # 2 CPU cores
            memory: "2Gi"    # 2 GB
        
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 3
          failureThreshold: 3
        
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
          timeoutSeconds: 2
          failureThreshold: 2
        
        lifecycle:
          preStop:
            exec:
              command: ["/bin/sh", "-c", "sleep 15"]  # Grace period
      
      terminationGracePeriodSeconds: 30
      
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - beve-app
              topologyKey: kubernetes.io/hostname
```

### Service

```yaml
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: beve-app
  labels:
    app: beve-app
spec:
  type: ClusterIP
  selector:
    app: beve-app
  ports:
  - port: 80
    targetPort: 8080
    protocol: TCP
    name: http
  sessionAffinity: None
```

### HorizontalPodAutoscaler

```yaml
# hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: beve-app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: beve-app
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 300  # 5 min cooldown
      policies:
      - type: Percent
        value: 50
        periodSeconds: 60
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Percent
        value: 100
        periodSeconds: 30
      - type: Pods
        value: 2
        periodSeconds: 30
      selectPolicy: Max
```

### ConfigMap

```yaml
# configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: beve-config
data:
  config.yaml: |
    buffer:
      pool_size: 10000
      max_size: 104857600
    
    limits:
      max_message_size: 104857600
      max_nesting_depth: 16
      max_array_size: 1000000
    
    performance:
      zero_copy: true
      use_arena: true
      arena_size: 4096
    
    logging:
      level: warn
      format: json
```

---

## Performance Tuning

### GOMAXPROCS Optimization

**Auto-detect CPU cores**:

```go
import "runtime"

func init() {
    // Use all available CPUs
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // Or set based on container limits (K8s)
    // GOMAXPROCS env var handled automatically
}
```

**Container CPU limits**:

```yaml
# In Kubernetes, GOMAXPROCS = limits.cpu
env:
- name: GOMAXPROCS
  valueFrom:
    resourceFieldRef:
      resource: limits.cpu
      divisor: "1"  # Round down
```

### Buffer Pool Sizing

**Rule of Thumb**: `pool_size = max_concurrent_requests × 1.5`

```go
// For 10,000 concurrent requests
poolSize := 15000

beve.SetBufferPoolSize(poolSize)
```

**Dynamic Sizing**:

```go
import "runtime"

func calculatePoolSize() int {
    numCPU := runtime.NumCPU()
    
    // Estimate concurrent requests per CPU
    requestsPerCPU := 1000
    
    return numCPU * requestsPerCPU * 2  // 2× headroom
}
```

### Memory Ballast

**Reduce GC frequency** with memory ballast:

```go
package main

import "runtime"

// Allocate 1GB ballast (not used, just reserves heap space)
var ballast = make([]byte, 1<<30)

func main() {
    // GC will trigger less frequently
    runtime.GC()  // Initial GC
    
    // Your application logic
    startServer()
}
```

**Trade-off**: Higher RSS, but more stable GC pauses.

### GC Tuning

**Set soft memory limit** (Go 1.19+):

```bash
export GOMEMLIMIT=8GiB  # Soft limit
```

**Result**: GC triggers before hitting 8GB, preventing OOM.

---

## Zero-Downtime Deployment

### Rolling Update Strategy

**Kubernetes**:

```yaml
strategy:
  type: RollingUpdate
  rollingUpdate:
    maxSurge: 1         # Add 1 new pod before removing old
    maxUnavailable: 0   # Never go below desired replicas
```

**Steps**:
1. New pod starts
2. Readiness probe passes
3. Traffic shifts to new pod
4. Old pod receives SIGTERM
5. Grace period (30s)
6. Old pod stops

### Graceful Shutdown

**Server Implementation**:

```go
package main

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: handler,
    }
    
    // Start server
    go func() {
        if err := server.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
    
    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down gracefully...")
    
    // Grace period (30s)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Stop accepting new requests
    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Shutdown error: %v", err)
    }
    
    // Flush BEVE buffers
    beve.FlushAll()
    
    log.Println("Server stopped")
}
```

### PreStop Hook

**Kubernetes**:

```yaml
lifecycle:
  preStop:
    exec:
      command: ["/bin/sh", "-c", "sleep 15"]
```

**Why?**: Delay pod termination to allow load balancer to remove endpoint.

---

## Load Balancing

### NGINX Configuration

```nginx
# nginx.conf
upstream beve_backend {
    least_conn;  # Route to least busy server
    
    server beve-app-1:8080 max_fails=3 fail_timeout=30s;
    server beve-app-2:8080 max_fails=3 fail_timeout=30s;
    server beve-app-3:8080 max_fails=3 fail_timeout=30s;
    
    keepalive 32;  # Connection pooling
}

server {
    listen 80;
    
    location / {
        proxy_pass http://beve_backend;
        proxy_http_version 1.1;
        
        # Timeouts
        proxy_connect_timeout 10s;
        proxy_send_timeout 30s;
        proxy_read_timeout 30s;
        
        # Headers
        proxy_set_header Connection "";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        
        # Buffer settings (important for large BEVE payloads)
        proxy_buffering on;
        proxy_buffer_size 128k;
        proxy_buffers 8 128k;
        proxy_busy_buffers_size 256k;
    }
    
    location /health {
        proxy_pass http://beve_backend/health;
        access_log off;  # Don't log health checks
    }
}
```

### AWS Application Load Balancer

**Target Group Settings**:

```json
{
  "Protocol": "HTTP",
  "Port": 8080,
  "HealthCheckProtocol": "HTTP",
  "HealthCheckPath": "/health",
  "HealthCheckIntervalSeconds": 30,
  "HealthCheckTimeoutSeconds": 5,
  "HealthyThresholdCount": 2,
  "UnhealthyThresholdCount": 3,
  "Matcher": {
    "HttpCode": "200"
  },
  "TargetGroupAttributes": [
    {
      "Key": "deregistration_delay.timeout_seconds",
      "Value": "30"
    },
    {
      "Key": "stickiness.enabled",
      "Value": "false"
    }
  ]
}
```

---

## Configuration Management

### Hierarchical Config

```go
// config/config.go
package config

import (
    "os"
    "strconv"
)

type Config struct {
    Buffer    BufferConfig
    Limits    LimitsConfig
    Server    ServerConfig
}

type BufferConfig struct {
    PoolSize  int
    MaxSize   int
    ZeroCopy  bool
    UseArena  bool
    ArenaSize int
}

func Load() (*Config, error) {
    return &Config{
        Buffer: BufferConfig{
            PoolSize:  getEnvInt("BUFFER_POOL_SIZE", 10000),
            MaxSize:   getEnvInt("MAX_BUFFER_SIZE", 100*1024*1024),
            ZeroCopy:  getEnvBool("BEVE_ZERO_COPY", true),
            UseArena:  getEnvBool("BEVE_USE_ARENA", true),
            ArenaSize: getEnvInt("BEVE_ARENA_SIZE", 4096),
        },
        Limits: LimitsConfig{
            MaxMessageSize:  getEnvInt64("MAX_MESSAGE_SIZE", 100*1024*1024),
            MaxNestingDepth: getEnvInt("MAX_NESTING_DEPTH", 16),
            MaxArraySize:    getEnvInt("MAX_ARRAY_SIZE", 1_000_000),
        },
        Server: ServerConfig{
            Port:         getEnvInt("PORT", 8080),
            ReadTimeout:  getEnvInt("READ_TIMEOUT_SEC", 30),
            WriteTimeout: getEnvInt("WRITE_TIMEOUT_SEC", 30),
        },
    }, nil
}

func getEnvInt(key string, def int) int {
    if val := os.Getenv(key); val != "" {
        if i, err := strconv.Atoi(val); err == nil {
            return i
        }
    }
    return def
}

func getEnvBool(key string, def bool) bool {
    if val := os.Getenv(key); val != "" {
        if b, err := strconv.ParseBool(val); err == nil {
            return b
        }
    }
    return def
}
```

### Hot Reload

```go
// config/watcher.go
package config

import (
    "log"
    "sync/atomic"
    "time"
)

type Watcher struct {
    current atomic.Value  // *Config
    stop    chan struct{}
}

func NewWatcher() *Watcher {
    w := &Watcher{
        stop: make(chan struct{}),
    }
    
    // Load initial config
    cfg, _ := Load()
    w.current.Store(cfg)
    
    // Watch for changes
    go w.watch()
    
    return w
}

func (w *Watcher) watch() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            cfg, err := Load()
            if err != nil {
                log.Printf("Config reload error: %v", err)
                continue
            }
            
            old := w.Get()
            if configChanged(old, cfg) {
                w.current.Store(cfg)
                log.Println("Config reloaded")
                
                // Apply changes
                applyConfig(cfg)
            }
        case <-w.stop:
            return
        }
    }
}

func (w *Watcher) Get() *Config {
    return w.current.Load().(*Config)
}

func applyConfig(cfg *Config) {
    // Update BEVE settings
    beve.SetBufferPoolSize(cfg.Buffer.PoolSize)
    // ... other settings
}
```

---

## Health Checks

### Health Endpoint

```go
// handlers/health.go
package handlers

import (
    "encoding/json"
    "net/http"
    "runtime"
    "sync/atomic"
    "time"
)

var (
    startTime = time.Now()
    reqCount  int64
)

type HealthResponse struct {
    Status    string            `json:"status"`
    Uptime    string            `json:"uptime"`
    Version   string            `json:"version"`
    Requests  int64             `json:"requests"`
    GoVersion string            `json:"go_version"`
    Memory    MemoryStats       `json:"memory"`
    Checks    map[string]string `json:"checks"`
}

type MemoryStats struct {
    Alloc      uint64 `json:"alloc"`
    TotalAlloc uint64 `json:"total_alloc"`
    Sys        uint64 `json:"sys"`
    NumGC      uint32 `json:"num_gc"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    resp := HealthResponse{
        Status:    "ok",
        Uptime:    time.Since(startTime).String(),
        Version:   "v1.3.0",
        Requests:  atomic.LoadInt64(&reqCount),
        GoVersion: runtime.Version(),
        Memory: MemoryStats{
            Alloc:      m.Alloc / 1024 / 1024,      // MB
            TotalAlloc: m.TotalAlloc / 1024 / 1024, // MB
            Sys:        m.Sys / 1024 / 1024,        // MB
            NumGC:      m.NumGC,
        },
        Checks: runHealthChecks(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(resp)
}

func runHealthChecks() map[string]string {
    checks := make(map[string]string)
    
    // Check buffer pool
    if beve.GetPoolStats().AvailableBuffers > 0 {
        checks["buffer_pool"] = "ok"
    } else {
        checks["buffer_pool"] = "exhausted"
    }
    
    // Check memory
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    if m.Alloc < 2*1024*1024*1024 {  // < 2GB
        checks["memory"] = "ok"
    } else {
        checks["memory"] = "high"
    }
    
    return checks
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
    // More strict checks for readiness
    checks := runHealthChecks()
    
    for _, status := range checks {
        if status != "ok" {
            w.WriteHeader(http.StatusServiceUnavailable)
            json.NewEncoder(w).Encode(map[string]string{
                "status": "not_ready",
                "checks": status,
            })
            return
        }
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
}
```

---

## Summary

### Deployment Checklist

- [ ] **Containerization**: Multi-stage Dockerfile, non-root user
- [ ] **Resource Limits**: CPU/memory requests & limits set
- [ ] **Health Checks**: Liveness & readiness probes configured
- [ ] **Environment Variables**: All configs externalized
- [ ] **Graceful Shutdown**: SIGTERM handler implemented
- [ ] **Buffer Pool Sizing**: Based on expected load
- [ ] **GOMAXPROCS**: Set to container CPU limits
- [ ] **Load Balancing**: Connection pooling enabled
- [ ] **Monitoring**: Prometheus metrics exposed
- [ ] **Logging**: Structured JSON logging
- [ ] **Autoscaling**: HPA configured with proper metrics
- [ ] **Zero-Downtime**: Rolling update strategy validated

### Performance Targets (Production)

| Metric | Target | Actual (Neoverse-N2) |
|--------|--------|----------------------|
| **Marshal Small** | < 1 μs | 694 ns ✅ |
| **Unmarshal Small** | < 2 μs | 805 ns ✅ |
| **Marshal Large** | < 150 μs | 103 μs ✅ |
| **Unmarshal Large** | < 300 μs | 230 μs ✅ |
| **Throughput** | > 10K ops/sec | ~50K ops/sec ✅ |
| **p99 Latency** | < 10 ms | ~2 ms ✅ |
| **Memory/Op** | < 1 KB | 600 bytes ✅ |

---

**Next**: [Monitoring Guide](monitoring.md) · [Security Guide](security.md) · [Troubleshooting](troubleshooting.md)
