package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
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
		go serve(conn)
		//conn.Close()
	}
}

func serve(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	i := 0
	method := ""
	url := ""
	for sc.Scan() {
		ln := sc.Text()
		if i == 0 {
			fields := strings.Fields(ln)
			method = fields[0]
			url = fields[1]
			fmt.Println("Method: ", method)
			fmt.Println("URL: ", url)
			i++
		}
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("Headers are over")
			break
		}
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	//fmt.Fprintf(conn, "Content-Length: %d\r\r", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	tpl.Execute(conn, "index.gohtml")
	//io.WriteString(conn, body)

}
