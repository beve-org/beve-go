# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-05-25 07:07:56 UTC

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
| benchmark-linux-amd-epyc-9v74-80-core-processor | AMD EPYC 9V74 80-Core Processor | Linux | [📄 Report](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.md) · [📊 JSON](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.json) · [📈 Chart](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 789ns | 513ns | 4.00μs | 2.13μs | 1.25μs |
| AMD EPYC 9V74 80-Core Processor | 733ns | 396ns | 3.78μs | 526ns | 1.15μs |
| Neoverse-N2 | 1.11μs | 773ns | 1.26μs | 737ns | 2.40μs |
| Unknown CPU | 1.47μs | 255ns | 3.60μs | 1.42μs | 1.44μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.44μs | 7.52μs | 2.37μs | 1.63μs |
| AMD EPYC 9V74 80-Core Processor | 986ns | 14.79μs | 2.77μs | 1.14μs |
| Neoverse-N2 | 1.20μs | 16.02μs | 7.71μs | 1.70μs |
| Unknown CPU | 1.42μs | 29.24μs | 5.40μs | 3.43μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (513ns) | 🥇 BEVE (1.44μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 9V74 80-Core Processor | 🥇 BEVE ZeroCopy (396ns) | 🥇 BEVE (986ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥈 CBOR (737ns) | 🥇 BEVE (1.20μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (255ns) | 🥇 BEVE (1.42μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 58.1% faster

### Platform Details

- **Apple M1 (Virtual)** (Darwin)
  - Architecture: arm64
  - Test Scenarios: 3

- **AMD EPYC 9V74 80-Core Processor** (Linux)
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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 513ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 789ns | 768 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.25μs | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 2.13μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.31μs | 1.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.00μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.44μs | 2.1K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.63μs | 440 | 12 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.37μs | 904 | 22 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.31μs | 3.4K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 7.52μs | 1.4K | 32 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.89μs | 4 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.22μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.33μs | 14.3K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.49μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 53.02μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 60.23μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.48μs | 23.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.00μs | 35.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 65.06μs | 41.7K | 788 |
| Medium Payload | 🥈 CBOR | Unmarshal | 86.36μs | 36.5K | 750 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.21μs | 43.8K | 567 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 84.28μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.37μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 234.67μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 366.30μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 536.48μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 587.47μs | 214.0K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 290.56μs | 273.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 386.63μs | 337.3K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 674.94μs | 346.6K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 796.45μs | 310.2K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.51ms | 527.0K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 9V74 80-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 396ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 526ns | 576 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 733ns | 1.5K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.15μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.66μs | 3.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.78μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 986ns | 1.7K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.14μs | 592 | 15 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.50μs | 4.7K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.77μs | 1.5K | 34 |
| Small Struct | 🥉 JSON | Unmarshal | 14.79μs | 4.7K | 83 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.63μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.68μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 13.14μs | 22.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 13.96μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.06μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 35.59μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 19.44μs | 31.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 30.03μs | 49.8K | 69 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 39.90μs | 33.3K | 610 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.97μs | 39.5K | 811 |
| Medium Payload | 🥉 JSON | Unmarshal | 153.84μs | 53.8K | 688 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 59.25μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.13μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 126.41μs | 208.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 156.34μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 271.57μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 333.68μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 193.38μs | 271.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 321.13μs | 542.7K | 574 |
| Large Payload | 🥈 MessagePack | Unmarshal | 442.28μs | 359.7K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 600.59μs | 316.9K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.54ms | 502.2K | 6.6K |

[📄 View full report](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥈 CBOR | Marshal | 737ns | 576 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 773ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.11μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.26μs | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.35μs | 1.6K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.40μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.20μs | 1.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.70μs | 872 | 21 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.53μs | 5.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.71μs | 4.6K | 98 |
| Small Struct | 🥉 JSON | Unmarshal | 16.02μs | 4.4K | 75 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.63μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.90μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 22.28μs | 24.6K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 31.86μs | 16.6K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 34.18μs | 25.4K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.65μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.52μs | 28.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.12μs | 44.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 42.65μs | 25.9K | 459 |
| Medium Payload | 🥈 CBOR | Unmarshal | 55.70μs | 24.7K | 512 |
| Medium Payload | 🥉 JSON | Unmarshal | 232.19μs | 65.7K | 863 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.05μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 122.42μs | 196.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 188.80μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 287.91μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 318.83μs | 225.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 417.09μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 231.74μs | 272.3K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 287.31μs | 374.2K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 523.12μs | 352.7K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 690.30μs | 335.2K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.00ms | 548.9K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 255ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 626ns | 684 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.42μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.44μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 1.47μs | 1.5K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.60μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.42μs | 1.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.33μs | 3.9K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.43μs | 1.7K | 38 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.40μs | 2.3K | 51 |
| Small Struct | 🥉 JSON | Unmarshal | 29.24μs | 7.6K | 102 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.95μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.54μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.31μs | 19.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.11μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.13μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.20μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.56μs | 28.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 53.81μs | 66.6K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 68.51μs | 35.8K | 661 |
| Medium Payload | 🥈 CBOR | Unmarshal | 94.62μs | 37.6K | 769 |
| Medium Payload | 🥉 JSON | Unmarshal | 246.88μs | 52.0K | 691 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.74μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 118.50μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 153.80μs | 206.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 213.43μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 276.77μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 495.58μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 271.13μs | 264.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 440.54μs | 569.5K | 588 |
| Large Payload | 🥈 MessagePack | Unmarshal | 648.24μs | 339.1K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 788.64μs | 290.9K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.67ms | 564.2K | 7.4K |

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

