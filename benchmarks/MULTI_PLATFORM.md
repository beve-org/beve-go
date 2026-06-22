# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-06-22 08:53:27 UTC

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
| Apple M1 (Virtual) | 950ns | 243ns | 3.39μs | 402ns | 2.61μs |
| AMD EPYC 7763 64-Core Processor | 996ns | 404ns | 3.40μs | 541ns | 1.20μs |
| Neoverse-N2 | 1.03μs | 704ns | 4.37μs | 625ns | 2.57μs |
| Unknown CPU | 2.12μs | 452ns | 4.27μs | 1.58μs | 3.70μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 853ns | 15.94μs | 2.94μs | 2.19μs |
| AMD EPYC 7763 64-Core Processor | 916ns | 16.30μs | 9.78μs | 5.76μs |
| Neoverse-N2 | 898ns | 8.46μs | 5.94μs | 2.99μs |
| Unknown CPU | 2.74μs | 14.26μs | 7.62μs | 6.56μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (243ns) | 🥇 BEVE (853ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (404ns) | 🥇 BEVE (916ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥈 CBOR (625ns) | 🥇 BEVE (898ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (452ns) | 🥇 BEVE (2.74μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 67.4% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 243ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 402ns | 384 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 950ns | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.56μs | 1.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.61μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 3.39μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 853ns | 1.3K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.19μs | 2.3K | 49 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.28μs | 3.2K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.94μs | 2.2K | 48 |
| Small Struct | 🥉 JSON | Unmarshal | 15.94μs | 4.3K | 70 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.78μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.74μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 13.95μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 23.95μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 35.58μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 40.55μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.42μs | 28.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.48μs | 32.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 41.57μs | 38.0K | 709 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.75μs | 32.0K | 656 |
| Medium Payload | 🥉 JSON | Unmarshal | 172.97μs | 51.5K | 714 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 58.26μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 81.07μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 162.31μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 180.90μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 341.08μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 471.74μs | 213.7K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 185.89μs | 259.2K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 333.54μs | 363.0K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 402.67μs | 352.6K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 461.03μs | 290.7K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.45ms | 554.2K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 404ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 541ns | 352 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 996ns | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.20μs | 1.0K | 6 |
| Small Struct | 🥉 Sonic | Marshal | 1.70μs | 2.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.40μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 916ns | 952 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.58μs | 2.0K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.76μs | 3.9K | 82 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.78μs | 4.3K | 91 |
| Small Struct | 🥉 JSON | Unmarshal | 16.30μs | 4.0K | 63 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 10.37μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.70μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.70μs | 22.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.62μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.20μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 50.32μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.75μs | 23.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.30μs | 63.9K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 61.64μs | 39.2K | 730 |
| Medium Payload | 🥈 CBOR | Unmarshal | 72.93μs | 33.5K | 688 |
| Medium Payload | 🥉 JSON | Unmarshal | 207.68μs | 44.5K | 588 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.24μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 128.98μs | 204.9K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 162.96μs | 215.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 215.15μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 348.11μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 423.79μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 257.42μs | 269.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 410.65μs | 580.1K | 602 |
| Large Payload | 🥈 MessagePack | Unmarshal | 583.53μs | 356.1K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 683.21μs | 290.3K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.49ms | 541.4K | 7.1K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥈 CBOR | Marshal | 625ns | 416 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 704ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.03μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.15μs | 1.5K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.57μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.37μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 898ns | 728 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.86μs | 4.5K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.99μs | 2.1K | 46 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.94μs | 3.5K | 75 |
| Small Struct | 🥉 JSON | Unmarshal | 8.46μs | 2.2K | 41 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.90μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.83μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.89μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 29.74μs | 22.1K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.60μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 38.78μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.46μs | 27.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.57μs | 52.4K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.99μs | 39.2K | 732 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.06μs | 29.6K | 608 |
| Medium Payload | 🥉 JSON | Unmarshal | 209.73μs | 58.6K | 780 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.45μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 104.91μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 185.01μs | 188.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 286.92μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 316.29μs | 225.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 393.47μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.29μs | 268.1K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 291.42μs | 401.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 523.47μs | 354.7K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 642.61μs | 308.3K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.95ms | 521.1K | 6.9K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 452ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.58μs | 1.3K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 2.12μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.18μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.70μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.27μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.74μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.33μs | 7.7K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.56μs | 4.4K | 92 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.62μs | 3.6K | 76 |
| Small Struct | 🥉 JSON | Unmarshal | 14.26μs | 3.7K | 53 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.65μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.94μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.67μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.33μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.60μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 43.08μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.96μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.79μs | 57.5K | 71 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 67.30μs | 35.2K | 652 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.13μs | 23.9K | 496 |
| Medium Payload | 🥉 JSON | Unmarshal | 261.97μs | 56.0K | 738 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.54μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.98μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 166.19μs | 222.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 207.86μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 291.14μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 511.52μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 270.58μs | 265.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 415.54μs | 526.1K | 575 |
| Large Payload | 🥈 MessagePack | Unmarshal | 682.99μs | 359.6K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 846.40μs | 317.0K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.52ms | 524.5K | 6.8K |

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

