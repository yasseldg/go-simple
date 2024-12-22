package dIter

import (
	"testing"
)

func TestInter_String(t *testing.T) {
	var inter Inter = New()
	expected := "Iter ( test ):"
	if result := inter.String("test"); result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestInter_Log(t *testing.T) {
	var inter Inter = New()
	inter.Log("test")
	// Check logs manually
}

func TestInter_SetError(t *testing.T) {
	var inter Inter = New()
	err := fmt.Errorf("test error")
	inter.SetError(err)
	if result := inter.Error(); result != err {
		t.Errorf("Expected %v, got %v", err, result)
	}
}

func TestInter_SetEmpty(t *testing.T) {
	var inter Inter = New()
	inter.SetEmpty(true)
	if result := inter.Empty(); !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestInter_Next(t *testing.T) {
	var inter Inter = New()
	if result := inter.Next(); !result {
		t.Errorf("Expected true, got %v", result)
	}

	inter.SetEmpty(true)
	if result := inter.Next(); result {
		t.Errorf("Expected false, got %v", result)
	}

	inter.SetEmpty(false)
	inter.SetError(fmt.Errorf("test error"))
	if result := inter.Next(); result {
		t.Errorf("Expected false, got %v", result)
	}
}
