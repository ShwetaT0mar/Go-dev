package main

import (
	"fmt"
	"net/http"
)

type something string

func (s something) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Anything you like!!")
	return
}

func main() {
	var a something
	http.ListenAndServe(":8080", a)
}
