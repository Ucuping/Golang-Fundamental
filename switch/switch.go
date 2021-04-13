package main

import "fmt"

func main() {
	// switch case condition
	number := 20

	switch number {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	}

	if number == 10 {
		fmt.Println("Number is ", number)
	}
}
