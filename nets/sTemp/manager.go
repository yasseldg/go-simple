package sTemp

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/yasseldg/go-simple/files/sFile"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type manager struct {
	Map

	tmpl_dir string
	exts     map[string]struct{}
	errs     []error
}

func NewManager(tmpl_dir string, exts ...string) *manager {
	tmpls := &manager{
		tmpl_dir: tmpl_dir,
		exts:     make(map[string]struct{}),
		errs:     make([]error, 0),
		Map:      NewMap(),
	}

	tmpls.AddExts(exts...)

	return tmpls
}

func (t *manager) Log() {
	msg := "Templates loaded:"
	for name, tmpl := range t.Map {
		msg += fmt.Sprintf("%%l   ( %s )%v", name, tmpl.DefinedTemplates())
	}
	sLog.Info(sLog.Lines(msg))

	if len(t.errs) == 0 {
		return
	}

	errs := "Templates errors:"
	for _, err := range t.errs {
		errs += fmt.Sprintf("%%l   %v", err.Error())
	}
	sLog.Error(sLog.Lines(errs))
}

func (t *manager) Merge(names ...string) (merged *template.Template, err error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("the names of the templates have not been specified")
	}

	errs := make([]error, 0)

	for _, name := range names {
		if name == "" {
			continue
		}

		tmpl := t.Get(name)
		if tmpl == nil {
			errs = append(errs, fmt.Errorf("template not found: ( %s )", name))
			continue
		}

		if merged == nil {
			merged, err = tmpl.Clone()
			if err != nil {
				errs = append(errs, fmt.Errorf("error when cloning the template ( %s ): %v", name, err))
				merged = nil
			}
			continue
		}

		for _, sub_tmpl := range tmpl.Templates() {

			_, err := merged.AddParseTree(sub_tmpl.Name(), sub_tmpl.Tree)
			if err != nil {
				errs = append(errs, fmt.Errorf("error when adding the parse tree of the template ( %s ): %v", name, err))
			}
		}
	}

	if len(errs) > 0 {
		return merged, fmt.Errorf("error getting the templates: %v", errs)
	}
	return merged, nil
}

func (t *manager) AddExts(exts ...string) {
	for _, ext := range exts {
		t.exts[ext] = struct{}{}
	}
}

func (t *manager) Load() error {
	t.errs = make([]error, 0)
	for ext := range t.exts {
		t.load(ext)
		if len(t.errs) > 0 {
			return fmt.Errorf("error loading templates: %v", t.errs)
		}
	}
	return nil
}

// private methods

func (t *manager) load(ext string) {

	tmpls, errs := getTmpls(t.tmpl_dir, ext, nil)
	if len(errs) > 0 {
		t.errs = append(t.errs, errs...)
	}

	for name, tmpl := range tmpls {
		t.Map.Add(name, tmpl)
	}

	t.Log()
}

// private functions

func getTmpls(base_dir, ext string, tmpl *template.Template) (Map, []error) {
	dirs, err := sFile.GetDirsNames(base_dir)
	if err != nil {
		return nil, []error{fmt.Errorf("error getting directory names: %v", err)}
	}

	return getDirsTmpl(dirs, base_dir, ext, tmpl)
}

func getDirsTmpl(dirs []string, base_dir, ext string, tmpl *template.Template) (Map, []error) {
	tmpls := make(Map)
	errs := make([]error, 0)

	for _, dir_name := range dirs {

		pattern := filepath.Join(base_dir, dir_name, fmt.Sprintf("*%s", ext)) // .html

		files, err := filepath.Glob(pattern)
		if err != nil {
			errs = append(errs, fmt.Errorf("filepath.Glob( %s ): %v", pattern, err))
			continue
		}

		if len(files) == 0 {
			continue
		}

		f_tmpls, f_errs := getFilesTmpl(files, ext, tmpl)
		if len(f_errs) > 0 {
			errs = append(errs, f_errs...)
		}

		if len(f_tmpls) == 0 {
			continue
		}

		for name, f_tmpl := range f_tmpls {

			if tmpl != nil { // is_sub_directory
				tmpls.Add(dir_name, f_tmpl)
				continue
			}

			if name == dir_name {

				tmpl, err := f_tmpl.Clone()
				if err != nil {
					errs = append(errs, fmt.Errorf("error when cloning the template  ( %s ): %v", name, err))
					continue
				}

				sub_tmpls, sub_errs := getTmpls(filepath.Join(base_dir, dir_name), ext, tmpl)
				if len(sub_errs) > 0 {
					errs = append(errs, sub_errs...)
				}

				for sub_name, sub_tmpl := range sub_tmpls {
					map_name := fmt.Sprintf("%s_%s", name, sub_name)

					tmpls.Add(map_name, sub_tmpl)
				}
				continue
			}

			map_name := fmt.Sprintf("%s_%s", dir_name, name)

			tmpls.Add(map_name, f_tmpl)
		}
	}

	return tmpls, errs
}

func getFilesTmpl(files []string, ext string, tmpl *template.Template) (Map, []error) {
	tmpls := make(Map)
	errs := make([]error, 0)

	make_new := true
	if tmpl != nil { // is_sub_directory
		make_new = false

		c_tmpl, err := tmpl.Clone()
		if err != nil {
			errs = append(errs, fmt.Errorf("error when cloning the template ( %s ): %v", tmpl.Name(), err))
			return nil, errs
		}
		tmpl = c_tmpl
	}

	for _, file := range files {

		file_name, find := strings.CutSuffix(filepath.Base(file), ext)
		if !find {
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			errs = append(errs, fmt.Errorf("os.ReadFile( %s ): %v", file, err))
			continue
		}

		if make_new {
			tmpl = template.New(file_name)
		}

		_, err = tmpl.Parse(string(content))
		if err != nil {
			errs = append(errs, fmt.Errorf("error parsing the file ( %s ): %v", file, err))
			continue
		}

		tmpls.Add(tmpl.Name(), tmpl)
	}

	return tmpls, errs
}
