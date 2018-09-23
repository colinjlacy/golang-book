package main

import (
	"math"
)

type Point struct { X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

func (point Point) Distance(q Point) float64 {
	return math.Hypot(point.X - q.X, point.Y - q.Y)
}

type Path []Point

func (path Path) Distance() (sum float64) {
	sum = 0.0
	for i, point := range path {
		if i > 0 {
			sum += path[i - 1].Distance(point)
		}
	}
	return
}

func (point *Point) ScaleBy(factor float64) {
	point.X *= factor
	point.Y *= factor
}

//func main() {
//	perim := Path{
//		{1,1},
//		{5,1},
//		{5,4},
//		{1,1},
//	}
//
//	fmt.Println(perim.Distance())
//}