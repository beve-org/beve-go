# Benchmark Scripts

## bench.sh - Parallel Benchmark Runner

Runs all BEVE benchmarks **in parallel** for significantly faster execution.

### Usage

```bash
# Run with default settings (8 parallel jobs)
./scripts/bench.sh

# Run with custom number of parallel jobs
BENCH_MAX_JOBS=4 ./scripts/bench.sh    # 4 parallel jobs
BENCH_MAX_JOBS=16 ./scripts/bench.sh   # 16 parallel jobs (for high-core CPUs)
```

### Performance

**Sequential (old):** ~15-20 minutes for full benchmark suite  
**Parallel (new):** ~2-4 minutes with 8 cores (5-8× faster!) ⚡

### Environment Variables

- `BENCH_MAX_JOBS`: Maximum number of parallel benchmark jobs (default: 8)

### Output

Generates two files in `benchmarks/`:
- `latest.md` - Human-readable markdown report
- `latest.json` - Machine-readable JSON data

### How It Works

1. **Job Control**: Maintains up to `MAX_JOBS` parallel benchmark processes
2. **Result Collection**: Each benchmark writes to a unique JSON file
3. **Aggregation**: All results are combined into final JSON/markdown output
4. **Progress Tracking**: Shows `[START]` and `[DONE]` messages for each benchmark

### Example Output

```
🚀 BEVE Parallel Benchmark Runner
📊 Running up to 8 benchmarks in parallel

[START] Small Struct · Marshal · BEVE
[START] Small Struct · Marshal · JSON
[START] Small Struct · Marshal · CBOR
...
[DONE] Small Struct · Marshal · BEVE
[DONE] Small Struct · Marshal · JSON
...
Waiting for all benchmarks to complete...
Combining results...

Benchmark report written to benchmarks/latest.md
Benchmark JSON written to benchmarks/latest.json
```

### Tips

- **CPU cores:** Set `MAX_JOBS` to your CPU core count for best performance
- **Memory:** Each job uses ~500MB RAM, monitor if using many parallel jobs
- **Accuracy:** Parallel execution may have slightly more variance than sequential

### Technical Details

- Uses bash job control (`&` and `wait`)
- Thread-safe result collection via unique files
- Atomic JSON aggregation at the end
- Proper cleanup on exit (SIGTERM, SIGINT)
