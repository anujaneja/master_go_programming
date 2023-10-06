package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	//exercise #1

	s1 := []string{"A", "B", "C"}

	for index, value := range s1 {
		fmt.Printf("index %d, value %s \n", index, value)
	}

	/*
		exercise #2
	*/
	mySlice := []float64{1.2, 5.6}

	//mySlice[2] = 6
	mySlice[1] = 6

	a := 10
	mySlice[0] = float64(a)

	mySlice[1] = 10.10

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

	/*
		exercise #3
	*/
	nums := []float64{1.1, 2.1, 3.1}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println("nums ", nums)

	n := []float64{5.1, 6.1}

	nums = append(nums, n...)

	fmt.Println("nums ", nums)

	/*
		exercise #4
	*/
	if len(os.Args) < 2 {
		log.Fatal("Please run with arguments (2-10 numbers)")
	}

	args := os.Args[1:]
	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue // can't convert string to float64...continue to next number
			}
			sum += num
			product *= num
		}
		fmt.Printf("Sum: %v Product: %v\n", sum, product)
	}

	/*
		Exercise #5
	*/
	nums1 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum1 := 0
	for index, value := range nums1[1 : len(nums1)-1] {
		fmt.Printf("index %d value %v\n", index, value)
		sum1 += value
	}
	fmt.Println("sum1 ", sum1)

	/*
		Exercise #6
	*/
	friends := []string{"Marry", "John", "Paul", "Diana"}

	friends1 := make([]string, len(friends))

	copy(friends1, friends)

	friends[0] = "Larry"

	fmt.Printf("friends %v\n", friends)
	fmt.Printf("friends1 %v\n", friends1)

	/*
		Exercise #7
	*/

	myfriends := []string{"Marry", "John", "Paul", "Diana"}
	myfriends1 := append(myfriends, "Larry", "Johna")

	myfriends[0] = "Harry"

	fmt.Printf("myfriends %v\n", myfriends)
	fmt.Printf("myfriends1 %v\n", myfriends1)

	/* Exercise #8
	 */
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := append(years[0:3], years[len(years)-3:len(years)]...)

	fmt.Println("newYears ", newYears)

}
