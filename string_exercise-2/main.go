package main

import "fmt"

func main() {
	var name = "Anuj Aneja"

	country := "India"

	fmt.Printf(`Your name: %s
Country: %s`, name, country)
	fmt.Println()
	fmt.Println(`He says: "Hello"`)
	fmt.Println(`C:\User\a.txt`)
	fmt.Println("C:\\User\\a.txt")

	var r rune = 'ă'
	fmt.Printf("r of type: %T\n", r)

	ma, m := "ma", "m"

	fmt.Println(ma + m + string(r))
	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	s2 := "Go is cool!"
	x := s2[0]
	fmt.Println(x)

	// s2[0] = 'x'

	// printing the number of runes in the string
	fmt.Println(len(s2))

	s := "你好 Go!"
	byteSlice := []byte(s)
	fmt.Printf("byteSlice: %#v\n", byteSlice)

	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("index: %d, byte: %v\n", i, byteSlice[i])
	}

	s = "你好 Go!"
	runSlice := []rune(s)
	fmt.Printf("rune slice: %#v\n", runSlice)

	for i, r := range runSlice {
		fmt.Printf("index: %d, r: %c\n", i, r)
	}
}
