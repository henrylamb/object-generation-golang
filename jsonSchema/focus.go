package jsonSchema

// Focus the idea for this is so that when a narrow focus request needs to be sent out to an LLM without needing all the additional information. From prior generation.
type Focus struct {
	Prompt string `json:"prompt"`
	//the fields value denotes the properties that will be extracted from the properties fields. These will only operate at a single level of generation.
	//the order in which the fields that are listed will be the order for which the currently generated information will be presented below the prompt value.
	Fields []string `json:"fields"`

	//KeepOriginal -- for keeping the original prompt in cases for lists where it would otherwise be removed from the context
	KeepOriginal bool `json:"keepOriginal,omitempty"`
}
