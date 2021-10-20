package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", all)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func all(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Println(err)
	}

	tpl.Execute(w, "index.gohtml")
}
