package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	//"io"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

func main() {
	db := OpenDatabase()

	filename := "test.csv"
	columns := []int{-1, 0, 1, -2}
	table := "ttt"

	FastImport(db, table, filename, columns)
}

func FastImport(db *sql.DB, table, filename string, columns []int) error {
	records, err := ReadDataFile(filename)
	if err != nil {
		return err
	}
	tmpfile := PickFields(records, columns)
	err = BatchInsert(db, table, tmpfile)
	os.Remove(tmpfile)
	return err
}

func ReadDataFile(filename string) ([][]string, error) {
	// Open File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create CSV Reader
	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1
	reader.Comma = ','
	reader.LazyQuotes = true

	return reader.ReadAll()
}

func PickFields(records [][]string, columns []int) string {
	rows := make([][]string, len(records))

	now := time.Now().Format("2006-01-02 15:04:05")

	for r, rec := range records {
		fields := make([]string, len(columns))
		for i, col := range columns {
			if col < 0 {
				if col == -1 {
					fields[i] = ""
				} else if col == -2 {
					fields[i] = now
				}
			} else {
				fields[i] = rec[col]
			}
		}
		rows[r] = fields
	}

	// C:/Users/Andy/AppData/Local/Temp/DDHHMMSS.csv
	filename := strings.Replace(os.TempDir(), "\\", "/", -1) + "/" + time.Now().Format("02030405")
	file, err := os.Create(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = '\t'
	w.WriteAll(rows)

	return filename
}

func BatchInsert(db *sql.DB, table, filename string) error {
	sql := fmt.Sprintf(`
       LOAD DATA INFILE '%s'
       IGNORE INTO TABLE %s
       FIELDS TERMINATED BY '\t'
       ENCLOSED BY '\"'
       LINES TERMINATED BY '\n'
       IGNORE 1 ROWS;
   `, filename, table)

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

var db *sql.DB

func OpenDatabase() *sql.DB {
	if db != nil {
		return db
	}

	var driver = "mysql"
	var host = "localhost"
	var port = "3306"
	var username = "root"
	var dbname = "test"

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, host, port, dbname)

	var err error

	db, err = sql.Open(driver, dsn)
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.Ping()
	if err != nil {
		panic("Error on connecting to database")
	}

	return db
}
