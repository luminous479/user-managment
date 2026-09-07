package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() {
	fmt.Printf("User #%d: %s <%s>\n", u.ID, u.Name, u.Email)
}

type UserRepository interface {
	CreateUser(name string, email string) (*User, error)
	ListUsers()
	GetUserByID(id int) (*User, error)
	UpdateUser(id int, name string, email string) error
	DeleteUser(id int) error
}

type UserService struct {
	users  []User
	nextId int
}

func (ser *UserService) CreateUser(name string, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}

	if email == "" {
		return nil, fmt.Errorf("Email cannot be empty")
	}
	ser.nextId++
	newUser := User{
		ID:    ser.nextId,
		Name:  name,
		Email: email,
	}

	ser.users = append(ser.users, newUser)
	return &ser.users[len(ser.users)-1], nil
}

func (ser *UserService) ListUsers() {
	for _, user := range ser.users {
		user.Display()
	}
}
func (ser *UserService) GetUserByID(id int) (*User, error) {
	for i := range ser.users {
		if ser.users[i].ID == id {
			return &ser.users[i], nil
		}
	}
	return nil, fmt.Errorf("User with ID %d not found", id)
}
func (ser *UserService) UpdateUser(id int, name string, email string) error {
	for i := range ser.users {
		if ser.users[i].ID == id {
			ser.users[i].Name = name
			ser.users[i].Email = email
			return nil
		}
	}
	return fmt.Errorf("User with ID %d not found", id)
}
func (ser *UserService) DeleteUser(id int) error {
	for i := range ser.users {
		if ser.users[i].ID == id {
			ser.users = append(ser.users[:i], ser.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID %d not found", id)
}

func main() {

	var repo UserRepository = &UserService{}

	_, err := repo.CreateUser("Elara", "elara@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	_, err = repo.CreateUser("Ada", "ada@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	_, err = repo.CreateUser("Luna", "luna@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	repo.ListUsers()
	err = repo.UpdateUser(1, "Elara Updated", "elara.updated@gmail.com")

	if err != nil {
		fmt.Println(err)
	}
	getUser, err := repo.GetUserByID(2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found:")
		getUser.Display()
	}

	err = repo.DeleteUser(2)

	if err != nil {
		fmt.Println(err)
	}

}
