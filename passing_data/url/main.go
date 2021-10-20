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
	v := r.FormValue("q")
	u := r.FormValue("r")
	io.WriteString(w, v)
	io.WriteString(w, u)
}
