package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1] //[0] is the program you execute, and [1] is the argument you enter.
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lan="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body> 
		</html>
		`)

	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

//run it with -> go run man.go eum602
