package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

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

		sc := bufio.NewScanner(conn)
		var resp string
		for sc.Scan() {
			text := sc.Text()
			fmt.Println("text = ", text)
			if text == "" {
				fmt.Println("breaking...")
				break
			}
			resp = resp + "\n" + text
		}
		fmt.Println("out of the loop")

		fmt.Println("Code got here.") // after some time the goroutine closes conn and we'll see this line, but not the next!
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
