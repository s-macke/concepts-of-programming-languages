import java.util.Arrays;
import java.util.Comparator;

class IntComparator implements Comparator<Integer> {
    @Override
    public int compare(Integer i1, Integer i2) {
        return i1.compareTo(i2);
    }
}

class SortDemo {
    public static void main(String[] args) {
        Integer[] array = { 3, 2, 1, 5, 8, 6 };
        Arrays.sort(array, new IntComparator());
        System.out.println(Arrays.toString(array));
    }
}