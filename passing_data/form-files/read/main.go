package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	s := ""
	fmt.Println("=============================", r.Method)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("ff")
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		defer f.Close()
		fmt.Println("FILE :", f, "Header :", h, "Error :", err)
		read, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, "Error reading the file", http.StatusInternalServerError)
		}
		s = string(read)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<html><form method="POST" encType="multipart/form-data"><input type ="file" name="ff"> <input type="submit"></form></html>`+s)
}
