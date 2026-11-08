import java.util.Arrays;

class SortLambdaDemo {
    public static void main(String[] args) {
        Integer[] array = { 3, 2, 1, 5, 8, 6 };
        Arrays.sort(array, (a, b) -> a.compareTo(b));
        System.out.println(Arrays.toString(array));
    }
}
