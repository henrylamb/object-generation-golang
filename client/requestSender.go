package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DefaultRequestSender is a default implementation of RequestSender
type DefaultRequestSender struct{
	client *http.Client
}

// NewDefaultRequestSender initializes a new DefaultRequestSender
func NewDefaultRequestSender(client *http.Client) *DefaultRequestSender {
	return &DefaultRequestSender{
		client: client,
	}
}

// SendRequestBody sends a JSON request and returns a response
func (rs *DefaultRequestSender) SendRequestBody(baseURL, token string, requestBody *RequestBody) (*http.Response, error) {
	url := baseURL + "/api/objectGen"

	// Serialize the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", err)
	}

	// Create an HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request and return the response
	resp, err := rs.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
