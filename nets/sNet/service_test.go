package sNet

import (
	"testing"
)

func TestNewService(t *testing.T) {
	service, err := NewService("test_env", "")
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	if service.env != "test_env" {
		t.Errorf("Expected env to be 'test_env', got %s", service.env)
	}
}

func TestService_GetUrl(t *testing.T) {
	service, err := NewService("test_env", "")
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	expectedUrl := "http://localhost:8080"
	if service.GetUrl() != expectedUrl {
		t.Errorf("Expected URL to be %s, got %s", expectedUrl, service.GetUrl())
	}
}

func TestService_SendObj(t *testing.T) {
	service, err := NewService("test_env", "")
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	err = service.SendObj("/test", map[string]string{"key": "value"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
