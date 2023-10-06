package main

import "fmt"

type duration int
type mile float64
type km float64

func main() {

	// var hour duration
	// _ = hour

	// fmt.Println(hour)

	// fmt.Printf("hour type is: %T\n", hour)

	// hour = 3600

	// fmt.Println(hour)

	var hour duration = 3600
	minute := 60
	fmt.Println(int(hour) != minute)

	const m2km = 1.609

	var mileBerlinToParis mile = 655.3

	var kmBerlinToParis km = km(mileBerlinToParis * m2km)

	fmt.Printf("kmBerlinToParis = %v", kmBerlinToParis)
}
