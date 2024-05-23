package fAccu

import (
	"os"

	"github.com/yasseldg/go-simple/data/dAccu"
)

type Inter interface {
	dAccu.Inter

	FilePath() string
	IsNew() bool
	SetNew(bool)

	OpenFile() (*os.File, error)
}
