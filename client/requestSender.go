package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RequestSender responsible for creating and sending HTTP requests
type RequestSender struct {
	client *Client
}

// NewRequestSender initializes a new RequestSender
func NewRequestSender(client *Client) *RequestSender {
	return &RequestSender{client: client}
}

// SendRequestBody sends a JSON request and returns a response
func (rs *RequestSender) SendRequestBody(requestBody *RequestBody) (*http.Response, error) {
	url := rs.client.BaseURL + "/api/objectGen"

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling definition: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+rs.client.Password)

	resp, err := rs.client.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
