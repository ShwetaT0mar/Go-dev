package main

import (
	"io"
	"net/http"
)

func imageReturns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<!DOCTYPE html><html><body>NOT FROM OUR SERVER<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg"></body></html>`)
}

func imageDoesNotReturn(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html><html><body>FROM OUR SERVER<img src="/stuff.png"></body></html>`)
}

func main() {
	http.HandleFunc("/", imageReturns)
	http.HandleFunc("/no", imageDoesNotReturn)
	http.ListenAndServe(":8080", nil)
}
