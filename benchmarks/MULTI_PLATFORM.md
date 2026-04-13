# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-04-13 05:37:44 UTC

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
| Apple M1 (Virtual) | 264ns | 494ns | 796ns | 1.37μs | 1.48μs |
| AMD EPYC 7763 64-Core Processor | 1.52μs | 525ns | 3.69μs | 2.77μs | 2.82μs |
| Neoverse-N2 | 886ns | 333ns | 1.68μs | 853ns | 2.31μs |
| Unknown CPU | 437ns | 679ns | 1.50μs | 1.91μs | 1.16μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.24μs | 12.94μs | 2.26μs | 3.37μs |
| AMD EPYC 7763 64-Core Processor | 2.07μs | 5.87μs | 1.16μs | 5.63μs |
| Neoverse-N2 | 1.68μs | 12.39μs | 2.34μs | 1.10μs |
| Unknown CPU | 762ns | 11.47μs | 7.34μs | 4.53μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (264ns) | 🥇 BEVE (1.24μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (525ns) | 🥈 CBOR (1.16μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (333ns) | 🥈 MessagePack (1.10μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (437ns) | 🥇 BEVE (762ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 60.9% faster

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
| Small Struct | 🥇 BEVE | Marshal | 264ns | 448 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 494ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 796ns | 512 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.37μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.48μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 1.50μs | 798 | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.24μs | 3.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.26μs | 1.6K | 36 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.82μs | 5.0K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.37μs | 4.4K | 94 |
| Small Struct | 🥉 JSON | Unmarshal | 12.94μs | 4.2K | 68 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.38μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.48μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 14.87μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 18.03μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 29.19μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 43.27μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 18.48μs | 25.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.18μs | 42.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 40.69μs | 36.4K | 674 |
| Medium Payload | 🥈 CBOR | Unmarshal | 49.16μs | 30.5K | 628 |
| Medium Payload | 🥉 JSON | Unmarshal | 159.14μs | 50.4K | 667 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 59.26μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 106.65μs | 196.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 176.55μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 333.79μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 348.02μs | 196.9K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 435.19μs | 213.9K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 265.92μs | 264.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 338.89μs | 347.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 476.26μs | 354.0K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 626.49μs | 311.5K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 535.9K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 525ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 891ns | 1.1K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.52μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.77μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.82μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.69μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.16μs | 224 | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.07μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.50μs | 3.8K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.63μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 5.87μs | 1.3K | 26 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.84μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.39μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.84μs | 18.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.18μs | 28.1K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.02μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 41.74μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.62μs | 25.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 48.33μs | 67.4K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.68μs | 33.9K | 623 |
| Medium Payload | 🥈 CBOR | Unmarshal | 72.63μs | 33.6K | 689 |
| Medium Payload | 🥉 JSON | Unmarshal | 210.07μs | 48.3K | 662 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.13μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 118.41μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 159.18μs | 224.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 235.41μs | 213.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 329.06μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 432.00μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 248.18μs | 262.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 372.46μs | 525.5K | 572 |
| Large Payload | 🥈 MessagePack | Unmarshal | 576.41μs | 344.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 722.09μs | 312.7K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.32ms | 541.5K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 333ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 853ns | 768 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 886ns | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.68μs | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.31μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.41μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.10μs | 304 | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.68μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.12μs | 2.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.34μs | 1.0K | 24 |
| Small Struct | 🥉 JSON | Unmarshal | 12.39μs | 3.9K | 58 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.59μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.77μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.03μs | 19.1K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 25.06μs | 18.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.70μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.81μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.48μs | 30.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.73μs | 44.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.57μs | 36.8K | 682 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.83μs | 30.4K | 627 |
| Medium Payload | 🥉 JSON | Unmarshal | 188.44μs | 50.5K | 684 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.68μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 110.05μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 186.46μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 275.19μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 299.82μs | 215.6K | 3 |
| Large Payload | 🥉 JSON | Marshal | 406.41μs | 221.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.71μs | 275.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 292.05μs | 399.5K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 521.10μs | 355.9K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 645.63μs | 310.6K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.03ms | 552.9K | 7.2K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 437ns | 384 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 679ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.16μs | 1.0K | 6 |
| Small Struct | 🥉 JSON | Marshal | 1.50μs | 576 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.51μs | 1.4K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.91μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 762ns | 376 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.53μs | 2.8K | 58 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.51μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.34μs | 3.4K | 72 |
| Small Struct | 🥉 JSON | Unmarshal | 11.47μs | 2.3K | 45 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.30μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.86μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.53μs | 20.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.28μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.10μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 50.89μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.01μs | 29.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.38μs | 50.9K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 76.82μs | 43.1K | 812 |
| Medium Payload | 🥈 CBOR | Unmarshal | 98.66μs | 39.0K | 804 |
| Medium Payload | 🥉 JSON | Unmarshal | 247.07μs | 53.5K | 685 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.36μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.66μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 164.43μs | 222.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 217.64μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 289.68μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 465.96μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 267.14μs | 259.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 443.99μs | 566.5K | 594 |
| Large Payload | 🥈 MessagePack | Unmarshal | 659.27μs | 340.2K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 813.21μs | 298.4K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.44ms | 495.3K | 6.5K |

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

