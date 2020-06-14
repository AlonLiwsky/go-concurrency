package main

import (
	"fmt"
)

func futureData(value float64) <-chan float64 {
	c := make(chan float64, 1)

	go func() {
		//heavy computational task o IO
		result := (value * 8 - 4 * 3)/5 + 1
		c <- result
	}()

	return c
}

func main() {
	future := futureData(42)

	// do other things

	fmt.Printf("response: %v", <- future)
}
