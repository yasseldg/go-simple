package fJson

import (
	"os"
	"testing"
)

type mockData struct {
	Field1 string
	Field2 int
}

func TestNewAccu(t *testing.T) {
	filePath := "testfile.json"
	defer os.Remove(filePath)

	dataFunc := func() any {
		return &mockData{
			Field1: "test",
			Field2: 123,
		}
	}

	accu, err := NewAccu(filePath, true, 10, dataFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if accu.FilePath() != filePath {
		t.Errorf("Expected file path to be %s, got %s", filePath, accu.FilePath())
	}

	if !accu.IsNew() {
		t.Errorf("Expected accu to be new")
	}

	if accu.Limit() != 10 {
		t.Errorf("Expected limit to be 10, got %d", accu.Limit())
	}

	if accu.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", accu.Count())
	}
}

func TestAccuSave(t *testing.T) {
	filePath := "testfile.json"
	defer os.Remove(filePath)

	dataFunc := func() any {
		return &mockData{
			Field1: "test",
			Field2: 123,
		}
	}

	accu, err := NewAccu(filePath, true, 10, dataFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = accu.save()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if stat.Size() == 0 {
		t.Errorf("Expected file size to be greater than 0")
	}
}
