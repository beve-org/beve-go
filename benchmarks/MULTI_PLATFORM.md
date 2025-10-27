# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-27 03:52:13 UTC

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
| Apple M1 (Virtual) | 1.80μs | 310ns | 2.39μs | 1.41μs | 709ns |
| AMD EPYC 7763 64-Core Processor | 639ns | 604ns | 5.55μs | 2.62μs | 4.19μs |
| Neoverse-N2 | 954ns | 563ns | 1.38μs | 1.77μs | 3.60μs |
| Unknown CPU | 990ns | 407ns | 3.61μs | 1.13μs | 1.21μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 751ns | 22.48μs | 3.08μs | 1.42μs |
| AMD EPYC 7763 64-Core Processor | 906ns | 10.83μs | 8.57μs | 1.26μs |
| Neoverse-N2 | 1.20μs | 16.03μs | 5.33μs | 5.00μs |
| Unknown CPU | 1.73μs | 27.77μs | 3.98μs | 5.31μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (310ns) | 🥇 BEVE (751ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥉 Sonic (440ns) | 🥇 BEVE (906ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (563ns) | 🥉 Sonic (1.13μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥉 Sonic (374ns) | 🥇 BEVE (1.73μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 54.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 310ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 709ns | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 1.41μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.80μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.39μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 4.10μs | 1.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 751ns | 728 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.42μs | 1.0K | 24 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.08μs | 1.9K | 43 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.96μs | 5.4K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 22.48μs | 7.3K | 92 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.75μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.79μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.73μs | 27.3K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.99μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.82μs | 19.3K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 46.97μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 19.53μs | 30.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.65μs | 36.6K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 53.11μs | 32.6K | 671 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.53μs | 41.1K | 770 |
| Medium Payload | 🥉 JSON | Unmarshal | 224.99μs | 70.5K | 893 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 62.20μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.94μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 145.13μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 291.63μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 434.30μs | 205.1K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 508.28μs | 214.2K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 210.79μs | 278.3K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 316.13μs | 358.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 405.94μs | 318.3K | 5.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 527.86μs | 310.4K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.96ms | 525.0K | 6.9K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 440ns | 357 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 604ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 639ns | 1.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.62μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.19μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.55μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 906ns | 888 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.26μs | 400 | 11 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.29μs | 1.3K | 7 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.57μs | 5.2K | 106 |
| Small Struct | 🥉 JSON | Unmarshal | 10.83μs | 2.4K | 47 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.76μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.16μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.25μs | 27.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.00μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.39μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.41μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.21μs | 29.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.87μs | 44.2K | 64 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.09μs | 35.4K | 655 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.54μs | 32.2K | 662 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.79μs | 46.6K | 611 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.45μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.41μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 155.76μs | 215.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 203.86μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 305.52μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 448.80μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 238.84μs | 272.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 360.86μs | 563.2K | 585 |
| Large Payload | 🥈 MessagePack | Unmarshal | 560.82μs | 343.1K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 739.00μs | 336.3K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.30ms | 535.1K | 7.1K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 563ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 954ns | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.38μs | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.77μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.05μs | 2.4K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.60μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.13μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.20μs | 1.7K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.00μs | 4.4K | 93 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.33μs | 3.1K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 16.03μs | 4.4K | 75 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.56μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.34μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.94μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.79μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 28.77μs | 20.9K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 44.14μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.36μs | 28.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.00μs | 37.2K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 56.32μs | 26.3K | 543 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 57.28μs | 42.5K | 797 |
| Medium Payload | 🥉 JSON | Unmarshal | 175.61μs | 47.0K | 633 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.88μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.49μs | 196.8K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 182.76μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 285.82μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 311.61μs | 223.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 362.53μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.87μs | 294.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 282.66μs | 394.2K | 205 |
| Large Payload | 🥈 MessagePack | Unmarshal | 505.38μs | 346.1K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 651.30μs | 318.1K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.91ms | 515.4K | 6.7K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 374ns | 268 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 407ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 990ns | 1.2K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.13μs | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.21μs | 1.0K | 6 |
| Small Struct | 🥉 JSON | Marshal | 3.61μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.73μs | 2.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.98μs | 1.5K | 35 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.25μs | 4.7K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.31μs | 3.6K | 77 |
| Small Struct | 🥉 JSON | Unmarshal | 27.77μs | 8.1K | 118 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.31μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.13μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.15μs | 22.3K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.48μs | 19.1K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 42.82μs | 19.3K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 44.35μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.96μs | 29.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 60.09μs | 69.6K | 80 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 70.34μs | 35.2K | 657 |
| Medium Payload | 🥈 CBOR | Unmarshal | 90.61μs | 36.9K | 756 |
| Medium Payload | 🥉 JSON | Unmarshal | 239.20μs | 57.4K | 767 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 71.17μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 138.36μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 183.25μs | 207.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 228.45μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 322.25μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 462.85μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 299.45μs | 278.1K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 500.00μs | 548.6K | 581 |
| Large Payload | 🥈 MessagePack | Unmarshal | 720.54μs | 367.5K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 879.79μs | 333.5K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.34ms | 538.8K | 7.1K |

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

