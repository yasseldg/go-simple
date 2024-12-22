package sNet

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(nil, nil)
	if client.httpClient != http.DefaultClient {
		t.Errorf("Expected default http client, got %v", client.httpClient)
	}
	if client.do == nil {
		t.Errorf("Expected default do function, got nil")
	}
}

func TestClient_Do(t *testing.T) {
	client := NewClient(nil, nil)
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
