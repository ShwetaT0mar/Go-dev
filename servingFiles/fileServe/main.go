package main

import (
	"io"
	"net/http"
)

func image(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html><img src="stuff.png"></html>`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/screen", image)
	http.ListenAndServe(":8080", nil)
}
