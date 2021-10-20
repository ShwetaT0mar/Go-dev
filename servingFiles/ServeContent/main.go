package main

import (
	"io"
	"net/http"
	"os"
)

func all(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html><img src="/image"</html>`)
}

func img(w http.ResponseWriter, r *http.Request) {
	im, err := os.Open("stuff.png")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	details, err := im.Stat()
	if err != nil {
		http.Error(w, "Image details not found", 504)
	}
	http.ServeContent(w, r, im.Name(), details.ModTime(), im)
}

func main() {
	http.HandleFunc("/", all)
	http.HandleFunc("/image", img)
	http.ListenAndServe(":8080", nil)
}
