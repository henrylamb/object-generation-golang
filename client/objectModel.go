package client

import "github.com/henrylamb/object-generation-golang/jsonSchema"

type RequestBody struct {
	Prompt     string                 `json:"prompt"`
	Definition *jsonSchema.Definition `json:"definition"`
}

// Create a response struct
type Response struct {
	Data    []byte  `json:"data"` //this data can then be marshalled into the apprioate object type.
	UsdCost float64 `json:"usdCost"`
}
