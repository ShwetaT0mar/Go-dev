package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Firstname string
	LastName  string
	Email     string
	password  []byte
}
type session struct {
	username     string
	lastActivity time.Time
}

var tpl *template.Template
var users = map[string]UserData{}
var sessions = map[string]session{}
var sessionsCleaned time.Time
var shouldSignUp bool

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	sessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	//check if user already exists
	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//process form data
	if r.Method == http.MethodPost {
		var u UserData
		fName := r.FormValue("fname")
		lName := r.FormValue("lname")
		email := r.FormValue("email")
		//encode password
		pass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("pass")), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), 501)
		}

		if _, ok := users[email]; ok {
			http.Error(w, "Email already used", http.StatusForbidden)
			return
		}
		//create session
		sid, _ := uuid.NewV4()
		http.SetCookie(w, &http.Cookie{
			Name:  "session_ID",
			Value: sid.String(),
		})
		//store user data
		sessions[sid.String()] = session{email, time.Now()}
		u = UserData{fName, lName, email, pass}
		users[email] = u
		fmt.Println("Data added to users", users[email])
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", shouldSignUp)

}

func login(w http.ResponseWriter, r *http.Request) {

	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		uName := r.FormValue("uName")
		pass := r.FormValue("pass")
		fmt.Println("UserName is::", uName)
		//check if userName exists
		user, ok := users[uName]
		fmt.Println("Does it exist?", ok)
		if !ok {
			shouldSignUp = true
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}
		//check if the passwords match
		err := bcrypt.CompareHashAndPassword(user.password, []byte(pass))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		//create a session
		sid, _ := uuid.NewV4()
		http.SetCookie(w, &http.Cookie{
			Name:  "session_ID",
			Value: sid.String(),
		})
		sessions[sid.String()] = session{uName, time.Now()}
		fmt.Println("Setting the cookie")
		fmt.Println("Trying to redirect")
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
func logout(w http.ResponseWriter, r *http.Request) {
	//check if already logged in
	if !AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//delete session
	c, _ := r.Cookie("session_ID")
	delete(sessions, c.Value)

	c = &http.Cookie{
		Name:   "session_ID",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(sessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func bar(w http.ResponseWriter, r *http.Request) {
	//check if already logged in
	if !AlreadyLoggedIn(r) {
		fmt.Println("Not logged in redirecting to Index")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//get user Data
	c, _ := r.Cookie("session_ID")
	un := sessions[c.Value]
	u := users[un.username]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}

func AlreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session_ID")
	if err != nil {
		return false
	}
	un := sessions[c.Value]
	_, ok := users[un.username]
	fmt.Println("Logged in?", ok)
	return ok
}

func cleanSessions() {
	for k, s := range sessions {
		if s.lastActivity.Sub(time.Now()) > (time.Second * 30) {
			delete(sessions, k)
		}
	}
}
