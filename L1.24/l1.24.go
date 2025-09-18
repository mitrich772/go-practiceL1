package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt((dx * dx) + (dy * dy))
}

func main(){
	p1 := NewPoint(0, 0)
	p2 := NewPoint(4, 3)
	fmt.Println(p1.Distance(p2))
}
