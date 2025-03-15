package main

import (
	"finance/database"
	"log"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	tmpls, err := template.ParseFiles("html/layout.html", "html/"+tmpl+".html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	tmpls.Execute(w, data)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Главная",
	}
	renderTemplate(w, "main", data)
}

func transactionsPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Транзакции",
	}
	renderTemplate(w, "transactions", data)
}

func main() {
	database.InitDB()
	defer database.DB.Close()

	// Подцепляем изображения
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("html/img"))))

	http.HandleFunc("/main/", mainPage)
	http.HandleFunc("/transactions/", transactionsPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
