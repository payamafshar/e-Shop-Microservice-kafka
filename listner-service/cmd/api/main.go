package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connect() (*amqp.Connection, error) {

	var counts int64
	var backoffs float64
	var connection *amqp.Connection
	time.Sleep(2 * time.Second)

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			counts++
			log.Println("RabbitMq is not ready")
		} else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println("errrrrrrrr", err)
			return nil, err
		}
		backoffs = time.Duration(math.Pow(float64(counts), 2)).Seconds()
		log.Println("Trying connect to rabbitMQ")
		time.Sleep(time.Duration(backoffs))
		continue
	}
	fmt.Println("conncetion *****", connection)
	return connection, nil
}

func main() {
	fmt.Println("hello from listner-service")
	var mutext sync.Mutex
	var wg sync.WaitGroup
	mutext.Lock()
	wg.Add(1)
	var c, err = amqp.Dial("amqp://guest:guest@rabbitmq")
	wg.Done()
	mutext.Unlock()

	wg.Wait()
	if err != nil {
		log.Println("err from connection to rabbitMQ******")
		log.Panic(err)

	} else {
		fmt.Println("kkkkkkkkkkkkkkk", c.LocalAddr())
		log.Println("Connected to rabbitMQ")
	}
	// defer rabbitmqConnection.Close() //when its not connected trying to close nil got panic
	defer c.Close()

}
