package main

import "fmt"

func main() {
	cities := [2]string{"hi", "bye"}
	_ = cities

	fmt.Println(cities)

	var grades = [3]float64{2.3, 4.4}
	_ = grades
	fmt.Println(grades)

	taskDone := [...]bool{true, false}
	fmt.Println(taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, v := range grades {
		fmt.Println("at index ", i, " Value is: ", v)
	}

	nums := [...]int{30, -1, -6, 90, -6}
	positiveIntegerCount := 0
	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			positiveIntegerCount++
		}
	}

	fmt.Println("positiveIntegerCount ", positiveIntegerCount)

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	// myArray[3] = 10.10

	fmt.Println(myArray)
}
