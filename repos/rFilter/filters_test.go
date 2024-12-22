package rFilter

import (
	"testing"
)

func TestNew(t *testing.T) {
	mockInterOper := &MockInterOper{}
	filters := New(mockInterOper)
	if filters == nil {
		t.Errorf("Expected Filters instance, got nil")
	}
}

func TestClone(t *testing.T) {
	mockInterOper := &MockInterOper{}
	filters := New(mockInterOper)
	clonedFilters := filters.Clone()
	if clonedFilters == nil {
		t.Errorf("Expected cloned Filters instance, got nil")
	}
}

func TestOper(t *testing.T) {
	mockInterOper := &MockInterOper{}
	filters := New(mockInterOper)
	if filters.Oper() == nil {
		t.Errorf("Expected InterOper instance, got nil")
	}
}

// MockInterOper is a mock implementation of the InterOper interface for testing purposes
type MockInterOper struct{}

func (m *MockInterOper) Append(field string, value interface{}) {}
func (m *MockInterOper) In(field string, values interface{})    {}
func (m *MockInterOper) Nin(field string, values interface{})   {}
func (m *MockInterOper) Gt(field string, value interface{})     {}
func (m *MockInterOper) Gte(field string, value interface{})    {}
func (m *MockInterOper) Lt(field string, value interface{})     {}
func (m *MockInterOper) Lte(field string, value interface{})    {}
func (m *MockInterOper) GtLt(field string, value1, value2 interface{}) {}
func (m *MockInterOper) GtLte(field string, value1, value2 interface{}) {}
func (m *MockInterOper) GteLt(field string, value1, value2 interface{}) {}
func (m *MockInterOper) GteLte(field string, value1, value2 interface{}) {}
func (m *MockInterOper) Clone_() InterOper { return &MockInterOper{} }
func (m *MockInterOper) Filters() []interface{} { return nil }
