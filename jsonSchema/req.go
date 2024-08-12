package jsonSchema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HTTPMethod string

// Constants for HTTP methods
const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

// RequestFormat defines the structure of the request
type RequestFormat struct {
	URL           string                 `json:"url"`
	Method        HTTPMethod             `json:"method"`
	Headers       map[string]string      `json:"headers,omitempty"`
	Body          map[string]interface{} `json:"body,omitempty"`
	Authorization string                 `json:"authorization,omitempty"`
	RequireFields []string               `json:"requirFields,omitempty"`
}

// ExecuteRequest executes an HTTP request based on the RequestFormat
func (d *Definition) ExecuteRequest(currentGen map[string]any) (*http.Response, error) {
	// Merge currentGen into the existing body
	if d.Req.Body == nil {
		d.Req.Body = make(map[string]interface{})
	}
	for key, value := range currentGen {
		d.Req.Body[key] = value
	}

	// Marshal the body to JSON
	body, err := json.Marshal(d.Req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest(string(d.Req.Method), d.Req.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range d.Req.Headers {
		req.Header.Set(key, value)
	}
	// Set authorization
	req.Header.Set("Authorization", d.Req.Authorization)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

func SendRequest(def *Definition, currentGen map[string]any) *Res {
	//the request that is sent out needs to send out a current map of the generated object.

	request, err := def.ExecuteRequest(currentGen)
	if err != nil {
		log.Println("failed to execute request", err)
		return nil
	}

	value, err := ExtractValue(request)
	if err != nil {
		log.Println("failed to extract value", err)
		return nil
	}

	return value
}
