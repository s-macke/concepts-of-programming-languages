# /// script
# requires-python = ">=3.10"
# dependencies = ["matplotlib"]
# ///
"""Plots the CSV files of Benchmark.java as separate SVG charts for the GPU lecture.

Run with: uv run plot.py [resultsDirectory] [imageDirectory]
"""
import csv
import re
import sys
from pathlib import Path

import matplotlib

matplotlib.use("Agg")
import matplotlib.pyplot as plt

results = Path(sys.argv[1] if len(sys.argv) > 1 else "results")
images = Path(sys.argv[2] if len(sys.argv) > 2 else "../../../docs/img")
images.mkdir(parents=True, exist_ok=True)

plt.rcParams.update({"font.size": 14, "svg.fonttype": "none"})


def read(path):
    lines = path.read_text().splitlines()
    info = lines[0].lstrip("# ")
    rows = list(csv.DictReader(lines[1:]))
    return info, rows


def cores_of(info):
    return int(re.search(r"cores=(\d+)", info).group(1))


def save(fig, info, name):
    fig.text(0.99, 0.01, info, ha="right", fontsize=9, color="gray")
    fig.tight_layout()
    fig.savefig(images / name)
    plt.close(fig)
    print("written:", images / name)


def scaling_chart(path, prime):
    info, rows = read(path)
    n = [int(r["elements"]) for r in rows]
    ms = [float(r["ms"]) for r in rows]

    fig, ax = plt.subplots(figsize=(10, 5.5))
    ax.plot(n, ms, "o-", markersize=4)
    ax.set_xlabel("number of elements")
    ax.set_ylabel("time [ms]")
    ax.set_title(f"CPU: every element is the prime number {prime:,}")
    ax.set_xlim(left=0)
    ax.set_ylim(bottom=0)
    ax.grid(alpha=0.3)
    save(fig, info, f"12-gpu-cpu-{prime}.svg")


def zoom_chart(path, prime, max_elements=128):
    info, rows = read(path)
    rows = [r for r in rows if int(r["elements"]) <= max_elements]
    n = [int(r["elements"]) for r in rows]
    ms = [float(r["ms"]) for r in rows]

    fig, ax = plt.subplots(figsize=(10, 5.5))
    ax.plot(n, ms, "o-", markersize=4)
    ax.set_xlabel("number of elements")
    ax.set_ylabel("time [ms]")
    ax.set_title(f"CPU, zoomed in: the prime number {prime:,}")
    ax.set_xticks(range(0, max_elements + 1, 8))
    ax.set_xlim(left=0)
    ax.set_ylim(bottom=0)
    ax.grid(alpha=0.3)
    save(fig, info, f"12-gpu-cpu-{prime}-zoom.svg")


def mix_chart(path):
    info, rows = read(path)
    fig, ax = plt.subplots(figsize=(10, 5.5))
    labels = [re.sub(r"\d{5,}", lambda m: f"{int(m.group()):,}", r["variant"]).replace(" (", "\n(") for r in rows]
    ax.bar(labels, [float(r["ms"]) for r in rows], color=["C0", "C1", "C2", "C3"])
    ax.set_ylabel("time [ms]")
    elements = re.search(r"elements=(\d+)", info).group(1)
    ax.set_title(f"CPU: {int(elements):,} elements")
    ax.grid(axis="y", alpha=0.3)
    save(fig, info, "12-gpu-cpu-mix.svg")


def comparison_chart(paths_and_primes):
    """The same scaling chart, with the larger prime number added for comparison."""
    fig, ax = plt.subplots(figsize=(10, 5.5))
    for (path, prime), marker in zip(paths_and_primes, ["o-", "s-"]):
        info, rows = read(path)
        ax.plot([int(r["elements"]) for r in rows], [float(r["ms"]) for r in rows],
                marker, markersize=4, label=f"prime number {prime:,}")
    ax.set_xlabel("number of elements")
    ax.set_ylabel("time [ms]")
    ax.set_title("CPU: the larger prime number in comparison")
    ax.set_xlim(left=0)
    ax.set_ylim(bottom=0)
    ax.legend()
    ax.grid(alpha=0.3)
    save(fig, info, f"12-gpu-cpu-{paths_and_primes[-1][1]}.svg")


