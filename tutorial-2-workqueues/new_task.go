package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func produce() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for {
		go func() {
			body := bodyFrom()
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Failed to publish a message")
			log.Printf("\x1b[%dm [x] Sent %s\x1b[0m", 36, body)

		}()
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func bodyFrom() string {
	return string(rand.Intn(100))
}
