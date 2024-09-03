package rMongo

import (
	"github.com/yasseldg/go-simple/repos/rMongo/internal/filter"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/manager"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/pipeline"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rSort"
)

func New() Inter {
	return manager.New()
}

func NewSort() rSort.Sorts {
	return *sort.New()
}

func NewFilter() rFilter.Filters {
	return *filter.New()
}

func NewPipeline() pipeline.Inter {
	return pipeline.New()
}
