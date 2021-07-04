package geoobject

import (
	"fmt"
	"image/color"
)

// Painter is used to paint GeoObjects.
type Painter interface {
	Paint()
}

// Paint is implemented by Circle
func (c Circle) Paint() {
	fmt.Printf("Painting circle with radius=%v at position=%v and color=%v\n", c.radius, c.pos, c.color)
}

// Paint is implemented by Rectangle
func (r Rectangle) Paint() {
	fmt.Printf("Painting rectangle with width=%v, height=%v at position=%v and color=%v\n", r.width, r.height, r.pos, r.color)
}

// Paint is implemented by Triangle
func (t Triangle) Paint() {
	fmt.Printf("Painting triangle with pos=%v, b=%v, c=%v and color=%v\n", t.pos, t.b, t.c, t.color)
}

func Paint() {
	// Polymorph slice of Painter objects
	objects := []Painter{
		// short initialization
		Circle{GeoObject{Position{1, 2}, color.Black}, 40},
		Rectangle{GeoObject{Position{1, 2}, color.Black}, 10, 10},
		Triangle{GeoObject{Position{10, 20}, color.Black}, Position{11, 21}, Position{12, 22}},

		// or with named identifiers
		Circle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: color.Black},
			radius: 40},
		Rectangle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: color.Black},
			width:  10,
			height: 10},
		Triangle{
			GeoObject: GeoObject{
				pos:   Position{x: 10, y: 20},
				color: color.Black},
			b: Position{x: 11, y: 21},
			c: Position{x: 12, y: 22}},
	}
	for _, v := range objects {
		v.Paint()
	}
}
