# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-05-11 06:33:36 UTC

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
| Apple M1 (Virtual) | 787ns | 497ns | 2.18μs | 1.60μs | 4.29μs |
| AMD EPYC 7763 64-Core Processor | 1.11μs | 821ns | 4.70μs | 2.82μs | 2.92μs |
| Neoverse-N2 | 1.34μs | 710ns | 734ns | 1.90μs | 765ns |
| Unknown CPU | 336ns | 766ns | 3.04μs | 1.30μs | 814ns |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 886ns | 18.64μs | 837ns | 1.97μs |
| AMD EPYC 7763 64-Core Processor | 806ns | 13.35μs | 6.11μs | 4.91μs |
| Neoverse-N2 | 1.66μs | 22.60μs | 4.37μs | 3.92μs |
| Unknown CPU | 1.35μs | 20.77μs | 5.58μs | 7.65μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (497ns) | 🥈 CBOR (837ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (821ns) | 🥇 BEVE (806ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (710ns) | 🥇 BEVE (1.66μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (336ns) | 🥇 BEVE (1.35μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 36.5% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 497ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 787ns | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.60μs | 2.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.18μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.02μs | 1.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.29μs | 8.2K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 837ns | 232 | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 886ns | 1.3K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.97μs | 1.3K | 30 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.54μs | 3.7K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 18.64μs | 7.5K | 100 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.97μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.77μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.70μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.89μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 28.46μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 40.38μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.71μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.59μs | 38.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 44.66μs | 35.9K | 664 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.00μs | 30.4K | 624 |
| Medium Payload | 🥉 JSON | Unmarshal | 168.22μs | 50.8K | 659 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 57.53μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.12μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 154.66μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 337.50μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 394.35μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 498.20μs | 214.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 228.14μs | 280.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 343.40μs | 347.6K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 408.56μs | 331.1K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 602.19μs | 320.7K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.04ms | 549.1K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 821ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.11μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.22μs | 1.6K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.82μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.92μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.70μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 806ns | 600 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.87μs | 2.4K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.91μs | 3.6K | 78 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.11μs | 3.3K | 71 |
| Small Struct | 🥉 JSON | Unmarshal | 13.35μs | 3.7K | 52 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.31μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.05μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.07μs | 19.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.83μs | 20.5K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 36.42μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.49μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.74μs | 31.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 35.63μs | 49.7K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 45.58μs | 26.7K | 477 |
| Medium Payload | 🥈 CBOR | Unmarshal | 80.66μs | 38.9K | 794 |
| Medium Payload | 🥉 JSON | Unmarshal | 280.23μs | 65.1K | 843 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.05μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.14μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 169.41μs | 224.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 206.52μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 331.20μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 483.68μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.52μs | 269.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 371.76μs | 556.6K | 583 |
| Large Payload | 🥈 MessagePack | Unmarshal | 572.62μs | 349.2K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 662.95μs | 285.8K | 5.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.56ms | 559.2K | 7.3K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 710ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 734ns | 288 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 765ns | 520 | 5 |
| Small Struct | 🥇 BEVE | Marshal | 1.34μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.74μs | 1.2K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.90μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.66μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.89μs | 2.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.92μs | 3.2K | 67 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.37μs | 2.4K | 53 |
| Small Struct | 🥉 JSON | Unmarshal | 22.60μs | 7.6K | 103 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.01μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.37μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.14μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.53μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 33.91μs | 25.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 43.58μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.85μs | 30.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.17μs | 45.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.36μs | 37.9K | 706 |
| Medium Payload | 🥈 CBOR | Unmarshal | 59.80μs | 28.7K | 588 |
| Medium Payload | 🥉 JSON | Unmarshal | 199.96μs | 58.4K | 732 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.46μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.30μs | 204.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 176.41μs | 180.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 282.05μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 313.01μs | 222.8K | 3 |
| Large Payload | 🥉 JSON | Marshal | 409.81μs | 229.8K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.55μs | 275.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 285.97μs | 382.3K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 499.17μs | 335.0K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 678.44μs | 332.7K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.86ms | 494.5K | 6.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 336ns | 192 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 766ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 814ns | 520 | 5 |
| Small Struct | 🥈 CBOR | Marshal | 1.30μs | 1.2K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.34μs | 2.3K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.04μs | 1.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.35μs | 1.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.44μs | 3.6K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.58μs | 2.8K | 60 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.65μs | 5.1K | 105 |
| Small Struct | 🥉 JSON | Unmarshal | 20.77μs | 4.5K | 77 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.95μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.91μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.50μs | 19.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.44μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.58μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.19μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.74μs | 26.3K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 54.04μs | 60.3K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 64.41μs | 33.3K | 613 |
| Medium Payload | 🥈 CBOR | Unmarshal | 79.74μs | 30.5K | 631 |
| Medium Payload | 🥉 JSON | Unmarshal | 240.57μs | 51.7K | 674 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 72.58μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.25μs | 188.4K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 161.78μs | 223.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 200.81μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 282.20μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 511.07μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 267.10μs | 262.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 462.06μs | 523.9K | 570 |
| Large Payload | 🥈 MessagePack | Unmarshal | 634.57μs | 316.8K | 5.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 824.07μs | 312.7K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.37ms | 509.8K | 6.5K |

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

