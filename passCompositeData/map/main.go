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
	stree := map[string]string{"Fuck": "RoRo", "kill": "Shibani", "Marry": "Foxy"}
	err := tpl.ExecuteTemplate(os.Stdout, "blah.gohtml", stree)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
