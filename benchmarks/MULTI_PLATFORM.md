# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 14:15:28 UTC

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
| Apple M1 (Virtual) | 995ns | 1.01μs | 721ns | 1.94μs | 3.40μs |
| AMD EPYC 7763 64-Core Processor | 777ns | 1.73μs | 7.00μs | 1.97μs | 1.55μs |
| Neoverse-N2 | 1.17μs | 948ns | 2.60μs | 1.29μs | 965ns |
| Unknown CPU | 2.03μs | 490ns | 3.86μs | 2.16μs | 4.00μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.74μs | 6.62μs | 3.67μs | 4.62μs |
| AMD EPYC 7763 64-Core Processor | 1.82μs | 31.65μs | 2.60μs | 8.35μs |
| Neoverse-N2 | 1.31μs | 23.64μs | 5.13μs | 3.71μs |
| Unknown CPU | 2.70μs | 15.07μs | 9.98μs | 2.87μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥉 JSON (721ns) | 🥇 BEVE (1.74μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (777ns) | 🥇 BEVE (1.82μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (948ns) | 🥇 BEVE (1.31μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (490ns) | 🥇 BEVE (2.70μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 38.3% faster

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
| Small Struct | 🥉 JSON | Marshal | 721ns | 560 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 995ns | 1.7K | 3 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.01μs | 288 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.94μs | 1.7K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.40μs | 8.3K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 5.68μs | 3.3K | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.74μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.37μs | 5.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.67μs | 3.1K | 65 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.62μs | 3.2K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 6.62μs | 1.3K | 28 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.59μs | 128 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.83μs | 21.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 15.66μs | 14.4K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.32μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.08μs | 20.8K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 39.66μs | 18.7K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.04μs | 23.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.86μs | 42.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.83μs | 33.2K | 608 |
| Medium Payload | 🥈 CBOR | Unmarshal | 57.94μs | 31.4K | 644 |
| Medium Payload | 🥉 JSON | Unmarshal | 237.60μs | 53.2K | 695 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 74.79μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 136.33μs | 181.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 189.51μs | 189.4K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 203.39μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 378.47μs | 213.8K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 413.97μs | 214.2K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 172.36μs | 251.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 271.06μs | 336.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 454.07μs | 355.3K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 555.60μs | 339.9K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.95ms | 524.4K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 777ns | 576 | 3 |
| Small Struct | 🥉 Sonic | Marshal | 830ns | 956 | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.55μs | 1.2K | 6 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.73μs | 291 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.97μs | 1.9K | 2 |
| Small Struct | 🥉 JSON | Marshal | 7.00μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.82μs | 2.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.60μs | 1.0K | 24 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.44μs | 3.9K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 8.35μs | 5.2K | 107 |
| Small Struct | 🥉 JSON | Unmarshal | 31.65μs | 7.6K | 103 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.26μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 16.62μs | 19.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.86μs | 19.2K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 23.86μs | 25.4K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 44.36μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 62.11μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.48μs | 31.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.23μs | 51.9K | 69 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.21μs | 39.3K | 738 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.61μs | 28.1K | 581 |
| Medium Payload | 🥉 JSON | Unmarshal | 205.25μs | 49.6K | 633 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 85.37μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 114.98μs | 180.5K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 163.78μs | 217.5K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 202.66μs | 189.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 299.17μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 444.55μs | 222.1K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 223.49μs | 258.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 364.36μs | 553.0K | 581 |
| Large Payload | 🥈 MessagePack | Unmarshal | 569.18μs | 351.0K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 707.38μs | 305.6K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.28ms | 534.7K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 948ns | 289 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 965ns | 1.2K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 1.17μs | 1.7K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 1.29μs | 1.3K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.60μs | 1.6K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 4.08μs | 3.3K | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.31μs | 2.1K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.88μs | 2.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.71μs | 2.9K | 63 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.13μs | 2.9K | 63 |
| Small Struct | 🥉 JSON | Unmarshal | 23.64μs | 7.7K | 107 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.62μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 9.76μs | 16.5K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.38μs | 20.6K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.53μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 31.40μs | 25.0K | 4 |
| Medium Payload | 🥉 JSON | Marshal | 38.58μs | 22.1K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.96μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 30.08μs | 39.4K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.83μs | 27.0K | 484 |
| Medium Payload | 🥈 CBOR | Unmarshal | 61.54μs | 29.3K | 602 |
| Medium Payload | 🥉 JSON | Unmarshal | 237.08μs | 68.2K | 892 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.38μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 108.81μs | 182.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 187.73μs | 190.7K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 278.38μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 310.62μs | 228.5K | 4 |
| Large Payload | 🥉 JSON | Marshal | 403.22μs | 224.8K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 233.92μs | 278.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 296.08μs | 405.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 523.49μs | 354.4K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 691.18μs | 333.6K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.89ms | 493.9K | 6.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 490ns | 290 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.03μs | 2.3K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.16μs | 1.7K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 2.90μs | 3.4K | 3 |
| Small Struct | 🥉 JSON | Marshal | 3.86μs | 1.7K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.00μs | 4.2K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.70μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.87μs | 1.3K | 29 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.31μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.98μs | 4.0K | 86 |
| Small Struct | 🥉 JSON | Unmarshal | 15.07μs | 3.8K | 57 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.09μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 16.64μs | 20.6K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 22.20μs | 25.1K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 26.45μs | 24.7K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 43.84μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.14μs | 20.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.08μs | 23.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 52.31μs | 62.3K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 77.76μs | 38.6K | 724 |
| Medium Payload | 🥈 CBOR | Unmarshal | 102.69μs | 36.6K | 751 |
| Medium Payload | 🥉 JSON | Unmarshal | 236.65μs | 44.9K | 628 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.80μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 124.85μs | 180.9K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 187.88μs | 218.7K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 221.96μs | 198.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 329.58μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 537.83μs | 215.4K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 332.51μs | 282.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 404.30μs | 497.9K | 558 |
| Large Payload | 🥈 MessagePack | Unmarshal | 651.91μs | 321.7K | 5.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 885.13μs | 311.8K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.82ms | 553.4K | 7.3K |

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

