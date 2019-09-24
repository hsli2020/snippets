package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Article struct
type Article struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	PubDate NullTime   `json:"pub_date"`
	Body    NullString `json:"body"`
	User    NullInt64  `json:"user"`
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
// func (ni *NullInt64) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ni.Int64)
// 	ni.Valid = (err == nil)
// 	return err
// }

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
// func (nb *NullBool) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nb.Bool)
// 	nb.Valid = (err == nil)
// 	return err
// }

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
// func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nf.Float64)
// 	nf.Valid = (err == nil)
// 	return err
// }

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
// func (ns *NullString) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ns.String)
// 	ns.Valid = (err == nil)
// 	return err
// }

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
// func (nt *NullTime) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nt.Time)
// 	nt.Valid = (err == nil)
// 	return err
// }

// MAIN program starts here
func main() {
	db, err := sql.Open("mysql", "user:pass@/test?charset=utf8")
	if err != nil {
		fmt.Println("database could not opened!!!!")
		fmt.Println(err.Error())
		return
	}

	// read articles
	rows, err := db.Query("SELECT * FROM Article")
	if err != nil {
		fmt.Println("Query failed.....")
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var a Article
		if err = rows.Scan(&a.ID, &a.Title, &a.PubDate, &a.Body, &a.User); err != nil {
			fmt.Println("Scanning failed.....")
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Article Instance := %#v\n", a)
		articleJSON, err := json.Marshal(&a)
		if err != nil {
			fmt.Errorf("Error while marshalling json: %s", err.Error())
			fmt.Println(err.Error())
			return
		} else {
			fmt.Printf("JSON Marshal := %s\n\n", articleJSON)
		}
	}

	db.Close()
}
