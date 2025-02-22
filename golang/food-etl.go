// https://github.com/353words/food
package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jszwec/csvutil"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schemaSQL string

//go:embed insert.sql
var insertSQL string

type Row struct {
	Business   string    `csv:"businessname" db:"business_name"`
	Licstatus  string    `csv:"licstatus" db:"license_status"`
	Result     string    `csv:"result" db:"result"`
	Violdesc   string    `csv:"violdesc" db:"description"`
	Violdttm   time.Time `csv:"violdttm" db:"time"`
	Violstatus string    `csv:"violstatus" db:"status"`
	Viollevel  string    `csv:"viollevel" db:"-"`
	Level      int       `db:"level"`
	Comments   string    `csv:"comments" db:"comments"`
	Address    string    `csv:"address" db:"address"`
	City       string    `csv:"city" db:"city"`
	Zip        string    `csv:"zip" db:"zip"`
}

func unmarshalTime(data []byte, t *time.Time) error {
	var err error
	*t, err = time.Parse("2006-01-02 15:04:05", string(data))
	return err
}

func parseLevel(value string) int {
	switch value {
	case "*":
		return 1
	case "**":
		return 2
	case "***":
		return 3
	}

	return -1
}

func ETL(csvFile io.Reader, tx *sqlx.Tx) (int, int, error) {
	r := csv.NewReader(csvFile)
	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		return 0, 0, err
	}
	dec.Register(unmarshalTime)
	numRecords := 0
	numErrors := 0

	for {
		numRecords++
		var row Row
		err = dec.Decode(&row)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error: %d: %s", numRecords, err)
			numErrors++
			continue
		}
		row.Level = parseLevel(row.Viollevel)
		if _, err := tx.NamedExec(insertSQL, &row); err != nil {
			return 0, 0, err
		}
	}

	return numRecords, numErrors, nil
}

func main() {
	file, err := os.Open("boston-food.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db, err := sqlx.Open("sqlite3", "./food.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(schemaSQL); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	numRecords, numErrors, err := ETL(file, tx)
	duration := time.Since(start)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	frac := float64(numErrors) / float64(numRecords)
	if frac > 0.1 {
		tx.Rollback()
		log.Fatalf("too many errors: %d/%d = %f", numErrors, numRecords, frac)
	}
	tx.Commit()
	fmt.Printf("%d records (%.2f errors) in %v\n", numRecords, frac, duration)
}

// insert.sql
INSERT INTO violations (
    business_name, license_status, result, description, time, status,
    level, comments, address, city, zip
) VALUES (
    :business_name, :license_status, :result, :description, :time, :status,
    :level, :comments, :address, :city, :zip
);

// query.sql
SELECT business_name, COUNT(business_name) as num_violations
FROM violations
WHERE license_status = 'Active' AND time >= '2016-01-01'
GROUP BY business_name
ORDER BY num_violations DESC
LIMIT 20

// schema.sql
CREATE TABLE IF NOT EXISTS violations (
    business_name TEXT,
    license_status TEXT,
    result TEXT,
    description TEXT,
    time TIMESTAMP,
    status TEXT,
    level INTEGER,
    comments TEXT,
    address TEXT,
    city TEXT,
    zip TEXT
);
