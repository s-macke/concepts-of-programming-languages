package geoobject

import (
	"github.com/fogleman/gg"
	"image/color"
)

type GGPainter interface {
	GGPaint(ctx *gg.Context)
}

func (c Circle) GGPaint(ctx *gg.Context) {
	ctx.SetColor(c.color)
	ctx.DrawCircle(c.pos.x, c.pos.y, c.radius)
	ctx.Stroke()
}

func (r Rectangle) GGPaint(ctx *gg.Context) {
	ctx.SetColor(r.color)
	ctx.DrawRectangle(r.pos.x, r.pos.y, r.width, r.height)
	ctx.Stroke()
}

func (t Triangle) GGPaint(ctx *gg.Context) {
	ctx.SetColor(t.color)
	ctx.DrawLine(t.pos.x, t.pos.y, t.b.x, t.b.y)
	ctx.DrawLine(t.b.x, t.b.y, t.c.x, t.c.y)
	ctx.DrawLine(t.c.x, t.c.y, t.pos.x, t.pos.y)
	ctx.Stroke()
}

func GGPaint() {
	context := gg.NewContext(100, 100)
	context.SetLineWidth(3)

	objects := []GGPainter{
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
				pos:   Position{x: 10, y: 5},
				color: color.RGBA{B: 255, A: 255}},
			b: Position{x: 50, y: 5},
			c: Position{x: 50, y: 35}},
	}

	for _, v := range objects {
		v.GGPaint(context)
	}

	err := context.SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}
