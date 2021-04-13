package main

import "fmt"

func main() {
	// if else condition
	score := 20
	var grade string

	if score <= 20 {
		grade = "You score 20-"
	} else if score >= 20 {
		grade = "You score 20+"
	} else {
		grade = "Nothing"
	}

	fmt.Println(grade)
}
