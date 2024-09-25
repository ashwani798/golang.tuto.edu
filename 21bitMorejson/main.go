package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"Price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to JSON video")
	//Encodejson()
	DecodeJson()
}

func Encodejson() {
	myCourses := []Course{
		{"reactJS Bootstrap", 299, "learnwithme.in", "abcdxyz", []string{"web-dev", "js"}},
		{"MERN Stack", 299, "learnwithme.in", "bcdxyz", []string{"fullstack", "js"}},
		{"Python Dev", 299, "learnwithme.in", "cdxyz", nil},
	}

	// Package this data into JSON format
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	// Print JSON output
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
		"coursename": "reactJS Bootstrap",
		"Price": 299,
		"website": "learnwithme.in",
		"tags": ["web-dev","js"]
	 }`)

	var myCourses Course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonDataFromWeb, &myCourses)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		} else {
			fmt.Printf("%#v\n", myCourses) // This will print the struct with field names
		}
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	// Decoding into a map[string]interface{} to handle dynamic JSON keys/values
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	// Print the map with detailed info
	fmt.Printf("%#v\n", myOnlineData)

	// Iterating over the key-value pairs in the map
	for k, v := range myOnlineData {
		// Use fmt.Printf to print formatted output
		fmt.Printf("key is %v and value is %v and type of data is: %T\n", k, v, v)
	}

}

// welcome to JSON video
// [
//         {
//                 "coursename": "reactJS Bootstrap",
//                 "Price": 299,
//                 "website": "learnwithme.in",
//                 "tags": [
//                         "web-dev",
//                         "js"
//                 ]
//         },
//         {
//                 "coursename": "MERN Stack",
//                 "Price": 299,
//                 "website": "learnwithme.in",
//                 "tags": [
//                         "fullstack",
//                         "js"
//                 ]
//         },
//         {
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "website": "learnwithme.in"
//         }
// ]
// PS C:\Users\Dell\golang-LCO\21bitMorejson>
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "website": "learnwithme.in"
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "website": "learnwithme.in"
//         }
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "coursename": "Python Dev",
//                 "Price": 299,
//                 "website": "learnwithme.in"
//         }
// ]
// JSON was valid
// main.Course{Name:"reactJS Bootstrap", Price:299, Platform:"learnwithme.in", Password:"", Tags:[]string{"web-dev", "js"}}
// map[string]interface {}{"Price":299, "coursename":"reactJS Bootstrap", "tags":[]interface {}{"web-dev", "js"}, "website":"learnwithme.in"}
// key is coursename and value is reactJS Bootstrap and type of data is: string
// key is Price and value is 299 and type of data is: float64
// key is website and value is learnwithme.in and type of data is: string
// key is tags and value is [web-dev js] and type of data is: []interface {}
