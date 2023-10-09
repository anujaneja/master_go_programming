package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
	grades grades
}

type grades struct {
	grade  int
	course string
}

func main() {
	me := person{name: "Anuj Aneja", age: 35, colors: []string{"red", "blue", "green"}}
	you := person{name: "Ram", age: 37, colors: []string{"Violet", "Orange", "green"}}

	fmt.Printf("me: %#v\n", me)
	fmt.Printf("me: %#v\n", you)

	me.name = "Andrei"
	colors := []string{"Green", "Black", "Orange"}
	you.colors = append(you.colors, colors...)
	fmt.Printf("me: %#v\n", me)
	fmt.Printf("me: %#v\n", you)

	for _, v := range me.colors {
		fmt.Println(v)
	}

	myPerson := person{
		name:   "Ram Kumar",
		age:    34,
		colors: []string{"red", "blue", "green"},
		grades: grades{
			grade:  98,
			course: "golang",
		},
	}

	fmt.Printf("myPerson: %#v", myPerson)
}
