package collection

import (
	"github.com/yasseldg/go-simple/repos/rMongo/internal/filter"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rSort"

	"go.mongodb.org/mongo-driver/mongo"
)

type Full struct {
	InterBase

	pipeline   mongo.Pipeline
	filter     rFilter.Inter
	sort       rSort.Inter
	projection rSort.Inter
	limit      int64
}

func NewFull(inter InterBase) *Full {
	return &Full{
		InterBase: inter,

		pipeline:   mongo.Pipeline{},
		filter:     filter.New(),
		sort:       sort.New(),
		projection: sort.New(),
	}
}

func (c *Full) Clone() *Full {
	full := NewFull(c.InterBase)
	full.Limit(c.limit)
	return full
}

func (c *Full) Pipeline(inter mongo.Pipeline) *Full {
	c.pipeline = inter
	return c
}

func (c *Full) Filters(inter rFilter.Inter) *Full {
	c.filter = inter
	return c
}

func (c *Full) Sorts(inter rSort.Inter) *Full {
	c.sort = inter
	return c
}

func (c *Full) Projections(inter rSort.Inter) *Full {
	c.projection = inter
	return c
}

func (c *Full) Limit(l int64) *Full {
	c.limit = l
	return c
}
