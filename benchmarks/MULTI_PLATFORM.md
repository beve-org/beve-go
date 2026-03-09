# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-03-09 04:51:35 UTC

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
| Apple M1 (Virtual) | 268ns | 235ns | 2.70μs | 460ns | 1.58μs |
| AMD EPYC 7763 64-Core Processor | 1.04μs | 268ns | 4.28μs | 2.26μs | 4.34μs |
| Neoverse-N2 | 1.04μs | 589ns | 4.85μs | 2.10μs | 2.44μs |
| Unknown CPU | 811ns | 349ns | 6.62μs | 2.33μs | 2.76μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 881ns | 14.49μs | 4.40μs | 1.86μs |
| AMD EPYC 7763 64-Core Processor | 1.28μs | 5.06μs | 6.91μs | 3.20μs |
| Neoverse-N2 | 1.72μs | 12.73μs | 5.99μs | 3.05μs |
| Unknown CPU | 2.65μs | 11.28μs | 10.27μs | 2.26μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (235ns) | 🥇 BEVE (881ns) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (268ns) | 🥇 BEVE (1.28μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (589ns) | 🥇 BEVE (1.72μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (349ns) | 🥈 MessagePack (2.26μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 83.0% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 235ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 268ns | 512 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 460ns | 384 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.58μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.70μs | 2.3K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 5.10μs | 3.1K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 881ns | 1.6K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.44μs | 1.5K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.86μs | 2.1K | 45 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.40μs | 3.9K | 82 |
| Small Struct | 🥉 JSON | Unmarshal | 14.49μs | 4.7K | 84 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.07μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.44μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.25μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 23.51μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 28.30μs | 19.3K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 37.15μs | 20.7K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 15.38μs | 22.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 30.61μs | 42.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 38.20μs | 35.9K | 671 |
| Medium Payload | 🥈 CBOR | Unmarshal | 48.43μs | 30.5K | 626 |
| Medium Payload | 🥉 JSON | Unmarshal | 204.00μs | 69.8K | 903 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.15μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 88.34μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 137.09μs | 180.4K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 208.43μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 322.73μs | 205.1K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 414.71μs | 213.8K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 160.27μs | 277.2K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 276.87μs | 350.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 385.64μs | 365.3K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 501.02μs | 311.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 1.74ms | 535.9K | 7.0K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 268ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.04μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.90μs | 2.8K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.26μs | 2.3K | 1 |
| Small Struct | 🥉 JSON | Marshal | 4.28μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 4.34μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.28μs | 1.6K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.20μs | 2.1K | 46 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.27μs | 7.5K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 5.06μs | 936 | 23 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.91μs | 3.9K | 83 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.73μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.21μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.87μs | 21.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 21.65μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 37.16μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 39.72μs | 19.3K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.51μs | 27.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 47.34μs | 67.5K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 56.80μs | 36.7K | 683 |
| Medium Payload | 🥈 CBOR | Unmarshal | 71.06μs | 32.8K | 669 |
| Medium Payload | 🥉 JSON | Unmarshal | 262.03μs | 66.1K | 860 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.76μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 115.29μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 163.82μs | 223.7K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 209.55μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 304.81μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 422.52μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 235.34μs | 262.7K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 365.30μs | 557.9K | 585 |
| Large Payload | 🥈 MessagePack | Unmarshal | 555.83μs | 343.9K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 739.19μs | 339.1K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.15ms | 503.3K | 6.5K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 589ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.04μs | 1.8K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.10μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.44μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.48μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 4.85μs | 3.1K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.72μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.05μs | 2.1K | 47 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.21μs | 5.0K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.99μs | 3.5K | 74 |
| Small Struct | 🥉 JSON | Unmarshal | 12.73μs | 3.9K | 59 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.59μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.92μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.93μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.99μs | 65.8K | 22 |
| Medium Payload | 🥉 Sonic | Marshal | 35.38μs | 25.1K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 42.81μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.27μs | 30.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.67μs | 37.6K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.89μs | 36.8K | 681 |
| Medium Payload | 🥈 CBOR | Unmarshal | 54.48μs | 24.4K | 503 |
| Medium Payload | 🥉 JSON | Unmarshal | 271.86μs | 81.9K | 1.0K |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.33μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 103.87μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 185.45μs | 196.9K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 291.63μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 325.02μs | 231.1K | 3 |
| Large Payload | 🥉 JSON | Marshal | 387.40μs | 213.4K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.14μs | 266.9K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 297.80μs | 406.3K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 486.57μs | 319.8K | 5.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 647.88μs | 311.0K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.11ms | 573.0K | 7.5K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 349ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 811ns | 768 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.22μs | 1.4K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.33μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.76μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 6.62μs | 3.1K | 1 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.26μs | 1.0K | 24 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.65μs | 3.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.38μs | 4.4K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 10.27μs | 4.8K | 103 |
| Small Struct | 🥉 JSON | Unmarshal | 11.28μs | 2.3K | 44 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.66μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.51μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.92μs | 19.4K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 24.48μs | 33.0K | 21 |
| Medium Payload | 🥈 CBOR | Marshal | 28.20μs | 20.5K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 43.88μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 33.17μs | 32.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 50.51μs | 50.6K | 61 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 76.11μs | 33.8K | 620 |
| Medium Payload | 🥈 CBOR | Unmarshal | 108.07μs | 34.2K | 700 |
| Medium Payload | 🥉 JSON | Unmarshal | 321.60μs | 62.9K | 787 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 85.11μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 146.09μs | 213.0K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 163.44μs | 215.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 222.02μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 277.70μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 481.17μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 293.61μs | 260.8K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 467.19μs | 564.5K | 582 |
| Large Payload | 🥈 MessagePack | Unmarshal | 789.23μs | 345.0K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 893.41μs | 313.6K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.81ms | 522.4K | 6.9K |

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

