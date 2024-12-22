package fIter

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	filePath := "test_file.txt"
	limit := 10
	iter := New(filePath, limit)

	if iter.file_path != filePath {
		t.Errorf("Expected file_path %s, got %s", filePath, iter.file_path)
	}

	if iter.limit != limit {
		t.Errorf("Expected limit %d, got %d", limit, iter.limit)
	}
}

func TestOpenFile(t *testing.T) {
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
		t.Errorf("Expected no error, got %s", err)
	}

	if iter.file == nil {
		t.Errorf("Expected file to be opened, got nil")
	}
}

func TestCloseFile(t *testing.T) {
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

	iter.CloseFile()

	if iter.file != nil {
		t.Errorf("Expected file to be closed, got %v", iter.file)
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
