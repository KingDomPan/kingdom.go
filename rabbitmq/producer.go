package main

// 使用rabbitmq的控制台管理界面创建好以下类型
// 1. exchangeType = fanout exchangeName = test-exchange
// 2. queueName = [test-queue1, test-queue2]

import "flag"
import "log"
import "fmt"
import "github.com/streadway/amqp"

var (
	uri          = flag.String("uri", "amqp://admin:panqd@192.168.99.100:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "fanout", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "message body is kingdompan's name", "Body of message")
	reliable     = flag.Bool("reliable", false, "Wait for the publisher confirmation before exiting")
)

func init() {
	flag.Parse()
}

func publish(amqpUri, exchangeName, exchangeType, routingKey, body string, reliable Bool) error {

	log.Println("dialing %q", amqpUri)
	connection, err := amqp.Dial(amqpUri)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	defer connection.Close()

	log.Println("get connection, get channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	if reliable {
		log.Println("enabling publishing confirms")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("Channel could not be put into confirm mode: %s", err)
		}
		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer confirmOne(confirms)
	}

	log.Println("declare exchange, publish %dB body (%q)", len(body), body)
	if err = channel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliverMode:     amqp.Transient, // 1不持久化 2持久化
			Priority:        0,
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return nil
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Println("waiting for confirmmation of one publishing")
	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %s", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery with delivery tag: %s", confirmed.DeliveryTag)
	}
}

func main() {
	if err := publish(*uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable); err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("publish %dB Ok", len(*body))
}
