package main

import "fmt"

func main() {
	var m1 map[float64]bool
	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1)

	m2 := map[int]string{1: "v1", 2: "v2"}
	m2[10] = "Abba"

	fmt.Printf(" m2 type: %T, m2 value: %#v\n", m2, m2)

	fmt.Println(m2[1])
	fmt.Println(m2[20])

	m12 := map[int]bool{}
	m12[5] = true

	m21 := map[int]int{3: 10, 4: 40}
	m31 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m21 == m31) invalid operation maps can't be compared!

	_, _ = m21, m31

	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 1)

	for k, v := range m {
		fmt.Printf("k : %v , v: %v\n", k, v)
	}

}
