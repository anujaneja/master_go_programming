package main

import (
	"fmt"
)

var version = "3.1"

func main() {
	/*
		//Coding exercise #1
		var a int
		var b float64
		var c bool
		var d string

		x := 20
		y := 15.5
		z := "Gopher!" // https://blog.golang.org/gopher

		fmt.Println(a, b, c, d)
		fmt.Println("x is", x, "y is", y, "z is", z)



		//Coding exercise #2
		var (
			a int
			b float64
			c bool
			d string
		)

		x, y, z := 20, 15.5, "Gopher!"

		fmt.Println(a, b, c, d, x, y, z)
	*/

	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y

	name := "Golang"
	fmt.Println(name)

	const daysOfWeek int = 7
	const lightOfSpeed float64 = 299792458
	const pi float64 = 3.14159

	const secPerDay int64 = 24 * 60 * 60
	const daysInYear int64 = 365

	fmt.Println(secPerDay * daysInYear)

	const x1 int = 10

	// declaring a constant of type slice int ([]int)
	//const m = []int{1: 3, 4: 5, 6: 8}
	//_ = m

	const a1 = 7
	const b float64 = 5.6
	const c = a1 * b

	// x2 := 8
	// const xc int = x2

	// const noIPv6 = math.Pow(2, 128)

	const (
		jun = iota + 6
		july
		aug
	)

	fmt.Println(jun, july, aug)
}