TABLE_STYLE = """<style>
    table.measured {
        width: auto;
        border-collapse: collapse;
        font-size: 26px;
        line-height: 1.3;
        font-family: inherit;
        margin: 20px 0;
    }
    table.measured th, table.measured td {
        border: 2px solid #444;
        padding: 8px 40px;
        text-align: right;
    }
    table.measured th {
        background: #eee;
    }
</style>
"""


def table_snippets(path, prime):
    """Writes one HTML table per reveal step: step k shows the first k values, the rest as '?'.
    The rows come from Benchmark.java (TABLE_ELEMENTS, fastest of many runs)."""
    info, rows = read(path)
    table = [(int(r["elements"]), float(r["ms"])) for r in rows]
    snippets = images.parent / "snippets"
    snippets.mkdir(parents=True, exist_ok=True)
    for k in range(len(table) + 1):
        lines = [TABLE_STYLE, '<table class="measured">',
                 f"<tr><th>elements</th><th>time for {prime:,} [ms]</th></tr>"]
        for i, (n, ms) in enumerate(table):
            value = f"{ms:.2f}" if i < k else "?"
            lines.append(f"<tr><td>{n}</td><td>{value}</td></tr>")
        lines.append("</table>")
        name = snippets / f"12-gpu-cpu-table-{k}.html"
        name.write_text("\n".join(lines) + "\n")
    print("written:", snippets / f"12-gpu-cpu-table-[0..{len(table)}].html")


# Default rows for the question "When does the GPU take longer?". Replace by measuring.
GPU_LARGE_ELEMENTS = [1_000, 2_000, 3_000, 4_000, 5_000, 10_000]


def read_optional(path):
    """Reads elements,ms from a CSV file, if it exists. Returns {elements: ms}."""
    if not path.exists():
        return {}
    info, rows = read(path)
    return {int(r["elements"]): float(r["ms"]) for r in rows}


def comparison_snippets(cpu_table_path, prime):
    """CPU vs. GPU table for the small element counts.
    Step 0: GPU column as '?'. Step 1: GPU values from results/gpu-table-<prime>.csv (empty if missing)."""
    info, rows = read(cpu_table_path)
    cpu = [(int(r["elements"]), float(r["ms"])) for r in rows]
    gpu = read_optional(results / f"gpu-table-{prime}.csv")
    snippets = images.parent / "snippets"
    for step in range(2):
        lines = [TABLE_STYLE, '<table class="measured">',
                 f"<tr><th>elements</th><th>CPU [ms]</th><th>GPU [ms]</th></tr>"]
        for n, ms in cpu:
            g = "?" if step == 0 else (f"{gpu[n]:.2f}" if n in gpu else "&nbsp;")
            lines.append(f"<tr><td>{n}</td><td>{ms:.2f}</td><td>{g}</td></tr>")
        lines.append("</table>")
        (snippets / f"12-gpu-compare-table-{step}.html").write_text("\n".join(lines) + "\n")
    print("written:", snippets / "12-gpu-compare-table-[0..1].html")


def gpu_large_snippet(prime):
    """GPU table for large element counts. Values from results/gpu-large-<prime>.csv (empty if missing)."""
    gpu = read_optional(results / f"gpu-large-{prime}.csv")
    rows = sorted(gpu) if gpu else GPU_LARGE_ELEMENTS
    lines = [TABLE_STYLE, '<table class="measured">',
             f"<tr><th>elements</th><th>GPU [ms]</th></tr>"]
    for n in rows:
        g = f"{gpu[n]:.2f}" if n in gpu else "&nbsp;"
        lines.append(f"<tr><td>{n:,}</td><td>{g}</td></tr>")
    lines.append("</table>")
    path = images.parent / "snippets" / "12-gpu-large-table.html"
    path.write_text("\n".join(lines) + "\n")
    print("written:", path)


def series(path):
    info, rows = read(path)
    return info, [int(r["elements"]) for r in rows], [float(r["ms"]) for r in rows]


