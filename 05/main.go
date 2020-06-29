package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.ParseGlob("./*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	file, err := os.Open("table.csv")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "forex.gohtml", records)

	if err != nil {
		log.Fatalln(err)
	}
}
