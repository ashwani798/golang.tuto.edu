package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
    //Maps in golang
	languages := make(map[string]string)

	languages["JS"] = "javaScript"
	languages["GO"] = "golang"
	languages["PY"] = "pyhton"

	fmt.Println("List of all languages : ", languages)
	fmt.Println("JS shorts for : ", languages["JS"]
    //List of all languages :  map[GO:golang JS:javaScript PY:pyhton]
    //JS shorts for :  javaScript

	delete(languages, "JS")
	fmt.Println("List of all languages : ", languages)
	//List of all languages :  map[GO:golang PY:pyhton]
	
	//****************************************************************************
	// Loops are interesting in golang
	for key, value := range languages {
		fmt.Println("For key %v, value is %v\n", key, value)
	}
    //For key %v, value is %v
    //GO golang
    //For key %v, value is %v
    //PY pyhton
}
