# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-04-27 05:50:13 UTC

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
| Apple M1 (Virtual) | 606ns | 586ns | 1.01μs | 685ns | 1000ns |
| AMD EPYC 9V74 80-Core Processor | 744ns | 450ns | 5.25μs | 781ns | 2.57μs |
| Neoverse-N2 | 979ns | 342ns | 1.20μs | 1.98μs | 3.74μs |
| Unknown CPU | 1.93μs | 758ns | 5.24μs | 2.82μs | 2.65μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.01μs | 18.49μs | 1.00μs | 1.81μs |
| AMD EPYC 9V74 80-Core Processor | 1.06μs | 24.27μs | 5.24μs | 1.28μs |
| Neoverse-N2 | 1.50μs | 24.18μs | 7.31μs | 5.10μs |
| Unknown CPU | 757ns | 13.23μs | 5.30μs | 5.27μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (586ns) | 🥈 CBOR (1.00μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 9V74 80-Core Processor | 🥇 BEVE ZeroCopy (450ns) | 🥇 BEVE (1.06μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (342ns) | 🥇 BEVE (1.50μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (758ns) | 🥇 BEVE (757ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 51.7% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 586ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 606ns | 1.5K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 685ns | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1000ns | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 1.01μs | 768 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.69μs | 2.3K | 2 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.00μs | 376 | 11 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.01μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.63μs | 1.7K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.81μs | 1.8K | 39 |
| Small Struct | 🥉 JSON | Unmarshal | 18.49μs | 7.5K | 101 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.32μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 8.22μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.85μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.26μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 27.67μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 43.61μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.69μs | 33.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.01μs | 31.9K | 31 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 43.87μs | 38.2K | 706 |
| Medium Payload | 🥈 CBOR | Unmarshal | 51.53μs | 33.8K | 693 |
| Medium Payload | 🥉 JSON | Unmarshal | 167.64μs | 52.1K | 684 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 54.15μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 84.80μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 154.09μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 256.38μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 367.82μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 496.03μs | 213.9K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 173.75μs | 268.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 296.92μs | 348.4K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 418.25μs | 365.2K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 596.35μs | 340.5K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.65ms | 488.6K | 6.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 9V74 80-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 450ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 744ns | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 781ns | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.09μs | 1.5K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.57μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.25μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.06μs | 1.3K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.28μs | 496 | 13 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.46μs | 3.7K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.24μs | 2.5K | 54 |
| Small Struct | 🥉 JSON | Unmarshal | 24.27μs | 7.7K | 105 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.56μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.17μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.08μs | 22.3K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.64μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.08μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.11μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.04μs | 29.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 46.20μs | 59.5K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 54.26μs | 34.5K | 639 |
| Medium Payload | 🥈 CBOR | Unmarshal | 76.52μs | 31.5K | 649 |
| Medium Payload | 🥉 JSON | Unmarshal | 208.97μs | 53.3K | 724 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 73.80μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.14μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 166.13μs | 216.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 200.30μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 327.48μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 436.33μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 228.66μs | 268.4K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 395.82μs | 550.1K | 575 |
| Large Payload | 🥈 MessagePack | Unmarshal | 529.18μs | 337.4K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 789.38μs | 328.4K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 1.99ms | 509.2K | 6.7K |

[📄 View full report](benchmark-linux-amd-epyc-9v74-80-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 342ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 979ns | 1.8K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.20μs | 576 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.31μs | 812 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.98μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.74μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.50μs | 2.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.84μs | 4.5K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.10μs | 4.3K | 90 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.31μs | 4.4K | 94 |
| Small Struct | 🥉 JSON | Unmarshal | 24.18μs | 7.8K | 110 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.89μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.41μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 21.79μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.24μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 38.15μs | 27.8K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 44.83μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.45μs | 26.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.02μs | 45.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.97μs | 33.6K | 620 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.94μs | 38.1K | 781 |
| Medium Payload | 🥉 JSON | Unmarshal | 219.76μs | 62.0K | 817 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 71.89μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.45μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 186.65μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 281.57μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 305.53μs | 214.7K | 3 |
| Large Payload | 🥉 JSON | Marshal | 383.76μs | 205.2K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.21μs | 280.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 289.17μs | 392.5K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 492.15μs | 323.0K | 5.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 645.61μs | 310.1K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.92ms | 513.5K | 6.8K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 758ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 1.70μs | 2.1K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.93μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.65μs | 4.1K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.82μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 5.24μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 757ns | 504 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.15μs | 4.4K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.27μs | 3.5K | 72 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.30μs | 2.4K | 52 |
| Small Struct | 🥉 JSON | Unmarshal | 13.23μs | 3.7K | 51 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.59μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.76μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.10μs | 27.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.69μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.06μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 57.55μs | 27.5K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.34μs | 27.8K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 57.95μs | 59.3K | 75 |
| Medium Payload | 🥈 CBOR | Unmarshal | 67.86μs | 24.9K | 515 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 81.38μs | 43.2K | 810 |
| Medium Payload | 🥉 JSON | Unmarshal | 296.12μs | 69.0K | 881 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 73.41μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 107.72μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 162.57μs | 216.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 226.58μs | 213.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 294.29μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 518.16μs | 229.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 312.99μs | 273.1K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 548.61μs | 535.9K | 576 |
| Large Payload | 🥈 MessagePack | Unmarshal | 706.58μs | 359.2K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 865.41μs | 325.1K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.55ms | 527.5K | 6.9K |

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

