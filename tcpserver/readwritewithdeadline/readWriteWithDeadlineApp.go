package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		fmt.Fprintf(conn, "I heard you say: %s\n", text)
	}

	defer conn.Close()
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}

func main() {
	li, err := net.Listen("tcp", ":8080") // A Listener is a generic network listener for stream-oriented protocols.
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err1 := li.Accept()
		if err1 != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}
