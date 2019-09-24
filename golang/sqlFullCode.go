package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	_ "github.com/lib/pq"
)

// Email type
type Email struct {

	// eg: steve.mcqueen@gmail.com
	Username string // steve.mcqueen
	Domain   string // gmail.com
	Valid    bool   // says the type is valid or not
}

// Person type
type Person struct {
	Name         string `json: "name"`
	EmailAddress Email  `json:"email"`
}

// Stringer implementation for Email
func (em *Email) String() string {
	if em.Valid {
		return fmt.Sprintf("%s@%s", em.Username, em.Domain)
	}

	return ""
}


//Scanner method for type Email
func (email *Email) Scan(value interface{}) error {

	// if value in the database is nil, create the email obj with Valid flag as false
	if value == nil {
		*email = Email{Valid: false}
		return nil
	}

	mid := strings.Index(value.(string), "@")                       //spilt the string in database by '@'
	username := value.(string)[:mid]                                // find username
	domain := value.(string)[mid+1:]                                //find domain
	*email = Email{Username: username, Domain: domain, Valid: true} // Create email obj pointer
	return nil
}

//Value method for type Email
func (email *Email) Value() (driver.Value, error) {
	if !email.Valid { // if email is not valid return nil (NULL)
		return nil, nil
	}
	return email.String(), nil //for a valid email returns the combined string
}

func main() {

	// Connect to database
  // insert your connect string here
	db, err := sql.Open("postgres", "<connectionString>")
	if err != nil {
		panic(err)
	}

	///////////////////////////////////////////////////////////////////////

	// read from database
	emailRead := Email{}
	rows := db.QueryRow("SELECT email FROM public.t_usertypes limit 1")

	err = rows.Scan(&emailRead)
	if err != nil {
		panic(err)
	}
	fmt.Println(emailRead)

	////////////////////////////////////////////////////////////////////////

	// insert email to database
	emailWrite := Email{Username: "charles.bronson", Domain: "gmail.com", Valid: true}

	query := `INSERT INTO public.t_usertypes (email) VALUES ($1) `
	_, err = db.Exec(query, &emailWrite)
	if err != nil {
		panic(err)
	}

}
