package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
	Name, Address, City, Zip string
}

type CaliHotel struct {
	Region string
	Hotel  []hotel
}

type calihotels []CaliHotel

func main() {

	a := CaliHotel{
		Region: "NORTH",
		Hotel:  []hotel{hotel{Name: "surya"}, hotel{Name: "Saya"}, hotel{Name: "Woods"}, hotel{Name: "Bhopa"}, hotel{Name: "Jhansi"}, hotel{Name: "Orchha"}, hotel{Name: "Delhi"}, hotel{Name: "Orchha2"}},
	}
	b := CaliHotel{
		Region: "SOUTH",
		Hotel:  []hotel{hotel{Name: "GreenGlen"}, hotel{Name: "OYOLIFE"}, hotel{Name: "OYOLIFE2"}, hotel{Name: "INDRANAGAR"}},
	}

	// c := calihotels{CaliHotel{
	// 	Region: "NORTH",
	// 	Hotel:  []hotel{{Name: "surya"}, {Name: "Saya"}, {Name: "Woods"}, {Name: "Bhopa"}, {Name: "Jhansi"}, {Name: "Orchha"}, {Name: "Delhi"}, {Name: "Orchha2"}},
	// },
	// 	CaliHotel{
	// 		Region: "SOUTH",
	// 		Hotel:  []hotel{{Name: "GreenGlen"}, {Name: "OYOLIFE"}, {Name: "OYOLIFE2"}, {Name: "INDRANAGAR"}},
	// 	},
	// }
	d := calihotels{a, b}
	err := tpl.Execute(os.Stdout, d)
	if err != nil {
		fmt.Println("error")
	}

	// type ab []string
	// abc := ab{"hey", "uo"}
	// fmt.Println{abc}

}
