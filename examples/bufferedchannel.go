package main

import (
	"fmt"
	"time"
)

//Buffered channel
func main() {
	c := make(chan string, 3)
	writeToChannel(c, "buffered")
	writeToChannel(c, "channel")
	writeToChannel(c, "test")

	time.Sleep(time.Second)

	//for m := range c {
	//	fmt.Println(m)
	//}

	size := len(c)
	for i := 0; i < size; i++ {
		fmt.Println(<-c)
	}

}

func writeToChannel(c chan string, message string) {
	c <- message
}
