package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func ServeMe(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "YOUR NAME IS==")
	fmt.Fprintf(w, "Shweta")
	io.WriteString(w, "Tomar")
}

func ServeDoggy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DOGGY DOGGY")
}

func ServeAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GOD LOVES ALL")
}

func main() {
	http.Handle("/", http.HandlerFunc(ServeAll))
	http.Handle("/dog/", http.HandlerFunc(ServeDoggy))
	http.Handle("/me/", http.HandlerFunc(ServeMe))

	http.ListenAndServe(":8080", nil)
}
