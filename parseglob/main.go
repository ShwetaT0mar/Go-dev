package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("error", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "temp1.stuff", nil)
	if err != nil {
		fmt.Println("error", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "temp2.stuu", nil)
	if err != nil {
		fmt.Println("error", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "temp3.stuff", nil)
	if err != nil {
		fmt.Println("error", err)
	}
}
