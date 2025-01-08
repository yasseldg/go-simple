package fCsv

import (
	"fmt"
	"io"
)

type IterFunc struct {
	Iter

	f func([]string) error
}

func NewIterFunc(file_path string, comma rune, f func([]string) error) (IterFunc, error) {

	iter, err := newIter(file_path, 0, comma)
	if err != nil {
		return IterFunc{}, fmt.Errorf("newIter: %s", err)
	}

	return IterFunc{
		Iter: iter,
		f:    f,
	}, nil
}

func (iter *IterFunc) Run() {
	if !iter.Iter.open() {
		return
	}

	for {
		line, err := iter.reader.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			iter.SetError(err)
			return
		}

		if err := iter.f(line); err != nil {
			iter.SetError(err)
			return
		}
	}
}
