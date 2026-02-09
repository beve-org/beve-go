# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-02-09 04:59:07 UTC

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
| Apple M1 (Virtual) | 1.90μs | 1.13μs | 4.32μs | 1.66μs | 1.95μs |
| AMD EPYC 7763 64-Core Processor | 1.39μs | 727ns | 4.66μs | 1.43μs | 4.32μs |
| Neoverse-N2 | 306ns | 670ns | 3.87μs | 947ns | 2.42μs |
| Unknown CPU | 2.23μs | 1.38μs | 5.92μs | 1.73μs | 2.51μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 719ns | 12.05μs | 10.24μs | 3.69μs |
| AMD EPYC 7763 64-Core Processor | 2.05μs | 21.59μs | 8.27μs | 2.06μs |
| Neoverse-N2 | 1.49μs | 21.92μs | 6.40μs | 1.59μs |
| Unknown CPU | 2.93μs | 9.46μs | 6.70μs | 6.55μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (1.13μs) | 🥇 BEVE (719ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (727ns) | 🥇 BEVE (2.05μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (306ns) | 🥇 BEVE (1.49μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥉 Sonic (831ns) | 🥇 BEVE (2.93μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 70.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.13μs | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.66μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.90μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.95μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 4.32μs | 2.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 9.44μs | 2.7K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 719ns | 600 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.69μs | 1.5K | 33 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.09μs | 5.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 10.24μs | 3.5K | 74 |
| Small Struct | 🥉 JSON | Unmarshal | 12.05μs | 2.1K | 37 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.97μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.54μs | 13.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 31.49μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 46.63μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 59.76μs | 18.7K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 77.28μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 34.08μs | 31.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.62μs | 45.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 67.27μs | 35.6K | 658 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.86μs | 33.4K | 685 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.00μs | 56.0K | 724 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.39μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 148.54μs | 196.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 229.32μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 309.59μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 526.05μs | 205.1K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 562.95μs | 206.2K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 261.72μs | 279.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 421.64μs | 340.6K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 691.00μs | 355.0K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 841.43μs | 320.4K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.28ms | 506.9K | 6.6K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 727ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 830ns | 926 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.39μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.43μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.32μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.66μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.05μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.06μs | 1.2K | 27 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.01μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.27μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 21.59μs | 4.8K | 88 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.88μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.99μs | 24.6K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.02μs | 22.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 25.35μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.35μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.47μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.91μs | 23.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.00μs | 58.4K | 73 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.37μs | 39.2K | 732 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.64μs | 36.2K | 744 |
| Medium Payload | 🥉 JSON | Unmarshal | 224.47μs | 58.5K | 732 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 83.67μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.52μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 158.92μs | 223.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 212.51μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 315.18μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 450.95μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.24μs | 259.2K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 356.75μs | 552.1K | 590 |
| Large Payload | 🥈 MessagePack | Unmarshal | 550.88μs | 339.2K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 699.98μs | 316.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.26ms | 541.6K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 306ns | 256 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 617ns | 286 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 670ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 947ns | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.42μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.87μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.49μs | 2.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.59μs | 832 | 20 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.60μs | 2.0K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.40μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 21.92μs | 7.5K | 100 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.81μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.42μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.27μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 23.25μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 28.51μs | 20.9K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 41.68μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.27μs | 25.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.55μs | 38.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.44μs | 39.8K | 742 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.50μs | 32.4K | 664 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.51μs | 55.8K | 719 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.06μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 106.97μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 193.49μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 269.26μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 298.83μs | 214.9K | 3 |
| Large Payload | 🥉 JSON | Marshal | 376.55μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.76μs | 273.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 299.91μs | 412.9K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 513.92μs | 348.6K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 658.86μs | 319.7K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.82ms | 480.2K | 6.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 831ns | 613 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.38μs | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.73μs | 1.2K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 2.23μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.51μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 5.92μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.93μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.35μs | 3.6K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.55μs | 4.2K | 88 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.70μs | 3.2K | 70 |
| Small Struct | 🥉 JSON | Unmarshal | 9.46μs | 2.2K | 41 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 10.77μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.15μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.45μs | 19.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.74μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 48.57μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 54.49μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 33.85μs | 32.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 67.00μs | 63.5K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 72.51μs | 37.9K | 707 |
| Medium Payload | 🥈 CBOR | Unmarshal | 93.63μs | 32.5K | 664 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.34μs | 42.2K | 587 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.04μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 140.41μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 201.72μs | 217.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 231.68μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 365.23μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 456.49μs | 197.0K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 329.84μs | 265.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 530.23μs | 549.7K | 574 |
| Large Payload | 🥈 MessagePack | Unmarshal | 749.56μs | 342.2K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 857.04μs | 325.4K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.38ms | 528.9K | 7.0K |

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

