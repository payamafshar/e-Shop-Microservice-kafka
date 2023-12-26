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
	err := godotenv.Load()
	fmt.Println("conncted to brooker service")
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	err = config.SetupServer(PORT)
	if err != nil {
		log.Panic(err)
	}

}
