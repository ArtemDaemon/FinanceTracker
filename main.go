package main

import (
	"finance/database"
	"log"
	"net/http"
	"text/template"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	// Подцепляем изображения
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("html/img"))))

	http.HandleFunc("/main/", mainPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/nav_menu.html")

	t.Execute(w, nil)
}
