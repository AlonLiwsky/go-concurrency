package main

func main(){
	r1 := make(chan int)
	r2 := make(chan int)
	finished := make(chan bool)
	go thread1(r1, r2, finished)
	go thread2(r1, r2, finished)
	<- finished
}

func thread1(resource1 chan int, resource2 chan int, finished chan bool) {
	<-resource2
	resource1 <- 1
	finished <- true
}

func thread2(resource1 chan int, resource2 chan int, finished chan bool) {
	<-resource1
	resource2 <- 1
	finished <- true
}
