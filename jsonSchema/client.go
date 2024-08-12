package jsonSchema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIKey  string
	BaseURL string
}

type RequestBody struct {
	Prompt     string      `json:"prompt"`
	Definition *Definition `json:"definition"`
}

// Create a response struct
type Response struct {
	Data    []byte  `json:"data"` //this data can then be marshalled into the apprioate object type.
	UsdCost float64 `json:"usdCost"`
}

func NewClient(apiKey, url string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: url,
	}
}

func (c *Client) SendHttpRequest(prompt string, definition *Definition) (*http.Response, error) {
	url := c.BaseURL
	if definition.Req != nil {
		url = definition.Req.URL
	}

	requestBody := RequestBody{
		Prompt:     prompt,
		Definition: definition,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling definition: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}

// Wrapper function to process the response and return the custom Response type
func (c *Client) SendRequest(prompt string, definition *Definition) (*Response, error) {
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
