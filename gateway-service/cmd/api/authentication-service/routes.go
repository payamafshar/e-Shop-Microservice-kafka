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
	authRoutes.POST("/test2", IncomingHandler)

}

// data type of incoming data from DownloadHandler
type WriteData struct {
	data string
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

	if err := ctx.ShouldBindJSON(&incomingData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secureURL := DownloadHandler(ctx)
	fmt.Println(secureURL)
	createAccountDto.url = secureURL
	writer, closeWriter := NewWriter[CreateAccountDto]("kafka:9092", "twitter.newTweets")
	writer1, closeWriter1 := NewWriter[WriteData]("kafka:9092", "twitter.newTweetss")
	err := writer.WriteBatch(ctx, createAccountDto)
	err = writer1.WriteBatch(ctx, incomingData)
	defer closeWriter1()
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

type WriteData struct {
	data string
}

func IncomingHandler(ctx *gin.Context) {
	var incomingData WriteData
	if err := ctx.ShouldBindJSON(&incomingData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secureURL := DownloadHandler(ctx)
	fmt.Println("ylylyl", secureURL)
	writer1, closeWriter1 := NewWriter[WriteData]("kafka:9092", "twitter.newTweetss")
	err := writer1.WriteBatch(ctx, incomingData)
	err = writer1.WriteBatch(ctx, incomingData)
	defer closeWriter1()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
	}
	ctx.JSON(http.StatusOK, &incomingData)
	defer closeWriter1()

}

// in this gateway trying to send data if we can get this particiluar message i mean addres And Hello its a key point
