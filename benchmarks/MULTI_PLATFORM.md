# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-07-06 07:06:59 UTC

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
| Apple M1 (Virtual) | 692ns | 540ns | 1.16μs | 415ns | 2.74μs |
| AMD EPYC 7763 64-Core Processor | 1.95μs | 242ns | 2.18μs | 949ns | 2.14μs |
| Neoverse-N2 | 235ns | 491ns | 2.12μs | 2.27μs | 2.37μs |
| Unknown CPU | 607ns | 401ns | 1.88μs | 780ns | 3.24μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 643ns | 14.23μs | 5.32μs | 1.43μs |
| AMD EPYC 7763 64-Core Processor | 1.53μs | 19.93μs | 9.18μs | 2.94μs |
| Neoverse-N2 | 669ns | 23.45μs | 5.57μs | 1.64μs |
| Unknown CPU | 1.11μs | 24.50μs | 10.30μs | 4.18μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥈 CBOR (415ns) | 🥇 BEVE (643ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (242ns) | 🥇 BEVE (1.53μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE (235ns) | 🥇 BEVE (669ns) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (401ns) | 🥇 BEVE (1.11μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 51.9% faster

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
| Small Struct | 🥈 CBOR | Marshal | 415ns | 416 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 540ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 692ns | 1.4K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.16μs | 704 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.74μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Marshal | 3.23μs | 1.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 643ns | 952 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.43μs | 1.2K | 27 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.14μs | 3.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.32μs | 4.7K | 100 |
| Small Struct | 🥉 JSON | Unmarshal | 14.23μs | 4.5K | 76 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 4.58μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 8.01μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 13.69μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 24.25μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 34.67μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 36.08μs | 18.6K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 14.52μs | 20.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 27.34μs | 34.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 40.62μs | 38.3K | 710 |
| Medium Payload | 🥈 CBOR | Unmarshal | 47.15μs | 33.2K | 685 |
| Medium Payload | 🥉 JSON | Unmarshal | 197.21μs | 64.4K | 858 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 56.84μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 81.00μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 130.50μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 185.68μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 329.15μs | 221.5K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 382.68μs | 205.3K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 158.56μs | 275.0K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 282.61μs | 346.0K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 369.96μs | 350.7K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 470.84μs | 314.5K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.65ms | 526.1K | 6.8K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 242ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 949ns | 256 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.95μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.14μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 2.18μs | 640 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.94μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.53μs | 1.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.06μs | 2.0K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.94μs | 1.5K | 34 |
| Small Struct | 🥈 CBOR | Unmarshal | 9.18μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 19.93μs | 4.2K | 68 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.43μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.21μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.37μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.34μs | 20.5K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 37.63μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 43.19μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.21μs | 25.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 43.55μs | 64.7K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 59.89μs | 38.9K | 728 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.67μs | 31.9K | 653 |
| Medium Payload | 🥉 JSON | Unmarshal | 260.89μs | 65.1K | 864 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 76.97μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.61μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 165.24μs | 215.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 223.53μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 338.45μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 451.69μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 244.61μs | 267.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 374.28μs | 549.6K | 592 |
| Large Payload | 🥈 MessagePack | Unmarshal | 570.42μs | 338.2K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 758.64μs | 336.4K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.28ms | 524.7K | 6.8K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 235ns | 160 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 491ns | 0 | 0 |
| Small Struct | 🥉 JSON | Marshal | 2.12μs | 1.2K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.27μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.37μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.73μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 669ns | 440 | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.64μs | 736 | 18 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.31μs | 5.7K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.57μs | 3.2K | 70 |
| Small Struct | 🥉 JSON | Unmarshal | 23.45μs | 7.7K | 106 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.54μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 10.45μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.58μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.28μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 30.35μs | 22.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 39.08μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.12μs | 28.1K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.34μs | 48.8K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 54.80μs | 39.4K | 730 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.95μs | 31.3K | 646 |
| Medium Payload | 🥉 JSON | Unmarshal | 193.44μs | 55.5K | 708 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.51μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 114.60μs | 196.8K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 184.03μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 281.56μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 293.38μs | 207.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 410.50μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 216.84μs | 252.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 282.23μs | 384.5K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 502.38μs | 335.7K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 687.47μs | 335.3K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.97ms | 532.1K | 7.0K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 401ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 607ns | 256 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 780ns | 576 | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.88μs | 704 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.52μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.24μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.11μs | 1.2K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.18μs | 2.5K | 53 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.87μs | 7.8K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 10.30μs | 5.1K | 105 |
| Small Struct | 🥉 JSON | Unmarshal | 24.50μs | 4.8K | 86 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.20μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.02μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.94μs | 20.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.44μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.81μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 52.15μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 30.83μs | 33.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 54.69μs | 67.8K | 80 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 77.50μs | 43.1K | 804 |
| Medium Payload | 🥈 CBOR | Unmarshal | 89.36μs | 33.2K | 682 |
| Medium Payload | 🥉 JSON | Unmarshal | 212.73μs | 42.4K | 585 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 78.84μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.32μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 161.34μs | 207.1K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 239.26μs | 204.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 287.44μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 499.08μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.12μs | 267.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 434.12μs | 538.7K | 576 |
| Large Payload | 🥈 MessagePack | Unmarshal | 655.12μs | 337.4K | 6.1K |
| Large Payload | 🥈 CBOR | Unmarshal | 843.67μs | 309.2K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.59ms | 535.8K | 7.0K |

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

