package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 5051

type Config struct{}

func main() {
	fmt.Println("hello from gateway")
	app := Config{}

	log.Printf("starting broker service on port %s")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
