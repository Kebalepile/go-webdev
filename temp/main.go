package main

import (
	// "fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	tpl *template.Template
	// FuncMap register go functions to use in template view
	// func can be custom or already built in
	fm = template.FuncMap{
		"drive": drive,
		"uc":    strings.ToUpper,
		"date":  date,
	}
)

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func drive(x string) string {
	return x + ",  started..."
}

func init() {
	// new undefined template is needed when adding func map
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	a := sage{
		Name:  "Buddha",
		Motto: "The blief of no beliefs.",
	}

	b := sage{
		Name:  "Gandhi",
		Motto: "Pimpin ain't easy.",
	}

	c := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	d := car{
		Manufacturer: "Nissan",
		Model:        "TIDA",
		Doors:        4,
	}

	data := struct {
		Wisdom    []sage
		Transport []car
		Nako      time.Time
	}{
		Wisdom:    []sage{a, b},
		Transport: []car{c, d},
		Nako:      time.Now(),
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln("occurde while executing template", err)
	}
}

func date(t time.Time) string {
	// return t.Format("01-02-2006") //m-d-y
	return t.Format("2006/01/02") //y/m/d
}
