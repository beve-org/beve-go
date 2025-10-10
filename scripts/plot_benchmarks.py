#!/usr/bin/env python3
"""Render benchmark summaries as charts.

This script consumes the JSON output produced by scripts/bench.sh and
emits a PNG (or any matplotlib-supported image format) that compares the
ns/op metric across codecs for each benchmark scenario.
"""
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

import matplotlib

matplotlib.use("Agg")  # Use a non-interactive backend suitable for CI
import matplotlib.pyplot as plt


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Plot benchmark results from JSON")
    parser.add_argument("input", type=Path, help="Path to benchmarks/latest.json")
    parser.add_argument("output", type=Path, help="Path where the chart image will be written")
    parser.add_argument(
        "--metric",
        choices=("ns_per_op", "bytes_per_op", "allocs_per_op"),
        default="ns_per_op",
        help="Which metric to visualise (default: ns_per_op)",
    )
    parser.add_argument(
        "--unit",
        choices=("auto", "ns", "us", "ms"),
        default="auto",
        help="Unit to display for the chosen metric (default: auto)",
    )
    return parser.parse_args()


def load_results(path: Path) -> dict:
    with path.open("r", encoding="utf-8") as fh:
        return json.load(fh)


def select_unit(metric: str, unit_option: str, values: list[float]) -> tuple[str, float]:
    if metric != "ns_per_op":
        # bytes_per_op and allocs_per_op already have intuitive units
        return metric, 1.0

    if unit_option == "ns" or not values:
        return "ns/op", 1.0
    if unit_option == "us":
        return "µs/op", 1e-3
    if unit_option == "ms":
        return "ms/op", 1e-6

    # auto-select: choose unit that keeps numbers within a readable range
    max_value = max(values)
    if max_value >= 1e6:
        return "ms/op", 1e-6
    if max_value >= 1e3:
        return "µs/op", 1e-3
    return "ns/op", 1.0


def build_groups(results: list[dict], metric: str) -> list[tuple[tuple[str, str], list[dict]]]:
    order = []
    grouped: dict[tuple[str, str], list[dict]] = {}
    for entry in results:
        value = entry.get(metric)
        if value is None:
            continue
        try:
            numeric = float(value)
        except (TypeError, ValueError):
            continue

        scenario = entry.get("scenario", "Unknown")
        operation = entry.get("operation", "Unknown")
        key = (scenario, operation)
        if key not in grouped:
            grouped[key] = []
            order.append(key)
        grouped[key].append({"codec": entry.get("codec", "?"), metric: numeric})

    return [(key, grouped[key]) for key in order]


def pick_colors(codecs: list[str]) -> list[str]:
    palette = {
        "BEVE": "#1f77b4",
        "JSON": "#ff7f0e",
        "Sonic": "#2ca02c",
        "MessagePack": "#9467bd",
        "CBOR": "#d62728",
    }
    default_cycle = plt.rcParams["axes.prop_cycle"].by_key().get("color", [])
    colors = []
    for idx, codec in enumerate(codecs):
        if codec in palette:
            colors.append(palette[codec])
        elif default_cycle:
            colors.append(default_cycle[idx % len(default_cycle)])
        else:
            colors.append(f"C{idx}")
    return colors


def render_chart(data: dict, output_path: Path, metric: str, unit_option: str) -> None:
    results = data.get("results", [])
    groups = build_groups(results, metric)
    if not groups:
        raise SystemExit("no benchmark groups found in JSON input")

    # Collect all metric values for unit selection
    all_values = [item[metric] for _, entries in groups for item in entries]
    unit_label, scale = select_unit(metric, unit_option, all_values)

    figure_height = max(4.0, 2.5 * len(groups))
    fig, axes = plt.subplots(len(groups), 1, figsize=(12, figure_height), squeeze=False)
    axes = axes.flatten()

    for ax, ((scenario, operation), entries) in zip(axes, groups):
        entries = sorted(entries, key=lambda e: e[metric])
        codecs = [entry["codec"] for entry in entries]
        values = [entry[metric] * scale for entry in entries]
        colors = pick_colors(codecs)

        y_positions = list(range(len(entries)))
        bars = ax.barh(y_positions, values, color=colors)
        ax.set_yticks(y_positions)
        ax.set_yticklabels(codecs)
        ax.invert_yaxis()  # Best performer at the top
        ax.set_xlabel(unit_label)
        ax.set_title(f"{scenario} · {operation}")
        ax.grid(axis="x", linestyle="--", linewidth=0.5, alpha=0.5)

        ax.bar_label(bars, labels=[f"{value:.3g}" for value in values], padding=4, fontsize=8)

        max_value = max(values) if values else 0
        if max_value > 0:
            ax.set_xlim(0, max_value * 1.15)

    environment = data.get("environment", {})
    meta_lines = []
    if environment:
        os_info = environment.get("os")
        cpu_info = environment.get("cpu")
        if os_info:
            meta_lines.append(str(os_info))
        if cpu_info:
            meta_lines.append(str(cpu_info))
    generated = data.get("generated_at")
    subtitle = " · ".join(meta_lines) if meta_lines else None

    title = "BEVE Benchmark Comparison"
    if subtitle:
        title = f"{title}\n{subtitle}"
    fig.suptitle(title, fontsize=16)
    if generated:
        fig.text(0.99, 0.01, f"Generated: {generated}", ha="right", va="bottom", fontsize=8)

    fig.tight_layout(rect=[0, 0.03, 1, 0.95])
    output_path.parent.mkdir(parents=True, exist_ok=True)
    fig.savefig(output_path, dpi=200)
    plt.close(fig)


def main() -> None:
    args = parse_args()
    data = load_results(args.input)
    try:
        render_chart(data, args.output, args.metric, args.unit)
    except Exception as exc:  # pragma: no cover - defensive for CI visibility
        print(f"error: {exc}", file=sys.stderr)
        raise


if __name__ == "__main__":
    main()
