package main

import (
	"errors"
	"fmt"
)

func main() {
	// function
	// input, proses, output
	sentence := printMyFunc("Saya mencoba")
	fmt.Println(sentence)
	// function multi return
	luas, panjang := calculate(10, 20)
	fmt.Println(luas)
	fmt.Println(panjang)
	// _, panjang := calculate(10, 20)
	// fmt.Println(luas)
	// fmt.Println(panjang)
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)
	total, err := calculation(10, 20, "+")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(total)
}

// input, proses, output
func printMyFunc(sentence string) string {
	newSentence := sentence + "Testting"
	return newSentence
}

// function multi return
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang * lebar)

	return luas, keliling
}

// function predevine return value
// func calculate(panjang, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang * lebar)

// 	return
// }

func sum(scores []int) int {
	var total int
	for _, score := range scores {
		total += score

	}
	return total
}

func calculation(number, numberTwo int, operation string) (int, error) {
	var total int
	var myError error

	switch operation {
	case "+":
		total = number + numberTwo
	case "-":
		total = number - numberTwo
	case "*":
		total = number * numberTwo
	case "/":
		total = number / numberTwo
	default:
		myError = errors.New("Unknown operation")
	}
	return total, myError
}
