package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func serve(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		text := sc.Text()
		fmt.Println("text = ", text)
		if text == "" {
			fmt.Println("breaking...")
			break
		}
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
		go serve(conn)
	}
}
