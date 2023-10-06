package main

import "fmt"

func main() {
	var age int = 10
	fmt.Println("Age:", age)

	var name = "Anuj"

	fmt.Println("Name is: ", name)

	s := "Sample String"
	fmt.Println(s)

	// Multiple variables declaration and assignment...
	car, cost := "Audi", 50000

	// _,_ = car, cost

	fmt.Println(car, cost)

	// car,cost := "BMW",40000 //Error becuase no new variable defined...

	car, year := "BMW", 40000 // Atleast one new variable defined...

	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	//multiple assignments

	var i, j int
	i, j = 5, 8

	j, i = i, j //swapping two varables...

	fmt.Println(i, j)

	sum := 5 + 2.3

	fmt.Println(sum)

}
