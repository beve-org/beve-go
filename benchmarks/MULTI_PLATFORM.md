# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 15:27:09 UTC

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
| Apple M1 (Virtual) | 2.15μs | 507ns | 1.49μs | 364ns | 1.90μs |
| AMD EPYC 7763 64-Core Processor | 1.31μs | 476ns | 2.00μs | 2.55μs | 1.09μs |
| Neoverse-N2 | 1.69μs | 423ns | 3.43μs | 506ns | 3.87μs |
| Unknown CPU | 2.28μs | 986ns | 5.83μs | 1.25μs | 5.66μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.92μs | 26.64μs | 6.40μs | 852ns |
| AMD EPYC 7763 64-Core Processor | 1.55μs | 8.20μs | 8.96μs | 4.65μs |
| Neoverse-N2 | 1.72μs | 2.55μs | 4.88μs | 5.83μs |
| Unknown CPU | 2.21μs | 32.18μs | 9.43μs | 1.60μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥈 CBOR (364ns) | 🥈 MessagePack (852ns) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (476ns) | 🥇 BEVE (1.55μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (423ns) | 🥇 BEVE (1.72μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (986ns) | 🥈 MessagePack (1.60μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 25.3% faster

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
| Small Struct | 🥈 CBOR | Marshal | 364ns | 352 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 507ns | 290 | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.49μs | 1.0K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.90μs | 2.2K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 2.15μs | 3.0K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 3.95μs | 2.2K | 3 |
| Small Struct | 🥈 MessagePack | Unmarshal | 852ns | 352 | 10 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.92μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.90μs | 4.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.40μs | 4.7K | 100 |
| Small Struct | 🥉 JSON | Unmarshal | 26.64μs | 7.8K | 110 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.24μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 12.91μs | 20.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.87μs | 20.6K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 44.21μs | 19.4K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 46.83μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 67.39μs | 27.7K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 34.74μs | 31.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 45.38μs | 36.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.42μs | 32.2K | 590 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.73μs | 33.7K | 692 |
| Medium Payload | 🥉 JSON | Unmarshal | 209.77μs | 44.5K | 584 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 63.80μs | 180 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 128.78μs | 189.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 178.24μs | 197.4K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 288.06μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 443.97μs | 214.2K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 524.89μs | 223.3K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 257.08μs | 283.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 337.24μs | 351.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 527.95μs | 353.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 766.72μs | 317.4K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.37ms | 542.7K | 7.1K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 476ns | 289 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 650ns | 686 | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.09μs | 1.2K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 1.31μs | 2.1K | 3 |
| Small Struct | 🥉 JSON | Marshal | 2.00μs | 1.0K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.55μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.55μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.11μs | 7.4K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.65μs | 3.5K | 72 |
| Small Struct | 🥉 JSON | Unmarshal | 8.20μs | 2.0K | 36 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.96μs | 4.7K | 101 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.58μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 15.87μs | 24.8K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 18.77μs | 25.6K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 21.79μs | 20.6K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 36.65μs | 16.7K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.17μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.61μs | 21.1K | 57 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.24μs | 65.6K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.52μs | 41.1K | 769 |
| Medium Payload | 🥈 CBOR | Unmarshal | 80.12μs | 37.1K | 761 |
| Medium Payload | 🥉 JSON | Unmarshal | 223.33μs | 50.9K | 687 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.13μs | 286 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 118.94μs | 188.7K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 168.00μs | 216.9K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 222.33μs | 213.8K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 304.23μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 436.80μs | 213.6K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.10μs | 265.8K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 360.64μs | 551.4K | 577 |
| Large Payload | 🥈 MessagePack | Unmarshal | 582.47μs | 359.9K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 727.62μs | 317.2K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.35ms | 541.9K | 7.1K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 423ns | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 506ns | 352 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.69μs | 3.0K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 1.70μs | 1.2K | 3 |
| Small Struct | 🥉 JSON | Marshal | 3.43μs | 2.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.87μs | 8.3K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.72μs | 3.4K | 4 |
| Small Struct | 🥉 JSON | Unmarshal | 2.55μs | 464 | 12 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.22μs | 5.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.88μs | 2.7K | 57 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.83μs | 5.2K | 105 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.13μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 11.27μs | 20.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 17.75μs | 19.2K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 30.10μs | 16.7K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.80μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 33.93μs | 25.1K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.16μs | 29.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.04μs | 33.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.29μs | 27.5K | 496 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.04μs | 30.1K | 620 |
| Medium Payload | 🥉 JSON | Unmarshal | 194.64μs | 53.0K | 707 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.90μs | 286 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 116.10μs | 206.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 188.70μs | 199.8K | 3 |
| Large Payload | 🥈 MessagePack | Marshal | 265.82μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 292.77μs | 219.7K | 4 |
| Large Payload | 🥉 JSON | Marshal | 382.22μs | 215.0K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 221.94μs | 272.0K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 278.50μs | 388.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 499.00μs | 337.2K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 653.25μs | 316.4K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.90ms | 504.4K | 6.6K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 986ns | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.25μs | 1.2K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.79μs | 2.3K | 3 |
| Small Struct | 🥇 BEVE | Marshal | 2.28μs | 2.6K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 5.66μs | 8.3K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.83μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.60μs | 544 | 14 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.96μs | 2.0K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.21μs | 3.0K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.43μs | 3.8K | 80 |
| Small Struct | 🥉 JSON | Unmarshal | 32.18μs | 8.0K | 116 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.87μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 17.84μs | 27.4K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 18.55μs | 25.2K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 19.47μs | 18.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.74μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.84μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.08μs | 28.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.66μs | 47.5K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 68.82μs | 36.6K | 679 |
| Medium Payload | 🥈 CBOR | Unmarshal | 95.14μs | 35.5K | 734 |
| Medium Payload | 🥉 JSON | Unmarshal | 251.82μs | 53.4K | 715 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 85.48μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 126.75μs | 197.2K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 157.11μs | 220.3K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 215.10μs | 189.6K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 294.89μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 486.93μs | 206.7K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 280.00μs | 279.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 436.27μs | 565.6K | 598 |
| Large Payload | 🥈 MessagePack | Unmarshal | 698.02μs | 338.8K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 865.81μs | 312.3K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.66ms | 566.9K | 7.3K |

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

