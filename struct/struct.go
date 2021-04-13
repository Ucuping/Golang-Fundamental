package main

import "fmt"

func main() {
	//struct
	// user := User{}
	// user.ID = 1
	// user.Email = "ucup@gmail.com"
	// user := User{ID: 1}
	user := User{1, "ucup", "ping", "ucup@gmail.com", true}
	result := printUser(user)
	fmt.Println(result)

	// embedded struct
	users := []User{user}
	group := Group{"Programer", user, users, true}

	printGroup(group)
}

// struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsAvtive  bool
}

// struct
func printUser(user User) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

// embedded struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func printGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Group: %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Member Name:")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
