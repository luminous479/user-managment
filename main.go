package main

import "fmt"


type User struct{

	ID int
	Name string
	Email string
}

func (u User) Display() {
	fmt.Printf("User #%d: %s <%s>\n", u.ID, u.Name, u.Email)
};

type UserRepository interface{

	CreateUser(name string, email string)
	ListUsers()
	GetUserByID(id int) *User
}

type UserService struct{
	users []User
}

func (ser *UserService)  CreateUser(name string, email string)  {
	newUser := User{
		ID: len(ser.users) +1,
		Name: name,
		Email: email,
	}
	ser.users = append(ser.users, newUser )
}

func (ser *UserService) ListUsers() {
	for _, user := range ser.users {
		user.Display()
	}
}
func (ser *UserService) GetUserByID(id int) *User {
	for i := range ser.users {
		if ser.users[i].ID == id {
			return &ser.users[i]
		}
	}
	return nil 
}

 
func main(){

	var repo UserRepository= &UserService{}

	repo.CreateUser("Elara","elara@gmail.com")
	repo.CreateUser("Ada","ada@gmail.com")
	repo.ListUsers()

	if repo.GetUserByID(2) == nil {
		fmt.Println("Not Found")
	}else{
		fmt.Println("Found")
	}
}