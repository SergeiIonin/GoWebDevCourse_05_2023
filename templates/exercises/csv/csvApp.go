package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

type records []record

func parseCSV(filePath string) records {
	src, err1 := os.Open(filePath)
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer src.Close()

	csvFile := csv.NewReader(src)
	rawRecs, err2 := csvFile.ReadAll()
	if err2 != nil {
		log.Fatalln(err1)
	}

	recs := make(records, 0, len(rawRecs))

	for i, row := range rawRecs {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		record := record{date, open}
		recs = append(recs, record)
	}

	return recs
}

//var tpl *template.Template

func foo(rw http.ResponseWriter, req *http.Request) { // req *http.Request is required so that func will be handler func(ResponseWriter, *Request)
	records := parseCSV("table.csv")

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(rw, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
