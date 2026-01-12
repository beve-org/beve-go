# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-01-12 04:11:47 UTC

This report consolidates benchmark results from multiple platforms tested in CI/CD.

## � Visual Comparisons

### Overall Performance Comparison
![Multi-Platform Comparison](charts/multi_platform_comparison.png)

### Performance Heatmap
![Performance Heatmap](charts/performance_heatmap.png)

### Memory Efficiency
![Memory Comparison](charts/memory_comparison.png)

---

## �🖥️ Tested Platforms

| Platform | CPU | OS | Artifacts |
|----------|-----|----|-----------| 
| benchmark-darwin-apple-m1-virtual | Apple M1 (Virtual) | Darwin | [📄 Report](benchmark-darwin-apple-m1-virtual/benchmark.md) · [📊 JSON](benchmark-darwin-apple-m1-virtual/benchmark.json) · [📈 Chart](benchmark-darwin-apple-m1-virtual/benchmark.png) |
| benchmark-linux-amd-epyc-7763-64-core-processor | AMD EPYC 7763 64-Core Processor | Linux | [📄 Report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [📊 JSON](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [📈 Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 632ns | 708ns | 2.02μs | 2.12μs | 1.29μs |
| AMD EPYC 7763 64-Core Processor | 796ns | 476ns | 1.85μs | 2.65μs | 4.42μs |
| Neoverse-N2 | 363ns | 790ns | 1.27μs | 890ns | 1.59μs |
| Unknown CPU | 455ns | 508ns | 5.80μs | 3.11μs | 1.75μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 988ns | 25.11μs | 9.31μs | 2.84μs |
| AMD EPYC 7763 64-Core Processor | 1.93μs | 14.47μs | 2.90μs | 6.92μs |
| Neoverse-N2 | 1.20μs | 15.95μs | 4.13μs | 1.96μs |
| Unknown CPU | 1.47μs | 11.43μs | 2.14μs | 5.52μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (632ns) | 🥇 BEVE (988ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (476ns) | 🥇 BEVE (1.93μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (363ns) | 🥇 BEVE (1.20μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (455ns) | 🥇 BEVE (1.47μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 72.3% faster

### Platform Details

- **Apple M1 (Virtual)** (Darwin)
  - Architecture: arm64
  - Test Scenarios: 3

- **AMD EPYC 7763 64-Core Processor** (Linux)
  - Architecture: x86_64
  - Test Scenarios: 3

- **Neoverse-N2** (Linux)
  - Architecture: aarch64
  - Test Scenarios: 3

- **Unknown CPU** (Windows)
  - Architecture: AMD64
  - Test Scenarios: 3

---

## 📋 Detailed Platform Results

### Apple M1 (Virtual) — Darwin

![Benchmark Chart](benchmark-darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 632ns | 1.2K | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 708ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.29μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 2.02μs | 1.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.12μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 4.33μs | 2.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 988ns | 2.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.84μs | 2.1K | 47 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.26μs | 4.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.31μs | 2.4K | 52 |
| Small Struct | 🥉 JSON | Unmarshal | 25.11μs | 8.0K | 117 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.65μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.34μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.44μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.59μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.70μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 51.67μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.41μs | 27.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.99μs | 41.6K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 37.55μs | 30.1K | 548 |
| Medium Payload | 🥈 CBOR | Unmarshal | 60.22μs | 37.6K | 775 |
| Medium Payload | 🥉 JSON | Unmarshal | 255.94μs | 65.4K | 875 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 54.49μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 82.65μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 212.99μs | 526.8K | 115 |
| Large Payload | 🥈 CBOR | Marshal | 220.73μs | 204.9K | 1 |
| Large Payload | 🥉 JSON | Marshal | 371.18μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 445.92μs | 213.8K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 222.09μs | 284.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 472.47μs | 341.1K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 481.23μs | 349.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 891.06μs | 316.9K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.83ms | 506.2K | 6.5K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 476ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 796ns | 768 | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.85μs | 896 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.96μs | 2.7K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.65μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.42μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.93μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.70μs | 4.2K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.90μs | 1.3K | 30 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.92μs | 5.2K | 106 |
| Small Struct | 🥉 JSON | Unmarshal | 14.47μs | 4.0K | 61 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.07μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.76μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.39μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.84μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.94μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 49.57μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.10μs | 29.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.19μs | 51.7K | 73 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 65.05μs | 43.9K | 828 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.84μs | 32.0K | 655 |
| Medium Payload | 🥉 JSON | Unmarshal | 229.85μs | 55.3K | 736 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.95μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 112.71μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 166.08μs | 231.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 203.87μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 330.09μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 438.69μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 261.70μs | 285.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 358.74μs | 538.0K | 568 |
| Large Payload | 🥈 MessagePack | Unmarshal | 553.60μs | 345.5K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 726.13μs | 324.8K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.21ms | 503.4K | 6.6K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 363ns | 352 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 772ns | 389 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 790ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 890ns | 768 | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.27μs | 640 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.59μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.20μs | 1.7K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.96μs | 1.1K | 25 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.24μs | 3.2K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.13μs | 2.2K | 49 |
| Small Struct | 🥉 JSON | Unmarshal | 15.95μs | 4.4K | 74 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.86μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.88μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.30μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 26.06μs | 18.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.80μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.30μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.85μs | 26.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.65μs | 44.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.51μs | 44.8K | 840 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.30μs | 32.2K | 661 |
| Medium Payload | 🥉 JSON | Unmarshal | 227.02μs | 68.2K | 854 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.75μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.54μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 194.27μs | 205.2K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 257.77μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 316.63μs | 231.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 382.14μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 228.51μs | 281.3K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 295.58μs | 406.6K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 492.80μs | 326.5K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 665.13μs | 322.4K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 1.88ms | 504.7K | 6.6K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 455ns | 352 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 508ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.17μs | 1.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.75μs | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 3.11μs | 3.1K | 1 |
| Small Struct | 🥉 JSON | Marshal | 5.80μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.47μs | 1.5K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.14μs | 704 | 18 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.52μs | 3.5K | 75 |
| Small Struct | 🥉 Sonic | Unmarshal | 6.23μs | 7.8K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 11.43μs | 2.3K | 44 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.19μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.19μs | 14.3K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.42μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.74μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.42μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 43.87μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.53μs | 28.4K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.38μs | 50.3K | 71 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.50μs | 30.2K | 549 |
| Medium Payload | 🥈 CBOR | Unmarshal | 94.34μs | 34.9K | 720 |
| Medium Payload | 🥉 JSON | Unmarshal | 286.74μs | 60.2K | 796 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.80μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 121.32μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 162.31μs | 207.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 252.17μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 286.87μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 466.13μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 270.46μs | 264.8K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 546.65μs | 528.2K | 571 |
| Large Payload | 🥈 MessagePack | Unmarshal | 712.85μs | 366.9K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 1.08ms | 348.5K | 7.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.84ms | 535.4K | 7.0K |

[📄 View full report](benchmark-windows-unknown-cpu/benchmark.md)

---

## 📚 Additional Resources

- [BEVE Specification](../SPECIFICATION.md)
- [Go Package Documentation](../README.md)
- [Translator Package](../translator/README.md)
- [Examples](../examples/)

**Legend:**
- 🥇 BEVE family (fastest)
- 🥈 CBOR/MessagePack (fast)
- 🥉 JSON/Sonic (standard)

