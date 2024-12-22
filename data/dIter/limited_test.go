package dIter

import (
	"testing"
)

func TestLimited_Add(t *testing.T) {
	limited := NewLimited(func(item int) int { return item })

	limited.Add(1, 2, 3)

	if limited.Count() != 3 {
		t.Errorf("Expected count to be 3, got %d", limited.Count())
	}
}

func TestLimited_Next(t *testing.T) {
	limited := NewLimited(func(item int) int { return item })

	limited.Add(1, 2, 3)

	if !limited.Next() {
		t.Error("Expected Next() to return true, got false")
	}

	if limited.Item() != 1 {
		t.Errorf("Expected item to be 1, got %d", limited.Item())
	}

	limited.Next()
	limited.Next()

	if limited.Next() {
		t.Error("Expected Next() to return false, got true")
	}
}

func TestLimited_Reset(t *testing.T) {
	limited := NewLimited(func(item int) int { return item })

	limited.Add(1, 2, 3)
	limited.Next()
	limited.Next()

	limited.Reset()

	if limited.Next() && limited.Item() != 1 {
		t.Errorf("Expected item to be 1 after reset, got %d", limited.Item())
	}
}

func TestLimited_Clone(t *testing.T) {
	limited := NewLimited(func(item int) int { return item })

	limited.Add(1, 2, 3)
	limited.Next()

	clone := limited.Clone()

	if clone.Count() != 3 {
		t.Errorf("Expected clone count to be 3, got %d", clone.Count())
	}

	if clone.Next() && clone.Item() != 1 {
		t.Errorf("Expected clone item to be 1, got %d", clone.Item())
	}
}
