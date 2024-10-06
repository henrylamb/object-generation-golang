package testingLite

import (
	"github.com/henrylamb/object-generation-golang/jsonSchema"
)

// TestConstructor separates code extraction and processing
func TestConstructor(assumption string, codePath string, model jsonSchema.ModelType, processor CodeProcessor) (*jsonSchema.Definition, string, error) {
	return processor.ProcessCode(assumption, codePath, model)
}
