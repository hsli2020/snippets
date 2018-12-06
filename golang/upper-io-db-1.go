package main

import (
	"context"
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
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

// Author represents an item from the "authors" table.
type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

// Subject represents an item from the "subjects" table.
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

	// The ctx variable is going to be passed to sess.Tx, you
	// can use any context.Context here. If you don't want to
	// use a special context you can also pass nil.
	ctx := context.Background()

	// sess.Tx requires a function, this function takes a
	// single sqlbuilder.Tx argument and returns an error. The
	// tx value is just like sess, except it lives within a
	// transaction.  If the function returns any error, the
	// transaction will be rolled back.
	err = sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
		// Anything you do here with the tx value will be part
		// of the transaction.
		cols, err := tx.Collections()
		if err != nil {
			return err
		}
		log.Printf("Cols: %#v", cols)

		// The booksTable value is valid only within the transaction.
		booksTable := tx.Collection("books")
		total, err := booksTable.Find().Count()
		if err != nil {
			return err
		}
		log.Printf("There are %d items in %s", total, booksTable.Name())

		var books []Book
		err = tx.SelectFrom("books").Limit(3).OrderBy(db.Raw("RANDOM()")).All(&books)
		if err != nil {
			return err
		}
		log.Printf("Books: %#v", books)

		res, err := tx.Query("SELECT * FROM books ORDER BY RANDOM() LIMIT 1")
		if err != nil {
			return err
		}

		var book Book
		err = sqlbuilder.NewIterator(res).One(&book)
		if err != nil {
			return err
		}
		log.Printf("Random book: %#v", book)

		// If the function returns no error the transaction is commited.
		return nil
	})

	if err != nil {
		log.Printf("sess.Tx: ", err)
	}
}

package main

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
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

// Author represents an item from the "authors" table.
type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

// Subject represents an item from the "subjects" table.
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

	var eaPoe Author

	// If you ever need to use raw SQL, use the Query,
	// QueryRow and Exec methods of sess:
	rows, err := sess.Query(`SELECT id, first_name, last_name FROM authors WHERE last_name = ?`, "Poe")
	if err != nil {
		log.Fatal("Query: ", err)
	}
	// This is a standard query that mimics the API from
	// database/sql.
	if !rows.Next() {
		log.Fatal("Expecting one row")
	}
	if err := rows.Scan(&eaPoe.ID, &eaPoe.FirstName, &eaPoe.LastName); err != nil {
		log.Fatal("Scan: ", err)
	}
	if err := rows.Close(); err != nil {
		log.Fatal("Close: ", err)
	}

	log.Printf("%#v", eaPoe)

	// Make sure you're using Exec or Query depending on the
	// specific situation.
	_, err = sess.Exec(`UPDATE authors SET first_name = ? WHERE id = ?`, "Edgar Allan", eaPoe.ID)
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}

	// The sqlbuilder package providers tools for working with
	// raw sql.Rows, such as the NewIterator function.
	rows, err = sess.Query(`SELECT * FROM books LIMIT 5`)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	// The NewIterator function takes a *sql.Rows value and
	// returns an iterator.
	iter := sqlbuilder.NewIterator(rows)

	// This iterator provides methods for iterating over data,
	// such as All, One, Next and friends.
	var books []Book
	if err := iter.All(&books); err != nil {
		log.Fatal("Query: ", err)
	}

	log.Printf("Books: %#v", books)
}

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

// Author represents an item from the "authors" table.
type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

// Subject represents an item from the "subjects" table.
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

	var eaPoe Author

	// Using sqlbuilder.Selector to get E .A. Poe from our
	// authors table.
	err = sess.SelectFrom("authors").
		Where("last_name", "Poe"). // Or Where("last_name = ?", "Poe")
		One(&eaPoe)
	if err != nil {
		log.Fatal("Query: ", err)
	}
	log.Printf("%#v", eaPoe)

	// The name says "Edgar Allen", let's fit it using
	// sqlbuilder.Updater:
	res, err := sess.Update("authors").
		Set("first_name = ?", "Edgar Allan"). // Or Set("first_name", "Edgar Allan").
		Where("id = ?", eaPoe.ID).            // Or Where("id", eaPoe.ID)
		Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}

	// Now let's create a new E. A. P. book.
	book := Book{
		Title:    "The Crow",
		AuthorID: eaPoe.ID,
	}
	res, err = sess.InsertInto("books").
		Values(book). // Or Columns(c1, c2, c2, ...).Values(v1, v2, v2, ...).
		Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}
	if res != nil {
		id, _ := res.LastInsertId()
		log.Printf("New book id: %d", id)
	}

	// Delete the book we just created (and any book with the
	// same name).
	q := sess.DeleteFrom("books").
		Where("title", "The Crow")
	log.Printf("Compiled query: %v", q)

	_, err = q.Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}
}
