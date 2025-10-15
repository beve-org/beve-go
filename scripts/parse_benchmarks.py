#!/usr/bin/env python3
"""
BEVE Benchmark Parser and Reporter
Parses Go benchmark output and generates JSON, MD, and PNG reports
"""

import sys
import re
import json
import platform
import subprocess
from datetime import datetime
from pathlib import Path

def parse_benchmarks(output_text):
    """Parse Go benchmark output"""
    # Pattern matches various Go benchmark formats
    patterns = [
        # Standard format: BenchmarkName-N iterations ns/op B/op allocs/op
        r'Benchmark(\w+?)_(BEVE|JSON|Sonic|CBOR|MessagePack)_(Marshal(?:ZeroCopy)?|Unmarshal)[-/](\d+)\s+(\d+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+)\s+allocs/op',
        # Alternative format without decimal points
        r'Benchmark(\w+?)_(BEVE|JSON|Sonic|CBOR|MessagePack)_(Marshal(?:ZeroCopy)?|Unmarshal)[-/](\d+)\s+(\d+)\s+(\d+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op',
    ]
    
    results = []
    seen = set()  # Avoid duplicates
    
    for pattern in patterns:
        for match in re.finditer(pattern, output_text):
            groups = match.groups()
            scenario, codec, operation = groups[0], groups[1], groups[2]
            nsop, bop, allocsop = float(groups[5]), float(groups[6]), float(groups[7])
            
            # Create unique key
            key = (scenario, codec, operation)
            if key in seen:
                continue
            seen.add(key)
            
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
                'Small': 'Small Struct'
            }
            scenario_formatted = scenario_map.get(scenario, scenario)
            
            results.append({
                'scenario': scenario_formatted,
                'codec': codec,
                'operation': operation,
                'ns_per_op': nsop,
                'bytes_per_op': int(bop),
                'allocs_per_op': int(allocsop)
            })
    
    return results


def get_system_info():
    """Get system information"""
    try:
        cpu_brand = subprocess.check_output(
            "sysctl -n machdep.cpu.brand_string 2>/dev/null || "
            "lscpu | awk -F: '/Model name/ {gsub(/^ +| +$/, \"\", $2); print $2; exit}' 2>/dev/null || "
            "echo 'Unknown CPU'",
            shell=True, text=True
        ).strip()
    except:
        cpu_brand = "Unknown CPU"
    
    os_name = platform.system()
    arch = platform.machine()
    
    try:
        go_version = subprocess.check_output(["go", "version"], text=True).strip()
    except:
        go_version = "Unknown Go version"
    
    return {
        'cpu': cpu_brand,
        'os': os_name,
        'arch': arch,
        'go_version': go_version,
        'timestamp': datetime.now().astimezone().strftime('%Y-%m-%dT%H:%M:%SZ')
    }


def generate_markdown(results, system_info, output_path):
    """Generate Markdown report"""
    lines = [
        f"# {system_info['cpu']} — {system_info['os']}",
        "",
        "_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._",
        "",
        "### Detailed Results",
        "",
        "| Scenario | Codec | Operation | ns/op | B/op | allocs/op |",
        "|----------|-------|-----------|-------|------|-----------|",
    ]
    
    # Sort by scenario, operation, then ns/op
    results_sorted = sorted(results, key=lambda x: (x['scenario'], x['operation'], x['ns_per_op']))
    
    for r in results_sorted:
        lines.append(
            f"| {r['scenario']} | {r['codec']} | {r['operation']} | "
            f"{r['ns_per_op']:.0f} | {r['bytes_per_op']} | {r['allocs_per_op']} |"
        )
    
    output_path.write_text('\n'.join(lines) + '\n', encoding='utf-8')
    print(f"✅ Generated Markdown: {output_path}")


def generate_json(results, system_info, output_path):
    """Generate JSON report"""
    data = {
        'timestamp': system_info['timestamp'],
        'environment': {
            'cpu': system_info['cpu'],
            'os': system_info['os'],
            'arch': system_info['arch'],
            'go_version': system_info['go_version']
        },
        'results': results
    }
    
    output_path.write_text(json.dumps(data, indent=2) + '\n', encoding='utf-8')
    print(f"✅ Generated JSON: {output_path}")


