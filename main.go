package main

import (
	"fmt"
	"log"
	"net/http"
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
	fmt.Fprintf(w, "Привет!")
}
