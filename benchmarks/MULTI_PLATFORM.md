# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-10-16 16:42:25 UTC

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
| Apple M1 (Virtual) | 695ns | 506ns | 2.29μs | 1.86μs | 1.83μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 1.39μs | 550ns | 1.64μs | 1.26μs | 1.53μs |
| Neoverse-N2 | 1.73μs | 722ns | 1.61μs | 2.58μs | 2.19μs |
| Unknown CPU | 1.25μs | 642ns | 5.71μs | 721ns | 3.24μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.32μs | 23.06μs | 897ns | 4.30μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 1.71μs | 12.04μs | 7.29μs | 5.03μs |
| Neoverse-N2 | 1.48μs | 6.51μs | 2.13μs | 4.11μs |
| Unknown CPU | 3.04μs | 27.94μs | 2.27μs | 1.95μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE ZeroCopy (506ns) | 🥈 CBOR (897ns) | 💾 BEVE ZeroCopy (2 allocs) |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 🥇 BEVE ZeroCopy (550ns) | 🥇 BEVE (1.71μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (722ns) | 🥇 BEVE (1.48μs) | 💾 BEVE ZeroCopy (2 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (642ns) | 🥈 MessagePack (1.95μs) | 💾 BEVE ZeroCopy (2 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 38.7% faster

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
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 506ns | 289 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 695ns | 1.6K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.83μs | 4.2K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 1.86μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.29μs | 1.7K | 2 |
| Small Struct | 🥉 Sonic | Marshal | 5.45μs | 3.3K | 3 |
| Small Struct | 🥈 CBOR | Unmarshal | 897ns | 232 | 7 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.32μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.55μs | 6.0K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.30μs | 2.1K | 47 |
| Small Struct | 🥉 JSON | Unmarshal | 23.06μs | 8.0K | 117 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 6.88μs | 134 | 2 |
| Medium Payload | 🥈 CBOR | Marshal | 15.22μs | 18.6K | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 15.59μs | 21.9K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 29.07μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 36.71μs | 19.4K | 9 |
| Medium Payload | 🥉 Sonic | Marshal | 51.11μs | 27.7K | 4 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.62μs | 25.7K | 58 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.90μs | 38.0K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 47.91μs | 37.5K | 700 |
| Medium Payload | 🥈 CBOR | Unmarshal | 51.20μs | 28.7K | 593 |
| Medium Payload | 🥉 JSON | Unmarshal | 224.29μs | 66.2K | 866 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 55.56μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 113.72μs | 197.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 156.00μs | 189.3K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 181.40μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 314.18μs | 197.5K | 9 |
| Large Payload | 🥉 Sonic | Marshal | 376.35μs | 214.7K | 4 |
| Large Payload | 🥇 BEVE | Unmarshal | 153.05μs | 264.0K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 344.07μs | 335.5K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 448.12μs | 357.1K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 589.30μs | 325.4K | 6.6K |
| Large Payload | 🥉 JSON | Unmarshal | 1.63ms | 489.3K | 6.4K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

![Benchmark Chart](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 550ns | 289 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 1.26μs | 1.3K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.39μs | 1.8K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 1.53μs | 2.2K | 7 |
| Small Struct | 🥉 JSON | Marshal | 1.64μs | 912 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 1.71μs | 2.3K | 3 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.71μs | 3.0K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.56μs | 3.8K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.03μs | 4.0K | 84 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.29μs | 4.4K | 95 |
| Small Struct | 🥉 JSON | Unmarshal | 12.04μs | 3.8K | 55 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.41μs | 134 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 13.90μs | 21.9K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 21.87μs | 28.3K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 23.39μs | 24.7K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 25.42μs | 33.1K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 44.99μs | 24.9K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.62μs | 26.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 42.22μs | 60.9K | 77 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 50.48μs | 33.2K | 610 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.41μs | 32.1K | 664 |
| Medium Payload | 🥉 JSON | Unmarshal | 212.51μs | 58.7K | 772 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 76.02μs | 233 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 113.48μs | 180.7K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 168.20μs | 217.8K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 191.13μs | 188.9K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 301.77μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 404.77μs | 205.2K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 235.13μs | 270.9K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 401.52μs | 570.1K | 602 |
| Large Payload | 🥈 MessagePack | Unmarshal | 559.01μs | 360.6K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 622.91μs | 283.9K | 5.8K |
| Large Payload | 🥉 JSON | Unmarshal | 1.96ms | 498.2K | 6.6K |

[📄 View full report](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 722ns | 290 | 2 |
| Small Struct | 🥉 Sonic | Marshal | 944ns | 590 | 3 |
| Small Struct | 🥉 JSON | Marshal | 1.61μs | 1.0K | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.73μs | 3.0K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 2.19μs | 4.2K | 8 |
| Small Struct | 🥈 CBOR | Marshal | 2.58μs | 3.2K | 2 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.48μs | 2.6K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.13μs | 856 | 21 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.11μs | 4.9K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.11μs | 3.3K | 71 |
| Small Struct | 🥉 JSON | Unmarshal | 6.51μs | 1.4K | 32 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.42μs | 141 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 10.11μs | 18.6K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.70μs | 21.9K | 2 |
| Medium Payload | 🥉 Sonic | Marshal | 25.87μs | 18.8K | 4 |
| Medium Payload | 🥈 MessagePack | Marshal | 30.87μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 35.49μs | 20.8K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.26μs | 35.5K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 33.26μs | 48.4K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 47.56μs | 31.4K | 571 |
| Medium Payload | 🥈 CBOR | Unmarshal | 75.66μs | 40.6K | 833 |
| Medium Payload | 🥉 JSON | Unmarshal | 164.88μs | 43.8K | 572 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 70.33μs | 286 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 107.96μs | 181.5K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 199.04μs | 206.9K | 3 |
| Large Payload | 🥈 MessagePack | Marshal | 270.88μs | 526.9K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 291.44μs | 209.0K | 4 |
| Large Payload | 🥉 JSON | Marshal | 396.19μs | 223.0K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 227.22μs | 271.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 293.42μs | 390.7K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 529.03μs | 361.8K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 585.17μs | 278.3K | 5.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.07ms | 561.5K | 7.3K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 642ns | 290 | 2 |
| Small Struct | 🥈 CBOR | Marshal | 721ns | 528 | 2 |
| Small Struct | 🥇 BEVE | Marshal | 1.25μs | 1.1K | 3 |
| Small Struct | 🥉 Sonic | Marshal | 1.85μs | 2.2K | 3 |
| Small Struct | 🥈 MessagePack | Marshal | 3.24μs | 4.2K | 8 |
| Small Struct | 🥉 JSON | Marshal | 5.71μs | 2.8K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.95μs | 784 | 19 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.27μs | 704 | 18 |
| Small Struct | 🥇 BEVE | Unmarshal | 3.04μs | 3.5K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.02μs | 4.4K | 9 |
| Small Struct | 🥉 JSON | Unmarshal | 27.94μs | 7.7K | 105 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.35μs | 138 | 2 |
| Medium Payload | 🥇 BEVE | Marshal | 12.97μs | 18.6K | 3 |
| Medium Payload | 🥉 Sonic | Marshal | 19.91μs | 27.8K | 4 |
| Medium Payload | 🥈 CBOR | Marshal | 21.83μs | 19.2K | 2 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.41μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 42.72μs | 19.4K | 9 |
| Medium Payload | 🥇 BEVE | Unmarshal | 29.22μs | 30.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 45.05μs | 56.2K | 73 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 60.68μs | 30.6K | 557 |
| Medium Payload | 🥈 CBOR | Unmarshal | 92.73μs | 35.0K | 723 |
| Medium Payload | 🥉 JSON | Unmarshal | 277.94μs | 63.2K | 805 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 84.10μs | 259 | 2 |
| Large Payload | 🥇 BEVE | Marshal | 123.94μs | 197.2K | 3 |
| Large Payload | 🥉 Sonic | Marshal | 190.66μs | 228.4K | 4 |
| Large Payload | 🥈 CBOR | Marshal | 219.82μs | 198.4K | 2 |
| Large Payload | 🥈 MessagePack | Marshal | 285.69μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 479.79μs | 215.3K | 9 |
| Large Payload | 🥇 BEVE | Unmarshal | 280.36μs | 284.8K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 425.42μs | 539.6K | 590 |
| Large Payload | 🥈 MessagePack | Unmarshal | 696.19μs | 366.5K | 6.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 840.45μs | 309.5K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.46ms | 531.6K | 6.9K |

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

