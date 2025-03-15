package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

type BalanceInfo struct {
	StartBalance   float64
	TotalIncome    float64
	TotalExpense   float64
	BalanceDelta   float64
	CurrentBalance float64
}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "finances.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Соединение с базой данных установлено")
	DB.SetMaxOpenConns(1)

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	err = runInitSQL()
	if err != nil {
		log.Fatal("Ошибка при инициализации структуры БД:", err)
	}
	fmt.Println("Структура данных инициализирована")
}

func runInitSQL() error {
	initSQL, err := os.ReadFile("database/sql/init.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(initSQL))
	return err
}

func GetTotalBalance() (BalanceInfo, error) {
	query, err := os.ReadFile("database/sql/get_total_balance.sql")
	if err != nil {
		return BalanceInfo{}, err
	}

	row := DB.QueryRow(string(query))
	var info BalanceInfo
	err = row.Scan(
		&info.StartBalance,
		&info.TotalIncome,
		&info.TotalExpense,
		&info.BalanceDelta,
		&info.CurrentBalance,
	)
	if err != nil {
		return BalanceInfo{}, err
	}
	return info, nil
}
