package authenticationservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	secureURL := DownloadHandler(ctx)
	fmt.Println(secureURL)
	createAccountDto.url = secureURL
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
	url       string `json:"url"`
}

func DownloadHandler(ctx *gin.Context) string {
	ctx.File("/home/helltion/e-Shop-Microservice-kafka/authentication-service/test.tsx")
	http.ServeFile(ctx.Writer, ctx.Request, "/home/helltion/e-Shop-Microservice-kafka/authentication-service/test.tsx")
	fmt.Println("Context From DownloadHandler", "c.Request.UserAgent")
	TEST := ctx.ContentType()
	return TEST

}

// in this gateway trying to send data if we can get this particiluar message i mean addres And Hello its a key point
