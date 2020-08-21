package main	// main.go

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	_db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = _db

	if err := db.AutoMigrate(&Grocery{}); err != nil {
		panic(err)
	}

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/allgroceries", AllGroceries)
	r.HandleFunc("/groceries/{name}", SingleGrocery)
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", r))
}

package main	// grocery.go

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

package main	// handlers.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

var groceries = []Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple",      Quantity: 6},
}

func AllGroceries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(w).Encode(groceries)
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, grocery := range groceries {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var grocery Grocery
	json.Unmarshal(reqBody, &grocery)
	groceries = append(groceries, grocery)

	json.NewEncoder(w).Encode(groceries)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)
		}
	}
}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updateGrocery Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries = append(groceries, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}
}

package main	// repository.go

func findGrocery(groceries string) (*Grocery, error) {
	var grocery Grocery
	if result := db.Where(&Grocery{Name: groceries}).First(&grocery); result.Error != nil {
		return nil, result.Error
	}
	return &grocery, nil
}

func findQuantity(quantity int) ([]Grocery, error) {
	var groceries []Grocery
	if result := db.Where("Quantity > ?", quantity).Find(&groceries); result.Error != nil {
		return nil, result.Error
	}
	return groceries, nil
}

func insertGrocery(grocery *Grocery) error {
	if result := db.Create(grocery); result.Error != nil {
		return result.Error
	}
	return nil
}