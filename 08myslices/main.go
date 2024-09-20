package main

import "fmt"

func main() {
	fmt.Println("welcome to slices like mango Slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList) //Type of fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana") //[Apple Orange Peach Mango Banana]
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //[Orange Peach Mango Banana]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
}
