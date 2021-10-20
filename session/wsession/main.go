package main

import (
	"fmt"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type UserData struct {
	UserName string
	First    string
	Last     string
}

var users = map[string]UserData{}
var sessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("Error while getting sid")
		if err == http.ErrNoCookie {
			sid, _ := uuid.NewV4()
			c = &http.Cookie{
				Name:  "session_id",
				Value: sid.String(),
			}
			fmt.Println("Creating sid")
			http.SetCookie(w, c)
		} else {
			return
		}
	}

	var u UserData
	if un, ok := sessions[c.Value]; ok {
		u = users[un]
	}
	if r.Method == http.MethodPost {
		fmt.Println("Processing Post Data")
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = UserData{un, f, l}
		fmt.Println("Saving data")
		sessions[c.Value] = un
		fmt.Println("Session and uName stored")
		fmt.Println(sessions[c.Value])
		users[un] = u

	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := sessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u, ok := users[un]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
