package main

import "fmt"

var taskDone bool // ok No need to intialized or use it.

func main() {
	// you must provide a type for each variable you declare or Go should be able to infer it
	var a int = 10
	var b = 15.5  // type inference(deduction)
	c := "Gopher" // short declaration type inference

	_, _, _ = a, b, c

	fmt.Println(a, b, c)

	// Go is a Statically and Strong Typed Programming Language
	// a = 3.14  // -> error. A variable cannot change it's type
	// a = b    // -> error. It's not allowed to assign a type to another type

	//** ZERO VALUES **//

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                         // initialized with 0
	var price float64                     // initialized with 0.0
	var name string                       // initialized with empty string -> ""
	var done bool                         // initialized with false
	fmt.Println(value, price, name, done) // -> 0 0.0 ""  false

	//  variable conventions
	var mv int        //max value
	var max_value int // NOT ok in Go

	var packetsReceived int //Not ok , name too long
	var b1 int              //instead ok

	const MAX_VALUE = 100 //Not ok
	const N = 100         // Better

	_, _, _, _ = mv, max_value, packetsReceived, b1

	// Camel Case should be use to create variables.
	maxValue := 1000   // Recommended
	max_values := 1000 // Not Recommended

	writeToDB := true // Ok, idiomatic

	_, _, _ = maxValue, max_value, writeToDB

}
