# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-11-10 03:53:09 UTC

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
| Apple M1 (Virtual) | 1.15μs | 209ns | 398ns | 729ns | 3.66μs |
| AMD EPYC 7763 64-Core Processor | 1.22μs | 832ns | 3.63μs | 2.24μs | 2.43μs |
| Neoverse-N2 | 920ns | 645ns | 829ns | 1.06μs | 1.66μs |
| Unknown CPU | 469ns | 804ns | 4.87μs | 440ns | 2.87μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.00μs | 14.95μs | 3.71μs | 4.77μs |
| AMD EPYC 7763 64-Core Processor | 1.24μs | 26.82μs | 5.11μs | 5.32μs |
| Neoverse-N2 | 1.32μs | 16.03μs | 7.69μs | 4.68μs |
| Unknown CPU | 1.82μs | 3.29μs | 9.10μs | 1.22μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (209ns) | 🥇 BEVE (1.00μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (832ns) | 🥇 BEVE (1.24μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (645ns) | 🥇 BEVE (1.32μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥈 CBOR (440ns) | 🥈 MessagePack (1.22μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** -10.6% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 209ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 398ns | 224 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 639ns | 279 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 729ns | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.15μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.66μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.00μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.52μs | 2.1K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.71μs | 1.6K | 37 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.77μs | 5.2K | 105 |
| Small Struct | 🥉 JSON | Unmarshal | 14.95μs | 4.6K | 79 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.81μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.26μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 20.09μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 20.34μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 32.11μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 46.92μs | 24.9K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.36μs | 29.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 31.97μs | 37.7K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.07μs | 37.2K | 688 |
| Medium Payload | 🥈 CBOR | Unmarshal | 50.18μs | 26.7K | 552 |
| Medium Payload | 🥉 JSON | Unmarshal | 169.76μs | 49.3K | 634 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.71μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 95.25μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 160.76μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 199.08μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 384.62μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 452.69μs | 213.9K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 180.27μs | 274.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 331.81μs | 356.7K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 453.57μs | 344.8K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 622.50μs | 326.9K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 1.91ms | 527.4K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 832ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.22μs | 2.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.87μs | 2.8K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.24μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.43μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.63μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.24μs | 1.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.86μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.11μs | 2.8K | 61 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.32μs | 4.1K | 87 |
| Small Struct | 🥉 JSON | Unmarshal | 26.82μs | 7.8K | 110 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.50μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.50μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.72μs | 27.8K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.76μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.20μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 44.12μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.09μs | 32.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.55μs | 64.2K | 78 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.50μs | 33.7K | 625 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.95μs | 27.6K | 570 |
| Medium Payload | 🥉 JSON | Unmarshal | 273.80μs | 67.7K | 891 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 83.70μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 112.41μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 159.21μs | 223.6K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 211.56μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 318.94μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 440.61μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 230.24μs | 255.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 353.73μs | 544.2K | 576 |
| Large Payload | 🥈 MessagePack | Unmarshal | 601.44μs | 374.3K | 6.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 708.51μs | 311.0K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.26ms | 530.4K | 6.9K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 645ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 829ns | 352 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 920ns | 1.5K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.06μs | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.66μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 3.39μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.32μs | 1.8K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.97μs | 2.7K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.68μs | 3.9K | 81 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.69μs | 4.7K | 99 |
| Small Struct | 🥉 JSON | Unmarshal | 16.03μs | 4.4K | 74 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.79μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.89μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.01μs | 18.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 23.43μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 32.98μs | 25.0K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 36.90μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.97μs | 26.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.73μs | 34.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.15μs | 34.6K | 642 |
| Medium Payload | 🥈 CBOR | Unmarshal | 66.94μs | 32.7K | 674 |
| Medium Payload | 🥉 JSON | Unmarshal | 187.90μs | 49.5K | 684 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.47μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.96μs | 213.2K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 189.83μs | 197.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 299.82μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 304.06μs | 217.4K | 3 |
| Large Payload | 🥉 JSON | Marshal | 396.28μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.12μs | 266.5K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 295.25μs | 399.1K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 518.68μs | 342.5K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 666.76μs | 319.0K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.03ms | 544.9K | 7.2K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥈 CBOR | Marshal | 440ns | 224 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 469ns | 384 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 804ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.84μs | 2.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.87μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.87μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.22μs | 304 | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.82μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.84μs | 2.1K | 8 |
| Small Struct | 🥉 JSON | Unmarshal | 3.29μs | 496 | 13 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.10μs | 4.4K | 93 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.84μs | 3 | 0 |
| Medium Payload | 🥉 Sonic | Marshal | 15.29μs | 18.8K | 3 |
| Medium Payload | 🥇 BEVE | Marshal | 16.64μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 26.54μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.07μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.16μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.27μs | 20.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.15μs | 61.0K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 71.04μs | 36.1K | 668 |
| Medium Payload | 🥈 CBOR | Unmarshal | 92.25μs | 33.7K | 689 |
| Medium Payload | 🥉 JSON | Unmarshal | 213.15μs | 41.2K | 536 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.57μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 121.90μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 172.60μs | 223.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 226.94μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 280.50μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 507.40μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 280.47μs | 271.3K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 417.21μs | 523.9K | 578 |
| Large Payload | 🥈 MessagePack | Unmarshal | 657.42μs | 336.8K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 807.58μs | 290.0K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.55ms | 508.9K | 6.7K |

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

