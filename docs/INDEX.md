# 📚 BEVE-Go Documentation Index

**Welcome to BEVE-Go!** This is your central hub for all documentation.

> **BEVE** (Binary Efficient Versatile Encoding) - High-performance binary serialization for Go, designed to be faster than JSON, MessagePack, and CBOR while maintaining JSON compatibility.

---

## � Quick Navigation

### For New Users
- 📦 [Installation](getting-started/installation.md) - Install BEVE-Go *(Coming Soon)*
- ⚡ [5-Minute Quick Start](getting-started/quick-start.md) - First encode/decode *(Coming Soon)*
- 📖 [Basic Usage](getting-started/basic-usage.md) - Common patterns *(Coming Soon)*
- 🔄 [Migrating from JSON](getting-started/json-migration.md) - Switch from `encoding/json` *(Coming Soon)*

### For API Consumers
- 🔧 [Encoding & Decoding](guides/encoding-decoding.md) - Marshal/Unmarshal APIs *(Coming Soon)*
- 🏷️ [Struct Tags](guides/struct-tags.md) - Tag system guide *(Coming Soon)*
- ⚡ [Performance Tuning](guides/performance.md) - Optimize your code *(Coming Soon)*
- 🧩 [Extensions](guides/extensions.md) - Extension system overview *(Coming Soon)*
- ⚠️ [Error Handling](guides/error-handling.md) - Error handling patterns *(Coming Soon)*

### For Performance Engineers
- 📊 [Benchmarks](performance/benchmarks.md) - Latest benchmark results *(Coming Soon)*
- 🚀 [Optimizations](performance/optimizations.md) - All optimizations explained *(Coming Soon)*
- ⚖️ [Comparison](performance/comparison.md) - BEVE vs JSON/CBOR/MessagePack *(Coming Soon)*
- 🔍 [Profiling Guide](performance/profiling.md) - Profile your code *(Coming Soon)*

### For Production Users
- 🚢 [Deployment Guide](production/deployment.md) - Docker/K8s setup *(Coming Soon)*
- 📡 [Monitoring](production/monitoring.md) - Metrics & observability *(Coming Soon)*
- 🔧 [Troubleshooting](production/troubleshooting.md) - Common issues *(Coming Soon)*
- ✅ [Best Practices](production/best-practices.md) - Production patterns *(Coming Soon)*

### For Contributors
- 💻 [Development Setup](contributing/development-setup.md) - Dev environment *(Coming Soon)*
- 🧪 [Testing Guidelines](contributing/testing.md) - How to test *(Coming Soon)*
- 📊 [Benchmarking Standards](contributing/benchmarking.md) - Write benchmarks *(Coming Soon)*
- 👀 [Code Review Checklist](contributing/code-review.md) - Review process *(Coming Soon)*
- 🚀 [Release Process](contributing/release-process.md) - Release workflow *(Coming Soon)*

---

## 📖 Documentation Sections

### 🎯 Getting Started *(Phase 2 - Coming Soon)*

Perfect for new users who want to get up and running quickly.

| Guide | Description | Time |
|-------|-------------|------|
| Installation | Install BEVE-Go and verify setup | 2 min |
| Quick Start | First encode/decode tutorial | 5 min |
| Basic Usage | Common usage patterns | 10 min |
| JSON Migration | Migrate from `encoding/json` | 15 min |

**Total Onboarding Time**: ~30 minutes

---

### 📚 User Guides *(Phase 3 - Coming Soon)*

Comprehensive guides for all major features.

| Guide | Description |
|-------|-------------|
| Encoding & Decoding | Marshal, Unmarshal, streaming, custom types |
| Struct Tags | Tag syntax, options, nested structs |
| Streaming | Stream encoding/decoding for large data |
| Performance | Zero-copy, buffer pooling, SIMD, profiling |
| Extensions | Extension system, MarshalAuto, typed arrays |
| Arena Allocator | Arena usage, pooling, performance |
| Error Handling | Error types, validation, recovery |

