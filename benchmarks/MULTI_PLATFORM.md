# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 16:15:01 UTC

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
| Apple M1 (Virtual) | 1.70μs | 294ns | 2.80μs | 2.52μs | 2.51μs |
| AMD EPYC 7763 64-Core Processor | 615ns | 664ns | 5.40μs | 826ns | 2.94μs |
| Neoverse-N2 | 1.21μs | 383ns | 2.23μs | 900ns | 4.12μs |
| Unknown CPU | 1.45μs | 580ns | 5.83μs | 628ns | 2.95μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 2.63μs | 9.07μs | 2.97μs | 4.87μs |
| AMD EPYC 7763 64-Core Processor | 1.80μs | 27.46μs | 4.82μs | 1.28μs |
| Neoverse-N2 | 657ns | 14.22μs | 5.31μs | 1.55μs |
| Unknown CPU | 1.00μs | 8.99μs | 3.76μs | 3.87μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (294ns) | 🥇 BEVE (2.63μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (615ns) | 🥈 MessagePack (1.28μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (383ns) | 🥇 BEVE (657ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (580ns) | 🥇 BEVE (1.00μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 62.2% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 294ns | 289 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.70μs | 3.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.51μs | 4.2K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.52μs | 2.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.80μs | 1.9K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 4.93μs | 2.2K | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.63μs | 3.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.97μs | 1.6K | 36 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.23μs | 5.5K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.87μs | 3.2K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 9.07μs | 2.1K | 37 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.97μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.98μs | 24.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.67μs | 16.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.57μs | 33.1K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 46.74μs | 22.1K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 64.54μs | 24.9K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 32.79μs | 34.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.90μs | 41.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 61.00μs | 36.4K | 669 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.07μs | 30.0K | 618 |
| Medium Payload | 🥉 JSON | Unmarshal | 263.93μs | 60.7K | 814 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.72μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 143.82μs | 197.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 159.08μs | 181.5K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 424.01μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 464.77μs | 213.8K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 592.87μs | 215.6K | 4 |
| Large Payload | 🥉 Sonic | Unmarshal | 320.07μs | 359.1K | 211 |
| Large Payload | 🥇 BEVE | Unmarshal | 364.28μs | 278.0K | 419 |
| Large Payload | 🥈 MessagePack | Unmarshal | 419.91μs | 335.3K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 513.78μs | 334.9K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.00ms | 527.4K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 615ns | 672 | 3 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 664ns | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 826ns | 656 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.44μs | 2.1K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.94μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.40μs | 3.2K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.28μs | 456 | 12 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.61μs | 1.9K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.80μs | 3.0K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.82μs | 2.2K | 48 |
| Small Struct | 🥉 JSON | Unmarshal | 27.46μs | 7.9K | 113 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.42μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.43μs | 20.7K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 16.65μs | 22.7K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 19.33μs | 18.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.98μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.69μs | 19.5K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.34μs | 28.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.09μs | 54.9K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.19μs | 35.0K | 646 |
| Medium Payload | 🥈 CBOR | Unmarshal | 84.42μs | 36.5K | 748 |
| Medium Payload | 🥉 JSON | Unmarshal | 213.79μs | 51.5K | 690 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.70μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 129.26μs | 197.1K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 179.38μs | 233.9K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 198.98μs | 189.3K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 321.16μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 452.59μs | 221.9K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.79μs | 276.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 369.17μs | 561.5K | 583 |
| Large Payload | 🥈 MessagePack | Unmarshal | 548.44μs | 341.0K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 755.36μs | 303.1K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.36ms | 552.9K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 383ns | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 900ns | 720 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.21μs | 1.8K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 1.50μs | 1.1K | 3 |
| Small Struct | 🥉 JSON | Marshal | 2.23μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.12μs | 8.3K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 657ns | 312 | 3 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.55μs | 688 | 17 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.40μs | 3.8K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.31μs | 2.9K | 63 |
| Small Struct | 🥉 JSON | Unmarshal | 14.22μs | 4.1K | 65 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.59μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 10.91μs | 19.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 17.23μs | 18.6K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 30.77μs | 16.7K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 33.18μs | 25.2K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.66μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.33μs | 30.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.81μs | 44.8K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.29μs | 36.4K | 679 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.04μs | 38.0K | 776 |
| Medium Payload | 🥉 JSON | Unmarshal | 202.22μs | 54.3K | 739 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.36μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 124.62μs | 198.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 189.40μs | 190.5K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 284.40μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 312.98μs | 218.8K | 4 |
| Large Payload | 🥉 JSON | Marshal | 399.14μs | 225.0K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.60μs | 260.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 301.96μs | 402.6K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 521.14μs | 346.3K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 659.55μs | 320.5K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.02ms | 545.9K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 580ns | 290 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 628ns | 496 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 993ns | 1.1K | 3 |
| Small Struct | 🥇 BEVE | Marshal | 1.45μs | 1.7K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.95μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.83μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.00μs | 728 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.56μs | 2.4K | 8 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.76μs | 1.3K | 31 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.87μs | 1.9K | 42 |
| Small Struct | 🥉 JSON | Unmarshal | 8.99μs | 2.1K | 36 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.13μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 15.19μs | 21.9K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 20.23μs | 25.2K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 22.28μs | 18.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.41μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.24μs | 20.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.48μs | 29.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.45μs | 56.1K | 71 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.49μs | 29.0K | 524 |
| Medium Payload | 🥈 CBOR | Unmarshal | 87.84μs | 31.0K | 642 |
| Medium Payload | 🥉 JSON | Unmarshal | 282.69μs | 61.1K | 784 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 75.26μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 117.58μs | 180.8K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 171.86μs | 237.0K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 227.70μs | 205.9K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 306.95μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 526.22μs | 224.2K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 294.38μs | 285.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 420.25μs | 536.6K | 566 |
| Large Payload | 🥈 MessagePack | Unmarshal | 666.41μs | 342.4K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 853.15μs | 313.7K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.58ms | 532.1K | 6.8K |

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

