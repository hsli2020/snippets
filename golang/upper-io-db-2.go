package main

import (
	"log"

	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}

// Customer represents an item from the "customers" table.
type Customer struct {
	ID        uint   `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	customersTable := sess.Collection("customers")

	// Creates a paginator with 10 items per page.
	p := customersTable.Find().
		OrderBy("last_name", "first_name").
		Paginate(10)

	var customers []Customer

	// Dump all results into the customers slice.
	err = p.Page(2).All(&customers)
	if err != nil {
		log.Fatal(err)
	}

	for i, customer := range customers {
		log.Printf("%d: %s, %s", i, customer.LastName, customer.FirstName)
	}

	totalNumberOfEntries, err := p.TotalEntries()
	if err != nil {
		log.Fatal(err)
	}

	totalNumberOfPages, err := p.TotalPages()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("total entries: %d, total pages: %d", totalNumberOfEntries, totalNumberOfPages)
}
