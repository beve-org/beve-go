# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-01-26 04:17:19 UTC

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
| Apple M1 (Virtual) | 1.53μs | 352ns | 3.12μs | 683ns | 1.75μs |
| AMD EPYC 7763 64-Core Processor | 1.09μs | 587ns | 988ns | 1.40μs | 4.71μs |
| Neoverse-N2 | 1.02μs | 299ns | 1.46μs | 2.35μs | 3.75μs |
| Unknown CPU | 1.62μs | 248ns | 1.61μs | 551ns | 1.07μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 845ns | 9.19μs | 2.94μs | 2.43μs |
| AMD EPYC 7763 64-Core Processor | 1.75μs | 21.67μs | 7.80μs | 1.16μs |
| Neoverse-N2 | 678ns | 22.36μs | 6.50μs | 5.11μs |
| Unknown CPU | 820ns | 27.78μs | 4.04μs | 7.01μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (352ns) | 🥇 BEVE (845ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (587ns) | 🥈 MessagePack (1.16μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (299ns) | 🥇 BEVE (678ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (248ns) | 🥇 BEVE (820ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 17.2% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 352ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 683ns | 320 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.53μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.75μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.78μs | 432 | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.12μs | 1.4K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 845ns | 408 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.42μs | 1.2K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.43μs | 1.2K | 27 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.94μs | 1.1K | 26 |
| Small Struct | 🥉 JSON | Unmarshal | 9.19μs | 2.2K | 42 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.29μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 16.65μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 26.79μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.32μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 52.93μs | 19.4K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 58.47μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.16μs | 28.3K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 42.98μs | 22.4K | 388 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.85μs | 42.8K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.38μs | 34.0K | 695 |
| Medium Payload | 🥉 JSON | Unmarshal | 231.53μs | 47.9K | 642 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.43μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 129.04μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 154.57μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 429.01μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 489.73μs | 205.1K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 620.07μs | 206.3K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 316.74μs | 294.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 365.71μs | 366.9K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 402.14μs | 342.3K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 731.17μs | 332.2K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.53ms | 569.2K | 7.5K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 587ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 988ns | 384 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.09μs | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.40μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.23μs | 3.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.71μs | 8.2K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.16μs | 304 | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.75μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.13μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.80μs | 4.4K | 94 |
| Small Struct | 🥉 JSON | Unmarshal | 21.67μs | 4.7K | 83 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.57μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.42μs | 19.1K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.71μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.10μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.97μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.92μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.76μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.27μs | 50.8K | 69 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.36μs | 35.2K | 649 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.24μs | 27.8K | 575 |
| Medium Payload | 🥉 JSON | Unmarshal | 194.19μs | 43.8K | 558 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.55μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 118.77μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 164.73μs | 224.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 212.03μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 324.03μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 440.64μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 240.91μs | 260.3K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 363.40μs | 525.0K | 570 |
| Large Payload | 🥈 MessagePack | Unmarshal | 596.65μs | 373.9K | 6.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 694.65μs | 302.5K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.43ms | 539.4K | 6.9K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 299ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.02μs | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.46μs | 704 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.70μs | 1.1K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.35μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.75μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 678ns | 376 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.12μs | 3.1K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.11μs | 4.4K | 92 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.50μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 22.36μs | 7.6K | 102 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.78μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.40μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.52μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.52μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 28.62μs | 21.0K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 42.71μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.79μs | 24.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.77μs | 40.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 51.11μs | 34.2K | 633 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.57μs | 30.8K | 629 |
| Medium Payload | 🥉 JSON | Unmarshal | 188.91μs | 52.7K | 690 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.86μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 109.36μs | 180.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 200.73μs | 205.2K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 274.88μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 298.14μs | 201.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 406.43μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 232.31μs | 273.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 291.46μs | 395.9K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 520.50μs | 353.8K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 665.60μs | 319.7K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.10ms | 565.2K | 7.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 248ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 551ns | 352 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.06μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.07μs | 1.0K | 6 |
| Small Struct | 🥉 JSON | Marshal | 1.61μs | 704 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.62μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 820ns | 472 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.28μs | 3.6K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.04μs | 1.7K | 39 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.01μs | 4.3K | 89 |
| Small Struct | 🥉 JSON | Unmarshal | 27.78μs | 7.3K | 94 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.84μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.67μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.45μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.09μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.14μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 54.20μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.93μs | 28.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.63μs | 48.8K | 71 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 77.68μs | 41.3K | 774 |
| Medium Payload | 🥈 CBOR | Unmarshal | 100.83μs | 38.6K | 791 |
| Medium Payload | 🥉 JSON | Unmarshal | 248.18μs | 46.3K | 646 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.14μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.33μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 152.95μs | 198.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 210.98μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 323.29μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 484.11μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 292.10μs | 270.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 441.13μs | 538.9K | 572 |
| Large Payload | 🥈 MessagePack | Unmarshal | 664.05μs | 341.0K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 899.43μs | 329.9K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.79ms | 571.5K | 7.5K |

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

