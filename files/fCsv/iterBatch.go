package fCsv

import (
	"fmt"
	"io"
)

type IterBatch struct {
	Iter

	Item  []string
	items [][]string
}

func NewIterBatch(file_path string, limit int, comma rune) (IterBatch, error) {

	iter, err := newIter(file_path, limit, comma)
	if err != nil {
		return IterBatch{}, fmt.Errorf("newIter: %s", err)
	}

	return IterBatch{
		Iter: iter}, nil
}

func (iter *IterBatch) Next() bool {
	if !iter.Inter.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.Item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	items := iter.read()
	if iter.Error() != nil {
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

func (iter *IterBatch) read() [][]string {
	if !iter.Iter.open() {
		return nil
	}

	var lines [][]string
	for i := 0; i < iter.Limit(); i++ {
		line, err := iter.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			iter.SetError(err)
			return nil
		}
		lines = append(lines, line)
	}
	return lines
}
