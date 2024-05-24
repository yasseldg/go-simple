package fIter

import (
	"fmt"
	"os"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
)

//  Base

type Base struct {
	dIter.Inter

	file *os.File

	file_path string
	limit     int
}

func New(file_path string, limit int) *Base {
	return &Base{
		Inter:     dIter.New(),
		file_path: file_path,
		limit:     limit,
	}
}

func (i Base) File() *os.File {
	return i.file
}

func (iter *Base) OpenFile() error {
	file, err := os.Open(iter.file_path)
	if err != nil {
		return fmt.Errorf("os.Open( %s ): %s", iter.file_path, err)
	}

	sLog.Info("OpenFile: %s", iter.file_path)

	iter.file = file
	return nil
}

func (iter *Base) CloseFile() {
	sLog.Info("CloseFile: %s ", iter.file_path)
	iter.file.Close()
}

func (iter *Base) Limit() int {
	return iter.limit
}
