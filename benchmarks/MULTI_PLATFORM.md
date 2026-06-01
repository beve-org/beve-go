# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-06-01 08:30:34 UTC

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
| Apple M1 (Virtual) | 1.21μs | 387ns | 2.80μs | 2.04μs | 2.87μs |
| AMD EPYC 7763 64-Core Processor | 871ns | 574ns | 3.10μs | 3.12μs | 3.31μs |
| Neoverse-N2 | 897ns | 298ns | 3.15μs | 2.00μs | 992ns |
| Unknown CPU | 2.12μs | 964ns | 2.33μs | 3.06μs | 4.68μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.53μs | 15.25μs | 3.79μs | 4.11μs |
| AMD EPYC 7763 64-Core Processor | 2.31μs | 35.42μs | 4.88μs | 7.71μs |
| Neoverse-N2 | 1.21μs | 25.01μs | 5.39μs | 5.33μs |
| Unknown CPU | 2.56μs | 8.64μs | 6.41μs | 7.87μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (387ns) | 🥇 BEVE (1.53μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (574ns) | 🥇 BEVE (2.31μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (298ns) | 🥉 Sonic (1.07μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (964ns) | 🥉 Sonic (799ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 52.3% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 387ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.21μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.60μs | 407 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.04μs | 3.1K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.80μs | 1.5K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.87μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.53μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.30μs | 5.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.79μs | 1.9K | 43 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.11μs | 3.6K | 76 |
| Small Struct | 🥉 JSON | Unmarshal | 15.25μs | 4.1K | 64 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.39μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.17μs | 27.3K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.67μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 24.85μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 44.48μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 48.06μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.70μs | 29.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.11μs | 28.9K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 46.66μs | 32.4K | 589 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.58μs | 39.2K | 806 |
| Medium Payload | 🥉 JSON | Unmarshal | 204.91μs | 54.3K | 684 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 64.80μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 112.78μs | 213.0K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 206.18μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 265.24μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 383.84μs | 196.9K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 554.92μs | 222.2K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 199.01μs | 252.2K | 415 |
| Large Payload | 🥉 Sonic | Unmarshal | 367.84μs | 374.6K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 467.29μs | 356.6K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 579.76μs | 276.5K | 5.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.30ms | 546.8K | 7.1K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 574ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 822ns | 755 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 871ns | 1.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.10μs | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 3.12μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.31μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.31μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.35μs | 3.8K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.88μs | 1.8K | 40 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.71μs | 4.7K | 99 |
| Small Struct | 🥉 JSON | Unmarshal | 35.42μs | 8.0K | 116 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.41μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.89μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.94μs | 25.4K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 26.60μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 42.92μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.34μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 30.22μs | 34.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.26μs | 46.4K | 63 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.32μs | 35.6K | 656 |
| Medium Payload | 🥈 CBOR | Unmarshal | 74.83μs | 33.1K | 680 |
| Medium Payload | 🥉 JSON | Unmarshal | 204.08μs | 46.6K | 620 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.51μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 127.48μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 171.57μs | 224.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 211.93μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 329.96μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 445.06μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 237.07μs | 259.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 378.39μs | 545.3K | 584 |
| Large Payload | 🥈 MessagePack | Unmarshal | 558.21μs | 336.3K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 740.80μs | 331.9K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.43ms | 579.7K | 7.5K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 298ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 778ns | 414 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 897ns | 1.5K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 992ns | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 2.00μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.15μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.07μs | 1.0K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.21μs | 1.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.33μs | 4.6K | 96 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.39μs | 3.1K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 25.01μs | 7.9K | 114 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.52μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.06μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.84μs | 18.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 26.68μs | 18.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.54μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.83μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.76μs | 23.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 30.93μs | 44.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.18μs | 35.7K | 660 |
| Medium Payload | 🥈 CBOR | Unmarshal | 56.02μs | 25.8K | 529 |
| Medium Payload | 🥉 JSON | Unmarshal | 215.43μs | 61.9K | 802 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.04μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 99.97μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 184.86μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 274.23μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 309.39μs | 223.3K | 3 |
| Large Payload | 🥉 JSON | Marshal | 393.54μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 225.19μs | 276.9K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 268.82μs | 351.5K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 530.85μs | 365.9K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 678.39μs | 333.3K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.03ms | 550.0K | 7.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 964ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.07μs | 1.2K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.12μs | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.33μs | 1.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 3.06μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.68μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Unmarshal | 799ns | 389 | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.56μs | 3.0K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.41μs | 2.4K | 52 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.87μs | 4.3K | 92 |
| Small Struct | 🥉 JSON | Unmarshal | 8.64μs | 1.4K | 29 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.31μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.78μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.05μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 29.31μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 40.12μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 62.28μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 35.33μs | 33.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 55.20μs | 60.6K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 82.42μs | 36.9K | 685 |
| Medium Payload | 🥈 CBOR | Unmarshal | 109.37μs | 36.1K | 738 |
| Medium Payload | 🥉 JSON | Unmarshal | 290.36μs | 56.4K | 709 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.54μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 137.10μs | 204.8K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 181.42μs | 215.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 239.11μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 320.52μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 519.13μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 332.89μs | 270.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 485.67μs | 571.3K | 598 |
| Large Payload | 🥈 CBOR | Unmarshal | 884.92μs | 332.1K | 6.8K |
| Large Payload | 🥈 MessagePack | Unmarshal | 1.35ms | 341.8K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 3.13ms | 567.3K | 7.4K |

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

