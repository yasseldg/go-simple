package dIter

import (
	"testing"
)

func TestNameConfig(t *testing.T) {
	base := New()
	nameConfig := NewNameConfig("TestName", base)

	if nameConfig.Name() != "TestName" {
		t.Errorf("Expected name to be 'TestName', got %s", nameConfig.Name())
	}

	if nameConfig.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", nameConfig.Count())
	}

	nameConfig.SetEmpty(true)
	if !nameConfig.Empty() {
		t.Errorf("Expected empty to be true")
	}

	nameConfig.SetError(nil)
	if nameConfig.Error() != nil {
		t.Errorf("Expected error to be nil")
	}
}
