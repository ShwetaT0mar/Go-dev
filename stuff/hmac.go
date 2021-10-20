package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	str := "A big string"
	fmt.Println(hash(str))
	str = "A bigger string"
	fmt.Println(hash(str))

}

func hash(s string) string {
	h := hmac.New(sha256.New, []byte("KEY"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
