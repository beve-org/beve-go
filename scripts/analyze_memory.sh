#!/bin/bash
# Memory Profiling Script for BEVE Go

set -e

echo "🔍 BEVE Go - Memory Allocation Analysis"
echo "========================================"
echo ""

cd /Users/mapletechnologies/go-workspace/src/github.com/meftunca/beve-go

# Run benchmarks with memory profiling
echo "📊 Running benchmarks with memory profiling..."
go test -run=^$ -bench=. -benchmem -benchtime=5000x -memprofile=mem.prof > bench_mem.txt 2>&1

echo ""
echo "✅ Benchmark completed. Analyzing results..."
echo ""

# Top allocations by bytes
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🔥 TOP 20 BENCHMARKS BY MEMORY USAGE (Bytes/op)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
grep "^Benchmark" bench_mem.txt | awk '{
    name=$1
    ns=$3
    bytes=$5
    allocs=$7
    print bytes, allocs, ns, name
}' | sort -rn | head -20 | awk '{
    printf "%-60s %10s B/op  %6s allocs/op  %10s ns/op\n", $4, $1, $2, $3
}'

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🔥 TOP 20 BENCHMARKS BY ALLOCATION COUNT (allocs/op)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
grep "^Benchmark" bench_mem.txt | awk '{
    name=$1
    ns=$3
    bytes=$5
    allocs=$7
    print allocs, bytes, ns, name
}' | sort -rn | head -20 | awk '{
    printf "%-60s %6s allocs/op  %10s B/op  %10s ns/op\n", $4, $1, $2, $3
}'

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📈 MEMORY HOTSPOT ANALYSIS (pprof)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
if [ -f mem.prof ]; then
    echo ""
    echo "Top 20 functions by cumulative allocation:"
    go tool pprof -top -cum -nodecount=20 mem.prof 2>/dev/null | grep -v "^$" | head -25
    
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "Top 20 functions by flat allocation:"
    go tool pprof -top -flat -nodecount=20 mem.prof 2>/dev/null | grep -v "^$" | head -25
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "💡 OPTIMIZATION TARGETS"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

echo ""
echo "Looking for high-allocation benchmarks (>1000 B/op or >10 allocs/op)..."
grep "^Benchmark" bench_mem.txt | awk '{
    name=$1
    bytes=$5
    allocs=$7
    if (bytes+0 > 1000 || allocs+0 > 10) {
        print name, bytes, allocs
    }
}' | awk '{
    printf "⚠️  %-60s %10s B/op  %6s allocs/op\n", $1, $2, $3
}'

echo ""
echo "✅ Analysis complete!"
echo ""
echo "Files generated:"
echo "  - bench_mem.txt (full benchmark output)"
echo "  - mem.prof (pprof memory profile)"
echo ""
echo "To explore interactively:"
echo "  go tool pprof -http=:8080 mem.prof"
