package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/old", old)
	http.HandleFunc("/new", new)
	http.HandleFunc("/", all)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)
}

func old(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/new", http.StatusPermanentRedirect)
}
func new(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method is:", r.Method)
	io.WriteString(w, "This is where you land.")
}

func all(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method is :", r.Method)
	io.WriteString(w, `<html><form action="/post" method="POST"><input type="submit"></form></html>`)
	io.WriteString(w, "This is the usual stuff.")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method is :", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}
