package main

func main() {
	// quando criar um chanel, e colocar um numero do lado, quer dizer que esta definindo buffer
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
