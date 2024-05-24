package fAccu

import (
	"fmt"
	"os"

	"github.com/yasseldg/go-simple/data/dAccu"
	"github.com/yasseldg/go-simple/files/sFile"
)

type Base struct {
	dAccu.Inter

	file_path string

	is_new bool
}

func New(file_path string, delete bool, limit int, save func() error) (*Base, error) {
	if delete {
		err := sFile.DeletePath(file_path)
		if err != nil {
			return nil, err
		}
	}

	b := &Base{
		file_path: file_path,
		is_new:    delete,
	}

	b.Inter = dAccu.New(limit, save)

	return b, nil
}

func (a *Base) FilePath() string {
	return a.file_path
}

func (a *Base) IsNew() bool {
	return a.is_new
}

func (a *Base) SetNew(b bool) {
	a.is_new = b
}

func (a *Base) OpenFile() (*os.File, error) {
	f, err := os.OpenFile(a.file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("os.OpenFile( %s ): %s", a.file_path, err)
	}
	return f, err
}
