package main

import "fmt"

// method
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsAvtive  bool
}

// method
// func printUser(user User) string {
// 	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
// }

// method
func (user User) printUser() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func main() {
	// method
	user := User{1, "ucup", "ping", "ucup@gmail.com", true}
	fmt.Println(user.printUser())
}
