package config

import (
	"fmt"
	authenticationservice "gateway-service/cmd/api/authentication-service"

	"github.com/gin-gonic/gin"
)

const webPort = "5051"

func SetupServer(PORT int) error {
	mux := gin.Default()
	gin.SetMode(gin.DebugMode)
	// mux := routes.ApplicationRouter()
	authenticationservice.SetupAuthRoutes(&mux.RouterGroup)
	err := mux.Run(fmt.Sprintf(":%d", PORT))

	if err != nil {
		return err
	}
	return nil
}
