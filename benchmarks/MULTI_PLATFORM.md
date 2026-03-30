# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-03-30 05:21:37 UTC

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
| Apple M1 (Virtual) | 779ns | 390ns | 5.52μs | 1.87μs | 3.64μs |
| AMD EPYC 7763 64-Core Processor | 1.62μs | 543ns | 2.21μs | 1.25μs | 2.42μs |
| Neoverse-N2 | 1.10μs | 696ns | 2.36μs | 1.78μs | 3.81μs |
| Unknown CPU | 1.12μs | 357ns | 2.55μs | 1.67μs | 1.08μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.25μs | 28.95μs | 3.60μs | 4.40μs |
| AMD EPYC 7763 64-Core Processor | 706ns | 25.02μs | 5.54μs | 1.91μs |
| Neoverse-N2 | 1.51μs | 25.23μs | 6.63μs | 4.50μs |
| Unknown CPU | 2.25μs | 33.23μs | 5.95μs | 1.05μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (390ns) | 🥇 BEVE (1.25μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (543ns) | 🥇 BEVE (706ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (696ns) | 🥇 BEVE (1.51μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (357ns) | 🥈 MessagePack (1.05μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 55.6% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 390ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 779ns | 640 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.87μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.64μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.52μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 6.52μs | 3.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.25μs | 1.7K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.34μs | 2.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.60μs | 2.0K | 45 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.40μs | 3.1K | 66 |
| Small Struct | 🥉 JSON | Unmarshal | 28.95μs | 8.0K | 117 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 11.62μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.85μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 13.36μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.92μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 45.27μs | 19.3K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 52.49μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.92μs | 28.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.67μs | 37.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 36.87μs | 32.7K | 603 |
| Medium Payload | 🥈 CBOR | Unmarshal | 53.61μs | 33.8K | 697 |
| Medium Payload | 🥉 JSON | Unmarshal | 175.05μs | 51.8K | 662 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 54.62μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.44μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 134.46μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 189.67μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 401.23μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 448.03μs | 213.7K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 256.09μs | 280.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 304.93μs | 337.5K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 556.98μs | 331.6K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 635.89μs | 326.0K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.36ms | 527.5K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 543ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.23μs | 1.6K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.25μs | 1.2K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.62μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.21μs | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.42μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 706ns | 440 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.74μs | 2.2K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.91μs | 976 | 23 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.54μs | 3.1K | 65 |
| Small Struct | 🥉 JSON | Unmarshal | 25.02μs | 7.6K | 102 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.07μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.07μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.21μs | 20.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.38μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.21μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.74μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.30μs | 29.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.84μs | 60.5K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.77μs | 34.9K | 642 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.37μs | 29.5K | 610 |
| Medium Payload | 🥉 JSON | Unmarshal | 268.15μs | 68.8K | 876 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.17μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.74μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 170.34μs | 240.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 203.03μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 325.35μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 431.25μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 246.83μs | 260.0K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 377.49μs | 544.9K | 571 |
| Large Payload | 🥈 MessagePack | Unmarshal | 559.52μs | 332.2K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 769.45μs | 347.3K | 7.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.23ms | 528.0K | 6.8K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 696ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.10μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.15μs | 677 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.78μs | 2.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.36μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.81μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.51μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.14μs | 3.3K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.50μs | 3.7K | 79 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.63μs | 4.0K | 85 |
| Small Struct | 🥉 JSON | Unmarshal | 25.23μs | 7.9K | 114 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.76μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.06μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.32μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.99μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 30.80μs | 22.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 39.30μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.79μs | 30.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.45μs | 51.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.11μs | 44.1K | 828 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.27μs | 32.7K | 673 |
| Medium Payload | 🥉 JSON | Unmarshal | 197.92μs | 58.0K | 728 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.45μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 105.92μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 193.15μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 289.06μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 313.40μs | 218.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 414.28μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.10μs | 282.8K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 279.45μs | 379.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 515.96μs | 350.6K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 663.37μs | 323.5K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 1.87ms | 494.9K | 6.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 357ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.08μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 1.12μs | 1.2K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.67μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.28μs | 2.7K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.55μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.05μs | 256 | 7 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.67μs | 1.3K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.25μs | 2.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.95μs | 2.5K | 54 |
| Small Struct | 🥉 JSON | Unmarshal | 33.23μs | 8.0K | 117 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.35μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.50μs | 24.6K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.79μs | 20.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.36μs | 33.0K | 21 |
| Medium Payload | 🥈 CBOR | Marshal | 27.42μs | 24.6K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 44.28μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.39μs | 26.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 51.39μs | 58.8K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 64.67μs | 33.9K | 627 |
| Medium Payload | 🥈 CBOR | Unmarshal | 87.65μs | 32.8K | 672 |
| Medium Payload | 🥉 JSON | Unmarshal | 265.41μs | 54.6K | 731 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.05μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.98μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 151.11μs | 206.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 221.15μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 276.93μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 492.71μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 268.29μs | 268.3K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 424.42μs | 546.7K | 586 |
| Large Payload | 🥈 MessagePack | Unmarshal | 662.30μs | 350.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 840.85μs | 313.8K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.55ms | 519.3K | 6.8K |

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

