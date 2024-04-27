package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"std_go_boilerplate/internal/controller"
	"testing"
)

func TestIsAlive(t *testing.T) {
	c := controller.NewBaseController()
	server := httptest.NewServer(http.HandlerFunc(c.HealthCheck))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	var body struct {
		Message string `json:"message"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		t.Fatalf("error pasring the body. Err: %v", err)
	}
	if body.Message != "I'm alive" {
		t.Errorf("expected response body to be %v; got %v", "I'm alive", body.Message)
	}
}
