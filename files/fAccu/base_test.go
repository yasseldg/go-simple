package fAccu

import (
	"os"
	"testing"

	"github.com/yasseldg/go-simple/data/dAccu"
)

func TestNew(t *testing.T) {
	saveFunc := func() error {
		return nil
	}

	filePath := "testfile.txt"
	defer os.Remove(filePath)

	base, err := New(filePath, true, 10, saveFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if base.FilePath() != filePath {
		t.Errorf("Expected file path to be %s, got %s", filePath, base.FilePath())
	}

	if !base.IsNew() {
		t.Errorf("Expected base to be new")
	}

	if base.Limit() != 10 {
		t.Errorf("Expected limit to be 10, got %d", base.Limit())
	}

	if base.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", base.Count())
	}
}

func TestBaseMethods(t *testing.T) {
	saveFunc := func() error {
		return nil
	}

	filePath := "testfile.txt"
	defer os.Remove(filePath)

	base, err := New(filePath, true, 10, saveFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	base.SetNew(false)
	if base.IsNew() {
		t.Errorf("Expected base to not be new")
	}

	f, err := base.OpenFile()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer f.Close()

	if f.Name() != filePath {
		t.Errorf("Expected file name to be %s, got %s", filePath, f.Name())
	}
}
