# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 18:19:47 UTC

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
| Apple M1 (Virtual) | 1.81μs | 537ns | 856ns | 2.06μs | 2.61μs |
| AMD EPYC 7763 64-Core Processor | 1.01μs | 324ns | 3.96μs | 1.75μs | 2.64μs |
| Neoverse-N2 | 834ns | 609ns | 2.57μs | 2.31μs | 703ns |
| Unknown CPU | 1.18μs | 590ns | 5.38μs | 1.19μs | 3.31μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 523ns | 30.38μs | 3.14μs | 8.28μs |
| AMD EPYC 7763 64-Core Processor | 1.29μs | 23.81μs | 6.55μs | 3.94μs |
| Neoverse-N2 | 802ns | 26.02μs | 2.83μs | 1.36μs |
| Unknown CPU | 1.94μs | 7.17μs | 6.62μs | 2.01μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (537ns) | 🥇 BEVE (523ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (324ns) | 🥇 BEVE (1.29μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (609ns) | 🥇 BEVE (802ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (590ns) | 🥇 BEVE (1.94μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 27.2% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 537ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 856ns | 448 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.81μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.06μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.61μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 4.13μs | 1.9K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 523ns | 312 | 3 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.14μs | 2.1K | 47 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.17μs | 4.1K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 8.28μs | 4.2K | 88 |
| Small Struct | 🥉 JSON | Unmarshal | 30.38μs | 7.9K | 112 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.21μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.04μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.16μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.64μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 49.12μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 54.76μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.59μs | 29.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.51μs | 25.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.01μs | 37.1K | 691 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.60μs | 30.7K | 631 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.62μs | 52.9K | 745 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 63.79μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 101.21μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 195.68μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 339.81μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 442.24μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 556.67μs | 214.6K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 209.52μs | 276.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 342.31μs | 379.9K | 213 |
| Large Payload | 🥈 CBOR | Unmarshal | 531.41μs | 291.3K | 5.9K |
| Large Payload | 🥈 MessagePack | Unmarshal | 683.29μs | 359.1K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 517.0K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 324ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.01μs | 1.5K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.75μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.97μs | 3.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.64μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.96μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.29μs | 1.7K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.94μs | 2.8K | 60 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.33μs | 7.7K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.55μs | 3.2K | 69 |
| Small Struct | 🥉 JSON | Unmarshal | 23.81μs | 7.5K | 99 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.97μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.41μs | 27.3K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.94μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.83μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.50μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.02μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.45μs | 21.1K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.75μs | 60.8K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 64.15μs | 42.3K | 793 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.10μs | 31.1K | 639 |
| Medium Payload | 🥉 JSON | Unmarshal | 229.24μs | 58.0K | 741 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 76.01μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.93μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 151.91μs | 207.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 197.86μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 308.33μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 454.13μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 228.40μs | 264.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 388.37μs | 585.6K | 593 |
| Large Payload | 🥈 MessagePack | Unmarshal | 546.03μs | 329.5K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 754.08μs | 301.2K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.42ms | 570.4K | 7.5K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 609ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 703ns | 520 | 5 |
| Small Struct | 🥇 BEVE | Marshal | 834ns | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.31μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.43μs | 1.9K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.57μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 802ns | 728 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.36μs | 496 | 13 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.83μs | 1.4K | 31 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.88μs | 5.0K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 26.02μs | 8.1K | 118 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.46μs | 10 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.43μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.73μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 31.52μs | 25.0K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.90μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.55μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.21μs | 27.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.04μs | 39.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 51.05μs | 35.6K | 660 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.10μs | 34.2K | 702 |
| Medium Payload | 🥉 JSON | Unmarshal | 173.27μs | 47.4K | 615 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.84μs | 105 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 105.46μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 201.77μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 283.71μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 306.64μs | 217.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 384.78μs | 213.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 221.08μs | 271.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 285.22μs | 392.0K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 523.33μs | 358.2K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 669.31μs | 333.1K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.95ms | 521.3K | 6.9K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 590ns | 1 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.18μs | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.19μs | 1.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.82μs | 2.4K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.31μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.38μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.94μs | 2.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.01μs | 688 | 17 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.14μs | 2.2K | 8 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.62μs | 2.7K | 56 |
| Small Struct | 🥉 JSON | Unmarshal | 7.17μs | 1.4K | 29 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.27μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.88μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.27μs | 27.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.55μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.44μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 44.11μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.25μs | 24.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 53.50μs | 64.8K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 72.98μs | 38.4K | 713 |
| Medium Payload | 🥈 CBOR | Unmarshal | 79.75μs | 30.1K | 621 |
| Medium Payload | 🥉 JSON | Unmarshal | 231.61μs | 49.0K | 624 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.75μs | 105 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 109.19μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 161.17μs | 215.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 224.33μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 285.23μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 489.33μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 336.76μs | 261.2K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 439.97μs | 570.8K | 592 |
| Large Payload | 🥈 MessagePack | Unmarshal | 675.06μs | 349.4K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 890.51μs | 340.3K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.65ms | 553.6K | 7.2K |

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

