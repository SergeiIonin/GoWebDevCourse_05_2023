package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

func writeHostname() string {
	host, _ := os.Hostname()
	msg := "Hostname is" + host + "\n"
	return msg
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/persons", persons)
	http.ListenAndServe(":80", nil)
}

func persons(w http.ResponseWriter, req *http.Request) {
	s := writeHostname()
	if len(os.Args) < 1 {
		http.Error(w, "datasource is not passed", http.StatusBadRequest)
		return
	}
	datasource := os.Args[1]
	db := connect(datasource)

	rows, err := db.Query(`select first_name from persons;`)
	check(err)
	defer rows.Close()

	var name string
	s += "\nRETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func connect(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Can't connect to datasource", err.Error())
		return nil
	}
	return db
}
