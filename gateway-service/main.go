package main

import (
	"gateway-service/cmd/api/config"
	"log"
	"os"
	"strconv"
)

func main() {
	PORT := os.Args[1]
	port, err := strconv.Atoi(PORT)
	err = config.SetupServer(port)
	if err != nil {
		log.Panic(err)
	}

}
