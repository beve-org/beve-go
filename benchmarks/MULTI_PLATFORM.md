# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-02-23 04:57:11 UTC

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
| Apple M1 (Virtual) | 663ns | 597ns | 2.38μs | 1.12μs | 2.23μs |
| AMD EPYC 7763 64-Core Processor | 823ns | 670ns | 3.51μs | 1.38μs | 1.03μs |
| Neoverse-N2 | 973ns | 496ns | 4.67μs | 2.21μs | 1.49μs |
| Unknown CPU | 1.11μs | 810ns | 6.31μs | 2.77μs | 834ns |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 2.32μs | 33.64μs | 6.68μs | 7.75μs |
| AMD EPYC 7763 64-Core Processor | 855ns | 10.44μs | 6.48μs | 6.11μs |
| Neoverse-N2 | 1.55μs | 3.04μs | 3.18μs | 3.44μs |
| Unknown CPU | 1.25μs | 30.11μs | 7.10μs | 1.14μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (597ns) | 🥇 BEVE (2.32μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (670ns) | 🥇 BEVE (855ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (496ns) | 🥉 Sonic (1.31μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (810ns) | 🥈 MessagePack (1.14μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 77.5% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 597ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 663ns | 640 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.12μs | 640 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.23μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 2.38μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 7.53μs | 3.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.32μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.17μs | 2.9K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.68μs | 3.2K | 68 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.75μs | 4.6K | 96 |
| Small Struct | 🥉 JSON | Unmarshal | 33.64μs | 7.5K | 99 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.74μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 16.07μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 24.07μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.69μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 45.49μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 62.21μs | 22.1K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.53μs | 29.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.34μs | 42.1K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 59.32μs | 23.6K | 486 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 61.08μs | 31.2K | 571 |
| Medium Payload | 🥉 JSON | Unmarshal | 225.95μs | 50.0K | 679 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 75.11μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 138.66μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 153.77μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 399.82μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 567.38μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 637.30μs | 214.5K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 282.53μs | 271.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 362.50μs | 365.8K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 607.44μs | 380.7K | 7.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 678.47μs | 327.9K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.58ms | 557.0K | 7.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 670ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 704ns | 791 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 823ns | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.03μs | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 1.38μs | 1.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.51μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 855ns | 728 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.66μs | 7.8K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.11μs | 4.8K | 102 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.48μs | 3.6K | 78 |
| Small Struct | 🥉 JSON | Unmarshal | 10.44μs | 2.3K | 45 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.11μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.13μs | 19.1K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 14.20μs | 19.5K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.80μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.21μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 41.47μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.71μs | 28.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.67μs | 44.1K | 63 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 51.94μs | 32.6K | 596 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.38μs | 31.9K | 662 |
| Medium Payload | 🥉 JSON | Unmarshal | 223.78μs | 55.0K | 720 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.58μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.56μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 160.61μs | 223.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 214.15μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 321.55μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 443.94μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 254.39μs | 279.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 364.26μs | 532.3K | 572 |
| Large Payload | 🥈 MessagePack | Unmarshal | 535.47μs | 314.5K | 5.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 735.97μs | 324.7K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.35ms | 540.2K | 7.1K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 496ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 973ns | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.49μs | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 2.21μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.92μs | 3.1K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.67μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.31μs | 1.3K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.55μs | 2.6K | 4 |
| Small Struct | 🥉 JSON | Unmarshal | 3.04μs | 552 | 15 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.18μs | 1.6K | 36 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.44μs | 2.5K | 55 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.86μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.15μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.91μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 25.27μs | 18.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.07μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.82μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.94μs | 25.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.99μs | 44.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.01μs | 34.4K | 635 |
| Medium Payload | 🥈 CBOR | Unmarshal | 60.25μs | 28.4K | 581 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.04μs | 59.4K | 817 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.85μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 104.11μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 195.66μs | 205.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 274.80μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 298.26μs | 206.8K | 3 |
| Large Payload | 🥉 JSON | Marshal | 369.33μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 236.24μs | 281.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 294.91μs | 396.8K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 545.48μs | 363.9K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 674.16μs | 322.7K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.12ms | 576.9K | 7.6K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 810ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 834ns | 520 | 5 |
| Small Struct | 🥇 BEVE | Marshal | 1.11μs | 1.4K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.71μs | 1.8K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.77μs | 3.1K | 1 |
| Small Struct | 🥉 JSON | Marshal | 6.31μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.14μs | 304 | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.25μs | 1.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.07μs | 3.5K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.10μs | 2.9K | 63 |
| Small Struct | 🥉 JSON | Unmarshal | 30.11μs | 7.7K | 105 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.60μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.61μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.46μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.45μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.06μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 53.19μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.68μs | 30.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 47.89μs | 56.9K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 66.25μs | 32.7K | 600 |
| Medium Payload | 🥈 CBOR | Unmarshal | 97.15μs | 34.5K | 712 |
| Medium Payload | 🥉 JSON | Unmarshal | 300.43μs | 61.4K | 806 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 85.95μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.19μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 169.23μs | 223.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 220.87μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 291.54μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 440.03μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 296.07μs | 278.0K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 441.44μs | 521.5K | 560 |
| Large Payload | 🥈 MessagePack | Unmarshal | 679.66μs | 340.8K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 837.42μs | 301.7K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.64ms | 527.1K | 7.0K |

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

