package main

import "fmt"

func main() {
	// slice
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "PS4")
	gamingConsole = append(gamingConsole, "PS3")
	gamingConsole = append(gamingConsole, "PS2")

	// gamingConsole := []string{"PS4", "PS3", "PS2"}

	for _, console := range gamingConsole {

		fmt.Println(console)
	}
}
