package reciver

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type WriteData struct {
	data string
}

func RecivedData() {

	writer, closeWriter := api.NewWriter[WriteData]("kafka:9092", "twitter.newTweets")
	reader, closeReader := api.NewReader[WriteData]("kafka:9092", "twitter.newTweets", "saver", func(incomingData WriteData) {
		// retry process
		fmt.Println("error, retrying ...")
		writer.WriteBatch(context.TODO(), incomingData.data)
	})
	defer closeReader()
	defer closeWriter()
	go reader.Read(func(incomingData WriteData) error {
		if incomingData.data != "" {
			fmt.Println("received a message: ", incomingData.data)
		}
		fmt.Println(incomingData.url)
		if rand.Intn(100) > 50 {
			return errors.New("a random error")
		}
		return nil
	})

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
	fmt.Println("Closing Kafka connections ...")

}
