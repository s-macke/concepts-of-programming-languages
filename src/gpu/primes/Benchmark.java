import java.io.IOException;
import java.io.PrintWriter;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Random;
import java.util.TreeSet;
import java.util.stream.IntStream;

// CPU benchmark for the prime number example of the GPU lecture.
// All numbers are 32-bit, so the same experiment runs on the GPU as well.
//
// Run with: java Benchmark.java [outputDirectory] [threads]
// threads: optional, limits the Parallel For to this many threads, e.g. 12 on an
// Apple M3 Max to use (mostly) only the 12 performance cores.
// Writes CSV files that plot.py turns into SVG charts.
public class Benchmark {

    static final int SMALL_PRIME = 1_000_003;
    static final int LARGE_PRIME = 2_000_003; // at most twice as large

    static final int MAX_ELEMENTS = 12_000;

    // Element counts for the table on the slides. With 12 threads (the 12 performance cores
    // of the Apple M3 Max) the steps are at 12 -> 13 and 24 -> 25.
    static final int[] TABLE_ELEMENTS = {1, 2, 11, 12, 13, 24, 25};
    static final int TABLE_REPETITIONS = 100;
    static final int MIX_ELEMENTS = 10_000;

    // Deliberately naive: tests all divisors up to n. So the work grows linearly with n,
    // and a prime number twice as large takes twice as long.
    static boolean isPrime(int n) {
        if (n < 2) return false;
        for (int d = 2; d < n; d++) {
            if (n % d == 0) return false;
        }
        return true;
    }

    // Tests all numbers with a Parallel For and returns the fastest runtime in milliseconds.
    // The fastest run is the one with the least disturbance by other programs.
    // Short runs are repeated more often to reduce the noise.
    static double measure(int[] numbers) {
        boolean[] result = new boolean[numbers.length];
        List<Double> times = new ArrayList<>();
        double total = 0;
        while (times.size() < 5 || (total < 200 && times.size() < 500)) {
            long start = System.nanoTime();
            IntStream.range(0, numbers.length).parallel().forEach(i -> {
                result[i] = isPrime(numbers[i]);
            });
            double ms = (System.nanoTime() - start) / 1e6;
            times.add(ms);
            total += ms;
        }
        for (boolean b : result) {
            if (!b) throw new IllegalStateException("not prime");
        }
        return times.stream().mapToDouble(Double::doubleValue).min().getAsDouble();
    }

    // Runs the Parallel For many times and returns the fastest run in milliseconds.
    // The fastest run is the one with the least disturbance by other programs.
    static double measureMin(int[] numbers, int repetitions) {
        boolean[] result = new boolean[numbers.length];
        double min = Double.MAX_VALUE;
        for (int r = 0; r < repetitions; r++) {
            long start = System.nanoTime();
            IntStream.range(0, numbers.length).parallel().forEach(i -> {
                result[i] = isPrime(numbers[i]);
            });
            min = Math.min(min, (System.nanoTime() - start) / 1e6);
        }
        return min;
    }

    static int[] filled(int count, int value) {
        int[] numbers = new int[count];
        Arrays.fill(numbers, value);
        return numbers;
    }

    // 5 % large primes, either randomly distributed or next to each other in one bunch.
    static int[] mixed(int count, boolean together) {
        int[] numbers = filled(count, SMALL_PRIME);
        int large = Math.max(1, count / 20);
        if (together) {
            for (int i = 0; i < large; i++) numbers[i] = LARGE_PRIME;
        } else {
            Random random = new Random(42);
            int placed = 0;
            while (placed < large) {
                int i = random.nextInt(count);
                if (numbers[i] != LARGE_PRIME) {
                    numbers[i] = LARGE_PRIME;
                    placed++;
                }
            }
        }
        return numbers;
    }

    // Every count from 1 to 128, then about 40 logarithmically distributed counts up to max.
    static List<Integer> counts(int max) {
        TreeSet<Integer> set = new TreeSet<>();
        for (int n = 1; n <= Math.min(128, max); n++) set.add(n);
        for (int k = 0; k <= 40; k++) {
            set.add((int) Math.round(Math.pow(max, k / 40.0)));
        }
        return new ArrayList<>(set);
    }

    public static void main(String[] args) throws IOException {
        Path out = Path.of(args.length > 0 ? args[0] : "results");
        Files.createDirectories(out);
        int cores = Runtime.getRuntime().availableProcessors();
        if (args.length > 1) {
            // Must be set before the first parallel stream: the calling thread helps,
            // so the pool needs one thread less.
            cores = Integer.parseInt(args[1]);
            System.setProperty("java.util.concurrent.ForkJoinPool.common.parallelism", String.valueOf(cores - 1));
        }
        System.out.println("CPU cores: " + cores);

        // Warm up for about 2 seconds of continuous load: the JIT compiler optimizes the code,
        // the clock speed goes up, and the operating system learns that these threads are
        // compute-heavy (on Apple Silicon: performance cores instead of efficiency cores).
        long warmupEnd = System.nanoTime() + 2_000_000_000L;
        while (System.nanoTime() < warmupEnd) {
            measureMin(filled(cores * 4, SMALL_PRIME), 1);
        }

        // The table for the slides: the fastest of many runs.
        List<String> table = new ArrayList<>();
        table.add("elements,ms");
        for (int n : TABLE_ELEMENTS) {
            table.add(n + "," + measureMin(filled(n, SMALL_PRIME), TABLE_REPETITIONS));
        }
        write(out.resolve("cpu-table-" + SMALL_PRIME + ".csv"), table, cores,
                "prime=" + SMALL_PRIME + ", fastest of " + TABLE_REPETITIONS + " runs");
        table.forEach(System.out::println);

        // Experiment 1 and 2: the same prime number in every element.
        for (int prime : new int[]{SMALL_PRIME, LARGE_PRIME}) {
            List<String> lines = new ArrayList<>();
            lines.add("elements,ms");
            for (int n : counts(MAX_ELEMENTS)) {
                double ms = measure(filled(n, prime));
                lines.add(n + "," + ms);
            }
            write(out.resolve("cpu-scaling-" + prime + ".csv"), lines, cores, "prime=" + prime);
            System.out.println("prime " + prime + " done");
        }

        // Experiment 3: 95 % small and 5 % large prime numbers.
        List<String> mix = new ArrayList<>();
        mix.add("variant,ms");
        mix.add("all " + SMALL_PRIME + "," + measure(filled(MIX_ELEMENTS, SMALL_PRIME)));
        mix.add("5% " + LARGE_PRIME + " (random)," + measure(mixed(MIX_ELEMENTS, false)));
        mix.add("5% " + LARGE_PRIME + " (together)," + measure(mixed(MIX_ELEMENTS, true)));
        mix.add("all " + LARGE_PRIME + "," + measure(filled(MIX_ELEMENTS, LARGE_PRIME)));
        write(out.resolve("cpu-mix.csv"), mix, cores, "elements=" + MIX_ELEMENTS);
        mix.forEach(System.out::println);
    }

    static void write(Path file, List<String> lines, int cores, String info) throws IOException {
        try (PrintWriter w = new PrintWriter(Files.newBufferedWriter(file))) {
            w.println("# cores=" + cores + ", " + info + ", os=" + System.getProperty("os.name")
                    + ", arch=" + System.getProperty("os.arch"));
            lines.forEach(w::println);
        }
    }
}
