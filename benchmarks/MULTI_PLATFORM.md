# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-20 03:50:06 UTC

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
| Apple M1 (Virtual) | 210ns | 373ns | 3.25μs | 2.71μs | 3.18μs |
| AMD EPYC 7763 64-Core Processor | 597ns | 875ns | 5.44μs | 1.36μs | 4.57μs |
| Neoverse-N2 | 1.09μs | 288ns | 1.26μs | 978ns | 2.41μs |
| Unknown CPU | 688ns | 751ns | 2.54μs | 1.37μs | 5.40μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 999ns | 6.80μs | 6.46μs | 2.26μs |
| AMD EPYC 7763 64-Core Processor | 1.11μs | 28.27μs | 1.11μs | 4.52μs |
| Neoverse-N2 | 898ns | 19.81μs | 3.06μs | 3.30μs |
| Unknown CPU | 1.83μs | 22.09μs | 5.42μs | 7.59μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (210ns) | 🥇 BEVE (999ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥉 Sonic (386ns) | 🥈 CBOR (1.11μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (288ns) | 🥇 BEVE (898ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥉 Sonic (582ns) | 🥉 Sonic (999ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 67.2% faster

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
| Small Struct | 🥇 BEVE | Marshal | 210ns | 176 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 373ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 2.71μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.18μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 3.25μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 5.57μs | 2.7K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 999ns | 888 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.21μs | 1.7K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.26μs | 1.3K | 29 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.46μs | 3.9K | 83 |
| Small Struct | 🥉 JSON | Unmarshal | 6.80μs | 1.4K | 29 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.58μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.54μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.51μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.35μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 45.02μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 46.99μs | 19.3K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.00μs | 24.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.75μs | 46.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.26μs | 39.7K | 743 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.16μs | 32.4K | 664 |
| Medium Payload | 🥉 JSON | Unmarshal | 236.90μs | 56.8K | 758 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 55.68μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 97.21μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 161.97μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 280.60μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 403.26μs | 213.4K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 645.14μs | 222.6K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 253.05μs | 266.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 343.89μs | 361.8K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 477.17μs | 328.2K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 641.82μs | 308.0K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.30ms | 485.3K | 6.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 386ns | 293 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 597ns | 512 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 875ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.36μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.57μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.44μs | 3.1K | 1 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.11μs | 232 | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.11μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.11μs | 7.4K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.52μs | 3.3K | 70 |
| Small Struct | 🥉 JSON | Unmarshal | 28.27μs | 8.0K | 116 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.41μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.38μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.10μs | 20.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.76μs | 21.8K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 36.84μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.92μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.99μs | 26.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.27μs | 55.9K | 74 |
| Medium Payload | 🥈 CBOR | Unmarshal | 58.68μs | 24.9K | 516 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 67.72μs | 42.3K | 789 |
| Medium Payload | 🥉 JSON | Unmarshal | 199.11μs | 46.4K | 611 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.38μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.83μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 153.97μs | 215.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 209.00μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 307.16μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 432.16μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 247.05μs | 286.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 353.33μs | 537.6K | 574 |
| Large Payload | 🥈 MessagePack | Unmarshal | 582.98μs | 366.6K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 702.22μs | 311.5K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.13ms | 514.9K | 6.6K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 288ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 978ns | 896 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.09μs | 2.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.26μs | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.73μs | 1.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.41μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 898ns | 824 | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.06μs | 1.5K | 34 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.30μs | 2.5K | 53 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.31μs | 5.5K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 19.81μs | 7.2K | 91 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.45μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.84μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.23μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.68μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 36.06μs | 27.9K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 37.27μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.02μs | 33.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.43μs | 44.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.56μs | 33.1K | 608 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.36μs | 32.4K | 663 |
| Medium Payload | 🥉 JSON | Unmarshal | 225.74μs | 64.3K | 841 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.29μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 100.74μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 178.98μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 287.03μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 316.54μs | 231.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 389.56μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 231.50μs | 281.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 293.02μs | 405.8K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 535.16μs | 364.2K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 687.85μs | 341.9K | 7.0K |
| Large Payload | 🥉 JSON | Unmarshal | 2.01ms | 546.2K | 7.2K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 582ns | 485 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 688ns | 576 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 751ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.37μs | 1.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.54μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 5.40μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Unmarshal | 999ns | 800 | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.83μs | 2.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.42μs | 2.3K | 51 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.59μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 22.09μs | 4.6K | 81 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.49μs | 5 | 0 |
| Medium Payload | 🥉 Sonic | Marshal | 15.22μs | 18.8K | 3 |
| Medium Payload | 🥇 BEVE | Marshal | 15.88μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 22.01μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.33μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 54.78μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.99μs | 29.7K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 47.35μs | 56.0K | 67 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 77.00μs | 41.5K | 782 |
| Medium Payload | 🥈 CBOR | Unmarshal | 79.15μs | 27.8K | 573 |
| Medium Payload | 🥉 JSON | Unmarshal | 266.28μs | 52.9K | 731 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.78μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 117.11μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 165.91μs | 215.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 217.18μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 333.15μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 486.86μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 330.62μs | 263.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 442.08μs | 547.1K | 586 |
| Large Payload | 🥈 MessagePack | Unmarshal | 674.69μs | 340.4K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 872.60μs | 314.5K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.72ms | 557.4K | 7.3K |

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

