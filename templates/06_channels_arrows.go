package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go publisher(ch)
	go consumer(ch)

	time.Sleep(3 * time.Second)

}

func publisher(out chan<- int) {
	for i := 0; i < 10; i++ {
		// with "<-" we start to insert a value to the channel
		// until someone takes it from us, we wait
		out <- i
	}
}

func consumer(in <-chan int) {
	for i := 0; i < 10; i++ {
		// with "<-" we start to receive a value from the channel
		// until someone gives one, we will wait
		fmt.Println(<-in)
	}
}
