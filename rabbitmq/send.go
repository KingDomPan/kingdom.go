package main

// The default exchange is implicitly bound to every queue, with a routing key equal to the queue name
import (
	"fmt"
	"log"

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

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := "panqd is kingdompan"
	err = ch.Publish(
		"",     // exchange 默认的exchange
		q.Name, // 队列名称
		false,  // mandatory
		false,  // immedaite
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	log.Printf(" [x] Send %s ", body)
	failOnError(err, "Failed to publish a message")
}
