package main

import "fergo.dev/H9/shapes"

type Octagon struct{}

func (t *Octagon) NumberOfSides() int {
	return 8
}

func (t *Octagon) Type() string {
	return "octagon"
}

func countSides(shapes []shapes.GeometricShape) int {
	count := 0

	for _, shape := range shapes {
		count += shape.NumberOfSides()
	}
	
	return count
}
