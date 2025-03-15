package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"finance/internal/models"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

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
	initSQL, err := os.ReadFile("internal/database/sql/init.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(initSQL))
	return err
}

func GetTotalBalance() (models.BalanceInfo, error) {
	query, err := os.ReadFile("internal/database/sql/get_total_balance.sql")
	if err != nil {
		return models.BalanceInfo{}, err
	}

	row := DB.QueryRow(string(query))
	var info models.BalanceInfo
	err = row.Scan(
		&info.StartBalance,
		&info.TotalIncome,
		&info.TotalExpense,
		&info.BalanceDelta,
		&info.CurrentBalance,
	)
	if err != nil {
		return models.BalanceInfo{}, err
	}
	return info, nil
}
