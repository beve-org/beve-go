# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-17 14:39:04 UTC

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
| Apple M1 (Virtual) | 756ns | 550ns | 5.98μs | 840ns | 2.71μs |
| AMD EPYC 7763 64-Core Processor | 1.45μs | 762ns | 3.35μs | 630ns | 1.61μs |
| Neoverse-N2 | 694ns | 388ns | 4.78μs | 2.40μs | 2.29μs |
| Unknown CPU | 1.99μs | 652ns | 3.69μs | 2.95μs | 3.88μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 2.07μs | 12.53μs | 5.96μs | 2.80μs |
| AMD EPYC 7763 64-Core Processor | 1.26μs | 17.59μs | 2.56μs | 5.21μs |
| Neoverse-N2 | 805ns | 8.07μs | 7.93μs | 5.69μs |
| Unknown CPU | 1.65μs | 27.29μs | 8.27μs | 6.82μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (550ns) | 🥇 BEVE (2.07μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥈 CBOR (630ns) | 🥉 Sonic (1.22μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (388ns) | 🥇 BEVE (805ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (652ns) | 🥇 BEVE (1.65μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 68.9% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 550ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 756ns | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 840ns | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.71μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 4.28μs | 1.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 5.98μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.07μs | 3.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.80μs | 1.9K | 42 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.51μs | 4.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.96μs | 3.2K | 69 |
| Small Struct | 🥉 JSON | Unmarshal | 12.53μs | 3.8K | 56 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.79μs | 0 | 0 |
| Medium Payload | 🥈 CBOR | Marshal | 14.28μs | 16.4K | 1 |
| Medium Payload | 🥇 BEVE | Marshal | 17.42μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.68μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 39.26μs | 19.3K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 40.86μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 18.14μs | 25.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.54μs | 37.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.97μs | 40.5K | 758 |
| Medium Payload | 🥈 CBOR | Unmarshal | 48.31μs | 29.9K | 614 |
| Medium Payload | 🥉 JSON | Unmarshal | 227.05μs | 72.2K | 934 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 75.64μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 121.24μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 156.99μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 263.22μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 442.97μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 487.16μs | 222.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 244.61μs | 266.2K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 394.04μs | 353.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 418.82μs | 349.0K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 730.14μs | 325.9K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 1.79ms | 495.8K | 6.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥈 CBOR | Marshal | 630ns | 448 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 762ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.30μs | 1.9K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.45μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.61μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 3.35μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.22μs | 1.3K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.26μs | 1.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.56μs | 1.1K | 25 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.21μs | 4.0K | 84 |
| Small Struct | 🥉 JSON | Unmarshal | 17.59μs | 4.4K | 75 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.69μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.28μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.44μs | 22.2K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.46μs | 16.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.09μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.31μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.38μs | 30.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 35.55μs | 50.9K | 73 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 57.36μs | 36.9K | 686 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.32μs | 29.7K | 606 |
| Medium Payload | 🥉 JSON | Unmarshal | 278.78μs | 72.1K | 922 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.28μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.96μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 156.46μs | 215.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 207.54μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 309.17μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 443.26μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.71μs | 273.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 342.14μs | 512.1K | 566 |
| Large Payload | 🥈 MessagePack | Unmarshal | 565.80μs | 350.4K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 705.52μs | 316.2K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.27ms | 538.2K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 388ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 694ns | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.29μs | 4.1K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.40μs | 3.1K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.22μs | 2.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.78μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 805ns | 600 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.66μs | 6.3K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.69μs | 4.8K | 101 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.93μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 8.07μs | 2.2K | 39 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.81μs | 7 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.34μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.89μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 25.96μs | 18.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.98μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.51μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.15μs | 30.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.07μs | 44.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.44μs | 43.6K | 818 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.42μs | 31.2K | 634 |
| Medium Payload | 🥉 JSON | Unmarshal | 155.83μs | 40.5K | 529 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.27μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 103.25μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 170.26μs | 172.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 274.55μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 317.78μs | 223.6K | 3 |
| Large Payload | 🥉 JSON | Marshal | 380.40μs | 205.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.09μs | 270.5K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 286.18μs | 393.3K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 527.26μs | 353.3K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 637.67μs | 306.8K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.10ms | 576.1K | 7.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 652ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.99μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.50μs | 2.7K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.95μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.69μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.88μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.65μs | 1.8K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.97μs | 7.8K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.82μs | 3.9K | 82 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.27μs | 3.6K | 78 |
| Small Struct | 🥉 JSON | Unmarshal | 27.29μs | 7.3K | 94 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.51μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.86μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.02μs | 18.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.70μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.17μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 53.81μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.22μs | 23.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.64μs | 54.1K | 72 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 76.50μs | 40.7K | 760 |
| Medium Payload | 🥈 CBOR | Unmarshal | 95.60μs | 35.5K | 726 |
| Medium Payload | 🥉 JSON | Unmarshal | 256.48μs | 55.5K | 722 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.18μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.43μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 166.37μs | 215.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 222.71μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 271.41μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 459.66μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 267.33μs | 262.3K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 415.23μs | 538.3K | 581 |
| Large Payload | 🥈 MessagePack | Unmarshal | 682.35μs | 354.9K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 824.11μs | 301.8K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 2.60ms | 535.6K | 7.0K |

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

