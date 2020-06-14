package main

import (
	"fmt"
	"math/rand"
	"time"
)

func server(serverID int) chan<- int {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(fmt.Sprintf("Server %d handled operation %d", serverID, v))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}()
	return c
}

func numbersGenerator(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	dataToProcess := numbersGenerator(10)
	server1 := server(1)
	server2 := server(2)

	for data := range dataToProcess {
		select {
		case server1 <- data:
		case server2 <- data:
		}
	}
	close(server1)
	close(server2)
}
