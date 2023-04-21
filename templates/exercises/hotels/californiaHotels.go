package main

type region string

const Southern region = "Southern"
const Central region = "Central"
const Northern region = "Northern"

type hotels struct {
	List   []hotel
	Region region
}

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func main() {

}
