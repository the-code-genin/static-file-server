package main

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

// Parse an incoming HTTP request
func ParseHTTPRequest(conn *net.Conn) (*HTTPRequest, error) {
	// Get the initial request parameters
	scanner := bufio.NewScanner(*conn)
	payload := make([]string, 0)
	for scanner.Scan() {
		buffer := scanner.Text()
		if buffer == "" {
			break
		}

		payload = append(payload, buffer)
	}

	// Create a new parsed HTTP request instance
	parsedHTTPRequest := NewHTTPRequest()


	// Parse the first line of the request
	firstLine := strings.SplitN(payload[0], " ", 3)
	if len(firstLine) != 3 {
		return parsedHTTPRequest, errors.New("invalid HTTP request")
	}


	// Extract the request method and path
	parsedHTTPRequest.RequestMethod = firstLine[0]
	parsedHTTPRequest.RequestPath = firstLine[1]


	// validate the request HTTP version
	requestHTTPVersion := strings.SplitN(firstLine[2], "/", 2)
	if len(requestHTTPVersion) != 2 {
		return parsedHTTPRequest, errors.New("invalid HTTP version")
	} else if requestHTTPVersion[0] != "HTTP" {
		return parsedHTTPRequest, errors.New("this server only supports HTTP")
	}
	parsedHTTPRequest.RequestHTTPVersion = requestHTTPVersion[1]


	// Parse the request headers
	for i := 1; i < len(payload); i++ {
		line := strings.SplitN(payload[i], ": ", 2)
		if len(line) != 2 {
			return parsedHTTPRequest, errors.New("incorrectly formatted headers")
		}

		parsedHTTPRequest.RequestHeaders[line[0]] = append(parsedHTTPRequest.RequestHeaders[line[0]], line[1])
	}


	// Return the parsed request
	return parsedHTTPRequest, nil
}