package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func consume(i int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	for {
		body := "Hello World! " + string(rand.Intn(100))
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf("\x1b[%dmProducer[%d] Sent %s\x1b[0m", 34, i, body)
		failOnError(err, "Failed to publish a message")

		t := rand.Intn(500)

		time.Sleep(time.Millisecond * time.Duration(t))
	}

}
