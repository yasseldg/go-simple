package dIter

import (
	"testing"
)

func TestBase_String(t *testing.T) {
	base := New()
	expected := "Iter ( test ):"
	if result := base.String("test"); result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBase_Log(t *testing.T) {
	base := New()
	base.Log("test")
	// Check logs manually
}

func TestBase_SetError(t *testing.T) {
	base := New()
	err := fmt.Errorf("test error")
	base.SetError(err)
	if result := base.Error(); result != err {
		t.Errorf("Expected %v, got %v", err, result)
	}
}

func TestBase_SetEmpty(t *testing.T) {
	base := New()
	base.SetEmpty(true)
	if result := base.Empty(); !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestBase_Next(t *testing.T) {
	base := New()
	if result := base.Next(); !result {
		t.Errorf("Expected true, got %v", result)
	}

	base.SetEmpty(true)
	if result := base.Next(); result {
		t.Errorf("Expected false, got %v", result)
	}

	base.SetEmpty(false)
	base.SetError(fmt.Errorf("test error"))
	if result := base.Next(); result {
		t.Errorf("Expected false, got %v", result)
	}
}
