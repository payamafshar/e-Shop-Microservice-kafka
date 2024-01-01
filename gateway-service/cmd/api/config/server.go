package config

import (
	"fmt"
	authenticationservice "gateway-service/cmd/api/authentication-service"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SetupServer(PORT int) error {
	mux := gin.Default()
	gin.SetMode(gin.DebugMode)
	kafkaWriter := getKafkaWriter("kafka:9092", "my-topic")
	// mux := routes.ApplicationRouter()
	authenticationservice.SetupAuthRoutes(&mux.RouterGroup, kafkaWriter)
	err := mux.Run(fmt.Sprintf("0.0.0.0:%d", PORT))
	defer kafkaWriter.Close()
	if err != nil {
		return err
	}
	return nil
}
