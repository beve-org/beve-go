# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-02-02 04:56:30 UTC

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
| Apple M1 (Virtual) | 585ns | 566ns | 985ns | 1.50μs | 1.82μs |
| AMD EPYC 7763 64-Core Processor | 805ns | 422ns | 2.54μs | 2.14μs | 2.76μs |
| Neoverse-N2 | 887ns | 338ns | 2.94μs | 1.52μs | 2.32μs |
| Unknown CPU | 1.02μs | 457ns | 2.56μs | 2.62μs | 1.57μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.96μs | 20.10μs | 2.11μs | 2.90μs |
| AMD EPYC 7763 64-Core Processor | 698ns | 21.10μs | 4.18μs | 2.60μs |
| Neoverse-N2 | 1.77μs | 18.46μs | 7.77μs | 3.40μs |
| Unknown CPU | 3.22μs | 23.06μs | 9.28μs | 6.36μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (566ns) | 🥉 Sonic (1.86μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (422ns) | 🥇 BEVE (698ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (338ns) | 🥇 BEVE (1.77μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (457ns) | 🥉 Sonic (1.92μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 59.7% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 566ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 585ns | 1.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 985ns | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.50μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.82μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 4.15μs | 2.1K | 2 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.86μs | 2.6K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.96μs | 1.7K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.11μs | 1.1K | 26 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.90μs | 3.1K | 65 |
| Small Struct | 🥉 JSON | Unmarshal | 20.10μs | 7.4K | 97 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.59μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.28μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.19μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.96μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 31.66μs | 16.6K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 36.69μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 14.86μs | 24.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.41μs | 36.6K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 38.22μs | 32.3K | 596 |
| Medium Payload | 🥈 CBOR | Unmarshal | 45.00μs | 29.8K | 613 |
| Medium Payload | 🥉 JSON | Unmarshal | 185.18μs | 62.3K | 814 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 63.16μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 89.77μs | 196.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 206.12μs | 526.8K | 115 |
| Large Payload | 🥈 CBOR | Marshal | 290.42μs | 188.5K | 1 |
| Large Payload | 🥉 JSON | Marshal | 326.65μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 416.66μs | 213.9K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 178.96μs | 280.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 298.68μs | 348.1K | 209 |
| Large Payload | 🥈 CBOR | Unmarshal | 668.90μs | 318.4K | 6.5K |
| Large Payload | 🥈 MessagePack | Unmarshal | 722.65μs | 356.5K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.99ms | 537.5K | 7.1K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 422ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 805ns | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.09μs | 1.5K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.14μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.54μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.76μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 698ns | 376 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.60μs | 1.6K | 35 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.99μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.18μs | 2.1K | 47 |
| Small Struct | 🥉 JSON | Unmarshal | 21.10μs | 4.8K | 88 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.78μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.52μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.12μs | 22.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.28μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.91μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.16μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.15μs | 24.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.09μs | 47.6K | 65 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.68μs | 40.1K | 748 |
| Medium Payload | 🥈 CBOR | Unmarshal | 68.01μs | 31.1K | 642 |
| Medium Payload | 🥉 JSON | Unmarshal | 205.29μs | 50.5K | 651 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.66μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.71μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 157.06μs | 215.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 199.91μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 309.82μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 467.73μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 239.07μs | 275.9K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 376.21μs | 572.8K | 589 |
| Large Payload | 🥈 MessagePack | Unmarshal | 548.12μs | 330.4K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 715.69μs | 318.5K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.39ms | 562.1K | 7.4K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 338ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 887ns | 1.5K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.52μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.32μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 2.73μs | 2.1K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.94μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.77μs | 3.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.40μs | 2.7K | 56 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.40μs | 5.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.77μs | 4.8K | 102 |
| Small Struct | 🥉 JSON | Unmarshal | 18.46μs | 4.8K | 86 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.23μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.12μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.64μs | 18.4K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 30.96μs | 16.6K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.47μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 34.03μs | 25.1K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.64μs | 32.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.53μs | 33.6K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 53.34μs | 24.1K | 499 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.72μs | 40.3K | 751 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.57μs | 62.3K | 828 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.29μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.33μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 183.29μs | 188.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 268.48μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 310.17μs | 222.8K | 3 |
| Large Payload | 🥉 JSON | Marshal | 406.52μs | 229.8K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 214.53μs | 257.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 296.50μs | 421.9K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 525.82μs | 360.9K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 647.84μs | 313.1K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.87ms | 507.1K | 6.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 457ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 727ns | 556 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.02μs | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.57μs | 1.0K | 6 |
| Small Struct | 🥉 JSON | Marshal | 2.56μs | 896 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.62μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.92μs | 1.3K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 3.22μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.36μs | 3.2K | 68 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.28μs | 3.5K | 74 |
| Small Struct | 🥉 JSON | Unmarshal | 23.06μs | 4.4K | 74 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.42μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.91μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 21.86μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.76μs | 22.1K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 50.83μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 56.83μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 42.22μs | 34.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.08μs | 48.8K | 69 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.40μs | 29.7K | 540 |
| Medium Payload | 🥈 CBOR | Unmarshal | 94.79μs | 33.8K | 700 |
| Medium Payload | 🥉 JSON | Unmarshal | 311.67μs | 61.4K | 836 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 86.68μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.23μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 165.94μs | 215.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 258.38μs | 213.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 285.36μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 479.35μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 302.65μs | 299.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 527.04μs | 555.8K | 590 |
| Large Payload | 🥈 MessagePack | Unmarshal | 786.18μs | 354.4K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 969.28μs | 310.5K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.89ms | 529.4K | 6.9K |

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

