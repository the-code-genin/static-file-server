package main;

// Representation of an http response
type HTTPResponse struct {
	ResponseHTTPVersion, ResponseCode, ResponseStatus string;
	ResponseHeaders map[string][]string;
	ResponseBody []byte;
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		ResponseHeaders: make(map[string][]string),
		ResponseBody: make([]byte, 0),
	}
}