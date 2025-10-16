# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 09:01:34 UTC

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
| benchmark-darwin-apple-m1-virtual | Apple M1 (Virtual) | Darwin | [📄 Report](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md) · [📊 JSON](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.json) · [📈 Chart](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png) |
| benchmark-linux-amd-epyc-7763-64-core-processor | AMD EPYC 7763 64-Core Processor | Linux | [📄 Report](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [📊 JSON](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [📈 Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmarks/benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmarks/benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmarks/benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmarks/benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmarks/benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 560ns | 1.17μs | 2.37μs | 1.37μs | 4.23μs |
| AMD EPYC 7763 64-Core Processor | 878ns | 742ns | 5.20μs | 2.21μs | 4.37μs |
| Neoverse-N2 | 1.55μs | 719ns | 4.77μs | 973ns | 1.48μs |
| Unknown CPU | 721ns | 783ns | 2.28μs | 2.06μs | 3.53μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.37μs | 7.53μs | 5.32μs | 3.99μs |
| AMD EPYC 7763 64-Core Processor | 943ns | 15.83μs | 4.74μs | 1.73μs |
| Neoverse-N2 | 1.05μs | 4.48μs | 1.31μs | 3.55μs |
| Unknown CPU | 1.64μs | 32.68μs | 2.63μs | 5.50μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (560ns) | 🥇 BEVE (1.37μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (742ns) | 🥇 BEVE (943ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (719ns) | 🥇 BEVE (1.05μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE (721ns) | 🥇 BEVE (1.64μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 73.8% faster

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

![Benchmark Chart](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 560ns | 736 | 3 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.17μs | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.37μs | 1.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.37μs | 1.7K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 2.76μs | 1.3K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 4.23μs | 4.2K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.37μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.42μs | 2.2K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.99μs | 2.1K | 47 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.32μs | 3.6K | 76 |
| Small Struct | 🥉 JSON | Unmarshal | 7.53μs | 2.0K | 35 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.74μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 15.01μs | 24.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 16.40μs | 20.6K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 30.22μs | 18.8K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.26μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 56.03μs | 24.9K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 17.00μs | 24.0K | 57 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.83μs | 33.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 75.84μs | 47.9K | 907 |
| Medium Payload | 🥈 CBOR | Unmarshal | 104.02μs | 32.2K | 667 |
| Medium Payload | 🥉 JSON | Unmarshal | 229.12μs | 65.7K | 843 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.52μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 135.13μs | 189.6K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 171.86μs | 197.7K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 326.38μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 516.65μs | 221.9K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 575.60μs | 223.6K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 249.26μs | 280.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 378.53μs | 348.0K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 483.35μs | 344.1K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 554.65μs | 305.2K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.27ms | 515.0K | 6.7K |

[📄 View full report](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 742ns | 289 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 811ns | 985 | 3 |
| Small Struct | 🥇 BEVE | Marshal | 878ns | 1.3K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.21μs | 2.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.37μs | 8.3K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.20μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 943ns | 952 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.73μs | 832 | 20 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.19μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.74μs | 2.4K | 52 |
| Small Struct | 🥉 JSON | Unmarshal | 15.83μs | 4.2K | 67 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.77μs | 134 | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 13.98μs | 19.2K | 4 |
| Medium Payload | 🥇 BEVE | Marshal | 14.57μs | 19.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.65μs | 21.9K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 32.68μs | 16.7K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.97μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.99μs | 27.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.30μs | 60.3K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.71μs | 31.9K | 585 |
| Medium Payload | 🥈 CBOR | Unmarshal | 68.28μs | 29.2K | 599 |
| Medium Payload | 🥉 JSON | Unmarshal | 197.96μs | 49.0K | 625 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.49μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 120.75μs | 196.9K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 159.24μs | 208.9K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 203.41μs | 189.2K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 291.61μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 419.43μs | 205.3K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.93μs | 269.1K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 341.86μs | 525.1K | 559 |
| Large Payload | 🥈 MessagePack | Unmarshal | 571.72μs | 349.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 709.31μs | 310.8K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.50ms | 602.4K | 7.8K |

[📄 View full report](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 719ns | 290 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 955ns | 608 | 3 |
| Small Struct | 🥈 CBOR | Marshal | 973ns | 848 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.48μs | 2.2K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 1.55μs | 3.0K | 3 |
| Small Struct | 🥉 JSON | Marshal | 4.77μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.05μs | 1.3K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.31μs | 328 | 10 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.93μs | 4.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.55μs | 2.7K | 56 |
| Small Struct | 🥉 JSON | Unmarshal | 4.48μs | 904 | 22 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.02μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 9.68μs | 16.5K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.98μs | 20.6K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 29.98μs | 22.3K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.01μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.95μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.51μs | 25.9K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.92μs | 37.8K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.15μs | 39.6K | 738 |
| Medium Payload | 🥈 CBOR | Unmarshal | 55.30μs | 25.5K | 525 |
| Medium Payload | 🥉 JSON | Unmarshal | 216.01μs | 61.0K | 806 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.06μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 114.48μs | 197.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 195.88μs | 191.0K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 277.57μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 306.27μs | 227.4K | 4 |
| Large Payload | 🥉 JSON | Marshal | 355.90μs | 198.3K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 233.44μs | 295.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 322.06μs | 408.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 531.07μs | 350.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 624.18μs | 290.0K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.91ms | 512.2K | 6.7K |

[📄 View full report](benchmarks/benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmarks/benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 721ns | 672 | 3 |
| Small Struct | 🥉 Sonic | Marshal | 779ns | 792 | 3 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 783ns | 290 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.06μs | 1.9K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.28μs | 1.0K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.53μs | 4.2K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.64μs | 1.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.63μs | 904 | 22 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.59μs | 4.4K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.50μs | 3.2K | 69 |
| Small Struct | 🥉 JSON | Unmarshal | 32.68μs | 7.9K | 114 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.55μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 12.71μs | 16.5K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 17.68μs | 19.5K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 21.72μs | 18.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.04μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 50.54μs | 19.4K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.97μs | 26.7K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 52.81μs | 66.1K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 67.81μs | 35.3K | 652 |
| Medium Payload | 🥈 CBOR | Unmarshal | 85.84μs | 32.2K | 663 |
| Medium Payload | 🥉 JSON | Unmarshal | 262.99μs | 55.5K | 727 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.81μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 114.72μs | 188.8K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 157.54μs | 218.9K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 219.28μs | 190.3K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 286.69μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 474.81μs | 207.2K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 276.17μs | 273.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 435.86μs | 553.3K | 588 |
| Large Payload | 🥈 MessagePack | Unmarshal | 642.31μs | 325.9K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 865.97μs | 311.1K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.57ms | 533.0K | 7.0K |

[📄 View full report](benchmarks/benchmark-windows-unknown-cpu/benchmark.md)

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

