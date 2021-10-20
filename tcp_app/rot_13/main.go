package main

import (
	"bufio"
	"fmt"
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
			log.Println("Error accepting connection:", err)
		}
		handle(conn)
		conn.Close()
	}

}

func handle(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		b := []byte(ln)
		convert(b)
		fmt.Println(string(b))
	}

}

func convert(arr []byte) {
	//new := make([]byte, len(arr))
	for i, v := range arr {
		if v <= 109 {
			arr[i] = v + 13
		} else {
			arr[i] = v - 13
		}

	}
	//return arr
}
