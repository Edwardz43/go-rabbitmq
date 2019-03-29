package main

func main() {

	for i := 0; i < 150; i++ {
		go produce(i)
	}

	for i := 0; i < 90; i++ {
		go consume(i)
	}
	consume(90)
}
