package main

import (
	"fmt"
	"math/big"
)

// Calculates the Collatz conjecture sequence
// https://en.wikipedia.org/wiki/Collatz_conjecture
func Collatz(number string) {
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)

	n := big.NewInt(0)
	n.SetString(number, 10)

	for {
		fmt.Println(n)
		if n.Bit(0) == 1 { // test if odd or even number by a simple check of the lowest bit
			// If odd multiply by 3 and add by one
			n.Mul(n, three)
			n.Add(n, one)
		} else {
			// Otherwise divide by 2
			n.Div(n, two)
		}
		if n.Cmp(one) == 0 {
			break
		}
	}
}

func main() {
	Collatz("1234982369834060ß9346ß093246123466")
}
