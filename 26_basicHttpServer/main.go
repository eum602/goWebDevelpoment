package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read the request
	request(conn)

	//write a response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0]
			fmt.Println("*** METHOD is ***", m)
		}
		if ln == "" { //according to the http spec after the request line and a headers there is a blank line
			//headers are done
			break
		}
		i++
	}

}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello world</strong></body></html>`
	//http structure
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\n\r")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n") //blank line
	fmt.Fprintf(conn, body)   //body
}
