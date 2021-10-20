package main

import (
	"log"
	"net/http"
	"text/template"
)

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}
func all(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, "Files not found", 404)
	}
	f := r.FormValue("fname")
	l := r.FormValue("lname")
	s := r.FormValue("sub") == "on"
	error := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if error != nil {
		http.Error(w, error.Error(), 500)
		log.Fatalln(error)
	}
}
