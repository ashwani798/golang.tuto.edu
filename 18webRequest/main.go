package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://pkg.go.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// LCO web request
	// Response is of type: *http.Response

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll((response.Body))

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}

//<!DOCTYPE html> file
