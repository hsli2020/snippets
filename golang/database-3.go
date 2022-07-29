package database

import (
	"database/sql"
	"time"
)

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

// ConnectSQL creates database pool for postgres
func ConnectSQL(dsn string) (*sql.DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	return db, nil
}

// NewDatabase opens connection with specified database
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
