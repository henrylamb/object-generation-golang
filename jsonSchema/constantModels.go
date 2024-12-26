package jsonSchema

type DataType string

const (
	Object  DataType = "object"
	Number  DataType = "number"
	Integer DataType = "integer"
	String  DataType = "string"
	Array   DataType = "array"
	Null    DataType = "null"
	Boolean DataType = "boolean"
	Map     DataType = "map"
	Byte    DataType = "byte" //this will be used for the audio and image data selection (if this is selected as byte then either Image or Audio must not be nil, if it is then nothing will occur and an empty byte will be returned. The same is true if both are filled.
)

type ModelType string

const (
	Gpt3              ModelType = "Gpt3"
	Gpt4              ModelType = "Gpt4"
	ClaudeSonnet      ModelType = "ClaudeSonnet"
	ClaudeHaiku       ModelType = "ClaudeHaiku"
	Llama70b          ModelType = "Llama70b"
	Gpt4Mini          ModelType = "Gpt4Mini"
	Llama405b         ModelType = "Llama405"
	Llama8b           ModelType = "Llama8b"
	O1                ModelType = "o1-preview"
	O1Mini            ModelType = "o1-mini"
	GeminiFlash       ModelType = "GeminiFlash"
	GeminiFlash8B     ModelType = "GeminiFlash8B"
	GeminiPro         ModelType = "GeminiPro"
	Llama8bInstant    ModelType = "Llama8bInstant"
	Llama70bVersatile ModelType = "Llama70bVersatile"
	Llama1B           ModelType = "Llama1B"
	Llama3B           ModelType = "Llama3B"
	Default           ModelType = "Default"
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
