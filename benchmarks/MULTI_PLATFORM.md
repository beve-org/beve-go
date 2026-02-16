# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-02-16 04:58:20 UTC

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
| Apple M1 (Virtual) | 2.15μs | 470ns | 3.54μs | 1.72μs | 2.73μs |
| AMD EPYC 7763 64-Core Processor | 1.30μs | 293ns | 2.38μs | 2.41μs | 2.16μs |
| Neoverse-N2 | 555ns | 195ns | 2.27μs | 2.41μs | 2.61μs |
| Unknown CPU | 546ns | 593ns | 3.75μs | 2.01μs | 1.12μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.82μs | 29.14μs | 4.48μs | 4.41μs |
| AMD EPYC 7763 64-Core Processor | 1.51μs | 35.48μs | 6.79μs | 6.61μs |
| Neoverse-N2 | 1.06μs | 6.29μs | 3.80μs | 1.95μs |
| Unknown CPU | 2.59μs | 7.57μs | 7.19μs | 7.06μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (470ns) | 🥇 BEVE (1.82μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (293ns) | 🥇 BEVE (1.51μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (195ns) | 🥇 BEVE (1.06μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (546ns) | 🥇 BEVE (2.59μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 61.4% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 470ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.72μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 2.15μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.73μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.54μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 4.40μs | 1.9K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.82μs | 1.8K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.45μs | 2.4K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.41μs | 2.8K | 58 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.48μs | 2.5K | 54 |
| Small Struct | 🥉 JSON | Unmarshal | 29.14μs | 7.9K | 113 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.58μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.68μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 25.42μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.27μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.59μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 46.86μs | 19.3K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.68μs | 27.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.08μs | 43.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 66.65μs | 39.9K | 747 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.56μs | 33.5K | 689 |
| Medium Payload | 🥉 JSON | Unmarshal | 226.59μs | 55.9K | 721 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.80μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 131.15μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 226.07μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 382.56μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 476.87μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 541.10μs | 214.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 299.73μs | 268.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 439.71μs | 343.4K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 701.91μs | 373.5K | 6.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 860.73μs | 309.3K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.63ms | 537.2K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 293ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 887ns | 763 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.30μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.16μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 2.38μs | 896 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.41μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.51μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.25μs | 7.4K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.61μs | 4.0K | 85 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.79μs | 3.6K | 76 |
| Small Struct | 🥉 JSON | Unmarshal | 35.48μs | 7.9K | 114 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.58μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.23μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 14.76μs | 19.5K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.64μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.17μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.11μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.28μs | 28.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.01μs | 64.3K | 79 |
| Medium Payload | 🥈 CBOR | Unmarshal | 58.78μs | 25.2K | 523 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 59.05μs | 36.6K | 682 |
| Medium Payload | 🥉 JSON | Unmarshal | 236.25μs | 53.5K | 747 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.11μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.49μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 156.65μs | 215.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 206.58μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 317.23μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 430.01μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 233.31μs | 252.9K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 358.19μs | 540.0K | 579 |
| Large Payload | 🥈 MessagePack | Unmarshal | 578.98μs | 352.7K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 681.72μs | 300.5K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.23ms | 515.8K | 6.7K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 195ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 555ns | 704 | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.27μs | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.41μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.61μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.90μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.06μs | 1.2K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.95μs | 1.1K | 25 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.41μs | 5.6K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.80μs | 2.0K | 44 |
| Small Struct | 🥉 JSON | Unmarshal | 6.29μs | 1.4K | 31 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.30μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.54μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 21.05μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.73μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 34.08μs | 25.4K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 36.32μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.04μs | 29.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.36μs | 39.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.61μs | 34.1K | 627 |
| Medium Payload | 🥈 CBOR | Unmarshal | 57.76μs | 26.7K | 549 |
| Medium Payload | 🥉 JSON | Unmarshal | 167.82μs | 46.9K | 588 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.53μs | 105 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 106.45μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 178.53μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 273.95μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 290.30μs | 207.3K | 3 |
| Large Payload | 🥉 JSON | Marshal | 380.76μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 220.96μs | 263.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 274.14μs | 369.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 519.36μs | 352.7K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 653.82μs | 317.2K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.99ms | 534.0K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 546ns | 416 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 593ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.12μs | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 2.01μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.51μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.75μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.59μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.62μs | 3.7K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.06μs | 3.8K | 80 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.19μs | 3.1K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 7.57μs | 1.4K | 29 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.18μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.21μs | 19.1K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.48μs | 22.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 31.95μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 40.52μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 61.04μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 34.56μs | 29.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 56.25μs | 63.5K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 80.90μs | 39.4K | 734 |
| Medium Payload | 🥈 CBOR | Unmarshal | 91.98μs | 30.8K | 635 |
| Medium Payload | 🥉 JSON | Unmarshal | 254.06μs | 46.9K | 627 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 87.59μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 133.35μs | 196.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 167.82μs | 215.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 244.02μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 290.16μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 458.84μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 291.49μs | 274.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 471.63μs | 549.2K | 572 |
| Large Payload | 🥈 MessagePack | Unmarshal | 689.74μs | 346.3K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 950.27μs | 298.0K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.75ms | 532.1K | 7.0K |

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

