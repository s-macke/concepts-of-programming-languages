package main

import (
	"fmt"
)

// Point is a two dimensional point in a cartesian coordinate system.
type Point struct{ x, y int }

// Point implements the fmt.Stringer interface.
func (p Point) String() string {
	return fmt.Sprintf("x=%v,y=%v", p.x, p.y)
}

// ColorPoint extends Point by adding a color field.
type ColorPoint struct {
	Point // Embedding simulates inheritance but it is (sort-of) delegation!
	c     int
}

// ColorPoint implements the fmt.Stringer interface.
func (p ColorPoint) String() string {
	return fmt.Sprintf("x=%v,y=%v,c=%v", p.x, p.y, p.c)
	// OR: return fmt.Sprintf("%v,c=%v", p.Point, p.c)  // Delegate to Point.String()
}

// END1 OMIT

func main() {
	var p = Point{1, 2}
	var cp = ColorPoint{Point{1, 2}, 3} // embeds Point
	fmt.Println(p)
	fmt.Println(cp)
	fmt.Println(cp.x) // access inherited field

	// p = cp       // does not work: No hierarchy, no polymorphism
	// p = cp.Point // works

	// s is an interface and supports Polymorphism
	var s fmt.Stringer
	s = p // check at compile time
	fmt.Println(s)
	s = cp
	fmt.Println(s)
}

// END2 OMIT
