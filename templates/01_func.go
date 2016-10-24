package main

import "fmt"

func main() {
	unexportedFunction()

	ExportedFunction()

	fmt.Println(funcWithReturn())

	fmt.Println(funcWithParameter(1))

	str := "Hello, World!"
	fmt.Printf("%v => %v\n", str, &str)
	fmt.Println(funcWithPointer(&str))
}

//unexportedFunction  method that only available in the package level
func unexportedFunction() {
}

//ExportedFunction method that available from outside of the package
func ExportedFunction() {
}

func funcWithReturn() int {
	return 42
}

// when a primitiv is passed as argument, it is passed by value,
// which means it's copy will be passed not the original one
// from this the slice, map, chan, array are exceptions
func funcWithParameter(parameter int) int {
	return parameter + 0
}

// * operator before the type means it's a pointer
// with this you can pass memory reference and not a copied value to the method
// for normal not pointer types you can still pass reference too, but with this
// you can explicitly say that you need the object it self, and not the value
func funcWithPointer(parameter *string) string {
	fmt.Printf("%v => %v\n", *parameter, parameter)
	return *parameter
}
