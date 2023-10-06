package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 3
	var f float64 = 3.2
	var x = float64(i)
	var y = int(f)

	fmt.Printf("x is of type %T : %f and y of type %T : %d\n", x, x, y, y)

	var i1 = 3
	var f1 = 3.2
	var s1, s2 = "3.14", "5"

	x1 := string(i1)
	_, _, _ = x1, f1, s1

	is, err := strconv.Atoi(s2)

	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}

}
