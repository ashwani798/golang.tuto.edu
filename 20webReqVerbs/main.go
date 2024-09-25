package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webReqVerbs")
	//performGetRequest()
	//perfromPostjsonRequest()
	perfromPostFromRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	//fmt.println(content)
	//fmt.print(string(content))
}

// Welcome to webReqVerbs
// Status code:  200
// Content length is:  43
// ByteCount is:  43
// {"message":"Hello from learnCodeonline.in"}

func perfromPostjsonRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	     {
	        "courseNmae":"Let's go for golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		 }
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// {"courseNmae":"Let's go for golang","price":0,"platform":"learnCodeOnline.in"}

func perfromPostFromRequest() {
	const myurl = "http://localhost:8000/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname", "ashu")
	data.Add("lastname", "pandey")
	data.Add("email", "ashu@go.dev")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// {"email":"ashu@go.dev","firstname":"ashu","lastname":"pandey"}
