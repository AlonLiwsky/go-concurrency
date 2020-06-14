package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c1 := make(chan string)
	c2 := make(chan string)
	go getResponse(c1)
	go getResponse(c2)

	select {
	case response := <-c1:
		fmt.Printf("Server 1 returned: %s", response)
	case response := <-c2:
		fmt.Printf("Server 2 returned: %s", response)
	}
}

func getResponse(c chan string) {
	random := rand.Intn(2000)
	time.Sleep(time.Duration(random) * time.Millisecond)
	c <- "Response"
}
