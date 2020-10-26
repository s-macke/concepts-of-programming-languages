package main

import "fmt"

type fnf func(fnf) fnf

// NAMER START OMIT
type namer struct {
	names map[string]string
}

func NewNamer() namer {
	return namer{names: map[string]string{}}
}

// register a known function (pointer) with a name
func (n *namer) register(x fnf, name string) {
	n.names[fmt.Sprintf("%p", x)] = name
}

// get the name of a function (pointer)
func (n *namer) nameOf(x fnf) string {
	s := n.names[fmt.Sprintf("%p", x)]
	if s == "" {
		return fmt.Sprintf("%p", x)
	}
	return s
}

// NAMER END OMIT

func main() {
	namer := NewNamer()

	// Combinators (pure functions, i.e. all variables are bound)
	I := func(x fnf) fnf { return x }                            // λx.x  = identity
	M := func(x fnf) fnf { return x(x) }                         // λx.xx = mocking bird
	T := func(x fnf) fnf { return func(y fnf) fnf { return x } } // λx.λy.x = kestrel/true/const
	//F := func(x fnf) fnf { return func(y fnf) fnf { return y } } // λx.λy.y = kite/false
	F := T(I)
	// COMB OMIT

	// λf.λx.λy.fyx = cardinal
	C := func(f fnf) fnf {
		return func(x fnf) fnf {
			return func(y fnf) fnf {
				return f(y)(x)
			}
		}
	}

	namer.register(I, "I")
	namer.register(M, "M")
	namer.register(C, "C")
	namer.register(T, "T")
	namer.register(F, "F")

	// Boolean Logic
	NOT := func(p fnf) fnf { return p(F)(T) }
	AND := func(p fnf) fnf { return func(q fnf) fnf { return p(q)(p) } }
	OR := func(p fnf) fnf { return func(q fnf) fnf { return p(p)(q) } }
	EQ := func(p fnf) fnf { return func(q fnf) fnf { return p(q)(NOT(q)) } }
	// BOOL OMIT

	namer.register(NOT, "NOT")
	namer.register(AND, "AND")
	namer.register(OR, "OR")
	namer.register(EQ, "EQ")

	// TEST COMB START OMIT
	fmt.Printf("I(I) = %s\n", namer.nameOf(I(I)))
	fmt.Printf("M(I) = %s\n", namer.nameOf(M(I)))
	fmt.Printf("T = %s\n", namer.nameOf(T))
	fmt.Printf("T(I) = %s\n", namer.nameOf(T(I)))
	fmt.Printf("T(M)(I) = %s\n", namer.nameOf(T(M)(I)))
	fmt.Printf("T(I)(I)(M) = %s\n", namer.nameOf(T(I)(I)(M)))
	//fmt.Printf("M(M) = %p\n", M(M)) // omega
	// TEST COMB END OMIT
	// fmt.Printf("C(K)(I)(M) = %s\n", namer.nameOf(C(K)(I)(M)))
	fmt.Print("\n")
	// TEST BOOL START OMIT
	fmt.Printf("NOT(T) = %s\n", namer.nameOf(NOT(T)))
	fmt.Printf("NOT(F) = %s\n", namer.nameOf(NOT(F)))
	fmt.Printf("AND(T)(T) = %s\n", namer.nameOf(AND(T)(T)))
	fmt.Printf("AND(T)(F) = %s\n", namer.nameOf(AND(T)(F)))
	fmt.Printf("AND(F)(T) = %s\n", namer.nameOf(AND(F)(T)))
	fmt.Printf("AND(F)(F) = %s\n", namer.nameOf(AND(F)(F)))
	fmt.Printf("OR(T)(T) = %s\n", namer.nameOf(OR(T)(T)))
	fmt.Printf("OR(T)(F) = %s\n", namer.nameOf(OR(T)(F)))
	fmt.Printf("OR(F)(T) = %s\n", namer.nameOf(OR(F)(T)))
	fmt.Printf("OR(F)(F) = %s\n", namer.nameOf(OR(F)(F)))
	fmt.Printf("EQ(T)(T) = %s\n", namer.nameOf(EQ(T)(T)))
	fmt.Printf("EQ(T)(F) = %s\n", namer.nameOf(EQ(T)(F)))
	fmt.Printf("EQ(F)(T) = %s\n", namer.nameOf(EQ(F)(T)))
	fmt.Printf("EQ(F)(F) = %s\n", namer.nameOf(EQ(F)(F)))
	// TEST BOOL END OMIT

	// Debugging Functions
	f := func(x fnf) fnf { fmt.Printf("f()\n"); return x }
	g := func(y fnf) fnf { fmt.Printf("g()\n"); return y }

	// select and call first function f(I)
	T(f)(g)(I)

	// select and call second function g(I)
	F(f)(g)(I)
	// DEBUG END OMIT
}
