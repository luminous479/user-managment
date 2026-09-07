package main

import (
	"fmt"
	service "github.com/luminous479/user-managment/internal/service"
	repositiory "github.com/luminous479/user-managment/internal/repo"
)

func main() {

	var repo repositiory.UserRepository = &service.UserService{}

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
