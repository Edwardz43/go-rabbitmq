package main

func main() {

	for i := 0; i < 150; i++ {
		go func(i int) { produce(i) }(i)
	}

	for i := 0; i < 90; i++ {
		go func(i int) { consume(i) }(i)
	}
	consume(90)
}
