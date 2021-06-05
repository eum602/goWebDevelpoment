package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed string
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//body
	bs := make([]byte, req.ContentLength) //'make' -> means make 'that slice of bytes' ([]byte) which is the input, also we pass
	// the length which is the capacity.
	//if the second argument is passed to 'make' then -> 'len(bs)' is equal to 'cap(bs)'
	//if you set ii := make([]int, 10, 30) then -> 'len(ii)' equals to 10 but 'cap(ii)' equals to 30
	//beneath a slice is an array, in the example '30' is set to the size of that underlined array
	//while '10' is set as the size of the slice which is being initizalized
	//in the last example -> ii[0] will return '0'
	// if I set ii[25] will throw an error 'index out range' because the slice was only set to '10' even if the underlined array was set to '30'
	//if I set bs:= make([]string, 0, 300) then I can append with something like bs.append("some")
	//bs -> slice o strings
	//ss -> slice of strings
	//ii -> slice of ints ... and so on
	req.Body.Read(bs) //read the content and put into the byte slice 'bs'
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
