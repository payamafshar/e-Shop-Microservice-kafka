package main

import (
	"gateway-service/cmd/api/config"
	"log"
	"os"
	"strconv"
)

func main() {
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	err = config.SetupServer(PORT)
	if err != nil {
		log.Panic(err)
	}

}
