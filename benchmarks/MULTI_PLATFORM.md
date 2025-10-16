# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 13:41:30 UTC

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
| Apple M1 (Virtual) | 812ns | 649ns | 2.21μs | 1.74μs | 1.61μs |
| AMD EPYC 7763 64-Core Processor | 1.77μs | 796ns | 1.70μs | 1.58μs | 2.89μs |
| Neoverse-N2 | 1.60μs | 438ns | 1.47μs | 1.18μs | 1.62μs |
| Unknown CPU | 2.17μs | 992ns | 1.46μs | 3.24μs | 2.12μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.19μs | 14.71μs | 4.09μs | 3.47μs |
| AMD EPYC 7763 64-Core Processor | 929ns | 14.70μs | 8.18μs | 3.94μs |
| Neoverse-N2 | 669ns | 20.59μs | 4.26μs | 4.44μs |
| Unknown CPU | 1.92μs | 5.06μs | 6.64μs | 3.25μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (649ns) | 🥇 BEVE (1.19μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (796ns) | 🥇 BEVE (929ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (438ns) | 🥇 BEVE (669ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (992ns) | 🥇 BEVE (1.92μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 0.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 649ns | 290 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 812ns | 2.1K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.61μs | 4.2K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 1.74μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.21μs | 1.7K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 2.46μs | 1.5K | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.19μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.26μs | 3.2K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.47μs | 4.0K | 86 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.09μs | 3.5K | 74 |
| Small Struct | 🥉 JSON | Unmarshal | 14.71μs | 4.6K | 80 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.68μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 8.70μs | 21.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 15.93μs | 21.9K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.59μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 34.13μs | 24.9K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 39.25μs | 22.0K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 18.01μs | 25.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.54μs | 45.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 33.88μs | 31.5K | 577 |
| Medium Payload | 🥈 CBOR | Unmarshal | 49.24μs | 32.3K | 669 |
| Medium Payload | 🥉 JSON | Unmarshal | 205.39μs | 70.1K | 895 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 55.56μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 89.63μs | 197.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 138.52μs | 205.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 183.67μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 306.16μs | 197.4K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 399.22μs | 223.0K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 164.49μs | 271.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 233.55μs | 319.5K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 380.78μs | 360.7K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 510.25μs | 322.2K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 1.62ms | 526.2K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 796ns | 290 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.42μs | 2.0K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 1.58μs | 1.6K | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.70μs | 848 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.77μs | 3.4K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.89μs | 4.2K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 929ns | 1.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.94μs | 2.5K | 55 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.28μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.18μs | 4.4K | 94 |
| Small Struct | 🥉 JSON | Unmarshal | 14.70μs | 4.0K | 62 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.54μs | 134 | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 13.57μs | 17.1K | 4 |
| Medium Payload | 🥇 BEVE | Marshal | 14.11μs | 21.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 27.81μs | 24.7K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 41.10μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 41.66μs | 19.5K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.03μs | 25.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.70μs | 73.4K | 81 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.31μs | 30.3K | 624 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 70.47μs | 47.5K | 900 |
| Medium Payload | 🥉 JSON | Unmarshal | 226.49μs | 52.0K | 716 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.61μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 118.40μs | 180.6K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 158.17μs | 200.7K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 211.63μs | 197.4K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 310.28μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 430.38μs | 205.3K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 232.72μs | 274.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 359.30μs | 540.3K | 570 |
| Large Payload | 🥈 MessagePack | Unmarshal | 581.57μs | 357.4K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 739.23μs | 319.9K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.10ms | 472.5K | 6.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 438ns | 289 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 824ns | 572 | 3 |
| Small Struct | 🥈 CBOR | Marshal | 1.18μs | 1.2K | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.47μs | 848 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.60μs | 3.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.62μs | 2.2K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 669ns | 376 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.95μs | 4.8K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.26μs | 2.3K | 51 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.44μs | 3.5K | 74 |
| Small Struct | 🥉 JSON | Unmarshal | 20.59μs | 7.3K | 93 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.92μs | 128 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 10.53μs | 18.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.93μs | 24.7K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.63μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 31.40μs | 25.1K | 4 |
| Medium Payload | 🥉 JSON | Marshal | 35.17μs | 19.4K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.83μs | 30.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.74μs | 34.4K | 29 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 46.51μs | 30.3K | 551 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.65μs | 30.4K | 623 |
| Medium Payload | 🥉 JSON | Unmarshal | 184.35μs | 51.3K | 660 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 71.63μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 105.85μs | 173.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 197.46μs | 207.3K | 3 |
| Large Payload | 🥈 MessagePack | Marshal | 278.00μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 311.54μs | 227.2K | 4 |
| Large Payload | 🥉 JSON | Marshal | 390.18μs | 225.4K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 232.12μs | 282.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 276.75μs | 365.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 518.27μs | 347.0K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 659.26μs | 320.6K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 2.02ms | 545.5K | 7.1K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 992ns | 289 | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.46μs | 656 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.72μs | 2.1K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.12μs | 2.2K | 7 |
| Small Struct | 🥇 BEVE | Marshal | 2.17μs | 3.0K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 3.24μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.92μs | 3.0K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.25μs | 1.8K | 40 |
| Small Struct | 🥉 JSON | Unmarshal | 5.06μs | 872 | 21 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.61μs | 7.8K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.64μs | 3.2K | 68 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.14μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.30μs | 18.6K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 16.47μs | 20.9K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 21.89μs | 18.5K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.39μs | 33.1K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 46.03μs | 20.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.70μs | 28.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.22μs | 42.3K | 61 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 69.61μs | 36.1K | 665 |
| Medium Payload | 🥈 CBOR | Unmarshal | 98.44μs | 38.5K | 797 |
| Medium Payload | 🥉 JSON | Unmarshal | 261.12μs | 56.4K | 730 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.07μs | 207 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 112.63μs | 189.3K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 160.02μs | 227.7K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 250.73μs | 206.0K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 290.03μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 485.10μs | 215.2K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 277.79μs | 276.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 420.96μs | 525.2K | 570 |
| Large Payload | 🥈 MessagePack | Unmarshal | 707.51μs | 353.5K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 902.03μs | 332.3K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.56ms | 531.2K | 7.0K |

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

