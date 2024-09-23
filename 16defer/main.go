package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	fmt.Println("Hello")
	myDefer()
}

// Hello
// 43210
// world2
// world1
// world
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
