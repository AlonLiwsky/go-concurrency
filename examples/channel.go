package main

import "fmt"

func main(){
	c := make(chan string)
	go writeToChannel(c)
	v := <- c
	fmt.Print(v)
}

func writeToChannel(c chan string) {
	c <- "Message"
}
