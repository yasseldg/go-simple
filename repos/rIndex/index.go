package rIndex

import (
	"github.com/yasseldg/go-simple/repos/rSort"
)

type Index struct {
	sort   rSort.Inter
	unique bool
}
type Indexes []Inter

func New(sort rSort.Inter, unique bool) *Index {
	return &Index{
		sort:   sort,
		unique: unique,
	}
}

func (i *Index) Sort() rSort.Inter {
	return i.sort
}

func (i *Index) Unique() bool {
	return i.unique
}
