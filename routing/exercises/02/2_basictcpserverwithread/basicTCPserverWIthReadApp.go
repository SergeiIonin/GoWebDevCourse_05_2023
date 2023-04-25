package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func serve(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	var method string
	var uri string
	i := 0
	for sc.Scan() {
		text := sc.Text()
		if i == 0 {
			parts := strings.Fields(text)
			method = parts[0]
			uri = parts[1]
		}
		fmt.Println("text = ", text)
		if text == "" {
			fmt.Println("breaking...")
			break
		}
		i++
	}
	switch method {
	case "GET":
		if uri == "/" {
			index(conn)
		} else if uri == "/apply" {
			apply(conn)
		} else {
			serverError(conn)
		}
	case "POST":
		if uri == "/apply" {
			applyProcess(conn)
		} else {
			serverError(conn)
		}
	default:
		serverError(conn)
	}
}

func serverError(conn net.Conn) {
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

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<h1>Apply now!</h1>
	<a href="/">index</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(conn)
	}
}
