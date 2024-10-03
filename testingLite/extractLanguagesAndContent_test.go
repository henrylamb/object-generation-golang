package testingLite

import (
	"encoding/json"
	"github.com/henrylamb/object-generation-golang/client"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"os"
	"testing"
)

func TestExtractLanguagesAndContents(t *testing.T) {
	// Create a temporary directory
	c := client.NewClient(os.Getenv("MULTIPLE_PASSWORD"), "http://localhost:2008")

	//construct single test
	definition, code, err := SingleUnitTestWrapper(WorkingAssumption, "./extractLanguagesAndContent.go", jsonSchema.Gpt3)
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
		t.Errorf("Failed to meet all the requirements. Expected Minimum: %v | Got: %v", ModerateTesting, *testVal)
		t.Errorf("Recommendation on how to fix this test: %v", testVal.Review.Feedback)
	}
}
