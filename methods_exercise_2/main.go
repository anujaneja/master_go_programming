package main

import "fmt"

type book struct {
	title string
	price float64
}

// method for book type
func (b *book) changePrice(p float64) {
	b.price = p
}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change

	var v vehicle = car{licenceNo: "1234", brand: "Toyota Fortuner"}
	fmt.Printf("car : License : %s and Name is: %s\n", v.License(), v.Name())

	//***************************************************************************//

	var empty interface{}
	fmt.Printf("Type: %T\n", empty)

	empty = 10
	fmt.Printf("Type: %T\n", empty)

	empty = 10.5
	fmt.Printf("Type: %T\n", empty)

	empty = []int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T\n", empty)

	empty = append(empty.([]int), 6)

	fmt.Printf("%v\n", empty)

	var x interface{}
	x = cube{edge: 5}
	v1 := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v1)
}
