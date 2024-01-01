package authenticationservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

func SetupAuthRoutes(group *gin.RouterGroup, kafkaWriter *kafka.Writer) {
	authRoutes := group.Group("auth")
	authRoutes.GET("/", handler(kafkaWriter))
	authRoutes.POST("/test", Testhandler(kafkaWriter))

}

func handler(kafkaWriter *kafka.Writer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		msg := kafka.Message{
			Key:   []byte("FromGetRoute"),
			Value: []byte("helllo"),
		}
		err := kafkaWriter.WriteMessages(ctx.Request.Context(), msg)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, ([]byte(err.Error())))
			log.Fatalln(err)
		}
	}
}
func Testhandler(kafkaWriter *kafka.Writer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto := new(CreateAccountDto)
		if err := ctx.ShouldBindJSON(dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg := kafka.Message{
			Key:   []byte("FromPostRoute"),
			Value: []byte(dto.Data),
		}
		fmt.Println(dto.Data)
		err := kafkaWriter.WriteMessages(ctx.Request.Context(), msg)
		ctx.JSON(http.StatusAccepted, &dto)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, ([]byte(err.Error())))
			log.Fatalln(err)
		}
	}
}

type CreateAccountDto struct {
	Data string `json:"data"`
}

func getKafkaTestWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// in this gateway trying to send data if we can get this particiluar message i mean addres And Hello its a key point
