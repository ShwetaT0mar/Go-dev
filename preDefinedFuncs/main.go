package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("pdf.gohtml"))
}

type Data struct {
	Person string
	Cool   bool
}

func main() {
	//arr := []string{"a", "b", "c"}
	// str := struct {
	// 	Name string
	// 	Age  []int
	// }{"Shweta", []int{0, 4, 10, 14, 17, 20, 24, 26}}

	a := Data{"", true}
	b := Data{"Engineer", true}
	c := Data{"Designers", false}
	Star := []Data{a, b, c}
	err := tpl.ExecuteTemplate(os.Stdout, "pdf.gohtml", Star)
	if err != nil {
		fmt.Println(err)
	}
}
