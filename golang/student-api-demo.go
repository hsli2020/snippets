package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Learning GoLang from Tutorial
// Source: https://www.youtube.com/watch?v=TkbhQQS3m_o
// 8.11.2021 / haapjari / CRUD API with GoLang

/* ---------------------------------------------------------- */

/* Data Structures */

type Student struct {
	ID    string `json: "id"`
	Name  string `json: "name"`
	Major *Major `json: "major"`
}

type Major struct {
	Subject string `json: "subject"`
	Status  string `json: "status"`
}

var students []Student

/* ---------------------------------------------------------- */

/* Get All */

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

/* Delete Method */

func deleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range students {

		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

/* Single Getter */

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range students {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

/* Create Method */

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(10000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

/* Update Method */

func updateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

/* Main Function */

func main() {
	r := mux.NewRouter()

	students = append(students, Student{ID: "1", Name: "John Doe", Major: &Major{Subject: "Computer Science", Status: "Active"}})
	students = append(students, Student{ID: "2", Name: "Dohn Joe", Major: &Major{Subject: "Finance", Status: "Inactive"}})

	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
