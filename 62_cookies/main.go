package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{ //using a composite literal to create a cookie
		Name:  "my-cookie-name",
		Value: "this values is abc",
	})
	fmt.Fprintln(w, "A cookie has been written, check your browser")
	fmt.Fprintln(w, "in chrome go to inspect / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your cookie #1 is: ", c1)

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your cookie #2 is: ", c2)

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your cookie #3 is: ", c3)
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "a cookie for general section",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "a cookie for a specific section",
	})
	fmt.Fprintln(w, "A cookie has been written, check your browser")
	fmt.Fprintln(w, "in chrome go to inspect / application / cookies")
}
