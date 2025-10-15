# Multi-Platform Benchmark Aggregation System

## Overview

This system automatically collects benchmark results from multiple CI/CD platforms, aggregates them into a unified report, and generates comprehensive visualizations.

## Architecture

```
CI/CD Pipeline
├── Platform 1 (Linux AMD64)
│   ├── Run benchmarks
│   └── Upload artifacts
├── Platform 2 (macOS ARM64)
│   ├── Run benchmarks
│   └── Upload artifacts
└── Platform 3 (Windows AMD64)
    ├── Run benchmarks
    └── Upload artifacts
         ↓
    Aggregation Job
    ├── Download all artifacts
    ├── Run aggregate_benchmarks.py → MULTI_PLATFORM.md
    ├── Run plot_multi_platform.py → PNG charts
    └── Commit to repository
```

## Components

### 1. **scripts/aggregate_benchmarks.py**

Main aggregation script that:
- ✅ Collects results from all platforms
- ✅ Generates comparison tables
- ✅ Creates performance champion rankings
- ✅ Calculates average improvements
- ✅ Produces detailed per-platform results
- ✅ Outputs unified `MULTI_PLATFORM.md`

**Features:**
- Cross-platform performance comparison
- Winners/champions table with emojis
- Summary statistics with averages
- Detailed platform breakdowns
- Automatic relative path handling

### 2. **scripts/plot_multi_platform.py**

Chart generation script that creates:
- ✅ **Multi-Platform Comparison** - Side-by-side bar charts
- ✅ **Performance Heatmap** - Speedup factors (BEVE vs competitors)
- ✅ **Memory Efficiency** - Allocation comparison

**Visual Features:**
- Color-coded by codec family
- Value labels on bars
- Heatmap with speedup multipliers
- Professional styling with grid
- 150 DPI high-quality output

### 3. **scripts/test_aggregation.sh**

Local testing script that:
- Creates fake artifacts for 3 platforms
- Runs aggregation pipeline
- Validates outputs
- Shows preview of results

**Usage:**
```bash
./scripts/test_aggregation.sh
```

## Output Structure

```
benchmarks/
├── MULTI_PLATFORM.md              # Main unified report
├── charts/                        # Comparison visualizations
│   ├── multi_platform_comparison.png
│   ├── performance_heatmap.png
│   └── memory_comparison.png
├── benchmark-linux-amd64/         # Platform-specific results
│   ├── benchmark.json
│   ├── benchmark.md
│   └── benchmark.png
├── benchmark-macos-arm64/
│   ├── benchmark.json
│   ├── benchmark.md
│   └── benchmark.png
└── benchmark-windows-amd64/
    ├── benchmark.json
    ├── benchmark.md
    └── benchmark.png
```

## Generated Report Sections

### MULTI_PLATFORM.md contains:

1. **📊 Visual Comparisons**
   - Overall performance comparison chart
   - Performance heatmap
   - Memory efficiency chart

2. **🖥️ Tested Platforms**
   - Table with links to artifacts
   - CPU, OS, and architecture info

3. **📊 Cross-Platform Performance Comparison**
   - Marshal performance table
   - Unmarshal performance table
   - All values in human-readable format (μs, ms)

4. **🏆 Performance Champions**
   - Fastest marshal per platform
   - Fastest unmarshal per platform
   - Most memory-efficient per platform
   - Emojis for codec families

5. **📈 Summary Statistics**
   - Total platforms tested
   - Average BEVE vs JSON improvement
   - Platform details (arch, scenarios)

6. **📋 Detailed Platform Results**
   - Per-platform detailed tables
   - Charts embedded
   - Emoji indicators
   - Links to full reports

## CI/CD Integration

### Workflow: `.github/workflows/benchmarks.yml`

```yaml
jobs:
  benchmark:
    # Runs on multiple platforms
    strategy:
      matrix:
        include:
          - label: linux-amd64
          - label: linux-arm64
          - label: macos-arm64
          - label: windows-amd64
    steps:
      - Run benchmarks
      - Upload artifacts

  aggregate:
    needs: benchmark
    steps:
      - Download all artifacts
      - Run aggregate_benchmarks.py
      - Run plot_multi_platform.py
      - Commit results
```

