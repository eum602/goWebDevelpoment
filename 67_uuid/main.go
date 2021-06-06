package main

import (
	"fmt"
	"github.com/satori/go.uuid" //the folder name is not the same as the package name
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			//Secure: true, // makes https connection as a must
			HttpOnly: true, //means can be accessed with javascript
		}
		http.SetCookie(res, cookie) //sets a cookie in the browser
	}

	fmt.Println(cookie)
}