def line_chart(lines, title, name, xmax=None):
    """lines: list of (label, elements, ms). Writes one SVG chart."""
    fig, ax = plt.subplots(figsize=(10, 5.5))
    for (label, n, ms), marker in zip(lines, ["o-", "s-", "^-"]):
        if xmax is not None:
            n, ms = zip(*[(a, b) for a, b in zip(n, ms) if a <= xmax])
        ax.plot(n, ms, marker, markersize=4, label=label)
    ax.set_xlabel("number of elements")
    ax.set_ylabel("time [ms]")
    ax.set_title(title)
    ax.set_xlim(left=0)
    ax.set_ylim(bottom=0)
    if len(lines) > 1:
        ax.legend()
    ax.grid(alpha=0.3)
    save(fig, "", name)


def gpu_mix_chart(path):
    info, rows = read(path)
    sizes = sorted({int(r["elements"]) for r in rows})
    fig, axes = plt.subplots(1, len(sizes), figsize=(12, 5.5))
    for ax, size in zip(axes, sizes):
        part = [r for r in rows if int(r["elements"]) == size]
        labels = [re.sub(r"\d{5,}", lambda m: f"{int(m.group()):,}", r["variant"]).replace(" (", "\n(") for r in part]
        ax.bar(labels, [float(r["ms"]) for r in part], color=["C0", "C1", "C2", "C3"])
        ax.set_title(f"GPU: {size:,} elements")
        ax.set_ylabel("time [ms]")
        ax.tick_params(axis="x", labelsize=9)
        ax.grid(axis="y", alpha=0.3)
    save(fig, info, "12-gpu-gpu-mix.svg")


def gpu_charts(small_prime, large_prime):
    small = results / f"gpu-scaling-{small_prime}.csv"
    if not small.exists():
        print("no GPU results, skipping GPU charts")
        return
    _, n, ms = series(small)
    _, n2, ms2 = series(results / f"gpu-scaling-{large_prime}.csv")
    _, cn, cms = series(results / f"cpu-scaling-{small_prime}.csv")
    line_chart([(f"prime number {small_prime:,}", n, ms)],
               f"GPU: every element is the prime number {small_prime:,}", f"12-gpu-gpu-{small_prime}.svg")
    line_chart([(f"prime number {small_prime:,}", n, ms)],
               f"GPU, zoomed in: the prime number {small_prime:,}", f"12-gpu-gpu-{small_prime}-zoom.svg", xmax=40_000)
    line_chart([(f"prime number {small_prime:,}", n, ms), (f"prime number {large_prime:,}", n2, ms2)],
               "GPU: the larger prime number in comparison", f"12-gpu-gpu-{large_prime}.svg")
    line_chart([("CPU", cn, cms), ("GPU", n, ms)],
               f"CPU and GPU: the prime number {small_prime:,}", "12-gpu-cpu-vs-gpu.svg", xmax=12_000)
    line_chart([("CPU", cn, cms), ("GPU", n, ms)],
               f"CPU and GPU, zoomed in: the prime number {small_prime:,}", "12-gpu-cpu-vs-gpu-zoom.svg", xmax=128)
    gpu_mix_chart(results / "gpu-mix.csv")


scalings = sorted(
    (path, int(re.search(r"cpu-scaling-(\d+)\.csv", path.name).group(1)))
    for path in results.glob("cpu-scaling-*.csv")
)
scalings.sort(key=lambda pp: pp[1])
small_path, small_prime = scalings[0]
scaling_chart(small_path, small_prime)  # step 1: only the small prime number
zoom_chart(small_path, small_prime)  # zoomed in
table_snippets(results / f"cpu-table-{small_prime}.csv", small_prime)  # the table that is revealed row by row
comparison_snippets(results / f"cpu-table-{small_prime}.csv", small_prime)  # CPU vs. GPU
gpu_large_snippet(small_prime)  # when does the GPU take longer?
comparison_chart(scalings)  # step 2: the larger prime number in comparison
mix_chart(results / "cpu-mix.csv")
gpu_charts(small_prime, scalings[-1][1])
