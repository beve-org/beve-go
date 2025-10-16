#!/usr/bin/env python3
"""
BEVE-Go Multi-Platform Benchmark Aggregator
Combines benchmark results from multiple CI/CD runs into a unified report
"""

import json
import shutil
import sys
from pathlib import Path
from typing import Dict, List, Any
from datetime import datetime


def format_value(value: Any) -> str:
    """Format numeric values for display"""
    if isinstance(value, float):
        if value >= 1_000_000:
            return f"{value/1_000_000:.2f}M"
        elif value >= 1_000:
            return f"{value/1_000:.2f}K"
        return f"{value:.2f}"
    if isinstance(value, int):
        if value >= 1_000_000:
            return f"{value/1_000_000:.2f}M"
        elif value >= 1_000:
            return f"{value/1_000:.1f}K"
        return str(value)
    if value in (None, ""):
        return "n/a"
    return str(value)


def format_ns_time(ns: float) -> str:
    """Format nanoseconds to human readable time"""
    if ns >= 1_000_000_000:
        return f"{ns/1_000_000_000:.2f}s"
    elif ns >= 1_000_000:
        return f"{ns/1_000_000:.2f}ms"
    elif ns >= 1_000:
        return f"{ns/1_000:.2f}μs"
    else:
        return f"{ns:.0f}ns"


def get_performance_emoji(codec: str) -> str:
    """Get emoji based on codec performance tier"""
    if codec in ["BEVE", "BEVE ZeroCopy"]:
        return "🥇"
    elif codec in ["CBOR", "MessagePack"]:
        return "🥈"
    elif codec in ["JSON", "Sonic"]:
        return "🥉"
    return "📊"


def create_comparison_table(all_platforms: List[Dict[str, Any]]) -> List[str]:
    """Create a cross-platform comparison table"""
    lines = [
        "## 📊 Cross-Platform Performance Comparison",
        "",
        "### Marshal Performance (Small Struct)",
        "",
        "| Platform | BEVE | BEVE ZeroCopy | JSON | CBOR | MessagePack |",
        "|----------|------|---------------|------|------|-------------|"
    ]
    
    for platform_data in all_platforms:
        cpu = platform_data["cpu_name"]
        results = platform_data["results"]
        
        # Extract small struct marshal results
        marshal_results = {}
        for entry in results:
            if entry.get("scenario") == "Small Struct" and entry.get("operation") == "Marshal":
                codec = entry.get("codec", "")
                ns = entry.get("ns_per_op")
                if codec and ns:
                    marshal_results[codec] = format_ns_time(ns)
        
        row = f"| {cpu} | " + " | ".join([
            marshal_results.get("BEVE", "n/a"),
            marshal_results.get("BEVE ZeroCopy", "n/a"),
            marshal_results.get("JSON", "n/a"),
            marshal_results.get("CBOR", "n/a"),
            marshal_results.get("MessagePack", "n/a")
        ]) + " |"
        
        lines.append(row)
    
    lines.extend([
        "",
        "### Unmarshal Performance (Small Struct)",
        "",
        "| Platform | BEVE | JSON | CBOR | MessagePack |",
        "|----------|------|------|------|-------------|"
    ])
    
    for platform_data in all_platforms:
        cpu = platform_data["cpu_name"]
        results = platform_data["results"]
        
        # Extract small struct unmarshal results
        unmarshal_results = {}
        for entry in results:
            if entry.get("scenario") == "Small Struct" and entry.get("operation") == "Unmarshal":
                codec = entry.get("codec", "")
                ns = entry.get("ns_per_op")
                if codec and ns:
                    unmarshal_results[codec] = format_ns_time(ns)
        
        row = f"| {cpu} | " + " | ".join([
            unmarshal_results.get("BEVE", "n/a"),
            unmarshal_results.get("JSON", "n/a"),
            unmarshal_results.get("CBOR", "n/a"),
            unmarshal_results.get("MessagePack", "n/a")
        ]) + " |"
        
        lines.append(row)
    
    lines.append("")
    return lines


