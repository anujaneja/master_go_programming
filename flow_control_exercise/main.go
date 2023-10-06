package main

import "fmt"

func main() {
	//exercise 1
	// Print num divisble by 7 between 1 to 50
	num := 1

	for num <= 50 {
		if num%7 == 0 {
			fmt.Println(num)
		}

		num++
	}

	//exercise 2
	num = 1
	fmt.Println("num after exercise 1: ", num)
	for num <= 50 {
		if num%7 != 0 {
			num++
			continue
		} else {
			fmt.Println(num)
			num++
		}

	}

	//exercise 3
	count := 0
	num = 1
	fmt.Println("num after exercise 2: ", num)
	for num <= 50 {
		if num%7 == 0 {
			fmt.Println(num)
			count++
		}

		if count == 3 {
			break
		}
		num++
	}

	num = 1
	fmt.Println("Number divisible by both 5 and 7 are: ")
	for num <= 500 {
		if num%5 == 0 && num%7 == 0 {
			fmt.Println(num)
		}
		num++
	}

	fmt.Println("Years from birth year")
	birthYear, currentYear := 1986, 2023

	for i := birthYear; i <= currentYear; i++ {
		fmt.Println(i)
	}

	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age < 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}

	//switch version
	age = 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")

	}
}
