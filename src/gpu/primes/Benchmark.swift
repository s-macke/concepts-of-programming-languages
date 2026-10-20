// GPU benchmark for the prime number example of the GPU lecture (Apple GPUs, Metal).
// The same naive prime test, the same 32-bit prime numbers and the same experiments
// as Benchmark.java, so CPU and GPU can be compared directly.
//
// Run with: swift Benchmark.swift [outputDirectory]
// Writes CSV files that plot.py turns into SVG charts.
// Only measures the GPU time (gpuStartTime to gpuEndTime), not the copying of the data.

import Foundation
import Metal

let smallPrime: UInt32 = 1_000_003
let largePrime: UInt32 = 2_000_003 // at most twice as large

let tableElements = [1, 2, 11, 12, 13, 24, 25] // the same rows as the CPU table
let largeElements = [1_000, 2_000, 5_000, 10_000, 20_000, 40_000]
let maxElements = 100_000
let mixElements = [1_000, 100_000] // case 1: everything runs at once, case 2: many more elements than lanes

// The kernel is the body of the Parallel For: one thread per element.
let source = """
#include <metal_stdlib>
using namespace metal;

kernel void isPrime(device const uint *numbers [[buffer(0)]],
                    device uint *result [[buffer(1)]],
                    constant uint &count [[buffer(2)]],
                    uint i [[thread_position_in_grid]]) {
    if (i >= count) return;
    uint n = numbers[i];
    uint prime = n >= 2 ? 1 : 0;
    for (uint d = 2; d < n; d++) {
        if (n % d == 0) {
            prime = 0;
            break;
        }
    }
    result[i] = prime;
}
"""

let device = MTLCreateSystemDefaultDevice()!
let library = try! device.makeLibrary(source: source, options: nil)
let pipeline = try! device.makeComputePipelineState(function: library.makeFunction(name: "isPrime")!)
let queue = device.makeCommandQueue()!
let threadsPerGroup = 256

// Runs the kernel once over all numbers and returns the GPU time in milliseconds.
func runOnce(_ numbers: [UInt32], check: Bool) -> Double {
    var data = numbers
    let input = device.makeBuffer(bytes: &data, length: 4 * data.count, options: .storageModeShared)!
    let output = device.makeBuffer(length: 4 * data.count, options: .storageModeShared)!
    var count = UInt32(data.count)

    let commands = queue.makeCommandBuffer()!
    let encoder = commands.makeComputeCommandEncoder()!
    encoder.setComputePipelineState(pipeline)
    encoder.setBuffer(input, offset: 0, index: 0)
    encoder.setBuffer(output, offset: 0, index: 1)
    encoder.setBytes(&count, length: 4, index: 2)
    encoder.dispatchThreads(MTLSize(width: data.count, height: 1, depth: 1),
                            threadsPerThreadgroup: MTLSize(width: threadsPerGroup, height: 1, depth: 1))
    encoder.endEncoding()
    commands.commit()
    commands.waitUntilCompleted()

    if check {
        let result = output.contents().bindMemory(to: UInt32.self, capacity: data.count)
        for i in 0..<data.count where result[i] != 1 {
            fatalError("element \(i) is not prime")
        }
    }
    return (commands.gpuEndTime - commands.gpuStartTime) * 1000
}


// The fastest of several runs: the one with the least disturbance by other programs.
func fastest(_ numbers: [UInt32], repetitions: Int = 5) -> Double {
    (0..<repetitions).map { r in runOnce(numbers, check: r == 0) }.min()!
}

func filled(_ count: Int, _ value: UInt32) -> [UInt32] {
    [UInt32](repeating: value, count: count)
}

