package service

import (
	"fmt"

	model "github.com/luminous479/user-managment/internal/model"
)

type UserService struct {
	users  []model.User
	nextId int
}

func (ser *UserService) CreateUser(name string, email string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}

	if email == "" {
		return nil, fmt.Errorf("Email cannot be empty")
	}
	ser.nextId++
	newUser := model.User{
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
func (ser *UserService) GetUserByID(id int) (*model.User, error) {
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
