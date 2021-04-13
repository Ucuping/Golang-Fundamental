package management

import "fmt"

// exported package
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsAvtive  bool
}

// exported package
func (user User) PrintUser() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}
