package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", all)
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("UUIDv4: %s\n", id)
		}

		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure:true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
