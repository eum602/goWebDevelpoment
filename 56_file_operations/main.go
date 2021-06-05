package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if _, err := os.Stat("./user"); os.IsNotExist(err) {
		err = os.Mkdir("./user", 0777)
		if err != nil {
			log.Fatalln("Error while creating a 'user' directory", err)
		}
	}
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

		dst, err := os.Create(filepath.Join("./user/", h.Filename)) //Create a file, it states: "From the current directory
		//go into the user folder and then in there store a file with a name that comes from 'h.Filename'"
		//os.Create -> creates "a pointer to a file"

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//1. enctype="multipart/form-data ---> indicates that user is uploading a file". Use this whenever you use any
	//kind of '<input type="file">' in your form

	//2. enctype="application/x-www-form-urlencoded" ---> this is the default enctype. If no enctype is pointed in a form
	//then the this default is automatically set. This kind "application/w-x-www-.." sets key value pairs separated by ampersands

	//3. enctype="text/plain" ---> Not reliable when interpreted by the computer, only use it for debugging but NO for production

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
