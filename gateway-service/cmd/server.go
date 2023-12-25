package cmd

import (
	"fmt"
	"gateway-service/cmd/api"
	"net/http"
)

const webPort = "5051"

func SetupServer(port int) error {
	mux := api.ApplicationRouter()

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux)
	fmt.Println("connected to server with port", port)
	if err != nil {
		return err
	}
	return nil
}
