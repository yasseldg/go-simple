package fIter

import (
	"fmt"
	"os"

	"github.com/yasseldg/go-simple/logs/sLog"
)

//  Iter

type Iter struct {
	file *os.File

	file_path string
	limit     int

	empty bool
	err   error
}

func NewIter(file_path string, limit int) Iter {
	return Iter{file_path: file_path, limit: limit}
}

func (i Iter) Next() bool {
	if i.empty {
		return false
	}

	if i.err != nil {
		return false
	}

	return true
}

func (i *Iter) SetError(e error) {
	i.err = e
}

func (i Iter) Error() error {
	return i.err
}

func (i *Iter) SetEmpty(e bool) {
	i.empty = e
}

func (i Iter) Empty() bool {
	return i.empty
}

func (i Iter) File() *os.File {
	return i.file
}

func (iter *Iter) OpenFile() error {
	file, err := os.Open(iter.file_path)
	if err != nil {
		return fmt.Errorf("os.Open( %s ): %s", iter.file_path, err)
	}

	sLog.Info("OpenFile: %s", iter.file_path)

	iter.file = file
	return nil
}

func (iter *Iter) CloseFile() {
	sLog.Info("CloseFile: %s ", iter.file_path)
	iter.file.Close()
}

func (iter *Iter) Limit() int {
	return iter.limit
}
