package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

func handle(conn net.Conn) { //User TELNET to test it
	err := conn.SetDeadline(time.Now().Add(10 * time.Second)) //every connection will only remain 10 seconds!
	if err != nil {
		log.Println(err)
	}

	counter++
	cnt := counter
	fmt.Println("counter ", cnt)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { //scanner.Scan returns a bool -> so it scans as long as there are LINES to read.
		ln := scanner.Text() //give me the line
		fmt.Println("room ", cnt, ": ", ln)
		fmt.Fprintf(conn, "I hear from the server %s\n", ln) //echoes the input to the issuer.
	}

	defer conn.Close()

	fmt.Println("code got here") //the code gets either because the client closes the connection or the SetDeadline reaches the configured time

}
