package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"net/http"
)

type Client struct {
	Password string
	BaseURL  string
}

// NewClient here the password that you would set in the request to ensure secure communication between servers. This value must be set as an environment variable as MULTIPLE_PASSWORD
func NewClient(password, url string) *Client {
	return &Client{
		Password: password,
		BaseURL:  url,
	}
}

func (c *Client) SendHttpRequest(prompt string, definition *jsonSchema.Definition) (*http.Response, error) {
	url := c.BaseURL

	requestBody := RequestBody{
		Prompt:     prompt,
		Definition: definition,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling definition: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/objectGen", url), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}

// Wrapper function to process the response and return the custom Response type
func (c *Client) SendRequest(prompt string, definition *jsonSchema.Definition) (*Response, error) {
	// Send the request
	resp, err := c.SendHttpRequest(prompt, definition)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Decode the response JSON into the Response struct
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
