package main

import (
	"fmt"
	"time"
)

func main() {

	// go command maked the passed function call executed as async goroutine
	// this means it could use different cpu and and may be executed in parralell
	go doSomeWork("first")
	go doSomeWork("second")
	go doSomeWork("third")

	time.Sleep(1 * time.Second)

}

func doSomeWork(msg string) {
	fmt.Println(msg)
}
