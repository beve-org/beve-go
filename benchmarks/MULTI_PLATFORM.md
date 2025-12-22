# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-12-22 04:06:28 UTC

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
| Apple M1 (Virtual) | 734ns | 590ns | 650ns | 644ns | 1.30μs |
| AMD EPYC 7763 64-Core Processor | 720ns | 616ns | 4.41μs | 1.86μs | 1.35μs |
| Neoverse-N2 | 1.04μs | 630ns | 4.61μs | 870ns | 1.49μs |
| Unknown CPU | 2.84μs | 905ns | 2.46μs | 528ns | 3.23μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.78μs | 18.84μs | 1.55μs | 5.82μs |
| AMD EPYC 7763 64-Core Processor | 1.21μs | 5.81μs | 1.58μs | 1.93μs |
| Neoverse-N2 | 1.10μs | 21.10μs | 3.98μs | 5.66μs |
| Unknown CPU | 2.04μs | 28.25μs | 7.35μs | 1.92μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (590ns) | 🥈 CBOR (1.55μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥉 Sonic (548ns) | 🥇 BEVE (1.21μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (630ns) | 🥇 BEVE (1.10μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥈 CBOR (528ns) | 🥈 MessagePack (1.92μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 33.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 590ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 644ns | 640 | 1 |
| Small Struct | 🥉 JSON | Marshal | 650ns | 384 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 734ns | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.30μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 4.00μs | 2.1K | 2 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.55μs | 712 | 18 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.78μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.62μs | 3.8K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.82μs | 4.4K | 92 |
| Small Struct | 🥉 JSON | Unmarshal | 18.84μs | 7.2K | 92 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.83μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.07μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 22.88μs | 20.5K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 32.71μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.77μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 55.15μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 16.58μs | 26.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.16μs | 31.3K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 48.48μs | 29.9K | 619 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 63.91μs | 40.8K | 764 |
| Medium Payload | 🥉 JSON | Unmarshal | 191.28μs | 50.1K | 673 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 61.94μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 112.35μs | 204.8K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 142.62μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 226.08μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 352.45μs | 196.9K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 465.86μs | 222.1K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 190.62μs | 274.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 289.77μs | 347.0K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 445.07μs | 356.1K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 553.59μs | 310.6K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.90ms | 570.5K | 7.5K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 548ns | 613 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 616ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 720ns | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.35μs | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 1.86μs | 1.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 4.41μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.21μs | 1.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.58μs | 424 | 12 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.65μs | 2.1K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.93μs | 976 | 23 |
| Small Struct | 🥉 JSON | Unmarshal | 5.81μs | 1.3K | 26 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.54μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.15μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.49μs | 25.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.90μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.98μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.00μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.40μs | 27.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.18μs | 61.9K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.08μs | 32.6K | 598 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.22μs | 34.3K | 706 |
| Medium Payload | 🥉 JSON | Unmarshal | 275.28μs | 66.7K | 879 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 83.43μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.43μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 155.97μs | 215.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 218.18μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 324.41μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 447.23μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 259.62μs | 289.9K | 415 |
| Large Payload | 🥉 Sonic | Unmarshal | 366.16μs | 541.9K | 573 |
| Large Payload | 🥈 MessagePack | Unmarshal | 573.79μs | 356.7K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 700.49μs | 304.5K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.34ms | 519.3K | 6.8K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 630ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 870ns | 768 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.04μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.49μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 2.73μs | 2.1K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.61μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.10μs | 1.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.43μs | 1.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.98μs | 2.2K | 48 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.66μs | 5.2K | 107 |
| Small Struct | 🥉 JSON | Unmarshal | 21.10μs | 7.4K | 97 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.13μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.26μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.06μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.86μs | 16.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.99μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.40μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.52μs | 25.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.80μs | 50.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 57.28μs | 41.6K | 778 |
| Medium Payload | 🥈 CBOR | Unmarshal | 61.11μs | 28.9K | 596 |
| Medium Payload | 🥉 JSON | Unmarshal | 250.04μs | 72.9K | 965 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.32μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 104.30μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 187.95μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 273.62μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 298.67μs | 216.6K | 3 |
| Large Payload | 🥉 JSON | Marshal | 382.30μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 218.59μs | 280.8K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 281.62μs | 400.0K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 498.89μs | 341.9K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 672.29μs | 334.4K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.80ms | 481.2K | 6.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥈 CBOR | Marshal | 528ns | 352 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 905ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 2.03μs | 2.7K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.46μs | 1.2K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 2.84μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.23μs | 4.1K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.92μs | 832 | 20 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.04μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.86μs | 7.3K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.35μs | 3.3K | 71 |
| Small Struct | 🥉 JSON | Unmarshal | 28.25μs | 7.5K | 101 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.70μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.33μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.66μs | 16.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.25μs | 33.0K | 21 |
| Medium Payload | 🥈 CBOR | Marshal | 30.76μs | 19.1K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 49.04μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 32.10μs | 29.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 59.18μs | 57.5K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 103.62μs | 39.7K | 742 |
| Medium Payload | 🥈 CBOR | Unmarshal | 141.92μs | 38.7K | 792 |
| Medium Payload | 🥉 JSON | Unmarshal | 298.00μs | 55.9K | 740 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 90.60μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 148.16μs | 196.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 209.77μs | 215.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 236.73μs | 213.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 372.54μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 560.30μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 297.70μs | 289.2K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 458.17μs | 600.7K | 607 |
| Large Payload | 🥈 MessagePack | Unmarshal | 655.82μs | 336.0K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 962.42μs | 319.8K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.64ms | 534.7K | 7.0K |

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

