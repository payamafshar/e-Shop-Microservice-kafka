package main

import (
	"fmt"
	"gateway-service/cmd"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const webPort = "5051"

type Config struct {
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println("conncted to brooker service")
	err = cmd.SetupServer(PORT)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("yesssssssssssssssssssssss")

}
