package client

import (
	"testing"

	"github.com/henrylamb/object-generation-golang/jsonSchema"
)

func TestGZipClient(t *testing.T) {
	// Create a new GZip client
	gzipClient := NewGZipClient("test-password", "https://api.example.com")

	// Verify that the client was created with the correct configuration
	if gzipClient.Password != "test-password" {
		t.Errorf("Expected password 'test-password', got '%s'", gzipClient.Password)
	}

	if gzipClient.BaseURL != "https://api.example.com" {
		t.Errorf("Expected BaseURL 'https://api.example.com', got '%s'", gzipClient.BaseURL)
	}

	// Verify that the RequestSender is of type GZipRequestSender
	if _, ok := gzipClient.RequestSender.(*GZipRequestSender); !ok {
		t.Errorf("Expected RequestSender to be of type *GZipRequestSender, got %T", gzipClient.RequestSender)
	}
}

func TestGZipClientSendRequest(t *testing.T) {
	// This test would require a mock server to fully test the gzip compression
	// For now, we'll just test that the method exists and can be called
	gzipClient := NewGZipClient("test-password", "https://api.example.com")

	// Create a simple definition for testing
	definition := &jsonSchema.Definition{
		Type: jsonSchema.Object,
		Properties: map[string]jsonSchema.Definition{
			"name": {
				Type: jsonSchema.String,
			},
		},
	}

	// This would fail in a real scenario without a server, but we're just testing the interface
	_, err := gzipClient.SendRequest("Generate a person object", definition)

	// We expect an error here since there's no actual server
	if err == nil {
		t.Error("Expected an error when calling SendRequest without a server")
	}
}

func TestGZipRequestSenderCreation(t *testing.T) {
	sender := NewGZipRequestSender()

	if sender == nil {
		t.Error("Expected NewGZipRequestSender to return a non-nil instance")
	}
}
