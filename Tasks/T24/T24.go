package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Dist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
func main() {
	p1 := New(0, 6)
	p2 := New(8, 0)

	fmt.Printf("p1 has x: %.2f, y: %.2f\n", p1.x, p1.y)
	fmt.Printf("p2 has x: %.2f, y: %.2f\n", p2.x, p2.y)

	fmt.Printf("Distance beetween two points: %.2f\n", Dist(p1, p2))

}
