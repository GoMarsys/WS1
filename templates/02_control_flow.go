package main

import (
	"fmt"
	"time"
)

func main() {
	switchExample()
	ifExample("Fizz")
	ifExample("Buzz")
	ifExample("FizzBuzz")
	ifExample("something else")
}

func ifExample(input string) {
	if input == "FizzBuzz" {
		fmt.Println("FizzBuzz")
	} else if input == "Fizz" {
		fmt.Println("Fizz")
	} else if input == "Buzz" {
		fmt.Println("Buzz")
	} else {
		fmt.Println("hmm...")
	}
}

func switchExample() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	case time.Monday:
		fmt.Println("It's Monday, don't forget to be AWESOME!")
	default:
		fmt.Println("it's a weekday")
	}
}
