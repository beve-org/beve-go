# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-03-23 05:02:02 UTC

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
| Apple M1 (Virtual) | 836ns | 491ns | 916ns | 1.43μs | 1.62μs |
| AMD EPYC 7763 64-Core Processor | 357ns | 384ns | 2.87μs | 2.15μs | 2.83μs |
| Neoverse-N2 | 385ns | 287ns | 1.86μs | 527ns | 1.54μs |
| Unknown CPU | 2.58μs | 1.04μs | 5.71μs | 2.98μs | 1.89μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.23μs | 9.07μs | 5.34μs | 2.91μs |
| AMD EPYC 7763 64-Core Processor | 1.09μs | 14.71μs | 4.32μs | 4.54μs |
| Neoverse-N2 | 1.73μs | 7.70μs | 3.15μs | 2.69μs |
| Unknown CPU | 1.58μs | 17.38μs | 4.06μs | 4.62μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (491ns) | 🥇 BEVE (1.23μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (357ns) | 🥇 BEVE (1.09μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (287ns) | 🥇 BEVE (1.73μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (1.04μs) | 🥇 BEVE (1.58μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 57.6% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 491ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 836ns | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 916ns | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.43μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.62μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 2.39μs | 1.4K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.23μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.16μs | 3.4K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.91μs | 3.3K | 71 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.34μs | 4.6K | 96 |
| Small Struct | 🥉 JSON | Unmarshal | 9.07μs | 3.7K | 51 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 4.82μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.69μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 12.92μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.84μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 29.56μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 48.43μs | 27.4K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 15.86μs | 27.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.03μs | 37.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 35.25μs | 34.3K | 633 |
| Medium Payload | 🥈 CBOR | Unmarshal | 60.24μs | 36.1K | 744 |
| Medium Payload | 🥉 JSON | Unmarshal | 129.10μs | 39.4K | 514 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 56.52μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 79.25μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 140.35μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 166.96μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 318.28μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 385.08μs | 213.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 164.74μs | 275.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 275.20μs | 368.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 365.03μs | 342.1K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 504.70μs | 334.2K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.74ms | 552.3K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 357ns | 288 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 384ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 2.15μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.19μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.83μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.87μs | 1.4K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.09μs | 1.2K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.88μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.32μs | 2.2K | 48 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.54μs | 3.4K | 72 |
| Small Struct | 🥉 JSON | Unmarshal | 14.71μs | 3.9K | 58 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.44μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.70μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.79μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.15μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.68μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 55.65μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.56μs | 23.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.12μs | 65.3K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.69μs | 32.1K | 583 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.78μs | 33.0K | 680 |
| Medium Payload | 🥉 JSON | Unmarshal | 237.24μs | 57.5K | 749 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.65μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.57μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 174.20μs | 232.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 202.24μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 344.71μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 448.18μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 250.50μs | 267.1K | 415 |
| Large Payload | 🥉 Sonic | Unmarshal | 396.43μs | 587.4K | 597 |
| Large Payload | 🥈 MessagePack | Unmarshal | 563.35μs | 341.3K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 737.21μs | 331.1K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.18ms | 493.2K | 6.6K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 287ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 385ns | 416 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 527ns | 288 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.44μs | 933 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.54μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 1.86μs | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.73μs | 3.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.69μs | 1.7K | 38 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.15μs | 1.6K | 35 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.33μs | 5.3K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 7.70μs | 2.1K | 37 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.95μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.39μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.96μs | 18.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 27.66μs | 19.7K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.47μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 34.50μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.46μs | 26.1K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.84μs | 50.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.14μs | 35.4K | 659 |
| Medium Payload | 🥈 CBOR | Unmarshal | 78.41μs | 39.8K | 814 |
| Medium Payload | 🥉 JSON | Unmarshal | 234.97μs | 68.3K | 875 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.09μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 118.60μs | 188.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 178.50μs | 180.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 312.13μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 325.36μs | 226.5K | 3 |
| Large Payload | 🥉 JSON | Marshal | 395.45μs | 213.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 236.00μs | 269.8K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 298.19μs | 400.2K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 503.78μs | 333.0K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 653.66μs | 313.0K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.88ms | 499.5K | 6.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 1.04μs | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.30μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.89μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 2.58μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.98μs | 3.1K | 1 |
| Small Struct | 🥉 JSON | Marshal | 5.71μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.58μs | 1.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.49μs | 3.6K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.06μs | 1.6K | 37 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.62μs | 2.5K | 55 |
| Small Struct | 🥉 JSON | Unmarshal | 17.38μs | 4.0K | 62 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.71μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.78μs | 24.6K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.90μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 25.32μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.01μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 52.08μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.25μs | 22.1K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 57.90μs | 67.2K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 67.79μs | 35.2K | 650 |
| Medium Payload | 🥈 CBOR | Unmarshal | 81.61μs | 28.7K | 594 |
| Medium Payload | 🥉 JSON | Unmarshal | 254.81μs | 51.1K | 679 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.89μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.73μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 177.87μs | 223.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 221.02μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 287.02μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 538.66μs | 237.8K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 284.93μs | 267.0K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 447.24μs | 547.1K | 588 |
| Large Payload | 🥈 MessagePack | Unmarshal | 705.50μs | 348.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 832.41μs | 289.7K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.52ms | 507.2K | 6.5K |

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

