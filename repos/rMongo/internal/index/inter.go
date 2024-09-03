package index

import "github.com/yasseldg/go-simple/repositorys/rSort"

type Inter interface {
	Sort() rSort.Inter
	Unique() bool
}
