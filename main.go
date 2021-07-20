package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on the port
	sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Response
	fmt.Println("Server running on http://127.0.0.1:8080/")

	// Start listening to requests
	for {
		// Accept a connection
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection concurrently
		go HandleConnection(&conn)
	}
}
