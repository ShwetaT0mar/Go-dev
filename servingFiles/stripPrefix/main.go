package main

import (
	"io"
	"net/http"
)

func all(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html><img src="/assets/stuff.png"</html>`)
}

func main() {
	http.HandleFunc("/", all)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
