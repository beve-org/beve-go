#!/usr/bin/env python3
"""
BEVE-Go Multi-Platform Benchmark Chart Generator
Creates a unified visualization comparing all platforms
"""

import json
import sys
from pathlib import Path
from typing import Dict, List, Any

try:
    import matplotlib
    matplotlib.use('Agg')  # Non-interactive backend
    import matplotlib.pyplot as plt
    import numpy as np
except ImportError:
    print("Error: matplotlib required. Install with: pip install matplotlib", file=sys.stderr)
    sys.exit(1)


def format_time(ns: float) -> str:
    """Format nanoseconds to readable string"""
    if ns >= 1_000_000:
        return f"{ns/1_000_000:.2f}ms"
    elif ns >= 1_000:
        return f"{ns/1_000:.2f}μs"
    else:
        return f"{ns:.0f}ns"


def create_multi_platform_chart(platforms: List[Dict[str, Any]], output_path: Path):
    """Create a comprehensive multi-platform comparison chart"""
    
    # Extract data for comparison
    platform_names = []
    marshal_data = {'BEVE': [], 'BEVE ZeroCopy': [], 'JSON': [], 'CBOR': [], 'MessagePack': []}
    unmarshal_data = {'BEVE': [], 'JSON': [], 'CBOR': [], 'MessagePack': []}
    
    for platform in platforms:
        platform_names.append(platform['cpu_name'])
        results = platform['results']
        
        # Extract marshal times
        for codec in marshal_data.keys():
            found = False
            for entry in results:
                if (entry.get('scenario') == 'Small Struct' and 
                    entry.get('operation') == 'Marshal' and 
                    entry.get('codec') == codec):
                    marshal_data[codec].append(entry.get('ns_per_op', 0) / 1000)  # Convert to μs
                    found = True
                    break
            if not found:
                marshal_data[codec].append(0)
        
        # Extract unmarshal times
        for codec in unmarshal_data.keys():
            found = False
            for entry in results:
                if (entry.get('scenario') == 'Small Struct' and 
                    entry.get('operation') == 'Unmarshal' and 
                    entry.get('codec') == codec):
                    unmarshal_data[codec].append(entry.get('ns_per_op', 0) / 1000)  # Convert to μs
                    found = True
                    break
            if not found:
                unmarshal_data[codec].append(0)
    
    # Create figure with subplots
    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(16, 8))
    fig.suptitle('BEVE-Go Multi-Platform Performance Comparison', fontsize=16, fontweight='bold')
    
    x = np.arange(len(platform_names))
    width = 0.15
    
    # Colors for each codec
    colors = {
        'BEVE ZeroCopy': '#2ecc71',  # Green
        'BEVE': '#3498db',           # Blue
        'JSON': '#e74c3c',           # Red
        'CBOR': '#f39c12',           # Orange
        'MessagePack': '#9b59b6'     # Purple
    }
    
    # Plot Marshal performance
    offset = -2 * width
    for codec, times in marshal_data.items():
        if any(t > 0 for t in times):
            bars = ax1.bar(x + offset, times, width, label=codec, color=colors.get(codec, '#95a5a6'))
            # Add value labels on bars
            for bar in bars:
                height = bar.get_height()
                if height > 0:
                    ax1.text(bar.get_x() + bar.get_width()/2., height,
                            f'{height:.2f}',
                            ha='center', va='bottom', fontsize=8)
            offset += width
    
    ax1.set_xlabel('Platform', fontsize=12, fontweight='bold')
    ax1.set_ylabel('Time (μs)', fontsize=12, fontweight='bold')
    ax1.set_title('Marshal Performance (Small Struct)\nLower is Better', fontsize=14)
    ax1.set_xticks(x)
    ax1.set_xticklabels(platform_names, rotation=15, ha='right')
    ax1.legend(loc='upper left', framealpha=0.9)
    ax1.grid(axis='y', alpha=0.3, linestyle='--')
    ax1.set_ylim(bottom=0)
    
    # Plot Unmarshal performance
    offset = -1.5 * width
    for codec, times in unmarshal_data.items():
        if any(t > 0 for t in times):
            bars = ax2.bar(x + offset, times, width, label=codec, color=colors.get(codec, '#95a5a6'))
            # Add value labels on bars
            for bar in bars:
                height = bar.get_height()
                if height > 0:
                    ax2.text(bar.get_x() + bar.get_width()/2., height,
                            f'{height:.2f}',
                            ha='center', va='bottom', fontsize=8)
            offset += width
    
    ax2.set_xlabel('Platform', fontsize=12, fontweight='bold')
    ax2.set_ylabel('Time (μs)', fontsize=12, fontweight='bold')
    ax2.set_title('Unmarshal Performance (Small Struct)\nLower is Better', fontsize=14)
    ax2.set_xticks(x)
    ax2.set_xticklabels(platform_names, rotation=15, ha='right')
    ax2.legend(loc='upper left', framealpha=0.9)
    ax2.grid(axis='y', alpha=0.3, linestyle='--')
    ax2.set_ylim(bottom=0)
    
    # Add footer with generation info
    fig.text(0.5, 0.02, 'Generated by BEVE-Go CI/CD Pipeline', 
             ha='center', fontsize=10, style='italic', color='gray')
    
    plt.tight_layout()
    plt.savefig(output_path, dpi=150, bbox_inches='tight', facecolor='white')
    plt.close()
    
    print(f"✅ Generated multi-platform chart: {output_path}")


