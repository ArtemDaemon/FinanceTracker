package database

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "finances.db")
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(1)

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	err = runInitSQL()
	if err != nil {
		log.Fatal("Ошибка при инициализации структуры БД:", err)
	}
}

func runInitSQL() error {
	initSQL, err := os.ReadFile("database/sql/init.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(initSQL))
	return err
}

func GetTotalBalance() (float64, error) {
	query, err := os.ReadFile("database/sql/get_total_balance.sql")
	if err != nil {
		return 0, err
	}

	row := DB.QueryRow(string(query))
	var balance float64
	err = row.Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
