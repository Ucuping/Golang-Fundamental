package main

import (
	"fmt"
	"golang-fundamental/calculation"
	"golang-fundamental/management"
)

func main() {
	fmt.Println("Hello World")
	// sentence := Test()

	// fmt.Println(sentence)

	// result := calculation.Add(20, 10)

	// fmt.Println(result)

	increase := calculation.Increase(10, 20)
	multiplication := calculation.Multiplication(5, 10)

	fmt.Println(increase)
	fmt.Println(multiplication)

	// exported package
	user := management.User{1, "ucup", "ping", "ucup@gmail.com", true}
	fmt.Println(user.PrintUser())
}
