// Package main - практика создания структур и шаблонов с нуля
package main

import (
	"log"
	"os"
	"text/template"
)

// Hotel - хранит сведения об одном отеле
type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

// Hotels - несколько отелей
type Hotels []Hotel

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.ParseGlob("./*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	h := Hotels{
		Hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		Hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}

	err := tmp.ExecuteTemplate(os.Stdout, "htplt.gohtml", h)

	if err != nil {
		log.Fatalln(err)
	}
}
