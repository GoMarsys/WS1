package main

import "fmt"

func main() {
	setValueByIndex()
	newSliceByAppend()
	fmt.Printf("\n")
}

func setValueByIndex() {

	fmt.Printf("\nsetValueByIndex:\n\n")

	slice := make([]int, 1)
	fmt.Println("the slice initial state")
	fmt.Printf("%T: %v\n", slice, slice)

	slice[0] = 42 // alter the original slice
	fmt.Println("the slice after set value to index position")
	fmt.Printf("%T: %v\n", slice, slice)

}

func newSliceByAppend() {

	fmt.Printf("\nnewSliceByAppend:\n\n")

	slice := make([]int, 0, 5)
	fmt.Println("the original slice initial state")
	fmt.Printf("%T: %v\n", slice, slice)

	newSlice := append(slice, 1) // return a new slice, not modify the original one
	fmt.Println("the original slice, after append")
	fmt.Printf("%T: %v\n", slice, slice)
	fmt.Println("the new slice created by append")
	fmt.Printf("%T: %v\n", newSlice, newSlice)

}

func slicingASlice() {

/}
