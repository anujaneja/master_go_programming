package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("info.txt")

	if err != nil {
		log.Fatal("err: ", err)
	}
	defer file.Close()

	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal("err: ", err)
	}

	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal("err: ", err)
	}

	file1, err := os.Open("info-1.txt")

	// content, err := io.ReadAll(file1)

	// _ = content

	if err != nil {
		log.Fatal("Err: ", err)
	}

	// fmt.Printf("Data as string %s\n", content)
	defer file1.Close()

	scanner := bufio.NewScanner(file1)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//checking for possible error

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bs := []byte("The Go gopher is an iconic mascot!")

	err = ioutil.WriteFile("info.txt", bs, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// f1.WriteString("The Go gopher is an iconic mascot!")
}
