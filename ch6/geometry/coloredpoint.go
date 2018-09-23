package main

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct{
	Point
	Color color.RGBA
}

func main() {
	var p ColoredPoint
	p.X = 1
	p.Y = 2
	p.Color = color.RGBA{255, 0, 0, 255}

	q := ColoredPoint{Point{5,4}, color.RGBA{0, 0, 255, 255}}

	p.ScaleBy(5)

	distance := p.Distance(q.Point)

	fmt.Println(distance)
}