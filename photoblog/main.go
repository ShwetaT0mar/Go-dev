package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	filepath "path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/pics/", http.StripPrefix("/pics", http.FileServer(http.Dir("./pics"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fName := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "pics", fName)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = appendValue(w, c, fName)
	}
	xs := strings.Split(c.Value, "|")
	//fmt.Println("Let's see what we got", xs)
	tpl.ExecuteTemplate(w, "index.gohtml", xs)

}

func appendValue(w http.ResponseWriter, c *http.Cookie, path string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, path) {
		s += "|" + path
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
