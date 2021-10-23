package main

import (
	"fmt"
)

// This is the key: A Recursive function definition for all functions!!!
type fnf func(fnf) fnf

// ToBool Returns an actual bool for a Church Boolean
func (f fnf) ToBool() bool {
	ID := func(x fnf) fnf { return x }
	var ret bool
	f(func(f fnf) fnf { ret = true; return f })(func(f fnf) fnf { ret = false; return f })(ID)
	return ret
}

func main() {

	// λx.x is a function which returns itself (the ID)
	//ID := func(x fnf) fnf { return x }

	// Boolean TRUE as function: λx.λy.x
	TRUE := fnf(func(x fnf) fnf {
		return func(y fnf) fnf {
			return x
		}
	})

	// Boolean FALSE as function: λx.λy.y
	FALSE := fnf(func(x fnf) fnf {
		return func(y fnf) fnf {
			return y
		}
	})

	// Boolean Logic
	NOT := func(p fnf) fnf { return p(FALSE)(TRUE) } // λb.b False True
	AND := func(p fnf) fnf { return func(q fnf) fnf { return p(q)(p) } }
	OR := func(p fnf) fnf { return func(q fnf) fnf { return p(p)(q) } }
	EQ := func(p fnf) fnf { return func(q fnf) fnf { return p(q)(NOT(q)) } }

	fmt.Println(NOT(FALSE).ToBool())
	fmt.Println(NOT(TRUE).ToBool())

	fmt.Println(AND(FALSE)(TRUE).ToBool())
	fmt.Println(OR(FALSE)(FALSE).ToBool())
	fmt.Println(EQ(TRUE)(TRUE).ToBool())
}
