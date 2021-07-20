package main;

// Representation of an http request
type HTTPRequest struct {
	RequestMethod, RequestPath, RequestHTTPVersion string;
	RequestHeaders map[string][]string;
	RequestBody []byte;
}

func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{
		RequestHeaders: make(map[string][]string),
		RequestBody: make([]byte, 0),
	}
}