package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
		handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	//respond(conn)
}

func request(conn net.Conn) {
	i := 0
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
		if i == 0 {
			method := strings.Fields(ln)[0]
			url := strings.Fields(ln)[1]
			if method == "GET" {
				handleGet(url, conn)
			}
			if method == "POST" {
				handlePost(url, conn)
			}
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, greetings string) {
	body := ``
	if greetings == "hola" {
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hola World</strong></body></html>`
	}

	if greetings == "hello" {
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func handleGet(url string, conn net.Conn) {
	pathFields := getPath(url)
	if pathFields[0] == "something" {
		if pathFields[1] == "hola" {
			respond(conn, "hola")
		}
		if pathFields[1] == "hello" {
			respond(conn, "hello")
		}

	}
}

func handlePost(url string, conn net.Conn) {

}
func getPath(url string) []string {
	path := ""
	arr := []string{}
	for i := 1; i <= len(url); i++ {
		if i == len(url) || string(url[i]) == "/" {
			fmt.Println("Good stuff")
			fmt.Println("Final Path is ::", path)
			fmt.Println("array is ", arr)
			arr = append(arr, path)
			path = ""
			if i == len(url) {
				break
			}
			continue
		}
		path = path + string(url[i])
	}
	return arr
}
