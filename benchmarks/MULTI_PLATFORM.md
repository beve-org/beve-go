# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-15 17:27:06 UTC

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
| benchmark-darwin-apple-m1-virtual | Apple M1 (Virtual) | Darwin | [📄 Report](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md) · [📊 JSON](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.json) · [📈 Chart](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png) |
| benchmark-linux-amd-epyc-7763-64-core-processor | AMD EPYC 7763 64-Core Processor | Linux | [📄 Report](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [📊 JSON](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [📈 Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| benchmark-linux-neoverse-n2 | Neoverse-N2 | Linux | [📄 Report](benchmarks/benchmark-linux-neoverse-n2/benchmark.md) · [📊 JSON](benchmarks/benchmark-linux-neoverse-n2/benchmark.json) · [📈 Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png) |
| benchmark-windows-unknown-cpu | Unknown CPU | Windows | [📄 Report](benchmarks/benchmark-windows-unknown-cpu/benchmark.md) · [📊 JSON](benchmarks/benchmark-windows-unknown-cpu/benchmark.json) · [📈 Chart](benchmarks/benchmark-windows-unknown-cpu/benchmark.png) |

---

## 📊 Cross-Platform Performance Comparison

### Marshal Performance (Small Struct)

| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |
|----------|------|---------------|------|------|-------------|
| Apple M1 (Virtual) | 791ns | 736ns | 6.35μs | 2.68μs | 3.53μs |
| AMD EPYC 7763 64-Core Processor | 2.31μs | 626ns | 3.59μs | 2.02μs | 3.19μs |
| Neoverse-N2 | 1.40μs | 497ns | 3.81μs | 1.43μs | 2.38μs |
| Unknown CPU | 2.58μs | 658ns | 6.67μs | 2.35μs | 5.47μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.77μs | 27.86μs | 1.74μs | 1.98μs |
| AMD EPYC 7763 64-Core Processor | 2.71μs | 10.47μs | 6.70μs | 3.43μs |
| Neoverse-N2 | 1.74μs | 10.39μs | 2.81μs | 2.06μs |
| Unknown CPU | 1.64μs | 26.36μs | 2.49μs | 6.60μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (736ns) | 🥈 CBOR (1.74μs) | 💾 BEVE ZeroCopy (2 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (626ns) | 🥇 BEVE (2.71μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (497ns) | 🥇 BEVE (1.74μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (658ns) | 🥇 BEVE (1.64μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 61.9% faster

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

![Benchmark Chart](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 736ns | 288 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 791ns | 1.8K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.68μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.53μs | 8.3K | 9 |
| Small Struct | 🥉 JSON | Marshal | 6.35μs | 3.2K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 8.41μs | 3.3K | 3 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.74μs | 896 | 22 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.77μs | 2.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.98μs | 1.5K | 34 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.43μs | 2.0K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 27.86μs | 7.5K | 101 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.38μs | 128 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 9.47μs | 20.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.68μs | 20.6K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.59μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.94μs | 27.6K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 51.44μs | 25.1K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.53μs | 30.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 37.08μs | 36.9K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.40μs | 37.2K | 693 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.31μs | 32.1K | 662 |
| Medium Payload | 🥉 JSON | Unmarshal | 192.71μs | 55.3K | 726 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 66.84μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 108.33μs | 181.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 167.34μs | 189.5K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 278.05μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 336.76μs | 213.6K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 463.85μs | 215.2K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 255.75μs | 279.2K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 382.15μs | 325.4K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 575.87μs | 359.2K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 760.66μs | 334.5K | 6.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.84ms | 497.9K | 6.5K |

[📄 View full report](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 626ns | 290 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.56μs | 2.0K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.02μs | 1.6K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.31μs | 3.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 3.19μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.59μs | 1.4K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.71μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.95μs | 3.7K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.43μs | 1.8K | 40 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.70μs | 2.3K | 51 |
| Small Struct | 🥉 JSON | Unmarshal | 10.47μs | 2.1K | 38 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 11.11μs | 128 | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 18.77μs | 25.8K | 4 |
| Medium Payload | 🥇 BEVE | Marshal | 20.11μs | 21.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.71μs | 18.6K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.97μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 53.02μs | 22.1K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.57μs | 28.9K | 59 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 41.48μs | 22.6K | 395 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.71μs | 64.3K | 79 |
| Medium Payload | 🥈 CBOR | Unmarshal | 87.89μs | 38.2K | 785 |
| Medium Payload | 🥉 JSON | Unmarshal | 214.47μs | 50.6K | 641 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 95.86μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 130.72μs | 188.8K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 162.88μs | 216.6K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 215.58μs | 205.6K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 301.38μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 432.42μs | 213.8K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 235.10μs | 267.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 367.21μs | 555.6K | 595 |
| Large Payload | 🥈 MessagePack | Unmarshal | 587.06μs | 364.5K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 781.21μs | 316.1K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.30ms | 522.4K | 6.9K |

[📄 View full report](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 497ns | 288 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.29μs | 906 | 3 |
| Small Struct | 🥇 BEVE | Marshal | 1.40μs | 2.3K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 1.43μs | 1.6K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.38μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 3.81μs | 2.4K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.74μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.01μs | 2.6K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.06μs | 1.2K | 28 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.81μs | 1.3K | 30 |
| Small Struct | 🥉 JSON | Unmarshal | 10.39μs | 2.5K | 49 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.93μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 10.04μs | 16.5K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 19.33μs | 20.6K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 28.31μs | 21.0K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 28.97μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.09μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.23μs | 28.6K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 28.78μs | 39.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.27μs | 43.1K | 813 |
| Medium Payload | 🥈 CBOR | Unmarshal | 73.13μs | 36.6K | 755 |
| Medium Payload | 🥉 JSON | Unmarshal | 212.82μs | 59.9K | 794 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 75.39μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 123.77μs | 189.6K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 183.11μs | 200.3K | 3 |
| Large Payload | 🥈 MessagePack | Marshal | 288.94μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 316.84μs | 231.4K | 4 |
| Large Payload | 🥉 JSON | Marshal | 368.41μs | 206.2K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 218.29μs | 265.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 278.89μs | 384.5K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 484.51μs | 328.7K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 658.74μs | 321.2K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.04ms | 559.3K | 7.3K |

[📄 View full report](benchmarks/benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmarks/benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 658ns | 289 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.70μs | 2.0K | 3 |
| Small Struct | 🥈 CBOR | Marshal | 2.35μs | 2.4K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.58μs | 1.7K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 5.47μs | 8.3K | 9 |
| Small Struct | 🥉 JSON | Marshal | 6.67μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.64μs | 2.1K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.49μs | 808 | 20 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.90μs | 7.8K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.60μs | 4.3K | 90 |
| Small Struct | 🥉 JSON | Unmarshal | 26.36μs | 7.3K | 93 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.79μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.66μs | 16.5K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 17.14μs | 20.9K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 26.19μs | 21.9K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.11μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.01μs | 22.1K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.07μs | 30.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 53.15μs | 63.8K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.82μs | 27.6K | 497 |
| Medium Payload | 🥈 CBOR | Unmarshal | 83.28μs | 27.8K | 575 |
| Medium Payload | 🥉 JSON | Unmarshal | 285.65μs | 60.1K | 802 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 95.29μs | 312 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 133.36μs | 189.2K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 157.75μs | 211.4K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 225.65μs | 206.1K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 283.85μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 488.26μs | 215.6K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 288.67μs | 274.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 428.12μs | 518.8K | 564 |
| Large Payload | 🥈 MessagePack | Unmarshal | 671.52μs | 341.2K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 854.17μs | 314.0K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.33ms | 482.4K | 6.3K |

[📄 View full report](benchmarks/benchmark-windows-unknown-cpu/benchmark.md)

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

