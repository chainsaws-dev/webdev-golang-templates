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

// Restaurant - описывает отдельный ресторан и его меню
type Restaurant struct {
	Name    string
	Address string
	City    string
	Tel     string
	Type    string
	Menu    Dishes
}

// Restaurants - Слайс ресторанов
type Restaurants []Restaurant

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

	r := Restaurants{
		{
			Name:    "Yoshi",
			Address: "4 Avenue de la Madone, 98000 Monaco",
			City:    "Monaco",
			Tel:     "+377 93 15 13 13",
			Type:    "Japanese restaurant",
			Menu:    d,
		},
		{
			Name:    "Le Louis XV-Alain Ducasse à l'Hôtel de Paris",
			Address: "Place du Casino, 98000 Monaco",
			City:    "Monaco",
			Tel:     "+377 98 06 88 64",
			Type:    "Fine dining restaurant",
			Menu:    d,
		},
	}

	err := tmp.ExecuteTemplate(os.Stdout, "rest.gohtml", r)

	if err != nil {
		log.Fatalln(err)
	}
}
