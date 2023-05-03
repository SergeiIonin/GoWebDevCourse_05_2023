package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
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
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func persons(w http.ResponseWriter, req *http.Request) {
	s := writeHostname()
	datasource := "admin:hellogo_@tcp(helloservice1db.ciyuywmlqmty.us-east-1.rds.amazonaws.com:3306)/hello?charset=utf8"
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
	fmt.Println("Connecting to db...")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Can't connect to datasource", err.Error())
		return nil
	}
	return db
}
