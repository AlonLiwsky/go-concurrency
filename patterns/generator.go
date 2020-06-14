package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j:= 0, 1; i < n ; i, j = i+j,i {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	// fib returns the fibonacci numbers lesser than 1000
	for i := range fib(1000) {
		// Consumer which consumes the data produced by the generator
		fmt.Println(i)
	}
}
