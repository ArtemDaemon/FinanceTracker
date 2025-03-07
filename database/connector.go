package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() {
	db, err := sql.Open("sqlite3", "finances.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
