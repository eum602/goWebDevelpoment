package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Run the server of lesson 22 before running this client ...")
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "I dialed you")
	fmt.Println("server called ...")
}
