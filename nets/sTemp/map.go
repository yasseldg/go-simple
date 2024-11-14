package sTemp

import (
	"html/template"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Map map[string]*template.Template

func NewMap() Map {
	return make(map[string]*template.Template)
}

func (m Map) Add(name string, tmpl *template.Template) {
	if _, ok := m[name]; ok {
		sLog.Warn("Template already exists: %s", name)
	}
	m[name] = tmpl
}

func (m Map) Get(name string) *template.Template {
	if tmpl, ok := m[name]; ok {
		return tmpl
	}
	return nil
}
