# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-05-04 06:00:54 UTC

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
| benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz | Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | Linux | [📄 Report](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.md) · [📊 JSON](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.json) · [📈 Chart](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 1.19μs | 434ns | 1.33μs | 1.99μs | 1.08μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 1.05μs | 495ns | 616ns | 1.74μs | 1.67μs |
| Neoverse-N2 | 925ns | 464ns | 3.84μs | 1.55μs | 3.55μs |
| Unknown CPU | 384ns | 829ns | 5.69μs | 2.15μs | 3.18μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 2.61μs | 16.39μs | 4.74μs | 5.76μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 1.57μs | 8.74μs | 8.71μs | 5.45μs |
| Neoverse-N2 | 1.66μs | 10.12μs | 4.83μs | 5.63μs |
| Unknown CPU | 2.35μs | 30.28μs | 8.88μs | 6.34μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (434ns) | 🥇 BEVE (2.61μs) | 💾 BEVE (1 allocs) |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 🥇 BEVE ZeroCopy (495ns) | 🥇 BEVE (1.57μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (464ns) | 🥇 BEVE (1.66μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE (384ns) | 🥉 Sonic (1.70μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 27.4% faster

### Platform Details

- **Apple M1 (Virtual)** (Darwin)
  - Architecture: arm64
  - Test Scenarios: 3

- **Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz** (Linux)
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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 434ns | 0 | 0 |
| Small Struct | 🥈 MessagePack | Marshal | 1.08μs | 1.0K | 6 |
| Small Struct | 🥉 Sonic | Marshal | 1.15μs | 464 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.19μs | 2.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.33μs | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.99μs | 1.5K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.61μs | 1.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.74μs | 2.3K | 51 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.76μs | 3.9K | 82 |
| Small Struct | 🥉 Sonic | Unmarshal | 6.38μs | 5.4K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 16.39μs | 2.3K | 43 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.23μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 15.62μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.65μs | 18.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 27.99μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 34.28μs | 18.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 50.01μs | 19.3K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.10μs | 27.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.45μs | 32.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 48.84μs | 33.0K | 606 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.43μs | 34.0K | 695 |
| Medium Payload | 🥉 JSON | Unmarshal | 209.36μs | 49.3K | 645 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 61.91μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 103.16μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 193.48μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 254.14μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 384.95μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 505.11μs | 222.3K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 228.96μs | 259.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 258.87μs | 329.8K | 211 |
| Large Payload | 🥈 CBOR | Unmarshal | 512.70μs | 322.0K | 6.6K |
| Large Payload | 🥈 MessagePack | Unmarshal | 539.23μs | 347.0K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.90ms | 526.0K | 6.9K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

![Benchmark Chart](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 495ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 616ns | 192 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.05μs | 1.5K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.16μs | 1.3K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.67μs | 2.1K | 7 |
| Small Struct | 🥈 CBOR | Marshal | 1.74μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.57μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.48μs | 3.7K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.45μs | 4.3K | 90 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.71μs | 5.1K | 105 |
| Small Struct | 🥉 JSON | Unmarshal | 8.74μs | 2.2K | 41 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.03μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.19μs | 19.1K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 14.52μs | 19.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.28μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.26μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 43.56μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 27.11μs | 34.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 48.07μs | 70.5K | 81 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.91μs | 41.4K | 773 |
| Medium Payload | 🥈 CBOR | Unmarshal | 70.10μs | 29.9K | 619 |
| Medium Payload | 🥉 JSON | Unmarshal | 194.09μs | 51.8K | 650 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.50μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.78μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 164.08μs | 216.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 204.82μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 310.10μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 431.75μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 233.78μs | 259.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 378.95μs | 536.6K | 590 |
| Large Payload | 🥈 MessagePack | Unmarshal | 563.11μs | 354.0K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 737.23μs | 330.9K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.02ms | 510.4K | 6.7K |

[📄 View full report](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 464ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 925ns | 1.8K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.55μs | 1.8K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.55μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 3.68μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.84μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.66μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.56μs | 6.4K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.83μs | 2.8K | 59 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.63μs | 5.1K | 104 |
| Small Struct | 🥉 JSON | Unmarshal | 10.12μs | 2.5K | 49 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.78μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.54μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.87μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.80μs | 16.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.24μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 41.77μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.03μs | 27.0K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.24μs | 32.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.61μs | 40.9K | 766 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.86μs | 35.4K | 726 |
| Medium Payload | 🥉 JSON | Unmarshal | 229.17μs | 66.6K | 863 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.90μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 107.50μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 194.06μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 273.30μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 315.55μs | 222.6K | 3 |
| Large Payload | 🥉 JSON | Marshal | 397.65μs | 221.7K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 222.84μs | 268.6K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 277.03μs | 359.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 510.10μs | 335.1K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 688.48μs | 334.7K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 2.06ms | 563.6K | 7.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 384ns | 208 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 829ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 2.15μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.26μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.18μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.69μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.70μs | 1.2K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.35μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.34μs | 4.0K | 86 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.88μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 30.28μs | 7.8K | 108 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.75μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.34μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.75μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 25.63μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.54μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 51.18μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.63μs | 27.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 42.42μs | 46.5K | 68 |
| Medium Payload | 🥈 CBOR | Unmarshal | 77.42μs | 27.1K | 560 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 85.46μs | 41.2K | 775 |
| Medium Payload | 🥉 JSON | Unmarshal | 295.30μs | 63.8K | 845 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 75.91μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 122.06μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 179.73μs | 215.9K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 250.04μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 329.64μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 504.48μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 290.22μs | 270.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 442.86μs | 538.9K | 579 |
| Large Payload | 🥈 MessagePack | Unmarshal | 685.49μs | 340.0K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 945.67μs | 328.8K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.57ms | 511.2K | 6.7K |

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

