# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-06-08 07:28:03 UTC

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
| Apple M1 (Virtual) | 1.04μs | 360ns | 3.68μs | 2.70μs | 3.70μs |
| AMD EPYC 7763 64-Core Processor | 1.29μs | 465ns | 910ns | 1.38μs | 3.49μs |
| Neoverse-N2 | 1.13μs | 494ns | 2.79μs | 775ns | 2.19μs |
| Unknown CPU | 864ns | 379ns | 3.38μs | 1.40μs | 3.10μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 952ns | 10.21μs | 5.30μs | 4.32μs |
| AMD EPYC 7763 64-Core Processor | 1.11μs | 3.93μs | 2.02μs | 3.60μs |
| Neoverse-N2 | 1.01μs | 12.81μs | 2.47μs | 3.54μs |
| Unknown CPU | 1.86μs | 27.08μs | 11.12μs | 5.32μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (360ns) | 🥇 BEVE (952ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (465ns) | 🥇 BEVE (1.11μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (494ns) | 🥇 BEVE (1.01μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (379ns) | 🥇 BEVE (1.86μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 40.9% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 360ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.04μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.70μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.68μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.70μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 5.53μs | 2.7K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 952ns | 1.3K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.32μs | 3.2K | 69 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.69μs | 5.6K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.30μs | 3.0K | 64 |
| Small Struct | 🥉 JSON | Unmarshal | 10.21μs | 2.4K | 46 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.34μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.61μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 25.05μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.50μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.17μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 49.35μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.63μs | 31.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 52.33μs | 45.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.82μs | 33.8K | 621 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.09μs | 28.5K | 583 |
| Medium Payload | 🥉 JSON | Unmarshal | 201.50μs | 48.5K | 676 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.36μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 127.75μs | 204.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 215.84μs | 213.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 334.69μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 485.03μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 553.78μs | 222.3K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 263.90μs | 264.1K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 373.87μs | 349.2K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 563.07μs | 328.9K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 683.78μs | 312.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.18ms | 556.5K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 465ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 583ns | 606 | 2 |
| Small Struct | 🥉 JSON | Marshal | 910ns | 352 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.29μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.38μs | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.49μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.11μs | 1.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.02μs | 760 | 19 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.60μs | 2.5K | 54 |
| Small Struct | 🥉 JSON | Unmarshal | 3.93μs | 680 | 18 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.25μs | 7.4K | 10 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.69μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.83μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.70μs | 20.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.02μs | 16.4K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 37.31μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.29μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.81μs | 31.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.01μs | 46.6K | 66 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 54.99μs | 34.0K | 624 |
| Medium Payload | 🥈 CBOR | Unmarshal | 61.08μs | 25.5K | 529 |
| Medium Payload | 🥉 JSON | Unmarshal | 265.00μs | 62.3K | 849 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.49μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.43μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 152.11μs | 207.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 210.49μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 322.01μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 433.22μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.09μs | 272.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 388.43μs | 587.0K | 603 |
| Large Payload | 🥈 MessagePack | Unmarshal | 595.89μs | 368.3K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 667.57μs | 290.0K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.35ms | 536.1K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 494ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 775ns | 640 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.13μs | 2.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.17μs | 734 | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.19μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.79μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.01μs | 1.2K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.47μs | 1.1K | 26 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.21μs | 5.3K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.54μs | 2.9K | 62 |
| Small Struct | 🥉 JSON | Unmarshal | 12.81μs | 3.9K | 60 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.30μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.98μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.12μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 26.07μs | 18.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.89μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 37.99μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.97μs | 30.4K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.87μs | 45.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.89μs | 40.2K | 749 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.57μs | 32.4K | 664 |
| Medium Payload | 🥉 JSON | Unmarshal | 190.28μs | 49.8K | 697 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.50μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 107.19μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 189.28μs | 197.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 271.37μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 302.43μs | 209.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 391.94μs | 213.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 217.18μs | 257.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 288.04μs | 397.7K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 506.83μs | 338.7K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 657.14μs | 318.8K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.99ms | 540.8K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 379ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 864ns | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.40μs | 1.4K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.48μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.10μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.38μs | 1.4K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.86μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.45μs | 2.4K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.32μs | 3.2K | 68 |
| Small Struct | 🥈 CBOR | Unmarshal | 11.12μs | 5.2K | 107 |
| Small Struct | 🥉 JSON | Unmarshal | 27.08μs | 7.4K | 98 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.88μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.48μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.64μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.32μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 24.96μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 49.88μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.37μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 52.55μs | 66.0K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 73.76μs | 38.2K | 711 |
| Medium Payload | 🥈 CBOR | Unmarshal | 86.80μs | 30.3K | 623 |
| Medium Payload | 🥉 JSON | Unmarshal | 282.42μs | 59.2K | 765 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 83.01μs | 92 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.92μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 160.74μs | 214.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 212.03μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 272.86μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 473.65μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.07μs | 272.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 428.34μs | 545.2K | 585 |
| Large Payload | 🥈 MessagePack | Unmarshal | 651.72μs | 337.0K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 881.47μs | 329.4K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.38ms | 484.1K | 6.4K |

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

