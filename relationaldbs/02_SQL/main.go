package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:hellogo_@tcp(helloservice1db.ciyuywmlqmty.us-east-1.rds.amazonaws.com:3306)/hello?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/persons", persons)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func persons(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select first_name from persons;`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Println(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`create table customers (name varchar(20));`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customers", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`insert into customers values ("James");`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select * from customers;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`update customers set name="Jimmy" where name="James";`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`delete from customers where name="Jimmy";`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
