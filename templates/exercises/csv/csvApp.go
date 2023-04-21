package main

import (
	"encoding/csv"
	"log"
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

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	records := parseCSV("table.csv")
	tpl.Execute(os.Stdout, records)
}
