package dAccu

import (
	"testing"
)

func TestBase(t *testing.T) {
	saveFunc := func() error {
		return nil
	}

	base := New(10, saveFunc)

	if base.Limit() != 10 {
		t.Errorf("Expected limit to be 10, got %d", base.Limit())
	}

	if base.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", base.Count())
	}

	base.Increase()
	if base.Count() != 1 {
		t.Errorf("Expected count to be 1, got %d", base.Count())
	}

	base.SetError(nil)
	if base.Error() != nil {
		t.Errorf("Expected error to be nil, got %v", base.Error())
	}

	base.Save()
	if !base.Empty() {
		t.Errorf("Expected base to be empty after save")
	}
}
