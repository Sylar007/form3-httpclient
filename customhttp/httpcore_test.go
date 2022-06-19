package customhttp

import (
	"testing"
)

func TestGetRequestBody_WhenBodyNil_ReturnError(t *testing.T) {
	// Initialization:
	client := httpClient{}

	// Execution
	body, err := client.getRequestBody(nil)

	// Validation
	if err != nil {
		t.Error("no error expected when passing a nil body")
	}

	if body != nil {
		t.Error("no body expected when passing a nil body")
	}
}

func TestGetRequestBody_WhenBodyWithJsonData_ReturnPass(t *testing.T) {
	// Initialization:
	client := httpClient{}

	// Execution
	requestBody := []string{"item1", "item2"}

	body, err := client.getRequestBody(requestBody)

	// Validation
	if err != nil {
		t.Error("no error expected when marshaling slice as json")
	}

	if string(body) != `["item1","item2"]` {
		t.Error("invalid json body obtained")
	}
}