---

### 🏛️ Architecture *(Phase 4 - Coming Soon)*

Deep technical documentation for understanding BEVE internals.

| Document | Description |
|----------|-------------|
| Overview | System architecture, component diagram |
| Encoder Design | Encoder internals, buffer management |
| Decoder Design | Decoder internals, reflection caching |
| Buffer Pooling | sync.Pool, arena allocator, memory reuse |
| Extension System | Extension architecture, binary format |
| SIMD Optimizations | SIMD instructions, platform detection |

---

### 📊 Performance *(Phase 5 - Coming Soon)*

All performance-related documentation consolidated in one place.

| Document | Description |
|----------|-------------|
| Benchmarks | Latest results across all platforms |
| Optimizations | All optimizations chronologically |
| Comparison | BEVE vs JSON/CBOR/MessagePack/Protobuf |
| Profiling | CPU/memory profiling, interpreting results |

**Key Metrics** (Apple M2 Max):
- Small unmarshal: **11.7× faster than JSON**
- Medium marshal: **4.0× faster than JSON**
- Large unmarshal: **9.4× faster than JSON**
- Zero-copy mode: **0 allocations, 277ns**

---

### � Production *(Phase 6 - Coming Soon)*

Production deployment and operations guides.

| Guide | Description |
|-------|-------------|
| Deployment | Docker, Kubernetes, load balancing |
| Monitoring | Prometheus, Grafana, alerting |
| Troubleshooting | Common issues, debug checklist |
| Best Practices | Configuration, capacity planning, security |

---

### 📚 API Reference *(Phase 7 - Coming Soon)*

Complete API documentation with examples.

| Package | Description |
|---------|-------------|
| core | Core encoding/decoding APIs |
| translator | JSON ↔ BEVE translation |
| wasm | WebAssembly bindings |
| bevegen | Code generator tool |

