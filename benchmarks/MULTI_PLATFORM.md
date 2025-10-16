# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 17:11:48 UTC

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
| Apple M1 (Virtual) | 2.76μs | 471ns | 3.86μs | 2.09μs | 2.34μs |
| AMD EPYC 7763 64-Core Processor | 1.20μs | 566ns | 2.74μs | 2.88μs | 2.46μs |
| Neoverse-N2 | 765ns | 432ns | 2.40μs | 1.43μs | 3.86μs |
| Unknown CPU | 928ns | 568ns | 1.44μs | 2.60μs | 1.29μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 2.08μs | 24.50μs | 9.37μs | 1.68μs |
| AMD EPYC 7763 64-Core Processor | 1.78μs | 5.80μs | 2.19μs | 3.11μs |
| Neoverse-N2 | 778ns | 11.01μs | 1.30μs | 4.87μs |
| Unknown CPU | 3.74μs | 32.21μs | 6.53μs | 6.95μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (471ns) | 🥈 MessagePack (1.68μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (566ns) | 🥉 Sonic (1.68μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (432ns) | 🥇 BEVE (778ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (568ns) | 🥇 BEVE (3.74μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 47.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 471ns | 288 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.09μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.34μs | 4.2K | 8 |
| Small Struct | 🥇 BEVE | Marshal | 2.76μs | 3.0K | 3 |
| Small Struct | 🥉 JSON | Marshal | 3.86μs | 1.9K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 4.71μs | 2.5K | 3 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.68μs | 680 | 17 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.08μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.49μs | 4.8K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.37μs | 3.9K | 84 |
| Small Struct | 🥉 JSON | Unmarshal | 24.50μs | 7.2K | 91 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.93μs | 134 | 2 |
| Medium Payload | 🥈 CBOR | Marshal | 16.58μs | 21.8K | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 16.76μs | 20.6K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.30μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.70μs | 19.4K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 47.85μs | 25.0K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.71μs | 29.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 35.50μs | 46.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 37.88μs | 36.0K | 668 |
| Medium Payload | 🥈 CBOR | Unmarshal | 51.32μs | 33.4K | 685 |
| Medium Payload | 🥉 JSON | Unmarshal | 208.33μs | 62.1K | 808 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 59.17μs | 180 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 116.54μs | 189.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 250.78μs | 206.2K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 331.59μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 474.83μs | 221.8K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 528.83μs | 215.7K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 264.98μs | 272.9K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 406.90μs | 381.0K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 668.98μs | 356.5K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 671.02μs | 286.5K | 5.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.39ms | 560.0K | 7.3K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 566ns | 289 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.20μs | 1.8K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 1.33μs | 2.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.46μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.74μs | 1.4K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.88μs | 2.8K | 2 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.68μs | 2.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.78μs | 3.4K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.19μs | 712 | 18 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.11μs | 2.1K | 45 |
| Small Struct | 🥉 JSON | Unmarshal | 5.80μs | 1.3K | 26 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.47μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 14.28μs | 22.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.98μs | 18.6K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 20.17μs | 21.2K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 28.42μs | 33.1K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 42.26μs | 18.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.27μs | 30.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 42.52μs | 66.2K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.60μs | 34.4K | 629 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.80μs | 31.5K | 646 |
| Medium Payload | 🥉 JSON | Unmarshal | 208.87μs | 51.7K | 660 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.11μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 116.19μs | 188.9K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 164.96μs | 216.9K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 223.08μs | 213.5K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 302.23μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 455.15μs | 221.9K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 224.85μs | 252.9K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 380.21μs | 579.3K | 601 |
| Large Payload | 🥈 MessagePack | Unmarshal | 595.64μs | 375.2K | 6.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 802.45μs | 324.9K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.30ms | 538.1K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 432ns | 288 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 765ns | 928 | 3 |
| Small Struct | 🥈 CBOR | Marshal | 1.43μs | 1.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.40μs | 1.4K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 3.83μs | 2.9K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 3.86μs | 8.3K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 778ns | 568 | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.30μs | 320 | 10 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.17μs | 5.3K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.87μs | 4.0K | 86 |
| Small Struct | 🥉 JSON | Unmarshal | 11.01μs | 3.7K | 52 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.42μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 12.70μs | 24.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.39μs | 19.2K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.80μs | 33.1K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 29.13μs | 22.2K | 4 |
| Medium Payload | 🥉 JSON | Marshal | 43.25μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.34μs | 28.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.15μs | 32.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.41μs | 39.8K | 744 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.12μs | 30.2K | 628 |
| Medium Payload | 🥉 JSON | Unmarshal | 182.81μs | 50.7K | 650 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 71.07μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 108.50μs | 189.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 185.38μs | 190.7K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 262.17μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 316.53μs | 237.8K | 4 |
| Large Payload | 🥉 JSON | Marshal | 370.89μs | 207.8K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 212.16μs | 253.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 289.97μs | 401.7K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 513.37μs | 345.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 630.01μs | 305.0K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.00ms | 539.8K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 568ns | 289 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 602ns | 522 | 3 |
| Small Struct | 🥇 BEVE | Marshal | 928ns | 800 | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.29μs | 1.2K | 6 |
| Small Struct | 🥉 JSON | Marshal | 1.44μs | 624 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.60μs | 2.5K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 3.74μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.79μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.53μs | 3.0K | 64 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.95μs | 3.9K | 81 |
| Small Struct | 🥉 JSON | Unmarshal | 32.21μs | 8.0K | 115 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.54μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.80μs | 19.2K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 17.82μs | 22.1K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 23.24μs | 19.2K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.58μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.59μs | 18.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.45μs | 29.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 51.54μs | 63.8K | 78 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 75.55μs | 40.2K | 749 |
| Medium Payload | 🥈 CBOR | Unmarshal | 80.60μs | 30.7K | 636 |
| Medium Payload | 🥉 JSON | Unmarshal | 242.07μs | 51.6K | 668 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.05μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 116.49μs | 189.1K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 162.12μs | 228.3K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 225.01μs | 190.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 288.05μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 487.55μs | 215.8K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 271.59μs | 271.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 435.01μs | 553.6K | 588 |
| Large Payload | 🥈 MessagePack | Unmarshal | 640.12μs | 324.9K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 874.98μs | 333.1K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.64ms | 566.6K | 7.3K |

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

