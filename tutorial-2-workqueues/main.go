package main

func main() {

	go func(i int) { consume(i) }(1)
	go func(i int) { consume(i) }(2)
	produce()

}
