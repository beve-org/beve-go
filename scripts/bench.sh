#!/usr/bin/env bash
# BEVE Benchmark Runner - Simplified Batch Edition
# 
# Runs all benchmarks in one go and generates comprehensive reports.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_DIR="${ROOT_DIR}/benchmarks"

ITERATIONS="${BENCH_ITERATIONS:-20000}"
TIMEOUT="${BENCH_TIMEOUT:-20m}"

echo "🚀 BEVE Benchmark Runner" >&2
echo "📊 Iterations: ${ITERATIONS}×, Timeout: ${TIMEOUT}" >&2
echo "" >&2

mkdir -p "${OUT_DIR}"

# Run ALL benchmarks in a single efficient command
echo "Running benchmarks..." >&2
benchmark_output=$(cd "${ROOT_DIR}" && go test \
  -bench="Benchmark(SmallStruct|Medium|Large)_(BEVE|JSON|Sonic|CBOR|MessagePack)_(Marshal|Unmarshal)" \
  -benchmem \
  -benchtime="${ITERATIONS}x" \
  -timeout="${TIMEOUT}" \
  -run=^$ \
  . 2>&1 || true)

echo "Benchmark completed, processing results..." >&2

# Save raw output
raw_file="${OUT_DIR}/latest_raw.txt"
echo "${benchmark_output}" > "${raw_file}"

echo "Parsing benchmark results..." >&2

# Process results with dedicated Python script
python3 "${ROOT_DIR}/scripts/parse_benchmarks.py" "${raw_file}" "${OUT_DIR}"

if [ $? -ne 0 ]; then
    echo "❌ Failed to parse benchmark results" >&2
    echo "Raw output saved to: ${raw_file}" >&2
    exit 1
fi

# Done! Exit before the old inline Python code
exit 0

# OLD CODE BELOW - KEPT FOR REFERENCE BUT NOT EXECUTED
python3 - "${raw_file}" "${OUT_DIR}" <<'PYTHON'
import sys
import re
import json
import platform
import subprocess
from datetime import datetime
from pathlib import Path
from collections import defaultdict

raw_file = Path(sys.argv[1])
out_dir = Path(sys.argv[2])

# Read benchmark output
with open(raw_file, 'r') as f:
    output = f.read()

# Parse benchmarks - More flexible pattern
pattern = r'Benchmark(\w+?)_(BEVE|JSON|Sonic|CBOR|MessagePack)_(Marshal(?:ZeroCopy)?|Unmarshal)[-/](\d+)\s+(\d+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+)\s+allocs/op'
results = []

for match in re.finditer(pattern, output):
    scenario, codec, operation, procs, iters, nsop, bop, allocsop = match.groups()
    
    # Handle ZeroCopy special case
    if 'ZeroCopy' in operation:
        codec = f"{codec} ZeroCopy"
        operation = operation.replace('ZeroCopy', '')
    
    # Format scenario name
    scenario_map = {
        'SmallStruct': 'Small Struct',
        'MediumPayload': 'Medium Payload', 
        'Medium': 'Medium Payload',
        'LargePayload': 'Large Payload',
        'Large': 'Large Payload',
        'Small': 'Small Struct'  # Some tests use 'Small' directly
    }
    scenario_formatted = scenario_map.get(scenario, scenario)
    
    results.append({
        'scenario': scenario_formatted,
        'codec': codec,
        'operation': operation,
        'nsop': float(nsop),
        'bop': int(bop),
        'allocsop': int(allocsop)
    })

if not results:
    print("⚠️  No benchmark results found!", file=sys.stderr)
    sys.exit(1)

# Get system info
cpu_brand = subprocess.getoutput("sysctl -n machdep.cpu.brand_string 2>/dev/null || lscpu | awk -F: '/Model name/ {gsub(/^ +| +$/, \"\", $2); print $2; exit}' 2>/dev/null || echo 'unknown'")
os_name = platform.system()
arch = platform.machine()
go_version = subprocess.getoutput("go version")
timestamp = datetime.now().astimezone().strftime('%Y-%m-%dT%H:%M:%SZ')

