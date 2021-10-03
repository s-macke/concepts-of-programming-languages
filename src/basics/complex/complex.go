package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	c := 3 + 4i
	fmt.Printf("c=%v\n", c)
	r, θ := cmplx.Polar(c) // Go supports variables with unicode characters
	fmt.Printf("r=%v, θ=%v\n", r, θ)
}
