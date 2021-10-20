package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	text := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<title>Hello World!</title></head>
	<body>
	<h1>` + name + `</h1></body>
	</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println("ERROR WHILE CREATING THE FILE", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(text))
	fmt.Println(text)
}
