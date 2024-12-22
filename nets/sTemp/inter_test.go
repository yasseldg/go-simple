package sTemp

import (
	"html/template"
	"testing"
)

type mockManager struct {
	templates map[string]*template.Template
}

func (m *mockManager) Get(name string) *template.Template {
	return m.templates[name]
}

func (m *mockManager) Add(name string, tmpl *template.Template) {
	m.templates[name] = tmpl
}

func (m *mockManager) Log() {}

func (m *mockManager) Merge(names ...string) (merged *template.Template, err error) {
	return nil, nil
}

func (m *mockManager) Load() error {
	return nil
}

func (m *mockManager) AddExts(exts ...string) {}

func TestInterManager(t *testing.T) {
	manager := &mockManager{templates: make(map[string]*template.Template)}

	tmpl1 := template.New("tmpl1")
	tmpl2 := template.New("tmpl2")

	manager.Add("tmpl1", tmpl1)
	manager.Add("tmpl2", tmpl2)

	if manager.Get("tmpl1") != tmpl1 {
		t.Errorf("Expected tmpl1, got %v", manager.Get("tmpl1"))
	}

	if manager.Get("tmpl2") != tmpl2 {
		t.Errorf("Expected tmpl2, got %v", manager.Get("tmpl2"))
	}
}
