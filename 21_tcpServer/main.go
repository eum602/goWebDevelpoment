package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080") //1. listen
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() //2. accept
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server \n") //3. read and write in the connection
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well I hope")

		conn.Close()
	}
}

//start with go run and then make telnet to the 8080 port
