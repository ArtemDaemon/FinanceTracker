package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite", "finances.db")
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(1)

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
