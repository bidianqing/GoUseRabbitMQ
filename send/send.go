package main

import (
	"context"

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

	for {

		body := "Hello World!"
		err = ch.PublishWithContext(context.Background(),
			"duck-direct", // exchange
			"red",         // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err)

	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
