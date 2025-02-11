/*
Copyright © 2025 ArtemDaemon <artem.daemon.official@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "finance-tracker",
	Short: "Приложение для ведения финансов",
	Long: `Приложение позволяет вести бюджет и расходы. Пример:

Тестовый пример`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
