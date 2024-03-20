package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"red",    // queue
		"golang", // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		fmt.Printf("Received a message: %s \n", d.Body)
	}
}
