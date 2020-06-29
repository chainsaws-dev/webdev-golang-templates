package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tmp *template.Template

var fm = template.FuncMap{
	"convdate":  convertdate,
	"convfloat": convertfloat,
}

func convertdate(s string) time.Time {
	res, err := time.Parse("2006-01-02", s)

	if err != nil {
		log.Println(err)
	}

	return res
}

func convertfloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)

	if err != nil {
		log.Println(err)
	}

	return res
}

func init() {

	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("forex.gohtml"))

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
