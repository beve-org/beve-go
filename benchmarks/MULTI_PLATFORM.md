# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-17 11:35:33 UTC

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
| Apple M1 (Virtual) | 1.07μs | 347ns | 3.47μs | 1.17μs | 1.14μs |
| AMD EPYC 7763 64-Core Processor | 1.62μs | 329ns | 4.36μs | 790ns | 2.77μs |
| Neoverse-N2 | 1.31μs | 709ns | 4.26μs | 1.64μs | 3.79μs |
| Unknown CPU | 1.42μs | 1.19μs | 3.19μs | 1.88μs | 6.60μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.88μs | 19.37μs | 5.30μs | 1.78μs |
| AMD EPYC 7763 64-Core Processor | 1.66μs | 14.30μs | 3.13μs | 3.35μs |
| Neoverse-N2 | 1.20μs | 22.59μs | 6.07μs | 2.69μs |
| Unknown CPU | 2.90μs | 11.52μs | 2.74μs | 6.64μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (347ns) | 🥈 MessagePack (1.78μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (329ns) | 🥇 BEVE (1.66μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (709ns) | 🥇 BEVE (1.20μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (1.19μs) | 🥈 CBOR (2.74μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 64.2% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 347ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.07μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.14μs | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 1.17μs | 1.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.47μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 4.84μs | 2.3K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.78μs | 640 | 16 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.88μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.68μs | 2.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.30μs | 2.4K | 52 |
| Small Struct | 🥉 JSON | Unmarshal | 19.37μs | 4.5K | 77 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.64μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.37μs | 27.3K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.76μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.88μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 43.02μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 46.98μs | 18.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.00μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.54μs | 32.4K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.12μs | 34.5K | 637 |
| Medium Payload | 🥈 CBOR | Unmarshal | 44.54μs | 27.8K | 574 |
| Medium Payload | 🥉 JSON | Unmarshal | 233.17μs | 54.6K | 704 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.79μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 96.92μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 165.55μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 303.32μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 415.27μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 484.38μs | 206.1K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 274.56μs | 270.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 414.23μs | 352.4K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 425.61μs | 330.7K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 803.93μs | 317.1K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.17ms | 519.4K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 329ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 463ns | 428 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 790ns | 640 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.62μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.77μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.36μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.66μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.69μs | 4.2K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.13μs | 1.4K | 32 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.35μs | 2.3K | 49 |
| Small Struct | 🥉 JSON | Unmarshal | 14.30μs | 2.4K | 48 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.41μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.14μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.75μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.04μs | 18.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.46μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.64μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.61μs | 30.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.88μs | 64.9K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.42μs | 35.8K | 658 |
| Medium Payload | 🥈 CBOR | Unmarshal | 78.16μs | 33.6K | 692 |
| Medium Payload | 🥉 JSON | Unmarshal | 220.20μs | 52.6K | 714 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.78μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.59μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 163.35μs | 223.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 210.07μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 331.36μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 453.90μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 239.56μs | 271.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 344.73μs | 525.3K | 548 |
| Large Payload | 🥈 MessagePack | Unmarshal | 571.66μs | 361.8K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 728.17μs | 307.5K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.24ms | 533.0K | 6.9K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 709ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.31μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.64μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.48μs | 2.7K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.79μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.26μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.20μs | 1.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.69μs | 1.9K | 41 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.79μs | 4.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.07μs | 3.6K | 77 |
| Small Struct | 🥉 JSON | Unmarshal | 22.59μs | 7.6K | 103 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.15μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.75μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.72μs | 18.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 25.11μs | 18.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.67μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 33.03μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Unmarshal | 23.71μs | 28.8K | 33 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.20μs | 32.7K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.35μs | 27.5K | 492 |
| Medium Payload | 🥈 CBOR | Unmarshal | 58.80μs | 27.7K | 570 |
| Medium Payload | 🥉 JSON | Unmarshal | 164.04μs | 42.3K | 580 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.89μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 106.48μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 193.55μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 272.97μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 300.74μs | 208.2K | 3 |
| Large Payload | 🥉 JSON | Marshal | 390.10μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 223.94μs | 270.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 282.30μs | 387.5K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 530.66μs | 360.2K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 641.86μs | 308.4K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.03ms | 550.3K | 7.2K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.19μs | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.42μs | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.88μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.85μs | 2.1K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.19μs | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 6.60μs | 8.2K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.74μs | 616 | 16 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.90μs | 3.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.64μs | 3.4K | 72 |
| Small Struct | 🥉 Sonic | Unmarshal | 7.10μs | 7.4K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 11.52μs | 2.1K | 38 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 10.94μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.79μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.77μs | 20.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 25.66μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 43.12μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 70.66μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 40.16μs | 31.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 55.62μs | 56.0K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 93.81μs | 42.1K | 788 |
| Medium Payload | 🥈 CBOR | Unmarshal | 120.72μs | 38.6K | 792 |
| Medium Payload | 🥉 JSON | Unmarshal | 315.64μs | 57.5K | 737 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 89.32μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 144.25μs | 188.4K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 190.63μs | 215.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 216.05μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 284.19μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 515.31μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 283.56μs | 275.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 435.13μs | 570.0K | 579 |
| Large Payload | 🥈 MessagePack | Unmarshal | 685.81μs | 370.5K | 6.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 895.95μs | 335.4K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.58ms | 518.5K | 6.9K |

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

