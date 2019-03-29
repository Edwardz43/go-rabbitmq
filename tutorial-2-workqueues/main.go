package main

func main() {

	go consume(1)
	go consume(2)
	produce()
}
