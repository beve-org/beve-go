# 🚀 BEVE-Go Multi-Platform Benchmark Results

**Generated:** 2025-11-03 03:50:49 UTC

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
| Apple M1 (Virtual) | 225ns | 562ns | 1.12μs | 697ns | 2.51μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 398ns | 385ns | 3.87μs | 1.31μs | 1.49μs |
| Neoverse-N2 | 800ns | 392ns | 2.50μs | 2.37μs | 3.62μs |
| Unknown CPU | 796ns | 486ns | 1.89μs | 2.29μs | 3.04μs |

### Unmarshal Performance (Small Struct)

| Platform | BEVE | JSON | CBOR | MessagePack |
|----------|------|------|------|-------------|
| Apple M1 (Virtual) | 1.55μs | 22.25μs | 4.16μs | 1.44μs |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 850ns | 7.57μs | 7.13μs | 4.14μs |
| Neoverse-N2 | 1.73μs | 8.68μs | 5.56μs | 2.88μs |
| Unknown CPU | 1.05μs | 23.85μs | 3.39μs | 4.07μs |

## 🏆 Performance Champions

| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |
|----------|----------------|-------------------|------------------|
| Apple M1 (Virtual) | 🥇 BEVE (225ns) | 🥈 MessagePack (1.44μs) | 💾 BEVE (1 allocs) |
| Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz | 🥇 BEVE ZeroCopy (385ns) | 🥇 BEVE (850ns) | 💾 BEVE (1 allocs) |
| Neoverse-N2 | 🥇 BEVE ZeroCopy (392ns) | 🥉 Sonic (1.63μs) | 💾 BEVE (1 allocs) |
| Unknown CPU | 🥇 BEVE ZeroCopy (486ns) | 🥇 BEVE (1.05μs) | 💾 BEVE (1 allocs) |

## 📈 Summary Statistics

**Total Platforms Tested:** 4

**Average BEVE vs JSON Improvement:** 73.9% faster

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
| Small Struct | 🥇 BEVE | Marshal | 225ns | 160 | 1 |
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 562ns | 0 | 0 |
| Small Struct | 🥈 CBOR | Marshal | 697ns | 576 | 1 |
| Small Struct | 🥉 JSON | Marshal | 1.12μs | 512 | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 2.51μs | 4.1K | 8 |
| Small Struct | 🥉 Sonic | Marshal | 6.57μs | 3.1K | 2 |
| Small Struct | 🥈 MessagePack | Unmarshal | 1.44μs | 600 | 15 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.55μs | 3.4K | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 2.13μs | 2.8K | 6 |
| Small Struct | 🥈 CBOR | Unmarshal | 4.16μs | 2.1K | 47 |
| Small Struct | 🥉 JSON | Unmarshal | 22.25μs | 7.7K | 107 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 5.85μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.37μs | 18.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 15.38μs | 21.8K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 26.84μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.59μs | 27.5K | 8 |
| Medium Payload | 🥉 Sonic | Marshal | 52.00μs | 24.9K | 3 |
| Medium Payload | 🥇 BEVE | Unmarshal | 20.91μs | 30.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 35.09μs | 46.3K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 35.44μs | 31.5K | 574 |
| Medium Payload | 🥈 CBOR | Unmarshal | 55.02μs | 36.7K | 758 |
| Medium Payload | 🥉 JSON | Unmarshal | 184.27μs | 42.3K | 560 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 59.44μs | 52 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 82.82μs | 196.6K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 157.19μs | 196.7K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 210.80μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 356.99μs | 213.3K | 8 |
| Large Payload | 🥉 Sonic | Marshal | 501.64μs | 222.4K | 3 |
| Large Payload | 🥇 BEVE | Unmarshal | 170.11μs | 270.4K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 324.35μs | 336.0K | 209 |
| Large Payload | 🥈 MessagePack | Unmarshal | 507.43μs | 343.7K | 6.3K |
| Large Payload | 🥈 CBOR | Unmarshal | 599.77μs | 313.5K | 6.4K |
| Large Payload | 🥉 JSON | Unmarshal | 2.01ms | 521.6K | 6.9K |

