# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-11-24 03:58:48 UTC

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
| Apple M1 (Virtual) | 225ns | 401ns | 3.39μs | 748ns | 1.72μs |
| AMD EPYC 7763 64-Core Processor | 363ns | 618ns | 3.17μs | 1.02μs | 1.55μs |
| Neoverse-N2 | 1.07μs | 666ns | 3.69μs | 983ns | 1.40μs |
| Unknown CPU | 727ns | 366ns | 4.13μs | 2.19μs | 3.41μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 932ns | 21.69μs | 4.08μs | 3.69μs |
| AMD EPYC 7763 64-Core Processor | 1.04μs | 17.98μs | 8.14μs | 1.31μs |
| Neoverse-N2 | 1.82μs | 7.10μs | 4.37μs | 2.43μs |
| Unknown CPU | 1.44μs | 12.84μs | 3.02μs | 6.23μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (225ns) | 🥇 BEVE (932ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE (363ns) | 🥇 BEVE (1.04μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥉 Sonic (665ns) | 🥇 BEVE (1.82μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (366ns) | 🥇 BEVE (1.44μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 83.8% faster

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
| Small Struct | 🥇 BEVE | Marshal | 225ns | 320 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 401ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 748ns | 1.2K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.72μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 3.39μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 5.28μs | 2.7K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 932ns | 2.1K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.69μs | 3.9K | 83 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.08μs | 2.8K | 60 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.46μs | 5.9K | 6 |
| Small Struct | 🥉 JSON | Unmarshal | 21.69μs | 7.8K | 108 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.64μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.35μs | 19.1K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 18.99μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 28.48μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 29.82μs | 19.3K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 43.67μs | 24.9K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 16.29μs | 26.3K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.72μs | 31.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 45.92μs | 37.5K | 697 |
| Medium Payload | 🥈 CBOR | Unmarshal | 59.42μs | 36.0K | 744 |
| Medium Payload | 🥉 JSON | Unmarshal | 149.93μs | 41.1K | 530 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 61.17μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 113.29μs | 204.9K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 201.50μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 272.51μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 410.19μs | 221.6K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 534.18μs | 222.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 292.83μs | 258.7K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 342.73μs | 340.8K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 585.96μs | 357.2K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 718.25μs | 312.9K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.22ms | 525.9K | 6.9K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE | Marshal | 363ns | 384 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 618ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 1.02μs | 896 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.08μs | 1.4K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.55μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 3.17μs | 1.8K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.04μs | 1.2K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.31μs | 448 | 12 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.29μs | 7.8K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 8.14μs | 4.7K | 100 |
| Small Struct | 🥉 JSON | Unmarshal | 17.98μs | 4.4K | 74 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.33μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.36μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.32μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 24.99μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.43μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 45.91μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.02μs | 29.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 38.38μs | 54.3K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.59μs | 29.9K | 540 |
| Medium Payload | 🥈 CBOR | Unmarshal | 68.06μs | 29.8K | 615 |
| Medium Payload | 🥉 JSON | Unmarshal | 206.63μs | 48.1K | 642 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.74μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 116.68μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 156.80μs | 215.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 208.70μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 316.49μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 452.57μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 241.10μs | 279.3K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 337.53μs | 506.6K | 561 |
| Large Payload | 🥈 MessagePack | Unmarshal | 556.92μs | 344.8K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 718.23μs | 322.5K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 2.32ms | 544.8K | 7.2K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥉 Sonic | Marshal | 665ns | 311 | 2 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 666ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 983ns | 896 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.07μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.40μs | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 3.69μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.82μs | 3.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.33μs | 4.1K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.43μs | 1.6K | 36 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.37μs | 2.4K | 53 |
| Small Struct | 🥉 JSON | Unmarshal | 7.10μs | 2.0K | 35 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.77μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.54μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 19.15μs | 20.5K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 29.46μs | 20.8K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.56μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.60μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.41μs | 30.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 25.58μs | 34.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 47.91μs | 32.0K | 586 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.60μs | 31.2K | 641 |
| Medium Payload | 🥉 JSON | Unmarshal | 218.06μs | 63.1K | 814 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 67.47μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 105.05μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 181.11μs | 188.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 275.57μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 299.38μs | 214.5K | 3 |
| Large Payload | 🥉 JSON | Marshal | 389.95μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 223.62μs | 269.6K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 286.67μs | 389.7K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 536.74μs | 368.1K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 635.50μs | 302.5K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 1.98ms | 533.6K | 7.0K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 366ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 484ns | 421 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 727ns | 768 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.19μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.41μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.13μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.44μs | 1.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.02μs | 1.2K | 27 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.54μs | 7.1K | 10 |
| Small Struct | 🥈 MessagePack | Unmarshal | 6.23μs | 3.9K | 83 |
| Small Struct | 🥉 JSON | Unmarshal | 12.84μs | 3.8K | 54 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.14μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.66μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 22.95μs | 25.1K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 28.26μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 43.11μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 49.94μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 28.72μs | 26.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.26μs | 55.8K | 76 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 62.61μs | 33.1K | 609 |
| Medium Payload | 🥈 CBOR | Unmarshal | 80.75μs | 32.8K | 678 |
| Medium Payload | 🥉 JSON | Unmarshal | 270.08μs | 70.3K | 885 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 72.78μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 128.92μs | 180.3K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 206.42μs | 225.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 266.69μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 337.61μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 497.37μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 298.78μs | 264.5K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 465.83μs | 550.8K | 576 |
| Large Payload | 🥈 MessagePack | Unmarshal | 669.34μs | 344.3K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 765.40μs | 290.1K | 5.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.23ms | 497.2K | 6.6K |

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

