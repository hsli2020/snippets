package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type Customer struct {
	Id   string
	Name string
}

func main() {
	doSqlDriverStuff()
}

func doSqlDriverStuff() {
	connStr := "user=postgres dbname=Customer password=viXLbzKq3seyOROsmirW sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if !doCustomersExist(db) {
		addCustomers(db)
	}
	customers := getCustomers(db)
	for _, customer := range customers {
		fmt.Printf("id=%s name=%s\n", customer.Id, customer.Name)
	}
	customer := getCustomer("f71d62ed-58b5-438b-8faa-3764499484fd", db)
	fmt.Printf("id=%s name=%s\n", customer.Id, customer.Name)
}

func doCustomersExist(db *sql.DB) bool {
	var customerCount int
	err := db.QueryRow(`select count(*) from "Customer" as customerCount`).
		Scan(&customerCount)
	if err != nil {
		log.Fatal(err)
	}
	return customerCount > 0
}

func addCustomer(name string, db *sql.DB) string {
	var customerId string
	id := uuid.New().String()
	err := db.QueryRow(`INSERT INTO "Customer"(id, name)
		VALUES($1,$2) RETURNING id`, id, name).Scan(&customerId)
	if err != nil {
		log.Fatal(err)
	}
	return customerId
}

func addCustomers(db *sql.DB) {
	var customers = [10]string{
		"James", "John", "Jimmy", "Does Not Start with J", "Hermoine",
		"Narcissus", "Hank", "Heather", "Bocephus", "Bob"}
	for _, customer := range customers {
		_ = addCustomer(customer, db)
	}
}

func getCustomers(db *sql.DB) []Customer {
	rows, err := db.Query(`SELECT id, name FROM "Customer"`)
	if err != nil {
		log.Fatal(err)
	}
	var customerArray []Customer
	for rows.Next() {
		var customer Customer
		var id string
		var name string
		err = rows.Scan(&id, &name)
		customer = Customer{Id: id, Name: name}
		customerArray = append(customerArray, customer)
	}
	defer db.Close()
	return customerArray
}

func getCustomer(id string, db *sql.DB) Customer {
	var customerId string
	var name string

	err := db.QueryRow(`SELECT id, name FROM "Customer" where "id"=$1`, id).
		Scan(&customerId, &name)

	if err != nil {
		log.Fatal(err)
	}

	return Customer{Id: customerId, Name: name}
}
