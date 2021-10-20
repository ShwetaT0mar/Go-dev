package main

import (
	"io"
	"net/http"
	"os"
)

func apic(w http.ResponseWriter, r *http.Request) {
	img, err := os.Open("stuff.png")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	defer img.Close()
	io.Copy(w, img)

}
func a(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html><img src="doggy"</html>`)
}
func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/doggy", apic)
	http.ListenAndServe(":8080", nil)
}
