package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func GetConnection() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	dbNew, err := sql.Open("sqlite", "./example.db")
	if err != nil {
		return nil, err
	}

	db = dbNew

	return db, err
}

func CloseConnection() error {
	if db != nil {
		err := db.Close()
		db = nil
		return err
	}

	return nil
}
