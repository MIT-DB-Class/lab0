package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a mock server
	srv := httptest.NewServer(http.HandlerFunc(HomeHandler))
	defer srv.Close()

	// Create an HTTP client
	client := &http.Client{}

	// Send a request to the mock server
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
}
