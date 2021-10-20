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
type Family struct {
	Noun string
	Link string
}

type Chart struct {
	Rchart []Relation
	Fchart []Family
}

func main() {
	tpl, err1 := template.ParseFiles("../b.gohtml")
	if err1 != nil {
		fmt.Println("ERROR", err1)
	}
	r := Relation{"Shibani's friend", "Fuck"}
	s := Relation{"Koko", "Kill"}
	t := Relation{"Foxy", "Marry"}
	u := Relation{"Shibani", "Fuck Harder"}
	stfu := []Relation{r, s, t, u}

	m := Family{"Seema", "mom"}
	d := Family{"Rony", "dad"}
	sd := Family{"Bobby", "stepDad"}
	fam := []Family{m, d, sd}

	item := Chart{stfu, fam}
	err := tpl.ExecuteTemplate(os.Stdout, "b.gohtml", item)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
