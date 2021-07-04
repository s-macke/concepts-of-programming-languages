package geoobject

import "image/color"

// Position is the position of the screen.
type Position struct {
	x, y float64
}

// GeoObject is the "BaseClass" of any geometrical object. All objects have a position and a color.
type GeoObject struct {
	pos   Position
	color color.Color
}

// Circle is a concrete GeoObject with a radius.
type Circle struct {
	GeoObject
	radius float64
}

// Rectangle is a concrete GeoObject with a width and a height.
type Rectangle struct {
	GeoObject
	width, height float64
}

// Triangle is a concrete GeoObject with three points (ABC).
// The coordinates of the three points are relative to the position of the object.
type Triangle struct {
	GeoObject
	b, c Position
}
