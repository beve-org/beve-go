# BEVE Go Documentation

Welcome to the BEVE Go documentation! This directory contains specifications, proposals, and design documents.

## 📚 Table of Contents

### Core Documentation

- **[SPECIFICATION.md](../SPECIFICATION.md)** - BEVE v1.0 Binary Format Specification
- **[README.md](../README.md)** - Main project README with quickstart and examples

### Proposals

#### 🎯 Active Proposals

- **[Extension Proposal: Essential Data Types](EXTENSION_PROPOSAL.md)**  
  Adds temporal types (timestamp, duration, interval), UUID/ULID, and regular expressions to BEVE.
  - Status: Draft
  - Impact: Native support for common data types, 30-55% space savings vs JSON
  - Priority: High

- **[Field Index Mode](proposals/FIELD_INDEX.md)** 🔥 **NEW**  
  Introduces integer field indexes instead of string keys for maximum performance.
  - Status: Draft
  - Impact: **49% faster** end-to-end serialization, 27% smaller intermediate payloads
  - Use cases: Microservices, API pagination, cache serialization, IoT telemetry
  - Priority: High

#### 📋 Proposal Guidelines

When creating a new proposal:
1. Copy `proposals/TEMPLATE.md` (if exists, or use existing proposals as reference)
2. Include: Abstract, Motivation, Specification, Performance Analysis, Trade-offs
3. Add benchmarks and real-world use cases
4. Submit for community discussion

### Benchmarks

- **[Multi-Platform Benchmarks](../benchmarks/MULTI_PLATFORM.md)**  
  Performance comparison across AMD64, ARM64, Windows, macOS, and Linux platforms.

### Analysis Documents

- **[CORE_ANALYSIS.md](../CORE_ANALYSIS.md)** - Deep dive into BEVE Go internals
- **[Phase Results](../)** - Development phase retrospectives (PHASE*_RESULTS.md files)

## 🎯 Current Focus Areas

### 1. Extension System (v1.4.0)

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
