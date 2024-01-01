package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fmt.Println("helllo from authentication 1 service")
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	groupID := os.Getenv("groupID")
	reader := getKafkaReader(kafkaURL, topic, groupID)
	mongoURL := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")
	collection := getMongoCollection(mongoURL, dbName, collectionName)
	fmt.Println("mongoUri", mongoURL)
	defer reader.Close()
	fmt.Println("start consuming ... !!")

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Println(err)
	}
	fmt.Println("asdasdasdasdsad", string(msg.Key))
	for {
		str2 := string(msg.Key)
		fmt.Println(string(msg.Key))
		fmt.Println("valueeeee", string(msg.Value))
		if str2 == "FromPostRoute" {

			fmt.Printf("FromPostRoute,%s \n", str2)
		}
		fmt.Println(collection)
		insertResult, err := collection.InsertOne(context.Background(), msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		time.Sleep(1 * time.Second)

	}

}
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB

	})
}
func getKafkaReaderTestMsg() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		GroupID:  "auth-service",
		Topic:    "my-topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
