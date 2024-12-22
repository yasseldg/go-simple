package dAccu

import (
	"testing"
	"github.com/yasseldg/go-simple/data/dTs"
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

	base := New(10, saveFunc)
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
	*Base
}

func (i *InterTsImpl) Add(inter dTs.Inter) {
	i.Increase()
}
