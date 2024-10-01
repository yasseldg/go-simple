package fCsv

import (
	"encoding/csv"
	"fmt"

	"github.com/yasseldg/go-simple/files/fIter"
)

type InterIter interface {
	fIter.Inter
}

type Iter struct {
	fIter.Inter

	reader *csv.Reader
	comma  rune
}

func newIter(file_path string, limit int, comma rune) (Iter, error) {
	return Iter{
		Inter: fIter.New(file_path, limit),
		comma: comma,
	}, nil
}

func (iter *Iter) open() bool {
	if iter.reader != nil {
		return true
	}

	err := iter.OpenFile()
	if err != nil {
		iter.SetError(fmt.Errorf("OpenFile: %s", err))
		return false
	}

	iter.reader = csv.NewReader(iter.File())
	iter.reader.Comma = iter.comma

	return true
}
