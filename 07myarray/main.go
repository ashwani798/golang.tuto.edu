package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var FruitList [4]string

	FruitList[0] = "Apple"
	FruitList[1] = "Orange"
	FruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", FruitList)
	//Fruit list is:  [Apple Orange  Peach]
	fmt.Println("Fruit list is: ", len(FruitList))
	//Fruit list is:  4

	var vegList = [5]string{"potato", "tomato", "onion"}
	fmt.Println("veg list is:", len(vegList))
	// veg list is: 5
}
