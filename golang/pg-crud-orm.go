package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/google/uuid"
)

type Customer struct {
	Id   string
	Name string
}

func main() {
	doOrmStuff()
}

func doOrmStuff() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "Customer",
		Password: "viXLbzKq3seyOROsmirW",
	})
	createSchema(db)
	if !doCustomersExist(db) {
		addCustomers(db)
	}
	customers := getCustomers(db)
	for _, customer := range customers {
		fmt.Printf("id=%s name=%s\n", customer.Id, customer.Name)
	}
	customer := getCustomer("1bbe4c12-b25a-4142-a4c5-3695e4905786", db)
	fmt.Printf("id=%s name=%s\n", customer.Id, customer.Name)
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Customer)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func doCustomersExist(db *pg.DB) bool {
	customers := getCustomers(db)
	return len(customers) > 0
}

func addCustomer(name string, db *pg.DB) string {
	id := uuid.New().String()
	customer := &Customer{
		Id:   id,
		Name: name,
	}
	err := db.Insert(customer)
	if err != nil {
		log.Fatal(err)
	}
	return customer.Id
}

func addCustomers(db *pg.DB) {
	var customers = [10]string{
		"James", "John", "Jimmy", "Does Not Start with J", "Hermoine", 
		"Narcissus", "Hank", "Heather", "Bocephus", "Bob"
	}
	for _, customer := range customers {
		_ = addCustomer(customer, db)
	}
}

func getCustomers(db *pg.DB) []Customer {
	var customers []Customer
	err := db.Model(&customers).Select()
	if err != nil {
		log.Fatal(err)
	}
	return customers
}

func getCustomer(id string, db *pg.DB) Customer {
	customer := &Customer{Id: id}
	err := db.Select(customer)
	if err != nil {
		log.Fatal(err)
	}
	return *customer
}
