package fIter

import (
	"os"

	"github.com/yasseldg/go-simple/data/dIter"
)

type Inter interface {
	dIter.Inter

	File() *os.File
	Limit() int
	OpenFile() error
	CloseFile()
}
