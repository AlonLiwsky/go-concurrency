package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(input1, input2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func numbersGenerator(n int) <-chan int {
	c := make(chan int)
	go func() {
		for{
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			c<- n
		}
	}()
	return c
}

func main(){
	input1 := numbersGenerator(1)
	input2 := numbersGenerator(2)

	c := fanIn(input1, input2)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
