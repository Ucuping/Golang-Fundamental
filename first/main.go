package main

import (
	"first/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// sentence := Test()

	// fmt.Println(sentence)

	result := calculation.Add(10, 20)

	fmt.Println(result)
}
