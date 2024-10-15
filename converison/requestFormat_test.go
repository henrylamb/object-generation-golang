package converison

import (
	"encoding/json"
	"github.com/henrylamb/object-generation-golang/client"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"github.com/henrylamb/object-generation-golang/testingLite"
	"log"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	// Create a temporary directory
	c := client.NewDefaultClient(os.Getenv("MULTIPLE_PASSWORD"), "http://localhost:2008")

	processor := testingLite.NewDefaultCodeProcessor()
	codePath := "./requestFormat.go"

	definition, code, err := testingLite.TestConstructor(testingLite.WorkingAssumption, codePath, jsonSchema.Gpt4, processor)
	if err != nil {
		t.Errorf("Error constructing test: %v", err)
	}

	response, err := c.SendRequest(code, definition)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Marshal res (map[string]any) into JSON bytes
	bytes, err := json.Marshal(response.Data)
	if err != nil {
		log.Println("Failed to marshal res:", err)
		t.Errorf("Error marshalling response: %v", err)
	}

	//unmarshal the response
	testVal := &testingLite.CodeTest{}
	err = json.Unmarshal(bytes, testVal)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	//Compare the values
	if !testingLite.TestComparison(testVal, &testingLite.ModerateTesting) {
		t.Errorf("Failed to meet all the requirements. Expected Minimum: %v | Got: %v", testingLite.ModerateTesting.Print(), testVal.Print())
		t.Errorf("Recommendation on how to fix this test: %v", testVal.Review.Feedback)
	}
	// Grade based on different testing levels
	grade := testingLite.EvaluateGrade(testVal)

	// Print the results with verbose mode
	testingLite.PrintResult(grade, testVal, false)
}
