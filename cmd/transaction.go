/*
Copyright © 2025 ArtemDaemon <artem.daemon.official@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Работа с транзакциями",
	Long: `Операции для работы с транзакциями. Примеры:

Тестовый пример`,
}

func init() {
	rootCmd.AddCommand(transactionCmd)
}
