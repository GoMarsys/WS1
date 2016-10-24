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

	// this will close the channel, so no more value can be put in
	// this will also cause range iteration to be terminated,
	// and to not expect any more value
	close(out)
}

func consumer(in <-chan int) {
	// range iteration on a channel cancel the iteration when the channel is closed
	// this is a good way to receive values from the channel, but the amount is unknown beforehand
	for n := range in {
		fmt.Println(n)
	}
}
