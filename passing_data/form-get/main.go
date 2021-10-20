package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", all)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("key")
	io.WriteString(w, `<html><form method="GET"><input type="text" name="key"><input type="submit"></form></html>`+v)
}
