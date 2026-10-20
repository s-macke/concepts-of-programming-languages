#!/bin/sh
# Runs all benchmarks for the prime number example of the GPU lecture
# and creates the charts and tables for the slides.
#
# NOTE: The parameters are set for an Apple M3 Max (CPU and GPU):
#   - CPU: 12 performance cores + 4 efficiency cores. The Java benchmark uses
#     only 12 threads, so macOS runs them on the 12 performance cores and the
#     steps in the charts are clean.
#   - GPU: 40 GPU cores, measured with Metal (swift). Runs only on macOS.
# On another machine, adjust CPU_THREADS (and the table rows in Benchmark.java).
#
# Requirements: Java 21+, Swift (Xcode Command Line Tools), uv
# Best results: restart, close all other apps, use the power adapter,
# no low power mode.
#
# Run with: ./run.sh

set -e
cd "$(dirname "$0")"

CPU_THREADS=12 # performance cores of the Apple M3 Max

echo "Benchmarks for an Apple M3 Max (CPU and GPU)."
if [ "$(uname -s)" != "Darwin" ] || [ "$(uname -m)" != "arm64" ]; then
    echo "Warning: this is not a Mac with Apple Silicon. The GPU benchmark needs Metal."
fi
CHIP=$(sysctl -n machdep.cpu.brand_string 2>/dev/null || echo unknown)
echo "This machine: $CHIP"
case "$CHIP" in
    *"M3 Max"*) ;;
    *) echo "Warning: the parameters are set for an Apple M3 Max, not for $CHIP." ;;
esac

echo
echo "1/4 CPU benchmark (Java, $CPU_THREADS threads, about 2 minutes)"
java Benchmark.java results "$CPU_THREADS"

echo
echo "2/4 GPU benchmark (Metal, about 6 minutes)"
swift Benchmark.swift results

echo
echo "3/4 Charts and tables"
uv run plot.py

echo
echo "4/4 Slides"
make -C ../../.. html

echo
echo "Done. Charts: docs/img/12-gpu-*.svg, tables: docs/snippets/12-gpu-*.html"
