package repo

import (

	model "github.com/luminous479/user-managment/internal/model"
)

type UserRepository interface {
	CreateUser(name string, email string) (*model.User, error)
	ListUsers()
	GetUserByID(id int) (*model.User, error)
	UpdateUser(id int, name string, email string) error
	DeleteUser(id int) error
}
