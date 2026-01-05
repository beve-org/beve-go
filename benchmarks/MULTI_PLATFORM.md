# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-01-05 04:17:06 UTC

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
| Apple M1 (Virtual) | 1.00μs | 529ns | 3.69μs | 1.91μs | 2.47μs |
| AMD EPYC 7763 64-Core Processor | 1.86μs | 800ns | 3.82μs | 1.54μs | 2.96μs |
| Neoverse-N2 | 711ns | 201ns | 1.46μs | 592ns | 2.44μs |
| Unknown CPU | 1.14μs | 252ns | 2.50μs | 1.62μs | 1.07μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.14μs | 16.24μs | 6.94μs | 1.86μs |
| AMD EPYC 7763 64-Core Processor | 1.98μs | 28.63μs | 1.61μs | 7.24μs |
| Neoverse-N2 | 1.05μs | 23.19μs | 3.84μs | 2.92μs |
| Unknown CPU | 1.15μs | 23.12μs | 8.94μs | 1.79μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (529ns) | 🥇 BEVE (1.14μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (800ns) | 🥈 CBOR (1.61μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (201ns) | 🥇 BEVE (1.05μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (252ns) | 🥇 BEVE (1.15μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 57.5% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 529ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.00μs | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.91μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.47μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.19μs | 1.3K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.69μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.14μs | 824 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.63μs | 961 | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.86μs | 832 | 20 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.94μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 16.24μs | 4.0K | 61 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.03μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.79μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 22.88μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.97μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 46.92μs | 16.6K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 57.45μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.51μs | 26.9K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 35.55μs | 26.7K | 478 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.50μs | 42.7K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 61.19μs | 35.2K | 724 |
| Medium Payload | 🥉 JSON | Unmarshal | 232.99μs | 57.9K | 732 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.98μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 123.56μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 231.99μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 306.78μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 545.37μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 565.71μs | 214.1K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 304.16μs | 284.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 377.98μs | 350.6K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 641.81μs | 363.9K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 773.25μs | 316.1K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.38ms | 527.6K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 800ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.22μs | 940 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.54μs | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.86μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.96μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.82μs | 1.5K | 1 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.61μs | 328 | 10 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.98μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.15μs | 2.0K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.24μs | 4.7K | 99 |
| Small Struct | 🥉 JSON | Unmarshal | 28.63μs | 7.4K | 96 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.09μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.08μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.25μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.92μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.21μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 38.42μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.63μs | 26.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.95μs | 48.6K | 71 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.00μs | 35.1K | 649 |
| Medium Payload | 🥈 CBOR | Unmarshal | 76.88μs | 36.9K | 753 |
| Medium Payload | 🥉 JSON | Unmarshal | 185.85μs | 42.5K | 576 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.49μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.88μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 173.91μs | 232.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 213.88μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 343.81μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 453.26μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 251.74μs | 273.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 369.12μs | 546.2K | 593 |
| Large Payload | 🥈 MessagePack | Unmarshal | 608.68μs | 371.1K | 6.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 738.07μs | 325.4K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.37ms | 551.7K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 201ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 592ns | 384 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 711ns | 1.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.13μs | 663 | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.46μs | 704 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.44μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.05μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.42μs | 1.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.92μs | 2.1K | 47 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.84μs | 2.1K | 45 |
| Small Struct | 🥉 JSON | Unmarshal | 23.19μs | 7.7K | 106 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.43μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.50μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.43μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.30μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 31.35μs | 25.0K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 33.48μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.67μs | 30.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.55μs | 50.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.16μs | 41.0K | 772 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.52μs | 31.2K | 642 |
| Medium Payload | 🥉 JSON | Unmarshal | 220.87μs | 63.0K | 835 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.98μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 105.61μs | 204.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 192.31μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 269.06μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 280.46μs | 198.5K | 3 |
| Large Payload | 🥉 JSON | Marshal | 380.83μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 216.45μs | 275.1K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 271.24μs | 372.8K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 514.51μs | 358.3K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 637.27μs | 310.6K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.90ms | 510.1K | 6.7K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 252ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.07μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Marshal | 1.14μs | 1.4K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.49μs | 1.5K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.62μs | 1.4K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.50μs | 1.2K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.15μs | 1.2K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.79μs | 688 | 17 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.02μs | 4.7K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.94μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 23.12μs | 4.7K | 85 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.32μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.97μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.13μs | 16.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.61μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.79μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 53.08μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.50μs | 24.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.22μs | 60.3K | 78 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.92μs | 21.2K | 437 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 88.58μs | 44.0K | 831 |
| Medium Payload | 🥉 JSON | Unmarshal | 290.65μs | 55.8K | 706 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 82.58μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.06μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 153.53μs | 206.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 221.13μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 265.56μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 506.38μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 259.59μs | 263.8K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 448.23μs | 558.7K | 586 |
| Large Payload | 🥈 MessagePack | Unmarshal | 695.46μs | 379.8K | 7.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 853.22μs | 321.0K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.43ms | 502.0K | 6.6K |

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

