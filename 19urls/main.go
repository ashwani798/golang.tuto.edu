package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.ddev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)
	//Welcome to handling urls in golang

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	// https
	// lco.ddev:3000
	// /learn
	// 3000
	// coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query()
	fmt.Printf("The type of params are: %T\n", qparams)
	// The type of params are: url.Values
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])
	// [reactjs]
	// [ghbj456ghb]

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}
	// param is:  [reactjs]
	// param is:  [ghbj456ghb]

	partsOfurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dec",
		Path:    "tuto",
		RawPath: "user=ashu",
	}

	anotherURL := partsOfurl.String()
	fmt.Println(anotherURL)
}
