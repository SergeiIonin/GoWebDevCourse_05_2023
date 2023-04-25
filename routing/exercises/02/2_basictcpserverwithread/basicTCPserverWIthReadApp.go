package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	var resp string
	for sc.Scan() {
		text := sc.Text()
		fmt.Println("text = ", text)
		if text == "" {
			break
		}
		resp = resp + "\n" + text
	}
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

		go handle(conn)

		fmt.Println("Code got here.") // if we start another connection, then goroutine closes the prev. one and we'll see this line, but not the next!
		io.WriteString(conn, "I see you connected.")

	}
}
