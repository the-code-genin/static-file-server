package main

import (
	"fmt"
	"net"
)

// Handles an HTTP connection
func HandleConnection(conn *net.Conn) {
	// Close connection gracefully at the end of the operation
	defer (*conn).Close()

	// Parse the incoming request from the connection
	parsedHTTPRequest, err := ParseHTTPRequest(conn)
	if err != nil {
		response := NewHTTPResponse()
		response.ResponseHTTPVersion = parsedHTTPRequest.RequestHTTPVersion
		response.ResponseCode = "500"
		response.ResponseStatus = "ERROR"
		response.ResponseBody = []byte("An error occured!")
		response.ResponseHeaders["Content-Type"] = []string{"text/plain"}
		response.ResponseHeaders["Content-Length"] = []string{fmt.Sprintf("%v", len(response.ResponseBody))}
		SendHTTPResponse(conn, response)
		return
	}

	// Parse the headers from the subsequent lines
	fmt.Printf("%s: %s\n", parsedHTTPRequest.RequestMethod, parsedHTTPRequest.RequestPath)

	// Return the response
	response := NewHTTPResponse()
	response.ResponseHTTPVersion = parsedHTTPRequest.RequestHTTPVersion
	response.ResponseCode = "200"
	response.ResponseStatus = "OK"
	response.ResponseBody = []byte("Hello world!")
	response.ResponseHeaders["Content-Type"] = []string{"text/plain"}
	response.ResponseHeaders["Content-Length"] = []string{fmt.Sprintf("%v", len(response.ResponseBody))}
	SendHTTPResponse(conn, response)
}
