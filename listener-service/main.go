package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitcon, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitcon.Close()

	// start listening for messages
	log.Println("Listening for and consuming Rabbitmq messages...")

	// create consumer
	consumer, err := event.NewConsumer(rabbitcon)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println("Error", err)
	}
}

func connect() (*amqp.Connection, error) {
	// backoff
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("Rabbitmq not yet ready,", err)
			counts++
		} else {
			log.Println("Connected to Rabbitmq")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off while trying to connecto to Rabbitmq...")
		time.Sleep(backoff)
		continue
	}

	return connection, nil
}
