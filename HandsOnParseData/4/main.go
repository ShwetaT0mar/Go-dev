package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
)

type Cell struct {
	Date string
	Open string
}
type AllData []Cell

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		fmt.Println(err)
	}
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	a := AllData{}
	for _, line := range data {
		fmt.Println("Date is :", line[0])
		a = append(a, Cell{Date: line[0], Open: line[1]})
	}
	err = tpl.Execute(os.Stdout, a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data)
}
