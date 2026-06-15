# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-06-15 08:53:12 UTC

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
| Apple M1 (Virtual) | 1.75μs | 186ns | 1.61μs | 2.72μs | 3.87μs |
| AMD EPYC 7763 64-Core Processor | 2.06μs | 865ns | 3.89μs | 3.54μs | 6.55μs |
| Neoverse-N2 | 613ns | 536ns | 2.12μs | 2.53μs | 2.38μs |
| Unknown CPU | 1.10μs | 565ns | 5.35μs | 1.37μs | 4.72μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.04μs | 28.57μs | 9.17μs | 2.33μs |
| AMD EPYC 7763 64-Core Processor | 1.54μs | 30.60μs | 2.76μs | 1.87μs |
| Neoverse-N2 | 1.30μs | 15.18μs | 1.44μs | 5.39μs |
| Unknown CPU | 2.25μs | 6.16μs | 10.19μs | 3.68μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (186ns) | 🥇 BEVE (1.04μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (865ns) | 🥇 BEVE (1.54μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (536ns) | 🥇 BEVE (1.30μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (565ns) | 🥇 BEVE (2.25μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 47.2% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 186ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 952ns | 215 | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.61μs | 1.0K | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.75μs | 2.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.72μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.87μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.04μs | 1.2K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.28μs | 2.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.33μs | 1.2K | 28 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.17μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 28.57μs | 7.3K | 94 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.71μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 16.70μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 29.49μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.52μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.24μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 58.03μs | 25.0K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.65μs | 25.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 47.03μs | 43.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.97μs | 29.8K | 542 |
| Medium Payload | 🥈 CBOR | Unmarshal | 92.43μs | 35.3K | 726 |
| Medium Payload | 🥉 JSON | Unmarshal | 291.37μs | 73.0K | 933 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.59μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 144.31μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 262.54μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 431.27μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 523.18μs | 205.2K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 634.28μs | 214.3K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 304.87μs | 278.3K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 405.68μs | 327.6K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 660.02μs | 354.4K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 752.46μs | 312.9K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.54ms | 554.2K | 7.3K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 865ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.64μs | 1.6K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.06μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 3.54μs | 3.1K | 1 |
| Small Struct | 🥉 JSON | Marshal | 3.89μs | 1.5K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 6.55μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.54μs | 1.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.78μs | 2.3K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.87μs | 880 | 21 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.76μs | 1.2K | 28 |
| Small Struct | 🥉 JSON | Unmarshal | 30.60μs | 7.9K | 113 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.88μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.80μs | 14.3K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 20.55μs | 18.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.77μs | 28.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 40.04μs | 19.3K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 41.11μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.62μs | 26.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 40.36μs | 58.1K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 45.68μs | 26.9K | 482 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.44μs | 27.6K | 570 |
| Medium Payload | 🥉 JSON | Unmarshal | 252.48μs | 59.9K | 763 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.60μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 123.52μs | 204.9K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 163.44μs | 223.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 219.35μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 343.52μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 421.07μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 248.10μs | 280.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 346.85μs | 500.1K | 564 |
| Large Payload | 🥈 MessagePack | Unmarshal | 601.88μs | 372.4K | 6.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 742.72μs | 333.8K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.52ms | 545.2K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 536ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 613ns | 896 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 925ns | 517 | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.12μs | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.38μs | 4.1K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.53μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.30μs | 2.1K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.36μs | 1.2K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.44μs | 424 | 12 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.39μs | 4.6K | 96 |
| Small Struct | 🥉 JSON | Unmarshal | 15.18μs | 4.3K | 71 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.74μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.79μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.46μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 30.57μs | 22.1K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.54μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 34.37μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Unmarshal | 24.77μs | 30.6K | 33 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.43μs | 35.6K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.16μs | 36.5K | 676 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.82μs | 33.3K | 681 |
| Medium Payload | 🥉 JSON | Unmarshal | 217.07μs | 59.3K | 811 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.48μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.64μs | 188.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 189.59μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 289.62μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 325.95μs | 218.2K | 3 |
| Large Payload | 🥉 JSON | Marshal | 417.13μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 233.22μs | 278.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 302.56μs | 408.0K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 510.79μs | 337.5K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 652.88μs | 311.6K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.89ms | 501.9K | 6.6K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 565ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.07μs | 1.3K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.10μs | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.37μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.72μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 5.35μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.25μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.68μs | 1.6K | 36 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.74μs | 4.4K | 9 |
| Small Struct | 🥉 JSON | Unmarshal | 6.16μs | 1.2K | 25 |
| Small Struct | 🥈 CBOR | Unmarshal | 10.19μs | 4.7K | 100 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.96μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.68μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.86μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 28.63μs | 27.3K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 33.38μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.37μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.49μs | 29.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.57μs | 57.4K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 73.22μs | 39.2K | 728 |
| Medium Payload | 🥈 CBOR | Unmarshal | 88.39μs | 33.9K | 699 |
| Medium Payload | 🥉 JSON | Unmarshal | 228.07μs | 46.0K | 596 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.93μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.53μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 154.04μs | 198.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 216.93μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 279.25μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 475.74μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.33μs | 267.2K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 453.84μs | 580.7K | 595 |
| Large Payload | 🥈 MessagePack | Unmarshal | 686.26μs | 361.5K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 852.60μs | 310.9K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.77ms | 564.4K | 7.3K |

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

