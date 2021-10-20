package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(conn, "Connected")
		fmt.Fprintf(conn, "Hello Hello")
		fmt.Fprintln(conn, "Hi Honey")

		conn.Close()
	}

}
