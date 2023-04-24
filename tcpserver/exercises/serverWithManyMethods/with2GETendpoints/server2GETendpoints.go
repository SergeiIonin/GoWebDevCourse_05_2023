package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/template"
)

type response struct {
	Title string
	H1    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func handle(conn net.Conn) {
	defer conn.Close()

	respond(conn)
}

func executeTemplate(tpl template.Template, resp response, filename string) int64 {
	f, _ := os.Create(filename) // it can be used later for caching
	errF := tpl.Execute(f, resp)
	if errF != nil {
		log.Fatalln("Error executing template, ", errF)
	}
	size, _ := f.Stat()
	return size.Size()
}

func getResponse(conn net.Conn, endpoint string) {
	switch endpoint {
	case "/hello":
		resp := response{
			Title: "Hello",
			H1:    "HEY!!!",
		}
		size := executeTemplate(*tpl, resp, "GET_hello_response.html")
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", size)
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		err := tpl.Execute(conn, resp)
		if err != nil {
			log.Fatalln("Error executing request, ", err)
		}
	case "/whatsup":
		resp := response{
			Title: "Whatsup",
			H1:    "whatsup!??",
		}
		size := executeTemplate(*tpl, resp, "GET_whatsup_response.html")
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", size)
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		err := tpl.Execute(conn, resp)
		if err != nil {
			log.Fatalln("Error executing request, ", err)
		}
	default:
		file, _ := os.Open("GET_unknown_response.html") // need to process possible error?
		stat, _ := file.Stat()
		size := stat.Size()
		fmt.Println("resp size, ", size)
		fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", size)
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		responseTpl, _ := template.ParseFiles("GET_unknown_response.html")
		responseTpl.Execute(conn, nil)
	}
}

func respond(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	var method string
	var uri string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		if i == 0 {
			line0 := strings.Fields(text)
			method = line0[0]
			uri = line0[1]
			fmt.Printf("HTTP METHOD %s\n", method)
			fmt.Printf("URI %s\n", uri)
		} else {
			fmt.Println(text)
		}
		i++
	}
	switch method {
	case "GET":
		getResponse(conn, uri)
	default:
		body := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
			"This endpoint is unknown or not implemented!" +
			`</h1>
		</body>
		</html>
	`)
		fmt.Fprint(conn, "HTTP/1.1 501 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}
}

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}

}
