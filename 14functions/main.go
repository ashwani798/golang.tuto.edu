package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	//Welcome to functions in golang
	greeter()
	greeterTwo()

	// result := adder(3, 4)

	// fmt.Println("result is: ", result)

	proRes, myAdd := proAdder(2, 4, 7, 17, 10, 2, 8)
	fmt.Println("pro result is: ", proRes)
	fmt.Println("Total Add is: ", myAdd)
	//pro result is:  50
}

// func adder(valOne int, valTwo int) int {
// 	return valOne + valTwo
// }
//result is:  7

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi, my name is ashu"
	//Total Add is:  hi, my name is ashu
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namste from INDIA")
	//Namste from INDIA
}
