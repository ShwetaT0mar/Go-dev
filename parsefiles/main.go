package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("stuff.gohtml")
	if err != nil {
		fmt.Println("ERROR", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	err = tpl.Execute(nf, nil)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	tpl, err = tpl.ParseFiles("temp1.gohtml", "temp2.crap")
	if err != nil {
		fmt.Println("ERROR", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "temp1.gohtml", nil)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "temp2.crap", nil)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("ERROR", err)
	}

}
