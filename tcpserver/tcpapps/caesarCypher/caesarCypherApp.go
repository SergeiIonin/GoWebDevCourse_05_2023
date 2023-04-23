package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		line := strings.ToLower(sc.Text())
		bytes := []byte(line)
		bytesRotated := rot13(bytes)
		fmt.Fprintf(conn, "%s - %s\n\n", line, bytesRotated) // prints to the connection (e.g. telnet)
	}
}

func rot13(bs []byte) []byte {
	var bytesRotated = make([]byte, len(bs))
	for i, cur := range bs {
		if cur <= 109 {
			bytesRotated[i] = cur + 13
		} else {
			bytesRotated[i] = cur - 13
		}
	}
	return bytesRotated
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() // blocks until smbd is connected
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}
