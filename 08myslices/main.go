package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices like mango Slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList) //Type of fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana") //[Apple Orange Peach Mango Banana]
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //[Orange Peach Mango Banana]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) //[Peach Mango]

	fruitList = append(fruitList[0:3])
	fmt.Println(fruitList) //[Peach Mango Banana]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) //[Mango Banana]

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 834
	highScore[2] = 454
	highScore[3] = 934
	//highScore[4] = 777

	highScore = append(highScore, 444, 555, 666)
	//[234 834 454 934 444 555 666]

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	//false
	//[234 444 454 555 666 834 934]

	//**********************************************************************
	// How to remove a value from slices based on index      ******

	var courses = []string{"HTML", "CSS", "javaScript", "reactJs", "Next.js", "Python", "golang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	//  append not only use for add the value as well as append use to remove the value
	fmt.Println(courses)
	//[HTML CSS javaScript reactJs Next.js Python golang]
	//[HTML CSS reactJs Next.js Python golang]
}
