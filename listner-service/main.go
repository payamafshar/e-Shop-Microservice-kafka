package main

import (
	"fmt"
	"os"
)

type User struct {
	Name     string `json:"first_name" `
	Email    string `json:"email" `
	LastName string `json:"last_name" `
}

func main() {
	PORT := os.Args[1]
	fmt.Println(PORT)
	fmt.Println("PORT IS FROM CMDARGUMNT", PORT)
}
