package main

import "fmt"

type User struct {
	Id       string
	Username string
	Fullname string
	Country  string
	Age      int8
}

// Get user age
func getUserAge(user User) string {
	return fmt.Sprintf("Age is : %d \n", user.Age)
}

func main() {
	user := User{Id: "123", Username: "Rijan", Fullname: "RI Rijan", Country: "BD", Age: 16}

	userAge := getUserAge(user)
	fmt.Print(userAge)
}
