package main

import "fmt"

func swapValues(x *float64, y *float64) {
	*x = 8.8
	*y = 5.5
}

func main() {
	x := 10.10
	fmt.Println("address of x: ", &x)

	ptr := &x
	fmt.Printf("Type of ptr: %T and value: %v\n", ptr, *ptr)

	fmt.Printf("Address of ptr: %v , value of ptr: %v\n", &ptr, *ptr)

	x1, y1 := 10, 2
	ptrx, ptry := &x1, &y1
	z1 := *ptrx / *ptry
	fmt.Println(z1)

	x2, y2 := 5.5, 8.8
	swapValues(&x2, &y2)

	fmt.Printf("values after swap: x2: %v, y2: %v", x2, y2)
}
