package main

import (
	"html/template"
	"net/http"
)

func main() {
	fs := http.StripPrefix("/resources", http.FileServer(http.Dir("./public")))
	http.Handle("/resources/", fs)
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
