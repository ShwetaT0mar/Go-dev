package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {

	v := r.FormValue("key")
	io.WriteString(w, `<html><form method="POST"><input type="text" name="key"><input type="submit"></form></html>`+v)
}
