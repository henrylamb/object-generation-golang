package client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

// GZipRequestSender is a request sender that compresses the request body using gzip
type GZipRequestSender struct{}

// NewGZipRequestSender initializes a new GZipRequestSender
func NewGZipRequestSender() *GZipRequestSender {
	return &GZipRequestSender{}
}

// SendRequestBody sends a gzip-compressed JSON request and returns a response
func (grs *GZipRequestSender) SendRequestBody(baseURL, token string, requestBody *RequestBody) (*http.Response, error) {
	url := baseURL + "/api/objectGen"

	// Serialize the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", err)
	}

	// Compress the JSON data using gzip
	var compressedData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedData)

	_, err = gzipWriter.Write(jsonData)
	if err != nil {
		gzipWriter.Close()
		return nil, fmt.Errorf("error writing to gzip writer: %v", err)
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing gzip writer: %v", err)
	}

	// Create an HTTP request with the compressed data
	req, err := http.NewRequest("POST", url, &compressedData)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers - including gzip encoding
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request and return the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
