package main

import (
	"fmt"
	"time"
)

type Pointer struct{
	Value *string
}

func main(){
	text := "text"
	pointer := &Pointer{Value:&text}
	go thread1(pointer)
	go thread2(pointer)
	time.Sleep(time.Second)
}

func thread1(pointer *Pointer) {
	if pointer.Value != nil {
		time.Sleep(1)
		fmt.Print(*pointer.Value)
	}
}

func thread2(pointer *Pointer) {
	pointer.Value = nil
}