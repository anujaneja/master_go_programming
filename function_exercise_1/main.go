package main

import (
	"fmt"
	"math"
	"strconv"
)

func cube(x float64) float64 {
	return math.Pow(x, 3)
}

func f1(n uint) (uint, uint) {
	result := fact(n)
	sum := uint((n * (n + 1)) / 2)

	return result, sum
}

func fact(n uint) uint {
	if n == 0 {
		return uint(1)
	} else {
		return n * fact(n-1)
	}
}

func myFunc(n rune) int64 {
	num1, _ := strconv.ParseInt(string(n), 0, 64)
	num2, _ := strconv.ParseInt(string(n)+string(n), 0, 64)
	num3, _ := strconv.ParseInt(string(n)+string(n)+string(n), 0, 64)

	return num1 + num2 + num3
}

func sum(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return sum
}

func sum_naked_return(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func main() {
	result := cube(3)
	fmt.Println(result)

	fmt.Println(f1(1))

	f, s := f1(2)

	fmt.Printf("factorial: %v , sum: %v\n", f, s)

	f, s = f1(6)

	fmt.Printf("factorial: %v , sum: %v\n", f, s)

	fmt.Println(myFunc('9'))

	fmt.Println("sum is", sum(1, 2, 3, 4, 5))

	fmt.Println("sum_naked_return is", sum_naked_return(1, 2, 3, 4, 5))
}
