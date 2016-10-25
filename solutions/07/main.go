/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern
*/

package main

import (
	"fmt"
)

func main() {

	in := gen()

	channels := make([]<-chan int, 0)
	for i := 0; i < 10; i++ {
		f := factorial(in)
		channels = append(channels, f)
	}

	final := make(chan int)
	go func(chs []<-chan int, out chan<- int) {
		for _, ch := range chs {
			for n := range ch {
				out <- n
			}
		}
		close(out)
	}(channels, final)

	for n := range final {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
