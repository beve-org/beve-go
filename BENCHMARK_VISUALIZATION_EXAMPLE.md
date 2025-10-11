# Benchmark Results Visualization - Example

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmarks/linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 978.5 | 723 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 1110 | 144 | 1 |
| Small Struct | Sonic | Marshal | 1311 | 1781 | 3 |
| Small Struct | JSON | Marshal | 1351 | 624 | 2 |
| Small Struct | MessagePack | Marshal | 1783 | 2176 | 7 |
| Small Struct | CBOR | Marshal | 3037 | 2834 | 2 |
| Small Struct | BEVE | Unmarshal | 1035 | 824 | 4 |
| Small Struct | CBOR | Unmarshal | 1637 | 424 | 12 |

---

## Apple M1 (Virtual) — Darwin

![Benchmark Chart](benchmarks/darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 959.6 | 145 | 1 |
| Small Struct | BEVE | Marshal | 1150 | 1938 | 2 |
| Small Struct | Sonic | Marshal | 1353 | 750 | 3 |

---

## Benefits

### Visual Hierarchy
1. **Chart First** - Quick visual comparison
2. **Caption** - Context about what the chart shows
3. **Detailed Table** - Exact metrics for analysis

### User Experience
- ✅ Easier to spot performance trends
- ✅ Quick platform comparison at a glance
- ✅ Detailed metrics still available below
- ✅ Better for README/documentation

### Before
```
## Platform

| Scenario | Codec | ... |
```

### After
```
## Platform

![Chart](path/to/chart.png)
_Performance visualization..._

### Detailed Results
| Scenario | Codec | ... |
```

This structure makes benchmark results more accessible and easier to understand!
