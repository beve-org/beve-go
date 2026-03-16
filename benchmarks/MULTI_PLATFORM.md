# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-03-16 05:14:03 UTC

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
| Apple M1 (Virtual) | 524ns | 296ns | 2.31μs | 1.44μs | 925ns |
| AMD EPYC 7763 64-Core Processor | 888ns | 908ns | 3.74μs | 1.79μs | 2.81μs |
| Neoverse-N2 | 977ns | 455ns | 944ns | 976ns | 2.40μs |
| Unknown CPU | 2.40μs | 723ns | 1.08μs | 1.36μs | 1.21μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 466ns | 8.76μs | 1.53μs | 2.76μs |
| AMD EPYC 7763 64-Core Processor | 1.88μs | 27.36μs | 1.52μs | 2.93μs |
| Neoverse-N2 | 1.68μs | 23.54μs | 7.15μs | 1.26μs |
| Unknown CPU | 1.38μs | 28.72μs | 4.11μs | 3.33μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (296ns) | 🥇 BEVE (466ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (888ns) | 🥈 CBOR (1.52μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (455ns) | 🥈 MessagePack (1.26μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (723ns) | 🥇 BEVE (1.38μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 7.0% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 296ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 524ns | 1.4K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 925ns | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.01μs | 517 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.44μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.31μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 466ns | 504 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.34μs | 1.5K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.53μs | 1.1K | 25 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.76μs | 3.6K | 78 |
| Small Struct | 🥉 JSON | Unmarshal | 8.76μs | 3.8K | 55 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.39μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 8.85μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 12.84μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.54μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 34.29μs | 20.7K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 34.89μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 15.28μs | 29.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 23.85μs | 34.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 34.53μs | 35.9K | 664 |
| Medium Payload | 🥈 CBOR | Unmarshal | 47.26μs | 35.5K | 732 |
| Medium Payload | 🥉 JSON | Unmarshal | 160.65μs | 59.0K | 762 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 53.92μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 79.56μs | 204.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 144.05μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 180.06μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 300.75μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 377.99μs | 221.5K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 157.33μs | 275.9K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 249.18μs | 342.2K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 370.76μs | 361.5K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 442.15μs | 295.4K | 6.0K |
| Large Payload | 🥉 JSON | Unmarshal | 1.52ms | 519.2K | 6.7K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 888ns | 1.4K | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 908ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.79μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.85μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.81μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.74μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.52μs | 424 | 12 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.60μs | 2.0K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.88μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.93μs | 1.9K | 41 |
| Small Struct | 🥉 JSON | Unmarshal | 27.36μs | 7.9K | 113 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.93μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 17.76μs | 27.3K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.63μs | 25.3K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.65μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.58μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.20μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.90μs | 34.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.07μs | 43.6K | 67 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 59.74μs | 37.6K | 703 |
| Medium Payload | 🥈 CBOR | Unmarshal | 76.47μs | 35.8K | 732 |
| Medium Payload | 🥉 JSON | Unmarshal | 238.59μs | 60.6K | 753 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.14μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.73μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 156.65μs | 215.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 214.51μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 329.90μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 439.24μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 255.97μs | 284.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 369.53μs | 525.9K | 568 |
| Large Payload | 🥈 MessagePack | Unmarshal | 573.00μs | 339.7K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 692.31μs | 298.7K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.21ms | 505.4K | 6.6K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 455ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 944ns | 416 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 976ns | 896 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 977ns | 1.5K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.40μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 2.82μs | 2.1K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.26μs | 448 | 12 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.68μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.87μs | 4.5K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.15μs | 4.4K | 93 |
| Small Struct | 🥉 JSON | Unmarshal | 23.54μs | 7.7K | 107 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.72μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.59μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.48μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 27.50μs | 19.4K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.86μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.46μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.80μs | 28.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.08μs | 37.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.91μs | 37.5K | 696 |
| Medium Payload | 🥈 CBOR | Unmarshal | 64.53μs | 31.7K | 651 |
| Medium Payload | 🥉 JSON | Unmarshal | 240.44μs | 72.2K | 911 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.50μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.26μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 182.48μs | 188.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 269.53μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 314.48μs | 225.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 402.07μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.84μs | 285.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 296.70μs | 409.8K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 504.08μs | 334.6K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 646.05μs | 310.4K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 569.5K | 7.4K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 723ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 1.08μs | 384 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.21μs | 1.0K | 6 |
| Small Struct | 🥉 Sonic | Marshal | 1.26μs | 1.4K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.36μs | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 2.40μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.38μs | 1.1K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.33μs | 1.7K | 38 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.11μs | 1.4K | 32 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.43μs | 4.4K | 9 |
| Small Struct | 🥉 JSON | Unmarshal | 28.72μs | 7.4K | 97 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.15μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.23μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.20μs | 18.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 30.81μs | 27.3K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.67μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.58μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.59μs | 28.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 54.21μs | 63.9K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 75.16μs | 37.2K | 693 |
| Medium Payload | 🥈 CBOR | Unmarshal | 80.84μs | 27.3K | 564 |
| Medium Payload | 🥉 JSON | Unmarshal | 260.56μs | 46.5K | 655 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 84.22μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 123.34μs | 196.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 175.79μs | 214.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 237.99μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 296.83μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 468.69μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 338.99μs | 273.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 445.02μs | 506.9K | 557 |
| Large Payload | 🥈 MessagePack | Unmarshal | 710.85μs | 332.2K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 885.89μs | 299.6K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.80ms | 542.9K | 7.1K |

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

