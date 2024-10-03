package client

import (
	"encoding/json"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"github.com/henrylamb/object-generation-golang/testingLite"
	"os"
	"testing"
)

func TestToolReq(t *testing.T) {
	// Create a temporary directory
	c := NewClient(os.Getenv("MULTIPLE_PASSWORD"), "http://localhost:2008")

	//construct single test
	definition, code, err := testingLite.SingleUnitTestWrapper(testingLite.WorkingAssumption, "./toolReq.go", jsonSchema.Gpt4)
	if err != nil {
		t.Errorf("Error constructing test: %v", err)
	}

	response, err := c.SendRequest(code, definition)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	//unmarshal the response
	testVal := &testingLite.CodeTest{}
	err = json.Unmarshal(response.Data, testVal)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	//Compare the values
	if !testingLite.TestComparison(testVal, &testingLite.ModerateTesting) {
		t.Errorf("Failed to meet all the requirements. Expected Minimum: %v | Got: %v", testingLite.ModerateTesting.Print(), testVal.Print())
		t.Errorf("Recommendation on how to fix this test: %v", testVal.Review.Feedback)
	}
}