def create_performance_heatmap(platforms: List[Dict[str, Any]], output_path: Path):
    """Create a heatmap showing relative performance across platforms"""
    
    fig, ax = plt.subplots(figsize=(14, 8))
    fig.suptitle('BEVE Performance vs Competitors (Relative Speed)', fontsize=16, fontweight='bold')
    
    platform_names = [p['cpu_name'] for p in platforms]
    codecs = ['JSON', 'CBOR', 'MessagePack', 'Sonic']
    operations = ['Marshal', 'Unmarshal']
    
    # Create data matrix (speedup factor: BEVE / competitor)
    data = []
    labels = []
    
    for operation in operations:
        for codec in codecs:
            row = []
            for platform in platforms:
                beve_time = None
                codec_time = None
                
                for entry in platform['results']:
                    if entry.get('scenario') == 'Small Struct' and entry.get('operation') == operation:
                        if entry.get('codec') == 'BEVE':
                            beve_time = entry.get('ns_per_op', 0)
                        elif entry.get('codec') == codec:
                            codec_time = entry.get('ns_per_op', 0)
                
                if beve_time and codec_time and beve_time > 0:
                    speedup = codec_time / beve_time
                    row.append(speedup)
                else:
                    row.append(0)
            
            if any(v > 0 for v in row):
                data.append(row)
                labels.append(f'{operation}\n{codec}')
    
    if not data:
        print("⚠️  No data for heatmap")
        return
    
    data_array = np.array(data)
    
    im = ax.imshow(data_array, cmap='RdYlGn', aspect='auto', vmin=0, vmax=max(3, data_array.max()))
    
    # Set ticks
    ax.set_xticks(np.arange(len(platform_names)))
    ax.set_yticks(np.arange(len(labels)))
    ax.set_xticklabels(platform_names, rotation=15, ha='right')
    ax.set_yticklabels(labels)
    
    # Add colorbar
    cbar = plt.colorbar(im, ax=ax)
    cbar.set_label('Speedup Factor (×)', rotation=270, labelpad=20)
    
    # Add text annotations
    for i in range(len(labels)):
        for j in range(len(platform_names)):
            value = data_array[i, j]
            if value > 0:
                text = ax.text(j, i, f'{value:.1f}×',
                             ha="center", va="center", color="black", fontweight='bold')
    
    ax.set_xlabel('Platform', fontsize=12, fontweight='bold')
    ax.set_ylabel('Operation / Codec', fontsize=12, fontweight='bold')
    ax.set_title('How much faster is BEVE?', fontsize=14, pad=20)
    
    plt.tight_layout()
    plt.savefig(output_path, dpi=150, bbox_inches='tight', facecolor='white')
    plt.close()
    
    print(f"✅ Generated performance heatmap: {output_path}")


