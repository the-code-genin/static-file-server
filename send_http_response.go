package main

import (
	"bufio"
	"fmt"
	"net"
)

// Send an HTTP Response
func SendHTTPResponse(conn *net.Conn, response *HTTPResponse) error {
	writer := bufio.NewWriter(*conn)
	defer writer.Flush()

	// Write the first line
	writer.WriteString(
		fmt.Sprintf(
			"HTTP/%s %s %s",
			response.ResponseHTTPVersion,
			response.ResponseCode,
			response.ResponseStatus,
		),
	)
	writer.WriteString("\r\n")


	// Write each header to the stream
	for key, values := range response.ResponseHeaders {
		for _, value := range values {
			writer.WriteString(
				fmt.Sprintf(
					"%s: %s",
					key,
					value,
				),
			)
			writer.WriteString("\r\n")
		}
	}
	writer.WriteString("\r\n")
	writer.WriteString("\r\n")


	// Write the body
	writer.Write(response.ResponseBody)
	writer.WriteString("\r\n")


	return nil
}
