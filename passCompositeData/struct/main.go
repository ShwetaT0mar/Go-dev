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
	tpl, err1 := template.ParseFiles("../blah.gohtml")
	if err1 != nil {
		fmt.Println("ERROR", err1)
	}
	r := Relation{"Shibani's friend", "Fuck"}
	err := tpl.ExecuteTemplate(os.Stdout, "blah.gohtml", r)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
