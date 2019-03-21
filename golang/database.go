package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/**
 * usage:
 *   config := map[string]string{
 *       "host":     "",	// default to localhost
 *       "port":     "",	// default to 3306
 *       "username": "",	// default to root
 *       "password": "",	// default to empty
 *       "dbname":   "",	// mandatory
 *   }
 *	 db := database.GetConnection(config)
 * or
 *	 db := database.GetConnection(nil)
 */
func GetConnection(config map[string]string) *sql.DB {
	if db != nil {
		return db
	}

	var driver = "mysql"
	var host = "localhost"
	var port = "3306"
	var username = "root"
	var password string
	var dbname string

	if s, ok := config["driver"]; ok {
		driver = s
	}
	if s, ok := config["host"]; ok {
		host = s
	}
	if s, ok := config["port"]; ok {
		port = s
	}
	if s, ok := config["username"]; ok {
		password = s
	}
	if s, ok := config["dbname"]; ok {
		dbname = s
	}

	var dsn string

	if password == "" {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, host, port, dbname)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, dbname)
	}

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

func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
