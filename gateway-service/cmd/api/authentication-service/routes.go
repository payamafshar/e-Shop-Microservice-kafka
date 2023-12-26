package authenticationservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(group *gin.RouterGroup) {
	authRoutes := group.Group("auth")
	authRoutes.GET("/", handler)

}

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"ASDASD": "TYESD"})
}
