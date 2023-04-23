package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		line := sc.Text()
		fields := strings.Fields(line)

		if len(fields) < 1 {
			continue
		}

		if len(fields) < 2 {
			fmt.Fprintf(conn, "at least key should be passed!")
			continue
		}

		command := fields[0]
		k := fields[1]
		switch command {
		case "SET":
			if len(fields) != 3 {
				fmt.Fprintf(conn, "value is not passed with SET command\r\n")
				continue
			}
			v := fields[2]
			data[k] = v
		case "GET":
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "DEL":
			delete(data, k)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND %s\r\n", command)
			continue
		}
	}
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
