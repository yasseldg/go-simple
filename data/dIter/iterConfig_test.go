package dIter

import (
	"testing"
)

func TestIterConfig(t *testing.T) {
	config := NewIterConfig("TestConfig")

	if config.Name() != "TestConfig" {
		t.Errorf("Expected name to be 'TestConfig', got %s", config.Name())
	}

	if config.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", config.Count())
	}

	config.Add(NewNameConfig("Config1", NewIterConfig("SubConfig1")))
	config.Add(NewNameConfig("Config2", NewIterConfig("SubConfig2")))

	if config.Count() != 1 {
		t.Errorf("Expected count to be 1, got %d", config.Count())
	}

	config.Reset()
	if config.Next() {
		t.Errorf("Expected Next() to return false, got true")
	}
}
