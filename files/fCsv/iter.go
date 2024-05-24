package fCsv

import (
	"encoding/csv"
	"io"

	"github.com/yasseldg/go-simple/files/fIter"
)

type InterIter interface {
	fIter.Inter
}

type Iter struct {
	fIter.Inter

	reader *csv.Reader
	comma  rune

	Item  []string
	items [][]string
}

func NewIter(file_path string, limit int, comma rune) (Iter, error) {
	return Iter{
		Inter: fIter.New(file_path, limit),
		comma: comma,
	}, nil
}

func (iter *Iter) Next() bool {
	if !iter.Inter.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.Item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	items, err := iter.readBatch()
	if err != nil {
		iter.SetError(err)
		return false
	}

	if len(items) == 0 {
		iter.Inter.SetEmpty(true)
		iter.CloseFile()
		return false
	}

	iter.items = items

	return iter.Next()
}

func (iter *Iter) readBatch() ([][]string, error) {
	if iter.reader == nil {
		err := iter.OpenFile()
		if err != nil {
			return nil, err
		}

		iter.reader = csv.NewReader(iter.File())
		iter.reader.Comma = iter.comma
	}

	var lines [][]string
	for i := 0; i < iter.Limit(); i++ {
		line, err := iter.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}
