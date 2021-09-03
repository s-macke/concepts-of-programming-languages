public class main {

static void inc(int count) {
    count++;
}

static class Count  {
    int count;
}

static void incStruct(Count count) {
    count.count++;
}

static void alterString(String s) {
    s = "world";
}

static void incArray(int[] count) {
    count[0]++;
}


public static void main(String... primitiveByValue) {
    int count = 1;
    inc(count);
    System.out.println(count);

    Count cclass = new Count();
    cclass.count = 1;
    incStruct(cclass);
    System.out.println(cclass.count);

    String s = "hello";
    alterString(s);
    System.out.println(s);

    int[1] countArray = {10};
    incArray(countArray);
    System.out.println(countArray[0]);



}

}