// 5 % large primes, either randomly distributed or next to each other in one bunch.
func mixed(_ count: Int, together: Bool) -> [UInt32] {
    var numbers = filled(count, smallPrime)
    let large = max(1, count / 20)
    if together {
        for i in 0..<large { numbers[i] = largePrime }
    } else {
        var generator = SplitMix64(seed: 42)
        var placed = 0
        while placed < large {
            let i = Int(generator.next() % UInt64(count))
            if numbers[i] != largePrime {
                numbers[i] = largePrime
                placed += 1
            }
        }
    }
    return numbers
}

// A small deterministic random generator, so every run uses the same distribution.
struct SplitMix64 {
    var state: UInt64
    init(seed: UInt64) { state = seed }
    mutating func next() -> UInt64 {
        state &+= 0x9E37_79B9_7F4A_7C15
        var z = state
        z = (z ^ (z >> 30)) &* 0xBF58_476D_1CE4_E5B9
        z = (z ^ (z >> 27)) &* 0x94D0_49BB_1331_11EB
        return z ^ (z >> 31)
    }
}

// Logarithmically distributed counts up to max, plus every 500 elements up to 40,000.
func counts(_ max: Int) -> [Int] {
    var set = Set<Int>()
    for k in 0...40 {
        set.insert(Int((pow(Double(max), Double(k) / 40)).rounded()))
    }
    for n in stride(from: 500, through: 40_000, by: 500) {
        set.insert(n)
    }
    return set.sorted()
}

func write(_ file: URL, _ lines: [String], _ info: String) {
    let header = "# gpu=\(device.name), simd=\(pipeline.threadExecutionWidth), \(info), os=macOS"
    try! ([header] + lines).joined(separator: "\n").appending("\n").write(to: file, atomically: true, encoding: .utf8)
}

let out = URL(fileURLWithPath: CommandLine.arguments.count > 1 ? CommandLine.arguments[1] : "results")
try! FileManager.default.createDirectory(at: out, withIntermediateDirectories: true)
print("GPU: \(device.name), SIMD width: \(pipeline.threadExecutionWidth)")

// Warm up for about 2 seconds, so the GPU reaches its clock speed.
let warmupEnd = Date().addingTimeInterval(2)
while Date() < warmupEnd {
    _ = runOnce(filled(10_000, smallPrime), check: false)
}

// The table for the slides: the fastest of 20 runs.
var table = ["elements,ms"]
for n in tableElements {
    table.append("\(n),\(fastest(filled(n, smallPrime), repetitions: 20))")
}
write(out.appendingPathComponent("gpu-table-\(smallPrime).csv"), table, "prime=\(smallPrime), fastest of 20 runs")
print(table.joined(separator: "\n"))

// When does the GPU take longer?
var large = ["elements,ms"]
for n in largeElements {
    large.append("\(n),\(fastest(filled(n, smallPrime), repetitions: 20))")
}
write(out.appendingPathComponent("gpu-large-\(smallPrime).csv"), large, "prime=\(smallPrime), fastest of 20 runs")
print(large.joined(separator: "\n"))

// Experiment 1 and 2: the same prime number in every element.
for prime in [smallPrime, largePrime] {
    var lines = ["elements,ms"]
    for n in counts(maxElements) {
        lines.append("\(n),\(fastest(filled(n, prime)))")
    }
    write(out.appendingPathComponent("gpu-scaling-\(prime).csv"), lines, "prime=\(prime)")
    print("prime \(prime) done")
}

// Experiment 3: 95 % small and 5 % large prime numbers, for a small and a large array.
var mix = ["variant,elements,ms"]
for n in mixElements {
    mix.append("all \(smallPrime),\(n),\(fastest(filled(n, smallPrime)))")
    mix.append("5% \(largePrime) (random),\(n),\(fastest(mixed(n, together: false)))")
    mix.append("5% \(largePrime) (together),\(n),\(fastest(mixed(n, together: true)))")
    mix.append("all \(largePrime),\(n),\(fastest(filled(n, largePrime)))")
}
write(out.appendingPathComponent("gpu-mix.csv"), mix, "mix")
print(mix.joined(separator: "\n"))
