package testingLite

import (
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
)

// TestConstructor separates code extraction and processing
func TestConstructor(assumption string, codePath string, model jsonSchema.ModelType, processor CodeProcessor) (*jsonSchema.Definition, string, error) {
	_, fileContents, err := extractLanguagesAndContents(codePath)
	if err != nil {
		return nil, "", fmt.Errorf("Error processing code path: %s", err.Error())
	}

	return processor.ProcessCode(assumption, fileContents, model)
}
