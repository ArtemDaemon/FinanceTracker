package main

import (
	"log"
	"net/http"
	"text/template"
)

// Полезное: https://go.dev/doc/articles/wiki/

func main() {
	http.HandleFunc("/main/", mainPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/nav_menu.html")

	// Отправить HTML
	t.Execute(w)
}
