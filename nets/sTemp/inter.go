package sTemp

import "html/template"

type InterManager interface {
	InterMap

	Log()

	Merge(names ...string) (merged *template.Template, err error)

	Load() error
	AddExts(exts ...string)
}

type InterMap interface {
	Get(name string) *template.Template
	Add(name string, tmpl *template.Template)
}
