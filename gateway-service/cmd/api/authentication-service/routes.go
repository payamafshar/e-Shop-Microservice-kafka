package authenticationservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(group *gin.RouterGroup) {

	authRoutes := group.Group("auth")
	// authRoutes.GET("/", handler)
	authRoutes.POST("/test", Testhandler)

}

// func handler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		msg := kafka.Message{
// 			Key:   []byte("FromGetRoute"),
// 			Value: []byte("helllo"),
// 		}

//			if err != nil {
//				ctx.JSON(http.StatusBadGateway, ([]byte(err.Error())))
//				log.Fatalln(err)
//			}
//		}
//	}
func Testhandler(ctx *gin.Context) {

	var createAccountDto CreateAccountDto
	if err := ctx.ShouldBindJSON(&createAccountDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	writer, closeWriter := NewWriter[CreateAccountDto]("kafka:9092", "twitter.newTweets")
	err := writer.WriteBatch(ctx, createAccountDto)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
	}
	ctx.JSON(http.StatusOK, &createAccountDto)
	defer closeWriter()
}

type CreateAccountDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// in this gateway trying to send data if we can get this particiluar message i mean addres And Hello its a key point
