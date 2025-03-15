package main

import (
	"finance/internal/database"
	"finance/internal/utils"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	tmpls, err := template.ParseFiles("web/templates/layout.html", "web/templates/"+tmpl+".html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	tmpls.Execute(w, data)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	balanceInfo, err := database.GetTotalBalance()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"Title":               "Главная",
		"StartBalance":        utils.FormatCurrency(balanceInfo.StartBalance),
		"TotalIncome":         utils.FormatCurrency(balanceInfo.TotalIncome),
		"TotalExpense":        utils.FormatCurrency(balanceInfo.TotalExpense),
		"BalanceDelta":        utils.FormatCurrency(balanceInfo.BalanceDelta),
		"BalanceDeltaColor":   utils.GetColorClass(balanceInfo.BalanceDelta),
		"CurrentBalance":      utils.FormatCurrency(balanceInfo.CurrentBalance),
		"CurrentBalanceColor": utils.GetColorClass(balanceInfo.CurrentBalance),
	}
	renderTemplate(w, "main", data)
}

func transactionsPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Транзакции",
	}
	renderTemplate(w, "transactions", data)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Println("Не удалось открыть браузер:", err)
	}
}

func main() {
	database.InitDB()
	defer database.DB.Close()

	// Подцепляем изображения
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/static/img"))))

	http.HandleFunc("/main/", mainPage)
	http.HandleFunc("/transactions/", transactionsPage)

	url := "http://localhost:8080/main"
	go func() {
		openBrowser(url)
	}()
	fmt.Println("Сервер запущен на", url)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
