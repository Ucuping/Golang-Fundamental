package main

import "fmt"

func main() {
	// map
	var language map[string]int
	language = map[string]int{}

	language["test"] = 10
	language["test1"] = 10
	language["test2"] = 10

	fmt.Println(language)
	fmt.Println(language["test"])

	myMap := map[string]string{
		"Go":         "Powerful",
		"JavaScript": "Beautiful",
		"python":     "Best",
	}

	for key, value := range myMap {

		fmt.Println("key: ", key, " value: ", value)
	}

	fmt.Println("===============")

	delete(myMap, "JavaScript")

	for key, value := range myMap {

		fmt.Println("key: ", key, " value: ", value)
	}

	value, isAvailable := myMap["php"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
