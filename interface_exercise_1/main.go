package main

import "fmt"

type money float64

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * .09
}

func (b *book) discount() {
	(*b).price = b.price * .9
}

func (m money) print() {
	fmt.Printf("money m is: %.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	var salary money = 4.332
	salary.print()

	fmt.Println(salary.printStr())

	var b book = book{price: 100, title: "My Book"}
	fmt.Println("BEfore discount....\n", b)
	fmt.Println("After discount...\n")
	fmt.Printf("Vat charges is: %f\n", b.vat())
	b.discount()
	fmt.Printf("Price after discount: %v\n", b)

}
