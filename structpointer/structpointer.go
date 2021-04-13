package main

import "fmt"

func main() {
	// struct pointer
	student := Student{1, "Yusup Supriyanto", 362055401025}
	fmt.Println(student.Name)
	change(&student)
	// method struct pointer
	student.change()
	fmt.Println(student.Name)
	game := Gamer{Name: "Game"}
	game.addGame("FIFA2021")
	game.addGame("GTA6")

	for _, g := range game.Game {
		fmt.Println(g)
	}
}

// struct pointer
type Student struct {
	ID   int
	Name string
	NIM  int
}
type Gamer struct {
	Name string
	Game []string
}

// struct pointer
func change(student *Student) {
	student.Name = student.Name + " A.md"
}

// method struct pointer
func (student *Student) change() {
	student.Name = student.Name + " A.md"
}
func (game *Gamer) addGame(name string) {
	game.Game = append(game.Game, name)
}