Also see: [GoDoc](https://godoc.org/github.com/beve-org/beve-go) for full API reference

---

### 🧩 Extensions *(Phase 8 - Coming Soon)*

Detailed documentation for all 8 BEVE extensions.

| Extension | Feature | Performance |
|-----------|---------|-------------|
| Extension 0 | Field Index | O(1) field access, 77ns |
| Extensions 1-3 | Typed Arrays | 25-48% size reduction |
| Extensions 4-6 | Timestamps/Duration/Interval | Nanosecond precision, 14-29 bytes |
| Extension 8 | UUID | 50% smaller, 0.3ns marshal |
| Extension 9 | RegExp | 173× faster with cache |

**[Extension Overview →](../EXTENSIONS_README.md)**

---

## 📋 Current Documentation (Available Now)

### Core Documentation

- **[SPECIFICATION.md](../SPECIFICATION.md)** - BEVE v1.0 Binary Format Specification
- **[README.md](../README.md)** - Main project README with quickstart and examples
- **[EXTENSIONS_README.md](../EXTENSIONS_README.md)** - Extension system documentation
- **[CHANGELOG.md](../CHANGELOG.md)** - Version history and updates

### Technical Reports (Archived)

Historical technical reports moved to [archive/](../archive/README.md):
- Performance optimization reports
- Test coverage improvements
- Arena allocator analysis

### Benchmarks

- **[Multi-Platform Benchmarks](../benchmarks/MULTI_PLATFORM.md)** - Performance comparison across platforms

### Proposals (Legacy)

- **[Extension Proposal](EXTENSION_PROPOSAL.md)** - Essential data types proposal
- **[Field Index Mode](proposals/FIELD_INDEX.md)** - Field index performance proposal

---

## 🎯 Current Focus Areas

### 1. Documentation Modernization (v1.3.0) - In Progress 🚧

Implementing extension types from `EXTENSION_PROPOSAL.md`:

**Phase 1** (High Priority):
- [ ] Timestamp (UTC + optional timezone)
- [ ] Duration
- [ ] UUID/ULID (128-bit identifier)

**Phase 2**:
- [ ] Regular Expression
- [ ] Interval type
- [ ] JavaScript/TypeScript library
- [ ] Python library

**Phase 3** (Low Priority):
- [ ] Recurring events (cron-like)
- [ ] Calendar-aware operations

### 2. Field Index Mode (v1.5.0-beta)

Experimental performance optimization from `proposals/FIELD_INDEX.md`:

**Benefits**:
- 49% faster end-to-end serialization
- 91% faster encode (no string allocation)
- 45% faster decode (array access vs hash lookup)
- 27% smaller intermediate payload

**Architecture**:
- Self-describing format with embedded schema
- File header: `"BEVE" + version + flags + schemas`
- Objects use integer indexes: `{0: "value", 1: 42, 2: true}`
- Backward compatible with normal BEVE

**Target Use Cases**:
- Microservice communication (stable schema)
- API pagination (10-500 items per page)
- Cache serialization (Redis, Memcached)
- Message queues (Kafka, RabbitMQ)
- IoT telemetry (small batches, fixed schema)

## 📊 Performance Philosophy

BEVE Go aims to be:
1. **Fastest** for medium/large payloads (current: 2-4× faster than competitors)
2. **Memory efficient** (current: 2-60× fewer allocations)
3. **Production ready** (comprehensive error handling, battle-tested)
4. **Flexible** (JSON-compatible, schema-less, extensible)

### Performance Targets

| Payload Size | Target Throughput | Current Status |
|--------------|-------------------|----------------|
| Small (<1KB) | Competitive | ⚠️ Optimizing (Phase 11) |
| Medium (1-10KB) | **Best-in-class** | ✅ **2-4× faster** |
| Large (>10KB) | **Best-in-class** | ✅ **2-4× faster** |

### Optimization Strategies

1. **Core BEVE** (v1.0-1.3): Fast path optimization, zero-copy, buffer pooling
2. **Extensions** (v1.4): Native temporal/UUID types, semantic compression
3. **Field Index** (v1.5): Schema-based encoding, 49% speedup for compatible workloads
4. **Compression** (v1.6): Optional LZ4/Zstd integration, 96-98% size reduction

## 🤝 Contributing

See proposals for areas needing implementation:
1. Review existing proposals in `proposals/`
2. Discuss on GitHub Discussions
3. Implement with benchmarks
4. Submit PR with documentation

### Key Contribution Areas

- **Extension Types**: Implement temporal, UUID, or RegExp extensions
- **Field Index Mode**: Prototype and benchmark field index encoder/decoder
- **Cross-Language**: JavaScript, Python, Rust implementations
- **Tooling**: CLI inspector (`beve-inspect`), schema migration tools
- **Documentation**: Examples, best practices, migration guides

## 📖 Related Resources

- **BEVE Specification** (reference): [github.com/stephenberry/glaze](https://github.com/stephenberry/glaze)
- **C++ Implementation**: Glaze library (original BEVE)
- **Community**: [GitHub Discussions](https://github.com/meftunca/beve-go/discussions)

## 🔗 Quick Links

| Document | Description | Audience |
|----------|-------------|----------|
| [SPECIFICATION.md](../SPECIFICATION.md) | Binary format reference | Implementers |
| [EXTENSION_PROPOSAL.md](EXTENSION_PROPOSAL.md) | Temporal types, UUID, RegExp | All users |
| [FIELD_INDEX.md](proposals/FIELD_INDEX.md) | Performance optimization | Performance-critical apps |
| [MULTI_PLATFORM.md](../benchmarks/MULTI_PLATFORM.md) | Cross-platform benchmarks | DevOps, evaluators |
| [README.md](../README.md) | Getting started guide | New users |

---

**Last Updated**: October 14, 2025  
**Documentation Version**: 1.3.0

For questions or suggestions, please open an issue or discussion on GitHub.