# Create platform-specific directory
platform_name = f"{os_name.lower()}-{arch.replace(' ', '-')}"
platform_dir = out_dir / f"benchmark-{platform_name}"
platform_dir.mkdir(parents=True, exist_ok=True)

# Generate Markdown report
md_lines = [
    f"# {cpu_brand} — {os_name}",
    "",
    f"![Benchmark Chart](benchmark-{platform_name}/benchmark.png)",
    "",
    "_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._",
    "",
    "### Detailed Results",
    "",
    "| Scenario | Codec | Operation | ns/op | B/op | allocs/op |",
    "|----------|-------|-----------|-------|------|-----------|",
]

# Sort by scenario, operation, then nsop
results_sorted = sorted(results, key=lambda x: (x['scenario'], x['operation'], x['nsop']))

for r in results_sorted:
    md_lines.append(f"| {r['scenario']} | {r['codec']} | {r['operation']} | {r['nsop']:.0f} | {r['bop']} | {r['allocsop']} |")

platform_md = platform_dir / "benchmark.md"
with open(platform_md, 'w') as f:
    f.write('\n'.join(md_lines))

# Generate JSON
json_data = {
    'timestamp': timestamp,
    'cpu': cpu_brand,
    'os': os_name,
    'architecture': arch,
    'go_version': go_version,
    'results': results
}

platform_json = platform_dir / "benchmark.json"
with open(platform_json, 'w') as f:
    json.dump(json_data, f, indent=2)

# Generate PNG chart
try:
    import matplotlib
    matplotlib.use('Agg')
    import matplotlib.pyplot as plt
    import numpy as np
    
    # Group by scenario
    scenarios = {}
    for r in results:
        if r['scenario'] not in scenarios:
            scenarios[r['scenario']] = []
        scenarios[r['scenario']].append(r)
    
    fig, axes = plt.subplots(1, 3, figsize=(18, 6), dpi=100)
    fig.suptitle(f'BEVE Performance Benchmark\n{cpu_brand} — {os_name}', 
                 fontsize=14, fontweight='bold')
    
    colors = {
        'BEVE': '#00D084',
        'BEVE ZeroCopy': '#00A86B', 
        'JSON': '#FF6B6B',
        'Sonic': '#4ECDC4',
        'CBOR': '#95E1D3',
        'MessagePack': '#F38181'
    }
    
    scenario_names = list(scenarios.keys())
    
    # Chart 1: Marshal performance (ns/op) - Log scale
    ax = axes[0]
    x = np.arange(len(scenario_names))
    width = 0.15
    
    codecs_marshal = sorted(set(r['codec'] for s in scenarios.values() for r in s if r['operation'] == 'Marshal'))
    for i, codec in enumerate(codecs_marshal):
        values = []
        for scenario_name in scenario_names:
            scenario_results = [r for r in scenarios[scenario_name] if r['codec'] == codec and r['operation'] == 'Marshal']
            values.append(scenario_results[0]['nsop'] if scenario_results else 0)
        
        ax.bar(x + i * width, values, width, label=codec, color=colors.get(codec, '#999'), alpha=0.85)
    
    ax.set_yscale('log')
    ax.set_ylabel('Time (ns/op)', fontweight='bold')
    ax.set_title('Marshal Performance (Log Scale)', fontweight='bold')
    ax.set_xticks(x + width * len(codecs_marshal) / 2)
    ax.set_xticklabels(scenario_names, fontsize=9)
    ax.legend(fontsize=8, loc='upper left')
    ax.grid(True, alpha=0.3, linestyle='--', axis='y')
    
    # Chart 2: Unmarshal performance (ns/op) - Log scale
    ax = axes[1]
    codecs_unmarshal = sorted(set(r['codec'] for s in scenarios.values() for r in s if r['operation'] == 'Unmarshal'))
    for i, codec in enumerate(codecs_unmarshal):
        values = []
        for scenario_name in scenario_names:
            scenario_results = [r for r in scenarios[scenario_name] if r['codec'] == codec and r['operation'] == 'Unmarshal']
            values.append(scenario_results[0]['nsop'] if scenario_results else 0)
        
        ax.bar(x + i * width, values, width, label=codec, color=colors.get(codec, '#999'), alpha=0.85)
    
    ax.set_yscale('log')
    ax.set_ylabel('Time (ns/op)', fontweight='bold')
    ax.set_title('Unmarshal Performance (Log Scale)', fontweight='bold')
    ax.set_xticks(x + width * len(codecs_unmarshal) / 2)
    ax.set_xticklabels(scenario_names, fontsize=9)
    ax.legend(fontsize=8, loc='upper left')
    ax.grid(True, alpha=0.3, linestyle='--', axis='y')
    
    # Chart 3: Memory allocation (B/op)
    ax = axes[2]
    all_codecs = sorted(set(r['codec'] for s in scenarios.values() for r in s))
    for i, codec in enumerate(all_codecs):
        values = []
        for scenario_name in scenario_names:
            scenario_results = [r for r in scenarios[scenario_name] if r['codec'] == codec]
            avg_bop = np.mean([r['bop'] for r in scenario_results]) if scenario_results else 0
            values.append(avg_bop)
        
        ax.bar(x + i * width, values, width, label=codec, color=colors.get(codec, '#999'), alpha=0.85)
    
    ax.set_ylabel('Memory (B/op)', fontweight='bold')
    ax.set_title('Average Memory Allocation', fontweight='bold')
    ax.set_xticks(x + width * len(all_codecs) / 2)
    ax.set_xticklabels(scenario_names, fontsize=9)
    ax.legend(fontsize=8, loc='upper left')
    ax.grid(True, alpha=0.3, linestyle='--', axis='y')
    
    plt.tight_layout()
    
    platform_png = platform_dir / "benchmark.png"
    plt.savefig(platform_png, bbox_inches='tight', dpi=100)
    print(f"✅ Chart saved to {platform_png}", file=sys.stderr)
    
