package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome devloper"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enetr the rating for our developer:")

	// commo Ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the rating %T", input)
}