def generate_chart(results, system_info, output_path):
    """Generate PNG chart"""
    try:
        import matplotlib
        matplotlib.use('Agg')
        import matplotlib.pyplot as plt
        import numpy as np
    except ImportError:
        print("⚠️  matplotlib not available, skipping chart generation", file=sys.stderr)
        return
    
    # Group by scenario
    scenarios = {}
    for r in results:
        if r['scenario'] not in scenarios:
            scenarios[r['scenario']] = []
        scenarios[r['scenario']].append(r)
    
    if not scenarios:
        print("⚠️  No scenarios to chart", file=sys.stderr)
        return
    
    # Create figure
    fig, axes = plt.subplots(1, 2, figsize=(14, 6), dpi=100)
    fig.suptitle(f"BEVE Performance Benchmark\n{system_info['cpu']} — {system_info['os']}", 
                 fontsize=14, fontweight='bold')
    
    colors = {
        'BEVE': '#00D084',
        'BEVE ZeroCopy': '#00A86B',
        'JSON': '#FF6B6B',
        'Sonic': '#4ECDC4',
        'CBOR': '#95E1D3',
        'MessagePack': '#F38181'
    }
    
    scenario_names = sorted(scenarios.keys())
    x = np.arange(len(scenario_names))
    width = 0.15
    
    # Chart 1: Marshal performance
    ax = axes[0]
    codecs_marshal = sorted(set(r['codec'] for s in scenarios.values() 
                               for r in s if r['operation'] == 'Marshal'))
    
    for i, codec in enumerate(codecs_marshal):
        values = []
        for scenario_name in scenario_names:
            scenario_results = [r for r in scenarios[scenario_name] 
                              if r['codec'] == codec and r['operation'] == 'Marshal']
            values.append(scenario_results[0]['ns_per_op'] if scenario_results else 0)
        
        if any(v > 0 for v in values):
            ax.bar(x + i * width, values, width, label=codec, 
                  color=colors.get(codec, '#999'), alpha=0.85)
    
    ax.set_yscale('log')
    ax.set_ylabel('Time (ns/op)', fontweight='bold')
    ax.set_title('Marshal Performance (Log Scale)', fontweight='bold')
    ax.set_xticks(x + width * len(codecs_marshal) / 2)
    ax.set_xticklabels(scenario_names, fontsize=9)
    ax.legend(fontsize=8, loc='upper left')
    ax.grid(True, alpha=0.3, linestyle='--', axis='y')
    
    # Chart 2: Unmarshal performance
    ax = axes[1]
    codecs_unmarshal = sorted(set(r['codec'] for s in scenarios.values() 
                                 for r in s if r['operation'] == 'Unmarshal'))
    
    for i, codec in enumerate(codecs_unmarshal):
        values = []
        for scenario_name in scenario_names:
            scenario_results = [r for r in scenarios[scenario_name] 
                              if r['codec'] == codec and r['operation'] == 'Unmarshal']
            values.append(scenario_results[0]['ns_per_op'] if scenario_results else 0)
        
        if any(v > 0 for v in values):
            ax.bar(x + i * width, values, width, label=codec,
                  color=colors.get(codec, '#999'), alpha=0.85)
    
    ax.set_yscale('log')
    ax.set_ylabel('Time (ns/op)', fontweight='bold')
    ax.set_title('Unmarshal Performance (Log Scale)', fontweight='bold')
    ax.set_xticks(x + width * len(codecs_unmarshal) / 2)
    ax.set_xticklabels(scenario_names, fontsize=9)
    ax.legend(fontsize=8, loc='upper left')
    ax.grid(True, alpha=0.3, linestyle='--', axis='y')
    
    plt.tight_layout()
    plt.savefig(output_path, dpi=150, bbox_inches='tight', facecolor='white')
    plt.close()
    
    print(f"✅ Generated Chart: {output_path}")


def main():
    if len(sys.argv) < 3:
        print("Usage: python parse_benchmarks.py <raw_file> <output_dir>", file=sys.stderr)
        sys.exit(1)
    
    raw_file = Path(sys.argv[1])
    out_dir = Path(sys.argv[2])
    
    if not raw_file.exists():
        print(f"❌ Error: Raw file not found: {raw_file}", file=sys.stderr)
        sys.exit(1)
    
    # Read benchmark output
    output_text = raw_file.read_text(encoding='utf-8')
    
    # Parse benchmarks
    results = parse_benchmarks(output_text)
    
    if not results:
        print("❌ Error: No benchmark results found in output!", file=sys.stderr)
        print("\nSearching for benchmark lines in output:", file=sys.stderr)
        for line in output_text.split('\n'):
            if 'Benchmark' in line and 'ns/op' in line:
                print(f"  Found: {line[:100]}", file=sys.stderr)
        sys.exit(1)
    
    print(f"✅ Parsed {len(results)} benchmark results")
    
    # Get system info
    system_info = get_system_info()
    
    # Generate outputs
    generate_json(results, system_info, out_dir / "latest.json")
    generate_markdown(results, system_info, out_dir / "latest.md")
    generate_chart(results, system_info, out_dir / "latest.png")
    
    print(f"\n🎉 Benchmark reports generated successfully!")
    print(f"📁 Output directory: {out_dir}")


if __name__ == "__main__":
    main()
