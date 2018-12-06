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

// Book represents an item from the "books" table.
type Book struct {
	ID        uint   `db:"id,omitempty"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id,omitempty"`
	SubjectID uint   `db:"subject_id,omitempty"`
}

type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

type Subject struct {
	ID       uint   `db:"id,omitempty"`
	Subject  string `db:"subject"`
	Location string `db:"location"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	sess.SetLogging(true)

	// The BookAuthorSubject type represents an element that
	// has columns from different tables.
	type BookAuthorSubject struct {
		// The book_id column was added to prevent collisions
		// with the other "id" columns from Author and Subject.
		BookID uint `db:"book_id"`

		Book    `db:",inline"`
		Author  `db:",inline"`
		Subject `db:",inline"`
	}

	// This is a query with a JOIN clause that was built using the SQL builder.
	q := sess.Select("b.id AS book_id", "*"). // Note how we set an alias for book.id.
							From("books AS b").
							Join("subjects AS s").On("b.subject_id = s.id").
							Join("authors AS a").On("b.author_id = a.id").
							OrderBy("a.last_name", "b.title")

	// The JOIN query above returns data from three different tables.
	var books []BookAuthorSubject
	if err := q.All(&books); err != nil {
		log.Fatal("q.All: ", err)
	}

	for _, book := range books {
		log.Printf("Book %d:\t%s. %q on %s",
			book.BookID, book.Author.LastName, book.Book.Title, book.Subject.Subject)
	}
}
