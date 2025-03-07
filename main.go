package main

import (
	"finance/database"
	"log"
	"net/http"
	"text/template"
)

func main() {
	database.Connect()

	// Подцепляем изображения
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("html/img"))))

	http.HandleFunc("/main/", mainPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/nav_menu.html")

	t.Execute(w, nil)
}
