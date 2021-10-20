package main

import (
	"fmt"
	"net/http"
)

type hndle int

func (h hndle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Shweta-Code", "Abhi")
	w.Header().Set("Content-Type", "Text/html")
	fmt.Fprintf(w, "Anything you like")
}

func main() {
	var a hndle
	http.ListenAndServe(":8080", a)
}
