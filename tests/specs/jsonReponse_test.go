package specs

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"std_go_boilerplate/internal/controller"
	"strings"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	// Create a simulated request recorder
	recorder := httptest.NewRecorder()

	// Create a test response
	data := map[string]interface{}{
		"message": "Hello, world!",
	}

	// Call the JsonResponse function
	controller.NewBaseController().JsonReponse(recorder, http.StatusOK, data)

	// Check HTTP status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code, got: %d, expected: %d", recorder.Code, http.StatusOK)
	}

	// Check content type
	contentType := recorder.Header().Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Errorf("Wrong content type, got: %s, expected: %s", contentType, expectedContentType)
	}

	// Check JSON response body
	expectedBody, _ := json.Marshal(data)
	expectedBodyString := string(expectedBody)

	// Trim leading and trailing white spaces and newlines
	expectedBodyString = strings.TrimSpace(expectedBodyString)
	recorderBodyString := strings.TrimSpace(recorder.Body.String())

	if recorderBodyString != expectedBodyString {
		t.Errorf("Incorrect response body, got: %s, expected: %s", recorderBodyString, expectedBodyString)
	}
}
