package jsonSchema

import (
	"encoding/json"
)

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

// Definition is a struct for describing a JSON Schema.
// It is fairly limited, and you may have better luck using a third-party library.
type Definition struct {
	// Type specifies the data type of the schema.
	Type DataType `json:"type,omitempty"`
	// Instruction is the instruction for what to generate.
	Instruction string `json:"instruction,omitempty"`
	// Properties describes the properties of an object, if the schema type is Object.
	Properties map[string]Definition `json:"properties"`
	// Required specifies which properties are required, if the schema type is Object.
	Required []string `json:"required,omitempty"`
	// Items specifies which data type an array contains, if the schema type is Array.
	Items *Definition `json:"items,omitempty"`
	// Model
	Model ModelType `json:"model,omitempty"`
	// ProcessingOrder this is the order of strings ie the fields of the parent property keys that need to be processed first before this field is processed
	ProcessingOrder []string `json:"processingOrder,omitempty"`

	// SystemPrompt allows the developer to spefificy their own system prompt so the processing. It operates current at the properties level.
	SystemPrompt *string `json:"systemPrompt,omitempty"`
	// Req allows a developer to send out a request at a given point to ensure that additional information can be extracted from external databases
	Req *RequestFormat
	// NarrowFocus
	NarrowFocus *Focus
	// ImprovementProcess --> so that the user can speficy when a super high quality completion is needed and it can be improved upon
	ImprovementProcess bool `json:"improvementProcess,omitempty"`
	//TODO all the below fields are in BETA
	// SelectFields has the aim of being able to select multiple pieice of information and when they are all present then continue with processing. Such that the selection of information can work like so:
	/*
		The system works as an absolute path that has to be selected. So starting from the top most object then down to the selected field(s)
		"car.color" --> this would fetch the information from the car field and then the color field.
		"cars.color" --> Would return the entire list of colours that have been generated so far
	*/
	SelectFields []string `json:"selectFields,omitempty"`

	// Choices For determining which of the property fields should be generated
	Choices *Choices `json:"choices,omitempty"`

	// Voters this is used for determining whether or not you want to have voters determine the qulaity of completions. Increases costs but improves quality. If avialible to your tier then turned on automatically.
	Voters bool `json:"voters,omitempty"`

	//Map is used here as so that a map of values can be created and then returned -- useful in the instruction creation process -- not sure how useful it is otherwise
	HashMap *HashMap
}

type Choices struct {
	Number  int      `json:"number,omitempty"`  //this denotes the number of choices that should be selected
	Options []string `json:"options,omitempty"` //this is the list of fields that will be chosen from
	/*
		How this works is that it needs to be in a of definitions which match with the properties field. From the properties fields the choice of those keys will be selected
		the information of what the overall object, the properties being selected along with the instruction and their type and the types that they contain if the object goes down further.
		the prompt will also be pass in so that the agent can make the best decesion possible

		Once the choices have been selected the choices that haven't been selected will be deleted from the remaining keys avialible in both the ordered and unordered keys.
	*/
}

// HashMap this can output a map of values and so whilst it may take up a single field it could output many fields
type HashMap struct {
	KeyInstruction  string      `json:"keyInstruction,omitempty"`
	FieldDefinition *Definition `json:"fieldDefinition,omitempty"`
}

func (d Definition) MarshalJSON() ([]byte, error) {
	if d.Properties == nil {
		d.Properties = make(map[string]Definition)
	}
	type Alias Definition
	return json.Marshal(struct {
		Alias
	}{
		Alias: (Alias)(d),
	})
}

// ToMap converts the Definition struct to a map representation
func (d Definition) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if d.Type != "" {
		result["type"] = d.Type
	}
	if d.Instruction != "" {
		result["instruction"] = d.Instruction
	}
	if d.Properties != nil && len(d.Properties) > 0 {
		propertiesMap := make(map[string]interface{})
		for key, value := range d.Properties {
			propertiesMap[key] = value.ToMap()
		}
		result["properties"] = propertiesMap
	}
	if len(d.Required) > 0 {
		result["required"] = d.Required
	}
	if d.Items != nil {
		result["items"] = d.Items.ToMap()
	}
	if d.Model != "" {
		result["model"] = d.Model
	}
	if len(d.ProcessingOrder) > 0 {
		result["processingOrder"] = d.ProcessingOrder
	}

	return result
}
