package main

import (
	"fmt"
)

type User struct {
	Name     string `json:"first_name" `
	Email    string `json:"email" `
	LastName string `json:"last_name" `
}

func main() {
	var user User
	fmt.Println("asdasdasd", user.Name)
}
