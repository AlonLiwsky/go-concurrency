package main

import (
	"fmt"
	"time"
)

type Pointer struct{
	Value *string
}

func main(){
	pointer := &Pointer{nil}
	go thread1(pointer)
	go thread2(pointer)
	time.Sleep(time.Second)
}

func thread1(pointer *Pointer) {
	value := "Data"
	*pointer = Pointer{Value: &value}
}

func thread2(pointer *Pointer) {
	fmt.Print(*pointer.Value)
}