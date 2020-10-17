package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"os"
)

type Position struct {
	x, y float64
}

type GeoObject struct {
	pos   Position
	color color.Color
}

type Circle struct {
	GeoObject
	radius float64
}

type Rectangle struct {
	GeoObject
	width, height float64
}

type Triangle struct {
	GeoObject
	p1, p2, p3 Position
}

type Painter interface {
	Paint(ctx *gg.Context)
}

func (c Circle) Paint(ctx *gg.Context) {
	ctx.SetColor(c.color)
	ctx.DrawCircle(c.pos.x, c.pos.y, c.radius)
	ctx.Stroke()
}

func (r Rectangle) Paint(ctx *gg.Context) {
	ctx.SetColor(r.color)
	ctx.DrawRectangle(r.pos.x, r.pos.y, r.width, r.height)
	ctx.Stroke()
}

func (c Triangle) Paint(ctx *gg.Context) {
	ctx.SetColor(c.color)
	ctx.DrawLine(c.p1.x, c.p1.y, c.p2.x, c.p2.y)
	ctx.DrawLine(c.p2.x, c.p2.y, c.p3.x, c.p3.y)
	ctx.DrawLine(c.p3.x, c.p3.y, c.p1.x, c.p1.y)
	ctx.Stroke()
}

func main() {
	context := gg.NewContext(100, 100)
	context.SetLineWidth(3)

	objects := []Painter{
		// or with named identifiers
		Circle{
			GeoObject: GeoObject{
				pos:   Position{x: 50, y: 50},
				color: color.RGBA{R: 255, A: 255}},
			radius: 40},
		Rectangle{
			GeoObject: GeoObject{
				pos:   Position{x: 60, y: 40},
				color: color.RGBA{G: 255, A: 255}},
			width:  20,
			height: 20},
		Triangle{
			GeoObject: GeoObject{
				pos:   Position{},
				color: color.RGBA{B: 255, A: 255}},
			p1: Position{x: 10, y: 5},
			p2: Position{x: 50, y: 5},
			p3: Position{x: 50, y: 35}},
	}

	for _, v := range objects {
		v.Paint(context)
	}

	err := context.SavePNG("out.png")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
