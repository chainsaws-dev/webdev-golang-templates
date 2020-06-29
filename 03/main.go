// Package main - практика создания структур и шаблонов с нуля
package main

import (
	"log"
	"os"
	"text/template"
)

// Dish - хранит сведения об одном блюде
type Dish struct {
	Name    string
	Descrip string
	Meal    string
	Price   float64
}

// Dishes - несколько блюд
type Dishes []Dish

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.ParseGlob("./*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	d := Dishes{
		Dish{
			Name:    "Fish Soup",
			Descrip: "Scandinavian delight",
			Meal:    "Lunch",
			Price:   3.22,
		},
		Dish{
			Name:    "Porrige",
			Descrip: "Looks like nuke",
			Meal:    "Breakfast",
			Price:   1.22,
		},
	}

	err := tmp.ExecuteTemplate(os.Stdout, "menu.gohtml", d)

	if err != nil {
		log.Fatalln(err)
	}
}
