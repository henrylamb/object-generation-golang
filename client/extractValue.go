package client

import (
	"encoding/json"
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"io/ioutil"
	"net/http"
)

// Res defines a structure with a single string field
type Res struct {
	Value string                 `json:"value"`
	Other map[string]interface{} `json:"Other"`
}

// ExtractValue extracts the value from the HTTP response and returns a Res struct
func ExtractValue(resp *http.Response) (*Res, error) {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var res Res
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return &res, nil
}

// ExtractValue extracts the value from the HTTP response and returns a Res struct
func ExtractTool(resp *http.Response) (*jsonSchema.SubordinateFunction, error) {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var res *jsonSchema.SubordinateFunction
	if err := json.Unmarshal(body, res); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return res, nil
}