def create_winners_table(all_platforms: List[Dict[str, Any]]) -> List[str]:
    """Create a table showing which codec wins on each platform"""
    lines = [
        "## 🏆 Performance Champions",
        "",
        "| Platform | Fastest Marshal | Fastest Unmarshal | Memory Efficient |",
        "|----------|----------------|-------------------|------------------|"
    ]
    
    for platform_data in all_platforms:
        cpu = platform_data["cpu_name"]
        results = platform_data["results"]
        
        # Find fastest marshal
        marshal_times = []
        for entry in results:
            if entry.get("scenario") == "Small Struct" and entry.get("operation") == "Marshal":
                codec = entry.get("codec", "")
                ns = entry.get("ns_per_op")
                if codec and ns and ns != "n/a":
                    marshal_times.append((codec, ns))
        
        fastest_marshal = min(marshal_times, key=lambda x: x[1]) if marshal_times else ("n/a", 0)
        
        # Find fastest unmarshal
        unmarshal_times = []
        for entry in results:
            if entry.get("scenario") == "Small Struct" and entry.get("operation") == "Unmarshal":
                codec = entry.get("codec", "")
                ns = entry.get("ns_per_op")
                if codec and ns and ns != "n/a":
                    unmarshal_times.append((codec, ns))
        
        fastest_unmarshal = min(unmarshal_times, key=lambda x: x[1]) if unmarshal_times else ("n/a", 0)
        
        # Find most memory efficient
        alloc_counts = []
        for entry in results:
            if entry.get("scenario") == "Small Struct":
                codec = entry.get("codec", "")
                allocs = entry.get("allocs_per_op")
                if codec and allocs and allocs != "n/a":
                    alloc_counts.append((codec, allocs))
        
        most_efficient = min(alloc_counts, key=lambda x: x[1]) if alloc_counts else ("n/a", 0)
        
        marshal_str = f"{get_performance_emoji(fastest_marshal[0])} {fastest_marshal[0]} ({format_ns_time(fastest_marshal[1])})"
        unmarshal_str = f"{get_performance_emoji(fastest_unmarshal[0])} {fastest_unmarshal[0]} ({format_ns_time(fastest_unmarshal[1])})"
        efficient_str = f"💾 {most_efficient[0]} ({most_efficient[1]} allocs)"
        
        lines.append(f"| {cpu} | {marshal_str} | {unmarshal_str} | {efficient_str} |")
    
    lines.append("")
    return lines


def create_summary_stats(all_platforms: List[Dict[str, Any]]) -> List[str]:
    """Create summary statistics"""
    lines = [
        "## 📈 Summary Statistics",
        "",
        f"**Total Platforms Tested:** {len(all_platforms)}",
        ""
    ]
    
    # Calculate average improvements
    beve_improvements = []
    for platform_data in all_platforms:
        results = platform_data["results"]
        
        beve_time = None
        json_time = None
        
        for entry in results:
            if entry.get("scenario") == "Small Struct" and entry.get("operation") == "Marshal":
                codec = entry.get("codec", "")
                ns = entry.get("ns_per_op")
                if codec == "BEVE" and ns:
                    beve_time = ns
                elif codec == "JSON" and ns:
                    json_time = ns
        
        if beve_time and json_time:
            improvement = ((json_time - beve_time) / json_time) * 100
            beve_improvements.append(improvement)
    
    if beve_improvements:
        avg_improvement = sum(beve_improvements) / len(beve_improvements)
        lines.extend([
            f"**Average BEVE vs JSON Improvement:** {avg_improvement:.1f}% faster",
            ""
        ])
    
    lines.extend([
        "### Platform Details",
        ""
    ])
    
    for platform_data in all_platforms:
        lines.extend([
            f"- **{platform_data['cpu_name']}** ({platform_data['os_name']})",
            f"  - Architecture: {platform_data.get('arch', 'unknown')}",
            f"  - Test Scenarios: {len(set(r.get('scenario') for r in platform_data['results']))}",
            ""
        ])
    
    return lines


