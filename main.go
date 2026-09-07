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
	UpdateUser(id int, name string, email string)
	DeleteUser(id int)

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
func (ser *UserService) UpdateUser(id int, name string, email string) {
	for i := range ser.users {
		if ser.users[i].ID == id {
			ser.users[i].Name = name
			ser.users[i].Email = email
			return
		}
	}
}
func (ser *UserService) DeleteUser(id int) {
	for i := range ser.users {
		if ser.users[i].ID == id {
			ser.users = append(ser.users[:i], ser.users[i+1:]...)
			return
		}
	}
}
 
func main(){

	var repo UserRepository= &UserService{}

	repo.CreateUser("Elara","elara@gmail.com")
	repo.CreateUser("Ada","ada@gmail.com")
	repo.CreateUser("Luna","luna@gmail.com")
	repo.ListUsers()
    repo.UpdateUser(1,"Elara Updated","elara.updated@gmail.com")
	if repo.GetUserByID(2) == nil {
		fmt.Println("Not Found")
	}else{
		fmt.Println("Found")
	}

    repo.DeleteUser(2)

}