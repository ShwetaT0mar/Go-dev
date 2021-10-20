package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, "File Not found", 404)
		return
	}

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)
	err = tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
