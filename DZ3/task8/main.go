package main

import "fmt"

func main() {
	shapes := []GeometricShape{
		&Dot{},
		&Henagon{},
		&Digon{},
		&RightTriangle{},
		&RightTriangle{},
		&EquilateralTriangle{},
	}
	fmt.Println(countTriangles(shapes)) // 3
	fmt.Println(countPolygons(shapes))  // 5
}

func countTriangles(shapes []GeometricShape) int {
	count := 0

	for _, shape := range shapes {
		_, ok := shape.(Triangle)
		if ok {
			count++
		}
	}
	return count
}

func countPolygons(shapes []GeometricShape) int {
	count := 0	

	for _, shape := range shapes {
		_, ok := shape.(Polygon)
		if ok {
			count++
		}
	}

	return count
}