### Key Features:
- ✅ Parallel platform execution
- ✅ Artifact-based result collection
- ✅ Automatic aggregation
- ✅ Chart generation
- ✅ Auto-commit with `[skip ci]`

## Visual Chart Details

### 1. Multi-Platform Comparison
- **Type**: Grouped bar chart
- **Data**: Marshal and Unmarshal times
- **Codecs**: BEVE ZeroCopy, BEVE, JSON, CBOR, MessagePack
- **Format**: Side-by-side comparison
- **Labels**: Time values on bars

### 2. Performance Heatmap
- **Type**: Color-coded matrix
- **Data**: Speedup factor (competitor/BEVE)
- **Scale**: Green (faster) to Red (slower)
- **Values**: Multiplier annotations (e.g., "2.5×")
- **Purpose**: Show how much faster BEVE is

### 3. Memory Comparison
- **Type**: Grouped bar chart
- **Data**: Allocations per operation
- **Metric**: Total memory allocations
- **Purpose**: Show memory efficiency

## Color Scheme

| Codec | Color | Hex |
|-------|-------|-----|
| BEVE ZeroCopy | Green | `#2ecc71` |
| BEVE | Blue | `#3498db` |
| JSON | Red | `#e74c3c` |
| CBOR | Orange | `#f39c12` |
| MessagePack | Purple | `#9b59b6` |

## Emoji Legend

- 🥇 **BEVE family** - Fastest
- 🥈 **CBOR/MessagePack** - Fast
- 🥉 **JSON/Sonic** - Standard
- 💾 **Memory efficient**
- 📄 **Report link**
- 📊 **JSON data**
- 📈 **Chart**

## Testing Locally

```bash
# Clean previous test
rm -rf artifacts dist

# Run test
./scripts/test_aggregation.sh

# View results
cat dist/MULTI_PLATFORM.md
open dist/charts/multi_platform_comparison.png
```

## Dependencies

```bash
# Python dependencies
pip install matplotlib numpy

# Or from requirements.txt
pip install -r requirements.txt
```

## Maintenance

### Adding New Platform
1. Add to matrix in `benchmarks.yml`
2. Ensure benchmark.json follows schema
3. Artifacts will be auto-aggregated

### Modifying Charts
Edit `scripts/plot_multi_platform.py`:
- `create_multi_platform_chart()` - Main comparison
- `create_performance_heatmap()` - Heatmap
- `create_memory_comparison()` - Memory chart

### Changing Report Format
Edit `scripts/aggregate_benchmarks.py`:
- `create_comparison_table()` - Comparison tables
- `create_winners_table()` - Champions table
- `create_summary_stats()` - Statistics section

## Troubleshooting

### No artifacts found
```bash
# Check artifacts directory structure
ls -R artifacts/

# Should contain: benchmark-*/benchmark.json
```

### Chart generation fails
```bash
# Install matplotlib
pip install matplotlib

# Test chart generation
python scripts/plot_multi_platform.py dist/charts
```

### Markdown encoding issues
```bash
# Ensure UTF-8 encoding in scripts
export PYTHONIOENCODING=utf-8
python scripts/aggregate_benchmarks.py
```

## Performance

- **Aggregation**: ~1-2 seconds
- **Chart generation**: ~3-5 seconds
- **Total pipeline**: ~30-45 minutes (includes benchmark runs)

## Future Enhancements

- [ ] Time-series tracking (historical trends)
- [ ] Interactive HTML charts
- [ ] Automatic regression detection
- [ ] Performance badge generation
- [ ] Slack/Discord notifications
- [ ] Comparison against previous runs

## Example Output

See [benchmarks/MULTI_PLATFORM.md](../benchmarks/MULTI_PLATFORM.md) for latest results.

### Sample Tables

**Cross-Platform Comparison:**
| Platform | BEVE | JSON |
|----------|------|------|
| Apple M2 Max | 530ns | 1.43μs |
| AMD EPYC 7763 | 612ns | 1.62μs |

**Performance Champions:**
| Platform | Fastest |
|----------|---------|
| Apple M2 Max | 🥇 BEVE ZeroCopy (530ns) |

## License

MIT License - Part of BEVE-Go project
