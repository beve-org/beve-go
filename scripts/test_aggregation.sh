#!/usr/bin/env bash
# Test script for aggregate_benchmarks.py
# Creates fake artifacts to test aggregation logic

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "🧪 Testing Multi-Platform Benchmark Aggregation"
echo "================================================"
echo ""

# Clean up previous test
rm -rf "${ROOT_DIR}/artifacts" "${ROOT_DIR}/dist"

# Create fake artifacts structure
mkdir -p "${ROOT_DIR}/artifacts/benchmark-macos-arm64"
mkdir -p "${ROOT_DIR}/artifacts/benchmark-linux-amd64"
mkdir -p "${ROOT_DIR}/artifacts/benchmark-windows-amd64"

# Create fake benchmark data for macOS ARM64
cat > "${ROOT_DIR}/artifacts/benchmark-macos-arm64/benchmark.json" <<'JSON'
{
  "environment": {
    "cpu": "Apple M2 Max",
    "os": "Darwin",
    "arch": "arm64"
  },
  "results": [
    {
      "scenario": "Small Struct",
      "codec": "BEVE ZeroCopy",
      "operation": "Marshal",
      "ns_per_op": 530,
      "bytes_per_op": 290,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Marshal",
      "ns_per_op": 1012,
      "bytes_per_op": 2596,
      "allocs_per_op": 3
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Marshal",
      "ns_per_op": 1434,
      "bytes_per_op": 1168,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Unmarshal",
      "ns_per_op": 719,
      "bytes_per_op": 1593,
      "allocs_per_op": 4
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Unmarshal",
      "ns_per_op": 7092,
      "bytes_per_op": 2408,
      "allocs_per_op": 47
    }
  ]
}
JSON

# Create fake markdown
echo "# macOS ARM64 Benchmark" > "${ROOT_DIR}/artifacts/benchmark-macos-arm64/benchmark.md"

# Create fake PNG (empty file)
touch "${ROOT_DIR}/artifacts/benchmark-macos-arm64/benchmark.png"

# Create fake benchmark data for Linux AMD64
cat > "${ROOT_DIR}/artifacts/benchmark-linux-amd64/benchmark.json" <<'JSON'
{
  "environment": {
    "cpu": "AMD EPYC 7763",
    "os": "Linux",
    "arch": "amd64"
  },
  "results": [
    {
      "scenario": "Small Struct",
      "codec": "BEVE ZeroCopy",
      "operation": "Marshal",
      "ns_per_op": 612,
      "bytes_per_op": 290,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Marshal",
      "ns_per_op": 1150,
      "bytes_per_op": 2596,
      "allocs_per_op": 3
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Marshal",
      "ns_per_op": 1620,
      "bytes_per_op": 1168,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Unmarshal",
      "ns_per_op": 820,
      "bytes_per_op": 1593,
      "allocs_per_op": 4
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Unmarshal",
      "ns_per_op": 8100,
      "bytes_per_op": 2408,
      "allocs_per_op": 47
    }
  ]
}
JSON

echo "# Linux AMD64 Benchmark" > "${ROOT_DIR}/artifacts/benchmark-linux-amd64/benchmark.md"
touch "${ROOT_DIR}/artifacts/benchmark-linux-amd64/benchmark.png"

# Create fake benchmark data for Windows AMD64
cat > "${ROOT_DIR}/artifacts/benchmark-windows-amd64/benchmark.json" <<'JSON'
{
  "environment": {
    "cpu": "Intel i7-12700K",
    "os": "Windows",
    "arch": "amd64"
  },
  "results": [
    {
      "scenario": "Small Struct",
      "codec": "BEVE ZeroCopy",
      "operation": "Marshal",
      "ns_per_op": 580,
      "bytes_per_op": 290,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Marshal",
      "ns_per_op": 1080,
      "bytes_per_op": 2596,
      "allocs_per_op": 3
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Marshal",
      "ns_per_op": 1520,
      "bytes_per_op": 1168,
      "allocs_per_op": 2
    },
    {
      "scenario": "Small Struct",
      "codec": "BEVE",
      "operation": "Unmarshal",
      "ns_per_op": 760,
      "bytes_per_op": 1593,
      "allocs_per_op": 4
    },
    {
      "scenario": "Small Struct",
      "codec": "JSON",
      "operation": "Unmarshal",
      "ns_per_op": 7500,
      "bytes_per_op": 2408,
      "allocs_per_op": 47
    }
  ]
}
JSON

echo "# Windows AMD64 Benchmark" > "${ROOT_DIR}/artifacts/benchmark-windows-amd64/benchmark.md"
touch "${ROOT_DIR}/artifacts/benchmark-windows-amd64/benchmark.png"

echo "✅ Created fake artifacts structure"
echo ""
echo "📁 Artifacts directory:"
tree "${ROOT_DIR}/artifacts" 2>/dev/null || find "${ROOT_DIR}/artifacts" -type f
echo ""

# Run the aggregation script
echo "🔧 Running aggregation script..."
echo ""
cd "${ROOT_DIR}"
python scripts/aggregate_benchmarks.py

echo ""
echo "🎨 Generating multi-platform charts..."
python scripts/plot_multi_platform.py dist/charts

echo ""
echo "✅ Aggregation complete!"
echo ""
echo "📊 Generated files:"
if [ -f "dist/MULTI_PLATFORM.md" ]; then
    echo "  ✅ dist/MULTI_PLATFORM.md ($(wc -l < dist/MULTI_PLATFORM.md) lines)"
else
    echo "  ❌ dist/MULTI_PLATFORM.md NOT FOUND"
fi

if [ -d "dist/benchmarks" ]; then
    echo "  ✅ dist/benchmarks/ directory"
    ls -1 dist/benchmarks/
else
    echo "  ❌ dist/benchmarks/ NOT FOUND"
fi

if [ -d "dist/charts" ]; then
    echo "  ✅ dist/charts/ directory"
    ls -lh dist/charts/
else
    echo "  ❌ dist/charts/ NOT FOUND"
fi

echo ""
echo "📄 Preview of MULTI_PLATFORM.md:"
echo "================================"
head -50 dist/MULTI_PLATFORM.md 2>/dev/null || echo "❌ Could not read file"

echo ""
echo "🧹 Cleaning up test artifacts..."
# Uncomment to clean up
# rm -rf "${ROOT_DIR}/artifacts" "${ROOT_DIR}/dist"

echo "✅ Test complete!"