def create_memory_comparison(platforms: List[Dict[str, Any]], output_path: Path):
    """Create a chart comparing memory allocations"""
    
    fig, ax = plt.subplots(figsize=(14, 6))
    fig.suptitle('Memory Allocations per Operation', fontsize=16, fontweight='bold')
    
    platform_names = [p['cpu_name'] for p in platforms]
    codecs = ['BEVE', 'JSON', 'CBOR', 'MessagePack']
    
    x = np.arange(len(platform_names))
    width = 0.2
    colors = {'BEVE': '#3498db', 'JSON': '#e74c3c', 'CBOR': '#f39c12', 'MessagePack': '#9b59b6'}
    
    offset = -1.5 * width
    for codec in codecs:
        allocs = []
        for platform in platforms:
            found = False
            for entry in platform['results']:
                if (entry.get('scenario') == 'Small Struct' and 
                    entry.get('codec') == codec):
                    allocs.append(entry.get('allocs_per_op', 0))
                    found = True
                    break
            if not found:
                allocs.append(0)
        
        if any(a > 0 for a in allocs):
            bars = ax.bar(x + offset, allocs, width, label=codec, color=colors.get(codec, '#95a5a6'))
            for bar in bars:
                height = bar.get_height()
                if height > 0:
                    ax.text(bar.get_x() + bar.get_width()/2., height,
                           f'{int(height)}',
                           ha='center', va='bottom', fontsize=9)
            offset += width
    
    ax.set_xlabel('Platform', fontsize=12, fontweight='bold')
    ax.set_ylabel('Allocations per Operation', fontsize=12, fontweight='bold')
    ax.set_title('Lower is Better', fontsize=14)
    ax.set_xticks(x)
    ax.set_xticklabels(platform_names, rotation=15, ha='right')
    ax.legend(loc='upper left', framealpha=0.9)
    ax.grid(axis='y', alpha=0.3, linestyle='--')
    ax.set_ylim(bottom=0)
    
    plt.tight_layout()
    plt.savefig(output_path, dpi=150, bbox_inches='tight', facecolor='white')
    plt.close()
    
    print(f"✅ Generated memory comparison: {output_path}")


def main():
    if len(sys.argv) < 2:
        print("Usage: python plot_multi_platform.py <output_dir>", file=sys.stderr)
        sys.exit(1)
    
    output_dir = Path(sys.argv[1])
    output_dir.mkdir(parents=True, exist_ok=True)
    
    artifacts_root = Path("artifacts")
    
    if not artifacts_root.exists():
        print(f"❌ Error: Artifacts directory not found", file=sys.stderr)
        sys.exit(1)
    
    # Load all platform data
    platforms = []
    benchmark_jsons = sorted(artifacts_root.glob("**/benchmark.json"))
    
    if not benchmark_jsons:
        print("❌ Error: No benchmark.json files found", file=sys.stderr)
        sys.exit(1)
    
    for json_path in benchmark_jsons:
        try:
            data = json.loads(json_path.read_text(encoding="utf-8"))
            env = data.get("environment", {})
            
            platforms.append({
                "cpu_name": env.get("cpu", "Unknown CPU"),
                "os_name": env.get("os", "Unknown OS"),
                "arch": env.get("arch", "unknown"),
                "results": data.get("results", [])
            })
        except Exception as e:
            print(f"⚠️  Warning: Failed to load {json_path}: {e}", file=sys.stderr)
    
    if not platforms:
        print("❌ Error: No valid platform data found", file=sys.stderr)
        sys.exit(1)
    
    print(f"📊 Generating charts for {len(platforms)} platforms...")
    
    # Generate all charts
    create_multi_platform_chart(platforms, output_dir / "multi_platform_comparison.png")
    create_performance_heatmap(platforms, output_dir / "performance_heatmap.png")
    create_memory_comparison(platforms, output_dir / "memory_comparison.png")
    
    print("✅ All charts generated successfully!")


if __name__ == "__main__":
    main()