[📄 View full report](benchmark-darwin-apple-m1-virtual/benchmark.md)

### Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

![Benchmark Chart](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 385ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 398ns | 416 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 1.31μs | 1.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 1.49μs | 2.1K | 7 |
| Small Struct | 🥉 Sonic | Marshal | 2.09μs | 2.8K | 2 |
| Small Struct | 🥉 JSON | Marshal | 3.87μs | 2.0K | 1 |
| Small Struct | 🥇 BEVE | Unmarshal | 850ns | 632 | 4 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.57μs | 1.9K | 8 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.14μs | 3.2K | 67 |
| Small Struct | 🥈 CBOR | Unmarshal | 7.13μs | 4.2K | 89 |
| Small Struct | 🥉 JSON | Unmarshal | 7.57μs | 2.1K | 36 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.57μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 11.13μs | 18.4K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 18.92μs | 25.3K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 20.64μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.64μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 46.56μs | 24.8K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 24.06μs | 28.3K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 44.46μs | 63.0K | 73 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 59.44μs | 40.4K | 753 |
| Medium Payload | 🥈 CBOR | Unmarshal | 65.94μs | 30.7K | 631 |
| Medium Payload | 🥉 JSON | Unmarshal | 188.87μs | 49.7K | 669 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 69.48μs | 26 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 107.33μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 160.97μs | 216.0K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 210.99μs | 213.2K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 290.35μs | 526.8K | 115 |
| Large Payload | 🥉 JSON | Marshal | 422.39μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 243.92μs | 282.6K | 419 |
| Large Payload | 🥉 Sonic | Unmarshal | 383.47μs | 562.2K | 593 |
| Large Payload | 🥈 MessagePack | Unmarshal | 542.32μs | 341.6K | 6.2K |
| Large Payload | 🥈 CBOR | Unmarshal | 666.07μs | 296.6K | 6.1K |
| Large Payload | 🥉 JSON | Unmarshal | 2.08ms | 550.2K | 7.2K |

[📄 View full report](benchmark-linux-intel-r-xeon-r-platinum-8370c-cpu-2-80ghz/benchmark.md)

### Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 392ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 800ns | 1.3K | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.37μs | 2.7K | 1 |
| Small Struct | 🥉 JSON | Marshal | 2.50μs | 1.4K | 1 |
| Small Struct | 🥉 Sonic | Marshal | 2.56μs | 1.9K | 2 |
| Small Struct | 🥈 MessagePack | Marshal | 3.62μs | 8.2K | 9 |
| Small Struct | 🥉 Sonic | Unmarshal | 1.63μs | 2.1K | 6 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.73μs | 3.4K | 4 |
| Small Struct | 🥈 MessagePack | Unmarshal | 2.88μs | 2.0K | 43 |
| Small Struct | 🥈 CBOR | Unmarshal | 5.56μs | 3.2K | 68 |
| Small Struct | 🥉 JSON | Unmarshal | 8.68μs | 2.2K | 42 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 7.68μs | 3 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 9.56μs | 16.4K | 1 |
| Medium Payload | 🥈 CBOR | Marshal | 17.27μs | 18.4K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 22.60μs | 33.0K | 21 |
| Medium Payload | 🥉 Sonic | Marshal | 28.07μs | 19.5K | 3 |
| Medium Payload | 🥉 JSON | Marshal | 39.62μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 22.34μs | 26.7K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 29.57μs | 40.2K | 33 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 57.58μs | 41.3K | 775 |
| Medium Payload | 🥈 CBOR | Unmarshal | 64.99μs | 31.6K | 653 |
| Medium Payload | 🥉 JSON | Unmarshal | 164.92μs | 44.3K | 571 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 65.38μs | 65 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 101.82μs | 180.4K | 1 |
| Large Payload | 🥈 CBOR | Marshal | 168.97μs | 172.3K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 272.24μs | 526.8K | 115 |
| Large Payload | 🥉 Sonic | Marshal | 306.53μs | 214.5K | 3 |
| Large Payload | 🥉 JSON | Marshal | 391.49μs | 213.3K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 215.90μs | 252.1K | 417 |
| Large Payload | 🥉 Sonic | Unmarshal | 297.76μs | 407.7K | 205 |
| Large Payload | 🥈 MessagePack | Unmarshal | 529.72μs | 359.9K | 6.6K |
| Large Payload | 🥈 CBOR | Unmarshal | 656.15μs | 319.4K | 6.5K |
| Large Payload | 🥉 JSON | Unmarshal | 1.91ms | 519.4K | 6.7K |

[📄 View full report](benchmark-linux-neoverse-n2/benchmark.md)

### Unknown CPU — Windows

![Benchmark Chart](benchmark-windows-unknown-cpu/benchmark.png)

_Performance visualization: lower is better._

| Scenario | Codec | Operation | Time | Memory | Allocations |
|----------|-------|-----------|------|--------|-------------|
| Small Struct | 🥇 BEVE ZeroCopy | Marshal | 486ns | 0 | 0 |
| Small Struct | 🥇 BEVE | Marshal | 796ns | 896 | 1 |
| Small Struct | 🥉 Sonic | Marshal | 1.50μs | 1.6K | 2 |
| Small Struct | 🥉 JSON | Marshal | 1.89μs | 768 | 1 |
| Small Struct | 🥈 CBOR | Marshal | 2.29μs | 2.3K | 1 |
| Small Struct | 🥈 MessagePack | Marshal | 3.04μs | 4.1K | 8 |
| Small Struct | 🥇 BEVE | Unmarshal | 1.05μs | 952 | 4 |
| Small Struct | 🥈 CBOR | Unmarshal | 3.39μs | 1.2K | 28 |
| Small Struct | 🥈 MessagePack | Unmarshal | 4.07μs | 2.4K | 52 |
| Small Struct | 🥉 Sonic | Unmarshal | 5.93μs | 7.8K | 10 |
| Small Struct | 🥉 JSON | Unmarshal | 23.85μs | 4.7K | 83 |
| Medium Payload | 🥇 BEVE ZeroCopy | Marshal | 9.24μs | 1 | 0 |
| Medium Payload | 🥇 BEVE | Marshal | 13.94μs | 21.8K | 1 |
| Medium Payload | 🥉 Sonic | Marshal | 16.84μs | 22.0K | 3 |
| Medium Payload | 🥈 CBOR | Marshal | 23.22μs | 20.5K | 1 |
| Medium Payload | 🥈 MessagePack | Marshal | 34.69μs | 65.8K | 22 |
| Medium Payload | 🥉 JSON | Marshal | 49.85μs | 22.0K | 8 |
| Medium Payload | 🥇 BEVE | Unmarshal | 25.91μs | 27.4K | 59 |
| Medium Payload | 🥉 Sonic | Unmarshal | 53.56μs | 70.0K | 82 |
| Medium Payload | 🥈 MessagePack | Unmarshal | 58.95μs | 30.3K | 549 |
| Medium Payload | 🥈 CBOR | Unmarshal | 91.30μs | 35.0K | 721 |
| Medium Payload | 🥉 JSON | Unmarshal | 241.13μs | 48.8K | 648 |
| Large Payload | 🥇 BEVE ZeroCopy | Marshal | 81.02μs | 79 | 0 |
| Large Payload | 🥇 BEVE | Marshal | 111.43μs | 188.5K | 1 |
| Large Payload | 🥉 Sonic | Marshal | 142.92μs | 189.8K | 3 |
| Large Payload | 🥈 CBOR | Marshal | 212.58μs | 196.8K | 1 |
| Large Payload | 🥈 MessagePack | Marshal | 269.29μs | 526.7K | 115 |
| Large Payload | 🥉 JSON | Marshal | 487.14μs | 221.5K | 8 |
| Large Payload | 🥇 BEVE | Unmarshal | 273.03μs | 271.2K | 418 |
| Large Payload | 🥉 Sonic | Unmarshal | 438.83μs | 573.0K | 608 |
| Large Payload | 🥈 MessagePack | Unmarshal | 616.12μs | 318.4K | 5.7K |
| Large Payload | 🥈 CBOR | Unmarshal | 816.17μs | 294.9K | 6.0K |
| Large Payload | 🥉 JSON | Unmarshal | 2.63ms | 533.0K | 7.1K |

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

