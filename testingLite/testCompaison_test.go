package testingLite

import (
	"encoding/json"
	"github.com/henrylamb/object-generation-golang/client"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"os"
	"testing"
)

func TestTestComparison(t *testing.T) {
	// Create a temporary directory
	c := client.NewDefaultClient(os.Getenv("MULTIPLE_PASSWORD"), "http://localhost:2008")

	processor := NewDefaultCodeProcessor()
	codePath := "./testComparison.go"

	definition, code, err := TestConstructor(WorkingAssumption, codePath, jsonSchema.Gpt4Mini, processor)
	if err != nil {
		t.Errorf("Error constructing test: %v", err)
	}

	response, err := c.SendRequest(code, definition)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	//unmarshal the response
	testVal := &CodeTest{}
	err = json.Unmarshal(response.Data, testVal)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	//Compare the values
	if !TestComparison(testVal, &ModerateTesting) {
		t.Errorf("Failed to meet all the requirements. Expected Minimum: %v | Got: %v", ModerateTesting.Print(), testVal.Print())
		t.Errorf("Recommendation on how to fix this test: %v", testVal.Review.Feedback)
	}

	// Grade based on different testing levels
	grade := EvaluateGrade(testVal)

	// Print the results with verbose mode
	PrintResult(grade, testVal, false)
}
