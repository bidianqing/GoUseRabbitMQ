package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	failOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err)
	defer ch.Close()

	ch.Qos(10, 0, false)
	msgs, err := ch.Consume(
		"red",    // queue
		"golang", // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err)

	for d := range msgs {
		msg := string(d.Body)
		fmt.Printf("Received a message: %v \n", msg)

		ch.Ack(d.DeliveryTag, false)
		if msg == "cancel" {
			ch.Cancel("golang", false)
		}
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
