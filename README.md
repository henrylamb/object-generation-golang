# Object Generation Golang

This module provides a client implementation for sending JSON definitions via HTTP POST requests. It is designed to be simple and easy to integrate into your existing Go projects.

For additional examples of how this library works, please visit: [https://go-multiple.com/console](https://go-multiple.com/console).

## Installation

```go get github.com/henrylamb/object-generation-golang```

To use this library you will need to import it in the below format:

```go
import (
"github.com/henrylamb/object-generation-golang/jsonSchema"
)

```

## Guide: Using Go Client to Send JSON Definitions

This guide demonstrates how to create a Go client that sends JSON definitions using HTTP POST requests.

### Step 1: Define the Client Struct

Start by defining a `Client` struct that will manage the API connection.

```go
type Client struct {
	APIKey  string // APIKey is the authentication token for the API.
	BaseURL string // BaseURL is the base endpoint for API requests.
}
```

# Guide: Using Go Client to Send JSON Definitions

This guide demonstrates how to create a Go client that sends JSON definitions using HTTP POST requests.

### Step 1: Define the Client Struct

Start by defining a `Client` struct that will manage the API connection.

```go
type Client struct {
	APIKey  string // APIKey is the authentication token for the API.
	BaseURL string // BaseURL is the base endpoint for API requests.
}
```

### Step 2: Initialize a New Client

Create a constructor function `NewClient` to initialize a new client instance with the API key and base URL.

```go
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://example.com/api", // Replace with your API's base URL
	}
}
```

### Step 3: Define the SendRequest Method

Implement a method `SendRequest` on the `Client` struct to send a POST request with a JSON-encoded definition.

```go
func (c *Client) SendRequest(definition *Definition) (*http.Response, error) {
	url := c.BaseURL
	if definition.Req != nil {
		url = definition.Req.URL
	}

	jsonData, err := json.Marshal(definition)
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
```

### Step 4: Example Usage

Demonstrate how to use the `Client` to send a definition using a sample `Definition` struct.

```go
// Example usage
func ExampleUsage() {
	// Initialize a new client with your API key
	apiKey := "your-api-key"
	client := NewClient(apiKey)

	// Define a sample definition
	definition := &Definition{
		Type:        "Object",
		Instruction: "Sample instruction for the definition.",
		Req: &RequestFormat{
			URL: "https://api.example.com/process",
		},
		Properties: map[string]Definition{
			"property1": {Type: "String", Instruction: "Description of property1"},
			"property2": {Type: "Number", Instruction: "Description of property2"},
		},
	}

	// Send the request
	resp, err := client.SendRequest(definition)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Process response as needed
	fmt.Printf("Response Status: %s\n", resp.Status)
	// Additional processing of response body, headers, etc.
}
```

### Conclusion

This guide provides a structured approach to creating a Go client for sending JSON definitions via HTTP POST requests. Ensure to adapt the `Definition` struct and example usage to fit your specific API requirements and data structures.