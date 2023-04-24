package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	respond(conn)
}

func respond(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	var host string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		if i == 1 {
			words := strings.Fields(text)
			host = words[1]
		}
		i++
	}
	body := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		host +
		`</h1>
		</body>
		</html>
	`)

	// unless we add any of the following 3 lines, then the browser will reject the response since it's not adhering HTTP
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")

	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
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
