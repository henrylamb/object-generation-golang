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
)

type ModelType string

const (
	Gpt3         ModelType = "Gpt3"
	Gpt4         ModelType = "Gpt4"
	ClaudeSonnet ModelType = "ClaudeSonnet"
	ClaudeHaiku  ModelType = "ClaudeHaiku"
	Llama70b     ModelType = "Llama70b"
	Gpt4Mini     ModelType = "Gpt4Mini"
	Llama405b    ModelType = "Llama405"
	Llama8b      ModelType = "Llama8b"
	Default      ModelType = "Default"
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
