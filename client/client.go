package client

import (
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"net/http"
)

// Client responsible for making HTTP requests
type Client struct {
	Password   string
	BaseURL    string
	HttpClient HttpClient
}

// HttpClient interface to abstract HTTP operations
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewDefaultClient(password, url string) *Client {
	return NewClient(password, url, &http.Client{})
}

// NewClient initializes a new Client instance
func NewClient(password, url string, httpClient HttpClient) *Client {
	return &Client{
		Password:   password,
		BaseURL:    url,
		HttpClient: httpClient,
	}
}

// SendRequest sends the prompt and definition, and returns the parsed response
func (c *Client) SendRequest(prompt string, definition *jsonSchema.Definition) (*Response, error) {
	requestSender := NewRequestSender(c)
	responseProcessor := NewResponseProcessor()

	requestBody := &RequestBody{
		Prompt:     prompt,
		Definition: definition,
	}

	resp, err := requestSender.SendRequestBody(requestBody)
	if err != nil {
		return nil, err
	}

	return responseProcessor.ProcessResponse(resp)
}
