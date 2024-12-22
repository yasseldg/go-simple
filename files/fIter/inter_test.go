package fIter

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	filePath := "test_file.txt"
	limit := 10
	iter := New(filePath, limit)

	// Create a test file
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %s", err)
	}
	file.Close()
	defer os.Remove(filePath)

	err = iter.OpenFile()
	if err != nil {
		t.Fatalf("Failed to open test file: %s", err)
	}

	if iter.File() == nil {
		t.Errorf("Expected file to be opened, got nil")
	}

	iter.CloseFile()

	if iter.File() != nil {
		t.Errorf("Expected file to be closed, got %v", iter.File())
	}
}

func TestLimit(t *testing.T) {
	filePath := "test_file.txt"
	limit := 10
	iter := New(filePath, limit)

	if iter.Limit() != limit {
		t.Errorf("Expected limit %d, got %d", limit, iter.Limit())
	}
}
