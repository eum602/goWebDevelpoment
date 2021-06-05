package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string //initialized as the zero value
	fmt.Println(req.Method)

	//if this method is POST then handle this form submission
	if req.Method == http.MethodPost {
		//open the file
		f, h, err := req.FormFile("q") //method for catching a file, returns a "File" and a "FileReader".
		//where "File" implements the io.Reader interface!! among other methods
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(f) //needs an io.Reader and returns a byte slice ([]byte)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs) //simply convert the slice of bytes into a string only to finally send it back to the user
	}

	//enctype="multipart/form-data ---> indicates that user is uploading a file"
	//<input type="file" name="q"  ---> also the input type has to be a file

	//finally either way the method run this
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	`+s) //attaching the content at the end
}
