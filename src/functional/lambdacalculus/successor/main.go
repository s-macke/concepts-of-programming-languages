package main

import "fmt"

type fnf func(fnf) fnf

func main() {
	I := func(x fnf) fnf { return x }

	// Functional Numbers 1
	ONCE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(x)
		}
	}

	// Functional Numbers 2
	TWICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(x))
		}
	}

	// Functional Numbers 3
	THRICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(f(x)))
		}
	}
	// NUMBER END OMIT

	// Functional Numbers SUCCESSOR(N) = N + 1
	SUCCESSOR := func(w fnf) fnf {
		return func(y fnf) fnf {
			return func(x fnf) fnf {
				return y(w)(y)(x)
			}
		}
	}

	Printer := func(x fnf) fnf { fmt.Print("."); return x }

	SUCCESSOR(ONCE)(Printer)(I)
	fmt.Println("SUCCESSOR(ONCE) = 2")

	SUCCESSOR(TWICE)(Printer)(I)
	fmt.Println("SUCCESSOR(TWICE) = 3")

	SUCCESSOR(THRICE)(Printer)(I)
	fmt.Println("SUCCESSOR(THRICE) = 4")
	// EOF OMIT
}
