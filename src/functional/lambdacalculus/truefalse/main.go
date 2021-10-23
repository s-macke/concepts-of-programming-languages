package main

import (
	"fmt"
)

// This is the key: A recursive function definition for all functions!!!
type fnf func(fnf) fnf

func main() {
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

	fmt.Println(TRUE.ToBool())
	fmt.Println(FALSE.ToBool())
}

// ToBool Returns an actual bool for a Church Boolean
func (f fnf) ToBool() bool {
	// λx.x is a function which returns itself (the ID)
	ID := func(x fnf) fnf { return x }
	var ret bool
	f(func(f fnf) fnf { ret = true; return f })(func(f fnf) fnf { ret = false; return f })(ID)
	return ret
}
