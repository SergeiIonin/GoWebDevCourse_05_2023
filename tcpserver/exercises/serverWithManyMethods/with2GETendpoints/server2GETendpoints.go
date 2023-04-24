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

	//request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	i := 0
	for sc.Scan() {
		text := sc.Text()
		if text == "" {
			break
		}
		if i == 0 {
			line0 := strings.Fields(text)
			fmt.Printf("HTTP METHOD %s\n", line0[0])
			fmt.Printf("URI %s\n", line0[1])
		} else {
			fmt.Println(text)
		}
		i++
	}
}

func getResponse(conn net.Conn, endpoint string) {
	var body string
	switch endpoint {
	case "/hello":
		body = fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello!</title>
		</head>
		<body>
		<h1>` +
			"Hello!" +
			`</h1>
		</body>
		</html>
	`)
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	case "/whatsup":
		body = fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Whatsup?</title>
		</head>
		<body>
		<h1>` +
			"whatsup?" +
			`</h1>
		</body>
		</html>
	`)
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	default:
		body = fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Unknown</title>
		</head>
		<body>
		<h1>` +
			"This endpoint is unknown!" +
			`</h1>
		</body>
		</html>
	`)
		fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
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
