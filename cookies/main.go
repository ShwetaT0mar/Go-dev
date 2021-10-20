package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var crowd int
var visits int

func init() {
	crowd = 1
	visits = 0
}
func main() {
	http.HandleFunc("/consumer", consumer)
	http.HandleFunc("/creep", creep)
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func consumer(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "Consumer", Value: "1"})
	io.WriteString(w, "You are going to give us your money")
}
func creep(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "Creep", Value: "1"})
	io.WriteString(w, "You are giving us leverage, please don't be so easy")
}
func all(w http.ResponseWriter, r *http.Request) {
	v, err := r.Cookie("visitor_id")
	if err != nil {
		if err == http.ErrNoCookie {
			v = &http.Cookie{
				Name:  "visitor_id",
				Value: strconv.Itoa(crowd) + ":1",
			}
			crowd++
		} else {
			fmt.Println(err.Error())
		}
	} else {
		i := strings.Index(v.Value, ":")
		runes := []rune(v.Value[i+1:])
		visits, err = strconv.Atoi(string(runes))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		v.Value = v.Value[:i+1] + strconv.Itoa(visits+1)

	}
	http.SetCookie(w, v)
	io.WriteString(w, "No of visits :"+strconv.Itoa(visits))
	for i, co := range r.Cookies() {
		fmt.Println("Cookie no", i, "is :", co.String())
	}
}
