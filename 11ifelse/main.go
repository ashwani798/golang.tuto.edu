package main

import "fmt"

func main() {
	fmt.Println("ifelse in golang")
	//ifelse in golang
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactuly 10 login count"
	}

	fmt.Println(result)
	//Exactuly 10 login count

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
		//Number is odd
	}

	if num := 3; num < 10 {
		fmt.Println("num is less then 10")
		//num is less then 10
	} else {
		fmt.Println("num is NOT less then 10")
	}

	// if err !=  nil[

	// ]
}
