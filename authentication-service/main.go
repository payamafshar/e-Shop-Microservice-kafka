package main

import (
	"authentication-service/cmd/api"
	"authentication-service/cmd/reciver"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func ClientOptions() *options.ClientOptions {
	// mongoUrl := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	return &options.ClientOptions{
		Hosts:   []string{"0.0.0.0"},
		AppName: &dbName,
		Dialer:  options.Client().Dialer,
	}
}

func getMongoCollection(mongoURL, dbName, collectionName string) *mongo.Collection {
	// client, err := mongo.Connect(context.Background() )
	clientOptions := ClientOptions()
	var cred options.Credential

	cred.Username = os.Getenv("MongoUsername")
	cred.Password = os.Getenv("MongoPassword")
	mongoUrl := os.Getenv("MongoUrl")
	clientOptions = clientOptions.ApplyURI(mongoUrl).SetAuth(cred)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ... !!")

	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	fmt.Println(collection)
	return collection
}

func main() {
	fmt.Println("connected to authentication service")
	fmt.Println("HELLLO AUTH SERVICEeeeee")

	writer, closeWriter := api.NewWriter[CreateAccountDto]("kafka:9092", "twitter.newTweets")
	reader, closeReader := api.NewReader[CreateAccountDto]("kafka:9092", "twitter.newTweets", "saver", func(tweet CreateAccountDto) {
		// retry process
		fmt.Println("error, retrying ...")
		writer.WriteBatch(context.TODO(), tweet)
	})
	defer closeReader()
	defer closeWriter()

	go reader.Read(func(dto CreateAccountDto) error {
		if dto.FirstName != "" {
			fmt.Println("received a message: ", dto.url)
		}
		fmt.Println(dto.url)
		if rand.Intn(100) > 50 {
			return errors.New("a random error")
		}
		return nil
	})
	go func() {
		reciver.Reciver()
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
	fmt.Println("Closing Kafka connections ...")
}

type CreateAccountDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	url       string `json:url`
}
