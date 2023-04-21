package main

import (
	"log"
	"os"
	"text/template"
)

// todo how can we enforce using only these constants for `Region`?
const Southern string = "Southern"
const Central string = "Central"
const Northern string = "Northern"

// type hotels []hotel is also possible
type hotels struct {
	List []hotel
}

type region struct {
	Region string
	Hotels []hotel
}

type hotel struct {
	Name, Address, City string
	Zip                 int
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml", "tpl_regions.gohtml"))
}

func main() {
	r := region{
		Region: Southern,
		Hotels: []hotel{
			{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     95612,
			},
			{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     95612,
			},
		},
	}

	regions := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
	}

	err1 := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	err2 := tpl.ExecuteTemplate(os.Stdout, "tpl_regions.gohtml", regions)
	if err1 != nil {
		log.Fatalln(err1)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}

}
