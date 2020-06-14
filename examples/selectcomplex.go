package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	receive := make(chan string)
	send := make(chan string)
	go getResponse(receive)
	go sendMessage(send)

	random := rand.Intn(2000)
	time.Sleep(time.Duration(random) * time.Millisecond)

	for {
		select {
		//channel assign receive
		case response := <-receive:
			fmt.Println(response + " received")

		//channel send
		case send <- "Message sent":

		//timeout (channel receive)
		case value := <-time.After(time.Millisecond * time.Duration(rand.Intn(1))):
			fmt.Println(value)
			return

		default:
			fmt.Println("No one was ready")

		}
	}

}

func getResponse(c chan<- string) {
	random := rand.Intn(2000)
	time.Sleep(time.Duration(random) * time.Millisecond)
	c <- "Response"
}

func sendMessage(c <-chan string) {
	random := rand.Intn(2000)
	time.Sleep(time.Duration(random) * time.Millisecond)
	fmt.Println(<-c)
}
