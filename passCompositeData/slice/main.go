package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../data/*"))
}

func main() {
	stree := []string{"RoRo", "Shibani", "Foxy"}
	err := tpl.ExecuteTemplate(os.Stdout, "blah.gohtml", stree)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
