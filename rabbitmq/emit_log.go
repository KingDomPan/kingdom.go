package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:panqd@192.168.99.100:5672/")
	failOnError(err, "Failed to connect to RabbitMq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to Publish a Message")
	log.Printf("[x] send %s", body)

}

func bodyFrom(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "hello"
	} else {
		return strings.Join(args[1:], " ")
	}
}
