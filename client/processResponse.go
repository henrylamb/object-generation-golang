package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ResponseProcessor responsible for processing the HTTP response
type ResponseProcessor struct{}

// NewResponseProcessor initializes a new ResponseProcessor
func NewResponseProcessor() ResponseProcessor {
	return ResponseProcessor{}
}

// ProcessResponse processes the response and returns the parsed Response struct
func (rp *ResponseProcessor) ProcessResponse(resp *http.Response) (*Response, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
