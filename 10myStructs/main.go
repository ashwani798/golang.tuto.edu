package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//tructs in golang

	// no inheritance in golang; No super or parent
	Ashu := User{"Ashu", "ashu@go.dev", true, 22}
	fmt.Println(Ashu)
	//{Ashu ashu@go.dev true 22}
	fmt.Printf("Ashu details are: %+v", Ashu)
	//Ashu details are: {Nmae:Ashu Emali:ashu@go.dev Status:true Age:22}
	fmt.Printf("Name is %v and email is %v.", Ashu.Name, Ashu.Emali)
	//Name is Ashu and email is ashu@go.dev.
}

type User struct {
	Name   string
	Emali  string
	Status bool
	Age    int
}