def main():
    artifacts_root = Path("artifacts")
    dist_root = Path("dist")
    bench_root = dist_root / "benchmarks"
    bench_root.mkdir(parents=True, exist_ok=True)

    if not artifacts_root.exists():
        print(f"❌ Error: Artifacts directory not found at {artifacts_root}", file=sys.stderr)
        sys.exit(1)

    benchmark_jsons = sorted(artifacts_root.glob("**/benchmark.json"))
    if not benchmark_jsons:
        print("❌ Error: No benchmark.json files found", file=sys.stderr)
        sys.exit(1)

    print(f"✅ Found {len(benchmark_jsons)} platform benchmark results")

    all_platforms = []
    summary_lines = [
        "# 🚀 BEVE-Go Multi-Platform Benchmark Results",
        "",
        f"**Generated:** {datetime.now().strftime('%Y-%m-%d %H:%M:%S UTC')}",
        "",
        "This report consolidates benchmark results from multiple platforms tested in CI/CD.",
        "",
        "## � Visual Comparisons",
        "",
        "### Overall Performance Comparison",
        "![Multi-Platform Comparison](charts/multi_platform_comparison.png)",
        "",
        "### Performance Heatmap",
        "![Performance Heatmap](charts/performance_heatmap.png)",
        "",
        "### Memory Efficiency",
        "![Memory Comparison](charts/memory_comparison.png)",
        "",
        "---",
        "",
        "## �🖥️ Tested Platforms",
        "",
        "| Platform | CPU | OS | Artifacts |",
        "|----------|-----|----|-----------| "
    ]

    # Process each platform
    for json_path in benchmark_jsons:
        cpu_dir = json_path.parent
        slug = cpu_dir.name
        data = json.loads(json_path.read_text(encoding="utf-8"))
        env = data.get("environment", {})
        cpu_name = env.get("cpu", slug)
        os_name = env.get("os", "unknown")
        arch = env.get("arch", "unknown")

        # Copy artifacts to dist
        dest_dir = bench_root / slug
        if dest_dir.exists():
            shutil.rmtree(dest_dir)
        dest_dir.mkdir(parents=True, exist_ok=True)

        for artifact_file in cpu_dir.glob("benchmark.*"):
            shutil.copy2(artifact_file, dest_dir / artifact_file.name)

        # Add to summary table
        summary_lines.append(
            f"| {slug} | {cpu_name} | {os_name} | "
            f"[📄 Report]({slug}/benchmark.md) · "
            f"[📊 JSON]({slug}/benchmark.json) · "
            f"[📈 Chart]({slug}/benchmark.png) |"
        )

        # Store platform data
        all_platforms.append({
            "slug": slug,
            "cpu_name": cpu_name,
            "os_name": os_name,
            "arch": arch,
            "results": data.get("results", [])
        })

    summary_lines.extend(["", "---", ""])

    # Add cross-platform comparison
    summary_lines.extend(create_comparison_table(all_platforms))
    summary_lines.extend(create_winners_table(all_platforms))
    summary_lines.extend(create_summary_stats(all_platforms))

    # Add detailed per-platform results
    summary_lines.extend(["---", "", "## 📋 Detailed Platform Results", ""])

    for platform_data in all_platforms:
        slug = platform_data["slug"]
        cpu_name = platform_data["cpu_name"]
        os_name = platform_data["os_name"]
        results = platform_data["results"]

        summary_lines.extend([
            f"### {cpu_name} — {os_name}",
            "",
            f"![Benchmark Chart]({slug}/benchmark.png)",
            "",
            "_Performance visualization: lower is better._",
            "",
            "| Scenario | Codec | Operation | Time | Memory | Allocations |",
            "|----------|-------|-----------|------|--------|-------------|"
        ])

        # Group and sort results
        grouped = {}
        order = []
        for entry in results:
            key = (entry.get("scenario", ""), entry.get("operation", ""))
            if key not in grouped:
                grouped[key] = []
                order.append(key)
            grouped[key].append(entry)

        for key in order:
            scenario, operation = key
            entries = sorted(
                grouped[key],
                key=lambda e: float("inf") if e.get("ns_per_op") in (None, "n/a") else e["ns_per_op"]
            )
            for entry in entries:
                emoji = get_performance_emoji(entry.get("codec", ""))
                summary_lines.append(
                    f"| {scenario} | {emoji} {entry.get('codec', '')} | {operation} | "
                    f"{format_ns_time(entry.get('ns_per_op', 0))} | "
                    f"{format_value(entry.get('bytes_per_op', 'n/a'))} | "
                    f"{format_value(entry.get('allocs_per_op', 'n/a'))} |"
                )

        summary_lines.extend(["", f"[📄 View full report]({slug}/benchmark.md)", ""])

    # Add footer
    summary_lines.extend([
        "---",
        "",
        "## 📚 Additional Resources",
        "",
        "- [BEVE Specification](../SPECIFICATION.md)",
        "- [Go Package Documentation](../README.md)",
        "- [Translator Package](../translator/README.md)",
        "- [Examples](../examples/)",
        "",
        "**Legend:**",
        "- 🥇 BEVE family (fastest)",
        "- 🥈 CBOR/MessagePack (fast)",
        "- 🥉 JSON/Sonic (standard)",
        ""
    ])

    # Write summary
    summary_path = dist_root / "MULTI_PLATFORM.md"
    summary_path.write_text("\n".join(summary_lines) + "\n", encoding="utf-8")
    print(f"✅ Generated multi-platform summary: {summary_path}")


if __name__ == "__main__":
    main()
