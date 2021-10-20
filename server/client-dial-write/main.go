package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hey I called you!")
}
