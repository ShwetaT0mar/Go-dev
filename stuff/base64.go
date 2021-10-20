package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	str := "Let's encode this str to base 64, \"and add more stuff steps 1) use base64 package then call stdencoding on the package then inside stndEncoding call the func encide to string, so basically base64.StdEncoding.EncodeToString(string) is the base 64 encoded value if you need something else like say decoding for example again base64.stdrncoding.DecodeString(string) this will return bytes "
	s := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(s)
	ab, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln("Some issues", err)
	}
	fmt.Println(string(ab))

}
