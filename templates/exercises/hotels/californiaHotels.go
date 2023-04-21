package main

import (
	"log"
	"os"
	"text/template"
)

type region string

// todo how can we enforce using only these constants for `Region`?
const Southern region = "Southern"
const Central region = "Central"
const Northern region = "Northern"

// type hotels []hotel is also possible
type hotels struct {
	List []hotel
}

type hotel struct {
	Name, Address, City string
	Region              region
	Zip                 int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotelsList := hotels{
		[]hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     95612,
				Region:  Southern,
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     95612,
				Region:  Southern,
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotelsList)
	if err != nil {
		log.Fatalln(err)
	}

}
