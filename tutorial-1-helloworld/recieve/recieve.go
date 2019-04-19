package main

import (
	"go-rabbitmq/config"
	"go-rabbitmq/lib/err"
	"log"

	"github.com/streadway/amqp"
)

var failOnError = err.FailOnError

func main() {
	conn, err := amqp.Dial(config.DNS)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("\x1b[%dmConsumer received a message: %s\x1b[0m", 35, d.Body)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
