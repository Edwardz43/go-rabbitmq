package main

import "go-rabbitmq/lib"

var failOnError = lib.FailOnError

func main() {
	for i := 0; i < 10; i++ {
		go consume(i)
	}
	produce()
}
