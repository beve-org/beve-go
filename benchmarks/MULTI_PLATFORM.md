# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 12:32:52 UTC

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
| Apple M1 (Virtual) | 1.04μs | 820ns | 699ns | 1.04μs | 3.15μs |
| AMD EPYC 7763 64-Core Processor | 1.66μs | 689ns | 4.22μs | 3.20μs | 1.66μs |
| Neoverse-N2 | 1.18μs | 539ns | 2.24μs | 2.02μs | 3.53μs |
| Unknown CPU | 2.73μs | 443ns | 5.14μs | 680ns | 5.08μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.34μs | 5.52μs | 1.06μs | 2.31μs |
| AMD EPYC 7763 64-Core Processor | 1.05μs | 6.52μs | 6.07μs | 5.23μs |
| Neoverse-N2 | 1.74μs | 17.74μs | 1.74μs | 4.11μs |
| Unknown CPU | 2.14μs | 8.08μs | 7.66μs | 5.65μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥉 JSON (699ns) | 🥈 CBOR (1.06μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (689ns) | 🥇 BEVE (1.05μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (539ns) | 🥉 Sonic (1.23μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (443ns) | 🥇 BEVE (2.14μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 26.5% faster

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
| Small Struct | 🥉 JSON | Marshal | 699ns | 496 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 820ns | 288 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.04μs | 1.3K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.04μs | 2.1K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 3.15μs | 8.3K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 5.28μs | 2.9K | 3 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.06μs | 432 | 12 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.34μs | 1.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.30μs | 3.3K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.31μs | 2.3K | 51 |
| Small Struct | 🥉 JSON | Unmarshal | 5.52μs | 1.4K | 30 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.69μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 11.83μs | 18.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 29.75μs | 24.7K | 2 |
| Medium Payload | 🥉 JSON | Marshal | 30.40μs | 18.7K | 9 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.13μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 55.58μs | 25.0K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.20μs | 27.5K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 38.36μs | 43.7K | 33 |
| Medium Payload | 🥈 CBOR | Unmarshal | 55.92μs | 26.5K | 550 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 65.33μs | 41.5K | 775 |
| Medium Payload | 🥉 JSON | Unmarshal | 237.78μs | 58.8K | 754 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 72.60μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 135.99μs | 189.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 167.11μs | 205.5K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 211.85μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 397.45μs | 222.1K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 481.85μs | 223.9K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 223.11μs | 265.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 280.16μs | 333.2K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 451.55μs | 363.6K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 804.73μs | 329.5K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 1.76ms | 532.8K | 6.9K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 689ns | 290 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.66μs | 3.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.66μs | 2.2K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.71μs | 2.3K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 3.20μs | 3.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.22μs | 2.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.05μs | 1.3K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.64μs | 3.8K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.23μs | 3.8K | 80 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.07μs | 3.1K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 6.52μs | 1.4K | 29 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.59μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 12.95μs | 19.3K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 14.36μs | 19.3K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 19.74μs | 18.6K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.67μs | 33.1K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 48.93μs | 25.0K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.12μs | 30.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.53μs | 45.5K | 65 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.68μs | 29.8K | 539 |
| Medium Payload | 🥈 CBOR | Unmarshal | 83.03μs | 38.5K | 786 |
| Medium Payload | 🥉 JSON | Unmarshal | 231.27μs | 54.5K | 726 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.19μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 117.47μs | 188.9K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 164.42μs | 208.7K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 221.46μs | 197.7K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 342.69μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 463.37μs | 222.1K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 242.66μs | 265.8K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 380.46μs | 548.6K | 586 |
| Large Payload | 🥈 MessagePack | Unmarshal | 578.95μs | 342.6K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 803.01μs | 345.1K | 7.0K |
| Large Payload | 🥉 JSON | Unmarshal | 2.34ms | 526.7K | 6.9K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 539ns | 290 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.18μs | 2.1K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.02μs | 2.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.24μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.53μs | 8.3K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 3.89μs | 3.3K | 3 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.23μs | 1.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.74μs | 616 | 16 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.74μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.11μs | 3.3K | 71 |
| Small Struct | 🥉 JSON | Unmarshal | 17.74μs | 4.6K | 82 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.28μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 11.58μs | 21.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.05μs | 19.2K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 29.64μs | 22.3K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.35μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 35.79μs | 20.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.48μs | 33.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.68μs | 33.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.59μs | 34.2K | 631 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.33μs | 34.0K | 698 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.54μs | 56.3K | 718 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.78μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 111.33μs | 182.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 182.20μs | 190.7K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 278.34μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 289.57μs | 209.9K | 4 |
| Large Payload | 🥉 JSON | Marshal | 396.16μs | 223.6K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.20μs | 277.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 285.42μs | 374.9K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 517.09μs | 354.9K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 673.68μs | 333.7K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 563.6K | 7.4K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 443ns | 290 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 680ns | 528 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.73μs | 3.4K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 2.86μs | 2.9K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 5.08μs | 8.3K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.14μs | 2.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.14μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.60μs | 3.9K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.65μs | 3.3K | 70 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.66μs | 3.2K | 69 |
| Small Struct | 🥉 JSON | Unmarshal | 8.08μs | 1.4K | 31 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.30μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 14.19μs | 19.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.72μs | 16.5K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 21.66μs | 27.9K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.32μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.94μs | 22.1K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 32.84μs | 30.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.12μs | 63.7K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 77.49μs | 38.8K | 726 |
| Medium Payload | 🥈 CBOR | Unmarshal | 89.58μs | 32.4K | 665 |
| Medium Payload | 🥉 JSON | Unmarshal | 210.94μs | 39.3K | 536 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.28μs | 286 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 124.31μs | 205.6K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 177.10μs | 227.8K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 219.14μs | 198.2K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 298.89μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 484.99μs | 214.9K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 288.82μs | 278.5K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 438.40μs | 571.0K | 591 |
| Large Payload | 🥈 MessagePack | Unmarshal | 690.59μs | 353.4K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 902.94μs | 333.3K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.64ms | 554.2K | 7.2K |

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

