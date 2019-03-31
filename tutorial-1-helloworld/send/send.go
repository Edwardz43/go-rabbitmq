package main

import (
	"flag"
	"go-rabbitmq/lib/err"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

var failOnError = err.FailOnError

var ip = flag.String("ip", "127.0.0.1", "server IP")

func main() {

	flag.Parse()

	conn, err := amqp.Dial("amqp://docker1:P@ssw0rd@" + *ip + ":5672/")

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
		log.Printf("\x1b[%dmProducer Sent %s\x1b[0m", 32, body)
		failOnError(err, "Failed to publish a message")

		t := rand.Intn(10)

		time.Sleep(time.Second * time.Duration(t))
	}

}