except ImportError:
    print("⚠️  matplotlib not installed, skipping chart generation", file=sys.stderr)

# Update MULTI_PLATFORM.md
multi_platform_md = out_dir / "MULTI_PLATFORM.md"
md_header = [
    "# Multi-Platform Benchmark Results",
    "",
    "| CPU | OS | Artifacts |",
    "|-----|----|-----------|",
]

# Scan all platform directories
platform_rows = []
for p_dir in sorted(out_dir.glob("benchmark-*")):
    if p_dir.is_dir():
        p_name = p_dir.name.replace('benchmark-', '')
        p_md = p_dir / "benchmark.md"
        if p_md.exists():
            with open(p_md, 'r') as f:
                first_line = f.readline().strip()
                cpu_info = first_line.replace('# ', '').replace(' —', ',').split(',')[0].strip()
                os_info = first_line.split('—')[-1].strip() if '—' in first_line else 'unknown'
            
            platform_rows.append(f"| {cpu_info} | {os_info} | [Markdown]({p_name}/benchmark.md) · [JSON]({p_name}/benchmark.json) · [PNG]({p_name}/benchmark.png) |")

# Add current platform's detailed report
with open(multi_platform_md, 'w') as f:
    f.write('\n'.join(md_header + platform_rows + [''] + md_lines))

print(f"✅ Benchmarks complete!", file=sys.stderr)
print(f"📝 Markdown: {platform_md}", file=sys.stderr)
print(f"📊 JSON: {platform_json}", file=sys.stderr)
print(f"📄 Multi-platform: {multi_platform_md}", file=sys.stderr)
PYTHON

echo "" >&2
echo "✅ Done!" >&2
