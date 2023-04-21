package main

import (
	"os"
	"text/template"
)

type item struct {
	Name, Description, Item string
	Price                   float64
}

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	items := items{
		item{
			Name:        "n1",
			Description: "d1",
			Item:        "i1",
			Price:       23.43,
		},
		item{
			Name:        "n2",
			Description: "d2",
			Item:        "i2",
			Price:       35.68,
		},
		item{
			Name:        "n3",
			Description: "d3",
			Item:        "i3",
			Price:       31.45,
		},
	}

	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", items)

}
