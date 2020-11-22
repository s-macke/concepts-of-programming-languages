package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var pointX = flag.Float64("x", -2.0, "X coordinate of starting point of mandelbrot or fix point for Julia (range: 2.0 to 2.0)")
var pointY = flag.Float64("y", -2.0, "Y coordinate of starting point of mandelbrot or fix point for Julia (range: 2.0 to 2.0)")
var zoom = flag.Float64("z", 1.0, "Zoom level (only working properly for mandelbrot)")
var julia = flag.Bool("julia", false, "Turn on Julia calculation")
var maxIter = flag.Int("maxIter", 51, "Max number of point iterations")
var imgSize = flag.Int("imgSize", 500, "Size of the image")

func main() {
	flag.Parse()

	fmt.Printf("X: %f\n", *pointX)
	fmt.Printf("Y: %f\n", *pointY)
	fmt.Printf("Zoom: %f\n", *zoom)
	fmt.Printf("Julia: %t\n", *julia)
	fmt.Printf("MaxIter: %d\n", *maxIter)
	fmt.Printf("ImgSize: %d\n", *imgSize)

	img := CalculateImage(*imgSize, *imgSize)
	WriteImage(img)
}

func CalculateImage(imgWidth int, imgHeight int) *image.NRGBA {
	var img = image.NewNRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	minCx := -2.0
	minCy := -2.0
	if !*julia {
		minCx = *pointX
		minCy = *pointY
	}
	maxSquAbs := 4.0 // maximum square of the absolute value
	// calculate step widths
	stepX := math.Abs(minCx-2.0) / float64(imgWidth) / *zoom
	stepY := math.Abs(minCy-2.0) / float64(imgHeight) / *zoom
	cx := 0.0
	cy := 0.0
	for px := 0; px < imgWidth; px++ {
		cx = minCx + float64(px)*stepX

		for py := 0; py < imgHeight; py++ {
			cy = minCy + float64(py)*stepY

			iterValue := PointIteration(cx, cy, maxSquAbs, *maxIter)

			c := ChooseColor(iterValue, *maxIter)
			img.Set(px, py, c)
		}
	}
	return img
}

func PointIteration(cx float64, cy float64, maxSquAbs float64, maxIter int) int {
	squAbs := 0.0
	iter := 0
	x := 0.0
	y := 0.0
	if *julia {
		x = cx
		y = cy
		cx = *pointX
		cy = *pointY
	}

	for squAbs <= maxSquAbs && iter < maxIter {
		xt := (x * x) - (y * y) + cx // z^2
		yt := (2.0 * x * y) + cy     // z^2
		x = xt
		y = yt
		iter++
		squAbs = (x * x) + (y * y)
	}
	return iter
}

func ChooseColor(iterValue int, maxIter int) *color.RGBA {
	val := uint8(iterValue)
	if iterValue == maxIter {
		return &color.RGBA{A: 0xff}
	}
	multi := uint8(255 / maxIter)
	return &color.RGBA{G: val * multi, A: 0xff}
}

func WriteImage(img *image.NRGBA) {
	file, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Printf("Could not create file %s", file.Name())
	}
	writer := bufio.NewWriter(file)
	png.Encode(writer, img)
	writer.Flush()
	file.Close()
}
