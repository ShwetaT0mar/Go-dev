package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	err := tpl.ExecuteTemplate(w, "index.gohtml", string(bs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
