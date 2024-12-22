package fAccu

import (
	"os"
	"testing"

	"github.com/yasseldg/go-simple/data/dAccu"
)

type mockTs struct{}

func (m *mockTs) String(name string) string { return "" }
func (m *mockTs) Log(name string)           {}
func (m *mockTs) SetError(e error)          {}
func (m *mockTs) Error() error              { return nil }
func (m *mockTs) Increase()                 {}
func (m *mockTs) Limit() int                { return 0 }
func (m *mockTs) Count() int                { return 0 }
func (m *mockTs) Empty() bool               { return false }
func (m *mockTs) Save()                     {}

func TestInterTs(t *testing.T) {
	saveFunc := func() error {
		return nil
	}

	base := dAccu.New(10, saveFunc)
	interTs := &InterTsImpl{Base: base}

	mock := &mockTs{}
	interTs.Add(mock)

	if interTs.Count() != 1 {
		t.Errorf("Expected count to be 1, got %d", interTs.Count())
	}

	if interTs.Limit() != 10 {
		t.Errorf("Expected limit to be 10, got %d", interTs.Limit())
	}

	if interTs.Error() != nil {
		t.Errorf("Expected error to be nil, got %v", interTs.Error())
	}

	if interTs.Empty() {
		t.Errorf("Expected interTs to not be empty")
	}
}

type InterTsImpl struct {
	*dAccu.Base
}

func (i *InterTsImpl) Add(inter dAccu.Inter) {
	i.Increase()
}

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
