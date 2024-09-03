package rIndex

import "github.com/yasseldg/go-simple/repos/rSort"

type Inter interface {
	Sort() rSort.Inter
	Unique() bool
}
