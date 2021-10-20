package main

import (
	"io"
	"net/http"
)

func all(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html><img src="/image"></html>`)
}

func image(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "stuff.go")
}

func main() {
	http.HandleFunc("/", all)
	http.HandleFunc("/image", image)
	http.ListenAndServe(":8080", nil)
}
