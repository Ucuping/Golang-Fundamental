package main

import "fmt"

func main() {
	//map slice
	sudents := []map[string]string{
		{"name": "Ucup", "score": "A"},
		{"name": "Praset", "score": "B"},
		{"name": "Mer", "score": "B"},
	}

	fmt.Println(sudents)

	// test
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// var total int

	// for _, score := range scores {
	// 	total += score
	// }

	// length := len(scores)
	// average := float64(total) / float64(length)

	// fmt.Println(average)

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var value []int

	for _, val := range scores {
		if val >= 90 {
			value = append(value, val)
		}
	}

	fmt.Println(value)
}
