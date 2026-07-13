# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-07-13 06:03:16 UTC

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
| benchmark-linux-amd-epyc-9v74-80-core-processor | AMD EPYC 9V74 80-Core Processor | Linux | [📄 Report](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.md) · [📊 JSON](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.json) · [📈 Chart](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 342ns | 348ns | 2.83μs | 1.04μs | 2.91μs |
| AMD EPYC 9V74 80-Core Processor | 1.65μs | 506ns | 4.53μs | 2.11μs | 4.07μs |
| Neoverse-N2 | 455ns | 505ns | 4.27μs | 2.30μs | 3.72μs |
| Unknown CPU | 1.53μs | 794ns | 4.76μs | 1.26μs | 1.21μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 913ns | 4.59μs | 1.19μs | 3.19μs |
| AMD EPYC 9V74 80-Core Processor | 688ns | 11.63μs | 9.12μs | 5.87μs |
| Neoverse-N2 | 786ns | 23.51μs | 4.17μs | 2.07μs |
| Unknown CPU | 1.93μs | 29.77μs | 3.16μs | 4.62μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (342ns) | 🥇 BEVE (913ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 9V74 80-Core Processor | 🥇 BEVE ZeroCopy (506ns) | 🥇 BEVE (688ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (455ns) | 🥇 BEVE (786ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (794ns) | 🥇 BEVE (1.93μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 77.2% faster

### Platform Details

- **Apple M1 (Virtual)** (Darwin)
  - Architecture: arm64
  - Test Scenarios: 3

- **AMD EPYC 9V74 80-Core Processor** (Linux)
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
| Small Struct | 🥇 BEVE | Marshal | 342ns | 416 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 348ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.04μs | 1.2K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.83μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.91μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 5.45μs | 2.7K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 913ns | 2.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.19μs | 656 | 17 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.19μs | 3.5K | 73 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.28μs | 5.2K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 4.59μs | 1.3K | 28 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.76μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.10μs | 27.3K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 23.69μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.93μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 32.19μs | 24.8K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 39.88μs | 22.0K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.39μs | 26.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 34.61μs | 35.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 65.37μs | 35.3K | 657 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.01μs | 34.2K | 699 |
| Medium Payload | 🥉 JSON | Unmarshal | 245.33μs | 66.6K | 879 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 62.54μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.54μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 154.06μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 222.31μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 384.90μs | 205.2K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 395.09μs | 205.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 182.47μs | 262.4K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 304.64μs | 331.8K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 384.07μs | 348.9K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 541.92μs | 324.8K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.00ms | 562.1K | 7.3K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 9V74 80-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 506ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.65μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.01μs | 2.7K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.11μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.07μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.53μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 688ns | 504 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.32μs | 1.3K | 7 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.87μs | 4.7K | 99 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.12μs | 4.2K | 89 |
| Small Struct | 🥉 JSON | Unmarshal | 11.63μs | 3.7K | 51 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.03μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.18μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.75μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.50μs | 18.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.30μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 50.28μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.49μs | 31.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.70μs | 60.5K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.01μs | 37.0K | 690 |
| Medium Payload | 🥈 CBOR | Unmarshal | 87.37μs | 35.9K | 740 |
| Medium Payload | 🥉 JSON | Unmarshal | 201.07μs | 47.9K | 660 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 71.94μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 125.31μs | 204.9K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 154.93μs | 208.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 209.79μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 329.25μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 412.19μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 239.09μs | 257.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 422.28μs | 550.4K | 588 |
| Large Payload | 🥈 MessagePack | Unmarshal | 597.25μs | 379.7K | 7.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 850.95μs | 327.8K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 1.92ms | 476.2K | 6.3K |

[📄 View full report](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 455ns | 512 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 505ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 768ns | 421 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.30μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.72μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.27μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 786ns | 568 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.03μs | 2.9K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.07μs | 1.2K | 28 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.17μs | 2.3K | 50 |
| Small Struct | 🥉 JSON | Unmarshal | 23.51μs | 7.7K | 107 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.02μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.34μs | 21.8K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 21.81μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.99μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 33.99μs | 25.0K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 38.77μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.62μs | 35.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.34μs | 49.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 51.89μs | 36.0K | 665 |
| Medium Payload | 🥈 CBOR | Unmarshal | 68.09μs | 33.7K | 694 |
| Medium Payload | 🥉 JSON | Unmarshal | 205.75μs | 57.4K | 753 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.80μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 106.64μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 192.16μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 296.62μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 305.22μs | 214.8K | 3 |
| Large Payload | 🥉 JSON | Marshal | 377.91μs | 205.2K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 238.07μs | 284.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 290.50μs | 388.1K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 520.21μs | 339.2K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 698.04μs | 338.2K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 561.1K | 7.4K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 794ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.21μs | 1.0K | 6 |
| Small Struct | 🥈 CBOR | Marshal | 1.26μs | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.53μs | 1.5K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.09μs | 2.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.76μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.93μs | 1.8K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.16μs | 1.2K | 27 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.28μs | 3.5K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.62μs | 2.4K | 52 |
| Small Struct | 🥉 JSON | Unmarshal | 29.77μs | 7.4K | 98 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.50μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.50μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.10μs | 18.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 26.38μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 38.49μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 52.34μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.06μs | 25.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 62.90μs | 69.7K | 80 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.18μs | 23.5K | 486 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 76.04μs | 36.8K | 680 |
| Medium Payload | 🥉 JSON | Unmarshal | 289.92μs | 55.3K | 735 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 80.80μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.35μs | 188.4K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 172.82μs | 222.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 212.79μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 310.40μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 477.07μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 300.74μs | 260.5K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 530.28μs | 591.3K | 601 |
| Large Payload | 🥈 MessagePack | Unmarshal | 717.43μs | 319.7K | 5.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 925.04μs | 295.7K | 6.0K |
| Large Payload | 🥉 JSON | Unmarshal | 2.54ms | 518.8K | 6.8K |

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

