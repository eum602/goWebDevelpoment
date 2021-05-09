package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var counter = 0

func main() {
	li, err := net.Listen("tcp", ":8080") //1. Listen
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept() //2. accept
		if err != nil {
			log.Println(err)
		}
		go handle(conn) //3. handle the connection -> launch a goroutine
	}
}

func handle(conn net.Conn) {
	counter++
	cnt := counter
	fmt.Println("counter ", cnt)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { //scanner.Scan returns a bool -> so it scans as long as there are LINES to read.
		ln := scanner.Text() //give me the line
		fmt.Println("room ", cnt, ": ", ln)
	}

	defer conn.Close()

	//never gets here because the connection is constantly opened, but if I open a connection with telnet and I
	//suddenly close it then the code reaches here
	fmt.Println("code got here")

}
