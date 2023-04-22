package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", ":8080") // A Listener is a generic network listener for stream-oriented protocols.
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	errConn := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if errConn != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	for {
		bs, err1 := io.ReadAll(conn)
		if err1 != nil {
			log.Fatalln("Exiting")
			panic(err1)
			//continue
		}
		time.Sleep(3 * time.Second)
		fmt.Println(string(bs))
	}

}
