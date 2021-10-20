package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hndle int

func (h hndle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		ContentLength int64
		Host          string
		Form          map[string][]string
		Header        http.Header
	}{
		Method:        r.Method,
		URL:           r.URL,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		Form:          r.Form,
		Header:        r.Header,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var a hndle
	http.ListenAndServe(":8080", a)
}
