package sTemp

import (
	"html/template"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := NewMap()
	if m == nil {
		t.Errorf("Expected non-nil map, got nil")
	}
}

func TestMap_Add(t *testing.T) {
	m := NewMap()
	tmpl := template.New("test")
	m.Add("test", tmpl)
	if m.Get("test") != tmpl {
		t.Errorf("Expected template 'test', got %v", m.Get("test"))
	}
}

func TestMap_Get(t *testing.T) {
	m := NewMap()
	tmpl := template.New("test")
	m.Add("test", tmpl)
	if m.Get("test") != tmpl {
		t.Errorf("Expected template 'test', got %v", m.Get("test"))
	}
	if m.Get("nonexistent") != nil {
		t.Errorf("Expected nil, got %v", m.Get("nonexistent"))
	}
}
