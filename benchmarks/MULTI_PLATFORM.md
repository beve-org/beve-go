# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-17 10:32:52 UTC

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
| Apple M1 (Virtual) | 1.77μs | 1.02μs | 2.10μs | 1.71μs | 2.45μs |
| AMD EPYC 7763 64-Core Processor | 1.00μs | 885ns | 4.57μs | 1.13μs | 4.04μs |
| Neoverse-N2 | 688ns | 777ns | 4.47μs | 1.00μs | 1.46μs |
| Unknown CPU | 903ns | 592ns | 4.03μs | 2.42μs | 2.70μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.26μs | 16.60μs | 5.13μs | 5.74μs |
| AMD EPYC 7763 64-Core Processor | 988ns | 13.22μs | 8.59μs | 4.34μs |
| Neoverse-N2 | 1.44μs | 25.89μs | 6.65μs | 5.83μs |
| Unknown CPU | 2.49μs | 16.85μs | 9.43μs | 3.10μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (1.02μs) | 🥇 BEVE (1.26μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (885ns) | 🥉 Sonic (969ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (688ns) | 🥇 BEVE (1.44μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (592ns) | 🥇 BEVE (2.49μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 64.0% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.02μs | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.71μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.77μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.10μs | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.20μs | 933 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.45μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.26μs | 1.1K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.03μs | 5.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.13μs | 2.7K | 57 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.74μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 16.60μs | 4.1K | 66 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.41μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.85μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.77μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.64μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 49.60μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 52.06μs | 20.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.41μs | 24.9K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.92μs | 31.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 61.80μs | 39.7K | 745 |
| Medium Payload | 🥈 CBOR | Unmarshal | 72.45μs | 29.6K | 605 |
| Medium Payload | 🥉 JSON | Unmarshal | 307.25μs | 69.8K | 906 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.44μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 103.03μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 195.99μs | 172.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 285.06μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 448.16μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 547.67μs | 222.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 253.36μs | 267.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 360.09μs | 348.3K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 524.60μs | 338.8K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 773.64μs | 309.1K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.39ms | 567.0K | 7.3K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 885ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.00μs | 1.8K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.13μs | 1.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.59μs | 2.4K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 4.04μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.57μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 969ns | 842 | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 988ns | 1.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.34μs | 3.2K | 68 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.59μs | 5.2K | 107 |
| Small Struct | 🥉 JSON | Unmarshal | 13.22μs | 3.8K | 56 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.09μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.72μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.25μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.62μs | 25.0K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.61μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 58.74μs | 33.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.27μs | 27.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.48μs | 49.3K | 62 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 47.33μs | 29.1K | 527 |
| Medium Payload | 🥈 CBOR | Unmarshal | 76.88μs | 37.8K | 774 |
| Medium Payload | 🥉 JSON | Unmarshal | 232.74μs | 56.2K | 728 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.00μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.17μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 162.06μs | 231.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 226.09μs | 213.2K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 300.50μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 468.87μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 226.60μs | 262.6K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 360.26μs | 564.4K | 599 |
| Large Payload | 🥈 MessagePack | Unmarshal | 579.15μs | 366.2K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 654.54μs | 291.4K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.33ms | 559.3K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 688ns | 1.0K | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 777ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.00μs | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.46μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.73μs | 1.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.47μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.44μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.98μs | 2.8K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.83μs | 4.8K | 103 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.65μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 25.89μs | 8.1K | 118 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.07μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.67μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.28μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 24.66μs | 16.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.20μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 41.44μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.24μs | 24.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.19μs | 45.6K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.55μs | 32.5K | 594 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.44μs | 35.2K | 727 |
| Medium Payload | 🥉 JSON | Unmarshal | 178.49μs | 50.5K | 633 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.94μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 110.51μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 190.20μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 279.54μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 307.66μs | 216.9K | 3 |
| Large Payload | 🥉 JSON | Marshal | 376.07μs | 205.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 220.81μs | 251.7K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 303.88μs | 412.2K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 531.65μs | 353.4K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 646.50μs | 305.9K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 1.95ms | 524.4K | 6.8K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 592ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 903ns | 1.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.03μs | 1.2K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.42μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.70μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.03μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.49μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.10μs | 1.7K | 37 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.50μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.43μs | 4.6K | 96 |
| Small Struct | 🥉 JSON | Unmarshal | 16.85μs | 4.0K | 63 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.91μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.43μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 13.36μs | 16.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.30μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.23μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 48.48μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.22μs | 30.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 52.06μs | 66.7K | 80 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 72.54μs | 38.3K | 710 |
| Medium Payload | 🥈 CBOR | Unmarshal | 93.44μs | 37.5K | 767 |
| Medium Payload | 🥉 JSON | Unmarshal | 243.52μs | 49.4K | 652 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.29μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.31μs | 196.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 155.98μs | 206.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 213.97μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 275.79μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 467.94μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 272.01μs | 273.8K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 404.56μs | 529.6K | 570 |
| Large Payload | 🥈 MessagePack | Unmarshal | 679.46μs | 350.6K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 822.77μs | 311.2K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.54ms | 505.8K | 6.7K |

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

