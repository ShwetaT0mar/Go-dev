package main

import (
	"fmt"
	"os"
	"text/template"
)

type Relation struct {
	Noun   string
	Action string
}

func main() {
	tpl, err1 := template.ParseFiles("../bleh.gohtml")
	if err1 != nil {
		fmt.Println("ERROR", err1)
	}
	r := Relation{"Shibani's friend", "Fuck"}
	s := Relation{"Koko", "Kill"}
	t := Relation{"Foxy", "Marry"}
	u := Relation{"Shibani", "Fuck Harder"}
	stfu := []Relation{r, s, t, u}
	err := tpl.ExecuteTemplate(os.Stdout, "bleh.gohtml", stfu)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
