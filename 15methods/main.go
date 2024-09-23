// methods
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
	//Ashu details are: {Nmae:Ashu Email:ashu@go.dev Status:true Age:22}
	fmt.Printf("Name is %v and email is %v.", Ashu.Name, Ashu.Email)
	//Name is Ashu and email is ashu@go.dev.
	Ashu.GetStatus()
	Ashu.NewMail()
	fmt.Printf("Name is %v and email is %v.", Ashu.Name, Ashu.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
	//Is user active:  true
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
	//Email of this user is:  test@go.dev
	//Name is Ashu and email is ashu@go.dev.
}
