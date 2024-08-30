package client

import "github.com/henrylamb/object-generation-golang/jsonSchema"

type ToolRequestBody struct {
	Prompt       string `json:"prompt"`
	SubFunctions []*jsonSchema.SubordinateFunction
}

// Create a response struct
type ToolResponse struct {
	Tool    *jsonSchema.SubordinateFunction
	UsdCost float64 `json:"usdCost"`
}
