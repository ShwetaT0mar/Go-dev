package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"text/template"
)

var fn = template.FuncMap{
	"up": strings.ToUpper,
	"sb": substring,
}

var maths = template.FuncMap{
	"rt": math.Sqrt,
	"sq": square,
}

func square(i float64) float64 {
	return i * i
}

func substring(s string) string {
	return strings.TrimSpace(s)[:3]
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(maths).Funcs(fn).ParseFiles("stuff.gohtml"))
}

type Monk struct {
	Name  string
	Level float64
}

func main() {
	a := Monk{"Shweta", 0}
	b := Monk{"Siddharth", 1}
	c := Monk{"Hymn", 2.7889}
	arr := []Monk{a, b, c}
	err := tpl.ExecuteTemplate(os.Stdout, "stuff.gohtml", arr)
	if err != nil {
		fmt.Println(err)
	}
}
