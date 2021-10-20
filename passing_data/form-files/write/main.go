package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	s := ""
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("ff")
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		defer f.Close()
		fmt.Println("file ", f, "file headers", h)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, "Going wrong", 501)
			return
		}
		s = string(bs)
		wr, err := os.Create(filepath.Join("./data/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = wr.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<html><form method="POST" encType="multipart/form-data"><input type="file" name="ff"><input type="submit"></form></html>`+s)
}
