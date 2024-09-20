package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var pointer *int
	// fmt.Println("value of pointers is ", pointer)

	myNumber := 23

	var pointer = &myNumber
	// for refrence already store value use & before the variable

	fmt.Println("value of actual pointer is ", pointer)
	//value of actual pointer is  0xc00008c098
	fmt.Println("value of actual pointer is ", *pointer)
	//value of actual pointer is  23

	*pointer = *pointer + 2
	fmt.Println("New value is: ", myNumber)
	//New value is:  25

}
