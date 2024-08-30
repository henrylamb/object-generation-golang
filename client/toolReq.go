package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"net/http"
)

func (c *Client) SendHttpToolRequest(prompt string, subFunctions []*jsonSchema.SubordinateFunction) (*http.Response, error) {
	url := c.BaseURL

	requestBody := ToolRequestBody{
		Prompt:       prompt,
		SubFunctions: subFunctions,
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
func (c *Client) SendToolRequest(prompt string, subFunctions []*jsonSchema.SubordinateFunction) (*jsonSchema.SubordinateFunction, error) {
	// Send the request
	resp, err := c.SendHttpToolRequest(prompt, subFunctions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Decode the response JSON into the Response struct
	var response jsonSchema.SubordinateFunction
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
