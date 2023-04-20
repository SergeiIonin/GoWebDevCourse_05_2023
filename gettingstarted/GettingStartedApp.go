package main

import "fmt"

type person struct {
	fname string
	lname string
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

type secretAgent struct {
	person
	isLicensed bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says "Good morning!"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says "Where's the whisky?'"`)
}

func main() {

	p := person{
		"John",
		"Smith",
	}
	sa := secretAgent{
		person{
			"Sterling",
			"Archer",
		},
		true,
	}

	p.speak()
	sa.speak()
	sa.person.speak() // NB!
	fmt.Println("-----")
	saySomething(p)
	saySomething(sa)

}
