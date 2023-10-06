package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x is %d y is %f z is %q\n", x, y, z)

	fmt.Printf("Type of y is %T and Type of score is %T\n", y, score)

	const x1 float64 = 1.422349587101

	fmt.Printf("x1= %.4f\n", x1)

	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape
}
