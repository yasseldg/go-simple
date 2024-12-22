package sTemp

import (
	"html/template"
	"testing"
)

func TestNewManager(t *testing.T) {
	manager := NewManager("templates", ".html")
	if manager.tmpl_dir != "templates" {
		t.Errorf("Expected tmpl_dir to be 'templates', got %s", manager.tmpl_dir)
	}
	if len(manager.exts) != 1 || manager.exts[".html"] == struct{}{} {
		t.Errorf("Expected exts to contain '.html', got %v", manager.exts)
	}
}

func TestManager_AddExts(t *testing.T) {
	manager := NewManager("templates")
	manager.AddExts(".html", ".tmpl")
	if len(manager.exts) != 2 || manager.exts[".html"] == struct{}{} || manager.exts[".tmpl"] == struct{}{} {
		t.Errorf("Expected exts to contain '.html' and '.tmpl', got %v", manager.exts)
	}
}

func TestManager_Load(t *testing.T) {
	manager := NewManager("templates", ".html")
	err := manager.Load()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(manager.Map) == 0 {
		t.Errorf("Expected templates to be loaded, got %v", manager.Map)
	}
}

func TestManager_Merge(t *testing.T) {
	manager := NewManager("templates", ".html")
	err := manager.Load()
	if err != nil {
		t.Fatalf("Failed to load templates: %v", err)
	}

	merged, err := manager.Merge("template1", "template2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if merged == nil {
		t.Errorf("Expected merged template, got nil")
	}
}

func TestManager_Get(t *testing.T) {
	manager := NewManager("templates", ".html")
	err := manager.Load()
	if err != nil {
		t.Fatalf("Failed to load templates: %v", err)
	}

	tmpl := manager.Get("template1")
	if tmpl == nil {
		t.Errorf("Expected template1 to be loaded, got nil")
	}
}

func TestManager_Add(t *testing.T) {
	manager := NewManager("templates", ".html")
	tmpl := template.New("template1")
	manager.Add("template1", tmpl)
	if manager.Get("template1") != tmpl {
		t.Errorf("Expected template1 to be added, got %v", manager.Get("template1"))
	}
}
