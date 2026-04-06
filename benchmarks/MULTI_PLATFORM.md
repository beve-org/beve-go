# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-04-06 05:20:24 UTC

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
| Apple M1 (Virtual) | 1.74μs | 538ns | 5.23μs | 620ns | 2.57μs |
| AMD EPYC 7763 64-Core Processor | 823ns | 571ns | 901ns | 814ns | 1.73μs |
| Neoverse-N2 | 1.31μs | 662ns | 2.73μs | 1.77μs | 2.49μs |
| Unknown CPU | 2.22μs | 345ns | 4.72μs | 2.66μs | 5.59μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.64μs | 6.41μs | 4.34μs | 1.37μs |
| AMD EPYC 7763 64-Core Processor | 1.12μs | 17.29μs | 6.89μs | 4.19μs |
| Neoverse-N2 | 1.29μs | 9.70μs | 5.98μs | 3.85μs |
| Unknown CPU | 1.83μs | 9.20μs | 4.90μs | 3.58μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (538ns) | 🥈 MessagePack (1.37μs) | 💾 BEVE (1 allocs) |
| AMD EPYC 7763 64-Core Processor | 🥇 BEVE ZeroCopy (571ns) | 🥇 BEVE (1.12μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (662ns) | 🥇 BEVE (1.29μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (345ns) | 🥇 BEVE (1.83μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 45.1% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 538ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 620ns | 640 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 1.74μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.57μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 3.53μs | 1.4K | 2 |
| Small Struct | 🥉 JSON | Marshal | 5.23μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.37μs | 640 | 16 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.64μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.15μs | 5.3K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.34μs | 2.1K | 47 |
| Small Struct | 🥉 JSON | Unmarshal | 6.41μs | 1.4K | 32 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.52μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.90μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 14.68μs | 19.1K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.67μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 40.63μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 55.64μs | 22.0K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 26.57μs | 33.1K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.47μs | 42.1K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 45.03μs | 36.6K | 679 |
| Medium Payload | 🥈 CBOR | Unmarshal | 52.16μs | 31.2K | 644 |
| Medium Payload | 🥉 JSON | Unmarshal | 242.72μs | 62.3K | 815 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 55.47μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 94.07μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 180.15μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 214.04μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 334.98μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 382.09μs | 197.5K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 200.37μs | 266.5K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 324.39μs | 323.5K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 593.29μs | 358.8K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 771.39μs | 339.7K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 2.03ms | 516.3K | 6.7K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 571ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 814ns | 640 | 1 |
| Small Struct | 🥇 BEVE | Marshal | 823ns | 896 | 1 |
| Small Struct | 🥉 JSON | Marshal | 901ns | 352 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.64μs | 2.1K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 1.73μs | 2.1K | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.12μs | 1.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.19μs | 3.1K | 65 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.25μs | 7.4K | 10 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.89μs | 3.9K | 83 |
| Small Struct | 🥉 JSON | Unmarshal | 17.29μs | 4.3K | 72 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.79μs | 0 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.11μs | 13.6K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.41μs | 20.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.43μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.30μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 37.30μs | 18.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 21.93μs | 23.3K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.57μs | 64.6K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 52.16μs | 29.3K | 533 |
| Medium Payload | 🥈 CBOR | Unmarshal | 74.61μs | 34.6K | 708 |
| Medium Payload | 🥉 JSON | Unmarshal | 214.24μs | 52.6K | 676 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.92μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.50μs | 188.6K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 168.31μs | 224.3K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 227.15μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 350.04μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 444.48μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 265.92μs | 252.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 378.39μs | 550.6K | 583 |
| Large Payload | 🥈 MessagePack | Unmarshal | 615.78μs | 377.8K | 6.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 712.66μs | 313.4K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.32ms | 539.8K | 7.1K |

[📄 View full report](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 662ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.31μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.77μs | 2.0K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.49μs | 4.1K | 8 |
| Small Struct | 🥉 JSON | Marshal | 2.73μs | 1.5K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 3.52μs | 2.8K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.29μs | 2.1K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.28μs | 5.7K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.85μs | 3.1K | 66 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.98μs | 3.6K | 76 |
| Small Struct | 🥉 JSON | Unmarshal | 9.70μs | 2.4K | 47 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.02μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.41μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.15μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 29.28μs | 22.1K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 31.70μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 38.38μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.14μs | 31.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.48μs | 46.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 55.78μs | 39.8K | 745 |
| Medium Payload | 🥈 CBOR | Unmarshal | 69.39μs | 34.9K | 718 |
| Medium Payload | 🥉 JSON | Unmarshal | 208.15μs | 59.5K | 776 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 68.74μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 104.45μs | 180.3K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 193.29μs | 205.1K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 267.86μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 291.50μs | 208.2K | 3 |
| Large Payload | 🥉 JSON | Marshal | 356.18μs | 197.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 213.79μs | 260.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 275.11μs | 371.9K | 211 |
| Large Payload | 🥈 MessagePack | Unmarshal | 476.14μs | 316.7K | 5.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 682.10μs | 338.4K | 6.9K |
| Large Payload | 🥉 JSON | Unmarshal | 1.90ms | 513.1K | 6.7K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 345ns | 0 | 0 |
| Small Struct | 🥉 Sonic | Marshal | 467ns | 293 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 2.22μs | 2.7K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.66μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 4.72μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 5.59μs | 8.2K | 9 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.83μs | 2.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.26μs | 1.9K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.58μs | 2.1K | 45 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.90μs | 1.9K | 43 |
| Small Struct | 🥉 JSON | Unmarshal | 9.20μs | 2.0K | 35 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.12μs | 5 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.25μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 21.28μs | 27.7K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 22.92μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 39.34μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 48.64μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 35.37μs | 30.9K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 61.12μs | 64.7K | 75 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 71.52μs | 34.6K | 638 |
| Medium Payload | 🥈 CBOR | Unmarshal | 90.70μs | 31.4K | 647 |
| Medium Payload | 🥉 JSON | Unmarshal | 287.75μs | 50.4K | 653 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 79.78μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 120.67μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 168.18μs | 199.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 270.95μs | 188.6K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 306.52μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 453.27μs | 196.9K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 329.88μs | 279.3K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 513.71μs | 541.5K | 574 |
| Large Payload | 🥈 MessagePack | Unmarshal | 729.64μs | 343.7K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 847.16μs | 298.0K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.73ms | 546.9K | 7.2K |

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

