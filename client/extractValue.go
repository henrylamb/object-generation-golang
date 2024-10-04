package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Res defines a structure with a single string field
type Res struct {
	Value string                 `json:"value"`
	Other map[string]interface{} `json:"Other"`
}

// extractValue extracts the value from the HTTP response and returns a Res struct
func extractValue(resp *http.Response) (*Res, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var res Res
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return &res, nil
}
