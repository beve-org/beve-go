# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-04-20 05:36:13 UTC

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
| Apple M1 (Virtual) | 355ns | 496ns | 1.46μs | 1.28μs | 917ns |
| AMD EPYC 7763 64-Core Processor | 800ns | 620ns | 4.17μs | 3.18μs | 2.27μs |
| Neoverse-N2 | 921ns | 674ns | 2.69μs | 1.94μs | 2.27μs |
| Unknown CPU | 300ns | 492ns | 1.93μs | 839ns | 2.07μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 999ns | 8.61μs | 1.67μs | 648ns |
| AMD EPYC 7763 64-Core Processor | 2.86μs | 4.01μs | 2.31μs | 3.18μs |
| Neoverse-N2 | 1.47μs | 17.45μs | 3.67μs | 3.32μs |
| Unknown CPU | 1.88μs | 14.01μs | 7.71μs | 1.25μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (355ns) | 🥈 MessagePack (648ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥉 Sonic (615ns) | 🥈 CBOR (2.31μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (674ns) | 🥇 BEVE (1.47μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (300ns) | 🥈 MessagePack (1.25μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 76.7% faster

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
| Small Struct | 🥇 BEVE | Marshal | 355ns | 704 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 496ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 917ns | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 1.28μs | 2.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.46μs | 1.2K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 5.15μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 648ns | 256 | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 999ns | 2.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.67μs | 1.1K | 25 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.87μs | 2.8K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 8.61μs | 3.8K | 54 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.84μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.60μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 14.13μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 20.88μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 28.95μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 41.48μs | 24.9K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 15.77μs | 31.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.55μs | 43.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 33.56μs | 35.0K | 646 |
| Medium Payload | 🥈 CBOR | Unmarshal | 52.26μs | 36.4K | 747 |
| Medium Payload | 🥉 JSON | Unmarshal | 167.89μs | 58.5K | 780 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 50.39μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 76.26μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 132.36μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 187.34μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 287.06μs | 196.9K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 380.25μs | 221.5K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 165.90μs | 254.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 258.66μs | 354.3K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 349.03μs | 329.8K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 430.31μs | 289.4K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.65ms | 564.6K | 7.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 615ns | 414 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 620ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 800ns | 1.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.27μs | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 3.18μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 4.17μs | 1.8K | 1 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.31μs | 480 | 13 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.86μs | 3.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.18μs | 1.1K | 26 |
| Small Struct | 🥉 JSON | Unmarshal | 4.01μs | 544 | 15 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.55μs | 7.0K | 10 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.98μs | 3 | 0 |
| Medium Payload | 🥉 Sonic | Marshal | 15.58μs | 20.9K | 3 |
| Medium Payload | 🥇 BEVE | Marshal | 19.59μs | 24.6K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 20.75μs | 18.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.49μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 52.17μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.33μs | 30.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 45.28μs | 64.6K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.34μs | 37.0K | 686 |
| Medium Payload | 🥈 CBOR | Unmarshal | 77.87μs | 36.2K | 748 |
| Medium Payload | 🥉 JSON | Unmarshal | 192.68μs | 45.3K | 591 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 83.53μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 132.10μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 164.15μs | 208.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 207.46μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 356.30μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 475.45μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 268.37μs | 280.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 401.81μs | 553.3K | 590 |
| Large Payload | 🥈 MessagePack | Unmarshal | 583.14μs | 351.5K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 738.93μs | 319.7K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.32ms | 537.9K | 7.0K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 674ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 921ns | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.91μs | 1.3K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.94μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.27μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.69μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.47μs | 2.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.32μs | 2.5K | 54 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.43μs | 5.4K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.67μs | 1.9K | 43 |
| Small Struct | 🥉 JSON | Unmarshal | 17.45μs | 4.6K | 81 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.93μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.35μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.79μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 21.38μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 32.39μs | 25.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 48.31μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.64μs | 28.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.18μs | 45.4K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 46.77μs | 30.0K | 545 |
| Medium Payload | 🥈 CBOR | Unmarshal | 68.49μs | 33.7K | 696 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.16μs | 53.3K | 724 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.14μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.64μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 184.32μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 285.81μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 306.04μs | 217.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 427.59μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 224.76μs | 268.5K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 295.98μs | 404.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 515.22μs | 352.1K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 647.79μs | 313.9K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.94ms | 516.3K | 6.8K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 300ns | 256 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 492ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 610ns | 933 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 839ns | 1.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.93μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.07μs | 4.1K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.25μs | 544 | 14 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.88μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.32μs | 3.5K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.71μs | 4.6K | 98 |
| Small Struct | 🥉 JSON | Unmarshal | 14.01μs | 4.1K | 64 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.92μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.85μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 12.83μs | 22.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 15.29μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.62μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 29.70μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.34μs | 31.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 39.92μs | 59.3K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.15μs | 32.0K | 586 |
| Medium Payload | 🥈 CBOR | Unmarshal | 58.35μs | 29.9K | 614 |
| Medium Payload | 🥉 JSON | Unmarshal | 236.95μs | 66.8K | 878 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 57.31μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 85.38μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 125.34μs | 223.6K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 163.87μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 211.06μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 350.63μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 213.56μs | 273.6K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 345.71μs | 539.4K | 569 |
| Large Payload | 🥈 MessagePack | Unmarshal | 555.45μs | 361.0K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 643.17μs | 323.1K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 1.96ms | 527.0K | 6.9K |

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

