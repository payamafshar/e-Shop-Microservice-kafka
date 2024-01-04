package main

import (
	"context"
	"fmt"
	"listnerApp/cmd/reader"
	"log"
	"net/http"
	"os/user"

	echo "github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type User struct {
	Name     string `json:"first_name" `
	Email    string `json:"email" `
	LastName string `json:"last_name" `
}
type Writer[T any] struct {
	w *kafka.Writer
}

func main() {

	e := echo.New()
	e.POST("/auth/test", func(c echo.Context) error {
		writer, err := reader.NewWriter[User]("kafka:9092", "myReader")
		u := new(User)
		writer.WriteBatch(context.TODO())
		reader, err := reader.NewReader[User]("kafka:9092", "myReader", "stringfy", func(u User) {
			loginUser, err := user.Current()
			if err != nil {
				fmt.Println("error, retrying ...")
				writer.WriteBatch(context.TODO(), u)

				return c.String(http.StatusOK, "ERR FROM rEADER")

			}

			fmt.Println(loginUser)
		})
		log.Println(&reader)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, u)

		return
	})
	e.Logger.Fatal(e.Start("0.0.0.0:5004"))
}
