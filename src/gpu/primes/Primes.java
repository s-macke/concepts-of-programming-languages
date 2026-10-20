import java.util.Arrays;
import java.util.stream.IntStream;

public class Primes {

    // START1 OMIT
    static boolean isPrime(int n) {
        if (n < 2) return false;
        for (int d = 2; d < n; d++) {
            if (n % d == 0) return false;
        }
        return true;
    }
    // END1 OMIT

    public static void main(String[] args) {
        int count = args.length > 0 ? Integer.parseInt(args[0]) : 10_000;

        // START2 OMIT
        int[] numbers = new int[count];
        Arrays.fill(numbers, 1_000_003); // a prime number
        boolean[] result = new boolean[count];

        IntStream.range(0, count).parallel().forEach(i -> {
            result[i] = isPrime(numbers[i]);
        });
        // END2 OMIT

        long start = System.nanoTime();
        IntStream.range(0, count).parallel().forEach(i -> {
            result[i] = isPrime(numbers[i]);
        });
        long ms = (System.nanoTime() - start) / 1_000_000;
        System.out.println(count + " numbers, all prime: " + allTrue(result) + ", " + ms + " ms");
    }

    static boolean allTrue(boolean[] values) {
        for (boolean v : values) {
            if (!v) return false;
        }
        return true;
    }
}
