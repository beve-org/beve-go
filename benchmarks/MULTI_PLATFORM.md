# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-12-01 04:12:03 UTC

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
| Apple M1 (Virtual) | 925ns | 329ns | 1.10μs | 1.37μs | 1.49μs |
| AMD EPYC 7763 64-Core Processor | 1.14μs | 665ns | 1.03μs | 2.41μs | 2.63μs |
| Neoverse-N2 | 842ns | 394ns | 3.46μs | 791ns | 4.01μs |
| Unknown CPU | 1.51μs | 447ns | 6.49μs | 806ns | 4.84μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 953ns | 17.02μs | 4.26μs | 2.02μs |
| AMD EPYC 7763 64-Core Processor | 1.69μs | 4.84μs | 1.40μs | 3.67μs |
| Neoverse-N2 | 1.26μs | 25.76μs | 3.74μs | 5.58μs |
| Unknown CPU | 1.38μs | 3.75μs | 3.41μs | 7.53μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (329ns) | 🥇 BEVE (953ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (665ns) | 🥈 CBOR (1.40μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (394ns) | 🥉 Sonic (1.20μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (447ns) | 🥇 BEVE (1.38μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 39.3% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 329ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 925ns | 1.5K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.10μs | 768 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.37μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.49μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 2.42μs | 1.4K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 953ns | 2.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.02μs | 2.1K | 46 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.49μs | 4.0K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.26μs | 3.7K | 79 |
| Small Struct | 🥉 JSON | Unmarshal | 17.02μs | 7.3K | 93 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.30μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.44μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 20.84μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.40μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.66μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 41.57μs | 20.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.17μs | 26.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.00μs | 36.4K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 38.48μs | 28.1K | 509 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.89μs | 30.6K | 631 |
| Medium Payload | 🥉 JSON | Unmarshal | 206.46μs | 49.6K | 644 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 74.27μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 156.84μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 217.02μs | 526.8K | 115 |
| Large Payload | 🥈 CBOR | Marshal | 227.98μs | 196.7K | 1 |
| Large Payload | 🥉 JSON | Marshal | 404.67μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 422.70μs | 213.7K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 186.30μs | 266.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 450.65μs | 356.1K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 528.44μs | 356.6K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 777.06μs | 309.7K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.98ms | 528.1K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 665ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 1.03μs | 416 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.14μs | 1.5K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.58μs | 2.3K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.41μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.63μs | 4.1K | 8 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.40μs | 368 | 11 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.69μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.46μs | 3.8K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.67μs | 2.5K | 54 |
| Small Struct | 🥉 JSON | Unmarshal | 4.84μs | 904 | 22 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.21μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 17.47μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.80μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.54μs | 27.7K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.61μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 36.68μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.95μs | 25.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 35.12μs | 51.4K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.47μs | 35.5K | 655 |
| Medium Payload | 🥈 CBOR | Unmarshal | 74.85μs | 32.0K | 661 |
| Medium Payload | 🥉 JSON | Unmarshal | 223.82μs | 55.0K | 719 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 74.51μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.62μs | 204.8K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 145.53μs | 207.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 207.41μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 316.12μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 468.51μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.54μs | 280.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 355.69μs | 557.3K | 592 |
| Large Payload | 🥈 MessagePack | Unmarshal | 577.95μs | 361.8K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 701.89μs | 314.8K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.50ms | 605.8K | 7.9K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 394ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 791ns | 576 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 842ns | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.20μs | 2.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.46μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.01μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.20μs | 1.2K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.26μs | 1.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.74μs | 2.0K | 44 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.58μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 25.76μs | 8.0K | 117 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.36μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.25μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.49μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.09μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 35.18μs | 27.9K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 39.56μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.48μs | 26.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.26μs | 30.8K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 48.29μs | 31.9K | 580 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.34μs | 33.4K | 684 |
| Medium Payload | 🥉 JSON | Unmarshal | 148.52μs | 39.0K | 505 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.23μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.78μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 195.25μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 275.09μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 280.48μs | 198.3K | 3 |
| Large Payload | 🥉 JSON | Marshal | 385.39μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.06μs | 270.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 298.19μs | 400.4K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 528.18μs | 351.3K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 628.74μs | 294.0K | 6.0K |
| Large Payload | 🥉 JSON | Unmarshal | 1.92ms | 512.3K | 6.7K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 447ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 806ns | 576 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.51μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.81μs | 2.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.84μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 6.49μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.38μs | 1.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.02μs | 2.0K | 8 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.41μs | 1.2K | 27 |
| Small Struct | 🥉 JSON | Unmarshal | 3.75μs | 552 | 15 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.53μs | 4.6K | 96 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.77μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.54μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.78μs | 20.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.91μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.63μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 49.24μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 32.43μs | 25.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.65μs | 55.4K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 79.82μs | 38.7K | 721 |
| Medium Payload | 🥈 CBOR | Unmarshal | 94.93μs | 34.5K | 706 |
| Medium Payload | 🥉 JSON | Unmarshal | 314.05μs | 61.8K | 833 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 84.81μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 132.63μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 185.21μs | 206.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 270.60μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 328.32μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 510.29μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 309.05μs | 273.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 444.16μs | 521.1K | 574 |
| Large Payload | 🥈 CBOR | Unmarshal | 858.29μs | 315.3K | 6.4K |
| Large Payload | 🥈 MessagePack | Unmarshal | 917.41μs | 339.7K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.35ms | 469.4K | 6.2K |

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

