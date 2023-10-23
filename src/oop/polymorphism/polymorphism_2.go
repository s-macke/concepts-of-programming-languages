package main

import "fmt"

type Point struct { x, y int }
func (p Point) String() string { return fmt.Sprintf("(%d, %d)", p.x, p.y) }

type ColorPoint struct { Point; color int }
func (cp ColorPoint) String() string { return fmt.Sprintf("(%d, %d, %d)", cp.x, cp.y, cp.color) }

func main() {
    p := Point{1, 2}
    cp := ColorPoint{Point{1, 2}, 3}
    var s fmt.Stringer = p; fmt.Println(s)  // prints "(1, 2)"
    s = cp; fmt.Println(s)  // prints "(1, 2, 3)"
}
