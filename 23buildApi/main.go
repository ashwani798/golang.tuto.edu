package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"_"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) Isempty() bool {
	return c.CourseName == ""
}

// Add your router setup and course initialization here
func main() {
	fmt.Println("API - LearncodeOnline.in")
	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Ashwani Pandey", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python Dev", CoursePrice: 399, Author: &Author{Fullname: "Ashwani Pandey", Website: "py.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API created by me</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// decode the request body
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// validate if the course data is empty
	if course.Isempty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//TODO: check only if title is duplicate
	//loop, title matches with course.coursename
	// generate a unique course ID using crypto/rand
	courseId, _ := rand.Int(rand.Reader, big.NewInt(100))
	course.CourseId = courseId.String()

	// append course into courses
	courses = append(courses, course)

	// return the created course
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("content-type", "application/json")

	// first - grap id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}

}
