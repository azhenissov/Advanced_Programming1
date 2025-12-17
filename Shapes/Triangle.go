package Shapes

import "math"

type Triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

func NewTriangle(side1, side2, side3 float64) *Triangle {
	return &Triangle{side1, side2, side3}
}

func (t *Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}

// Area By Heron's Formula
func (t *Triangle) Area() float64 {
	half := t.Perimeter() / 2.0
	return math.Sqrt(half * (half - t.side1) * (half - t.side2) * (half - t.side3))
}
