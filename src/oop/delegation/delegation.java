class A {
    public void foo() {
        System.out.print("a.foo ");
        bar();
    }

    public void bar() {
        System.out.println("a.bar");
    }
}

class B extends A {
    @Override
    public void foo() {
        System.out.print("b.foo ");
        bar();
    }
    @Override
    public void bar() {
        System.out.println("b.bar");
    }

}

public class Main {
    public static void main(String[] args) {
        // A a = new A();
        // a.foo();

        B b = new B();
        b.foo(); // Will print "b.foo b.bar"
    }
}