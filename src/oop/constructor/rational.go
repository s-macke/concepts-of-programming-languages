package main

import "fmt"

// Rational represents a rational number numerator/denominator.
type Rational struct {
	numerator   int
	denominator int
}

// Constructor
func NewRational(numerator int, denominator int) Rational {
	if denominator == 0 {
		panic("division by zero")
	}
	return Rational{
		numerator:   numerator,
		denominator: denominator,
	}
}

// END OMIT
func main() {
	rat := NewRational(1, 2)
	fmt.Println(rat)
}
