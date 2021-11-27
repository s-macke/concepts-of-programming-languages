public class Add
{
    public int Add(int a, int b) {
        return a + b;
    }

    public int Add(int a, int b, int c) {
        return a + b + c;
    }

    public int ifthen(int a, int b, int c) {
        int d;
        if (a > b) {
            d = a + c;
        } else {
            d = b + c;
        }
        return d;
    }

    public int forloop() {
        int a = 0;
        for(int i=0; i<10; i++) {
            a = a + i;
        }
        return a;
    }


}
