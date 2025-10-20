package main

import (
	"fmt"
)

// Point is a two dimensional point in a cartesian coordinate system.
type Point struct{ x, y int }

// ColorPoint extends Point by adding a color field.
type ColorPoint struct {
	Point // Embedding simulates inheritance but it is (sort-of) delegation!
	c     int
}

// END1 OMIT

func main() {
	var p = Point{1, 2}
	var cp = ColorPoint{Point{1, 2}, 3}

	fmt.Println(p)
	fmt.Println(cp)

	fmt.Println(cp.x)       // access inherited field
	fmt.Println(cp.Point.x) // access inherited field via composition

	//p = cp // does not work: No type hierarchy, no polymorphism
	p = cp.Point // works

}

// END2 OMIT
