# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-01-19 04:12:32 UTC

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
| Apple M1 (Virtual) | 1.91μs | 526ns | 2.18μs | 428ns | 1.45μs |
| AMD EPYC 7763 64-Core Processor | 2.06μs | 980ns | 6.51μs | 3.28μs | 1.89μs |
| Neoverse-N2 | 818ns | 237ns | 3.94μs | 1.30μs | 2.57μs |
| Unknown CPU | 628ns | 349ns | 6.08μs | 2.42μs | 3.03μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.16μs | 12.04μs | 5.56μs | 4.64μs |
| AMD EPYC 7763 64-Core Processor | 1.92μs | 8.60μs | 6.89μs | 4.49μs |
| Neoverse-N2 | 1.68μs | 16.20μs | 7.41μs | 2.00μs |
| Unknown CPU | 762ns | 18.49μs | 1.41μs | 5.08μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥈 CBOR (428ns) | 🥇 BEVE (1.16μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥉 Sonic (636ns) | 🥇 BEVE (1.92μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (237ns) | 🥇 BEVE (1.68μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (349ns) | 🥇 BEVE (762ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 62.4% faster

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
| Small Struct | 🥈 CBOR | Marshal | 428ns | 352 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 526ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.45μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 1.91μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.18μs | 1.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.69μs | 1.3K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.16μs | 1.7K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.81μs | 1.2K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.64μs | 3.6K | 76 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.56μs | 3.2K | 68 |
| Small Struct | 🥉 JSON | Unmarshal | 12.04μs | 3.9K | 59 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.97μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.07μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 28.34μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.75μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 44.50μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 53.64μs | 22.0K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.34μs | 32.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 49.53μs | 42.8K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 54.04μs | 27.9K | 502 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.02μs | 25.7K | 530 |
| Medium Payload | 🥉 JSON | Unmarshal | 229.87μs | 51.0K | 659 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 72.88μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 138.50μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 199.43μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 369.38μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 489.05μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 632.26μs | 206.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 296.35μs | 261.5K | 415 |
| Large Payload | 🥉 Sonic | Unmarshal | 412.27μs | 353.5K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 638.74μs | 352.4K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 844.28μs | 315.8K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.53ms | 546.8K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 636ns | 350 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 980ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.89μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 2.06μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 3.28μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 6.51μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.92μs | 2.1K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.54μs | 4.2K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.49μs | 2.8K | 60 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.89μs | 2.9K | 63 |
| Small Struct | 🥉 JSON | Unmarshal | 8.60μs | 2.1K | 36 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.21μs | 0 | 0 |
| Medium Payload | 🥉 Sonic | Marshal | 15.11μs | 20.9K | 3 |
| Medium Payload | 🥇 BEVE | Marshal | 17.37μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 21.27μs | 19.1K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 36.72μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.07μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.02μs | 27.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.14μs | 57.7K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 64.19μs | 38.8K | 729 |
| Medium Payload | 🥈 CBOR | Unmarshal | 78.12μs | 36.8K | 755 |
| Medium Payload | 🥉 JSON | Unmarshal | 227.68μs | 52.8K | 716 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.53μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.49μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 162.94μs | 223.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 202.44μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 338.00μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 436.83μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 245.81μs | 277.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 343.56μs | 508.2K | 563 |
| Large Payload | 🥈 MessagePack | Unmarshal | 588.74μs | 359.1K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 739.74μs | 323.6K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.29ms | 514.9K | 6.7K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 237ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 754ns | 407 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 818ns | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.30μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.57μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.94μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.68μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.00μs | 1.2K | 27 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.00μs | 5.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.41μs | 4.4K | 95 |
| Small Struct | 🥉 JSON | Unmarshal | 16.20μs | 4.4K | 75 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.24μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.15μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.78μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 26.43μs | 19.0K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.94μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 36.17μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.80μs | 27.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.53μs | 35.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.80μs | 32.8K | 602 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.34μs | 31.3K | 648 |
| Medium Payload | 🥉 JSON | Unmarshal | 222.72μs | 63.1K | 828 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.81μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.52μs | 196.8K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 179.84μs | 180.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 274.86μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 326.60μs | 235.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 410.33μs | 221.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 234.29μs | 293.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 294.90μs | 404.7K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 526.42μs | 349.9K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 651.19μs | 308.3K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.08ms | 569.2K | 7.4K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 349ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 628ns | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.09μs | 2.8K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.42μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.03μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 6.08μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 762ns | 344 | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.41μs | 320 | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.08μs | 3.3K | 71 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.18μs | 7.0K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 18.49μs | 4.2K | 68 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.34μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.50μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.35μs | 19.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.48μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.38μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 45.82μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.20μs | 33.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.10μs | 54.3K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 69.26μs | 35.8K | 666 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.34μs | 27.2K | 560 |
| Medium Payload | 🥉 JSON | Unmarshal | 271.36μs | 57.6K | 768 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.54μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.57μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 160.86μs | 222.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 232.86μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 286.48μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 523.05μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.53μs | 273.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 432.61μs | 543.8K | 574 |
| Large Payload | 🥈 MessagePack | Unmarshal | 681.53μs | 345.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 846.92μs | 302.8K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.59ms | 519.1K | 6.8K |

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

