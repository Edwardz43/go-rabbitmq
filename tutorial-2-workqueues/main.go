package main

import "go-rabbitmq/lib/err"

var failOnError = err.FailOnError

func main() {
	for i := 0; i < 10; i++ {
		go consume(i)
	}
	produce()
}
