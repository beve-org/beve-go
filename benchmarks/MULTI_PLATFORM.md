# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-11-17 03:49:07 UTC

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
| Apple M1 (Virtual) | 1.14μs | 851ns | 6.68μs | 2.84μs | 2.01μs |
| AMD EPYC 7763 64-Core Processor | 819ns | 850ns | 4.08μs | 1.30μs | 2.86μs |
| Neoverse-N2 | 485ns | 665ns | 3.82μs | 1.84μs | 2.40μs |
| Unknown CPU | 2.20μs | 572ns | 2.71μs | 927ns | 4.39μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.06μs | 10.94μs | 4.36μs | 4.16μs |
| AMD EPYC 7763 64-Core Processor | 861ns | 5.62μs | 6.42μs | 1.05μs |
| Neoverse-N2 | 1.18μs | 5.35μs | 4.09μs | 2.27μs |
| Unknown CPU | 797ns | 21.27μs | 8.60μs | 7.42μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (851ns) | 🥇 BEVE (1.06μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (819ns) | 🥇 BEVE (861ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (485ns) | 🥇 BEVE (1.18μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (572ns) | 🥇 BEVE (797ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 67.3% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 851ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.14μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.91μs | 926 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.01μs | 4.1K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.84μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 6.68μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.06μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.79μs | 5.0K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.16μs | 3.8K | 80 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.36μs | 2.3K | 51 |
| Small Struct | 🥉 JSON | Unmarshal | 10.94μs | 2.3K | 44 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.97μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.41μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.51μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.11μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 45.98μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 61.89μs | 24.9K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.21μs | 29.4K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 35.59μs | 23.3K | 407 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.72μs | 42.7K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 57.38μs | 26.2K | 542 |
| Medium Payload | 🥉 JSON | Unmarshal | 217.55μs | 56.8K | 735 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 59.80μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 110.54μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 199.75μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 304.20μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 393.16μs | 205.1K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 524.46μs | 214.5K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 265.61μs | 257.3K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 365.97μs | 340.1K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 572.86μs | 355.4K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 720.30μs | 314.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 509.9K | 6.7K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 819ns | 704 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 850ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.30μs | 1.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.43μs | 1.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.86μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.08μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 861ns | 760 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.05μs | 256 | 7 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.43μs | 7.8K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 5.62μs | 1.2K | 25 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.42μs | 3.2K | 69 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.01μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.38μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.11μs | 22.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.11μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.82μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 45.12μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.88μs | 27.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.44μs | 57.8K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.88μs | 37.5K | 698 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.21μs | 23.6K | 486 |
| Medium Payload | 🥉 JSON | Unmarshal | 244.08μs | 57.8K | 768 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.26μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.23μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 159.94μs | 215.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 203.75μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 321.91μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 450.54μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 250.00μs | 276.3K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 370.81μs | 544.5K | 572 |
| Large Payload | 🥈 MessagePack | Unmarshal | 566.60μs | 337.6K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 700.17μs | 302.9K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.35ms | 542.3K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 485ns | 640 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 665ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.26μs | 798 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.84μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.40μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.82μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.18μs | 1.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.27μs | 1.5K | 33 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.36μs | 3.9K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.09μs | 2.2K | 49 |
| Small Struct | 🥉 JSON | Unmarshal | 5.35μs | 1.3K | 27 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.23μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.94μs | 14.3K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.01μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 26.51μs | 18.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.32μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 36.68μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 19.99μs | 22.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.92μs | 38.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.05μs | 42.8K | 806 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.69μs | 33.7K | 691 |
| Medium Payload | 🥉 JSON | Unmarshal | 211.31μs | 61.1K | 790 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 63.66μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 102.10μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 182.71μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 251.79μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 301.66μs | 214.3K | 3 |
| Large Payload | 🥉 JSON | Marshal | 379.89μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 226.24μs | 281.1K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 271.56μs | 367.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 501.82μs | 346.0K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 638.90μs | 311.7K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.81ms | 475.9K | 6.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 572ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 927ns | 768 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.32μs | 1.8K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.20μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.71μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.39μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 797ns | 408 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.79μs | 7.3K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.42μs | 5.2K | 107 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.60μs | 3.9K | 84 |
| Small Struct | 🥉 JSON | Unmarshal | 21.27μs | 4.5K | 77 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.53μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.15μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.27μs | 20.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 26.73μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.43μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 45.47μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.71μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.37μs | 56.9K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 68.21μs | 34.7K | 638 |
| Medium Payload | 🥈 CBOR | Unmarshal | 90.19μs | 34.3K | 705 |
| Medium Payload | 🥉 JSON | Unmarshal | 280.82μs | 59.8K | 769 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.73μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 126.36μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 153.64μs | 206.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 250.81μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 278.50μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 505.25μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.83μs | 272.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 411.07μs | 522.8K | 573 |
| Large Payload | 🥈 MessagePack | Unmarshal | 866.25μs | 345.4K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 1.04ms | 331.0K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.66ms | 543.8K | 7.1K |

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

