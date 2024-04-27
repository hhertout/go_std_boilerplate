package specs

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"std_go_boilerplate/internal/controller"
	"testing"
)

type TestStruct struct {
	Message string `json:"message"`
}

func TestBindJsonBody_Success(t *testing.T) {
	// Create a test JSON payload
	payload := TestStruct{Message: "Hello, world!"}
	payloadBytes, _ := json.Marshal(payload)

	// Create a test request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	controller := &controller.BaseController{}
	var decodedPayload TestStruct

	err = controller.BindJsonBody(req, &decodedPayload)
	if err != nil {
		t.Fatalf("BindJsonBody failed: %v", err)
	}

	// Check if the decoded payload matches the original payload
	if decodedPayload.Message != payload.Message {
		t.Errorf("Decoded payload does not match, got: %s, expected: %s", decodedPayload.Message, payload.Message)
	}
}

func TestBindJsonBody_EmptyBody(t *testing.T) {
	// Create an empty request
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	controller := &controller.BaseController{}
	var decodedPayload TestStruct
	err = controller.BindJsonBody(req, &decodedPayload)

	if err != nil {
		t.Fatal("Expected an error but got nil")
	}

	// Check if the decoded payload matches the original payload
	if decodedPayload.Message != "" {
		t.Errorf("Decoded payload does not match, got: %s, expected an empty string", decodedPayload.Message)
	}
}

func TestBindJsonBody_InvalidJson(t *testing.T) {
	// Create a test request with invalid JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte("{")))
	if err != nil {
		t.Fatal(err)
	}

	controller := &controller.BaseController{}
	var decodedPayload TestStruct
	err = controller.BindJsonBody(req, &decodedPayload)

	if err == nil {
		t.Fatal("Expected an error but got nil")
	}

	// Check if the error is the expected one
	expectedErr := errors.New("unexpected EOF")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Unexpected error, got: %v, expected: %v", err, expectedErr)
	}
}
