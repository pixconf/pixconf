package xerror

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestErrorSingle(t *testing.T) {
	code := http.StatusNotFound
	message := http.StatusText(code)

	response := ErrorSingle(code, message)
	if response.Code != code {
		t.Errorf("Expected code %d, but got %d", code, response.Code)
	}

	if len(response.Errors) != 1 {
		t.Errorf("Expected 1 error message, but got %d", len(response.Errors))
	}

	if response.Errors[0].Message != message {
		t.Errorf("Expected error message '%s', but got '%s'", message, response.Errors[0].Message)
	}
}

func TestResponse_Marshal(t *testing.T) {
	code := http.StatusBadRequest
	message := http.StatusText(code)

	response := ErrorSingle(code, message)

	data, err := response.Marshal()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	var parsedResponse Response
	err = json.Unmarshal(data, &parsedResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %s", err)
	}

	if parsedResponse.Code != code {
		t.Errorf("Expected code %d, but got %d", code, parsedResponse.Code)
	}

	if len(parsedResponse.Errors) != 1 {
		t.Errorf("Expected 1 error message, but got %d", len(parsedResponse.Errors))
	}

	if parsedResponse.Errors[0].Message != message {
		t.Errorf("Expected error message '%s', but got '%s'", message, parsedResponse.Errors[0].Message)
	}
}
