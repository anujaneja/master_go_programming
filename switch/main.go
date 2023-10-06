/////////////////////////////////
// Switch Statement
// Go Playground: https://play.golang.org/p/FI-zCGPMtmA
/////////////////////////////////

// Go converts switch statements into if statements behind the scenes automatically.
// The purpose of a switch statement is to make very long if statements more readable.

package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "Python": //values must be comparable (compare string to string)
		fmt.Println("Learning Python...")
	case "go", "golang":
		fmt.Println("Learning golang...")
	default:
		fmt.Println("Learning other programming language...")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even")
	case n%2 != 0:
		fmt.Println("Odd")
	default:
		fmt.Println("Unreachable!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default

	switch n := 10; { //default bool comparison
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
