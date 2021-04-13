package main

import "fmt"

func main() {
	// looping
	fmt.Println("I am learning Golang")
	for i := 1; i <= 10; i++ {
		fmt.Println("I am is amazing", i)
	}

	// while
	i := 1
	for i <= 10 {
		fmt.Println("I am amazing", i)
		i++
	}

	// foreach
	title := "Go the best language"
	for index, letter := range title {
		// fmt.Println("index:", index, " letter:", letter)
		fmt.Println("index:", index, " letter:", string(letter))
	}
	for _, letter := range title {
		fmt.Println("Golang:", string(letter))
	}
	for index, letter := range title {
		// if index%2 == 0 {
		// 	fmt.Println("index:", index, " letter:", string(letter))
		// }
		letterString := string(letter)
		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("index:", index, " letter:", string(letter))
		// }
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index:", index, " letter:", string(letter))
		}
	}
}
