# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-03-02 04:48:33 UTC

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
| Apple M1 (Virtual) | 688ns | 495ns | 993ns | 1.49μs | 2.83μs |
| AMD EPYC 7763 64-Core Processor | 1.11μs | 288ns | 2.82μs | 3.13μs | 4.66μs |
| Neoverse-N2 | 337ns | 394ns | 2.91μs | 694ns | 1.53μs |
| Unknown CPU | 2.56μs | 395ns | 5.69μs | 2.88μs | 1.35μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 717ns | 13.00μs | 2.78μs | 3.36μs |
| AMD EPYC 7763 64-Core Processor | 2.25μs | 16.38μs | 7.53μs | 2.38μs |
| Neoverse-N2 | 907ns | 13.87μs | 3.55μs | 4.59μs |
| Unknown CPU | 2.58μs | 31.79μs | 11.24μs | 6.96μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (495ns) | 🥇 BEVE (717ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (288ns) | 🥇 BEVE (2.25μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (337ns) | 🥇 BEVE (907ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (395ns) | 🥉 Sonic (902ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 58.7% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 495ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 688ns | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 993ns | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.49μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.83μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 3.73μs | 2.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 717ns | 1.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.78μs | 2.1K | 46 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.30μs | 5.7K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.36μs | 4.0K | 85 |
| Small Struct | 🥉 JSON | Unmarshal | 13.00μs | 4.3K | 71 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.70μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.27μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 13.13μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.21μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 35.02μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 38.68μs | 22.0K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 15.73μs | 27.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 24.45μs | 30.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 35.01μs | 32.9K | 607 |
| Medium Payload | 🥈 CBOR | Unmarshal | 50.89μs | 34.5K | 706 |
| Medium Payload | 🥉 JSON | Unmarshal | 157.93μs | 49.0K | 642 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 57.47μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 78.23μs | 196.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 150.50μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 190.11μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 324.95μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 394.06μs | 205.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 195.91μs | 282.6K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 270.28μs | 335.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 387.96μs | 345.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 457.08μs | 294.0K | 6.0K |
| Large Payload | 🥉 JSON | Unmarshal | 1.73ms | 536.3K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 288ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.11μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.53μs | 2.1K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.82μs | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 3.13μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.66μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.25μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.38μs | 1.4K | 31 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.49μs | 3.7K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.53μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 16.38μs | 4.1K | 66 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.86μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.59μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.41μs | 20.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.47μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.62μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.61μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.36μs | 23.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.46μs | 50.4K | 68 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 51.52μs | 31.8K | 581 |
| Medium Payload | 🥈 CBOR | Unmarshal | 76.74μs | 32.9K | 677 |
| Medium Payload | 🥉 JSON | Unmarshal | 223.88μs | 53.2K | 685 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.67μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 109.95μs | 172.1K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 149.04μs | 198.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 208.25μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 331.10μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 454.90μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 245.41μs | 268.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 373.85μs | 556.3K | 584 |
| Large Payload | 🥈 MessagePack | Unmarshal | 562.34μs | 335.9K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 752.37μs | 332.2K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.28ms | 509.3K | 6.7K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 337ns | 320 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 394ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 694ns | 480 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.53μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 2.91μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.58μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 907ns | 952 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.22μs | 5.5K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.55μs | 1.9K | 42 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.59μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 13.87μs | 4.1K | 65 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.53μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.09μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.66μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 23.26μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 28.00μs | 20.9K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 43.45μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.61μs | 23.2K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.82μs | 37.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 54.49μs | 37.8K | 703 |
| Medium Payload | 🥈 CBOR | Unmarshal | 72.49μs | 36.6K | 750 |
| Medium Payload | 🥉 JSON | Unmarshal | 198.05μs | 57.3K | 717 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.19μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 112.49μs | 196.8K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 185.44μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 282.93μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 300.19μs | 206.8K | 3 |
| Large Payload | 🥉 JSON | Marshal | 384.17μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 234.98μs | 270.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 310.65μs | 413.2K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 541.70μs | 366.3K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 649.96μs | 311.0K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.96ms | 526.8K | 6.9K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 395ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.26μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.35μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 2.56μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.88μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 5.69μs | 2.3K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 902ns | 570 | 5 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.58μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.96μs | 3.5K | 73 |
| Small Struct | 🥈 CBOR | Unmarshal | 11.24μs | 4.2K | 89 |
| Small Struct | 🥉 JSON | Unmarshal | 31.79μs | 7.6K | 104 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 11.07μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 18.93μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 24.31μs | 20.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 31.34μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.80μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 71.33μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.57μs | 25.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 57.18μs | 58.6K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 82.61μs | 37.5K | 695 |
| Medium Payload | 🥈 CBOR | Unmarshal | 106.52μs | 35.0K | 718 |
| Medium Payload | 🥉 JSON | Unmarshal | 300.50μs | 57.0K | 777 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 84.59μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 134.49μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 175.01μs | 223.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 226.97μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 309.26μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 535.65μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 286.32μs | 265.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 462.11μs | 546.8K | 582 |
| Large Payload | 🥈 MessagePack | Unmarshal | 707.75μs | 344.5K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 961.88μs | 337.6K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.62ms | 480.3K | 6.4K |

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

