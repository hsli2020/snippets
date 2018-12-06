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

type Book struct { // Book represents an item from the "books" table.
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	sess.SetLogging(true)

	booksTable := sess.Collection("books")

	// A result set that has only one element.
	res := booksTable.Find(4267)

	// Get the element with the given ID.
	var book Book
	err = res.One(&book)
	if err != nil {
		log.Fatal("Find: ", err)
	}

	log.Printf("Book: %#v", book)

	// Change a property.
	book.Title = "New title"

	log.Printf("Book (modified): %#v", book)

	// Update the result set (which only has one element).
	if err := res.Update(book); err != nil {
		log.Printf("Update: %v\n", err)
		log.Printf("This is OK, this is a restricted sandbox with a read-only database.")
	}

	// Delete the result set (which only has one element).
	if err := res.Delete(); err != nil {
		log.Printf("Delete: %v\n", err)
		log.Printf("This is OK, this is a restricted sandbox with a read-only database.")
	}
}
