# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2026-06-29 07:26:30 UTC

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
| Apple M1 (Virtual) | 652ns | 419ns | 3.34μs | 316ns | 2.36μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 1.52μs | 457ns | 3.36μs | 1.96μs | 3.38μs |
| Neoverse-N2 | 1.38μs | 262ns | 2.40μs | 2.44μs | 1.43μs |
| Unknown CPU | 535ns | 311ns | 891ns | 2.52μs | 2.17μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 877ns | 9.13μs | 1.26μs | 3.87μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 2.27μs | 23.84μs | 6.41μs | 3.45μs |
| Neoverse-N2 | 1.15μs | 21.77μs | 2.38μs | 1.87μs |
| Unknown CPU | 661ns | 15.29μs | 1.64μs | 5.12μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥈 CBOR (316ns) | 🥇 BEVE (877ns) | 💾 BEVE (1 allocs) |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 🥇 BEVE ZeroCopy (457ns) | 🥇 BEVE (2.27μs) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (262ns) | 🥉 Sonic (1.15μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (311ns) | 🥇 BEVE (661ns) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 54.4% faster

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
| Small Struct | 🥈 CBOR | Marshal | 316ns | 256 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 419ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 652ns | 1.5K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.21μs | 1.2K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 2.36μs | 8.2K | 9 |
| Small Struct | 🥉 JSON | Marshal | 3.34μs | 2.7K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 877ns | 1.8K | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.26μs | 616 | 16 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.12μs | 5.9K | 6 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.87μs | 4.7K | 100 |
| Small Struct | 🥉 JSON | Unmarshal | 9.13μs | 3.7K | 51 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.18μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 7.41μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 12.90μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.25μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 30.85μs | 20.7K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 41.78μs | 24.8K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 18.52μs | 28.8K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 26.59μs | 37.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 38.83μs | 39.8K | 744 |
| Medium Payload | 🥈 CBOR | Unmarshal | 47.45μs | 32.2K | 662 |
| Medium Payload | 🥉 JSON | Unmarshal | 176.52μs | 48.3K | 662 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 53.98μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 75.47μs | 188.5K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 138.96μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 186.69μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 322.56μs | 213.4K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 379.02μs | 205.2K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 169.89μs | 269.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 261.04μs | 317.7K | 207 |
| Large Payload | 🥈 MessagePack | Unmarshal | 385.77μs | 368.5K | 6.8K |
| Large Payload | 🥈 CBOR | Unmarshal | 491.18μs | 308.5K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 1.78ms | 539.7K | 7.2K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

![Benchmark Chart](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 457ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.52μs | 2.0K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.96μs | 1.8K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.96μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.36μs | 1.4K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.38μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 2.27μs | 2.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 3.45μs | 1.5K | 34 |
| Small Struct | 🥉 Sonic | Unmarshal | 4.01μs | 4.2K | 9 |
| Small Struct | 🥈 CBOR | Unmarshal | 6.41μs | 3.1K | 67 |
| Small Struct | 🥉 JSON | Unmarshal | 23.84μs | 4.8K | 88 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.19μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.10μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 15.26μs | 19.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 18.64μs | 18.5K | 1 |
| Medium Payload | 🥉 JSON | Marshal | 35.27μs | 18.7K | 8 |
| Medium Payload | 🥈 MessagePack | Marshal | 36.55μs | 65.8K | 22 |
| Medium Payload | 🥇 BEVE | Unmarshal | 23.23μs | 25.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 41.82μs | 61.3K | 70 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 53.44μs | 35.5K | 663 |
| Medium Payload | 🥈 CBOR | Unmarshal | 61.10μs | 27.9K | 571 |
| Medium Payload | 🥉 JSON | Unmarshal | 169.37μs | 44.4K | 577 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.23μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 125.25μs | 196.7K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 168.16μs | 209.2K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 206.68μs | 205.0K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 303.71μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 423.01μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 245.84μs | 264.7K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 370.37μs | 521.6K | 578 |
| Large Payload | 🥈 MessagePack | Unmarshal | 525.40μs | 328.7K | 5.9K |
| Large Payload | 🥈 CBOR | Unmarshal | 712.37μs | 329.8K | 6.7K |
| Large Payload | 🥉 JSON | Unmarshal | 2.04ms | 527.7K | 6.9K |

[📄 View full report](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 262ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 1.38μs | 2.7K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.43μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 1.48μs | 933 | 2 |
| Small Struct | 🥉 JSON | Marshal | 2.40μs | 1.4K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.44μs | 3.1K | 1 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.15μs | 1.1K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.15μs | 1.5K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.87μs | 1.0K | 24 |
| Small Struct | 🥈 CBOR | Unmarshal | 2.38μs | 1.1K | 25 |
| Small Struct | 🥉 JSON | Unmarshal | 21.77μs | 7.5K | 99 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.09μs | 6 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.53μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 16.57μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 30.20μs | 22.3K | 3 |
| Medium Payload | 🥈 MessagePack | Marshal | 32.87μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 38.00μs | 20.7K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.90μs | 29.2K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 32.43μs | 45.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 49.41μs | 34.0K | 624 |
| Medium Payload | 🥈 CBOR | Unmarshal | 57.14μs | 26.8K | 549 |
| Medium Payload | 🥉 JSON | Unmarshal | 169.45μs | 45.5K | 597 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 73.44μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 107.84μs | 188.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 188.09μs | 188.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 288.16μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 326.28μs | 232.4K | 3 |
| Large Payload | 🥉 JSON | Marshal | 378.97μs | 205.2K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 229.86μs | 274.4K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 287.78μs | 376.0K | 213 |
| Large Payload | 🥈 MessagePack | Unmarshal | 522.08μs | 352.5K | 6.4K |
| Large Payload | 🥈 CBOR | Unmarshal | 659.19μs | 318.2K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.99ms | 542.5K | 7.0K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 311ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 535ns | 256 | 1 |
| Small Struct | 🥉 JSON | Marshal | 891ns | 288 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.17μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 2.37μs | 2.1K | 2 |
| Small Struct | 🥈 CBOR | Marshal | 2.52μs | 2.3K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 661ns | 312 | 3 |
| Small Struct | 🥈 CBOR | Unmarshal | 1.64μs | 328 | 10 |
| Small Struct | 🥉 Sonic | Unmarshal | 3.09μs | 3.7K | 9 |
| Small Struct | 🥈 MessagePack | Unmarshal | 5.12μs | 2.8K | 60 |
| Small Struct | 🥉 JSON | Unmarshal | 15.29μs | 3.8K | 57 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 8.36μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 12.43μs | 16.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 19.66μs | 24.9K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 25.76μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.48μs | 33.0K | 21 |
| Medium Payload | 🥉 JSON | Marshal | 51.87μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 32.71μs | 31.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 48.17μs | 55.1K | 74 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 71.87μs | 36.2K | 667 |
| Medium Payload | 🥈 CBOR | Unmarshal | 88.34μs | 30.3K | 627 |
| Medium Payload | 🥉 JSON | Unmarshal | 251.55μs | 47.7K | 654 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 77.75μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 119.21μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 178.44μs | 223.6K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 236.60μs | 188.5K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 290.79μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 498.18μs | 205.1K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 299.50μs | 275.5K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 466.12μs | 556.5K | 599 |
| Large Payload | 🥈 MessagePack | Unmarshal | 684.78μs | 357.9K | 6.5K |
| Large Payload | 🥈 CBOR | Unmarshal | 868.44μs | 306.9K | 6.3K |
| Large Payload | 🥉 JSON | Unmarshal | 2.50ms | 504.7K | 6.6K |

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

