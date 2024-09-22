package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	//Welcome to loops in golang

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thrusday", "Saturday"}

	fmt.Println(days)
	//[Sunday Tuesday Wednesday Thrusday Saturday]

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
		// index is 0 and value is Sunday
		// index is 1 and value is Tuesday
		// index is 2 and value is Wednesday
		// index is 3 and value is Thrusday
		// index is 4 and value is Saturday
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 4 {
			rougueValue++
			break //value is :  1
		} //value is :  2
		fmt.Println("value is : ", rougueValue) //value is :  3
		rougueValue++
	}

lco:
	fmt.Println("jumping at golangLearning")
	//value is :  1
	//jumping at golangLearning

}
