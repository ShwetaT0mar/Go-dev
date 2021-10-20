package main

import (
	"bufio"
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
			log.Println(err)
		}
		sc := bufio.NewScanner(conn)
		for sc.Scan() {
			ln := sc.Text()
			fmt.Println(ln)
			io.WriteString(conn, sc.Text())
			if ln == "" {
				fmt.Println("Headers are over")
				break
			}
		}
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
