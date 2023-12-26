package main

import (
	"fmt"
	"gateway-service/cmd/api/config"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(PORT)
	fmt.Println("conncted to brooker service")
	err = config.SetupServer(PORT)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("yesssssssssssssssssssssss")

}
