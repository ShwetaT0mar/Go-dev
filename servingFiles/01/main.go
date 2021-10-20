package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "FOO RAN")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Println(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", "This is for dog")

}

func toby(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}

func main() {
	//http.Handle("/get", http.StripPrefix("/get", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)
}
