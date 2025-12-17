package main

import (
	"fmt"

	"github.com/azhenissov/Advanced_Programming1/Shapes"
)

func main() {
	shapes := []Shapes.Shape{
		Shapes.NewRectangle(4, 5),
		Shapes.NewCircle(3),
		Shapes.NewSquare(5),
		Shapes.NewTriangle(3, 4, 5),
	}

	for i, shape := range shapes {
		fmt.Printf("Shape: %d\n", i+1)
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n\n", shape.Perimeter())
	}
}
