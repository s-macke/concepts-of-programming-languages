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
    public void foo() {
        System.out.print("b.foo ");
        super.foo();
    }

    public void bar() {
        System.out.println("b.bar");
    }
}

public class Main {
    public static void main(String[] args) {
        A a = new A();
        a.foo();

        B b = new B();
        b.foo(); // "b.foo a.foo a.bar" or "b.foo a.foo b.bar"?
    }
}