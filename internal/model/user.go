package model

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() {
	fmt.Printf("User #%d: %s <%s>\n", u.ID, u.Name, u.Email)
}
