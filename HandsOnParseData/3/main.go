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

type Item struct {
	Name, Desc string
	Price      float32
}

type Menu struct {
	Breakfast []Item
	Lunch     []Item
	Dinner    []Item
}

type Restaraunt []Menu

func main() {

	// menu := Menu{
	// 	Breakfast: []Item{Item{Name: "Dosa", Price: 19}, {Name: "Poha", Price: 25}, {Name: "Dahi", Price: 25}},
	// 	Lunch:     []Item{Item{Name: "RajmaChawal", Price: 35}, Item{Name: "ChholeChawal", Price: 40}},
	// 	Dinner:    []Item{Item{Name: "Daal Roti", Price: 28}, Item{Name: "Sabzi Roti", Price: 23}},
	// }

	rest := Restaraunt{
		Menu{
			Breakfast: []Item{Item{Name: "Dosa1", Price: 19}, {Name: "Poha", Price: 25}, {Name: "Dahi", Price: 25}},
			Lunch:     []Item{Item{Name: "RajmaChawal1", Price: 35}, Item{Name: "ChholeChawal", Price: 40}},
			Dinner:    []Item{Item{Name: "Daal Roti1", Price: 28}, Item{Name: "Sabzi Roti", Price: 23}},
		},
		Menu{
			Breakfast: []Item{Item{Name: "Dosa2", Price: 19}, {Name: "Poha", Price: 25}, {Name: "Dahi", Price: 25}},
			Lunch:     []Item{Item{Name: "RajmaChawal2", Price: 35}, Item{Name: "ChholeChawal", Price: 40}},
			Dinner:    []Item{Item{Name: "Daal Roti2", Price: 28}, Item{Name: "Sabzi Roti", Price: 23}},
		},
		Menu{
			Breakfast: []Item{Item{Name: "Dosa3", Price: 19}, {Name: "Poha", Price: 25}, {Name: "Dahi", Price: 25}},
			Lunch:     []Item{Item{Name: "RajmaChawal3", Price: 35}, Item{Name: "ChholeChawal", Price: 40}},
			Dinner:    []Item{Item{Name: "Daal Roti3", Price: 28}, Item{Name: "Sabzi Roti", Price: 23}},
		},
	}

	err := tpl.Execute(os.Stdout, rest)
	if err != nil {
		fmt.Println(err)
	}

}
