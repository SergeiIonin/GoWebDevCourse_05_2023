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

	request(conn)

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
		fmt.Println(text)
		if i == 0 {
			method := strings.Fields(text)
			fmt.Printf("HTTP METHOD %s\n", method[0])
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

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
