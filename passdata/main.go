package main

import (
	"fmt"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("data/*"))
}

func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "blah.gohtml", "Money is the anthem; To success")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
