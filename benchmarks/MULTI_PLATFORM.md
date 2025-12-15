# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-12-15 04:06:01 UTC

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
| Apple M1 (Virtual) | 830ns | 273ns | 1.03μs | 444ns | 968ns |
| AMD EPYC 7763 64-Core Processor | 919ns | 376ns | 1.64μs | 1.89μs | 2.73μs |
| Neoverse-N2 | 1.35μs | 762ns | 4.38μs | 1.88μs | 3.59μs |
| Unknown CPU | 894ns | 851ns | 4.95μs | 2.68μs | 2.90μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 571ns | 17.58μs | 3.32μs | 3.89μs |
| AMD EPYC 7763 64-Core Processor | 952ns | 3.02μs | 10.81μs | 7.72μs |
| Neoverse-N2 | 1.45μs | 11.20μs | 7.90μs | 3.20μs |
| Unknown CPU | 1.88μs | 11.33μs | 2.44μs | 5.25μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (273ns) | 🥇 BEVE (571ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (376ns) | 🥇 BEVE (952ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (762ns) | 🥇 BEVE (1.45μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (851ns) | 🥇 BEVE (1.88μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 53.7% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 273ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 444ns | 448 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 830ns | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 968ns | 2.1K | 7 |
| Small Struct | 🥉 JSON | Marshal | 1.03μs | 704 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.63μs | 2.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 571ns | 696 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.05μs | 5.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.32μs | 2.5K | 55 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.89μs | 1.9K | 42 |
| Small Struct | 🥉 JSON | Unmarshal | 17.58μs | 4.8K | 88 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.13μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 8.81μs | 20.5K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 24.03μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.16μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 32.15μs | 22.0K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 37.21μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 16.71μs | 28.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 36.78μs | 41.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 39.13μs | 33.8K | 622 |
| Medium Payload | 🥈 CBOR | Unmarshal | 63.82μs | 39.7K | 814 |
| Medium Payload | 🥉 JSON | Unmarshal | 192.87μs | 62.7K | 814 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 58.16μs | 39 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 104.58μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 156.99μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 249.51μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 370.19μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 534.03μs | 222.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 176.82μs | 280.1K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 312.03μs | 337.0K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 484.22μs | 345.8K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 595.65μs | 303.1K | 6.2K |
| Large Payload | 🥉 JSON | Unmarshal | 1.79ms | 493.9K | 6.5K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 376ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 888ns | 684 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 919ns | 1.0K | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.64μs | 704 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.89μs | 896 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.73μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 952ns | 312 | 3 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.17μs | 2.1K | 8 |
| Small Struct | 🥉 JSON | Unmarshal | 3.02μs | 488 | 13 |
| Small Struct | 🥈 MessagePack | Unmarshal | 7.72μs | 4.8K | 103 |
| Small Struct | 🥈 CBOR | Unmarshal | 10.81μs | 5.2K | 107 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.78μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 14.15μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 17.84μs | 25.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.98μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 35.61μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 55.18μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.43μs | 33.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 38.74μs | 59.0K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.02μs | 35.5K | 658 |
| Medium Payload | 🥈 CBOR | Unmarshal | 62.60μs | 27.8K | 572 |
| Medium Payload | 🥉 JSON | Unmarshal | 196.49μs | 47.6K | 615 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.14μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 123.81μs | 204.8K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 159.66μs | 215.4K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 196.43μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 311.87μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 438.22μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 245.23μs | 290.0K | 416 |
| Large Payload | 🥉 Sonic | Unmarshal | 349.68μs | 540.6K | 571 |
| Large Payload | 🥈 MessagePack | Unmarshal | 548.32μs | 331.2K | 6.0K |
| Large Payload | 🥈 CBOR | Unmarshal | 702.34μs | 315.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.20ms | 520.1K | 6.7K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 762ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.35μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.88μs | 2.0K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.43μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.59μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 4.38μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.45μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.87μs | 4.0K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.20μs | 2.3K | 50 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.90μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 11.20μs | 3.7K | 52 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.77μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.26μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 23.25μs | 27.3K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 29.73μs | 22.2K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.50μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.09μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.11μs | 31.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.92μs | 34.5K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.10μs | 32.5K | 598 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.33μs | 38.7K | 796 |
| Medium Payload | 🥉 JSON | Unmarshal | 181.38μs | 49.4K | 652 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.56μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 108.74μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 184.99μs | 188.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 283.49μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 303.78μs | 214.5K | 3 |
| Large Payload | 🥉 JSON | Marshal | 401.14μs | 221.6K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 232.26μs | 271.6K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 296.76μs | 401.0K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 490.62μs | 316.5K | 5.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 681.46μs | 329.7K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 559.1K | 7.4K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 851ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 894ns | 1.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.68μs | 2.7K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.69μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.90μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 4.95μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.88μs | 1.8K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.44μs | 816 | 20 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.25μs | 3.1K | 66 |
| Small Struct | 🥉 Sonic | Unmarshal | 6.63μs | 7.5K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 11.33μs | 2.4K | 47 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.82μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 19.65μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 20.57μs | 20.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 38.25μs | 24.6K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 51.77μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 56.29μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 31.65μs | 22.0K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 60.49μs | 64.5K | 79 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 76.10μs | 31.5K | 574 |
| Medium Payload | 🥈 CBOR | Unmarshal | 83.31μs | 32.0K | 659 |
| Medium Payload | 🥉 JSON | Unmarshal | 280.26μs | 61.8K | 832 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 88.98μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 162.50μs | 196.7K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 238.47μs | 188.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 247.57μs | 227.0K | 3 |
| Large Payload | 🥈 MessagePack | Marshal | 482.25μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 609.75μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 423.63μs | 263.6K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 492.90μs | 549.2K | 567 |
| Large Payload | 🥈 MessagePack | Unmarshal | 761.71μs | 364.0K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 849.69μs | 300.6K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.92ms | 560.4K | 7.3K |

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

